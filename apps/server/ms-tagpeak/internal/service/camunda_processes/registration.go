package camunda_processes

import (
	"context"
	"fmt"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
	"ms-tagpeak/internal/service"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"os"
	"os/signal"
	"syscall"
)

const (
	HaveTransactions = "have-transactions"
)

func HandleHaveTransactions() {
	fmt.Println("***CAMUNDA*** Start Have Transactions")

	w, _ := camundaPKG.StartDynamicWorker(camundaPKG.InjectEnvOnKey(HaveTransactions), func(client worker.JobClient, job entities.Job) {

		//GET Rewards by transaction uuid
		variables, err := job.GetVariablesAsMap()

		userUUID := variables["userUuid"].(string)
		haveTransactions, err := service.UserHaveTransactions(userUUID)

		variablesToSend := map[string]interface{}{
			"haveTransactions": haveTransactions,
		}

		request, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromMap(variablesToSend)
		if err != nil {
			return
		}

		_, err = request.Send(context.Background())
		if err != nil {
			return
		}

	})

	// Set up a barrier for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Have Transactions worker...")
		w.Close()
	}()

}
