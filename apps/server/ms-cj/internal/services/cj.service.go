package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"ms-cj/internal/dto"
	"ms-cj/internal/responses"
	"ms-cj/pkg/dotenv"
	"ms-cj/pkg/logster"
	"ms-cj/pkg/utils"
	"net/http"
	"time"
)

type CJAffiliateClient struct {
	APIToken   string
	BaseURL    string
	HTTPClient *http.Client
}

func CjAffiliateClient(apiToken string, url string) *CJAffiliateClient {
	return &CJAffiliateClient{
		APIToken:   apiToken,
		BaseURL:    url,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func executeQuery[T any](c *CJAffiliateClient, query string) (*T, error) {
	logster.StartFuncLog()

	// Prepare request body
	reqBody := dto.GraphQLRequest{
		Query: query,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check HTTP status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	// Parse GraphQL response
	var gqlResp responses.CJAffiliateResponse[T]
	if err := json.Unmarshal(body, &gqlResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check for GraphQL errors
	if len(gqlResp.Errors) > 0 {
		return &gqlResp.Data, fmt.Errorf("GraphQL errors: %+v", gqlResp.Errors)
	}

	logster.EndFuncLog()
	return &gqlResp.Data, nil
}

func GetCjTransaction(dtoParam dto.CjTransactionDTO) (*responses.CJAffiliateData, error) {
	logster.StartFuncLog()

	// Create a new instance of the CJAffiliateClient
	client := CjAffiliateClient(dtoParam.ApiToken, dotenv.GetEnv("TRANSACTIONS_URL"))

	// Create a custom query based on the dtoParam
	customQuery := utils.MapCustomCommissionQuery(dtoParam)

	// Execute the custom query
	resp, err := executeQuery[responses.CJAffiliateData](client, customQuery)
	if err != nil {
		return nil, err
	}

	logster.EndFuncLog()
	return resp, nil
}

func GetAdIdStore(dto dto.AdIdDTO) (*responses.CJAffiliateDataAid, error) {
	logster.StartFuncLog()

	// Create a new instance of the CJAffiliateClient
	client := CjAffiliateClient(dto.ApiToken, dotenv.GetEnv("ADID_URL"))

	// Create a query based on the dto
	query := utils.MapCustomAdIdQuery(dto)

	// Execute the custom query
	resp, err := executeQuery[responses.CJAffiliateDataAid](client, query)
	if err != nil {
		return nil, err
	}

	logster.EndFuncLog()
	return resp, nil
}
