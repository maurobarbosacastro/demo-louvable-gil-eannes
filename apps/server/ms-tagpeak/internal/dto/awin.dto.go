package dto

type TransactionAwinDTO struct {
	ID                              int64             `json:"id"`
	URL                             string            `json:"url"`
	AdvertiserID                    int64             `json:"advertiserId"`
	PublisherID                     int64             `json:"publisherId"`
	CommissionSharingPublisherID    *int64            `json:"commissionSharingPublisherId"`
	CommissionSharingSelectedRateID *int64            `json:"commissionSharingSelectedRatePublisherId"`
	Campaign                        *string           `json:"campaign"`
	SiteName                        string            `json:"siteName"`
	CommissionStatus                string            `json:"commissionStatus"`
	CommissionAmount                MonetaryAmount    `json:"commissionAmount"`
	SaleAmount                      MonetaryAmount    `json:"saleAmount"`
	IPHash                          string            `json:"ipHash"`
	CustomerCountry                 string            `json:"customerCountry"`
	ClickRefs                       ClickReferences   `json:"clickRefs"`
	ClickDate                       string            `json:"clickDate"`
	TransactionDate                 string            `json:"transactionDate"`
	ValidationDate                  string            `json:"validationDate"`
	Type                            string            `json:"type"`
	DeclineReason                   *string           `json:"declineReason"`
	VoucherCodeUsed                 bool              `json:"voucherCodeUsed"`
	VoucherCode                     *string           `json:"voucherCode"`
	LapseTime                       int               `json:"lapseTime"`
	Amended                         bool              `json:"amended"`
	AmendReason                     *string           `json:"amendReason"`
	OldSaleAmount                   *MonetaryAmount   `json:"oldSaleAmount"`
	OldCommissionAmount             *MonetaryAmount   `json:"oldCommissionAmount"`
	ClickDevice                     string            `json:"clickDevice"`
	TransactionDevice               string            `json:"transactionDevice"`
	CustomerAcquisition             *string           `json:"customerAcquisition"`
	PublisherURL                    string            `json:"publisherUrl"`
	AdvertiserCountry               string            `json:"advertiserCountry"`
	OrderRef                        *string           `json:"orderRef"`
	CustomParameters                *[]DefaultStruct  `json:"customParameters"`
	TransactionParts                []TransactionPart `json:"transactionParts"`
	PaidToPublisher                 bool              `json:"paidToPublisher"`
	PaymentID                       int               `json:"paymentId"`
	TransactionQueryID              int               `json:"transactionQueryId"`
	TrackedCurrencyAmount           MonetaryAmount    `json:"trackedCurrencyAmount"`
	OriginalSaleAmount              *MonetaryAmount   `json:"originalSaleAmount"`
	AdvertiserCost                  MonetaryAmount    `json:"advertiserCost"`
	BasketProducts                  *string           `json:"basketProducts"`
}

type MonetaryAmount struct {
	Amount   *float64 `json:"amount"`
	Currency *string  `json:"currency"`
}

type ClickReferences struct {
	ClickRef string `json:"clickRef"`
}

type TransactionPart struct {
	CommissionGroupID   int           `json:"commissionGroupId"`
	Amount              float64       `json:"amount"`
	CommissionAmount    float64       `json:"commissionAmount"`
	AdvertiserCost      *float64      `json:"advertiserCost"`
	CommissionGroupCode string        `json:"commissionGroupCode"`
	CommissionGroupName string        `json:"commissionGroupName"`
	TrackedParts        []TrackedPart `json:"trackedParts"`
}

type TrackedPart struct {
	Code     string  `json:"code"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type DefaultStruct struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
