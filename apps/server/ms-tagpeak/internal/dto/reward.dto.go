package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateRewardDTO struct {
	TransactionUUID   uuid.UUID `json:"transactionUuid,omitempty" validate:"required"`
	Isin              string    `json:"isin,omitempty" validate:"required"`
	Conid             string    `json:"conid,omitempty" validate:"required"`
	CurrencySource    string    `json:"currency,omitempty" validate:"required"`
	State             string    `json:"state,omitempty"`                            // Enum: 'LIVE', 'STOPPED', 'FINISHED', 'REQUESTED', 'PAID'
	InitialPrice      float64   `json:"initialPrice,omitempty" validate:"required"` // One-to-one relationship (optional)
	EndDate           time.Time `json:"endDate,omitempty" validate:"required"`      // One-to-one relationship (optional)
	Type              string    `json:"type,omitempty" validate:"required"`         // Enum: 'INVESTMENT', 'FIXED'
	Title             *string   `json:"title,omitempty"`
	Details           *string   `json:"details,omitempty"`
	CurrentRewardUser *float64  `json:"-"`
	Origin            string    `json:"-"` // Enum: 'PURCHASE', 'REFERRAL' -  DEFAULT 'PURCHASE'
}

type UpdateRewardDTO struct {
	Isin                *string    `json:"isin,omitempty"`
	Conid               *string    `json:"conid,omitempty"`
	InitialReward       *float64   `json:"initialReward,omitempty"`
	CurrentRewardSource *float64   `json:"currentRewardSource,omitempty"`
	State               *string    `json:"state,omitempty"` // Enum: 'LIVE', 'STOPPED', 'FINISHED', 'REQUESTED', 'PAID'
	InitialPrice        *float64   `json:"initialPrice,omitempty"`
	EndDate             *time.Time `json:"endDate,omitempty"`
	Type                *string    `json:"type,omitempty"` // Enum: 'INVESTMENT', 'FIXED'
	InitialDate         *time.Time `json:"initialDate,omitempty"`
	Title               *string    `json:"title,omitempty"`
	Details             *string    `json:"details,omitempty"`
	OverridePrice       *float64   `json:"overridePrice,omitempty"`
	Origin              *string    `json:"origin,omitempty"` // Enum: 'PURCHASE', 'REFERRAL' -  DEFAULT 'PURCHASE'

}

type RewardBulkEditReq struct {
	Uuids         []string   `json:"uuids"`
	InitialPrice  *float64   `json:"initialPrice"`
	EndDate       *time.Time `json:"endDate"`
	InitialDate   *time.Time `json:"initialDate"`
	Status        *string    `json:"status"`
	Isin          *string    `json:"isin"`
	Conid         *string    `json:"conid"`
	OverridePrice *float64   `json:"overridePrice"` //Source price
}

type CreateRewardBulkDTO struct {
	TransactionUUIDs []uuid.UUID `json:"transactionUuids,omitempty" `
	Isin             string      `json:"isin,omitempty" `
	Conid            string      `json:"conid,omitempty" `
	CurrencySource   string      `json:"currency,omitempty" `
	State            string      `json:"status,omitempty"`
	InitialDate      *time.Time  `json:"initialDate"`             // Enum: 'LIVE', 'STOPPED', 'FINISHED', 'REQUESTED', 'PAID'
	InitialPrice     float64     `json:"initialPrice,omitempty" ` // One-to-one relationship (optional)
	EndDate          time.Time   `json:"endDate,omitempty" `      // One-to-one relationship (optional)
	Type             string      `json:"type,omitempty" `         // Enum: 'INVESTMENT', 'FIXED'
	Title            *string     `json:"title,omitempty"`
	Details          *string     `json:"details,omitempty"`
	UserUUID         string      `json:"userUuid,omitempty"`
	OverridePrice    *float64    `json:"overridePrice,omitempty"` //Source price
	Origin           string      `json:"origin,omitempty"`        // Enum: 'PURCHASE', 'REFERRAL'
}
