package models

import (
	"github.com/google/uuid"
)

type CurrencyExchangeRate struct {
	Uuid  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Base  *string   `gorm:"type:text" json:"base"`
	Rates string    `gorm:"type:jsonb" json:"rates"`
	BaseEntity
}
