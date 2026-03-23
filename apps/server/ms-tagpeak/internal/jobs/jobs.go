package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"github.com/samber/lo"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/files"
	"ms-tagpeak/pkg/logster"
	"strconv"
	"time"
)

func StartJobs() {
	logster.Info("Starting jobs...")

	loc, err := time.LoadLocation("Europe/Lisbon")
	if err != nil {
		logster.Error(err, "Error loading timezone")
	}
	logster.Info(fmt.Sprintf("Timezone: %s", loc.String()))

	go getCurrencyExchangeRateFromAPIJob(loc)
	go checkValidIbansToDeleteFile(loc)
	go updateRewardsHistory(loc)
	go updateShopifyActionableTransactionsJob(loc)
}

func getCurrencyExchangeRateFromAPIJob(location *time.Location) {
	logster.StartFuncLog()

	timer := dotenv.GetEnv("CURRENCY_EXCHANGE_RATE_JOB")

	c := cron.NewWithLocation(location)

	// Error handling for cron job addition
	err := c.AddFunc(timer, func() {
		logster.Info("getCurrencyExchangeRateFromAPIJob Started")

		fixerApiKey := dotenv.GetEnv("FIXER_API_KEY")

		//CALL API TO GET CURRENCY EXCHANGE RATES in FixerDTO
		rates, err := service.GetFixerExchangeRates(fixerApiKey)
		if err != nil {
			logster.Error(err, "Error getting Fixer API rates")
			return
		}

		//Marshal the rates to string
		ratesToString, err := json.Marshal(rates.Rates)
		if err != nil {
			logster.Error(err, "Error marshaling rates")
			return
		}

		//Create currency exchange rate
		c := dto.CreateCurrencyExchangeRateDTO{
			Base:  rates.Base,
			Rates: string(ratesToString),
		}

		//Save currency exchange rate
		_, err = service.CreateCurrencyExchangeRate(c)
		if err != nil {
			logster.Error(err, "Error creating currency exchange rate")
			return
		}

		logster.Info("getCurrencyExchangeRateFromAPIJob Ended")
	})
	if err != nil {
		logster.Error(err, "Error adding cron job getCurrencyExchangeRateFromAPIJob")
		return
	}

	// Start the cron scheduler
	c.Start()

	select {}
}

func checkValidIbansToDeleteFile(location *time.Location) {
	logster.StartFuncLog()

	timer := dotenv.GetEnv("IBAN_CHECK_TIMER")
	logster.Info("Timer: " + timer)

	c := cron.NewWithLocation(location)

	err := c.AddFunc(timer, func() {
		logster.Info("checkValidIbansToDeleteFile Started")

		ibanReplacementUuid := dotenv.GetEnv("IBAN_FILE_REPLACEMENT_UUID")
		logster.Info("ibanReplacementUuid: " + ibanReplacementUuid)

		config, err := service.GetConfiguration("user_payment_iban_valid_time")
		if err != nil {
			logster.Error(err, "Error getting user_payment_iban_valid_time configuration")
			return
		}
		maxDays, _ := strconv.Atoi(config.Value)

		now := time.Now()
		maxTime := now.AddDate(0, 0, -maxDays)
		logster.Info("Max Time: " + maxTime.String())

		//Call function on the database that will set the file_uuid to another value for
		//all user_payment_method with state "VALIDATED" and older than x(maxDays) days
		replacedFiles, err := service.ReplaceIbanStatementFile(ibanReplacementUuid, maxDays)
		if err != nil {
			logster.Error(err, "Error updating")
			return
		}
		logster.Info(fmt.Sprintf("Replaced files: %#v", lo.Map(replacedFiles, func(file models.AffectedRecord, _ int) string { return file.OriginalFileUUID.String() })))

		//Delete current iban file and add a new "confidential" one
		for _, file := range replacedFiles {
			err = files.DeleteFile(file.OriginalFileUUID.String())
			if err != nil {
				logster.Error(err, fmt.Sprintf("Error deleting file %s", file.OriginalFileUUID.String()))
				return
			}
		}

		logster.Info("checkValidIbansToDeleteFile End")
	})

	if err != nil {
		logster.Error(err, "Error adding cron job checkValidIbansToDeleteFile")
		return
	}

	// Start the cron scheduler
	c.Start()

	select {}
}

func updateRewardsHistory(location *time.Location) {
	logster.StartFuncLog()

	timer := dotenv.GetEnv("IBKR_UPDATE_TIMER")
	logster.Info("Timer: " + timer)

	c := cron.NewWithLocation(location)

	err := c.AddFunc(timer, func() {
		logster.Info("updateRewardsHistory Started")
		keycloak := auth.KeycloakInstance

		service.UpdateRewardsHistory(keycloak)

		logster.Info("updateRewardsHistory Ended")
	})

	if err != nil {
		logster.Error(err, "Error adding cron job updateRewardsHistory")
		return
	}

	c.Start()
	select {}
}

func updateShopifyActionableTransactionsJob(location *time.Location) {
	logster.StartFuncLog()

	timer := dotenv.GetEnv("SHOPIFY_ACTIONABLE_TRANSACTIONS_UPDATE_TIMER")
	logster.Info("Timer: " + timer)

	c := cron.NewWithLocation(location)

	err := c.AddFunc(timer, func() {
		logster.Info("updateShopifyActionableTransactionsJob Started")

		err := service.UpdateShopifyActionableTransactionsState()
		if err != nil {
			logster.Error(err, "Error updating shopify actionable transactions state")
			logster.Info("updateShopifyActionableTransactionsJob Ended")
			return
		}

		logster.Info("updateShopifyActionableTransactionsJob Ended")
	})

	if err != nil {
		logster.Error(err, "Error adding cron job updateShopifyActionableTransactionsJob")
		return
	}

	c.Start()
	select {}
}
