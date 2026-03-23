package response_object

import "ms-shopify/internal/models"

type SetupRO struct {
	AlreadyExists  bool        `json:"alreadyExists"`
	DiscountCodeId string      `json:"discountCodeId"`
	Shop           models.Shop `json:"shop"`
}
