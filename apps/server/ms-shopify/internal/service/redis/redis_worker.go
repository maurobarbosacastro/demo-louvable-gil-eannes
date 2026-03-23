package redis

import (
	"fmt"
	"ms-shopify/pkg/logster"
	"ms-shopify/pkg/redisclient"
)

const (
	ShopifyRedisQueueKey       = "shopify_orders_queue"
	FailedShopifyRedisQueueKey = "failed_shopify_orders_queue"
	DeleteShopifyRedisQueueKey = "delete_shopify_orders_queue"
)

func PushToRedisQueue(queue string, topic string, shop string, orderData string) {
	client := redisclient.GetRedisClient()
	ctx := redisclient.GetContext()

	stringToPush := fmt.Sprintf("%s|%s|%s", topic, shop, orderData)

	err := client.LPush(ctx, queue, stringToPush).Err()
	if err != nil {
		logster.Error(err, "Error pushing to queue")
		return
	}

	logster.Info(fmt.Sprintf("Pushed to Redis queue %s: %s", queue, stringToPush))
}
