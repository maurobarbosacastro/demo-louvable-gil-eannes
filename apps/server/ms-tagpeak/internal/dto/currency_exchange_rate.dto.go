package dto

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type CreateCurrencyExchangeRateDTO struct {
	Base  string `json:"base" validate:"required"`
	Rates string `json:"rates" validate:"required"`
}

type UpdateCurrencyExchangeRateDTO struct {
	Base  *string `json:"base"`
	Rates *string `json:"rates"`
}

type FixerDTO struct {
	Success   bool            `json:"success"`
	Timestamp int64           `json:"timestamp"`
	Base      string          `json:"base"`
	Date      string          `json:"date"`
	Rates     json.RawMessage `json:"rates"`
}

type CurrencyExchangeRateDTO struct {
	Uuid      *uuid.UUID         `json:"uuid"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
	CreatedAt time.Time          `json:"createdAt"`
}
