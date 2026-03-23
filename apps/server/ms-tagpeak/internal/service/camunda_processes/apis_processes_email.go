package camunda_processes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/service"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"os"
	"os/signal"
	"syscall"
)

const (
	GetAllUsers = "get-all-users"

	DynamicQuery = "dynamic-query" // Purchase, User, End Date Begin Date
)

func HandleGetAllUsers() {
	fmt.Println("***CAMUNDA*** Start Get All Users")

	w, _ := camundaPKG.StartDynamicWorker(GetAllUsers, func(client worker.JobClient, job entities.Job) {

		getKeycloakAdminToken(auth.KeycloakInstance)

		variables, err := job.GetVariablesAsMap()
		var levelFilter *string

		//Get the membershipLevel from the variables
		if level, ok := variables["level"].(string); ok {
			levelFilter = &level
		}
		fmt.Printf("Filter: %v", levelFilter)

		user, err := service.GetAllUsersUuids(auth.KeycloakInstance, levelFilter)
		fmt.Printf("Users length: %v", len(user))

		if err != nil {
			fmt.Printf("Error -  %v\n", err)
		}

		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(map[string]interface{}{"users": user})
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = request.Send(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Get Email Data worker...")
		w.Close()
	}()

}

func HandlerDynamicCamunda() {
	fmt.Println("***CAMUNDA*** Start Get Cashback With Filters")

	w, _ := camundaPKG.StartDynamicWorker(DynamicQuery, func(client worker.JobClient, job entities.Job) {

		getKeycloakAdminToken(auth.KeycloakInstance)

		var (
			userFilter  *string = nil
			userFilters []string
		)

		variables, err := job.GetVariablesAsMap()

		// This is the query to be executed
		// Required
		query := variables["query"].(string)

		// This is the user filter to be executed
		// Optional
		if user, ok := variables["userFilter"].(string); ok {
			userFilter = &user
		}

		// This is the user filters to be executed
		// Optional
		if userFilter == nil {
			if users, ok := variables["usersFilter"].([]interface{}); ok {
				userFilters = convertInterfaceToString(users)
			}
		}

		res, err := service.DynamicQueryCamunda(query, userFilter, userFilters)

		variablesToSend, err := structToMap(res)
		fmt.Printf("Variables: %v \n", variablesToSend)

		emailVar := map[string]interface{}{
			"email":    variablesToSend,
			"userUuid": userFilter,
		}

		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(emailVar)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = request.Send(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Get Email Data worker...")
		w.Close()
	}()

}

// StructToMap converts a []struct to a map[string]interface{}
func structToMap(data interface{}) (map[string]interface{}, error) {
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON to map[string]interface{}
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func convertInterfaceToString(input []interface{}) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = fmt.Sprint(v)
	}
	return result
}
