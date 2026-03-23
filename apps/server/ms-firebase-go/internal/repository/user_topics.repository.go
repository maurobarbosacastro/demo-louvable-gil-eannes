package repository

import (
	"errors"
	"ms-firebase-go/internal/db"
	"ms-firebase-go/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func AddUserToTopic(userUUID string, topicUUID uuid.UUID, createdBy string) error {
	dbInstance := db.GetDB()

	userTopic := models.UserTopics{
		UserUUID:  userUUID,
		TopicUUID: topicUUID,
		BaseEntity: models.BaseEntity{
			CreatedBy: createdBy,
		},
	}

	if err := dbInstance.Create(&userTopic).Error; err != nil {
		var pge *pgconn.PgError
		if errors.As(err, &pge) && pge.Code == "23505" {
			// duplicate key -> handle gracefully
			return nil
		}
		return err
	}

	return nil
}

func AddUserToTopics(userUUID string, topicUUIDs []uuid.UUID, createdBy string) error {
	dbInstance := db.GetDB()

	userTopics := make([]models.UserTopics, len(topicUUIDs))
	for i, topicUUID := range topicUUIDs {
		userTopics[i] = models.UserTopics{
			UserUUID:  userUUID,
			TopicUUID: topicUUID,
			BaseEntity: models.BaseEntity{
				CreatedBy: createdBy,
			},
		}
	}

	if err := dbInstance.Create(&userTopics).Error; err != nil {
		return err
	}

	return nil
}

func RemoveUserFromTopic(userUUID string, topicUUID uuid.UUID) error {
	dbInstance := db.GetDB()

	err := dbInstance.Delete(&models.UserTopics{}, "user_uuid = ? AND topic_uuid = ?", userUUID, topicUUID).Error
	if err != nil {
		return err
	}

	return nil
}

func RemoveUserFromAllTopics(userUUID string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Delete(&models.UserTopics{}, "user_uuid = ?", userUUID).Error
	if err != nil {
		return err
	}

	return nil
}

func GetTopicsByUserUUID(userUUID string) ([]uuid.UUID, error) {
	dbInstance := db.GetDB()

	var userTopics []models.UserTopics
	err := dbInstance.Where("user_uuid = ?", userUUID).Find(&userTopics).Error
	if err != nil {
		return nil, err
	}

	topicUUIDs := make([]uuid.UUID, len(userTopics))
	for i, ut := range userTopics {
		topicUUIDs[i] = ut.TopicUUID
	}

	return topicUUIDs, nil
}

func GetUsersByTopicUUID(topicUUID uuid.UUID) ([]string, error) {
	dbInstance := db.GetDB()

	var userTopics []models.UserTopics
	err := dbInstance.Where("topic_uuid = ?", topicUUID).Find(&userTopics).Error
	if err != nil {
		return nil, err
	}

	userUUIDs := make([]string, len(userTopics))
	for i, ut := range userTopics {
		userUUIDs[i] = ut.UserUUID
	}

	return userUUIDs, nil
}

func CheckUserTopicExists(userUUID string, topicUUID uuid.UUID) (bool, error) {
	dbInstance := db.GetDB()

	var count int64
	err := dbInstance.Model(&models.UserTopics{}).Where("user_uuid = ? AND topic_uuid = ?", userUUID, topicUUID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
