package service

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
)

func GetLanguage(uuid uuid.UUID) (*models.Language, error) {
	res, err := repository.GetLanguage(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.CustomErrorStruct{}.NotFoundError("Payment Method", uuid)
		}
		return nil, err
	}
	return &res, nil
}

func GetAllLanguages(pag pagination.PaginationParams) (*pagination.PaginationResult, error) {
	res, err := repository.GetAllLanguages(pag)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateLanguage(dtoParam dto.CreateLanguageDTO, uuidUser string) (*models.Language, error) {
	model := utils.LanguageDtoToModel(&dtoParam)
	model.CreatedBy = uuidUser

	res, err := repository.CreateLanguage(model)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func UpdateLanguage(dtoParam dto.UpdateLanguageDTO, uuid uuid.UUID, uuidUser string) (*models.Language, error) {

	toUpdate, err := repository.GetLanguage(uuid)
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

	res, err := repository.UpdateLanguage(toUpdate)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func DeleteLanguage(uuid uuid.UUID, uuidUser string) error {
	err := repository.DeleteLanguage(uuid, uuidUser)
	if err != nil {
		return err
	}
	return nil
}

func LanguageCodeExist(code string) (bool, error) {
	exists, err := repository.LanguageCodeExist(code)
	if err != nil {
		return false, err
	}
	return exists, nil
}
