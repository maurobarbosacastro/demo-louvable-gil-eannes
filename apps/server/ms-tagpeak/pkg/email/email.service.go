package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"net/http"
)

type SendEmailDTO struct {
	To         string                 `json:"to"`
	Dictionary map[string]interface{} `json:"dictionary"`
}

type SendEmailAttachment struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
	MimeType string `json:"mimeType"`
}

type SendRawEmailDTO struct {
	To          string                `json:"to"`
	Subject     string                `json:"subject"`
	Body        string                `json:"body"`
	ReplyTo     string                `json:"replyTo,omitempty"`
	Attachments []SendEmailAttachment `json:"attachments,omitempty"`
}

func SendEmail(emailDto SendEmailDTO, template string) (*string, map[string]string) {
	logster.StartFuncLog()

	//Parse the struct to JSON
	jsonData, err := json.Marshal(emailDto)
	if err != nil {
		logster.Error(err, "Error marshaling struct")
		logster.EndFuncLog()
		return nil, map[string]string{"error": "Error marshaling struct"}
	}

	// Create the HTTP request to the email data
	msEmailUrl := dotenv.GetEnv("MS_EMAIL_URL")
	url := msEmailUrl + "api/emails/" + template

	// Create the HTTP request to send the file
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		logster.Error(err, "Error creating HTTP request")
		logster.EndFuncLog()
		return nil, map[string]string{"error": "Could not create HTTP request"}
	}

	// Set Content-Type header for form submission
	req.Header.Set("Content-Type", "application/json")

	logster.Info(fmt.Sprintf("Email url: %v", url))
	logster.Info(fmt.Sprintf("Email data: %v", string(jsonData)))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logster.Error(err, "Error making request")
		logster.EndFuncLog()
		return nil, map[string]string{"error": "Error making request"}
	}
	defer resp.Body.Close()

	logster.EndFuncLogMsg(fmt.Sprintf("Response Status: %v", resp.Status))
	return &resp.Status, nil
}

func SendRawEmail(dto SendRawEmailDTO) (*string, map[string]string) {
	logster.StartFuncLog()

	jsonData, err := json.Marshal(dto)
	if err != nil {
		logster.Error(err, "Error marshaling struct")
		logster.EndFuncLog()
		return nil, map[string]string{"error": "Error marshaling struct"}
	}

	msEmailUrl := dotenv.GetEnv("MS_EMAIL_URL")
	url := msEmailUrl + "api/raw"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		logster.Error(err, "Error creating HTTP request")
		logster.EndFuncLog()
		return nil, map[string]string{"error": "Could not create HTTP request"}
	}

	req.Header.Set("Content-Type", "application/json")

	logster.Info(fmt.Sprintf("Raw email url: %v", url))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logster.Error(err, "Error making request")
		logster.EndFuncLog()
		return nil, map[string]string{"error": "Error making request"}
	}
	defer resp.Body.Close()

	logster.EndFuncLogMsg(fmt.Sprintf("Response Status: %v", resp.Status))
	return &resp.Status, nil
}
