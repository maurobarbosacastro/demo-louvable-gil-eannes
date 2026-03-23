package models

import (
	"github.com/google/uuid"
)

// Not created on the DB because Keycloak handles it

const (
	SOURCE_SHOPIFY = "Shopify"
)

type User struct {
	Uuid                  uuid.UUID `json:"uuid"`
	Email                 string    `json:"email"`
	Country               string    `json:"country"`
	ReferralCode          string    `json:"referralCode"`
	Balance               float64   `json:"balance"`
	CreatedAt             int64     `json:"createdAt"`
	FirstName             string    `json:"firstName"`
	LastName              string    `json:"lastName"`
	Currency              string    `json:"currency"`
	DisplayName           string    `json:"displayName"`
	BirthDate             string    `json:"birthDate"`
	Groups                []string  `json:"groups"`
	IsVerified            bool      `json:"isVerified"`
	OnboardingFinished    bool      `json:"onboardingFinished"`
	ProfilePicture        string    `json:"profilePicture"`
	TransactionPercentage *float64  `json:"transactionPercentage"`
	RewardPercentage      *float64  `json:"rewardPercentage"`
	LegacyId              *string   `json:"legacyId,omitempty"`
	Newsletter            bool      `json:"newsletter"`
	Source                string    `json:"source"`
	CurrencySelected      bool      `json:"currencySelected"`
	InfluencerAmount      float64   `json:"influencerAmount,omitempty"`
}

type MembershipStatus struct {
	Level                   *string  `json:"level"`
	PercentageOnTransaction *float64 `json:"percentageOnTransaction"`
	PercentageOnReward      *float64 `json:"percentageOnReward"`
}
