package webhooks

import "time"

type ShopifyOrderDelete struct {
	ID int64 `json:"id"`
}

type ShopifyOrder struct {
	ID                                    int64                 `json:"id"`
	AdminGraphqlAPIID                     string                `json:"admin_graphql_api_id"`
	AppID                                 int                   `json:"app_id"`
	BrowserIP                             string                `json:"browser_ip"`
	BuyerAcceptsMarketing                 bool                  `json:"buyer_accepts_marketing"`
	CancelReason                          *string               `json:"cancel_reason"`
	CancelledAt                           *string               `json:"cancelled_at"`
	CartToken                             string                `json:"cart_token"`
	CheckoutID                            int64                 `json:"checkout_id"`
	CheckoutToken                         string                `json:"checkout_token"`
	ClientDetails                         ClientDetails         `json:"client_details"`
	ClosedAt                              *string               `json:"closed_at"`
	Company                               *string               `json:"company"`
	ConfirmationNumber                    string                `json:"confirmation_number"`
	Confirmed                             bool                  `json:"confirmed"`
	ContactEmail                          string                `json:"contact_email"`
	CreatedAt                             string                `json:"created_at"`
	Currency                              string                `json:"currency"`
	CurrentShippingPriceSet               PriceSet              `json:"current_shipping_price_set"`
	CurrentSubtotalPrice                  string                `json:"current_subtotal_price"`
	CurrentSubtotalPriceSet               PriceSet              `json:"current_subtotal_price_set"`
	CurrentTotalAdditionalFeesSet         *PriceSet             `json:"current_total_additional_fees_set"`
	CurrentTotalDiscounts                 string                `json:"current_total_discounts"`
	CurrentTotalDiscountsSet              PriceSet              `json:"current_total_discounts_set"`
	CurrentTotalDutiesSet                 *PriceSet             `json:"current_total_duties_set"`
	CurrentTotalPrice                     string                `json:"current_total_price"`
	CurrentTotalPriceSet                  PriceSet              `json:"current_total_price_set"`
	CurrentTotalTax                       string                `json:"current_total_tax"`
	CurrentTotalTaxSet                    PriceSet              `json:"current_total_tax_set"`
	CustomerLocale                        string                `json:"customer_locale"`
	DeviceID                              *string               `json:"device_id"`
	DiscountCodes                         []DiscountCode        `json:"discount_codes"`
	DutiesIncluded                        bool                  `json:"duties_included"`
	Email                                 string                `json:"email"`
	EstimatedTaxes                        bool                  `json:"estimated_taxes"`
	FinancialStatus                       string                `json:"financial_status"`
	FulfillmentStatus                     *string               `json:"fulfillment_status"`
	LandingSite                           string                `json:"landing_site"`
	LandingSiteRef                        *string               `json:"landing_site_ref"`
	LocationID                            *int64                `json:"location_id"`
	MerchantBusinessEntityID              string                `json:"merchant_business_entity_id"`
	MerchantOfRecordAppID                 *int                  `json:"merchant_of_record_app_id"`
	Name                                  string                `json:"name"`
	Note                                  *string               `json:"note"`
	NoteAttributes                        []interface{}         `json:"note_attributes"`
	Number                                int                   `json:"number"`
	OrderNumber                           int                   `json:"order_number"`
	OrderStatusURL                        string                `json:"order_status_url"`
	OriginalTotalAdditionalFeesSet        *PriceSet             `json:"original_total_additional_fees_set"`
	OriginalTotalDutiesSet                *PriceSet             `json:"original_total_duties_set"`
	PaymentGatewayNames                   []string              `json:"payment_gateway_names"`
	Phone                                 *string               `json:"phone"`
	PoNumber                              *string               `json:"po_number"`
	PresentmentCurrency                   string                `json:"presentment_currency"`
	ProcessedAt                           time.Time             `json:"processed_at"`
	Reference                             *string               `json:"reference"`
	ReferringSite                         string                `json:"referring_site"`
	SourceIdentifier                      *string               `json:"source_identifier"`
	SourceName                            string                `json:"source_name"`
	SourceURL                             *string               `json:"source_url"`
	SubtotalPrice                         string                `json:"subtotal_price"`
	SubtotalPriceSet                      PriceSet              `json:"subtotal_price_set"`
	Tags                                  string                `json:"tags"`
	TaxExempt                             bool                  `json:"tax_exempt"`
	TaxLines                              []interface{}         `json:"tax_lines"`
	TaxesIncluded                         bool                  `json:"taxes_included"`
	Test                                  bool                  `json:"test"`
	Token                                 string                `json:"token"`
	TotalCashRoundingPaymentAdjustmentSet PriceSet              `json:"total_cash_rounding_payment_adjustment_set"`
	TotalCashRoundingRefundAdjustmentSet  PriceSet              `json:"total_cash_rounding_refund_adjustment_set"`
	TotalDiscounts                        string                `json:"total_discounts"`
	TotalDiscountsSet                     PriceSet              `json:"total_discounts_set"`
	TotalLineItemsPrice                   string                `json:"total_line_items_price"`
	TotalLineItemsPriceSet                PriceSet              `json:"total_line_items_price_set"`
	TotalOutstanding                      string                `json:"total_outstanding"`
	TotalPrice                            string                `json:"total_price"`
	TotalPriceSet                         PriceSet              `json:"total_price_set"`
	TotalShippingPriceSet                 PriceSet              `json:"total_shipping_price_set"`
	TotalTax                              string                `json:"total_tax"`
	TotalTaxSet                           PriceSet              `json:"total_tax_set"`
	TotalTipReceived                      string                `json:"total_tip_received"`
	TotalWeight                           int                   `json:"total_weight"`
	UpdatedAt                             time.Time             `json:"updated_at"`
	UserID                                *int64                `json:"user_id"`
	BillingAddress                        Address               `json:"billing_address"`
	Customer                              Customer              `json:"customer"`
	DiscountApplications                  []DiscountApplication `json:"discount_applications"`
	Fulfillments                          []interface{}         `json:"fulfillments"`
	LineItems                             []LineItem            `json:"line_items"`
	PaymentTerms                          *interface{}          `json:"payment_terms"`
	Refunds                               []interface{}         `json:"refunds"`
	ShippingAddress                       Address               `json:"shipping_address"`
	ShippingLines                         []ShippingLine        `json:"shipping_lines"`
	Returns                               []interface{}         `json:"returns"`
}

type ClientDetails struct {
	AcceptLanguage *string `json:"accept_language"`
	BrowserHeight  *int    `json:"browser_height"`
	BrowserIP      string  `json:"browser_ip"`
	BrowserWidth   *int    `json:"browser_width"`
	SessionHash    *string `json:"session_hash"`
	UserAgent      string  `json:"user_agent"`
}

type Money struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

type PriceSet struct {
	ShopMoney        Money `json:"shop_money"`
	PresentmentMoney Money `json:"presentment_money"`
}

type DiscountCode struct {
	Code   string `json:"code"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type Address struct {
	ID           *int64   `json:"id,omitempty"`
	CustomerID   *int64   `json:"customer_id,omitempty"`
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Company      *string  `json:"company"`
	Address1     string   `json:"address1"`
	Address2     string   `json:"address2"`
	City         string   `json:"city"`
	Province     string   `json:"province"`
	Country      string   `json:"country"`
	Zip          string   `json:"zip"`
	Phone        *string  `json:"phone"`
	Name         string   `json:"name"`
	ProvinceCode string   `json:"province_code"`
	CountryCode  string   `json:"country_code"`
	CountryName  *string  `json:"country_name,omitempty"`
	Latitude     *float64 `json:"latitude,omitempty"`
	Longitude    *float64 `json:"longitude,omitempty"`
	Default      *bool    `json:"default,omitempty"`
}

type Customer struct {
	ID                  int64         `json:"id"`
	Email               string        `json:"email"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	FirstName           string        `json:"first_name"`
	LastName            string        `json:"last_name"`
	State               string        `json:"state"`
	Note                *string       `json:"note"`
	VerifiedEmail       bool          `json:"verified_email"`
	MultipassIdentifier *string       `json:"multipass_identifier"`
	TaxExempt           bool          `json:"tax_exempt"`
	Phone               *string       `json:"phone"`
	Currency            string        `json:"currency"`
	TaxExemptions       []interface{} `json:"tax_exemptions"`
	AdminGraphqlAPIID   string        `json:"admin_graphql_api_id"`
	DefaultAddress      Address       `json:"default_address"`
}

type DiscountApplication struct {
	TargetType       string `json:"target_type"`
	Type             string `json:"type"`
	Value            string `json:"value"`
	ValueType        string `json:"value_type"`
	AllocationMethod string `json:"allocation_method"`
	TargetSelection  string `json:"target_selection"`
	Code             string `json:"code"`
}

type DiscountAllocation struct {
	Amount                   string   `json:"amount"`
	AmountSet                PriceSet `json:"amount_set"`
	DiscountApplicationIndex int      `json:"discount_application_index"`
}

type LineItem struct {
	ID                         int64                `json:"id"`
	AdminGraphqlAPIID          string               `json:"admin_graphql_api_id"`
	AttributedStaffs           []interface{}        `json:"attributed_staffs"`
	CurrentQuantity            int                  `json:"current_quantity"`
	FulfillableQuantity        int                  `json:"fulfillable_quantity"`
	FulfillmentService         string               `json:"fulfillment_service"`
	FulfillmentStatus          *string              `json:"fulfillment_status"`
	GiftCard                   bool                 `json:"gift_card"`
	Grams                      int                  `json:"grams"`
	Name                       string               `json:"name"`
	Price                      string               `json:"price"`
	PriceSet                   PriceSet             `json:"price_set"`
	ProductExists              bool                 `json:"product_exists"`
	ProductID                  int64                `json:"product_id"`
	Properties                 []interface{}        `json:"properties"`
	Quantity                   int                  `json:"quantity"`
	RequiresShipping           bool                 `json:"requires_shipping"`
	SalesLineItemGroupID       *int64               `json:"sales_line_item_group_id"`
	SKU                        string               `json:"sku"`
	Taxable                    bool                 `json:"taxable"`
	Title                      string               `json:"title"`
	TotalDiscount              string               `json:"total_discount"`
	TotalDiscountSet           PriceSet             `json:"total_discount_set"`
	VariantID                  int64                `json:"variant_id"`
	VariantInventoryManagement string               `json:"variant_inventory_management"`
	VariantTitle               *string              `json:"variant_title"`
	Vendor                     string               `json:"vendor"`
	TaxLines                   []interface{}        `json:"tax_lines"`
	Duties                     []interface{}        `json:"duties"`
	DiscountAllocations        []DiscountAllocation `json:"discount_allocations"`
}

type ShippingLine struct {
	ID                            int64         `json:"id"`
	CarrierIdentifier             *string       `json:"carrier_identifier"`
	Code                          string        `json:"code"`
	CurrentDiscountedPriceSet     PriceSet      `json:"current_discounted_price_set"`
	DiscountedPrice               string        `json:"discounted_price"`
	DiscountedPriceSet            PriceSet      `json:"discounted_price_set"`
	IsRemoved                     bool          `json:"is_removed"`
	Phone                         *string       `json:"phone"`
	Price                         string        `json:"price"`
	PriceSet                      PriceSet      `json:"price_set"`
	RequestedFulfillmentServiceID *int64        `json:"requested_fulfillment_service_id"`
	Source                        string        `json:"source"`
	Title                         string        `json:"title"`
	TaxLines                      []interface{} `json:"tax_lines"`
	DiscountAllocations           []interface{} `json:"discount_allocations"`
}
