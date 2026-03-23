package service

import (
	"encoding/json"
	"fmt"
	"io"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"net/http"
	"strings"
	"time"
)

// GetAwinTransactions Function to get transactions from the API
func GetAwinTransactions(dateType string) (*[]dto.TransactionAwinDTO, error) {
	logster.StartFuncLog()

	// Set up the Variables for the API Call
	awinToken := dotenv.GetEnv("AWIN_TOKEN")
	startDate := time.Now().Add(-7 * 24 * time.Hour).Format("2006-01-02T15:04:05Z")
	endDate := time.Now().Format("2006-01-02T15:04:05Z")

	url := fmt.Sprintf("https://api.awin.com/publishers/1256165/transactions/?accessToken=%s&startDate=%s&endDate=%s&dateType=%s", awinToken, startDate, endDate, dateType)
	logster.Info(fmt.Sprintf("Getting transactions from %s to %s for dateType %s", startDate, endDate, dateType))
	// Awin Test URL Endpoint
	// url := "http://localhost:3721/awin"

	// Make the API request (3 times with a delay of 2 seconds if the API request fails)
	res, err := callAwinAPI(url)
	if err != nil {
		logster.Error(err, "Error API call")
		SendNotificationToDiscord(err.Error())
		return nil, err
	}
	if res == nil {
		logster.Warn("No new transactions")
		return &[]dto.TransactionAwinDTO{}, nil
	}

	// Check response status code
	if res.StatusCode != http.StatusOK {
		logster.Error(nil, fmt.Sprintf("Error status code: %d", res.StatusCode))
		SendNotificationToDiscord(fmt.Sprintf("Error status code: %d", res.StatusCode))
		return &[]dto.TransactionAwinDTO{}, fmt.Errorf("error status code: %d", res.StatusCode)
	}

	// Read the response body
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		logster.Error(err, "Error reading response body")
		return &[]dto.TransactionAwinDTO{}, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the response body to TransactionAwinDTO
	var transactions []dto.TransactionAwinDTO
	err = json.Unmarshal(responseBody, &transactions)
	if err != nil {
		logster.Error(err, "Error parsing JSON data")
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Transaction size %d", len(transactions)))
	return &transactions, nil
}

// callAwinAPI Function to call the Awin API 3 times with a delay of 2 seconds if the API request fails
func callAwinAPI(url string) (*http.Response, error) {
	logster.StartFuncLog()

	delay := 2
	/*
	   , err := strconv.Atoi(dotenv.GetEnv("API_CALL_DELAY"))
	   	if err != nil {
	   		fmt.Printf("Error converting string to int: %v\n", err)
	   		return nil, err
	   	}
	*/

	attempts := 3
	/*
	    err := strconv.Atoi(dotenv.GetEnv("API_ATTEMPTS"))
	   	if err != nil {
	   		fmt.Printf("Error converting string to int: %v\n", err)
	   		return nil, err
	   	}
	*/

	for attempt := 1; attempt <= attempts; attempt++ {

		// Make the API request
		res, err := http.Get(url)
		if err != nil {
			logster.Error(err, fmt.Sprintf("Attempt %d: Failed to make API request", attempt))
		} else {
			// Check response status code
			if res.StatusCode == http.StatusOK {
				fmt.Printf("callAwinAPI - Attempt %d: Success\n", attempt)
				logster.Info(fmt.Sprintf("Attempt %d: Success", attempt))
				logster.EndFuncLog()
				return res, nil // Success
			}

			// Close the response body if the status code is not OK
			err = res.Body.Close()
			if err != nil {
				fmt.Printf("callAwinAPI - Attempt %d: Error closing response body: %v\n", attempt, err)
				logster.Error(err, fmt.Sprintf("Attempt %d", attempt))
				logster.EndFuncLog()
				return nil, err
			}
			logster.Info(fmt.Sprintf("Attempt %d: Received non-OK status code: %d", attempt, res.StatusCode))
		}

		// Wait before retrying (if not the last attempt)
		if attempt < 3 {
			logster.Info(fmt.Sprintf("Retrying in %v", time.Duration(delay)*time.Second))
			time.Sleep(time.Duration(delay) * time.Second)
		}

		if attempt == attempts {
			logster.Error(nil, "All attempts failed (3/3)")
			return nil, fmt.Errorf("All attempts failed  (3/3)")
		}
	}

	logster.EndFuncLog()
	return nil, nil
}

type DiscordContent struct {
	Content string `json:"content"`
}

func SendNotificationToDiscord(message string) {
	logster.StartFuncLog()

	notificationEnabled := dotenv.GetEnv("DISCORD_NOTIFICATION_ENABLED")
	if notificationEnabled == "false" || notificationEnabled == "" {
		logster.Warn("DISCORD_NOTIFICATION_ENABLED not set or false")
		logster.EndFuncLogMsg("Not enabled or env not set")
		return
	}

	discordUrl := dotenv.GetEnv("DISCORD_WEBHOOK")

	if discordUrl == "" {
		logster.Warn("DISCORD_WEBHOOK not set")
		logster.EndFuncLog()
		return
	}

	content := DiscordContent{Content: message}
	var body []byte
	var err error
	body, err = json.Marshal(content)
	if err != nil {
		logster.Error(err, "Error marshalling Discord content")
		logster.EndFuncLog()
		return
	}

	_, errSend := http.Post(
		discordUrl,
		"application/json",
		strings.NewReader(string(body)),
	)

	if errSend != nil {
		logster.Error(err, "Error sending notification to Discord")
		logster.EndFuncLog()
		return
	}
	logster.EndFuncLog()
}
