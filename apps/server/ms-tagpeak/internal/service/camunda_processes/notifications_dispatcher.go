package camunda_processes

import (
	"context"
	"fmt"
	"ms-tagpeak/external/notifications"
	"ms-tagpeak/internal/service"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
	"github.com/google/uuid"
)

const (
	SendNotificationAction = "send-notification-action"
)

func HandleSendNotificationAction() {
	fmt.Println("***CAMUNDA*** Start Send Notification Action")

	w, _ := camundaPKG.StartDynamicWorker(SendNotificationAction,
		func(client worker.JobClient, job entities.Job) {
			logster.StartFuncLogMsg(job.ProcessInstanceKey)

			variables, err := job.GetVariablesAsMap()
			if err != nil {
				logster.Error(err, "Error getting variables")
				logster.EndFuncLog()
				return
			}
			logster.Debug(fmt.Sprintf("Variables: %+v", variables))

			rewardUuid := variables["rewardUuid"].(string)
			rewardOwner := variables["rewardOwner"].(string)
			state := variables["investmentState"].(string)
			logster.Info(fmt.Sprintf("rewardUuid: %s | rewardOwner: %s", rewardUuid, rewardOwner))

			reward, errGetReward := service.GetReward(uuid.MustParse(rewardUuid))
			if errGetReward != nil {
				logster.Error(errGetReward, "Error getting reward")
				logster.EndFuncLog()
			}

			transaction, errGetTransaction := service.GetTransactionByReward(uuid.MustParse(rewardUuid))
			if errGetTransaction != nil {
				logster.Error(errGetTransaction, "Error getting reward")
				logster.EndFuncLog()
			}
			storeVisit, errGetStoreVisit := service.GetStoreVisitByTransaction(transaction.Uuid)
			if errGetStoreVisit != nil {
				logster.Error(errGetStoreVisit, "Error getting store visit")
				logster.EndFuncLog()
			}

			var notificationsContent notifications.NotificationContent
			switch state {
			case "LIVE":
				notificationsContent = *notifications.GetNotificationContent(state, storeVisit.Reference, storeVisit.Store.Name)
			case "TRACKED":
				notificationsContent = *notifications.GetNotificationContent(state, storeVisit.Store.Name)
			case "FINISHED":
				notificationsContent = *notifications.GetNotificationContent(state, storeVisit.Store.Name, storeVisit.Reference, reward.CurrentRewardUser)
			case "STOPPED":
				notificationsContent = *notifications.GetNotificationContent(state)
			}

			notification := notifications.CreateNotificationDto{
				Title:      notificationsContent.Title,
				Content:    notificationsContent.Body,
				UserTarget: utils.Ptr([]string{rewardOwner}),
				Date:       time.Now(),
				State:      utils.Ptr(string(notifications.NotificationStateScheduled)),
			}

			notifCreated, errNotifCreated := notifications.CreateNotification(notification)
			if errNotifCreated != nil {
				logster.Error(errNotifCreated, "Error creating notification")
				logster.EndFuncLog()
			}

			logster.Debug(fmt.Sprintf("Notification created: %+v", notifCreated))

			notifications.SendNotification(notifCreated.UUID.String())

			_, err = client.NewCompleteJobCommand().JobKey(job.GetKey()).Send(context.Background())
			if err != nil {
				logster.Error(err, "Error completing job")
				logster.EndFuncLog()
				return
			}

			logster.EndFuncLog()
		})

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("***CAMUNDA*** Shutting down Send Notification Action worker...")
		w.Close()
	}()
}
