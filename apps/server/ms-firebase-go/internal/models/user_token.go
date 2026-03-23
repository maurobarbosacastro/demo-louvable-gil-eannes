package models

import (
	"github.com/google/uuid"
)

type UserToken struct {
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	UserUUID string    `gorm:"type:text;not null" json:"userUuid"`
	Token    string    `gorm:"type:text;not null" json:"token"`
	Topics   []Topic   `gorm:"many2many:user_token_topics;" json:"-"`

	BaseEntity

	// Unique constraint on user_uuid and token combination
	// This is defined at the table level in the migration
}
