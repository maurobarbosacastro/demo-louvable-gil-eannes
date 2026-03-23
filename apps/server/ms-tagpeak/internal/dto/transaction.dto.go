package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreateTransactionDTO struct {
	SourceId         string     `json:"sourceId,omitempty" validate:"required"`
	AmountSource     float64    `json:"amount,omitempty" validate:"required"`
	CurrencySource   string     `json:"currency,omitempty" validate:"required"`
	State            string     `json:"state,omitempty" validate:"required"` // Enum: 'TRACKED', 'VALIDATED', 'REJECTED'
	CommissionSource float64    `json:"commission,omitempty" validate:"required"`
	OrderDate        time.Time  `json:"orderDate,omitempty" validate:"required"`
	StoreVisitUUID   *uuid.UUID `json:"storeVisitUuid,omitempty" validate:"required"`
}

type CamundaCreateTransactionDTO struct {
	SourceId         string     `json:"sourceId,omitempty" validate:"required"`
	AmountSource     float64    `json:"amount,omitempty" validate:"required"`
	CurrencySource   string     `json:"currency,omitempty" validate:"required"`
	CommissionSource float64    `json:"commission,omitempty" validate:"required"`
	OrderDate        time.Time  `json:"orderDate,omitempty" validate:"required"`
	StoreVisitUUID   *uuid.UUID `json:"storeVisitUuid,omitempty" validate:"required"`
	UserUUID         uuid.UUID  `json:"userUuid,omitempty" validate:"required"`
	Reference        string     `json:"reference,omitempty"`
	State            string     `json:"state,omitempty"`
}

type UpdateTransactionDTO struct {
	CurrencySource   *string    `json:"currencySource,omitempty"`
	State            *string    `json:"state,omitempty"` // Enum: 'TRACKED', 'VALIDATED', 'REJECTED'
	CommissionTarget *float64   `json:"commissionTarget,omitempty"`
	OrderDate        *time.Time `json:"orderDate,omitempty"`
	ExitClick        *string    `json:"exitClick,omitempty"`
	AmountSource     *float64   `json:"-"`
	CommissionSource *float64   `json:"-"`
	Uuids            *[]string  `json:"uuids,omitempty"`
	Cashback         *float64   `json:"cashback,omitempty"`
}

type TransactionFiltersDTO struct {
	State          *string    `json:"state,omitempty" query:"state"`
	User           *string    `json:"userUuid,omitempty" query:"userUuid"`
	StoreUuid      *string    `json:"storeUuid,omitempty" query:"storeUuid"`
	StoreVisitUuid *string    `json:"storeVisitUuid,omitempty" query:"storeVisitUuid"`
	StartDate      *time.Time `json:"StartDate,omitempty" query:"startDate"`
	EndDate        *time.Time `json:"endDate,omitempty" query:"endDate"`
}
