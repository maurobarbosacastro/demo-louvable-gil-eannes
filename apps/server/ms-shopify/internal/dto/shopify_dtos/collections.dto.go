package shopify_dtos

// Search for collection by handle
type CollectionsSearch struct {
	Title string `json:"title"`
}

type Identifier struct {
	CustomId CustomId `json:"customId"`
}

type CustomId struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Namespace string `json:"namespace"`
}

// Create collection
type CreateCollectionInput struct {
	Title string `json:"title"`
}

type Metafield struct {
	Namespace string `json:"namespace"`
	Key       string `json:"key"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

type CreateCollection struct {
	Input CreateCollectionInput `json:"input"`
}

// Add products to collection
type CollectionAddProducts struct {
	ID         string   `json:"id"`
	ProductIds []string `json:"productIds"`
}
