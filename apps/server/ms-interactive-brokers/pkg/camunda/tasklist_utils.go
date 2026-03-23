package camunda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"ms-interactive-brokers/internal/dto"
	"ms-interactive-brokers/pkg/dotenv"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	CREATED   = "CREATED"
	COMPLETED = "COMPLETED"
	CANCELED  = "CANCELED"
	FAILED    = "FAILED"
)

func CompleteTask(taskId string) (bool, error) {
	fmt.Printf("***CAMUNDA*** Starting process to complete task %s\n", taskId)
	// Check if camundaToken is valid
	if camundaToken == nil || *camundaToken == "" {
		log.Fatal("Error: camundaToken is nil or empty")
	}

	url := dotenv.GetEnv("TASKLIST_URL") + "/v1/tasks/" + taskId + "/complete"
	method := "PATCH"

	// Adjust payload if necessary
	payload := strings.NewReader(`{"variables": [] }`)

	// Set HTTP client with timeout
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Printf("Failed to create HTTP request: %v", err)
		return false, err
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+*camundaToken)

	// Perform the request
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to complete HTTP request: %v", err)
		return false, err
	}
	defer res.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return false, err
	}

	log.Printf("***CAMUNDA*** Completed task. Response: %s", string(body))
	return true, nil
}

func GetTaskByStateAndProcessInstanceKey(state string, processInstanceKey int64) (dto.TaskDTO, error) {
	fmt.Printf("***CAMUNDA*** Get task by state %s and process instance key %d\n", state, processInstanceKey)

	url := dotenv.GetEnv("TASKLIST_URL") + "/v1/tasks/search"
	method := "POST"

	// Create a map for the payload
	payloadData := map[string]string{
		"state":              state,
		"processInstanceKey": strconv.FormatInt(processInstanceKey, 10),
	}

	// Convert the map to JSON
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Failed to marshal payload:", err)
		return dto.TaskDTO{}, nil
	}

	// Verify the token and if invalid, get a new one
	VerifyCamundaToken()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payloadBytes))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return dto.TaskDTO{}, err
	}

	// Check if camundaToken is set
	if camundaToken == nil || *camundaToken == "" {
		fmt.Println("Error: camundaToken is nil or empty")
		return dto.TaskDTO{}, nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+*camundaToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return dto.TaskDTO{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return dto.TaskDTO{}, err
	}

	//Add this to check if the task is empty and print the error returned by Camunda
	var task []interface{}
	err = json.Unmarshal(body, &task)
	fmt.Println("Task: ", task)

	// The error here is not very specific
	var taskDTOs []dto.TaskDTO
	err = json.Unmarshal(body, &taskDTOs)
	if err != nil {
		return dto.TaskDTO{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(taskDTOs) == 0 {
		return dto.TaskDTO{}, fmt.Errorf("no tasks found in the response")
	}

	fmt.Printf("***CAMUNDA*** Found %d tasks\n", len(taskDTOs))
	// Return the first task (or handle as per your requirement)
	return taskDTOs[0], nil
}
