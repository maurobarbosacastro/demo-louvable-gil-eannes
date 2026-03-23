package models

import "github.com/google/uuid"

type StoreVisit struct {
	Uuid      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	User      *string   `gorm:"size:255;not null" json:"user"`
	Reference *string   `gorm:"size:255" json:"reference"`
	Purchase  bool      `gorm:"type:bool" json:"purchase"`

	StoreUUID *uuid.UUID `gorm:"type:uuid;" json:"-"`

	Store Store `gorm:"foreignKey:StoreUUID" json:"store"`

	BaseEntity
}
