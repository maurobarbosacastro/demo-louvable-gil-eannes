package models

import (
	"github.com/google/uuid"
	"time"
)

// Withdrawal represents the withdrawal table
type Withdrawal struct {
	Uuid                     uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();primaryKey" json:"uuid"`
	User                     string     `gorm:"type:text;not null" json:"user"`
	UserMethod               uuid.UUID  `gorm:"type:uuid;not null" json:"-"`
	AmountSource             float64    `gorm:"type:double precision;not null" json:"amountSource" `
	AmountTarget             float64    `gorm:"type:double precision;not null" json:"amountTarget" `
	Details                  *string    `gorm:"type:text" json:"details"`
	State                    string     `gorm:"type:varchar(20);check:state IN ('PENDING', 'COMPLETED', 'REJECTED')" json:"state"`
	CompletionDate           *time.Time `gorm:"type:timestamp" json:"completionDate" `
	CurrencyExchangeRateUUID uuid.UUID  `gorm:"type:uuid" json:"currencyExchangeRateUuid"`
	CurrencySource           string     `gorm:"type:text" json:"currencySource"`
	CurrencyTarget           string     `gorm:"type:text" json:"currencyTarget"`
	BaseEntity

	// Associations
	UserPaymentMethod    UserPaymentMethod    `gorm:"foreignKey:UserMethod"`
	CurrencyExchangeRate CurrencyExchangeRate `gorm:"foreignKey:CurrencyExchangeRateUUID" json:"-"`
}
