package camunda

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func HandleKeycloakRegistration() {
	fmt.Println("***CAMUNDA*** Start Keycloak Registration")
	w, _ := StartWorker(KeycloakRegistration)

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Keycloak Registration worker...")
		w.Close()
	}()
}

func HandleEmailVerification() {
	fmt.Println("***CAMUNDA*** Start Email Verification")
	w, _ := StartWorker(CheckEmailVerified)

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Email Verification worker...")
		w.Close()
	}()
}
