package camunda_processes

import (
	"context"
	"fmt"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/service"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/email"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
)

const (
	GetEmailData    = "get-email-data"
	SendEmailAction = "send-email-action"
)

/*
### VARAIABLES NEEDED ###
Email template code
Store uuid
Store name
User name
*/

func HandleGetEmailData() {
	fmt.Println("***CAMUNDA*** Start Get Email Data")

	w, _ := camundaPKG.StartDynamicWorker(GetEmailData, func(client worker.JobClient, job entities.Job) {
		logster.StartFuncLogMsg(job.ProcessInstanceKey)

		getKeycloakAdminToken(auth.KeycloakInstance)

		variables, err := job.GetVariablesAsMap()

		variablesObject := map[string]interface{}{}

		if err != nil {
			logster.Error(err, "Error getting variables")
			logster.EndFuncLog()
			return
		}

		// Get User Information
		userUuid := variables["userUuid"].(string)
		userUuidSubstitute := dotenv.GetEnv("USER_SUBSTITUTION_UUID")
		if userUuid == userUuidSubstitute {
			logster.Warn("Cannot send email for errored awin transactions")

			emailVar := map[string]interface{}{
				"abort": true,
			}

			request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(emailVar)
			if err != nil {
				logster.Error(err, "Error completing job")
				logster.EndFuncLog()
				return
			}

			_, err = request.Send(context.Background())
			if err != nil {
				logster.Error(err, "Error completing job")
				logster.EndFuncLog()
				return
			}

			return
		}

		user, err := service.GetUserById(userUuid, auth.KeycloakInstance)

		var userEmail *string = nil
		var userNewsletterSubscription *bool = nil
		isNewsletter := false

		if isNewsletterAux, ok := variables["isNewsletter"].(bool); ok {
			isNewsletter = isNewsletterAux
		}

		if emailVars, ok := variables["email"].(map[string]interface{}); ok {
			variablesObject = emailVars
		}

		if err == nil && user != nil {
			variablesObject["firstName"] = &user.FirstName
			variablesObject["lastName"] = &user.LastName
			userEmail = &user.Email
			userNewsletterSubscription = &user.Newsletter
		}

		if storeUuid, ok := variables["storeUuid"].(string); ok {
			store, err := service.GetStore(utils.ParseIDToUUID(storeUuid))
			if err == nil {
				variablesObject["store"] = store.Name
			}
		}

		if transactionUuid, ok := variables["transactionUuid"].(string); ok {
			storeVisit, err := service.GetStoreVisitByTransaction(utils.ParseIDToUUID(transactionUuid))
			if err == nil {
				variablesObject["refId"] = storeVisit.Reference
			}

			storeByTransaction, err := service.GetStoreByTransaction(utils.ParseIDToUUID(transactionUuid))
			if err == nil {
				variablesObject["store"] = storeByTransaction.Name
			}
		}

		if rewardUuid, ok := variables["rewardUuid"].(string); ok {
			rewards, err := service.GetReward(utils.ParseIDToUUID(rewardUuid))
			if err == nil {
				variablesObject["value"] = rewards.CurrentRewardUser
			}
		}

		emailVar := map[string]interface{}{
			"email":        variablesObject,
			"userEmail":    userEmail,
			"newsletterOn": userNewsletterSubscription,
			"isNewsletter": isNewsletter,
			"abort":        false,
		}

		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(emailVar)
		if err != nil {
			logster.Error(err, "Error completing job")
			logster.EndFuncLog()
			return
		}

		_, err = request.Send(context.Background())
		if err != nil {
			logster.Error(err, "Error sending request")
			logster.EndFuncLog()
			return
		}
		logster.EndFuncLog()
	})

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Get Email Data worker...")
		w.Close()
	}()
}

func HandleSendEmailAction() {
	fmt.Println("***CAMUNDA*** Start Send Email Action")

	w, _ := camundaPKG.StartDynamicWorker(SendEmailAction, func(client worker.JobClient, job entities.Job) {
		logster.StartFuncLogMsg(job.ProcessInstanceKey)

		// GET Rewards by transaction uuid
		variables, err := job.GetVariablesAsMap()
		if err != nil {
			logster.Error(err, "Error getting variables")
			logster.EndFuncLog()
			return
		}

		templateCode := variables["template_code"].(string)
		logster.Info(fmt.Sprintf("TemplateCode: %v", templateCode))

		emailV, _ := variables["email"].(interface{})
		logster.Info(fmt.Sprintf("Email: %v", emailV))

		_, errMessage := email.SendEmail(email.SendEmailDTO{To: variables["userEmail"].(string), Dictionary: emailV.(map[string]interface{})}, templateCode)
		if errMessage != nil {
			logster.Error(nil, fmt.Sprintf("Error - %v", errMessage))
		}

		_, err = client.NewCompleteJobCommand().JobKey(job.GetKey()).Send(context.Background())
		if err != nil {
			logster.Error(err, "Error completing job")
			logster.EndFuncLog()
			return
		}
		logster.EndFuncLog()
	})

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Get Email Data worker...")
		w.Close()
	}()
}
