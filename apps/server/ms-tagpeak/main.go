package main

import (
	"fmt"
	_ "ms-tagpeak/docs"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/config"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/jobs"
	"ms-tagpeak/internal/service"
	camundaService "ms-tagpeak/internal/service/camunda_processes"
	"ms-tagpeak/internal/service/redis"
	"ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/redisclient"
	"time"
)

func StartCamunda() {
	logster.StartFuncLog()

	camunda.SetCamundaClientId(dotenv.GetEnv("ZEEBE_CLIENT_ID"))
	camunda.SetCamundaClientSecret(dotenv.GetEnv("ZEEBE_CLIENT_SECRET"))
	camunda.SetCamundaTokenUrl(dotenv.GetEnv("ZEEBE_AUTHORIZATION_SERVER_URL"))

	// Initialize the Zeebe client
	camunda.CreateClient()
	//camunda.CreateClientLocal() // CALL THIS TO USE CAMUNDA LOCALLY
	camundaClient := camunda.GetCamundaClient()

	if camundaClient == nil {
		logster.Fatal(nil, "Failed to initialize Zeebe client")
	}

	// Get and set the token for tasklist operations
	token, err := camunda.GetToken(camunda.GetCamundaClientId(), camunda.GetCamundaClientSecret())
	if err != nil {
		logster.Fatal(err, "Failed to get Camunda token")
	}
	camunda.SetCamundaToken(token)

	// Wait for Zeebe readiness
	err = camunda.WaitForZeebeReadiness(*camundaClient, 30*time.Second)
	if err != nil {
		logster.Fatal(err, "Zeebe readiness check failed")
	}

	// Deploy all BPMN files
	camunda.DeployProcesses(*camundaClient, camunda.LoadProcessDefinitionsFromDir())

	// Start workers in separate goroutines
	go camunda.HandleKeycloakRegistration()
	go camunda.HandleEmailVerification()
	go camundaService.HandleHaveTransactions()

	//***CASHBACK TRANSACTIONS***
	// -> JOB Check Transactions
	go camundaService.HandleCheckTransactions()
	go camundaService.HandleCheckTransactionIfNew()
	go camundaService.HandleCreateNewTransaction()
	go camundaService.HandleConvertToTagpeakCurrency()
	go camundaService.HandleUpdateTransaction()
	go camundaService.HandleCheckUserMembershipLevel()
	go camundaService.HandleCheckNotificationSend()

	// -> JOB Update Expired Rewards
	go camundaService.HandleCheckInvestments()
	go camundaService.HandleUpdateInvestment()

	// -> Start Investment
	go camundaService.HandleCheckIfInvestmentExists()
	go camundaService.HandleTransactionChangeState()

	// -> End Investment
	go camundaService.HandleNotifyTagPeak()
	go camundaService.HandleUpdateReward()
	go camundaService.HandleCreditUser()

	// -> Credit Referral
	go camundaService.HandleCreateNewRewardReferral()
	go camundaService.HandleCreditUserReferral()

	// -> Credit Influencer Referral
	go camundaService.HandleGetDataCreditInfluencer()
	go camundaService.HandleCreditNewInfluencerReward()

	// -> Notifications
	go camundaService.HandleGetEmailData()
	go camundaService.HandleSendEmailAction()
	go camundaService.HandleSendNotificationAction()

	//Dynamic Handlers
	go camundaService.HandlerDynamicCamunda()
	go camundaService.HandleGetAllUsers()

	logster.EndFuncLog()
}

func StartRedis() {
	logster.StartFuncLog()

	redisclient.SetRedisClient(redisclient.Config{
		Addr:     dotenv.GetEnv("REDIS_HOST"),
		Password: dotenv.GetEnv("REDIS_PASSWORD"),
		DB:       0,
	})

	go redis.ProcessCjTransactionsRedisQueue()
	go redis.ProcessShopifyOrderQueue()
	go redis.ProcessShopifyDeleteQueue()

	logster.EndFuncLog()
}

func main() {

	// Initialize environment and services
	dotenv.InitDotenv()

	currentEnv := dotenv.GetEnv("ENV")
	loggerLevel := dotenv.GetEnv("LOGGER_LEVEL")
	logster.InitLogster(currentEnv, loggerLevel)
	logster.Info(fmt.Sprintf("Env Logger level: %s", loggerLevel))

	db.InitDB()
	auth.InitAuth()
	service.LoadConfigurations()
	service.InitConfigs()

	// Update al users membership level after migration. SHOULD ONLY RUN ONCE -> TO BE REMOVED AFTER MIGRATION
	updateMembershipLevel := dotenv.GetEnv("UPDATE_MEMBERSHIP_LEVEL")
	logster.Warn("Update membership level: " + updateMembershipLevel)

	if updateMembershipLevel == "true" {
		//go service.ManageAllUsersMembershipLevel(auth.KeycloakInstance)
		updateMembershipLevel = "false"
	}

	// Start Redis
	StartRedis()

	//Start Camunda
	go StartCamunda()

	// Start the HTTP server (last step)
	jobs.StartJobs()

	//Must be last instruction!!
	config.InitServer()

}
