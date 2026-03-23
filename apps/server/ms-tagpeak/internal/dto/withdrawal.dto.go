package dto

import (
	"github.com/google/uuid"
	"time"
)

// WithdrawalDTO is a data transfer object for the Withdrawal struct
type WithdrawalDTO struct {
	Uuid           *uuid.UUID        `json:"uuid"`
	UserID         *string           `json:"userId"`
	MethodID       *PaymentMethodDTO `json:"method"`
	Amount         *float64          `json:"amount"`
	Details        *string           `json:"details,omitempty"` // Optional field
	State          *string           `json:"state"`
	CompletionDate *time.Time        `json:"completionDate,omitempty"` // Optional, can be null
	Rate           *float64          `json:"rate"`
	CurrencySource *string           `json:"currencySource"`
	CurrencyTarget *string           `json:"currencyTarget"`
}

type CreateWithdrawalDTO struct {
	MethodID uuid.UUID `json:"methodId" validate:"required"`
	Details  *string   `json:"details,omitempty"`
}

type UpdateWithdrawalDTO struct {
	UserID         *string    `json:"userId"`
	UserMethodID   *uuid.UUID `json:"userMethodId"`
	Amount         *float64   `json:"amount"`
	Details        *string    `json:"details,omitempty"` // Optional field
	State          *string    `json:"state"`
	CompletionDate *time.Time `json:"completionDate,omitempty"` // Optional, can be null
	Rate           *float64   `json:"rate"`
	CurrencySource *string    `json:"currencySource"`
	CurrencyTarget *string    `json:"currencyTarget"`
}

type WithdrawalFiltersDTO struct {
	StartDate *string `json:"startDate,omitempty" query:"startDate"`
	EndDate   *string `json:"endDate,omitempty" query:"endDate"`
	User      *string `json:"userUuid,omitempty" query:"userUuid"`
	State     string  `json:"state,omitempty" query:"state"`
}

type BulkUpdateStateRequestDTO struct {
	Uuids []string `json:"uuids"`
	UpdateWithdrawalDTO
}
