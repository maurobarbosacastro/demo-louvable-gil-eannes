package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ms-firebase-go/pkg/dotenv"
	"net/http"
)

type SendEmailDTO struct {
	To         string                 `json:"to"`
	Dictionary map[string]interface{} `json:"dictionary"`
}

func SendEmail(emailDto SendEmailDTO, template string) (*string, map[string]string) {

	//Parse the struct to JSON
	jsonData, err := json.Marshal(emailDto)
	if err != nil {
		fmt.Printf("Error marshaling struct: %v\n", err)
		return nil, map[string]string{"error": "Error marshaling struct"}
	}

	// Create the HTTP request to the email data
	msEmailUrl := dotenv.GetEnv("MS_EMAIL_URL")
	url := msEmailUrl + "api/emails/" + template
	// Create the HTTP request to send the file
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, map[string]string{"error": "Could not create HTTP request"}
	}

	// Set Content-Type header for form submission
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nil, map[string]string{"error": "Error making request"}
	}
	defer resp.Body.Close()

	fmt.Printf("Response Status: %v\n", resp.Status)
	return &resp.Status, nil
}
