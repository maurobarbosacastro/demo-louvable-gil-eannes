package models

import "github.com/google/uuid"

type StoreCategory struct {
	CategoryCode string    `gorm:"type:text;primaryKey"`
	StoreUUID    uuid.UUID `gorm:"type:uuid;primaryKey"`

	Category Category `gorm:"foreignKey:CategoryCode; references:code" json:"category"`
	Store    Store    `gorm:"foreignKey:StoreUUID; references:uuid" json:"store"`
}
