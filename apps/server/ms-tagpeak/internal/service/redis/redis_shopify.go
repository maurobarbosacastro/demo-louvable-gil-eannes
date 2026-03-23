package redis

import (
	"fmt"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto/webhooks"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/logster"
	"strconv"
	"time"
)

func ProcessWebhookRedisQueue(domain string, shopifyOrder webhooks.ShopifyOrder) error {
	logster.Debug(fmt.Sprintf("ShopifyOrder: %+v", shopifyOrder))

	keycloak := auth.KeycloakInstance

	transaction, err := repository.GetTransactionBySourceId(strconv.FormatInt(shopifyOrder.ID, 10))
	if err != nil {
		logster.Error(err, "Error getting transaction")
		logster.EndFuncLog()
		return err
	}

	//Prevent further updates once it validated/rejected
	if transaction != nil && transaction.State != "TRACKED" {
		logster.Error(nil, fmt.Sprintf("Transaction is not updatable when in %s only in TRACKED state", transaction.State))
		logster.EndFuncLog()
		return err
	}

	transactionCamunda, err := service.HandleWebhookCreateOrUpdate(shopifyOrder, transaction, keycloak, domain)

	if err != nil {
		logster.Error(err, "Error handling webhook data")
		logster.EndFuncLog()
		return err
	}

	if shopifyOrder.CancelledAt != nil {
		transactionCamunda.State = "rejected"
	}

	service.StartCamundaProcessForShopifyOrder(transactionCamunda)

	time.Sleep(5 * time.Second)

	logster.EndFuncLog()
	return nil
}
