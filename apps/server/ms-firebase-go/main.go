package main

import (
	"fmt"
	"ms-firebase-go/external/firebase"
	"ms-firebase-go/internal/auth"
	"ms-firebase-go/internal/config"
	"ms-firebase-go/internal/db"
	"ms-firebase-go/pkg/dotenv"
	"ms-firebase-go/pkg/logster"

	_ "ms-firebase-go/docs"
)

// @title ms-firebase-go Swagger API
// @version 1.0
// @description Description
// @host localhost:8080
func main() {
	// LOGGER
	currentEnv := dotenv.GetEnv("ENV")
	loggerLevel := dotenv.GetEnv("LOGGER_LEVEL")
	logster.InitLogster(currentEnv, loggerLevel)
	logster.Info(fmt.Sprintf("Env Logger level: %s", loggerLevel))

	dotenv.InitDotenv()
	db.InitDB()
	auth.InitAuth()

	// Initialize Firebase
	firebase.Init()

	// jobs.StartJobs()

	config.InitServer()
}
