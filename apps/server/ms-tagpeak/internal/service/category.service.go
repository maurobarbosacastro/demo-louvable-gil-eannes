package service

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"strings"
)

func GetCategory(uuid uuid.UUID) (models.Category, error) {
	res, err := repository.GetCategory(uuid)
	if err != nil {
		return models.Category{}, err
	}
	return res, nil
}

func GetCategoryByCode(code string) (models.Category, error) {
	res, err := repository.GetCategoryByCode(code)
	if err != nil {
		return models.Category{}, err
	}
	return res, nil
}

func GetAllCategories(pag pagination.PaginationParams, filters *dto.CategoryFiltersDTO) (*pagination.PaginationResult, error) {
	log.Info("*** START service.GetAllCategories ***")
	res, err := repository.GetAllCategoriesWithPagination(pag, filters)
	if err != nil {
		return &pagination.PaginationResult{}, err
	}

	log.Info("*** END service.GetAllCategories ***")
	return res, nil
}

func CreateCategory(dtoParam dto.CreateCategoryDTO, uuidUser string) (models.Category, error) {

	if dtoParam.Code == nil {
		catCode := strings.ToLower(dtoParam.Name)
		catCode = strings.ReplaceAll(catCode, " ", "_")
		dtoParam.Code = &catCode
	}

	model := utils.CreateCategoryDtoToModel(&dtoParam)
	model.CreatedBy = uuidUser

	res, err := repository.CreateCategory(model)
	if err != nil {
		return models.Category{}, err
	}
	return res, nil
}

func UpdateCategory(dtoParam dto.UpdateCategoryDTO, uuid uuid.UUID, uuidUser string) (models.Category, error) {

	toUpdate, err := repository.GetCategory(uuid)
	if err != nil {
		return models.Category{}, err
	}

	if dtoParam.Name != nil {
		toUpdate.Name = dtoParam.Name
	}

	toUpdate.UpdatedBy = &uuidUser

	res, err := repository.UpdateCategory(toUpdate)
	if err != nil {
		return models.Category{}, err
	}
	return res, nil
}

func DeleteCategory(uuid uuid.UUID, uuidUser string) error {
	_, err := repository.DeleteCategory(uuid, uuidUser)
	if err != nil {
		return err
	}
	return nil
}

func CategoryCodeExist(code string) (bool, error) {
	exists, err := repository.CategoryCodeExist(code)
	if err != nil {
		return false, err
	}
	return exists, nil
}
