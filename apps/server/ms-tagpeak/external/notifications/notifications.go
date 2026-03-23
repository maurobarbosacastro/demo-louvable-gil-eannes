package notifications

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type NotificationState string

const (
	NotificationStateDraft     NotificationState = "draft"
	NotificationStateScheduled NotificationState = "scheduled"
	NotificationStateProcessed NotificationState = "processed"
	NotificationStateError     NotificationState = "error"
	NotificationStateCancelled NotificationState = "cancelled"
)

type JoinTableNotificationState string

const (
	JoinTableNotificationStatePending   JoinTableNotificationState = "pending"
	JoinTableNotificationStateDelivered JoinTableNotificationState = "delivered"
	JoinTableNotificationStateCanceled  JoinTableNotificationState = "canceled"
	JoinTableNotificationStateError     JoinTableNotificationState = "error"
)

type UserToken struct {
	UUID     uuid.UUID `json:"uuid"`
	UserUUID string    `json:"userUuid"`
	Token    string    `json:"token"`
	BaseEntity
}

type Topic struct {
	UUID   uuid.UUID   `json:"uuid"`
	Name   string      `json:"name"`
	Target []UserToken `json:"target,omitempty"`
	BaseEntity
}

type Notification struct {
	UUID   uuid.UUID         `json:"uuid"`
	Title  string            `json:"title"`
	Body   string            `json:"body"`
	Target []UserToken       `json:"target,omitempty"`
	Topics []Topic           `json:"topics,omitempty"`
	Date   time.Time         `json:"date"`
	State  NotificationState `json:"state"`
	BaseEntity
}

type NotificationResponse struct {
	Success   bool      `json:"success"`
	MessageID string    `json:"message_id,omitempty"`
	Error     string    `json:"error,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	RequestID string    `json:"request_id,omitempty"`
}

type TokenSave struct {
	UserUUID string `json:"userUuid"`
	Token    string `json:"token"`
}

type CreateNotificationDto struct {
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	UserTarget  *[]string `json:"targets,omitempty"`
	TopicTarget *[]string `json:"topics,omitempty"`
	Date        time.Time `json:"date"`
	State       *string   `json:"state,omitempty"`
}

type AddRemoveTokenToTopicDto struct {
	Token string `json:"token"`
}

type AddRemoveUserToTopicDto struct {
	UserUuid string `json:"userUuid"`
}
