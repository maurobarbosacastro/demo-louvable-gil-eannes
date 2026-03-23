package models

import "github.com/google/uuid"

type StoreCountry struct {
	CountryCode string    `gorm:"type:text;primaryKey"`
	StoreUUID   uuid.UUID `gorm:"type:uuid;primaryKey"`

	Country Country `gorm:"foreignKey:CountryCode; references:abbreviation" json:"country"`
	Store   Store   `gorm:"foreignKey:StoreUUID; references:uuid" json:"store"`
}
