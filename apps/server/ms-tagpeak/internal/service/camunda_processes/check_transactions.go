package camunda_processes

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
	"github.com/google/uuid"
)

const (
	CheckTransactions        = "check-transactions"
	CheckTransactionIfNew    = "check-transaction-if-new"
	CreateTransactionCamunda = "create-transaction"
	ConvertCurrency          = "convert-currency"
	UpdateTransactionCamunda = "update-transaction"
	CheckNotificationSend    = "check-notification-send"
)

// *** FLOW ***
// Starts in HandleCheckTransactions
// If there are transactions, sends them to the CheckTransactionIfNew process
// If there are no new transactions, updates the transaction
// If there are new transactions create new ones
// After creating the new ones, sends them to the ConvertCurrency process

func HandleCheckTransactions() {
	fmt.Println("***CAMUNDA*** Start Check Transactions")
	w, _ := camundaPKG.StartDynamicWorker(CheckTransactions, func(client worker.JobClient, job entities.Job) {
		logster.StartFuncLogMsg(job.ProcessInstanceKey)
		logster.Info(fmt.Sprintf("Job key: %v", job.Key))

		variables, err := job.GetVariablesAsMap()
		logster.Info(fmt.Sprintf("Variables: %+v", variables))
		if err != nil {
			logster.Error(err, "Error getting variables as map")
			logster.EndFuncLog()
			return
		}

		awinEnabled := dotenv.GetEnv("AWIN_ENABLED")
		logster.Warn(fmt.Sprintf("AwinEnabled: %s", awinEnabled))
		if awinEnabled == "false" {

			variables["transactions"] = []dto.CamundaCreateTransactionDTO{}
			request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(variables)
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
			return
		}
		awinRefPrefix := dotenv.GetEnv("AWIN_REF_PREFIX")

		getKeycloakAdminToken(auth.KeycloakInstance)

		var panicError error
		var storeVisit *models.StoreVisit

		var dateType string
		dateType, ok := variables["dateType"].(string)
		if !ok {
			dateType = "transaction"
		}

		awinTransactions, err := service.GetAwinTransactions(dateType)
		if err != nil {
			// Logs are already shown in the service call

			variables["transactions"] = []dto.CamundaCreateTransactionDTO{}
			request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(variables)
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
			return
		}
		logster.Info(fmt.Sprintf("Awin transactions: %+v", awinTransactions))
		var transactions []dto.CamundaCreateTransactionDTO

		for _, awinTransaction := range *awinTransactions {
			panicError = nil

			func() {
				defer func() {
					if r := recover(); r != nil {
						panicError = fmt.Errorf("panic during awin ref handling: %v", r)
					}
				}()

				awinRef := awinTransaction.ClickRefs.ClickRef
				if awinRef == "" {
					storeVisit = nil
					logster.Panic(nil, "Awin ref is empty")
					return
				}
				ref, sectionCheck := utils.DecodeStoreVisitRef(awinRef)
				logster.Info(fmt.Sprintf("Ref: %s | Section: %s", ref, sectionCheck))

				if strings.Contains(ref, "-") {
					if !strings.HasPrefix(ref, awinRefPrefix+"-") {
						logster.Panic(fmt.Errorf("Reference prefix mismatch. Expected '%s-' but got '%s'.", awinRefPrefix, ref), "Skipping transaction from different environment.")
						return // Skip this transaction
					}
				}

				storeVisitTmp, err := service.GetStoreVisitByRef(ref)
				if err != nil {
					logster.Panic(err, "Error getting store visit")
					return
				}
				storeVisit = storeVisitTmp

				// This ensures awin transactions made in another environment are not processed if not in the right env.
				if sectionCheck != "" {
					if !utils.ValidateAwinDecodedStoreVisitRef(sectionCheck, *storeVisit.User) {
						storeVisit = nil
						logster.Panic(nil, "Invalid section check")
						return
					}
				}
			}()

			if panicError != nil {
				logster.Error(panicError, "Error handling click ref")
				storeVisit = nil
				continue
			}

			transactionDate, _ := time.Parse("2006-01-02T15:04:05Z", awinTransaction.TransactionDate+"Z")

			amount := float64(0)
			currency := "EUR"
			commission := float64(0)
			if awinTransaction.SaleAmount.Amount != nil {
				amount = *awinTransaction.SaleAmount.Amount
			}
			if awinTransaction.SaleAmount.Currency != nil {
				currency = *awinTransaction.SaleAmount.Currency
			}
			if awinTransaction.CommissionAmount.Amount != nil {
				commission = *awinTransaction.CommissionAmount.Amount
			}

			if storeVisit != nil {
				transactions = append(transactions, dto.CamundaCreateTransactionDTO{
					SourceId:         strconv.FormatInt(awinTransaction.ID, 10),
					AmountSource:     amount,
					CurrencySource:   currency,
					CommissionSource: commission,
					OrderDate:        transactionDate,
					StoreVisitUUID:   &storeVisit.Uuid,
					UserUUID:         utils.ParseIDToUUID(*storeVisit.User),
					Reference:        *storeVisit.Reference,
					State:            awinTransaction.CommissionStatus,
				})
			} /* else {
			    //Try to find the transaction by sourceId if it was created by hand in the bo or db.
			    sourceId := strconv.FormatInt(awinTransaction.ID, 10)

			    transaction, errTrx := service.GetTransactionBySourceId(sourceId)
			    if errTrx != nil {
			        logster.Error(errTrx, "Error getting transaction by sourceId")
			    }
			    if transaction != nil && transaction.StoreVisit != nil {
			        transactions = append(transactions, dto.CamundaCreateTransactionDTO{
			            SourceId:         strconv.FormatInt(awinTransaction.ID, 10),
			            AmountSource:     amount,
			            CurrencySource:   currency,
			            CommissionSource: commission,
			            OrderDate:        transactionDate,
			            StoreVisitUUID:   &transaction.StoreVisit.Uuid,
			            UserUUID:         uuid.MustParse(transaction.User),
			            Reference:        *transaction.StoreVisit.Reference,
			            State:            awinTransaction.CommissionStatus,
			        })
			    }
			}*/
			storeVisit = nil
		}

		variables["transactions"] = transactions
		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(variables)
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
		fmt.Println("***CAMUNDA*** Shutting down Check Transactions worker...")
		w.Close()
	}()
}

func HandleCheckTransactionIfNew() {
	fmt.Println("***CAMUNDA*** Start Process Transaction")
	w, _ := camundaPKG.StartDynamicWorker(CheckTransactionIfNew, func(client worker.JobClient, job entities.Job) {
		logster.StartFuncLogMsg(job.ProcessInstanceKey)
		logster.Info(fmt.Sprintf("Job key: %v", job.Key))

		transactionVariables := map[string]interface{}{}

		transactions := job.GetVariables()

		getKeycloakAdminToken(auth.KeycloakInstance)

		err := json.Unmarshal([]byte(transactions), &transactionVariables)
		if err != nil {
			logster.Error(err, "Error unmarshalling transactions")
		}
		logster.Info(fmt.Sprintf("Variables: %+v", transactionVariables))

		loopCounter := transactionVariables["loopCounter"].(float64)
		allTransactions := transactionVariables["transactions"].([]interface{})

		//Get transactions by reference
		transaction := allTransactions[int(loopCounter)-1].(map[string]interface{})

		sourceId, sourceIdExists := transaction["sourceId"].(string)

		var value *models.Transaction
		if sourceIdExists {
			logster.Info("SourceId: " + sourceId)

			//Check if the transaction has been already processed by click ref and sourceId
			//Transactions have to have the sourceId, if not, it will be a new transaction
			v, err := service.GetTransactionBySourceId(sourceId)
			if err != nil {
				logster.Error(err, "Error getting transaction by sourceId")
			}

			if v != nil {
				value = v
				transaction["transactionUuid"] = value.Uuid
			} else {
				logster.Info("Nothing found by sourceId, going by store visit uuid")
				svUuid, okVarSv := transaction["storeVisitUuid"].(string)
				if okVarSv {
					trx, errTrx := service.GetTransactionByStoreVisitUuid(svUuid)

					if errTrx != nil {
						logster.Error(errTrx, "Error getting transaction by store visit")
					}

					//Only use the transaction if no source id is set.
					if trx != nil && trx.SourceId == "" {
						value = trx
						transaction["transactionUuid"] = trx.Uuid
					}

				}
			}
		} else {
			logster.Info("No sourceId, going by store visit uuid")
			svUuid, okVarSv := transaction["storeVisitUuid"].(string)
			if okVarSv {
				trx, errTrx := service.GetTransactionByStoreVisitUuid(svUuid)
				if errTrx != nil {
					logster.Error(errTrx, "Error getting transaction by store visit")
				}
				if trx != nil {
					value = trx
					transaction["transactionUuid"] = trx.Uuid
				}
			}
		}

		var isProcessed bool

		if value != nil {
			isProcessed = value.IsProcessed
		} else {
			isProcessed = false
		}

		transactionVariablesToSend := map[string]interface{}{
			"newTransaction":   value == nil,
			"processed":        isProcessed,
			"transactionValue": transaction,
		}

		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(transactionVariablesToSend)
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
		fmt.Println("***CAMUNDA*** Shutting down Check Transactions worker...")
		w.Close()
	}()
}

func HandleCreateNewTransaction() {
	fmt.Println("***CAMUNDA*** Start Create New Transaction")
	w, _ := camundaPKG.StartDynamicWorker(CreateTransactionCamunda, func(client worker.JobClient, job entities.Job) {
		logster.StartFuncLogMsg(job.ProcessInstanceKey)
		logster.Info(fmt.Sprintf("Job key: %v", job.Key))

		transactionVariables := map[string]interface{}{}

		transactions := job.GetVariables()

		err := json.Unmarshal([]byte(transactions), &transactionVariables)
		if err != nil {
			logster.Error(err, "Error unmarshalling transactions")
			logster.EndFuncLog()
		}

		transactionCamunda := transactionVariables["transactionValue"].(map[string]interface{})
		transactionDate := transactionCamunda["orderDate"].(string)
		logster.Info(fmt.Sprintf("New - Transaction: %v", transactionCamunda))
		logster.Info(fmt.Sprintf("New - Transaction date: %v", transactionDate))

		date, _ := time.Parse("2006-01-02T15:04:05Z", transactionDate)

		var storeVisitUUID *uuid.UUID
		if transactionCamunda["storeVisitUuid"] != nil {
			tmp, _ := uuid.Parse(transactionCamunda["storeVisitUuid"].(string))
			storeVisitUUID = utils.UuidPointer(tmp)
		} else {
			storeVisitUUID = nil
		}

		amountSource := float64(0.0)
		if transactionCamunda["amount"] != nil {
			amountSource = transactionCamunda["amount"].(float64)
		}

		createDTO := dto.CreateTransactionDTO{
			SourceId:         transactionCamunda["sourceId"].(string),
			AmountSource:     amountSource,
			CurrencySource:   transactionCamunda["currency"].(string),
			OrderDate:        date,
			StoreVisitUUID:   storeVisitUUID,
			CommissionSource: 0.0,
		}

		if transactionCamunda["commission"] != nil {
			createDTO.CommissionSource = transactionCamunda["commission"].(float64)
		}

		var storeVisit *models.StoreVisit
		if storeVisitUUID != nil {
			storeVisit, err = service.GetStoreVisit(*storeVisitUUID)
			if err != nil {
				logster.Error(err, "Error getting store visit")
			}
		}

		userId := transactionCamunda["userUuid"].(string)

		getKeycloakAdminToken(auth.KeycloakInstance)

		transaction, err := service.CreateTransaction(createDTO, storeVisit, userId)
		if err != nil {
			logster.Error(err, "Error creating transaction")
			logster.EndFuncLog()
			return
		}

		var storeUuid *uuid.UUID
		if storeVisit != nil {
			storeUuid = storeVisit.StoreUUID
		}
		transactionValue := map[string]interface{}{
			"transactionUuid": transaction.Uuid,
			"storeUuid":       storeUuid,
			"userUuid":        userId,
		}

		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(transactionValue)
		if err != nil {
			logster.Error(err, "Error creating job")
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
		fmt.Println("***CAMUNDA*** Shutting down Check Transactions worker...")
		w.Close()
	}()
}

func HandleUpdateTransaction() {
	fmt.Println("***CAMUNDA*** Start Update Transaction")
	w, _ := camundaPKG.StartDynamicWorker(UpdateTransactionCamunda, func(client worker.JobClient, job entities.Job) {
		logster.StartFuncLogMsg(job.ProcessInstanceKey)

		//Update Transaction
		transactionVariables := map[string]interface{}{}

		transactions := job.GetVariables()

		getKeycloakAdminToken(auth.KeycloakInstance)

		err := json.Unmarshal([]byte(transactions), &transactionVariables)
		if err != nil {
			logster.Error(err, "Error unmarshalling transactions")
			logster.EndFuncLog()
			return
		}

		transactionCamunda := transactionVariables["transactionValue"].(map[string]interface{})
		currency := transactionCamunda["currency"].(string)
		statusAux := transactionCamunda["state"].(string)
		logster.Info(fmt.Sprintf("Update - Transaction: %v", transactionCamunda))

		amountSource := float64(0.0)
		if transactionCamunda["amount"] != nil {
			amountSource = transactionCamunda["amount"].(float64)
		}

		commissionSource := 0.0
		if transactionCamunda["commission"] != nil {
			commissionSource = transactionCamunda["commission"].(float64)
		}

		var transactionStatus string
		switch strings.ToLower(statusAux) {
		case "approved", "validated":
			transactionStatus = "VALIDATED"
			break
		case "rejected":
			transactionStatus = "REJECTED"
			break
		case "declined":
			transactionStatus = "REJECTED"
			break
		default:
			transactionStatus = "TRACKED"
		}

		transactionToUpdate := dto.UpdateTransactionDTO{
			CurrencySource:   &currency,
			AmountSource:     &amountSource,
			CommissionSource: &commissionSource,
			State:            &transactionStatus,
		}

		transactionUUID, err := uuid.Parse(transactionCamunda["transactionUuid"].(string))
		if err != nil {
			logster.Error(err, fmt.Sprintf("Error parsing transactionUuid %v", transactionCamunda["transactionUuid"]))
			logster.EndFuncLog()
			return
		}
		logster.Info(fmt.Sprintf("Treating transaction %s", transactionUUID.String()))

		if transactionStatus == "VALIDATED" {
			t, _ := service.GetTransaction(transactionUUID)
			if t.StoreVisitUUID != nil {
				logster.Info(fmt.Sprintf("Setting store visit %s purchase flag to true", *t.StoreVisitUUID))
				_, err = service.UpdateStoreVisit(dto.UpdateStoreVisitDTO{Purchase: utils.BoolPointer(true)}, *t.StoreVisitUUID, "camunda")
				if err != nil {
					logster.Error(err, "Error updating purchase flag for store visit. Proceeding without it.")
				}
			}

		}

		transaction, err := service.UpdateTransaction(transactionToUpdate, transactionUUID, nil, false)
		if err != nil {
			logster.Error(err, "Error updating transaction")
			logster.EndFuncLog()
			return
		}

		userId := transactionCamunda["userUuid"].(string)

		transactionValue := map[string]interface{}{
			"transactionUuid":  transaction.Uuid,
			"transactionState": transaction.State,
			"userUuid":         userId,
		}

		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(transactionValue)
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
		fmt.Println("***CAMUNDA*** Shutting down Check Transactions worker...")
		w.Close()
	}()
}

func HandleConvertToTagpeakCurrency() {
	fmt.Println("***CAMUNDA*** Start Convert To Tagpeak Currency")
	w, _ := camundaPKG.StartDynamicWorker(ConvertCurrency, func(client worker.JobClient, job entities.Job) {
		logster.StartFuncLogMsg(job.ProcessInstanceKey)

		transactionVariables := map[string]interface{}{}

		transactions := job.GetVariables()

		err := json.Unmarshal([]byte(transactions), &transactionVariables)
		if err != nil {
			logster.Error(err, "Error unmarshalling transactions")
			logster.EndFuncLog()
			return
		}

		transactionUUID, err := uuid.Parse(transactionVariables["transactionUuid"].(string))
		if err != nil {
			logster.Error(err, "Error parsing transactionUuid")
			logster.EndFuncLog()
			return
		}
		logster.Info(fmt.Sprintf("Convert - Transaction %s", transactionUUID.String()))

		userUuid := transactionVariables["userUuid"].(string)

		getKeycloakAdminToken(auth.KeycloakInstance)

		service.ConvertToTagpeakCurrency(transactionUUID, userUuid, auth.KeycloakInstance)

		_, err = client.NewCompleteJobCommand().JobKey(job.GetKey()).Send(context.Background())
		if err != nil {
			fmt.Println(err)
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
		fmt.Println("***CAMUNDA*** Shutting down Check Transactions worker...")
		w.Close()
	}()
}

func HandleCheckNotificationSend() {
	fmt.Println("***CAMUNDA*** Start Check Notification Send")
	w, _ := camundaPKG.StartDynamicWorker(CheckNotificationSend, func(client worker.JobClient, job entities.Job) {
		logster.StartFuncLogMsg(job.ProcessInstanceKey)
		logster.Info(fmt.Sprintf("Job key: %v", job.Key))

		transactionVariables := map[string]interface{}{}

		transactions := job.GetVariables()

		err := json.Unmarshal([]byte(transactions), &transactionVariables)
		if err != nil {
			logster.Error(err, "Error unmarshalling transactions")
			variablesToSend := map[string]interface{}{
				"canSend": true,
			}
			logster.EndFuncLogMsg(fmt.Sprintf("Variables: %+v", variablesToSend))
			camundaPKG.EndProcess(context.Background(), client, job.GetKey(), &variablesToSend)
		}

		//Get user from the transaction
		userUuid, _ := transactionVariables["userUuid"].(string)
		user, errGetUser := service.GetUserById(userUuid, auth.KeycloakInstance)
		if errGetUser != nil || user == nil {
			logster.Error(errGetUser, "Error getting user by id")
			variablesToSend := map[string]interface{}{
				"canSend": true,
			}
			logster.EndFuncLogMsg(fmt.Sprintf("Variables: %+v", variablesToSend))
			camundaPKG.EndProcess(context.Background(), client, job.GetKey(), &variablesToSend)
		}

		userr := *user
		//If the user does not come from Shopify, send notification
		if userr.Source != models.SOURCE_SHOPIFY {
			logster.Info("User is not from Shopify, sending notification")
			variablesToSend := map[string]interface{}{
				"canSend": true,
			}
			logster.EndFuncLogMsg(fmt.Sprintf("Variables: %+v", variablesToSend))
			camundaPKG.EndProcess(context.Background(), client, job.GetKey(), &variablesToSend)
		}

		//Check if this is the first transaction of the user by checking if the user has more than one transaction.
		//If has more than one, we send the notifications. If not, we don't send them and they are sent along with the welcome email from keycloak.
		hasMoreThanOne, errHasMoreThanOne := service.UserHasMoreThanOneTransaction(userr.Uuid.String())
		if errHasMoreThanOne != nil {
			logster.Error(errHasMoreThanOne, "Error checking if user has more than one transaction")
			variablesToSend := map[string]interface{}{
				"canSend": true,
			}
			logster.EndFuncLogMsg(fmt.Sprintf("Variables: %+v", variablesToSend))
			camundaPKG.EndProcess(context.Background(), client, job.GetKey(), &variablesToSend)
		}

		var variablesToSend map[string]interface{}
		if hasMoreThanOne {
			variablesToSend = map[string]interface{}{
				"canSend": true,
			}
		} else {
			variablesToSend = map[string]interface{}{
				"canSend": false,
			}
		}
		logster.EndFuncLogMsg(fmt.Sprintf("Variables: %+v", variablesToSend))
		camundaPKG.EndProcess(context.Background(), client, job.GetKey(), &variablesToSend)

	})

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Check Notification Send worker...")
		w.Close()
	}()
}
