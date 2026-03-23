package shopify_models

// Search for collection
type CollectionsResponse struct {
	Collections Collections `json:"collections"`
}

type Collections struct {
	Nodes []CollectionNode `json:"nodes"`
}

type CollectionNode struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Create collection
type CollectionCreateResponse struct {
	CollectionCreate CollectionCreate `json:"collectionCreate"`
}

type CollectionCreate struct {
	Collection Collection            `json:"collection"`
	UserErrors []UserErrorCollection `json:"userErrors"`
}

type Collection struct {
	ID string `json:"id"`
}

type UserErrorCollection struct {
	Field     []string `json:"field"`
	Message   string   `json:"message"`
	ExtraInfo *string  `json:"extraInfo"`
	Code      *string  `json:"code"`
}

// Add products to collection
type CollectionAddProductsV2Response struct {
	CollectionAddProductsV2 CollectionAddProductsV2 `json:"collectionAddProductsV2"`
}

type CollectionAddProductsV2 struct {
	Job        *Job                   `json:"job"`
	UserErrors []UserErrorCollection  `json:"userErrors"`
}

type Job struct {
	ID string `json:"id"`
}
