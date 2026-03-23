package service

import (
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
)

func CreateConfiguration(c models.Configuration) (models.Configuration, error) {
	dbInstance := db.GetDB()

	if err := dbInstance.Create(&c).Error; err != nil {
		return models.Configuration{}, err
	}
	return c, nil
}

func ConfigurationExistsByCode(code string) (bool, error) {
	dbInstance := db.GetDB()

	var saved bool
	err := dbInstance.Model(&models.Configuration{}).
		Select("1").
		Where("code = ?", code).
		Scan(&saved).
		Error

	if err != nil {
		return false, err
	}

	return saved, nil
}

func GetAllConfigurations() ([]models.Configuration, error) {
	dbInstance := db.GetDB()
	var configurations []models.Configuration

	err := dbInstance.Find(&configurations).Error
	if err != nil {
		return nil, err
	}
	return configurations, nil
}

func GetConfigurations(pag pagination.PaginationParams) (*pagination.PaginationResult, error) {
	dbInstance := db.GetDB()
	var configurations []models.Configuration
	var res pagination.PaginationResult

	res.Limit = pag.Limit
	res.Page = pag.Page
	res.Sort = pag.Sort

	if err := dbInstance.Scopes(pagination.Paginate(&configurations, &res, nil, dbInstance)).Find(&configurations).Error; err != nil {
		return &pagination.PaginationResult{}, err
	}
	res.Data = configurations

	return &res, nil
}

func GetConfiguration(id int) (models.Configuration, error) {
	dbInstance := db.GetDB()

	var configuration models.Configuration
	err := dbInstance.Where("id = ?", id).First(&configuration).Error
	if err != nil {
		return models.Configuration{}, err
	}

	return configuration, nil
}

func GetConfigurationByCode(code string) (models.Configuration, error) {
	dbInstance := db.GetDB()

	var configuration models.Configuration
	err := dbInstance.Where("code = ?", code).First(&configuration).Error
	if err != nil {
		return models.Configuration{}, err
	}

	return configuration, nil
}

func UpdateConfiguration(update models.Configuration) (models.Configuration, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&update).Error
	if err != nil {
		return models.Configuration{}, err
	}

	return update, nil

}

func DeleteConfiguration(id int, user string) (models.Configuration, error) {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Configuration{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.Configuration{}, "id = ?", id).Error

	if err != nil {
		return models.Configuration{}, err
	}

	return models.Configuration{}, nil
}
