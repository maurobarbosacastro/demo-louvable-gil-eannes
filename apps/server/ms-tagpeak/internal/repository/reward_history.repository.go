package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
)

func GetRewardHistory(rewardUuid uuid.UUID, pagDTO pagination.PaginationParams) (pagination.PaginationResult, error) {
	dbInstance := db.GetDB()

	var resPag pagination.PaginationResult
	var history []models.RewardHistory

	filters := map[string]interface{}{
		"reward_uuid": rewardUuid,
	}

	resPag.Limit = pagDTO.Limit
	resPag.Page = pagDTO.Page
	resPag.Sort = pagDTO.Sort

	if err := dbInstance.Scopes(pagination.Paginate(&models.RewardHistory{}, &resPag, filters, dbInstance)).
		Where("reward_uuid = ?", rewardUuid).
		Find(&history).Error; err != nil {
		resPag.Data = nil
		return resPag, err
	}
	resPag.Data = history
	return resPag, nil
}

func CreateRewardHistory(model models.RewardHistory) (models.RewardHistory, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.RewardHistory{}, err
	}

	return model, nil
}

func GetRewardHistoryGraph(rewardUuid uuid.UUID) ([]models.RewardHistory, error) {
	dbInstance := db.GetDB()

	var history []models.RewardHistory

	err := dbInstance.Where("reward_uuid = ?", rewardUuid).
		Order("created_at asc").
		Find(&history).Error
	if err != nil {
		return nil, err
	}

	return history, nil
}

func RunCopyLatestRewards(rewardUuids []string) error {
	dbInstance := db.GetDB()

	sql := `SELECT copy_latest_rewards(ARRAY[`
	for i, rewardUuid := range rewardUuids {
		if i > 0 {
			sql += ", "
		}
		sql += "'" + rewardUuid + "'::uuid"
	}
	sql += "])"

	err := dbInstance.Exec(sql).Error

	if err != nil {
		return err
	}

	return nil
}
