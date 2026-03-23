package models

import "github.com/google/uuid"

type Category struct {
	Uuid uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name *string   `gorm:"size:100;not null" json:"name"` // Category name, not null
	Code *string   `gorm:"text; unique;" json:"code"`
	BaseEntity
}
