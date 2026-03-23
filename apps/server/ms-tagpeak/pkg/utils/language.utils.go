package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
)

func LanguageDtoToModel(c *dto.CreateLanguageDTO) models.Language {
	return models.Language{
		Name: &c.Name,
		Code: &c.Code,
	}
}
