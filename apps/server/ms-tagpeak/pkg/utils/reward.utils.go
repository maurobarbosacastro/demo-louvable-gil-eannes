package utils

import (
	"fmt"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/logster"
)

func RewardDtoToModel(c *dto.CreateRewardDTO) models.Reward {
	var title string
	var details string
	var currentRewardUser float64

	if c.Title != nil {
		title = *c.Title
	}

	if c.Details != nil {
		details = *c.Details
	}

	if c.CurrentRewardUser != nil {
		currentRewardUser = *c.CurrentRewardUser
	}

	return models.Reward{
		TransactionUUID:   c.TransactionUUID,
		Isin:              c.Isin,
		Conid:             c.Conid,
		CurrencySource:    c.CurrencySource,
		State:             c.State,
		InitialPrice:      c.InitialPrice,
		EndDate:           c.EndDate,
		Type:              c.Type,
		Title:             title,
		Details:           details,
		CurrentRewardUser: currentRewardUser,
		Origin:            c.Origin,
	}
}

func RewardBulkDtoToModel(c *dto.CreateRewardBulkDTO) []models.Reward {
	var title string
	var details string

	if c.Title != nil {
		title = *c.Title
	}

	if c.Details != nil {
		details = *c.Details
	}

	var rewards []models.Reward

	for i := range c.TransactionUUIDs {
		rewards = append(rewards, models.Reward{
			TransactionUUID: c.TransactionUUIDs[i],
			Isin:            c.Isin,
			Conid:           c.Conid,
			CurrencySource:  c.CurrencySource,
			State:           c.State,
			InitialPrice:    c.InitialPrice,
			EndDate:         c.EndDate,
			Type:            c.Type,
			Title:           title,
			Details:         details,
			Origin:          c.Origin,
		})
	}

	return rewards

}

func MapRewardRO(reward models.Reward) response_object.RewardCashbackRO {
	return response_object.RewardCashbackRO{
		Uuid:                &reward.Uuid,
		Isin:                &reward.Isin,
		Conid:               &reward.Conid,
		CurrentRewardTarget: &reward.CurrentRewardTarget,
		CurrentRewardSource: &reward.CurrentRewardSource,
		CurrentRewardUser:   &reward.CurrentRewardUser,
		Status:              &reward.State,
		PriceDayZero:        &reward.InitialPrice,
		Title:               &reward.Title,
		EndDate:             &reward.EndDate,
		StartDate:           &reward.CreatedAt,
		Origin:              reward.Origin,
	}
}

func GetRewardByCurrencyRate(rateSource float64, value float64, rate float64) float64 {
	logster.Info(fmt.Sprintf("GetRewardByCurrencyRate - rateSource: %f | value: %f | rate: %f", rateSource, value, rate))
	return (value * rate) / rateSource
}
