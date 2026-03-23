package repository

import (
	"ms-firebase-go/internal/db"
	"ms-firebase-go/internal/dto"
	"ms-firebase-go/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateNotification(notification models.Notification) (*models.Notification, error) {
	dbInstance := db.GetDB()

	if err := dbInstance.Create(&notification).Error; err != nil {
		return nil, err
	}
	return &notification, nil
}

func UpdateNotification(id uuid.UUID, updateFields dto.NotificationUpdateFields) (*models.Notification, error) {
	dbInstance := db.GetDB()

	var notification models.Notification
	if err := dbInstance.Where("uuid = ?", id).First(&notification).Error; err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})

	if updateFields.Title != nil {
		updates["title"] = *updateFields.Title
	}
	if updateFields.Content != nil {
		updates["body"] = *updateFields.Content
	}
	if updateFields.Date != nil {
		updates["date"] = *updateFields.Date
	}
	if updateFields.State != nil {
		updates["state"] = *updateFields.State
	}

	if len(updates) > 0 {
		if err := dbInstance.Model(&notification).Updates(updates).Error; err != nil {
			return nil, err
		}
	}

	if err := dbInstance.Preload("Target").Where("uuid = ?", id).First(&notification).Error; err != nil {
		return nil, err
	}

	return &notification, nil
}

func GetNotificationByID(id uuid.UUID) (*models.Notification, error) {
	dbInstance := db.GetDB()

	var notification models.Notification
	err := dbInstance.Preload("Target").Preload("Topics").Where("uuid = ?", id).First(&notification).Error
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

type NotificationFilters struct {
	Target         []uuid.UUID
	State          []models.NotificationState
	Date           *time.Time
	UserTokenState []models.JoinTableNotificationState
	TopicState     []models.JoinTableNotificationState
}

func GetAllNotifications(filters *NotificationFilters) ([]models.Notification, error) {
	dbInstance := db.GetDB()
	var notifications []models.Notification

	query := dbInstance.Model(&models.Notification{}).Preload("Target").Preload("Topics")

	if filters != nil {
		if len(filters.Target) > 0 {
			query = query.Joins("JOIN notification_user_tokens ON notification_user_tokens.notification_uuid = notification.uuid").
				Where("notification_user_tokens.user_token_uuid IN ?", filters.Target)
		}

		if len(filters.State) > 0 {
			query = query.Where("notification.state IN ?", filters.State)
		}

		if filters.Date != nil {
			query = query.Where("notification.date >= ?", filters.Date)
		}

		if len(filters.UserTokenState) > 0 {
			query = query.Joins("JOIN notification_user_tokens nut ON nut.notification_uuid = notification.uuid").
				Where("nut.state IN ?", filters.UserTokenState)
		}

		if len(filters.TopicState) > 0 {
			query = query.Joins("JOIN notification_topics nt ON nt.notification_uuid = notification.uuid").
				Where("nt.state IN ?", filters.TopicState)
		}
	}

	err := query.Distinct().Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func ReplaceNotificationAssociations(notificationID uuid.UUID, targets *[]models.UserToken, topics *[]models.Topic) error {
	dbInstance := db.GetDB()

	var notification models.Notification
	if err := dbInstance.Where("uuid = ?", notificationID).First(&notification).Error; err != nil {
		return err
	}

	if targets != nil {
		// Replace target associations
		if err := dbInstance.Model(&notification).Association("Target").Replace(targets); err != nil {
			return err
		}
	}

	if topics != nil {
		// Replace topic associations
		if err := dbInstance.Model(&notification).Association("Topics").Replace(topics); err != nil {
			return err
		}
	}

	return nil
}

// UpdateNotificationUserTokenState updates the state of a specific notification-user_token relationship
func UpdateNotificationUserTokenState(notificationUuid uuid.UUID, userTokenID uuid.UUID, state models.JoinTableNotificationState) error {
	dbInstance := db.GetDB()

	result := dbInstance.Model(&models.NotificationUserTokens{}).
		Where("notification_uuid = ? AND user_token_uuid = ?", notificationUuid, userTokenID).
		Update("state", state)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateNotificationTopicState updates the state of a specific notification-topic relationship
func UpdateNotificationTopicState(notificationUuid uuid.UUID, topicID uuid.UUID, state models.JoinTableNotificationState) error {
	dbInstance := db.GetDB()

	result := dbInstance.Model(&models.NotificationTopics{}).
		Where("notification_uuid = ? AND topic_uuid = ?", notificationUuid, topicID).
		Update("state", state)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetNotificationUserTokenStates gets all user token relationships for a notification with their states
func GetNotificationUserTokenStates(notificationID uuid.UUID) ([]models.NotificationUserTokens, error) {
	dbInstance := db.GetDB()

	var relations []models.NotificationUserTokens
	err := dbInstance.Where("notification_uuid = ?", notificationID).Find(&relations).Error

	return relations, err
}

// GetNotificationTopicStates gets all topic relationships for a notification with their states
func GetNotificationTopicStates(notificationID uuid.UUID) ([]models.NotificationTopics, error) {
	dbInstance := db.GetDB()

	var relations []models.NotificationTopics
	err := dbInstance.Where("notification_uuid = ?", notificationID).Find(&relations).Error

	return relations, err
}

func GetNotificationByDate(date time.Time) (*[]models.Notification, error) {
	dbInstance := db.GetDB()

	// Make sure the date is parsed as we want. If it fails, we use the provided one and hope it's right
	dateStr := date.Format(time.RFC3339)

	var notifications []models.Notification
	if err := dbInstance.Model(&models.Notification{}).
		Preload("Target").
		Preload("Topics").
		Where("date = ?", dateStr).
		Find(&notifications).Error; err != nil {
		return nil, err
	}
	return &notifications, nil
}
