package dto

import (
	"github.com/google/uuid"
)

type StoreCategoryDTO struct {
	Categories []string `json:"categories"`
}

type StoreUploadResponseDTO struct {
	Message string      `json:"message"`
	Stores  []uuid.UUID `json:"stores_uuid"`
}
