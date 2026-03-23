package utils

import (
	"fmt"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/logster"
)

func CurrencyExchangeRateDtoToModel(c *dto.CreateCurrencyExchangeRateDTO) models.CurrencyExchangeRate {
	return models.CurrencyExchangeRate{
		Base:  &c.Base,
		Rates: c.Rates,
	}
}

func GetAmountByCurrencyRate(rateSource float64, value float64, rate float64) float64 {
	logster.Info(fmt.Sprintf("GetRewardByCurrencyRate - rateSource: %f | value: %f | rate: %f", rateSource, value, rate))
	return (value * rate) / rateSource
}
