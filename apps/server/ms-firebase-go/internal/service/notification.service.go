package service

import (
	"context"
	"fmt"
	"ms-firebase-go/external/firebase"
	"ms-firebase-go/internal/dto"
	"ms-firebase-go/internal/models"
	"ms-firebase-go/internal/repository"
	"ms-firebase-go/pkg/logster"
	"ms-firebase-go/pkg/utils"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

func SaveNotifications(notifDto dto.CreateNotificationDto) (*models.Notification, error) {
	logster.StartFuncLog()

	var userTargets []models.UserToken
	var topicTargets []models.Topic
	var err error

	if notifDto.UserTarget != nil {
		userTargets, err = repository.GetUserTokensByUserUUIDs(*notifDto.UserTarget)
		if err != nil {
			logster.Error(err, "Error while getting targets token")
			logster.EndFuncLog()
			return nil, err
		}
	}

	if notifDto.TopicTarget != nil {
		err = nil
		topicTargets, err = repository.GetTopicsByUUIDs(*notifDto.TopicTarget)
		if err != nil {
			logster.Error(err, "Error while getting topics")
			logster.EndFuncLog()
			return nil, err
		}
	}

	state := models.NotificationStateDraft
	if notifDto.State != nil {
		state = *notifDto.State
	}

	model := models.Notification{
		Title: notifDto.Title,
		Body:  notifDto.Content,
		Date:  notifDto.Date,
		State: state,
		BaseEntity: models.BaseEntity{
			CreatedAt: time.Now(),
			CreatedBy: "",
		},
	}

	if len(userTargets) != 0 {
		model.Target = userTargets
	}

	if len(topicTargets) != 0 {
		model.Topics = topicTargets
	}

	modelCreated, err := repository.CreateNotification(model)

	logster.EndFuncLogMsg(fmt.Sprintf("model created: %+v", modelCreated))
	return modelCreated, err
}

func UpdateNotification(notifUuid uuid.UUID, notifDto dto.NotificationUpdateFields) (*models.Notification, error) {
	logster.StartFuncLogMsg(notifUuid)

	notification, errUpdateNotifciation := repository.UpdateNotification(notifUuid, notifDto)
	if errUpdateNotifciation != nil {
		logster.Error(errUpdateNotifciation, "Error updating notification")
		logster.EndFuncLog()
		return nil, errUpdateNotifciation
	}

	var userTargets *[]models.UserToken
	var topicTargets *[]models.Topic

	if notifDto.UserTarget != nil {
		t, err := repository.GetUserTokensByUserUUIDs(*notifDto.UserTarget)
		if err != nil {
			logster.Error(err, "Error while getting targets token")
			logster.EndFuncLog()
			return nil, err
		}
		userTargets = &t
	}

	if notifDto.TopicTarget != nil {
		t, err := repository.GetTopicsByUUIDs(*notifDto.TopicTarget)
		if err != nil {
			logster.Error(err, "Error while getting topics")
			logster.EndFuncLog()
			return nil, err
		}
		topicTargets = &t
	}

	errReplaceTargets := repository.ReplaceNotificationAssociations(notification.UUID, userTargets, topicTargets)
	if errReplaceTargets != nil {
		logster.Error(errReplaceTargets, "Error replacing targets")
		logster.EndFuncLog()
		return nil, errReplaceTargets
	}

	notification, errGetNotifciation := repository.GetNotificationByID(notification.UUID)

	if errGetNotifciation != nil {
		logster.Error(errGetNotifciation, "Error getting notification")
		logster.EndFuncLog()
		return nil, errGetNotifciation
	}

	logster.EndFuncLog()
	return notification, nil
}

func GetNotificationByUuid(notifUuid uuid.UUID) (*models.Notification, error) {
	logster.StartFuncLog()

	notification, err := repository.GetNotificationByID(notifUuid)
	if err != nil {
		logster.Error(err, "Error getting notification")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLogMsg(notification.UUID)
	return notification, nil
}

func UpdateNotificationUserTokenState(notificationID, userTokenID uuid.UUID, state models.JoinTableNotificationState) error {
	logster.StartFuncLogMsg(notificationID)

	err := repository.UpdateNotificationUserTokenState(notificationID, userTokenID, state)
	if err != nil {
		logster.Error(err, "Error updating notification user token state")
		logster.EndFuncLog()
		return err
	}

	logster.EndFuncLog()
	return nil
}

func UpdateNotificationTopicState(notificationID, topicID uuid.UUID, state models.JoinTableNotificationState) error {
	logster.StartFuncLogMsg(notificationID)

	err := repository.UpdateNotificationTopicState(notificationID, topicID, state)
	if err != nil {
		logster.Error(err, "Error updating notification topic state")
		logster.EndFuncLog()
		return err
	}

	logster.EndFuncLog()
	return nil
}

func GetNotificationWithRelationStates(notificationID uuid.UUID) (*dto.NotificationWithRelationStatesDto, error) {
	logster.StartFuncLogMsg(notificationID)

	// Get the notification
	notification, err := repository.GetNotificationByID(notificationID)
	if err != nil {
		logster.Error(err, "Error getting notification")
		logster.EndFuncLog()
		return nil, err
	}

	// Get user token relations with states
	userTokenRelations, err := repository.GetNotificationUserTokenStates(notificationID)
	if err != nil {
		logster.Error(err, "Error getting notification user token states")
		logster.EndFuncLog()
		return nil, err
	}

	// Get topic relations with states
	topicRelations, err := repository.GetNotificationTopicStates(notificationID)
	if err != nil {
		logster.Error(err, "Error getting notification topic states")
		logster.EndFuncLog()
		return nil, err
	}

	result := &dto.NotificationWithRelationStatesDto{
		Notification:       *notification,
		UserTokenRelations: userTokenRelations,
		TopicRelations:     topicRelations,
	}

	logster.EndFuncLog()
	return result, nil
}

func ProcessCurrentNotifications() {
	now := time.Now().UTC().Truncate(time.Second)
	logster.StartFuncLogMsg(now)

	// Get notifications
	notifications, errGetNotifications := repository.GetNotificationByDate(now)

	if errGetNotifications != nil {
		logster.Error(errGetNotifications, "Error gettings notifications")
		return
	}

	if len(*notifications) == 0 {
		logster.Info("No notificiations to send")
		return
	}
	logster.Infof("Notifications to process: %d", len(*notifications))

	notifUuuids := lo.Map(*notifications, func(item models.Notification, index int) string {
		return item.UUID.String()
	})

	logster.Infof("Notifications UUIDs to update: %+v", notifUuuids)
	logster.Debugf("%+v", *notifications)

	lo.ForEach(*notifications, func(notification models.Notification, index int) {
		logster.Infof("Handling notification %s | Title: %s | Body: %s", notification.UUID, notification.Title, notification.Body)

		// TODO: This has a limitation since firebase, for singular tokens, processed by batches of 500 so this would need to be upgraded in the future

		if len(notification.Target) > 0 {
			logster.Infof("Targets # : %d", len(notification.Target))

			lo.ForEach(notification.Target, func(target models.UserToken, index int) {
				logster.Infof("Sending notitification for target %s | user: %s", target.Token, target.UserUUID)

				SendNotificationToTarget(notification, target)
			})

		}

		if len(notification.Topics) > 0 {
			logster.Infof("Topics # : %d", len(notification.Topics))

			lo.ForEach(notification.Topics, func(topic models.Topic, index int) {
				logster.Infof("Sending notification for topic %s | Name: %s", topic.UUID, topic.Name)
				SendNotificationToTopic(notification, topic)
			})

		}

		if len(notification.Target) == 0 && len(notification.Topics) == 0 {
			logster.Info("No targets/topics set for this notification")
		}

		_, errUpdate := UpdateNotification(
			notification.UUID,
			dto.NotificationUpdateFields{
				State: utils.Ptr(models.NotificationStateProcessed),
			})

		if errUpdate != nil {
			logster.Error(errUpdate, "Error updating notification status")
		}
	})
}

func SendNotificationToTarget(notification models.Notification, target models.UserToken) {
	_, errSend := firebase.SendNotification(
		context.Background(),
		target.Token,
		&dto.MessageDTO{
			Title: notification.Title,
			Body:  notification.Body,
		})
	newState := models.JoinTableNotificationStateDelivered

	if errSend != nil {
		logster.Error(errSend, "Error sending notification to target")
		newState = models.JoinTableNotificationStateError
	}

	errUpdate := repository.UpdateNotificationUserTokenState(
		notification.UUID,
		target.UUID,
		newState,
	)

	if errUpdate != nil {
		logster.Error(errUpdate, "Error seting notification to delivered")
	}
}

func SendNotificationToTopic(notification models.Notification, topic models.Topic) {
	_, errSendResponse := firebase.SendGroupMessage(
		context.Background(),
		topic.UUID.String(),
		&dto.MessageDTO{
			Title: notification.Title,
			Body:  notification.Body,
		})

	newState := models.JoinTableNotificationStateDelivered

	if errSendResponse != nil {
		logster.Error(errSendResponse, "Error sending notification")
		newState = models.JoinTableNotificationStateError
	}

	errUpdate := repository.UpdateNotificationTopicState(
		notification.UUID,
		topic.UUID,
		newState,
	)

	if errUpdate != nil {
		logster.Error(errUpdate, "Error while updating notification <-> user_tokens")
	}

	if len(notification.Target) == 0 && len(notification.Topics) == 0 {
		logster.Info("No targets/topics set for this notification")
	}

	_, errUpdate = UpdateNotification(
		notification.UUID,
		dto.NotificationUpdateFields{
			State: utils.Ptr(models.NotificationStateProcessed),
		})

	if errUpdate != nil {
		logster.Error(errUpdate, "Error updating notification status")
	}
}
