package response_object

type CJAffiliateData struct {
	PublisherCommissions PublisherCommissions `json:"publisherCommissions"`
}

type PublisherCommissions struct {
	Count           int                `json:"count"`
	PayloadComplete bool               `json:"payloadComplete"`
	Records         []CommissionRecord `json:"records"`
}

// CommissionRecord represents a single commission record
type CommissionRecord struct {
	ActionStatus                   string             `json:"actionStatus"`
	ActionTrackerName              string             `json:"actionTrackerName"`
	ActionType                     string             `json:"actionType"`
	AdvertiserID                   string             `json:"advertiserId"`
	AdvertiserName                 string             `json:"advertiserName"`
	AID                            string             `json:"aid"`
	CommissionID                   string             `json:"commissionId"`
	ClickDate                      string             `json:"clickDate"`
	ClickReferringURL              string             `json:"clickReferringURL"`
	ConcludingDeviceName           string             `json:"concludingDeviceName"`
	CorrectionReason               string             `json:"correctionReason"`
	Country                        string             `json:"country"`
	Coupon                         string             `json:"coupon"`
	EventDate                      string             `json:"eventDate"`
	Items                          []CommissionItem   `json:"items"`
	LockingDate                    string             `json:"lockingDate"`
	OrderID                        string             `json:"orderId"`
	Original                       bool               `json:"original"`
	PostingDate                    string             `json:"postingDate"`
	PubCommissionAmountPubCurrency string             `json:"pubCommissionAmountPubCurrency"`
	PubCommissionAmountUSD         string             `json:"pubCommissionAmountUsd"`
	SaleAmountPubCurrency          string             `json:"saleAmountPubCurrency"`
	SaleAmountUSD                  string             `json:"saleAmountUsd"`
	VerticalAttributes             VerticalAttributes `json:"verticalAttributes"`
	WebsiteName                    string             `json:"websiteName"`
	ShopperId                      string             `json:"shopperId"`
	ValidationStatus               string             `json:"validationStatus"`
}

// CommissionItem represents an individual item within a commission
type CommissionItem struct {
	Quantity                     int     `json:"quantity"`
	PerItemSaleAmountAdvCurrency float64 `json:"perItemSaleAmountAdvCurrency"`
	PerItemSaleAmountPubCurrency float64 `json:"perItemSaleAmountPubCurrency"`
	PerItemSaleAmountUSD         float64 `json:"perItemSaleAmountUsd"`
	TotalCommissionAdvCurrency   float64 `json:"totalCommissionAdvCurrency"`
	TotalCommissionPubCurrency   float64 `json:"totalCommissionPubCurrency"`
	TotalCommissionUSD           float64 `json:"totalCommissionUsd"`
}

// VerticalAttributes represents a CJAffiliateResponseAdditional vertical-specific attributes
type VerticalAttributes struct {
	BusinessUnit string `json:"businessUnit"`
}

type AdIdResponse struct {
	ShoppingProductFeeds ShoppingProductFeeds `json:"shoppingProductFeeds"`
}

type ShoppingProductFeeds struct {
	TotalCount int                   `json:"totalCount"`
	Count      int                   `json:"count"`
	ResultList []ShoppingProductFeed `json:"resultList"`
}

type ShoppingProductFeed struct {
	SourceFeedType    string `json:"sourceFeedType"`
	ProductCount      int    `json:"productCount"`
	AdvertiserCountry string `json:"advertiserCountry"`
	LastUpdated       string `json:"lastUpdated"`
	AdvertiserName    string `json:"advertiserName"`
	Language          string `json:"language"`
	Currency          string `json:"currency"`
	AdId              string `json:"adId"`
	FeedName          string `json:"feedName"`
	AdvertiserId      string `json:"advertiserId"`
}
