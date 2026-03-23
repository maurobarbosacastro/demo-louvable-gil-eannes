package shopify_dtos

// CodeDiscountNodeByCode
type CodeDiscountNodeByCode struct {
	Code string `json:"code"`
}

// DiscountCodeBxgyCreate
type BasicCodeBxgyDiscount struct {
	BasicCodeDiscount BasicCodeDiscount `json:"basicCodeDiscount"`
}

type BasicCodeDiscount struct {
	Title             string            `json:"title"`
	Code              string            `json:"code"`
	StartsAt          string            `json:"startsAt"`
	CustomerSelection CustomerSelection `json:"customerSelection"`
	CustomerGets      CustomerGets      `json:"customerGets"`
}

type CustomerSelection struct {
	All bool `json:"all"`
}

type CustomerGets struct {
	Value Value `json:"value"`
	Items Items `json:"items"`
}

type Value struct {
	Percentage float64 `json:"percentage"`
}

type Items struct {
	Collections Collections `json:"collections"`
}

type Collections struct {
	Add string `json:"add"`
}
