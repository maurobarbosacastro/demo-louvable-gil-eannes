package service

import (
	"github.com/google/uuid"
	"mime/multipart"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/utils"
	"path/filepath"
)

func SaveFile(fileHeader *multipart.FileHeader, uuidUser string) (*models.File, error) {
	model := utils.MapModelDTO(fileHeader, filepath.Ext(fileHeader.Filename))
	model.CreatedBy = uuidUser

	res, err := repository.SaveFile(model)
	if err != nil {
		return &models.File{}, err
	}

	return &res, nil
}

func GetFileByUuid(uuid uuid.UUID) (*models.File, error) {
	res, err := repository.GetFileByUuid(uuid)
	if err != nil {
		return nil, err
	}

	return res, nil
}
