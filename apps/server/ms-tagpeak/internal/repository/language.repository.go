package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
)

func GetLanguage(uuid uuid.UUID) (models.Language, error) {
	dbInstance := db.GetDB()

	var model models.Language
	// Execute the query and capture the error
	err := dbInstance.Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.Language{}, err
	}

	return model, nil
}

func GetAllLanguages(pagDTO pagination.PaginationParams) (*pagination.PaginationResult, error) {

	dbInstance := db.GetDB()
	var model []models.Language
	var res pagination.PaginationResult

	res.Limit = pagDTO.Limit
	res.Page = pagDTO.Page
	res.Sort = pagDTO.Sort

	err := dbInstance.Scopes(pagination.Paginate(&model, &res, nil, dbInstance)).Find(&model).Error
	res.Data = model

	if err != nil {
		return &pagination.PaginationResult{}, err
	}

	return &res, nil
}

func CreateLanguage(model models.Language) (models.Language, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.Language{}, err
	}

	return model, nil
}

func UpdateLanguage(model models.Language) (models.Language, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&model).Error
	if err != nil {
		return models.Language{}, err
	}

	return model, nil
}

func DeleteLanguage(uuid uuid.UUID, user string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Language{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.Language{}, "uuid = ?", uuid).Error

	if err != nil {
		return err
	}

	return nil
}

func LanguageCodeExist(code string) (bool, error) {
	dbInstance := db.GetDB()

	count := int64(0)
	// Execute the query and capture the error
	err := dbInstance.Model(&models.Language{}).
		Scopes(utils.ActiveScope()).
		Where("code = ?", code).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
