package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
)

func GetPaymentMethod(uuid uuid.UUID) (*models.PaymentMethod, error) {
	dbInstance := db.GetDB()

	var model models.PaymentMethod
	// Execute the query and capture the error
	err := dbInstance.Where("uuid = ?", uuid).First(&model).Error

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func GetAllPaymentMethods() ([]models.PaymentMethod, error) {

	dbInstance := db.GetDB()
	var model []models.PaymentMethod

	err := dbInstance.Find(&model).Error
	if err != nil {
		return []models.PaymentMethod{}, err
	}

	return model, nil
}

func CreatePaymentMethod(model models.PaymentMethod) (models.PaymentMethod, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.PaymentMethod{}, err
	}

	return model, nil
}

func UpdatePaymentMethod(model models.PaymentMethod) (models.PaymentMethod, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&model).Error
	if err != nil {
		return models.PaymentMethod{}, err
	}

	return model, nil
}

func DeletePaymentMethod(uuid uuid.UUID, user string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.PaymentMethod{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.PaymentMethod{}, "uuid = ?", uuid).Error

	if err != nil {
		return err
	}

	return nil
}
