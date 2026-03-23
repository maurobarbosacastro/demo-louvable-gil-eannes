package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
)

func RewardHistoryDtoToModel(c *dto.RewardHistoryDTO) models.RewardHistory {
	return models.RewardHistory{
		RewardUUID: *c.RewardUUID,
		Rate:       *c.Rate,
		Units:      *c.Units,
		CashReward: *c.CashReward,
	}
}
