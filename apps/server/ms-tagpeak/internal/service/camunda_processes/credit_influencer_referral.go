package camunda_processes

import (
	"context"
	"fmt"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/service"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

const (
	GetDataCreditInfluencer   = "get-data-credit-influencer"
	CreditNewInfluencerReward = "create-new-reward-influencer-referral"
)

func HandleGetDataCreditInfluencer() {
	fmt.Println("***CAMUNDA*** Start Get Data Credit Influencer")

	w, _ := camundaPKG.StartDynamicWorker(GetDataCreditInfluencer, func(client worker.JobClient, job entities.Job) {
		logster.StartFuncLogMsg(job.ProcessInstanceKey)

		//Get Admin Token
		getKeycloakAdminToken(auth.KeycloakInstance)

		//Get rewards balance for user by uuid
		variables, _ := camundaPKG.GetVariables(job)
		userUuid, _ := variables["userUuid"].(string)
		transactionUuid, _ := variables["transactionUuid"].(string)

		referral, errGetReferral := service.GetReferralByInvitee(uuid.MustParse(userUuid))
		if errGetReferral != nil || referral.Uuid.String() == "00000000-0000-0000-0000-000000000000" {
			logster.Error(errGetReferral, "Error getting referral by invitee")
			variablesToSend := map[string]interface{}{
				"hasBeenReferred": false,
			}
			camundaPKG.EndProcess(context.Background(), client, job.GetKey(), &variablesToSend)
			return
		}

		referrerUuid := referral.ReferrerUUID
		user, err := service.GetUserById(referrerUuid.String(), auth.KeycloakInstance)
		if err != nil {
			logster.Error(err, "Error getting user by referrer uuid")
			variablesToSend := map[string]interface{}{
				"hasBeenReferred": false,
			}
			camundaPKG.EndProcess(context.Background(), client, job.GetKey(), &variablesToSend)
			return

		}

		_, found := lo.Find(user.Groups, func(item string) bool { return item == *service.Configuration.MembershipLevels.Influencer })

		logster.Info(fmt.Sprintf("Found: %v", found))
		if !found {
			logster.Warn(fmt.Sprintf("User referrer of %s is not an influencer", userUuid))
			variablesToSend := map[string]interface{}{
				"hasBeenReferred": false,
			}
			camundaPKG.EndProcess(context.Background(), client, job.GetKey(), &variablesToSend)
			return
		}

		variablesToSend := map[string]interface{}{
			"hasBeenReferred": true,
			"transactionUuid": transactionUuid,
			"transactionUser": userUuid,
			"referrerUuid":    referrerUuid,
		}

		logster.EndFuncLogMsg(fmt.Sprintf("Variables: %+v", variablesToSend))
		camundaPKG.EndProcess(context.Background(), client, job.GetKey(), &variablesToSend)
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

func HandleCreditNewInfluencerReward() {
	fmt.Println("***CAMUNDA*** Start Credit Influencer")
	w, _ := camundaPKG.StartDynamicWorker(CreditNewInfluencerReward, func(client worker.JobClient, job entities.Job) {
		logster.StartFuncLogMsg(job.ProcessInstanceKey)

		//Get Admin Token
		getKeycloakAdminToken(auth.KeycloakInstance)

		//Get rewards balance for user by uuid
		variables, _ := camundaPKG.GetVariables(job)
		transactionUuid, _ := variables["transactionUuid"].(string)
		//userUuid, _ := variables["transactionUser"].(string)
		referrerUuid, _ := variables["referrerUuid"].(string)

		logster.Info(fmt.Sprintf("Variables: %+v", variables))

		rewardCreated, errorCreation := service.CreateRewardForInfluencer(
			referrerUuid,
			transactionUuid,
			auth.KeycloakInstance,
		)

		if errorCreation != nil {
			logster.Error(errorCreation, "Error creating reward for influencer")
			camundaPKG.EndProcess(context.Background(), client, job.GetKey(), nil)
		}

		logster.Info(fmt.Sprintf("Reward created: %s", rewardCreated.Uuid.String()))

		//Create reward history revenue
		var createDto = dto.CreateReferralRevenueDTO{
			Amount:          rewardCreated.CurrentRewardUser,
			RewardUUID:      &rewardCreated.Uuid,
			TransactionUUID: &rewardCreated.TransactionUUID,
			ReferralUUID:    rewardCreated.Uuid,
			CreatedBy:       "camunda",
		}

		revenueModel := utils.CreateReferralRevenueDto(createDto)
		// Create the revenue
		_, err := repository.CreateReferralRevenue(revenueModel)
		if err != nil {
			logster.Error(err, "Error creating revenue")
		}

		variableToSend := map[string]interface{}{
			"newRewardUuid":  rewardCreated.Uuid.String(),
			"userUuid":       rewardCreated.User,
			"amountToCredit": rewardCreated.CurrentRewardUser,
		}

		camundaPKG.EndProcess(context.Background(), client, job.GetKey(), &variableToSend)
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
