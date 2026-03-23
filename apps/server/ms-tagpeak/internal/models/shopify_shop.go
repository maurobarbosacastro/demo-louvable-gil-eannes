package models

import "github.com/google/uuid"

type ShopifyShop struct {
	Uuid      uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	ShopUuid  uuid.UUID  `gorm:"type:uuid;" json:"shopUuid"`
	UserUuid  uuid.UUID  `gorm:"type:uuid;" json:"userUuid"`
	StoreUuid *uuid.UUID `gorm:"type:uuid;" json:"storeUuid"`
}

type ShopStats struct {
	TotalOrders int64   `json:"totalOrders"`
	TotalAmount float64 `json:"totalAmount"`
	Currency    string  `json:"currency"`
}
