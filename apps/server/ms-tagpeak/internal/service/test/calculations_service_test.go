package test

import (
	"ms-tagpeak/internal/service"
	"testing"
)

// TestCalculateTransactionCashback calls service.CalculateTransactionCashback with a given storeCashbackValue, type INVESTMENT, transaction fee and transaction amount
// checking for a valid result
func TestCalculateTransactionCashback(t *testing.T) {
	t.Run("should calculate cashback", func(t *testing.T) {
		storeCashbackValue := float64(10)
		storeCashbackType := "INVESTMENT"
		transactionFee := float64(2)
		transactionAmount := float64(100)
		result := service.CalculateTransactionCashback(storeCashbackValue, storeCashbackType, transactionFee, transactionAmount)
		if result != float64(8) {
			t.Errorf("Expected %f, got %f", float64(8), result)
		}
	})
}

// TestRewardUnitInvested calls service.RewardUnitInvested with a given cashback, price day zero and expected result
// checking for a valid result
func TestRewardUnitInvested(t *testing.T) {
	t.Run("should calculate reward unit invested", func(t *testing.T) {
		cashback := float64(8)
		priceDayZero := float64(0.65)
		result := service.RewardUnitInvested(cashback, priceDayZero)
		if result != float64(12.31) {
			t.Errorf("Expected %f, got %f", float64(12.31), result)
		}
	})
}

// TestCalculateCurrentCashRewardValue calls service.CalculateCurrentCashRewardValue with a given transaction amount, asset units, investment price today and member percentage
// checking for a valid result
func TestCalculateCurrentCashRewardValue(t *testing.T) {
	t.Run("should calculate current cash reward value", func(t *testing.T) {
		transactionAmount := float64(100)
		assetUnits := float64(12.3)
		investmentPriceToday := float64(0.72)
		memberPercentage := float64(50)
		result := service.CalculateCurrentCashRewardValue(transactionAmount, assetUnits, investmentPriceToday, memberPercentage)
		if result != float64(4.43) {
			t.Errorf("Expected %f, got %f", float64(4.43), result)
		}
	})
}
