package redis

import (
	"encoding/json"
	"fmt"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/dto/webhooks"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/redisclient"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	CjTransactionKey           = "cj_transaction_queue"
	ShopifyRedisQueueKey       = "shopify_orders_queue"
	FailedShopifyRedisQueueKey = "failed_shopify_orders_queue"
	DeleteShopifyRedisQueueKey = "delete_shopify_orders_queue"
)

func PushToRedisQueue(queue string, transactions *response_object.CommissionRecord) {
	client := redisclient.GetRedisClient()
	ctx := redisclient.GetContext()

	transactionJSON, err := json.Marshal(transactions)
	if err != nil {
		logster.Error(err, "Error marshaling transaction to JSON")
		return
	}

	err = client.LPush(ctx, queue, transactionJSON).Err()
	if err != nil {
		logster.Error(err, "Error pushing to queue")
		return
	}

	logster.Info(fmt.Sprintf("Pushed to Redis queue %s: %v", queue, transactions))
}

func PushToShopifyRedisQueue(queue string, topic string, shop string, orderData string) {
	client := redisclient.GetRedisClient()
	ctx := redisclient.GetContext()

	stringToPush := fmt.Sprintf("%s|%s|%s", topic, shop, orderData)

	err := client.LPush(ctx, queue, stringToPush).Err()
	if err != nil {
		logster.Error(err, "Error pushing to queue")
		logster.Info(stringToPush)
		return
	}

	logster.Info(fmt.Sprintf("Pushed to Redis queue %s: %s", queue, stringToPush))
}

func ProcessCjTransactionsRedisQueue() {
	client := redisclient.GetRedisClient()
	ctx := redisclient.GetContext()

	for {
		result, err := client.BRPop(ctx, 0*time.Second, CjTransactionKey).Result()
		if err != nil {
			logster.Error(err, "Error popping from queue")
			continue
		}

		value := result[1]

		var transaction response_object.CommissionRecord
		err = json.Unmarshal([]byte(value), &transaction)
		if err != nil {
			logster.Error(err, "failed to unmarshal Redis message")
			continue
		}

		errProcess := ProcessCjTransaction(&transaction)

		if errProcess != nil {
			logster.Error(errProcess, "Error processing queue item")
			PushToRedisQueue(CjTransactionKey, &transaction)
			continue
		}

		logster.EndFuncLog()
	}
}

func ProcessShopifyOrderQueue() {
	client := redisclient.GetRedisClient()
	ctx := redisclient.GetContext()

	for {
		result, err := client.BRPop(ctx, 0*time.Second, ShopifyRedisQueueKey).Result()
		if err != nil {
			logster.Error(err, "Error popping from queue")
			continue
		}

		orderData := result[1]
		orderDataSplit := strings.Split(orderData, "|")
		topic := orderDataSplit[0]
		shop := orderDataSplit[1]
		data := orderDataSplit[2]

		logster.Info(fmt.Sprintf("Processing order for : Topic %s | Shop %s", topic, shop))

		var shopifyOrder webhooks.ShopifyOrder
		err = json.Unmarshal([]byte(data), &shopifyOrder)
		if err != nil {
			logster.Error(err, "failed to unmarshal Redis message")
			continue
		}

		errProcess := ProcessWebhookRedisQueue(shop, shopifyOrder)

		if errProcess != nil {
			logster.Error(errProcess, "Error processing queue item")
			PushToShopifyRedisQueue(FailedShopifyRedisQueueKey, topic, shop, orderData)
			continue
		}

		logster.EndFuncLog()
	}
}

func ProcessShopifyDeleteQueue() {
	client := redisclient.GetRedisClient()
	ctx := redisclient.GetContext()

	for {
		result, err := client.BRPop(ctx, 0*time.Second, DeleteShopifyRedisQueueKey).Result()
		if err != nil {
			logster.Error(err, "Error popping from queue")
			continue
		}

		orderData := result[1]
		orderDataSplit := strings.Split(orderData, "|")
		topic := orderDataSplit[0]
		shop := orderDataSplit[1]
		data := orderDataSplit[2]

		logster.Info(fmt.Sprintf("Processing order for : Topic %s | Shop %s", topic, shop))

		var shopifyOrderDelete webhooks.ShopifyOrderDelete
		err = json.Unmarshal([]byte(data), &shopifyOrderDelete)
		if err != nil {
			logster.Error(err, "failed to unmarshal Redis message")
			PushToShopifyRedisQueue(FailedShopifyRedisQueueKey, topic, shop, orderData)
			continue
		}

		logster.Info(fmt.Sprintf("%+v", shopifyOrderDelete))

		transaction, err := repository.GetTransactionBySourceId(strconv.FormatInt(shopifyOrderDelete.ID, 10))
		if err != nil {
			logster.Error(err, "Error getting transaction")
			logster.EndFuncLog()
			PushToShopifyRedisQueue(FailedShopifyRedisQueueKey, topic, shop, orderData)
			continue
		}

		transactionCamunda := &dto.CamundaCreateTransactionDTO{
			SourceId:         strconv.FormatInt(shopifyOrderDelete.ID, 10),
			AmountSource:     transaction.AmountSource,
			CurrencySource:   transaction.CurrencySource,
			CommissionSource: transaction.CommissionSource,
			OrderDate:        transaction.OrderDate,
			StoreVisitUUID:   transaction.StoreVisitUUID,
			UserUUID:         uuid.MustParse(transaction.User),
			Reference:        "",
			State:            "rejected",
		}

		service.StartCamundaProcessForShopifyOrder(transactionCamunda)

	}
}

const (
	BalanceLockTimeout = 30 * time.Second      // SET NX expiration
	BalanceLockRetries = 20                    // Max retry attempts
	BalanceLockDelay   = 50 * time.Millisecond // Delay between retries
)

// balanceLockKey - Lock key template
func balanceLockKey(userUUID string) string {
	return fmt.Sprintf("balance_lock:%s", userUUID)
}

// HandleKeycloakBalanceUpdate - Thread-safe balance update with SET NX lock
// Drop-in replacement for manual GetUserById + UpdateUser pattern
func HandleKeycloakBalanceUpdate(userUUID string, amountToAdd float64, keycloak *constants.Keycloak) error {
	lockKey := balanceLockKey(userUUID)
	lockValue := fmt.Sprintf("%d", time.Now().UnixNano()) // Unique lock value

	client := redisclient.GetRedisClient()
	ctx := redisclient.GetContext()

	// Try to acquire lock with retries
	var lockAcquired bool
	for attempt := 0; attempt < BalanceLockRetries; attempt++ {
		// SET NX with expiration
		acquired, err := client.SetNX(ctx, lockKey, lockValue, BalanceLockTimeout).Result()
		if err != nil {
			// Redis unavailable - fallback: proceed without lock (risk accepted)
			logster.Warn(fmt.Sprintf("Redis unavailable for balance lock (user %s): %v. Proceeding without lock.", userUUID, err))
			lockAcquired = false
			break
		}

		if acquired {
			lockAcquired = true
			logster.Info(fmt.Sprintf("Balance lock acquired for user %s (attempt %d)", userUUID, attempt+1))
			break
		}

		// Lock held by another process - wait and retry
		if attempt < BalanceLockRetries-1 {
			time.Sleep(BalanceLockDelay)
		}
	}

	if !lockAcquired {
		logster.Warn(fmt.Sprintf("Failed to acquire balance lock for user %s after %d retries. Proceeding without lock.", userUUID, BalanceLockRetries))
	}

	// Defer lock release (only if acquired)
	if lockAcquired {
		defer func() {
			released, err := client.Del(ctx, lockKey).Result()
			if err != nil {
				logster.Error(err, fmt.Sprintf("Failed to release balance lock for user %s", userUUID))
			} else if released == 0 {
				logster.Warn(fmt.Sprintf("Balance lock for user %s was already released or expired", userUUID))
			} else {
				logster.Info(fmt.Sprintf("Balance lock released for user %s", userUUID))
			}
		}()
	}

	// Critical section: Read -> Modify -> Write (protected by lock)
	startTime := time.Now()

	// Read current balance from Keycloak
	user, err := service.GetUserById(userUUID, keycloak)
	if err != nil {
		return fmt.Errorf("failed to get user from Keycloak: %w", err)
	}

	// Calculate new balance
	oldBalance := user.Balance
	newBalance := oldBalance + amountToAdd

	// Update Keycloak
	balanceStr := strconv.FormatFloat(newBalance, 'f', 2, 64)
	userUUIDParsed, err := uuid.Parse(userUUID)
	if err != nil {
		return fmt.Errorf("failed to parse user UUID: %w", err)
	}
	_, err = service.UpdateUser(userUUIDParsed, dto.UpdateUserDto{Balance: &balanceStr}, keycloak)
	if err != nil {
		return fmt.Errorf("failed to update user balance in Keycloak: %w", err)
	}

	duration := time.Since(startTime)
	logster.Info(fmt.Sprintf("Balance updated for user %s: %.2f → %.2f (+%.2f) in %dms [lock=%v]",
		userUUID, oldBalance, newBalance, amountToAdd, duration.Milliseconds(), lockAcquired))

	return nil
}
