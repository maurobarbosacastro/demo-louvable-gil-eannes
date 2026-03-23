package response_object

type LastPriceBulkRO struct {
	ConId         string `json:"conid"`
	LastPrice     string `json:"lastPrice"`
	LastPriceType string `json:"lastPriceType"`
}
