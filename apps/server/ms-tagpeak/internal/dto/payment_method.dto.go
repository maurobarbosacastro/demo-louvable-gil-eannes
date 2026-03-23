package dto

import "github.com/google/uuid"

// PaymentMethodDTO is a simplified data transfer object for the PaymentMethod struct
type PaymentMethodDTO struct {
	Uuid *uuid.UUID `json:"uuid"`
	Name *string    `json:"name"`
	Code *string    `json:"code"`
}

type CreatePaymentMethodDTO struct {
	Name *string `json:"name" validate:"required"`
	Code *string `json:"code" validate:"required"`
}

type UpdatePaymentMethodDTO struct {
	Name *string `json:"name"`
	Code *string `json:"code"`
}
