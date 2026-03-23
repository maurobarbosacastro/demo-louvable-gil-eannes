package models

import "github.com/google/uuid"

type Language struct {
	Uuid uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name *string   `gorm:"size:255;not null" json:"name"` // Country name, not null
	Code *string   `gorm:"size:255" json:"code"`          // eCommerce platform name, nullable
	BaseEntity
}
