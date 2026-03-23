package service

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/utils"
)

func GetPaymentMethod(uuid uuid.UUID) (*models.PaymentMethod, error) {
	res, err := repository.GetPaymentMethod(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.CustomErrorStruct{}.NotFoundError("Payment Method", uuid)
		}
		return nil, err
	}
	return res, nil
}

func GetAllPaymentMethods() (*[]models.PaymentMethod, error) {
	res, err := repository.GetAllPaymentMethods()
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func CreatePaymentMethod(dtoParam dto.CreatePaymentMethodDTO, uuidUser string) (*models.PaymentMethod, error) {

	model := utils.PaymentMethodDtoToModel(&dtoParam)
	model.CreatedBy = uuidUser

	res, err := repository.CreatePaymentMethod(model)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func UpdatePaymentMethod(dtoParam dto.UpdatePaymentMethodDTO, uuid uuid.UUID, uuidUser string) (*models.PaymentMethod, error) {

	toUpdate, err := repository.GetPaymentMethod(uuid)
	if err != nil {
		return nil, err
	}

	if dtoParam.Name != nil {
		toUpdate.Name = dtoParam.Name
	}
	if dtoParam.Code != nil {
		toUpdate.Code = dtoParam.Code
	}

	toUpdate.UpdatedBy = &uuidUser

	res, err := repository.UpdatePaymentMethod(*toUpdate)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func DeletePaymentMethod(uuid uuid.UUID, uuidUser string) error {
	err := repository.DeletePaymentMethod(uuid, uuidUser)
	if err != nil {
		return err
	}
	return nil
}
