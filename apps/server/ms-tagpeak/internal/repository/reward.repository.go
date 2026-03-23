package service

import (
	"fmt"
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"time"
)

func GetReward(uuid uuid.UUID) (models.Reward, error) {
	dbInstance := db.GetDB()

	var model models.Reward
	// Use Preload to load the History relationship
	err := dbInstance.Preload("History").Preload("CurrencyExchangeRate").Preload("Transaction").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.Reward{}, err
	}

	return model, nil
}

func GetAllRewards() ([]models.Reward, error) {

	dbInstance := db.GetDB()
	var model []models.Reward

	// Use Preload to load the History relationship
	err := dbInstance.Preload("History").Preload("CurrencyExchangeRate").Preload("Transaction").Find(&model).Error
	if err != nil {
		return []models.Reward{}, err
	}

	return model, nil
}

func CreateReward(model models.Reward) (models.Reward, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.Reward{}, err
	}

	return model, nil
}

func UpdateReward(model models.Reward) (models.Reward, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&model).Error
	if err != nil {
		return models.Reward{}, err
	}

	return model, nil
}

func DeleteReward(uuid uuid.UUID, user string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Reward{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.Reward{}, "uuid = ?", uuid).Error

	if err != nil {
		return err
	}

	return nil
}

func GetTransactionByReward(uuid uuid.UUID) (models.Transaction, error) {
	dbInstance := db.GetDB()

	var model models.Reward
	// Execute the query and capture the error
	err := dbInstance.Preload("Transaction").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.Transaction{}, err
	}

	return model.Transaction, nil
}

func GetCurrencyExchangeRateByReward(uuid uuid.UUID) (models.CurrencyExchangeRate, error) {
	dbInstance := db.GetDB()

	var model models.Reward
	// Execute the query and capture the error
	err := dbInstance.Preload("CurrencyExchangeRate").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.CurrencyExchangeRate{}, err
	}

	return model.CurrencyExchangeRate, nil
}

func GetExpiredRewards() (*[]string, error) {
	dbInstance := db.GetDB()

	var res []string
	// Use Preload to load the History relationship
	today := time.Now().Format("2006-01-02")
	err := dbInstance.Model(&models.Reward{}).Select("uuid").Where("end_date = ? and state='LIVE'", today).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func BulkUpdateRewards(uuids []string, data map[string]interface{}) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Reward{}).
		Where("uuid IN ?", uuids). // Filter rows by UUIDs
		Updates(data).             // Provide fields to update
		Error

	if err != nil {
		fmt.Printf("Error updating records: %v\n", err)
		return err
	} else {
		fmt.Println("Records updated successfully")
		return nil
	}
}

func IsRewardsWithTransactionSaved(transactionUUID uuid.UUID) (bool, error) {
	dbInstance := db.GetDB()

	var saved bool
	err := dbInstance.Model(&models.Reward{}).
		Select("1").
		Where("transaction_uuid = ?", transactionUUID).
		Scan(&saved).
		Error

	if err != nil {
		return false, err
	}

	return saved, nil
}

func BulkTransactionRewardUpdate(req []map[string]interface{}) error {
	dbInstance := db.GetDB()

	tx := dbInstance.Begin()

	for _, reward := range req {
		err := tx.Model(&models.Reward{}).
			Where("uuid = ?", reward["uuid"]).
			Updates(reward).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func GetSumReferralRewards(uuidUser string) (float64, error) {
	dbInstance := db.GetDB()

	var rewards []uuid.UUID

	var res float64
	err := dbInstance.Model(&models.Reward{}).
		Select("uuid").
		Where("\"user\" = ?", uuidUser).
		Where("state = 'FINISHED'").
		Scan(&rewards).Error

	if err != nil {
		return 0, err
	}

	err = dbInstance.Model(&models.ReferralRevenueHistory{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("reward_uuid IN (?)", rewards).
		Scan(&res).Error

	if err != nil {
		return 0, err
	}

	return res, nil
}

func CreateBulkRewards(rewardModels []models.Reward) error {
	dbInstance := db.GetDB()

	var err error

	err = dbInstance.Model(&models.Reward{}).Create(rewardModels).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserLiveRewardSum(userUuid string) (float64, error) {
	dbInstance := db.GetDB()

	var sum float64

	// Perform the query to calculate the sum
	err := dbInstance.Model(&models.Reward{}).
		Select("COALESCE(SUM(current_reward_user), 0)").
		Where("\"user\" = ? AND state = ?", userUuid, "LIVE").
		Scan(&sum).Error

	if err != nil {
		return 0, fmt.Errorf("error calculating sum: %w", err)
	}

	return sum, nil
}

func GetRewardsByStateAndUser(state string, userUuid string) ([]uuid.UUID, error) {
	var rewards []uuid.UUID
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Reward{}).
		Select("reward.uuid").
		Where("state = ? AND \"user\" = ?", state, userUuid).
		Scan(&rewards).Error

	if err != nil {
		return nil, fmt.Errorf("error getting rewards: %w", err)
	}

	return rewards, nil
}

func UpdateRewardsState(state string, rewards []uuid.UUID, withdrawalUuid string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Reward{}).
		Where("uuid IN ?", rewards).
		Updates(map[string]interface{}{
			"state":           state,
			"withdrawal_uuid": withdrawalUuid,
		}).Error

	if err != nil {
		return fmt.Errorf("error updating rewards: %w", err)
	}

	return nil
}

func GetUsersWithTotalAmountReward() ([]response_object.UsersWithTotalAmountReward, error) {
	dbInstance := db.GetDB()

	var res []response_object.UsersWithTotalAmountReward

	err := dbInstance.Raw(`select "user", coalesce(sum(current_reward_user), 0) as total from reward where reward.legacy_id is not null and state = 'FINISHED' group by "user"`).
		Scan(&res).Error

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetRewardsByState(state string) ([]models.Reward, error) {
	dbInstance := db.GetDB()

	var res []models.Reward

	err := dbInstance.
		Preload("Transaction").
		Where("state = ?", state).
		Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateCurrentRewardValue(
	rewardUuid uuid.UUID,
	valueSource float64,
	valueUser float64,
	valueTarget float64) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Reward{}).
		Where("uuid = ?", rewardUuid).
		Update("current_reward_user", valueUser).
		Update("current_reward_source", valueSource).
		Update("current_reward_target", valueTarget).
		Error

	if err != nil {
		return err
	}

	return nil
}

func GetRewardsByCurrencies() ([]response_object.RewardByCurrencies, error) {
	dbInstance := db.GetDB()

	var res []response_object.RewardByCurrencies

	sql := `SELECT * FROM currency_dashboard`
	err := dbInstance.Raw(sql).
		Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
