package models

import "github.com/google/uuid"

type Country struct {
	Uuid         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Abbreviation *string   `gorm:"size:10"          json:"abbreviation"`
	Currency     *string   `gorm:"size:50"          json:"currency"`
	Flag         *string   `gorm:"size:100"         json:"flag"`
	Name         *string   `gorm:"size:100;not null" json:"name"`
	Enabled      *bool     `gorm:"column:enabled" json:"enabled"`
	BaseEntity
}
