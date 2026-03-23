package dto

import (
	"github.com/google/uuid"
)

type ReferralDTO struct {
	Uuid                       *uuid.UUID `json:"uuid"`
	ReferrerUUID               *uuid.UUID `json:"referrerUUID"`
	InviteeUUID                *uuid.UUID `json:"inviteeUUID"`
	SuccessfulFirstTransaction *bool      `json:"successfulFirstTransaction"`
}

type ReferralClickDTO struct {
	Uuid         *uuid.UUID `json:"uuid"`
	ReferrerUUID *uuid.UUID `json:"referrerUUID"`
	Code         *string    `json:"code"`
}

type CreateReferralClickDTO struct {
	ReferrerUUID *uuid.UUID `json:"referrerUUID"`
	Code         string     `json:"code"`
}

type ReferralRevenueDTO struct {
	Uuid            *uuid.UUID `json:"uuid"`
	ReferralUUID    *uuid.UUID `json:"referralUUID"`
	Amount          *float64   `json:"amount"`
	RewardUUID      *uuid.UUID `json:"rewardUUID"`
	TransactionUUID *uuid.UUID `json:"transactionUUID"`
}

type CreateReferralRevenueDTO struct {
	ReferralUUID    uuid.UUID  `json:"referralUUID"`
	Amount          float64    `json:"amount"`
	RewardUUID      *uuid.UUID `json:"rewardUUID"`
	TransactionUUID *uuid.UUID `json:"transactionUUID"`
	CreatedBy       string     `json:"createdBy"`
}
