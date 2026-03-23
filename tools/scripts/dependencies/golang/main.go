package main

import (
	"fmt"
	"ms-changeme/internal/auth"
	"ms-changeme/internal/config"
	"ms-changeme/internal/db"
	"ms-changeme/pkg/dotenv"
	"ms-changeme/pkg/logster"

	_ "ms-changeme/docs"
)

// @title ms-changeme Swagger API
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

	config.InitServer()
}
