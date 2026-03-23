package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
	"time"
)

func CreateReferralClick(referralClick models.ReferralClicks) (models.ReferralClicks, error) {
	dbInstance := db.GetDB()

	err := dbInstance.Model(models.ReferralClicks{}).Create(&referralClick).Error
	if err != nil {
		return models.ReferralClicks{}, err
	}

	return referralClick, nil
}

func CountReferralClickByCode(code string) (int64, error) {
	dbInstance := db.GetDB()

	var count int64

	err := dbInstance.Model(models.ReferralClicks{}).
		Where("code = ?", code).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetClicksByMonth(code string, startDate time.Time, endDate time.Time) (float64, error) {
	dbInstance := db.GetDB()

	var count int64

	err := dbInstance.Model(models.ReferralClicks{}).
		Where("code = ?", code).
		Where("created_at >= ?", startDate).
		Where("created_at <= ?", endDate).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return float64(count), nil
}

func UpdateReferralClickReferral(referralClick uuid.UUID, referralUUID uuid.UUID) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(models.ReferralClicks{}).
		Where("uuid = ?", referralClick).
		Updates(map[string]interface{}{
			"referral_uuid": referralUUID,
		}).Error

	if err != nil {
		return err
	}

	return nil
}
