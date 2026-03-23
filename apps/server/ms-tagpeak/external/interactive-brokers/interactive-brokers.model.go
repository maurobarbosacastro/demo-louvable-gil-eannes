package interactive_brokers

type GetLastPriceBulkDto struct {
	Conids []string `json:"conids"`
}

type GetLastPriceBulkRO struct {
	Conid         string `json:"conid"`
	LastPrice     string `json:"lastPrice"`
	LastPriceType string `json:"lastPriceType"`
}

type UserMemberReward struct {
	UserUuid string
	Group    string
	Amount   *float64
	Currency string
}
