package utils

import (
	"mime/multipart"
	"ms-tagpeak/internal/models"
)

func MapModelDTO(file *multipart.FileHeader, extension string) models.File {
	return models.File{
		Name:      &file.Filename,
		Extension: &extension,
	}
}
