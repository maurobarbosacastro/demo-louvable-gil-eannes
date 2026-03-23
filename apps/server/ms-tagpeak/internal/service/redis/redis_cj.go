package redis

import (
	"fmt"
	"math"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/internal/service"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"strconv"
	"strings"
	"time"
)

func ProcessCjTransaction(cjTransaction *response_object.CommissionRecord) error {
	logster.StartFuncLogMsg(cjTransaction.OrderID)
	var transaction dto.CamundaCreateTransactionDTO

	cjShopperId := cjTransaction.ShopperId
	storeVisit, err := service.GetStoreVisitByRef(cjShopperId)
	if err != nil {
		logster.Error(err, "Error getting store visit")
		return err
	}

	transactionDate, err := time.Parse("2006-01-02T15:04:05Z", cjTransaction.EventDate)
	if err != nil {
		logster.Error(err, "Error parsing transaction date")
		return err
	}

	amountSource, err := strconv.ParseFloat(cjTransaction.SaleAmountUSD, 64)
	if err != nil {
		logster.Error(err, "Error parsing transaction amount")
		return err
	}

	commissionSource, err := strconv.ParseFloat(cjTransaction.PubCommissionAmountUSD, 64)
	if err != nil {
		logster.Error(err, "Error parsing commission amount")
		return err
	}

	state := strings.ToLower(cjTransaction.ValidationStatus)

	switch state {
	case "accepted":
		state = "approved"
	case "automated":
		state = "pending"
	}

	ignore := false

	if !cjTransaction.Original {
		logster.Warn("Not original transaction, it's a correction. Searching for original transaction and updating it.")
		func() {
			//Since it's not a new transaction, it should exist by sourceId
			transactionToUpdate, err := service.GetTransactionBySourceId(cjTransaction.OrderID)

			if err != nil {
				ignore = true
				logster.Error(err, fmt.Sprintf("Error searching for transaction with source id %s", cjTransaction.OrderID))
				return
			}

			updateAmount := transactionToUpdate.AmountSource - math.Abs(amountSource)
			logster.Info(fmt.Sprintf("Transaction amount %f", updateAmount))

			// Negative or 0 value is a full refund, so the transaction is declined.
			if updateAmount <= 0 {
				logster.Info("Full refund so the transaction is declined but do not update the amounts.")
				state = "declined"
				return
			}

			// Update amounts with new values.
			amountSource = updateAmount
			commissionSource = transactionToUpdate.CommissionSource - math.Abs(commissionSource)
		}()
	}

	if ignore {
		logster.EndFuncLogMsg("Ignored transaction.")
		return nil
	}

	transaction = dto.CamundaCreateTransactionDTO{
		SourceId:         cjTransaction.OrderID,
		AmountSource:     amountSource,
		CurrencySource:   "USD", // Once we are using the value in USD, after we will change it to the right currency
		CommissionSource: commissionSource,
		OrderDate:        transactionDate,
		StoreVisitUUID:   &storeVisit.Uuid,
		UserUUID:         utils.ParseIDToUUID(*storeVisit.User),
		Reference:        *storeVisit.Reference,
		State:            state,
	}

	StartCamundaProcessForCJ(&transaction)

	logster.EndFuncLog()
	return nil
}

func StartCamundaProcessForCJ(transactionCamunda *dto.CamundaCreateTransactionDTO) {
	logster.StartFuncLog()
	trx := map[string]interface{}{
		"loopCounter":  1,
		"transactions": [1]interface{}{transactionCamunda},
	}

	resp := camundaPKG.StartProcessInstance(
		camundaPKG.InjectEnvOnKey("sub-process-transaction"),
		*camundaPKG.GetCamundaClient(),
		trx,
	)
	logster.Info(fmt.Sprintf("Process started -Process Instance Key : %d", resp.ProcessInstanceKey))

	logster.EndFuncLog()
}
