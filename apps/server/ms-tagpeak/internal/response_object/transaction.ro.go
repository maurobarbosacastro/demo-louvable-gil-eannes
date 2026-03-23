package response_object

import (
	"time"

	"github.com/google/uuid"
)

type StoreCashbackRO struct {
	Uuid              uuid.UUID `json:"uuid,omitempty"`
	Name              string    `json:"name,omitempty"`
	Logo              *string   `json:"logo,omitempty"`
	PercentageCashout *float64  `json:"percentageCashout,omitempty"`
	CashbackValue     *float64  `json:"cashbackValue,omitempty"`
	CashbackType      *string   `json:"cashbackType,omitempty"`
}

type CashbackRO struct {
	Uuid              uuid.UUID         `json:"uuid,omitempty"`
	ExitId            *string           `json:"exitId,omitempty"`
	Store             *StoreCashbackRO  `json:"store,omitempty"`
	UserUuid          string            `json:"userUuid,omitempty"`
	UserName          string            `json:"userName,omitempty"`
	Email             string            `json:"email,omitempty"`
	Date              time.Time         `json:"date,omitempty"`
	AmountSource      float64           `json:"amountSource"`
	AmountTarget      float64           `json:"amountTarget"`
	AmountUser        float64           `json:"amountUser"`
	CurrencySource    string            `json:"currencySource"`
	CurrencyTarget    string            `json:"currencyTarget"`
	NetworkCommission float64           `json:"networkCommission,omitempty"`
	Status            string            `json:"status,omitempty"`
	Reward            *RewardCashbackRO `json:"reward"`
	Cashback          float64           `json:"cashback"`
}

func (c CashbackRO) Shopify() CashbackRO {
	c.Reward = nil
	return c
}

type CashbackDetailRO struct {
	Uuid                     uuid.UUID        `json:"uuid"`
	AmountSource             float64          `json:"amountSource"`
	AmountTarget             float64          `json:"amountTarget"`
	AmountUser               float64          `json:"amountUser"`
	CurrencySource           string           `json:"currencySource"`
	CurrencyTarget           string           `json:"currencyTarget"`
	State                    string           `json:"state"`
	CommissionSource         float64          `json:"commissionSource"`
	CommissionTarget         float64          `json:"commissionTarget"`
	CommissionUser           float64          `json:"commissionUser"`
	OrderDate                time.Time        `json:"orderDate"`
	User                     string           `json:"user"`
	Store                    *StoreCashbackRO `json:"store,omitempty"`
	StoreVisitUUID           *uuid.UUID       `json:"-"`
	StoreVisit               *StoreVisitRO    `json:"storeVisit"`
	CurrencyExchangeRateUUID uuid.UUID        `json:"currencyExchangeRateUuid"`
	Cashback                 float64          `json:"cashback"`

	BaseEntityRO
}

type CashbackViewRO struct {
	TransactionUUID        uuid.UUID  `json:"transactionUuid"`
	User                   string     `json:"user"`
	Email                  string     `json:"email"`
	StoreVisitUUID         uuid.UUID  `json:"storeVisitUuid"`
	RefId                  *string    `json:"refId"`
	StoreUUID              *uuid.UUID `json:"storeUuid"`
	StoreName              *string    `json:"storeName"`
	StoreLogo              *string    `json:"storeLogo"`
	Date                   time.Time  `json:"date"`
	AmountSource           float64    `json:"amountSource"`
	AmountTarget           float64    `json:"amountTarget"`
	AmountUser             float64    `json:"amountUser"`
	CurrencySource         string     `json:"currencySource"`
	CurrencyTarget         string     `json:"currencyTarget"`
	NetworkCommission      float64    `json:"networkCommission"`
	Status                 string     `json:"status"`
	Cashback               float64    `json:"cashback"`
	RewardUUID             uuid.UUID  `json:"rewardUuid"`
	ISIN                   string     `json:"isin"`
	Conid                  string     `json:"conid"`
	CurrentRewardSource    float64    `json:"currentRewardSource"`
	CurrentRewardTarget    float64    `json:"currentRewardTarget"`
	CurrentRewardUser      float64    `json:"currentRewardUser"`
	InitialPrice           float64    `json:"initialPrice"`
	Title                  string     `json:"title"`
	EndDate                time.Time  `json:"endDate"`
	StartDate              time.Time  `json:"startDate"`
	Origin                 string     `json:"origin"`
	StorePercentageCashout float64    `json:"storePercentageCashout"`
	StoreCashbackValue     float64    `json:"storeCashbackValue"`
	StoreCashbackType      string     `json:"storeCashbackType"`
	StoppedAt              *time.Time `json:"stoppedAt"`
	StoreVisitReference    *string    `json:"storeVisitReference"`
}

type ShopifyActionableTransactionRO struct {
	TransactionUUID             uuid.UUID `json:"transactionUuid"`
	StoreUUID                   uuid.UUID `json:"storeUuid"`
	StoreName                   string    `json:"storeName"`
	TransactionState            string    `json:"transactionState"`
	OrderDate                   time.Time `json:"orderDate"`
	AverageRewardActivationTime string    `json:"averageRewardActivationTime"`
	RemainingTime               int       `json:"remainingTime"`
}
