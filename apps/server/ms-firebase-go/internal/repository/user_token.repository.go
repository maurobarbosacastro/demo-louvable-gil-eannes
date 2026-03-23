package repository

import (
	"ms-firebase-go/internal/db"
	"ms-firebase-go/internal/models"

	"github.com/google/uuid"
)

func CreateUserToken(userToken models.UserToken) (*models.UserToken, error) {
	dbInstance := db.GetDB()

	if err := dbInstance.Create(&userToken).Error; err != nil {
		return nil, err
	}
	return &userToken, nil
}

func UpdateUserToken(userToken models.UserToken) (*models.UserToken, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&userToken).Error
	if err != nil {
		return nil, err
	}

	return &userToken, nil
}

func GetUserTokenByToken(token string) (*models.UserToken, error) {
	dbInstance := db.GetDB()

	var userToken models.UserToken
	err := dbInstance.Where("token = ?", token).First(&userToken).Error
	if err != nil {
		return nil, err
	}

	return &userToken, nil
}

func GetUserTokenByUserUuid(userUuid string) (*models.UserToken, error) {
	dbInstance := db.GetDB()

	var userToken models.UserToken
	err := dbInstance.Where("user_uuid = ?", userUuid).First(&userToken).Error
	if err != nil {
		return nil, err
	}

	return &userToken, nil
}

func GetAllUserTokensByUser(userUUID string) (*[]models.UserToken, error) {
	dbInstance := db.GetDB()
	var userTokens []models.UserToken

	err := dbInstance.Where("user_uuid = ?", userUUID).Find(&userTokens).Error
	if err != nil {
		return nil, err
	}
	return &userTokens, nil
}

func GetAllUserTokens() ([]models.UserToken, error) {
	dbInstance := db.GetDB()
	var userTokens []models.UserToken

	err := dbInstance.Find(&userTokens).Error
	if err != nil {
		return nil, err
	}
	return userTokens, nil
}

func GetUserTokensByUserUUIDs(uuids []string) ([]models.UserToken, error) {
	dbInstance := db.GetDB()
	var userTokens []models.UserToken

	err := dbInstance.Where("user_uuid IN ?", uuids).Find(&userTokens).Error
	if err != nil {
		return nil, err
	}

	return userTokens, nil
}

func GetUserTokensByTokens(tokens []string) ([]models.UserToken, error) {
	dbInstance := db.GetDB()
	var userTokens []models.UserToken

	err := dbInstance.Where("token IN ?", tokens).Find(&userTokens).Error
	if err != nil {
		return nil, err
	}

	return userTokens, nil
}

func DeleteUserToken(id uuid.UUID) error {
	dbInstance := db.GetDB()

	err := dbInstance.Delete(&models.UserToken{}, "uuid = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func RemoveUserTokenFromAllTopics(userToken models.UserToken) (*[]models.Topic, error) {
	dbInstance := db.GetDB()

	var topics []models.Topic
	err := dbInstance.Model(&userToken).Association("Topics").Find(&topics)
	if err != nil {
		return nil, err
	}

	if len(topics) > 0 {
		err = dbInstance.Model(&userToken).Association("Topics").Delete(&topics)
		if err != nil {
			return nil, err
		}
	}

	return &topics, nil
}
