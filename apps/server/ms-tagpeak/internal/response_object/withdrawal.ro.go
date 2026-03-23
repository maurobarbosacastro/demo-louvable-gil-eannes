package response_object

import (
	"time"
)

type WithdrawalRO struct {
	Uuid           string               `json:"uuid"`
	User           SimpleUserDto        `json:"user"`
	AmountSource   float64              `json:"amountSource,omitempty"`
	AmountTarget   float64              `json:"amountTarget"`
	CurrencyTarget *string              `json:"currencyTarget,omitempty"`
	Details        *string              `json:"details"`
	State          string               `json:"state"`
	CompletionDate *time.Time           `json:"completionDate"`
	CreatedAt      *time.Time           `json:"createdAt,omitempty"`
	PaymentMethod  *UserPaymentMethodRo `json:"paymentMethod,omitempty"`

	BaseEntityRO
}

type WithdrawalStatsRO struct {
	PaidWithdrawals float64 `json:"paidWithdrawals"`
	AmountRewards   float64 `json:"amountRewards"`
	AmountReferrals float64 `json:"amountReferrals"`
}
