package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"ms-tagpeak/pkg/logster"
	"strings"
)

// CalculateTransactionCashback
//
//	Investimento = (store.cashbackPercentage * Preço sem IVA) - (2% * Preço sem IVA)
//	Exemple:
//	(10% cashback * 100 euros preço sem iva ) - (2% * 100 euros preço sem iva) = 8 euros para investir
func CalculateTransactionCashback(storeCashbackValue float64, storeCashbackType string, transactionFee float64, transactionAmount float64) float64 {
	logster.StartFuncLog()
	cashbackValue := decimal.NewFromFloat(storeCashbackValue / 100)
	if strings.ToLower(storeCashbackType) == "fixed" {
		cashbackValue = decimal.NewFromFloat(storeCashbackValue).Div(decimal.NewFromFloat(100)).Mul(decimal.NewFromFloat(90))
		res, _ := cashbackValue.Round(2).Float64()
		return res
	}

	fee := decimal.NewFromFloat(transactionFee / 100)
	amount := decimal.NewFromFloat(transactionAmount)

	commission := cashbackValue.Mul(amount)

	res, _ := commission.Sub(fee.Mul(amount)).Round(2).Float64()
	logster.EndFuncLogMsg(fmt.Sprintf("storeCashbackValue: %f, storeCashbackType: %s, transactionFee: %f, transactionAmount: %f - Result: %f", storeCashbackValue, storeCashbackType, transactionFee, transactionAmount, res))
	return res
}

// RewardUnitInvested
//
// Unidades = Investimento / Preço do ativo no dia 0
// Exemple:
// 8 euros para investir / 0.65 (preço de uma unidade do ativo financeiro que queremos comprar) = 12.3 unidades
func RewardUnitInvested(cashback float64, priceDayZero float64) float64 {
	logster.StartFuncLog()
	dCashback := decimal.NewFromFloat(cashback)
	dPriceDayZero := decimal.NewFromFloat(priceDayZero)

	res, _ := dCashback.Div(dPriceDayZero).Round(2).Float64()
	logster.EndFuncLogMsg(fmt.Sprintf("RewardUnitInvested - cashback: %f, priceDayZero: %f - Result: %f", cashback, priceDayZero, res))
	return res
}

// CalculateCurrentCashRewardValue
//
// Current Reward = max (0.5% * Preço sem IVA ; Unidades * Preço ativo no dia x * Member Status %)
// Exemple:
// Queremos dar sempre um mínimo garantido ao cliente de 0.5%, caso o investimento corra mal. Logo é o valor máximo entre 0.5% do preço sem iva, ou as unidades (12.3) * 0.72 (preço dia x) * 50% (dependendo do loyalty status de cliente, 50, 55, ou 60%) = 4.4 euros
func CalculateCurrentCashRewardValue(transactionAmount float64, assetUnits float64, investmentPriceToday float64, memberPercentage float64) float64 {
	logster.StartFuncLog()
	amount := decimal.NewFromFloat(transactionAmount)
	//Divided by 10000 because we get 50% so divide by 100 to get 0.5% then by 100 remove percentages
	minPercentage := decimal.NewFromFloat(memberPercentage / 10000)
	minAmount := amount.Mul(minPercentage)

	rewardAmount, _ := minAmount.Round(2).Float64()

	units := decimal.NewFromFloat(assetUnits)
	todayPrice := decimal.NewFromFloat(investmentPriceToday)
	memberPerc := decimal.NewFromFloat(memberPercentage / 100)

	currentCashReward := units.Mul(todayPrice).Mul(memberPerc)

	if currentCashReward.GreaterThan(minAmount) {
		rewardAmount, _ = currentCashReward.Round(2).Float64()
	}

	logster.EndFuncLogMsg(fmt.Sprintf("CalculateCurrentCashRewardValue - transactionAmount: %f, assetUnits: %f, investmentPriceToday: %f, memberPercentage: %f - Result: %f", transactionAmount, assetUnits, investmentPriceToday, memberPercentage, rewardAmount))
	return rewardAmount
}
