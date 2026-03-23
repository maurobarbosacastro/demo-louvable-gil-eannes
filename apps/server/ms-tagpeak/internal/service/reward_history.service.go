package service

import (
	"errors"
	"github.com/google/uuid"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"sort"
)

func GetRewardHistory(rewardUuid uuid.UUID, pagParams pagination.PaginationParams) (*pagination.PaginationResult, error) {
	res, err := repository.GetRewardHistory(rewardUuid, pagParams)
	if err != nil {
		return nil, err
	}

	var aux []dto.RewardHistoryDTO

	if dataSlice, ok := res.Data.([]models.RewardHistory); ok {
		for _, item := range dataSlice {
			aux = append(aux, dto.RewardHistoryDTO{
				UUID:       &item.UUID,
				Rate:       &item.Rate,
				Units:      &item.Units,
				CashReward: &item.CashReward,
				CreatedAt:  &item.CreatedAt,
			})
		}

		res.Data = aux
	} else {
		return nil, errors.New("res.Data is not of type []models.Store")
	}

	return &res, nil
}

func CreateRewardHistory(dtoParam dto.RewardHistoryDTO, uuidUser string) (*models.RewardHistory, error) {

	model := utils.RewardHistoryDtoToModel(&dtoParam)
	model.CreatedBy = uuidUser

	res, err := repository.CreateRewardHistory(model)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetRewardHistoryGraph(rewardUuid uuid.UUID) ([]dto.RewardHistoryDTO, error) {
	res, err := repository.GetRewardHistoryGraph(rewardUuid)
	if err != nil {
		return nil, err
	}

	latestByDate := make(map[string]models.RewardHistory)
	dateFormat := "2006-01-02"

	for _, rewardHistory := range res {
		dateStr := rewardHistory.CreatedAt.Format(dateFormat)

		// Update the map with new rewardHistory or update the rewardHistory with newest rewardHistory
		if existing, exists := latestByDate[dateStr]; !exists ||
			rewardHistory.CreatedAt.After(existing.CreatedAt) {
			latestByDate[dateStr] = rewardHistory
		}
	}

	aux := make([]dto.RewardHistoryDTO, 0, len(latestByDate))

	for _, item := range latestByDate {
		aux = append(aux, dto.RewardHistoryDTO{
			UUID:       &item.UUID,
			Rate:       &item.Rate,
			Units:      &item.Units,
			CashReward: &item.CashReward,
			CreatedAt:  &item.CreatedAt,
		})
	}

	sort.Slice(aux, func(i, j int) bool {
		return aux[i].CreatedAt.Before(*aux[j].CreatedAt)
	})

	return aux, nil
}

func RunCopyLatestRewards(rewardUuids []string) error {
	return repository.RunCopyLatestRewards(rewardUuids)
}
