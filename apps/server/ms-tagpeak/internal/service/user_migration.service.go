package service

import (
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
)

func CreateUserMigration(model models.UserMigration) error {
	err := repository.CreateUserMigration(model)
	if err != nil {
		return err
	}
	return nil
}
