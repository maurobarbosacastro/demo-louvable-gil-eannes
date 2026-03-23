package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
)

func PaymentMethodDtoToModel(c *dto.CreatePaymentMethodDTO) models.PaymentMethod {
	return models.PaymentMethod{
		Name: c.Name,
		Code: c.Code,
	}
}
