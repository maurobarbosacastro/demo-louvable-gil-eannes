package models

import "github.com/google/uuid"

// UserPaymentMethod
// Information What is it?
// Is a string of an object that contains the information of the payment method
// In case of payment method is "BANK" will save
// {"iban": string, "country": country_uuid, "bankName": string, "bankAddress": string, "bankAccountTitle": string}
type UserPaymentMethod struct {
	Uuid        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	User        string    `gorm:"type:text" json:"user"`
	Information *string   `gorm:"type:jsonb" json:"information"`
	State       string    `gorm:"type:varchar(20);check:state IN ('PENDING', 'VALIDATED', 'REJECTED')" json:"state"`

	PaymentMethodUUID *uuid.UUID `gorm:"type:uuid;" json:"-"`
	FileUUID          *uuid.UUID `gorm:"type:uuid" json:"-"`

	PaymentMethod PaymentMethod `gorm:"foreignKey:PaymentMethodUUID" json:"paymentMethod"`
	File          File          `gorm:"foreignKey:FileUUID" json:"file"`

	BaseEntity
}

type AffectedRecord struct {
	Uuid             uuid.UUID `gorm:"column:uuid"`
	FileUUID         uuid.UUID `gorm:"column:file_uuid"`
	OriginalFileUUID uuid.UUID `gorm:"column:original_file_uuid"`
}
