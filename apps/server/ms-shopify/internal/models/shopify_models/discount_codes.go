package shopify_models

// CodeDiscountNodeByCode

type CodesCount struct {
	Count int `json:"count"`
}

type CodeDiscountBasic struct {
	TypeName     string     `json:"__typename"`
	CodesCount   CodesCount `json:"codesCount"`
	ShortSummary string     `json:"shortSummary"`
}

type CodeDiscountNodeByCode struct {
	CodeDiscount CodeDiscountBasic `json:"codeDiscount"`
	ID           string            `json:"id"`
}

type CodeDiscountNodeByCodeResponse struct {
	CodeDiscountNodeByCode CodeDiscountNodeByCode `json:"codeDiscountNodeByCode"`
}

// DiscountCodeBxgyCreate
type DiscountCodeBasicCreateResponse struct {
	DiscountCodeBasicCreate DiscountCodeBasicCreate `json:"discountCodeBasicCreate"`
}

type DiscountCodeBasicCreate struct {
	CodeDiscountNode CodeDiscountNode `json:"codeDiscountNode"`
	UserErrors       []UserError      `json:"userErrors"`
}

type CodeDiscountNode struct {
	ID           string       `json:"id"`
	CodeDiscount CodeDiscount `json:"codeDiscount"`
}

type CodeDiscount struct {
	Title             string            `json:"title"`
	StartsAt          string            `json:"startsAt"`
	CustomerSelection CustomerSelection `json:"customerSelection"`
	CustomerGets      CustomerGets      `json:"customerGets"`
}

type CustomerSelection struct {
	AllCustomers bool `json:"allCustomers"`
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

type UserError struct {
	Field     []string `json:"field"`
	Message   string   `json:"message"`
	ExtraInfo *string  `json:"extraInfo"`
	Code      *string  `json:"code"`
}
