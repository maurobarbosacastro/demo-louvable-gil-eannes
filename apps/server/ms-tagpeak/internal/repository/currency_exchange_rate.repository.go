package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
)

func GetCurrencyExchangeRate(uuid uuid.UUID) (models.CurrencyExchangeRate, error) {
	dbInstance := db.GetDB()

	var model models.CurrencyExchangeRate
	// Execute the query and capture the error
	err := dbInstance.Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.CurrencyExchangeRate{}, err
	}

	return model, nil
}

func GetAllCurrencyExchangeRates() ([]models.CurrencyExchangeRate, error) {

	dbInstance := db.GetDB()
	var model []models.CurrencyExchangeRate

	err := dbInstance.Find(&model).Error
	if err != nil {
		return []models.CurrencyExchangeRate{}, err
	}

	return model, nil
}

func CreateCurrencyExchangeRate(model models.CurrencyExchangeRate) (models.CurrencyExchangeRate, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.CurrencyExchangeRate{}, err
	}

	return model, nil
}

func UpdateCurrencyExchangeRate(model models.CurrencyExchangeRate) (models.CurrencyExchangeRate, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&model).Error
	if err != nil {
		return models.CurrencyExchangeRate{}, err
	}

	return model, nil
}

func DeleteCurrencyExchangeRate(uuid uuid.UUID, user string) error {
	dbInstance := db.GetDB()
	err := dbInstance.Model(&models.CurrencyExchangeRate{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.CurrencyExchangeRate{}, "uuid = ?", uuid).Error

	if err != nil {
		return err
	}

	return nil
}

func GetLatestCurrencyExchangeRate() (*models.CurrencyExchangeRate, error) {
	dbInstance := db.GetDB()

	var latestCurrencyExchangeRate models.CurrencyExchangeRate
	err := dbInstance.Model(&models.CurrencyExchangeRate{}).Order("created_at desc").Last(&latestCurrencyExchangeRate).Error
	if err != nil {
		return nil, err
	}
	return &latestCurrencyExchangeRate, nil
}
