package response_object

import (
	"github.com/google/uuid"
)

type SimpleUserDto struct {
	Uuid  uuid.UUID `json:"uuid"`
	Email string    `json:"email,omitempty"`
	Name  string    `json:"name,omitempty"`
}

type UserDto struct {
	Uuid                  uuid.UUID   `json:"uuid"`
	Email                 string      `json:"email"`
	Country               string      `json:"country"`
	ReferralCode          string      `json:"referralCode"`
	Balance               float64     `json:"balance"`
	CreatedAt             int64       `json:"createdAt"`
	FirstName             string      `json:"firstName"`
	LastName              string      `json:"lastName"`
	Currency              string      `json:"currency"`
	DisplayName           string      `json:"displayName"`
	BirthDate             string      `json:"birthDate"`
	Groups                []string    `json:"groups"`
	IsVerified            bool        `json:"isVerified"`
	OnboardingFinished    bool        `json:"onboardingFinished"`
	Referral              *InviteeDto `json:"referral,omitempty"`
	ProfilePicture        *string     `json:"profilePicture"`
	TransactionPercentage *float64    `json:"transactionPercentage"`
	RewardPercentage      *float64    `json:"rewardPercentage"`
	LegacyId              *string     `json:"legacyId,omitempty"`
	Newsletter            bool        `json:"newsletter,omitempty"`
}

type UserReferralRevenueInfoDto struct {
	Uuid                       uuid.UUID `json:"uuid"`
	FirstName                  string    `json:"firstName"`
	LastName                   string    `json:"lastName"`
	ProfilePicture             string    `json:"profilePicture"`
	ReferredValue              float64   `json:"referredValue"`
	FirstTransactionSuccessful bool      `json:"firstTransactionSuccessful"`
	DisplayName                string    `json:"displayName,omitempty"`
}
