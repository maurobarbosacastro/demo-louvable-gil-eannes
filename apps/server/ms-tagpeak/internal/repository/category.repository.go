package service

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"strings"
)

func GetCategory(uuid uuid.UUID) (models.Category, error) {
	dbInstance := db.GetDB()

	var category models.Category
	// Execute the query and capture the error
	err := dbInstance.Where("uuid = ?", uuid).First(&category).Error
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func GetCategoryByCode(code string) (models.Category, error) {
	dbInstance := db.GetDB()

	var category models.Category
	// Execute the query and capture the error
	err := dbInstance.Where("code = ?", code).First(&category).Error
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func GetAllCategoriesWithPagination(pagDTO pagination.PaginationParams, filters *dto.CategoryFiltersDTO) (*pagination.PaginationResult, error) {
	log.Info("*** START repository.GetAllCategoriesWithPagination ***")
	dbInstance := db.GetDB()
	var categories []models.Category
	var res pagination.PaginationResult

	// Set pagination details
	res.Limit = pagDTO.Limit
	res.Page = pagDTO.Page
	res.Sort = pagDTO.Sort

	dbWithFilters := dbInstance

	if pagDTO.Sort != "" {
		dbWithFilters = dbWithFilters.Order(pagDTO.Sort)
	}

	if filters.Name != "" {
		dbWithFilters = dbWithFilters.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(filters.Name)+"%")
	}

	err := dbWithFilters.Scopes(pagination.Paginate(&categories, &res, nil, dbInstance)).Find(&categories).Error

	res.Data = categories
	log.Infof("*** Found %v categories ***", len(categories))

	if err != nil {
		log.Errorf("*** Occurred an error getting categories: %v ***", err)
		return &pagination.PaginationResult{}, err
	}

	log.Info("*** END repository.GetAllCategoriesWithPagination ***")
	return &res, nil
}

func CreateCategory(category models.Category) (models.Category, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&category).Error
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func UpdateCategory(category models.Category) (models.Category, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&category).Error
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func DeleteCategory(uuid uuid.UUID, user string) (models.Category, error) {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Category{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.Category{}, "uuid = ?", uuid).Error
	if err != nil {
		return models.Category{}, err
	}

	return models.Category{}, nil
}

func CategoryCodeExist(code string) (bool, error) {
	dbInstance := db.GetDB()

	count := int64(0)
	// Execute the query and capture the error
	err := dbInstance.Model(&models.Category{}).
		Scopes(utils.ActiveScope()).
		Where("code = ?", code).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
