package repository

import (
	"ms-changeme/internal/models"
	"ms-changeme/internal/service"
)

func GetExample() models.Example {
	return service.GetExample()
}
