package camunda_processes

import (
	"context"
	"fmt"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/internal/service/redis"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
	"github.com/google/uuid"
)

const (
	CreateNewRewardReferral = "create-new-reward-referral"
	CreditUserReferral      = "credit-user-referral"
)

func HandleCreateNewRewardReferral() {
	fmt.Println("***CAMUNDA*** Start Credit User")
	w, _ := camundaPKG.StartDynamicWorker(CreateNewRewardReferral, func(client worker.JobClient, job entities.Job) {

		//Get Admin Token
		getKeycloakAdminToken(auth.KeycloakInstance)

		//Get rewards balance for user by uuid
		rewardVariables, err := camundaPKG.GetVariables(job)
		rewardUUID, err := uuid.Parse(rewardVariables["rewardUuid"].(string))
		if err != nil {
			fmt.Println("Error parsing rewardUuid: ", err)
			return
		}

		reward, err := service.GetReward(rewardUUID)
		if err != nil {
			fmt.Println("Error getting reward value: ", err)
			return
		}

		newRewardOwner, err := uuid.Parse(rewardVariables["referredByWho"].(string))
		if err != nil {
			fmt.Println("Error parsing userUUID: ", err)
			return
		}

		user, err := service.GetUserById(newRewardOwner.String(), auth.KeycloakInstance)
		if err != nil {
			fmt.Println("Error getting user: ", err)
			return
		}

		var currentReward float64

		//Knowing that the percentage is setting was 5 10 and not 0.05 ou 0.1
		// User have Reward Percentage the initial price is the value multiplicated by the percentage.
		if user.RewardPercentage != nil {
			currentReward = reward.CurrentRewardUser * (*user.RewardPercentage / 100)
		} else {
			// Get the membership, convert the percentage to float and multiply the value to get the initial price
			membership := service.PercentageBaseOnMembership(service.GetMembershipLevel(user.Groups))
			currentReward = reward.CurrentRewardUser * (float64(*membership.PercentageOnReward) / 100)
		}

		//Check if user has reward percentage
		// If not, use the default percentage groups

		title := "Reward by referral"
		newReward := dto.CreateRewardDTO{
			TransactionUUID:   reward.TransactionUUID,
			Isin:              reward.Isin,
			CurrencySource:    reward.CurrencyUser, //We need to use the original reward CurrencyUser because we calculate the reward based on CurrentRewardUser and not the source.
			State:             "FINISHED",
			CurrentRewardUser: &currentReward,
			EndDate:           reward.EndDate,
			Type:              "FIXED",
			Title:             &title,
			Origin:            "REFERRAL",
		}

		createReward, err := service.CreateRewardFromReferral(newReward, *user)
		if err != nil {
			fmt.Println("Error creating reward from referral: ", err)
			return
		}

		referrer, err := service.GetReferralByInvitee(utils.ParseIDToUUID(reward.User))
		if err != nil {
			fmt.Println("Error getting referral by invitee: ", err)
			return
		}

		var createDto = dto.CreateReferralRevenueDTO{
			Amount:          createReward.CurrentRewardUser,
			RewardUUID:      &createReward.Uuid,
			TransactionUUID: &createReward.TransactionUUID,
			ReferralUUID:    referrer.Uuid,
			CreatedBy:       "camunda",
		}

		revenueModel := utils.CreateReferralRevenueDto(createDto)

		// Create the revenue
		_, err = repository.CreateReferralRevenue(revenueModel)
		if err != nil {
			fmt.Println(err)
			return
		}

		rewards := map[string]interface{}{
			"newRewardUuid": createReward.Uuid,
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
		fmt.Println("***CAMUNDA*** Shutting down Credit User worker...")
		w.Close()
	}()
}

func HandleCreditUserReferral() {
	fmt.Println("***CAMUNDA*** Start Credit User")
	w, _ := camundaPKG.StartDynamicWorker(CreditUserReferral, func(client worker.JobClient, job entities.Job) {

		//Get Admin Token
		getKeycloakAdminToken(auth.KeycloakInstance)

		//Get rewards balance for user by uuid
		rewardVariables, err := camundaPKG.GetVariables(job)
		rewardUUID, err := uuid.Parse(rewardVariables["newRewardUuid"].(string))
		if err != nil {
			fmt.Println("Error parsing rewardUuid: ", err)
			return
		}

		reward, err := service.GetReward(rewardUUID)
		if err != nil {
			fmt.Println("Error getting reward value: ", err)
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
