package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"strings"
)

func GetPartner(uuid uuid.UUID) (models.Partner, error) {
	res, err := repository.GetPartner(uuid)
	if err != nil {
		return models.Partner{}, err
	}
	return res, nil
}

func GetAllPartners(pag pagination.PaginationParams, filters dto.PartnerFiltersDTO) (*pagination.PaginationResult, error) {

	res, err := repository.GetAllPartnersWithPagination(pag, filters)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreatePartner(dtoParam dto.CreatePartnerDTO, uuidUser string) (models.Partner, error) {

	if dtoParam.Code == nil {
		partnerCode := strings.ToLower(dtoParam.Name)
		partnerCode = strings.ReplaceAll(partnerCode, " ", "_")
		dtoParam.Code = &partnerCode
	}

	model := utils.CreatePartnerDtoToModel(&dtoParam)
	model.CreatedBy = uuidUser

	res, err := repository.CreatePartner(model)
	if err != nil {
		return models.Partner{}, err
	}
	return res, nil
}

func UpdatePartner(dtoParam dto.UpdatePartnerDTO, uuid uuid.UUID, uuidUser string) (models.Partner, error) {

	toUpdate, err := repository.GetPartner(uuid)
	if err != nil {
		return models.Partner{}, err
	}

	if dtoParam.Name != nil {
		toUpdate.Name = dtoParam.Name
	}
	if dtoParam.ECommercePlatform != nil {
		toUpdate.ECommercePlatform = dtoParam.ECommercePlatform
	}
	if dtoParam.ValidationPeriod != nil {
		toUpdate.ValidationPeriod = dtoParam.ValidationPeriod
	}
	if dtoParam.DeepLink != nil {
		toUpdate.DeepLink = dtoParam.DeepLink
	}
	if dtoParam.DeepLinkIdentifier != nil {
		toUpdate.DeepLinkIdentifier = dtoParam.DeepLinkIdentifier
	}
	if dtoParam.SubIdentifier != nil {
		toUpdate.SubIdentifier = dtoParam.SubIdentifier
	}
	if dtoParam.PercentageTagpeak != nil {
		toUpdate.PercentageTagpeak = dtoParam.PercentageTagpeak
	}
	if dtoParam.PercentageInvested != nil {
		toUpdate.PercentageInvested = dtoParam.PercentageInvested
	}

	toUpdate.UpdatedBy = &uuidUser

	res, err := repository.UpdatePartner(toUpdate)
	if err != nil {
		return models.Partner{}, err
	}
	return res, nil
}

func DeletePartner(uuid uuid.UUID, uuidUser string) error {

	err := repository.DeletePartner(uuid, uuidUser)
	if err != nil {
		return err
	}

	return nil
}

func CodeAlreadyExists(code string) (bool, error) {
	exists, err := repository.CodeAlreadyExists(code)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func GetPartnerByCode(code *string) (models.Partner, error) {
	res, err := repository.GetPartnerByCode(code)
	if err != nil {
		return models.Partner{}, err
	}
	return res, nil
}
