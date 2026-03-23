package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetCurrencyExchangeRate(uuid uuid.UUID) (*dto.CurrencyExchangeRateDTO, error) {
	res, err := repository.GetCurrencyExchangeRate(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.CustomErrorStruct{}.NotFoundError("Currency Exchange Rate", uuid)
		}
		return nil, err
	}

	// Convert the rates from json.RawMessage to map[string]float64
	// to be able to get the Rate by currency currencyExchangeRateDTO.Rates[currency] for example
	var rawJsonToMap map[string]float64
	err = json.Unmarshal(json.RawMessage(res.Rates), &rawJsonToMap)
	if err != nil {
		return nil, err
	}

	// Return the currency exchange rate DTO
	currencyExchangeRateDTO := dto.CurrencyExchangeRateDTO{
		Uuid:      &res.Uuid,
		Base:      *res.Base,
		Rates:     rawJsonToMap,
		CreatedAt: res.CreatedAt,
	}

	return &currencyExchangeRateDTO, nil
}

func GetAllCurrencyExchangeRates() (*[]models.CurrencyExchangeRate, error) {
	res, err := repository.GetAllCurrencyExchangeRates()
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func CreateCurrencyExchangeRate(dtoParam dto.CreateCurrencyExchangeRateDTO) (*models.CurrencyExchangeRate, error) {
	model := utils.CurrencyExchangeRateDtoToModel(&dtoParam)
	model.CreatedBy = "system"

	res, err := repository.CreateCurrencyExchangeRate(model)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func UpdateCurrencyExchangeRate(dtoParam dto.UpdateCurrencyExchangeRateDTO, uuid uuid.UUID, uuidUser string) (*models.CurrencyExchangeRate, error) {
	toUpdate, err := repository.GetCurrencyExchangeRate(uuid)
	if err != nil {
		return nil, err
	}

	if dtoParam.Base != nil {
		toUpdate.Base = dtoParam.Base
	}
	if dtoParam.Rates != nil {
		toUpdate.Rates = *dtoParam.Rates
	}

	toUpdate.UpdatedBy = &uuidUser

	res, err := repository.UpdateCurrencyExchangeRate(toUpdate)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func DeleteCurrencyExchangeRate(uuid uuid.UUID, uuidUser string) error {
	err := repository.DeleteCurrencyExchangeRate(uuid, uuidUser)
	if err != nil {
		return err
	}
	return nil
}

// GetLatestCurrencyExchangeRateFromDatabase Get The latest currency exchange rate From database
func GetLatestCurrencyExchangeRateFromDatabase() (*dto.CurrencyExchangeRateDTO, error) {
	logster.StartFuncLog()

	// Get the latest currency exchange rate from the database
	res, err := repository.GetLatestCurrencyExchangeRate()
	if err != nil {
		return nil, err
	}

	currencyExchangeRateDTO, err := convertRatesFromJsonToMap(*res)
	if err != nil {
		return nil, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("From: %s", res.CreatedAt.String()))
	return currencyExchangeRateDTO, nil
}

func convertRatesFromJsonToMap(res models.CurrencyExchangeRate) (*dto.CurrencyExchangeRateDTO, error) {
	// Convert the rates from json.RawMessage to map[string]float64
	// to be able to get the Rate by currency currencyExchangeRateDTO.Rates[currency] for example
	var rawJsonToMap map[string]float64
	err := json.Unmarshal(json.RawMessage(res.Rates), &rawJsonToMap)
	if err != nil {
		return nil, err
	}

	// Return the currency exchange rate DTO
	currencyExchangeRateDTO := dto.CurrencyExchangeRateDTO{
		Uuid:      &res.Uuid,
		Base:      *res.Base,
		Rates:     rawJsonToMap,
		CreatedAt: res.CreatedAt,
	}

	return &currencyExchangeRateDTO, nil
}

// callFixerAPI Function to call the Fixer API 3 times with a delay of 2 seconds if the API request fails
func callFixerAPI(apiKey string) (*http.Response, error) {
	delay, err := strconv.Atoi(dotenv.GetEnv("FIXER_DELAY"))
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
		return nil, err
	}

	attempts, err := strconv.Atoi(dotenv.GetEnv("FIXER_ATTEMPTS"))
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
		return nil, err
	}

	// API Request
	url := fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%s", apiKey)
	// res, err := http.Get(url)

	for attempt := 1; attempt <= attempts; attempt++ {
		// Make the API request
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("Attempt %d: Failed to make API request: %v\n", attempt, err)
		} else {
			// Check response status code
			if res.StatusCode == http.StatusOK {
				return res, nil // Success
			}

			// Close the response body if the status code is not OK
			err := res.Body.Close()
			if err != nil {
				return nil, err
			}
			fmt.Printf("Attempt %d: Received non-OK status code: %d\n", attempt, res.StatusCode)
		}

		// Wait before retrying (if not the last attempt)
		if attempt < 3 {
			fmt.Printf("Retrying in %v...\n", time.Duration(delay)*time.Second)
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}

	return nil, nil
}

// GetFixerExchangeRates Function to get exchange rates from the API
func GetFixerExchangeRates(apiKey string) (dto.FixerDTO, error) {
	// API Request
	res, err := callFixerAPI(apiKey)

	if res.StatusCode != http.StatusOK {
		return dto.FixerDTO{}, fmt.Errorf("error status code: %d", res.StatusCode)
	}

	// Read the response body
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return dto.FixerDTO{}, fmt.Errorf("error reading response body: %w", err)
	}

	// Unmarshal the response body to FixerDTO
	var fixerResponse dto.FixerDTO
	err = json.Unmarshal(responseBody, &fixerResponse)
	if err != nil {
		return dto.FixerDTO{}, fmt.Errorf("error decoding JSON response: %w", err)
	}

	return fixerResponse, nil
}
