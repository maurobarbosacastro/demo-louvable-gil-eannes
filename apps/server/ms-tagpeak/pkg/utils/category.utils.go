package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
)

func CreateCategoryDtoToModel(c *dto.CreateCategoryDTO) models.Category {
	return models.Category{
		Name: &c.Name,
		Code: c.Code,
	}
}
