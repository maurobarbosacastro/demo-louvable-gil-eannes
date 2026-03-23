package repository

import (
	"ms-interactive-brokers/internal/models"
	"ms-interactive-brokers/internal/service"
)

func GetExample() models.Example {
	return service.GetExample()
}
