package dto

import (
	"ms-shopify/internal/models"
)

type ShopifyStoreExistsDto struct {
	Shop string `json:"shop"`
}

type CreateShopifyShopDTO struct {
	Shop string `json:"shop" validate:"required"`
	User string `json:"user" validate:"required"`
}

func (dto CreateShopifyShopDTO) ToModel() models.Shop {
	return models.Shop{
		Url: dto.Shop,
		BaseEntity: models.BaseEntity{
			CreatedBy: dto.User,
		},
	}
}

type UpdateShopifyShopDTO struct {
	AccessToken      *string `json:"accessToken"`
	InstallationDone *bool   `json:"installationDone"`
	State            *string `json:"state"`
}

type InstallSetupDto struct {
	Products *[]string `json:"products,omitempty"`
}
