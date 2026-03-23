package shopify_models

// Get publications
type PublicationResponse struct {
	Publications Publications `json:"publications"`
}
type Publications struct {
	Nodes []PublicationNode `json:"nodes"`
}

type PublicationNode struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Catalog interface{} `json:"catalog"` // Use interface{} because catalog is null
}

// PublishCollectionToPublication
type PublishablePublishResponse struct {
	PublishablePublish PublishablePublish `json:"publishablePublish"`
}

type PublishablePublish struct {
	Publishable Publishable `json:"publishable"`
	UserErrors  []string    `json:"userErrors"`
}

type Publishable struct {
	PublishedOnPublication bool `json:"publishedOnPublication"`
}
