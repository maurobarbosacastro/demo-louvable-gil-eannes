package models

import (
	"github.com/google/uuid"
)

type Topic struct {
	UUID   uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name   string      `gorm:"type:text;not null;uniqueIndex" json:"name"`
	Target []UserToken `gorm:"many2many:user_token_topics;" json:"target"`
	BaseEntity
}

type TopicWithoutTarget struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
	BaseEntity
}

func (ut *Topic) ToTopicWithoutTarget() TopicWithoutTarget {
	return TopicWithoutTarget{
		UUID: ut.UUID,
		Name: ut.Name,
		BaseEntity: BaseEntity{
			CreatedAt: ut.CreatedAt,
			CreatedBy: ut.CreatedBy,
			UpdatedAt: ut.UpdatedAt,
			UpdatedBy: ut.UpdatedBy,
		},
	}
}
