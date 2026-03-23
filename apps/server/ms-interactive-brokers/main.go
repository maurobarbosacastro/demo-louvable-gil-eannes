package main

import (
	"fmt"
	"ms-interactive-brokers/internal/auth"
	"ms-interactive-brokers/internal/config"
	"ms-interactive-brokers/internal/service/ibkr"
	"ms-interactive-brokers/pkg/dotenv"
	"ms-interactive-brokers/pkg/logster"

	_ "ms-interactive-brokers/docs"
)

// @title ms-interactive-brokers Swagger API
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
	auth.InitAuth()

	// Load configuration
	ibkrConfig, err := ibkr.LoadConfig()
	if err != nil {
		logster.Fatal(err, "Failed to load configuration")
	}

	// Validate access token and secret
	if ibkrConfig.AccessToken == "" || ibkrConfig.AccessTokenSecret == "" {
		logster.Fatal(nil, "Please set the access token and access token secret generated in the self-service portal. "+
			"The self-service portal can be found at: https://ndcdyn.interactivebrokers.com/sso/Login?action=OAUTH&RL=1&ip2loc=US")
	}
	ibkr.Appconfig = ibkrConfig

	config.InitServer()
}
