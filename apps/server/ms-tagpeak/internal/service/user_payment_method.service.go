package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
)

func CreateUserPaymentMethod(dto dto.CreateUserPaymentMethodDTO, paymentMethod *models.PaymentMethod, file *models.File, uuidUser string) (*models.UserPaymentMethod, error) {
	model, err := utils.UserPaymentMethodDtoToModel(&dto)
	if err != nil {
		return nil, err
	}

	model.User = uuidUser
	model.CreatedBy = uuidUser
	model.PaymentMethod = *paymentMethod
	model.File = *file
	model.State = "PENDING"

	res, err := repository.CreateUserPaymentMethod(*model)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func GetUserPaymentMethodsByUserUuidPaginated(uuidUser string, pag pagination.PaginationParams) (*pagination.PaginationResult, error) {
	resPaymentMethods, err := repository.GetUserPaymentMethodsByUserUuidPaginated(uuidUser, pag)
	if err != nil {
		return nil, err
	}

	var res []response_object.UserPaymentMethodRo
	for _, rPayment := range resPaymentMethods.Data.([]models.UserPaymentMethod) {
		modelDto := response_object.MapUserPaymentMethod(&rPayment)

		if modelDto.Country != nil && *modelDto.Country != "" {
			country, err := repository.GetCountry(utils.ParseIDToUUID(*modelDto.Country))
			if err != nil {
				return nil, err
			}
			modelDto.Country = country.Name
		}

		if modelDto.BankCountry != nil && *modelDto.BankCountry != "" {
			bankCountry, err := repository.GetCountry(utils.ParseIDToUUID(*modelDto.BankCountry))
			if err != nil {
				return nil, err
			}

			modelDto.BankCountry = bankCountry.Name
		}

		res = append(res, *modelDto)
	}

	resPaymentMethods.Data = res

	return resPaymentMethods, err
}

func GetUserPaymentMethodById(uuid uuid.UUID) (*response_object.UserPaymentMethodRo, error) {
	res, err := repository.GetUserPaymentMethodsById(uuid)
	if err != nil {
		return nil, err
	}

	modelDto := response_object.MapUserPaymentMethod(&res)

	if modelDto.Country != nil && *modelDto.Country != "" {
		country, err := repository.GetCountry(utils.ParseIDToUUID(*modelDto.Country))
		if err != nil {
			return nil, err
		}
		modelDto.Country = country.Name
	}

	if modelDto.BankCountry != nil && *modelDto.BankCountry != "" {
		bankCountry, err := repository.GetCountry(utils.ParseIDToUUID(*modelDto.BankCountry))
		if err != nil {
			return nil, err
		}

		modelDto.BankCountry = bankCountry.Name
	}

	return modelDto, err
}

func GetUserPaymentMethodByIdUnscoped(uuid uuid.UUID) (*response_object.UserPaymentMethodRo, error) {
	res, err := repository.GetUserPaymentMethodsByIdUnscoped(uuid)
	if err != nil {
		return nil, err
	}

	modelDto := response_object.MapUserPaymentMethod(&res)

	if modelDto.Country != nil && *modelDto.Country != "" {
		country, err := repository.GetCountry(utils.ParseIDToUUID(*modelDto.Country))
		if err != nil {
			return nil, err
		}
		modelDto.Country = country.Name
	}

	if modelDto.BankCountry != nil && *modelDto.BankCountry != "" {
		bankCountry, err := repository.GetCountry(utils.ParseIDToUUID(*modelDto.BankCountry))
		if err != nil {
			return nil, err
		}

		modelDto.BankCountry = bankCountry.Name
	}

	return modelDto, err
}

func DeleteUserPaymentMethod(uuid uuid.UUID, uuidUser string) error {
	err := repository.DeleteUserPaymentMethod(uuid, uuidUser)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUserPaymentMethodState(id *uuid.UUID, s string) error {
	err := repository.UpdateUserPaymentMethodState(id, s)
	if err != nil {
		return err
	}
	return nil
}

func ReplaceIbanStatementFile(uuidReplacement string, intervalInDays int) ([]models.AffectedRecord, error) {
	res, err := repository.UpdateUserPaymentFileUuid(uuidReplacement, intervalInDays)

	if err != nil {
		return nil, err
	}

	return res, nil
}
