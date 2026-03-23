package camunda

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/pb"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"log"
	"ms-interactive-brokers/pkg/dotenv"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/zbc"
)

type ProcessDefinition struct {
	FileName  string
	ProcessID string
}

// Declare enum values
const (
	RegistrationFlow     = "registration-flow"
	KeycloakRegistration = "keycloak-registration"
	CheckEmailVerified   = "check-email"
	UserMigrationFlow    = "user-migration-flow"
)

var camundaClient *zbc.Client
var camundaToken *string
var parsedResponse map[string]interface{}

var clientId string
var clientSecret string
var tokenUrl string

func InjectEnvOnKey(processID string) string {
	env := dotenv.GetEnv("ENV")

	if env == "pre" || env == "qa" {
		return strings.ToUpper(env) + "_" + processID
	}

	return processID
}

func GetCamundaClientId() string {
	return clientId
}

func GetCamundaClientSecret() string {
	return clientSecret
}

func GetCamundaTokenUrl() string {
	return tokenUrl
}

func SetCamundaClientId(id string) {
	clientId = id
}

func SetCamundaClientSecret(secret string) {
	clientSecret = secret
}

func SetCamundaTokenUrl(url string) {
	tokenUrl = url
}

func StartProcessInstance(processId string, client zbc.Client, variables map[string]interface{}) *pb.CreateProcessInstanceResponse {
	command, err := client.NewCreateInstanceCommand().
		BPMNProcessId(processId).
		LatestVersion().VariablesFromMap(variables)
	if err != nil {
		panic(fmt.Errorf("***CAMUNDA*** failed to create process instance command for client [%s]", client))
	}

	process, err := command.Send(context.Background())
	if err != nil {
		fmt.Printf("***CAMUNDA*** Error creating user register: %s \n", err)
	}

	if err != nil {
		fmt.Printf("***CAMUNDA*** Error creating user register: %s \n", err)
	}

	fmt.Printf("***CAMUNDA*** started process instance [%d] with {\"%s\": \"%v\"}\n", process.GetProcessInstanceKey(), "user", variables)
	return process
}

// StartWorker creates a new Zeebe worker for a task type
func StartWorker(taskType string) (worker.JobWorker, entities.Job) {

	fmt.Printf("***CAMUNDA*** Start Worker with taskType: %s\n", taskType)

	c := *camundaClient

	var usedJob entities.Job

	jobWorker := c.NewJobWorker().
		JobType(InjectEnvOnKey(taskType)).
		Handler(func(client worker.JobClient, job entities.Job) {
			usedJob = job
			fmt.Printf("***CAMUNDA*** Activated job with key: %d\n", job.Key)

			_, err := client.NewCompleteJobCommand().JobKey(job.Key).Send(context.Background())
			if err != nil {
				return
			}
		}).
		Open()

	return jobWorker, usedJob
}

func StartDynamicWorker(taskType string, jobHandler func(client worker.JobClient, job entities.Job)) (worker.JobWorker, entities.Job) {
	fmt.Printf("***CAMUNDA*** Start Worker with taskType: %s\n", taskType)

	c := *camundaClient

	var usedJob entities.Job

	jobWorker := c.NewJobWorker().
		JobType(InjectEnvOnKey(taskType)).
		Handler(jobHandler).
		Open()

	return jobWorker, usedJob
}

func CreateClient() zbc.Client {
	fmt.Printf(" Creating client with clientId %s\n", clientId)
	audience := dotenv.GetEnv("ZEEBE_AUDIENCE")

	credentials, err := zbc.NewOAuthCredentialsProvider(&zbc.OAuthProviderConfig{
		ClientID:               clientId,
		ClientSecret:           clientSecret,
		AuthorizationServerURL: tokenUrl,
		Audience:               audience,
		Scope:                  "openid",
	})

	if err != nil {
		log.Fatalf("Failed to create OAuth credentials provider: %v", err)
	}

	config := zbc.ClientConfig{
		GatewayAddress:         dotenv.GetEnv("ZEEBE_GATEWAY_ADDRESS"),
		CredentialsProvider:    credentials,
		UsePlaintextConnection: true,
	}

	log.Printf("Client configured with GatewayAddress: %s and AuthorizationServerURL: %s",
		config.GatewayAddress, tokenUrl)

	client, err := zbc.NewClient(&config)
	if err != nil {
		panic(err)
	}
	camundaClient = &client
	return client
}

func CreateClientLocal() zbc.Client {
	fmt.Printf(" Creating client with clientId %s\n", clientId)

	config := zbc.ClientConfig{
		GatewayAddress:         "localhost:26500",
		UsePlaintextConnection: true,
	}

	log.Printf("Client configured with GatewayAddress: %s and AuthorizationServerURL: %s",
		config.GatewayAddress, tokenUrl)

	client, err := zbc.NewClient(&config)
	if err != nil {
		panic(err)
	}
	camundaClient = &client
	return client
}

func GetCamundaClient() *zbc.Client {
	return camundaClient
}

func DeployProcesses(client zbc.Client, processes []ProcessDefinition) {
	for _, process := range processes {
		DeployProcess(client, process.FileName, process.ProcessID)
	}
}

func DeployProcess(client zbc.Client, fileName, processID string) {
	definition := ReadFile(fileName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	command := client.NewDeployResourceCommand().AddResource(definition, fileName)
	response, err := command.Send(ctx)
	if err != nil {
		fmt.Printf("Deployment failed for process [%s]: %s\n", fileName, err)
	}

	for _, deployment := range response.Deployments {
		fmt.Printf("Deployed process [%s] with key [%d] from file [%s]\n", deployment.GetProcess().BpmnProcessId, deployment.GetProcess().ProcessDefinitionKey, fileName)
	}
}

// Used to know if zeebe is connected and ready to receive jobs
func WaitForZeebeReadiness(client zbc.Client, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		_, err := client.NewTopologyCommand().Send(ctx)
		if err == nil {
			fmt.Printf("*** ZEEBE IS READY *** clientId: %s\n", clientId)
			return nil
		}

		log.Printf("Waiting for Zeebe readiness: %v", err)
		time.Sleep(2 * time.Second) // Retry after 2 seconds
	}
}

// Function to load dinamically all BPMN files from the resources/bpmn directory
func LoadProcessDefinitionsFromDir() []ProcessDefinition {
	var definitions []ProcessDefinition
	env := dotenv.GetEnv("ENV")

	// Get the current working directory of the application
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("***CAMUNDA*** Error getting current directory: %s", err)
	}

	var bpmnDir string
	// Specify the relative path to the BPMN directory from the current working directory
	// If env is pre or qa, the BPMN files are in the resources/bpmn/pre or resources/bpmn/qa directory
	// If env is dev/prod, the BPMN files are in the resources/bpmn directory
	if env == "pre" || env == "qa" {
		bpmnDir = filepath.Join(dir, "resources", "bpmn", env)
	} else {
		bpmnDir = filepath.Join(dir, "resources", "bpmn", "default")
	}

	err = filepath.Walk(bpmnDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".bpmn" {
			processID := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
			definitions = append(definitions, ProcessDefinition{
				FileName:  path,
				ProcessID: processID,
			})
		}
		return nil
	})

	if err != nil {
		log.Fatalf("***CAMUNDA*** Error loading BPMN files from directory [%s]: %s", bpmnDir, err)
	}

	log.Printf("***CAMUNDA*** Loaded [%d] BPMN files from directory [%s]", len(definitions), bpmnDir)
	return definitions
}

func ReadFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file [%s]: %s", filePath, err)
	}
	return data
}

func GetToken(clientID, clientSecret string) (string, error) {

	// Prepare the form data
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")

	// Create the HTTP request
	req, err := http.NewRequest("POST", tokenUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	parsedResponse, err = parseResponse(resp)

	accessToken := parsedResponse["access_token"].(string)

	fmt.Println("Got token successfully")
	// Return the accessToken as a string
	return accessToken, nil
}

func SetCamundaToken(token string) {
	camundaToken = &token
}

func GetCamundaToken() *string {
	return camundaToken
}

// Aux function to parse response to map to extract the access token
func parseResponse(resp *http.Response) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	defer resp.Body.Close()

	// Parse JSON into a map
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return result, nil
}

func GetVariables(job entities.Job) (map[string]interface{}, error) {

	res := map[string]interface{}{}

	variables := job.GetVariables()

	err := json.Unmarshal([]byte(variables), &res)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}

func VerifyCamundaToken() {

	parsedToken, err := jwt.Parse(*camundaToken, func(token *jwt.Token) (interface{}, error) {
		return token.Raw, nil
	})

	// If there was an error parsing or the token is invalid, get a new token
	if err != nil || !parsedToken.Valid {
		fmt.Println("Token invalid or expired. Fetching a new one...")

		// Fetch a new token
		newToken, err := GetToken(GetCamundaClientId(), GetCamundaClientSecret())
		if err != nil {
			fmt.Println("Error getting new Camunda token:", err)
			return
		}

		// Set the new token
		camundaToken = &newToken
		return
	}

	// If the token is valid, continue with the same token
	fmt.Println("Token is valid.")

}
