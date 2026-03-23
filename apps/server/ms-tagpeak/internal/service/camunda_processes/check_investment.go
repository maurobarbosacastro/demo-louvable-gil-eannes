package camunda_processes

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-tagpeak/internal/service"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"os"
	"os/signal"
	"syscall"

	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
)

const (
	CheckExpiredRewards  = "check-expired-rewards"
	UpdateExpiredRewards = "update-expired-rewards"
)

// *** FLOW ***
// Starts in HandleCheckTransactions
// If there are expired rewards, sends them to the UpdateExpiredRewards process
// If there are no expired rewards, sends an empty array and the process ends
// After updating the rewards, the process starts the End Investment process for each expired reward

func HandleCheckInvestments() {
	fmt.Println("***CAMUNDA*** Start Check Expired Rewards")
	w, _ := camundaPKG.StartDynamicWorker(CheckExpiredRewards, func(client worker.JobClient, job entities.Job) {

		expiredRewards, err := service.GetExpiredRewards()
		if err != nil {
			fmt.Println("***CAMUNDA*** Error getting expired rewards")
			return
		}

		rewardsToSend := map[string]interface{}{}

		if len(*expiredRewards) == 0 {
			//Always add an empty string
			rewardsToSend["rewards"] = []string{}
			fmt.Println("***CAMUNDA*** No expired rewards")
		} else {
			rewardsToSend["rewards"] = expiredRewards
			fmt.Println("***CAMUNDA*** Expired rewards found")
		}

		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(rewardsToSend)
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
		fmt.Println("***CAMUNDA*** Shutting down Check Transactions worker...")
		w.Close()
	}()
}

func HandleUpdateInvestment() {
	fmt.Println("***CAMUNDA*** Start Update Expired Rewards")
	w, _ := camundaPKG.StartDynamicWorker(UpdateExpiredRewards, func(client worker.JobClient, job entities.Job) {

		isExpiredOn, err := service.GetConfiguration("expired")
		if err != nil {
			fmt.Println("***CAMUNDA*** Error getting expired configuration")
			return
		}

		variables := map[string]interface{}{}

		expiredRewards := job.GetVariables()

		err = json.Unmarshal([]byte(expiredRewards), &variables)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Convert the interface{} to a string slice
		rewardsInterface := variables["rewards"].([]interface{})
		rewards := make([]string, len(rewardsInterface))

		for i, v := range rewardsInterface {
			if str, ok := v.(string); ok {
				rewards[i] = str
			} else {
				fmt.Printf("Non-string value at index %d: %v\n", i, v)
			}
		}

		err = service.UpdateRewardsState(isExpiredOn.Value, rewards)
		if err != nil {
			fmt.Println("Error updating rewards state: ", err)
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
		fmt.Println("***CAMUNDA*** Shutting down Check Transactions worker...")
		w.Close()
	}()
}
