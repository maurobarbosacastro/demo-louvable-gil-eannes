package camunda_processes

import (
	"context"
	"fmt"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/internal/service/redis"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
	"github.com/google/uuid"
)

const (
	EndInvestment = "end-investment"
	NotifyTagpeak = "notify-tagpeak-actor"
	UpdateReward  = "update-reward-status"
	CreditUser    = "credit-value-to-user"
)

// *** FLOW ***
// Process starts in  CheckInvestment or in Stop Reward
// Next Notify the tagpeak actor
// If Reward is finished, sends the reward to the last step (Update User Balance)
// If Reward is Expired or  Stopped waits for the manual task to validate the reward
// After validated updates the reward status to FINISHED
// Updates the user balance

func EndInvestmentProcess(reward map[string]interface{}) {
	//START CAMUNDA PROCESS
	camundaPKG.StartProcessInstance(camundaPKG.InjectEnvOnKey(EndInvestment), *camundaPKG.GetCamundaClient(), reward)
}

func HandleNotifyTagPeak() {
	fmt.Println("***CAMUNDA*** Start Notify TagPeak Admin")
	w, _ := camundaPKG.StartDynamicWorker(NotifyTagpeak, func(client worker.JobClient, job entities.Job) {

		//TODO: Implement notifications
		//TODO: Send email to Admin with last day resume (Maybe change this one to the CheckInvestments job)

		endInvestmentVariables, err := camundaPKG.GetVariables(job)

		rewardUUID, err := uuid.Parse(endInvestmentVariables["rewardUuid"].(string))
		if err != nil {
			fmt.Println(err)
		}

		reward, err := service.GetReward(rewardUUID)
		if err != nil {
			fmt.Println("Error getting reward", err)
		}

		//Check if the user have referral code
		referredByWho, err := service.ReferredByWho(utils.ParseIDToUUID(reward.User))
		if err != nil {
			fmt.Println("Error getting referredByWho: ", err)
			return
		}

		//Variables to continue the process with the referral
		//rewardUUID is already in the process
		rewards := map[string]interface{}{
			"referredByWho":   referredByWho,
			"hasBeenReferred": referredByWho != nil,
			"investmentState": reward.State,
			"rewardOwner":     reward.User,
			"userUuid":        reward.User,
			"transactionUuid": reward.TransactionUUID,
		}

		//Only save if Status is Expired or Stopped
		if reward.State == "STOPPED" || reward.State == "EXPIRED" {
			//Save process
			camundaProcessDto := dto.CreateCamundaProcessDto{
				FieldUUID:          rewardUUID,
				ProcessInstanceKey: job.ProcessInstanceKey,
				ProcessId:          EndInvestment,
			}

			// Saved process for manual task
			_, err = service.CreateCamundaProcess(camundaProcessDto)
			if err != nil {
				fmt.Println("Error creating camunda process: ", err)
			}

		}

		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(rewards)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = request.Send(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}

	})

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Notify TagPeak Admin worker...")
		w.Close()
	}()
}

func HandleUpdateReward() {
	fmt.Println("***CAMUNDA*** Start Update Reward")
	w, _ := camundaPKG.StartDynamicWorker(UpdateReward, func(client worker.JobClient, job entities.Job) {

		//Get Rewards uuid
		endInvestmentVariables, err := camundaPKG.GetVariables(job)

		rewardUUID, err := uuid.Parse(endInvestmentVariables["rewardUuid"].(string))
		if err != nil {
			fmt.Println(err)
			return
		}
		rewardOwner, err := uuid.Parse(endInvestmentVariables["rewardOwner"].(string))
		if err != nil {
			fmt.Println(err)
			return
		}

		err = service.EditStateReward(rewardUUID, "FINISHED", "camunda")
		if err != nil {
			fmt.Println("Error putting the reward in FINISHED state: ", err)
			return
		}

		//Check if the user have referral code
		referredByWho, err := service.ReferredByWho(rewardOwner)
		if err != nil {
			fmt.Println("Error getting referredByWho: ", err)
			return
		}

		//Get transaction for email
		transaction, err := service.GetTransactionByReward(rewardUUID)
		if err != nil {
			fmt.Printf("Error getting transaction value: %v", err)
			return
		}

		//Variables to continue the process with the referral
		//rewardUUID is already in the process
		rewards := map[string]interface{}{
			"referredByWho":   referredByWho,
			"hasBeenReferred": referredByWho != nil,
			"transactionUuid": transaction.Uuid,
		}

		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(rewards)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = request.Send(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}

	})

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Update Reward worker...")
		w.Close()
	}()
}

func HandleCreditUser() {
	fmt.Println("***CAMUNDA*** Start Credit User")
	w, _ := camundaPKG.StartDynamicWorker(CreditUser, func(client worker.JobClient, job entities.Job) {

		//Get Keycloak token
		getKeycloakAdminToken(auth.KeycloakInstance)

		//Get rewards balance for user by uuid
		rewardVariables, err := camundaPKG.GetVariables(job)
		rewardUUID, err := uuid.Parse(rewardVariables["rewardUuid"].(string))
		if err != nil {
			fmt.Printf("Error parsing rewardUuid: %v", err)
			return
		}

		reward, err := service.GetReward(rewardUUID)
		if err != nil {
			fmt.Printf("Error getting reward value: %v", err)
			return
		}

		//update user balance
		err = redis.HandleKeycloakBalanceUpdate(reward.User, reward.CurrentRewardUser, auth.KeycloakInstance)
		if err != nil {
			fmt.Printf("Error updating balance: %v", err)
			return
		}

		_, err = client.NewCompleteJobCommand().JobKey(job.GetKey()).Send(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}

	})

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Credit User worker...")
		w.Close()
	}()
}

func getKeycloakAdminToken(keycloak *constants.Keycloak) {
	logster.StartFuncLog()

	now := time.Now()

	if now.After(keycloak.TokenExpireDate) {
		logster.Warn("Token expired, getting new one")
		token, err := keycloak.Client.LoginAdmin(
			keycloak.Ctx,
			dotenv.GetEnv("TAGPEAK_ADMIN_USERNAME"),
			dotenv.GetEnv("TAGPEAK_ADMIN_PASSWORD"),
			"master",
		)

		if err != nil {
			logster.Error(err, "Error getting Keycloak admin token")
		}

		keycloak.AdminToken = token
		keycloak.TokenExpireDate = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}

	logster.EndFuncLog()
}
