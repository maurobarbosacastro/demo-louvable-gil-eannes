package service

import (
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
)

func CreateUserMigration(model models.UserMigration) error {
	dbInstance := db.GetDB()

	err := dbInstance.Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}
