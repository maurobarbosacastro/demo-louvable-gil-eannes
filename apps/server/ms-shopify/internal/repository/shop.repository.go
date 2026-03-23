package repository

import (
	"github.com/google/uuid"
	"ms-shopify/internal/db"
	"ms-shopify/internal/models"
)

func GetShopByUrl(url string) (*models.Shop, error) {
	dbInstance := db.GetDB()
	var shop models.Shop

	err := dbInstance.Where("url = ?", url).First(&shop).Error

	if err != nil {
		return nil, err
	}

	return &shop, nil
}

func CreateShop(body models.Shop) (*models.Shop, error) {
	dbInstance := db.GetDB()

	err := dbInstance.Create(&body).Error
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func GetByUuid(uuid uuid.UUID) (*models.Shop, error) {
	dbInstance := db.GetDB()
	var shop models.Shop

	err := dbInstance.Where("uuid = ?", uuid).First(&shop).Error
	if err != nil {
		return nil, err
	}

	return &shop, nil
}

func UpdateShop(shop models.Shop) (*models.Shop, error) {
	dbInstance := db.GetDB()

	err := dbInstance.Save(&shop).Error
	if err != nil {
		return nil, err
	}

	return &shop, nil
}
