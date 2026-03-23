package response_object

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ms-tagpeak/internal/models"
	"time"
)

type RewardCashbackRO struct {
	Uuid                *uuid.UUID `json:"uuid,omitempty"`
	Isin                *string    `json:"isin,omitempty"`
	Conid               *string    `json:"conid,omitempty"`
	CurrentRewardSource *float64   `json:"currentRewardSource,omitempty"`
	CurrentRewardTarget *float64   `json:"currentRewardTarget,omitempty"`
	CurrentRewardUser   *float64   `json:"currentRewardUser,omitempty"`
	Status              *string    `json:"state,omitempty"` // Enum: 'LIVE', 'STOPPED', 'FINISHED', 'REQUESTED', 'PAID'
	PriceDayZero        *float64   `json:"initialPrice,omitempty"`
	Title               *string    `json:"title,omitempty"`
	StartDate           *time.Time `json:"startDate,omitempty"`
	EndDate             *time.Time `json:"endDate,omitempty"`
	Origin              string     `json:"origin,omitempty"` // Enum: 'REFERRAL', 'PURCHASE'
	StoppedAt           *time.Time `json:"stoppedAt,omitempty"`
}

type UsersWithTotalAmountReward struct {
	User  uuid.UUID `json:"user"`
	Total float64   `json:"total"`
}

type UsersWithTotalAmountRewardFinal struct {
	Updated []UsersWithTotalAmountReward `json:"updated"`
	Error   []UsersWithTotalAmountReward `json:"error"`
}

type RewardDetailRO struct {
	Uuid                     uuid.UUID              `json:"uuid"`
	User                     string                 `json:"user"`
	TransactionUUID          uuid.UUID              `json:"transactionUuid"`
	Isin                     string                 `json:"isin"`
	Conid                    string                 `json:"conid"`
	InitialReward            float64                `json:"initialReward"`
	CurrentRewardSource      float64                `json:"currentRewardSource"`
	CurrentRewardTarget      float64                `json:"currentRewardTarget"`
	CurrentRewardUser        float64                `json:"currentRewardUser"`
	CurrencyExchangeRateUUID uuid.UUID              `json:"currencyExchangeRateUuid"`
	CurrencySource           string                 `json:"currencySource"`
	CurrencyTarget           string                 `json:"currencyTarget"`
	CurrencyUser             string                 `json:"currencyUser"`
	State                    string                 `json:"state"`
	InitialPrice             float64                `json:"initialPrice"`
	EndDate                  time.Time              `json:"endDate"`
	AssetUnits               float64                `json:"assetUnits"`
	History                  []models.RewardHistory `json:"history"`
	Type                     string                 `json:"type"`
	Title                    string                 `json:"title"`
	Details                  string                 `json:"details"`
	OverridePrice            *float64               `json:"overridePrice"`
	WithdrawalUuid           *uuid.UUID             `json:"withdrawalUuid"`
	Origin                   string                 `json:"origin"`
	StoppedAt                *time.Time             `json:"stoppedAt"`
	CreatedAt                time.Time              `json:"createdAt"`
	CreatedBy                string                 `json:"createdBy"`
	UpdatedAt                *time.Time             `json:"updatedAt"`
	UpdatedBy                *string                `json:"updatedBy"`
	Deleted                  bool                   `json:"deleted"`
	DeletedAt                *gorm.DeletedAt        `json:"deletedAt" swaggertype:"string" swaggertype:"string" format:"date-time"`
	DeletedBy                *string                `json:"deletedBy"`
	MinimumReward            float64                `json:"minimumReward"`
}

func (r RewardDetailRO) FromReward(reward models.Reward) RewardDetailRO {
	r.Uuid = reward.Uuid
	r.User = reward.User
	r.TransactionUUID = reward.TransactionUUID
	r.Isin = reward.Isin
	r.Conid = reward.Conid
	r.InitialReward = reward.InitialReward
	r.CurrentRewardSource = reward.CurrentRewardSource
	r.CurrentRewardTarget = reward.CurrentRewardTarget
	r.CurrentRewardUser = reward.CurrentRewardUser
	r.CurrencyExchangeRateUUID = reward.CurrencyExchangeRateUUID
	r.CurrencySource = reward.CurrencySource
	r.CurrencyTarget = reward.CurrencyTarget
	r.CurrencyUser = reward.CurrencyUser
	r.State = reward.State
	r.InitialPrice = reward.InitialPrice
	r.EndDate = reward.EndDate
	r.AssetUnits = reward.AssetUnits
	r.History = reward.History
	r.Type = reward.Type
	r.Title = reward.Title
	r.Details = reward.Details
	r.OverridePrice = reward.OverridePrice
	r.WithdrawalUuid = reward.WithdrawalUuid
	r.Origin = reward.Origin
	r.StoppedAt = reward.StoppedAt
	r.CreatedAt = reward.CreatedAt
	r.CreatedBy = reward.CreatedBy
	r.UpdatedAt = reward.UpdatedAt
	r.UpdatedBy = reward.UpdatedBy
	r.Deleted = reward.Deleted
	r.DeletedAt = reward.DeletedAt
	r.DeletedBy = reward.DeletedBy
	r.MinimumReward = 0

	return r
}
