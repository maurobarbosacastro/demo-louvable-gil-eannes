package main

import (
	"ms-cj/internal/config"
	"ms-cj/internal/jobs"
	"ms-cj/pkg/dotenv"
	"ms-cj/pkg/logster"
	"ms-cj/pkg/redisclient"
)

func StartRedis() {
	logster.StartFuncLog()

	redisclient.SetRedisClient(redisclient.Config{
		Addr:     dotenv.GetEnv("REDIS_HOST"),
		Password: dotenv.GetEnv("REDIS_PASSWORD"),
		DB:       0,
	})

	logster.EndFuncLog()
}

func main() {
	logster.InitLogster(dotenv.GetEnv("ENV"), dotenv.GetEnv("LOGGER_LEVEL"))
	dotenv.InitDotenv()

	// Start Redis
	StartRedis()

	// Start jobs
	jobs.StartJobs()

	config.InitServer()
}
