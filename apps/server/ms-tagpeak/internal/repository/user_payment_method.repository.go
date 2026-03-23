package service

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
)

func CreateUserPaymentMethod(model models.UserPaymentMethod) (models.UserPaymentMethod, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.UserPaymentMethod{}, err
	}

	return model, nil
}

func GetUserPaymentMethodsByUserUuid(uuidUser string) ([]models.UserPaymentMethod, error) {
	dbInstance := db.GetDB()
	var model []models.UserPaymentMethod

	// Try fetching records with state = 'VALIDATED'
	err := dbInstance.Preload("PaymentMethod").Preload("File").
		Where("\"user\" = ?", uuidUser).
		Where("state = 'VALIDATED'").
		Order("created_at DESC").
		Find(&model).
		Error

	if err != nil {
		return []models.UserPaymentMethod{}, err
	}

	// If no results, try fetching with state = 'PENDING'
	if len(model) == 0 {
		err = dbInstance.Preload("PaymentMethod").Preload("File").
			Where("\"user\" = ?", uuidUser).
			Where("state = 'PENDING'").
			Order("created_at DESC").
			Find(&model).
			Error

		if err != nil {
			return []models.UserPaymentMethod{}, err
		}
	}

	return model, nil
}

func GetUserPaymentMethodsByUserUuidPaginated(uuidUser string,
	pageDTO pagination.PaginationParams) (*pagination.PaginationResult, error) {

	dbInstance := db.GetDB()
	var model []models.UserPaymentMethod
	var res pagination.PaginationResult

	res.Limit = pageDTO.Limit
	res.Page = pageDTO.Page
	res.Sort = pageDTO.Sort

	err := dbInstance.
		Scopes(pagination.Paginate(&model, &res, nil, dbInstance)).
		Preload("PaymentMethod").
		Preload("File").
		Where("\"user\" = ?", uuidUser).
		Find(&model).
		Error

	if err != nil {
		return &pagination.PaginationResult{}, err
	}
	res.Data = model

	return &res, nil
}

func GetUserPaymentMethodsById(uuid uuid.UUID) (models.UserPaymentMethod, error) {
	dbInstance := db.GetDB()
	var model models.UserPaymentMethod

	err := dbInstance.Preload("PaymentMethod").Preload("File").Where("uuid = ?", uuid).First(&model).Error

	if err != nil {
		return models.UserPaymentMethod{}, err
	}

	return model, nil
}

func GetUserPaymentMethodsByIdUnscoped(uuid uuid.UUID) (models.UserPaymentMethod, error) {
	dbInstance := db.GetDB()
	var model models.UserPaymentMethod

	err := dbInstance.Unscoped().
		Preload("PaymentMethod").
		Preload("File").
		Where("uuid = ?", uuid).
		First(&model).Error

	if err != nil {
		return models.UserPaymentMethod{}, err
	}

	return model, nil
}

func DeleteUserPaymentMethod(uuid uuid.UUID, uuidUser string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.UserPaymentMethod{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": uuidUser,
		}).Error

	err = dbInstance.Delete(&models.UserPaymentMethod{Uuid: uuid}).Error

	if err != nil {
		return err
	}

	return nil

}

func UpdateUserPaymentMethodState(id *uuid.UUID, s string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.UserPaymentMethod{}).
		Where("uuid = ?", id).
		Updates(map[string]interface{}{
			"state": s,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateUserPaymentFileUuid(uuidReplacement string, intervalInDays int) ([]models.AffectedRecord, error) {
	dbInstance := db.GetDB()

	result, err := clearOldFileUUIDs(dbInstance, intervalInDays, &uuidReplacement)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func clearOldFileUUIDs(db *gorm.DB, daysInterval int, replacementUUID *string) ([]models.AffectedRecord, error) {
	var results []models.AffectedRecord

	query := "SELECT * FROM clear_old_file_uuids(?, ?::uuid, 'job')"
	err := db.Raw(query, daysInterval, replacementUUID).Scan(&results).Error
	if err != nil {
		return nil, fmt.Errorf("failed to execute function: %w", err)
	}

	return results, nil
}
