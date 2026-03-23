package models

import (
	"github.com/google/uuid"
)

// PaymentMethod represents the payment_method table
type PaymentMethod struct {
	Uuid uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name *string   `gorm:"type:text;not null" json:"name"`
	Code *string   `gorm:"type:text;unique;not null" json:"code"`
	BaseEntity
}
