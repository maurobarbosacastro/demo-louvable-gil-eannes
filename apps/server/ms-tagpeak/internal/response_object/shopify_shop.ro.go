package response_object

type ShopifyShopExistRO struct {
	Uuid   *string `json:"uuid"`
	Exists bool    `json:"exists"`
}
