package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

// Custom join table for notification and user_token with state
type NotificationUserTokens struct {
	NotificationUUID uuid.UUID                  `gorm:"type:uuid;primaryKey" json:"notificationUuid"`
	UserTokenUUID    uuid.UUID                  `gorm:"type:uuid;primaryKey" json:"userTokenUuid"`
	State            JoinTableNotificationState `gorm:"type:join_table_notification_state;not null;default:pending" json:"state"`
	CreatedAt        time.Time                  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt        *time.Time                 `gorm:"autoUpdateTime" json:"updatedAt"`
}

// Custom join table for notification and topic with state
type NotificationTopics struct {
	NotificationUUID uuid.UUID                  `gorm:"type:uuid;primaryKey" json:"notificationUuid"`
	TopicUUID        uuid.UUID                  `gorm:"type:uuid;primaryKey" json:"topicUuid"`
	State            JoinTableNotificationState `gorm:"type:join_table_notification_state;not null;default:pending" json:"state"`
	CreatedAt        time.Time                  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt        *time.Time                 `gorm:"autoUpdateTime" json:"updatedAt"`
}

type Notification struct {
	UUID   uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Title  string            `gorm:"type:text;not null" json:"title"`
	Body   string            `gorm:"type:text;not null" json:"body"`
	Target []UserToken       `gorm:"many2many:notification_user_tokens;joinForeignKey:NotificationUUID;joinReferences:UserTokenUUID" json:"target"`
	Topics []Topic           `gorm:"many2many:notification_topics;joinForeignKey:NotificationUUID;joinReferences:TopicUUID" json:"topics"`
	Date   time.Time         `gorm:"not null" json:"date"`
	State  NotificationState `gorm:"type:notification_state;not null;default:draft" json:"state"`
	BaseEntity
}

// SetupJoinTable configures the custom join table for notification user tokens
func (Notification) SetupJoinTable(db *gorm.DB) {
	err := db.SetupJoinTable(&Notification{}, "Target", &NotificationUserTokens{})
	if err != nil {
		panic("Failed to setup NotificationUserToken join table: " + err.Error())
	}

	err = db.SetupJoinTable(&Notification{}, "Topics", &NotificationTopics{})
	if err != nil {
		panic("Failed to setup NotificationTopic join table: " + err.Error())
	}
}
