package utils

import (
	"encoding/json"
	"fmt"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
)

func UserPaymentMethodDtoToModel(dto *dto.CreateUserPaymentMethodDTO) (*models.UserPaymentMethod, error) {
	information, err := json.Marshal(MapInformation(dto))
	if err != nil {
		return nil, fmt.Errorf("failed to convert information: %w", err)
	}

	info := string(information)

	return &models.UserPaymentMethod{
		Information: &info,
	}, nil
}

func MapInformation(dtoParams *dto.CreateUserPaymentMethodDTO) dto.Information {
	return dto.Information{
		BankName:         dtoParams.BankName,
		BankAddress:      dtoParams.BankAddress,
		Country:          dtoParams.Country,
		BankAccountTitle: dtoParams.BankAccountTitle,
		Iban:             dtoParams.Iban,
		Vat:              dtoParams.Vat,
		BankCountry:      dtoParams.BankCountry,
	}
}
