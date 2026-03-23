package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
)

func SaveFile(file models.File) (models.File, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&file).Error

	if err != nil {
		return models.File{}, err
	}

	return file, nil
}

func GetFileByUuid(uuid uuid.UUID) (*models.File, error) {
	dbInstance := db.GetDB()

	var model models.File

	err := dbInstance.Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return &models.File{}, err
	}

	return &model, nil
}
