package service

import (
	"ms-changeme/internal/db"
	"ms-changeme/internal/models"
)

func GetExample() models.Example {
	dbInstance := db.GetDB()

	var example models.Example
	dbInstance.First(&example)

	return example
}

func GetExamples() []models.Example {
	dbInstance := db.GetDB()
	var examples []models.Example
	dbInstance.Find(&examples)

	return examples
}
