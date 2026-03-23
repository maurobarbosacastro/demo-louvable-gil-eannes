package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
)

func CreatePartnerDtoToModel(c *dto.CreatePartnerDTO) models.Partner {
	return models.Partner{
		Name:               &c.Name,
		Code:               c.Code,
		ECommercePlatform:  &c.ECommercePlatform,
		ValidationPeriod:   &c.ValidationPeriod,
		DeepLink:           &c.DeepLink,
		DeepLinkIdentifier: &c.DeepLinkIdentifier,
		SubIdentifier:      &c.SubIdentifier,
		PercentageTagpeak:  &c.PercentageTagpeak,
		PercentageInvested: &c.PercentageInvested,
	}
}
