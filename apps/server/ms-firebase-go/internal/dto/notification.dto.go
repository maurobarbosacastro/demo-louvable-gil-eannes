package dto

import (
	"ms-firebase-go/internal/models"
	"time"
)

type CreateNotificationDto struct {
	Title       string                    `json:"title"`
	Content     string                    `json:"content"`
	UserTarget  *[]string                 `json:"targets,omitempty"`
	TopicTarget *[]string                 `json:"topics,omitempty"`
	Date        time.Time                 `json:"date"`
	State       *models.NotificationState `json:"state,omitempty"`
}

type NotificationUpdateFields struct {
	Title       *string                   `json:"title,omitempty"`
	Content     *string                   `json:"content,omitempty"`
	UserTarget  *[]string                 `json:"target,omitempty"`
	TopicTarget *[]string                 `json:"topics,omitempty"`
	Date        *time.Time                `json:"date,omitempty"`
	State       *models.NotificationState `json:"state,omitempty"`
}

type NotificationSendUser struct {
	UserUuid string `json:"userUuid"`
}

type NotificationSendTopic struct {
	Topic string `json:"topic"`
}

type UpdateNotificationUserTokenStateDto struct {
	State models.JoinTableNotificationState `json:"state"`
}

type UpdateNotificationTopicStateDto struct {
	State models.JoinTableNotificationState `json:"state"`
}

type NotificationWithRelationStatesDto struct {
	models.Notification
	UserTokenRelations []models.NotificationUserTokens `json:"userTokenRelations,omitempty"`
	TopicRelations     []models.NotificationTopics     `json:"topicRelations,omitempty"`
}
