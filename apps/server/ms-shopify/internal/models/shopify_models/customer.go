package shopify_models

type CustomerResponse struct {
	Customer CustomerSubResponse `json:"customer"`
}

type CustomerSubResponse struct {
	Email string `json:"email"`
}
