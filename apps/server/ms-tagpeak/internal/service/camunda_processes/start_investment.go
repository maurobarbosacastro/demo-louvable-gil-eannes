package camunda_processes

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/service"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"os"
	"os/signal"
	"syscall"

	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
	"github.com/google/uuid"
)

const (
	StartInvestment         = "start-investment"
	CheckIfInvestmentExists = "check-if-investment-exists"
	TransactionChangeState  = "transaction-change-state"
	CheckMembershipLevel    = "check-membership-level"
)

// *** FLOW ***
// Process starts in  CheckTransactions in the HandleUpdateTransaction
// Validates the transaction state
// If the transactions is Rejected ends the process
// If the transactions is Valided, starts the investment process
// Manual task to create the investment
// After the investment is created, sends the transaction to the HandleCheckIfInvestmentExists

func StartInvestmentProcess(transaction map[string]interface{}) {
	//START CAMUNDA PROCESS
	process := camundaPKG.StartProcessInstance(camundaPKG.InjectEnvOnKey(StartInvestment), *camundaPKG.GetCamundaClient(), transaction)

	camundaProcessDto := dto.CreateCamundaProcessDto{
		FieldUUID:          transaction["transactionUuid"].(uuid.UUID),
		ProcessInstanceKey: process.ProcessInstanceKey,
		ProcessId:          StartInvestment,
	}

	// Saved process for manual task
	_, err := service.CreateCamundaProcess(camundaProcessDto)
	if err != nil {
		fmt.Println("Error creating camunda process: ", err)
		return
	}
}

func HandleTransactionChangeState() {
	fmt.Println("***CAMUNDA*** Start Transaction Change State")

	w, _ := camundaPKG.StartDynamicWorker(TransactionChangeState, func(client worker.JobClient, job entities.Job) {

		//GET Rewards by transaction uuid
		investmentVariables := map[string]interface{}{}

		variables := job.GetVariables()

		err := json.Unmarshal([]byte(variables), &investmentVariables)
		if err != nil {
			fmt.Println(err)
			return
		}

		transactionUUIDString := investmentVariables["transactionUuid"].(string)
		transactionState := investmentVariables["transactionState"].(string)
		transactionUUID, err := uuid.Parse(transactionUUIDString)

		if transactionState == "VALIDATED" {
			camundaProcessDto := dto.CreateCamundaProcessDto{
				FieldUUID:          transactionUUID,
				ProcessInstanceKey: job.ProcessInstanceKey,
				ProcessId:          StartInvestment,
			}

			// Saved process for manual task
			_, err = service.CreateCamundaProcess(camundaProcessDto)
			if err != nil {
				fmt.Println("Error creating camunda process: ", err)
				return
			}
		}

		err = service.UpdateTransactionToProcessed(transactionUUID, true)
		if err != nil {
			fmt.Println("Error updating transaction to processed: ", err)
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

func HandleCheckIfInvestmentExists() {
	fmt.Println("***CAMUNDA*** Start Check Transactions")
	w, _ := camundaPKG.StartDynamicWorker(CheckIfInvestmentExists, func(client worker.JobClient, job entities.Job) {

		//GET Rewards by transaction uuid
		investmentVariables := map[string]interface{}{}

		variables := job.GetVariables()

		err := json.Unmarshal([]byte(variables), &investmentVariables)
		if err != nil {
			fmt.Println(err)
			return
		}

		transactionUUIDString := investmentVariables["transactionUuid"].(string)

		transactionUUID, err := uuid.Parse(transactionUUIDString)
		if err != nil {
			fmt.Println(err)
			return
		}

		res := service.IsRewardsWithTransactionSaved(transactionUUID)

		// Log if Rewards with Transaction exists
		fmt.Println("Is Rewards with Transaction Saved: ", res)

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

func HandleCheckUserMembershipLevel() {
	fmt.Println("***CAMUNDA*** Start Check User Membership Level")
	w, _ := camundaPKG.StartDynamicWorker(CheckMembershipLevel, func(client worker.JobClient, job entities.Job) {

		getKeycloakAdminToken(auth.KeycloakInstance)

		transactionVariables := map[string]interface{}{}

		transactions := job.GetVariables()

		err := json.Unmarshal([]byte(transactions), &transactionVariables)
		if err != nil {
			fmt.Println(err)
			return

		}
		userUuid := transactionVariables["userUuid"].(string)

		err = service.ManageUserMembershipLevel(userUuid, auth.KeycloakInstance)
		if err != nil {
			fmt.Println(err)
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
		fmt.Println("***CAMUNDA*** Shutting down Check User Membership Level worker...")
		w.Close()
	}()
}
