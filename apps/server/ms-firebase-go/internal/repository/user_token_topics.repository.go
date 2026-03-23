package repository

import (
	"ms-firebase-go/internal/db"
	"ms-firebase-go/internal/models"

	"github.com/google/uuid"
)

func AddUserTokensToTopic(topicID uuid.UUID, userTokens []models.UserToken) error {
	dbInstance := db.GetDB()

	var topic models.Topic
	if err := dbInstance.Where("uuid = ?", topicID).First(&topic).Error; err != nil {
		return err
	}

	if err := dbInstance.Model(&topic).Association("Target").Append(userTokens); err != nil {
		return err
	}

	return nil
}

func RemoveUserTokensFromTopic(topicID uuid.UUID, userTokens []models.UserToken) error {
	dbInstance := db.GetDB()

	var topic models.Topic
	if err := dbInstance.Where("uuid = ?", topicID).First(&topic).Error; err != nil {
		return err
	}

	if err := dbInstance.Model(&topic).Association("Target").Delete(&userTokens); err != nil {
		return err
	}

	return nil
}
