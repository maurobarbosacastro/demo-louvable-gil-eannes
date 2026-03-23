package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
)

func GetPartner(uuid uuid.UUID) (models.Partner, error) {
	dbInstance := db.GetDB()

	var model models.Partner
	// Execute the query and capture the error
	err := dbInstance.Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.Partner{}, err
	}

	return model, nil
}

func GetPartnerByCode(code *string) (models.Partner, error) {
	dbInstance := db.GetDB()

	var model models.Partner
	// Execute the query and capture the error
	err := dbInstance.Where("code = ?", code).First(&model).Error
	if err != nil {
		return models.Partner{}, err
	}

	return model, nil
}

func GetAllPartners() ([]models.Partner, error) {

	dbInstance := db.GetDB()
	var model []models.Partner

	err := dbInstance.Find(&model).Error
	if err != nil {
		return []models.Partner{}, err
	}

	return model, nil
}

func GetAllPartnersWithPagination(pagDTO pagination.PaginationParams, filters dto.PartnerFiltersDTO) (*pagination.PaginationResult, error) {

	dbInstance := db.GetDB()
	var models []models.Partner
	var res pagination.PaginationResult

	res.Limit = pagDTO.Limit
	res.Page = pagDTO.Page
	res.Sort = pagDTO.Sort

	err := dbInstance.Scopes(pagination.Paginate(&models, &res, filters, dbInstance)).Find(&models).Error
	if err != nil {
		return nil, err
	}
	res.Data = models

	return &res, nil
}

func CreatePartner(model models.Partner) (models.Partner, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.Partner{}, err
	}

	return model, nil
}

func UpdatePartner(model models.Partner) (models.Partner, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&model).Error
	if err != nil {
		return models.Partner{}, err
	}

	return model, nil
}

func DeletePartner(uuid uuid.UUID, user string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Partner{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.Partner{}, "uuid = ?", uuid).Error

	if err != nil {
		return err
	}

	return nil
}

func CodeAlreadyExists(code string) (bool, error) {
	dbInstance := db.GetDB()

	count := int64(0)
	// Execute the query and capture the error
	err := dbInstance.Model(&models.Partner{}).
		Scopes(utils.ActiveScope()).
		Where("code = ?", code).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
