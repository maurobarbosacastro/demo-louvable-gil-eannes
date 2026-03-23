package main

import (
	"fmt"
	"ms-shopify/internal/auth"
	"ms-shopify/internal/config"
	"ms-shopify/internal/db"
	"ms-shopify/pkg/dotenv"
	"ms-shopify/pkg/logster"
	"ms-shopify/pkg/redisclient"

	_ "ms-shopify/docs"
)

// @title ms-shopify Swagger API
// @version 1.0
// @description Description
// @host localhost:8080
func main() {

	//LOGGER
	currentEnv := dotenv.GetEnv("ENV")
	loggerLevel := dotenv.GetEnv("LOGGER_LEVEL")
	logster.InitLogster(currentEnv, loggerLevel)
	logster.Info(fmt.Sprintf("Env Logger level: %s", loggerLevel))

	dotenv.InitDotenv()
	db.InitDB()
	auth.InitAuth()

	redisclient.SetRedisClient(redisclient.Config{
		Addr:     dotenv.GetEnv("REDIS_HOST"),
		Password: dotenv.GetEnv("REDIS_PASSWORD"),
		DB:       0,
	})

	config.InitServer()
}
