package dto

import "github.com/google/uuid"

type CreateUserPaymentMethodDTO struct {
	PaymentMethod    uuid.UUID  `json:"paymentMethod,omitempty" validate:"required"`
	BankName         string     `json:"bankName,omitempty"`
	BankAddress      string     `json:"bankAddress,omitempty"`
	BankCountry      string     `json:"bankCountry,omitempty"`
	Country          string     `json:"country,omitempty"`
	BankAccountTitle string     `json:"bankAccountTitle,omitempty"`
	Iban             string     `json:"iban,omitempty"`
	IbanStatement    *uuid.UUID `json:"ibanStatement,omitempty"`
	Vat              string     `json:"vat,omitempty"`
}

type Information struct {
	BankName         string `json:"bankName"`
	BankAddress      string `json:"bankAddress"`
	Country          string `json:"country"`
	BankAccountTitle string `json:"bankAccountTitle"`
	Iban             string `json:"iban"`
	BankCountry      string `json:"bankCountry"`
	Vat              string `json:"vat"`
}
