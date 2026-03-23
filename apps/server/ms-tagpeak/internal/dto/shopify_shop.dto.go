package dto

import "github.com/google/uuid"

type ShopifyStoreExistsDto struct {
	Shop string `json:"shop"`
}

type GetInstallShopifyDto struct {
	Shop        string `json:"shop"`
	AccessToken string `json:"accessToken"`
}

type CreateShopDTO struct {
	Shop      string  `json:"shop" validate:"required"`
	Email     string  `json:"email" validate:"required,email"`
	Password  string  `json:"password" validate:"required,min=8"`
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	Country   *string `json:"country"`
	Currency  string  `json:"currency"`
}

type CreateShopifyStoreDTO struct {
	Name         string   `json:"name" validate:"required"`
	Percentage   float64  `json:"percentage" validate:"required"`
	ReturnPeriod int      `json:"returnPeriod" validate:"required"`
	Url          string   `json:"url" validate:"required"`
	ShopUuid     string   `json:"shopUuid" validate:"required"`
	Description  *string  `json:"description"`
	Countries    []string `json:"countries" validate:"required"`
	Categories   []string `json:"categories" validate:"required"`
}

type UpdateShopifyShopDTO struct {
	StoreUuid *uuid.UUID `json:"storeUuid"`
}
