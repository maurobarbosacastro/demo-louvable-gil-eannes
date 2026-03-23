package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"time"
)

func GetReferral(uuid uuid.UUID) (models.Referral, error) {
	dbInstance := db.GetDB()

	var referral models.Referral

	err := dbInstance.
		Where("referrer_uuid = ?", uuid).
		Find(&referral).Error

	if err != nil {
		return models.Referral{}, err
	}

	return referral, nil
}

func GetReferralByInvitee(uuid uuid.UUID) (*models.Referral, error) {
	dbInstance := db.GetDB()

	var referral *models.Referral

	err := dbInstance.
		Where("invitee_uuid = ?", uuid).
		Find(&referral).Error

	if err != nil {
		return nil, err
	}

	return referral, nil
}

func GetAllReferralByUserUuid(uuid uuid.UUID) ([]models.Referral, error) {
	dbInstance := db.GetDB()

	var referral []models.Referral

	err := dbInstance.
		Where("referrer_uuid = ?", uuid).
		Find(&referral).Error

	if err != nil {
		return []models.Referral{}, err
	}

	return referral, nil
}

func GetAllReferralByUserUuidWithPagination(pagDTO pagination.PaginationParams, uuid uuid.UUID) (*pagination.PaginationResult, error) {
	dbInstance := db.GetDB()

	var referral []models.Referral
	var res pagination.PaginationResult

	// Set pagination details
	res.Limit = pagDTO.Limit
	res.Page = pagDTO.Page
	res.Sort = pagDTO.Sort

	query := dbInstance.Model(&models.Referral{}).Where("referrer_uuid = ?", uuid)

	// Apply pagination and execute the query
	err := query.Scopes(pagination.Paginate(&referral, &res, nil, query)).
		Find(&referral).Error

	if err != nil {
		return &pagination.PaginationResult{}, err
	}

	// Transform the results into DTOs
	var modelDto []dto.ReferralDTO
	for _, refer := range referral {
		modelDto = append(modelDto, utils.BuildReferralDTO(refer))
	}
	res.Data = modelDto

	return &res, nil
}

func CreateReferral(referral models.Referral) (models.Referral, error) {
	dbInstance := db.GetDB()

	err := dbInstance.Create(&referral).Error

	if err != nil {
		return models.Referral{}, err
	}

	return referral, nil
}

func UpdateReferralFirstTransaction(referral models.Referral) (models.Referral, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Model(models.Referral{}).Where("uuid = ?", referral.Uuid).Update("successful_first_transaction", referral.SuccessfulFirstTransaction).Error
	if err != nil {
		return models.Referral{}, err
	}

	return referral, nil
}

func CountReferralByUserRegistered(referrerUUID uuid.UUID) (int64, error) {
	dbInstance := db.GetDB()

	var count int64

	err := dbInstance.Model(models.Referral{}).Where("referrer_uuid = ?", referrerUUID).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, err
}

func CountReferralByFirstTransaction(referrerUUID uuid.UUID, firstTransaction bool) (int64, error) {
	dbInstance := db.GetDB()

	var count int64

	err := dbInstance.Model(models.Referral{}).Where("referrer_uuid = ?", referrerUUID).
		Where("successful_first_transaction = ?", firstTransaction).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, err
}

func GetRegisteredByMonth(referrerUUID uuid.UUID, startDate time.Time, endDate time.Time) (float64, error) {
	dbInstance := db.GetDB()

	var count int64

	err := dbInstance.Model(models.Referral{}).
		Where("referrer_uuid = ?", referrerUUID).
		Where("created_at >= ?", startDate).
		Where("created_at <= ?", endDate).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return float64(count), nil
}

func GetFirstPurchaseByMonth(referrerUUID uuid.UUID, startDate time.Time, endDate time.Time, firstPurchase bool) (float64, error) {
	dbInstance := db.GetDB()

	var count int64

	err := dbInstance.Model(models.Referral{}).
		Where("referrer_uuid = ?", referrerUUID).
		Where("successful_first_transaction = ?", firstPurchase).
		Where("created_at >= ?", startDate).
		Where("created_at <= ?", endDate).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return float64(count), nil
}
