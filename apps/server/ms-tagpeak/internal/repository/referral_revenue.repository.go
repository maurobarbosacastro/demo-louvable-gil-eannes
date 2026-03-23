package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
	"time"
)

func CreateReferralRevenue(referralRevenue models.ReferralRevenueHistory) (models.ReferralRevenueHistory, error) {
	dbInstance := db.GetDB()

	err := dbInstance.Model(models.ReferralRevenueHistory{}).Create(&referralRevenue).Error
	if err != nil {
		return models.ReferralRevenueHistory{}, err
	}

	return referralRevenue, nil
}

func GetRevenueByMonth(referrerUUID uuid.UUID, startDate time.Time, endDate time.Time) (float64, error) {
	dbInstance := db.GetDB()

	var totalAmount float64

	subQuery := dbInstance.Model(&models.Referral{}).
		Select("uuid").
		Where("referrer_uuid = ?", referrerUUID)

	err := dbInstance.
		Model(&models.ReferralRevenueHistory{}).
		Where("referral_uuid in (?)", subQuery).
		Where("created_at >= ?", startDate).
		Where("created_at <= ?", endDate).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalAmount).Error

	if err != nil {
		return 0, err
	}

	return totalAmount, nil
}

func GetReferralsByReferrerUuid(referrerUUID uuid.UUID) ([]models.Referral, error) {
	dbInstance := db.GetDB()

	var referrals []models.Referral

	err := dbInstance.Model(&models.Referral{}).
		Where("referrer_uuid = ?", referrerUUID).
		Find(&referrals).Error

	if err != nil {
		return nil, err
	}

	return referrals, nil
}

func GetNumberOfSuccessfulReferralsByReferrerUuid(referrerUUID uuid.UUID) (int64, error) {
	dbInstance := db.GetDB()

	var count int64

	err := dbInstance.Model(&models.Referral{}).
		Where("referrer_uuid = ?", referrerUUID).
		Where("successful_first_transaction = ?", true).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetAmountByReferralUuid(referralUUID uuid.UUID) (float64, error) {
	dbInstance := db.GetDB()

	var amount float64

	err := dbInstance.Model(&models.ReferralRevenueHistory{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("referral_uuid = ?", referralUUID).
		Scan(&amount).Error

	if err != nil {
		return amount, err
	}

	return amount, nil
}

func GetAllReferralRevenueByReferrerUuid(referrerUUID uuid.UUID) (*float64, error) {
	dbInstance := db.GetDB()

	var amount float64

	subQuery := dbInstance.Model(&models.Referral{}).
		Select("uuid").
		Where("referrer_uuid = ?", referrerUUID)

	err := dbInstance.
		Model(&models.ReferralRevenueHistory{}).
		Where("referral_uuid in (?)", subQuery).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&amount).Error

	if err != nil {
		return nil, err
	}

	return &amount, nil
}
