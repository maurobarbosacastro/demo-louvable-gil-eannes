package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
)

func CreateShopifyShop(model models.ShopifyShop) (models.ShopifyShop, error) {
	dbInstance := db.GetDB()

	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.ShopifyShop{}, err
	}

	return model, nil
}

func GetShopifyShopByUuid(uuid uuid.UUID) (*models.ShopifyShop, error) {
	dbInstance := db.GetDB()
	var model models.ShopifyShop

	err := dbInstance.Where("shop_uuid = ?", uuid).First(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func GetShopifyShopByUserUuid(uuid uuid.UUID) (*models.ShopifyShop, error) {
	dbInstance := db.GetDB()
	var model models.ShopifyShop

	err := dbInstance.Where("user_uuid = ?", uuid).First(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func GetShopifyShopByStoreUuid(uuid uuid.UUID) (*models.ShopifyShop, error) {
	dbInstance := db.GetDB()
	var model models.ShopifyShop

	err := dbInstance.Where("store_uuid = ?", uuid).First(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func SetStoreToShopifyShop(uuid uuid.UUID, storeUuid uuid.UUID) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.ShopifyShop{}).
		Where("shop_uuid = ?", uuid).
		Update("store_uuid", storeUuid).Error
	if err != nil {
		return err
	}

	return nil
}

func GetShopifyStoreStats(uuid string) (models.ShopStats, error) {
	dbInstance := db.GetDB()
	var model models.ShopStats

	sql := `SELECT * FROM shopify_order_summary where shop_uuid = ?`
	err := dbInstance.Raw(sql, uuid).Scan(&model).Error
	if err != nil {
		return models.ShopStats{}, err
	}
	return model, nil
}
