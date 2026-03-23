package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
)

func CountryDtoToModel(c *dto.CreateCountryDTO) models.Country {
	return models.Country{
		Abbreviation: &c.Abbreviation,
		Currency:     &c.Currency,
		Flag:         &c.Flag,
		Name:         &c.Name,
		Enabled:      &c.Enabled,
	}
}
