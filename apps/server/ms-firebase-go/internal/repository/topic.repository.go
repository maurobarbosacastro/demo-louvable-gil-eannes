package repository

import (
	"ms-firebase-go/internal/db"
	"ms-firebase-go/internal/models"

	"github.com/google/uuid"
)

func CreateTopic(topic models.Topic) (*models.Topic, error) {
	dbInstance := db.GetDB()

	if err := dbInstance.Create(&topic).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

func GetTopicByUUID(id uuid.UUID) (*models.Topic, error) {
	dbInstance := db.GetDB()

	var topic models.Topic
	err := dbInstance.Preload("Target").Where("uuid = ?", id).First(&topic).Error
	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func GetTopicByName(name string) (*models.Topic, error) {
	dbInstance := db.GetDB()

	var topic models.Topic
	err := dbInstance.Preload("Target").Where("name = ?", name).First(&topic).Error
	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func GetAllTopics() ([]models.Topic, error) {
	dbInstance := db.GetDB()
	var topics []models.Topic

	err := dbInstance.Preload("Target").Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func GetTopicsByUUIDs(ids []string) ([]models.Topic, error) {
	dbInstance := db.GetDB()
	var topics []models.Topic

	err := dbInstance.Preload("Target").Where("uuid IN ?", ids).Find(&topics).Error
	if err != nil {
		return nil, err
	}

	return topics, nil
}
