package redis

import (
	"encoding/json"
	"fmt"
	"ms-cj/internal/responses"
	"ms-cj/pkg/logster"
	"ms-cj/pkg/redisclient"
)

const (
	CjTransactionKey = "cj_transaction_queue"
)

func PushTransactionToRedisQueue(queue string, transaction *responses.CommissionRecord) {
	client := redisclient.GetRedisClient()
	ctx := redisclient.GetContext()

	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		logster.Error(err, "Error marshaling transaction to JSON")
		return
	}

	err = client.LPush(ctx, queue, transactionJSON).Err()
	if err != nil {
		logster.Error(err, "Error pushing to queue")
		return
	}

	logster.Info(fmt.Sprintf("Pushed to Redis queue %s: %v", queue, transaction))
}
