package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	TransactionStateTracked   = "TRACKED"
	TransactionStateValidated = "VALIDATED"
	TransactionStateRejected  = "REJECTED"
)

type Transaction struct {
	Uuid                     uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	AmountSource             float64    `gorm:"type:double precision" json:"amountSource"`
	AmountTarget             float64    `gorm:"type:double precision" json:"amountTarget"`
	AmountUser               float64    `gorm:"type:double precision" json:"amountUser"`
	CurrencySource           string     `gorm:"type:text" json:"currencySource"`
	CurrencyTarget           string     `gorm:"type:text" json:"currencyTarget"`
	State                    string     `gorm:"type:varchar(20);check:state IN ('TRACKED', 'VALIDATED', 'REJECTED')" json:"state"`
	CommissionSource         float64    `gorm:"type:double precision" json:"commissionSource"`
	CommissionTarget         float64    `gorm:"type:double precision" json:"commissionTarget"`
	CommissionUser           float64    `gorm:"type:double precision" json:"commissionUser"`
	OrderDate                time.Time  `gorm:"type:date" json:"orderDate"`
	User                     string     `gorm:"type:text" json:"user"`
	StoreUUID                *uuid.UUID `gorm:"type:uuid" json:"storeUuid"`
	StoreVisitUUID           *uuid.UUID `gorm:"type:uuid" json:"."`
	CurrencyExchangeRateUUID uuid.UUID  `gorm:"type:uuid" json:"currencyExchangeRateUuid"`
	IsProcessed              bool       `gorm:"type:boolean" json:"-"`
	Cashback                 float64    `gorm:"type:double precision" json:"cashback"`
	SourceId                 string     `gorm:"type:text" json:"sourceId"`
	ManualCommission         *float64   `gorm:"type:double precision" json:"manualCommission"`
	BaseEntity

	Store                *Store               `gorm:"foreignKey:StoreUUID;" json:"-"`
	StoreVisit           *StoreVisit          `gorm:"foreignKey:StoreVisitUUID;" json:"storeVisit"`
	CurrencyExchangeRate CurrencyExchangeRate `gorm:"foreignKey:CurrencyExchangeRateUUID" json:"-"`
}

func (f *Transaction) AfterFind(tx *gorm.DB) (err error) {
	if f.ManualCommission != nil {
		f.CommissionTarget = *f.ManualCommission
	}
	return
}

type TransactionStateUpdate struct {
	TrxUUID             string    `gorm:"column:trx_uuid"`
	TrxSourceID         string    `gorm:"column:trx_source_id"`
	TrxAmountSource     float64   `gorm:"column:trx_amount_source"`
	TrxCurrencySource   string    `gorm:"column:trx_currency_source"`
	TrxCommissionSource float64   `gorm:"column:trx_commission_source"`
	TrxOrderDate        time.Time `gorm:"column:trx_order_date"`
	TrxStoreVisitUUID   uuid.UUID `gorm:"column:trx_store_visit_uuid"`
	TrxUser             string    `gorm:"column:trx_user"`
}

type TransactionsCurrencyUserUpdate struct {
	TransactionUUID         string  `json:"transaction_uuid"`
	OriginalAmountUser      float64 `json:"original_amount_user"`
	ConvertedAmountUser     float64 `json:"converted_amount_user"`
	OriginalCommissionUser  float64 `json:"original_commission_user"`
	ConvertedCommissionUser float64 `json:"converted_commission_user"`
}
