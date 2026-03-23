package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	RewardStateLive        = "LIVE"
	RewardStateStopped     = "STOPPED"
	RewardStateFinished    = "FINISHED"
	RewardStateRequested   = "REQUESTED"
	RewardStatePaid        = "PAID"
	RewardTypeInvestment   = "INVESTMENT"
	RewardTypeFixed        = "FIXED"
	RewardOriginPurchase   = "PURCHASE"
	RewardOriginReferral   = "REFERRAL"
	RewardOriginCommission = "COMMISSION"
)

type Reward struct {
	Uuid                     uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	User                     string          `gorm:"type:text" json:"user"`
	TransactionUUID          uuid.UUID       `gorm:"type:uuid" json:"transactionUuid"`
	Isin                     string          `gorm:"type:text" json:"isin"`
	Conid                    string          `gorm:"type:text" json:"conid"`
	InitialReward            float64         `gorm:"type:double precision" json:"initialReward"`
	CurrentRewardSource      float64         `gorm:"type:double precision" json:"currentRewardSource"`
	CurrentRewardTarget      float64         `gorm:"type:double precision" json:"currentRewardTarget"`
	CurrentRewardUser        float64         `gorm:"type:double precision" json:"currentRewardUser"`
	CurrencyExchangeRateUUID uuid.UUID       `gorm:"type:uuid" json:"currencyExchangeRateUuid"`
	CurrencySource           string          `gorm:"type:text" json:"currencySource"`
	CurrencyTarget           string          `gorm:"type:text" json:"currencyTarget"`
	CurrencyUser             string          `gorm:"type:text" json:"currencyUser"`
	State                    string          `gorm:"type:varchar(20);check:state IN ('LIVE', 'STOPPED', 'EXPIRED', FINISHED', 'REQUESTED', 'PAID')" json:"state"`
	InitialPrice             float64         `gorm:"type:double precision" json:"initialPrice"`
	EndDate                  time.Time       `gorm:"type:date" json:"endDate"`
	AssetUnits               float64         `gorm:"type:double precision" json:"assetUnits"`
	History                  []RewardHistory `gorm:"foreignKey:RewardUUID" json:"history"` // One-to-many relationship
	Type                     string          `gorm:"type:varchar(20);check:state IN ('INVESTMENT', 'FIXED')" json:"type"`
	Title                    string          `gorm:"type:text" json:"title"`
	Details                  string          `gorm:"type:text" json:"details"`
	OverridePrice            *float64        `gorm:"type:double precision" json:"overridePrice"`
	WithdrawalUuid           *uuid.UUID      `gorm:"type:uuid" json:"withdrawalUuid"`
	Origin                   string          `gorm:"origin:text" json:"origin"`
	StoppedAt                *time.Time      `gorm:"type:date" json:"stoppedAt"`
	BaseEntity

	Transaction          Transaction          `gorm:"foreignKey:TransactionUUID" json:"-"`
	CurrencyExchangeRate CurrencyExchangeRate `gorm:"foreignKey:CurrencyExchangeRateUUID" json:"-"`
	Withdrawal           Withdrawal           `gorm:"foreignKey:WithdrawalUuid" json:"-"`
}

type CreateReward struct {
	Uuid                     uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	User                     string          `gorm:"type:text" json:"user"`
	TransactionUUID          uuid.UUID       `gorm:"type:uuid" json:"transactionUUID"`
	Isin                     string          `gorm:"type:text" json:"isin"`
	Conid                    string          `gorm:"type:text" json:"conid"`
	InitialReward            float64         `gorm:"type:double precision" json:"initialReward"`
	CurrentRewardSource      float64         `gorm:"type:double precision" json:"currentRewardSource"`
	CurrentRewardTarget      float64         `gorm:"type:double precision" json:"currentRewardTarget"`
	CurrentRewardUser        float64         `gorm:"type:double precision" json:"currentRewardUser"`
	CurrencyExchangeRateUUID uuid.UUID       `gorm:"type:uuid" json:"currencyExchangeRateUUID"`
	CurrencySource           string          `gorm:"type:text" json:"currencySource"`
	CurrencyTarget           string          `gorm:"type:text" json:"currencyTarget"`
	CurrencyUser             string          `gorm:"type:text" json:"currencyUser"`
	State                    string          `gorm:"type:varchar(20);check:state IN ('LIVE', 'STOPPED', 'EXPIRED', FINISHED', 'REQUESTED', 'PAID')" json:"state"`
	InitialPrice             float64         `gorm:"type:double precision" json:"initialPrice"`
	EndDate                  time.Time       `gorm:"type:date" json:"endDate"`
	AssetUnits               float64         `gorm:"type:float64" json:"assetUnits"`
	History                  []RewardHistory `gorm:"foreignKey:RewardUUID" json:"-"` // One-to-many relationship
	Type                     string          `gorm:"type:varchar(20);check:state IN ('INVESTMENT', 'FIXED')" json:"type"`
	Title                    string          `gorm:"type:text" json:"title"`
	Details                  string          `gorm:"type:text" json:"details"`
	OverridePrice            *float64        `gorm:"type:double precision" json:"overridePrice"`
	WithdrawalUuid           *uuid.UUID      `gorm:"type:uuid" json:"withdrawalUuid"`
	Origin                   string          `gorm:"origin:text" json:"origin"`
	BaseEntity

	Transaction          Transaction          `gorm:"foreignKey:TransactionUUID" json:"-"`
	CurrencyExchangeRate CurrencyExchangeRate `gorm:"foreignKey:CurrencyExchangeRateUUID" json:"-"`
	Withdrawal           Withdrawal           `gorm:"foreignKey:WithdrawalUuid" json:"-"`
}
