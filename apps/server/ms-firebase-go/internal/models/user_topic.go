package models

import (
	"github.com/google/uuid"
)

type UserTopics struct {
	UserUUID  string    `gorm:"type:text;primaryKey;not null" json:"userUuid"`
	TopicUUID uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"topicUuid"`
	BaseEntity
}
