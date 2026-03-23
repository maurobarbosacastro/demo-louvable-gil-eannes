package response_object

type DashboardRO struct {
	CashbackSection   CashbackSection   `json:"cashbackSection"`
	IndicatorsSection IndicatorsSection `json:"indicatorsSection"`
}

type CashbackSection struct {
	TotalValidatedCashbacks int64 `json:"totalValidatedCashbacks"`
	TotalStoppedCashbacks   int64 `json:"totalStoppedCashbacks"`
	TotalPaidCashbacks      int64 `json:"totalPaidCashbacks"`
	TotalRequestedCashbacks int64 `json:"totalRequestedCashbacks"`
}

type IndicatorsSection struct {
	TotalUsers               TotalUsers               `json:"totalUsers"`
	ActiveUsers              ActiveUsers              `json:"activeUsers"`
	NumTransactions          NumTransactions          `json:"numTransactions"`
	TotalGMV                 TotalGMV                 `json:"totalGMV"`
	AverageTransactionAmount AverageTransactionAmount `json:"averageTransactionAmount"`
	TotalRevenue             TotalRevenue             `json:"totalRevenue"`
}

type TotalUsers struct {
	AllTime          int     `json:"allTime"`
	LastMonth        int     `json:"lastMonth"`
	CompareLastMonth string  `json:"compareLastMonth"`
	PercentageChange float64 `json:"percentageChange"`
}

type ActiveUsers struct {
	AllTime          int     `json:"allTime"`
	Last12Months     int     `json:"last12Months"`
	CompareLastMonth string  `json:"compareLastMonth"`
	PercentageChange float64 `json:"percentageChange"`
}

type NumTransactions struct {
	AllTime          int64   `json:"allTime"`
	CurrentMonth     int     `json:"currentMonth"`
	CompareLastMonth string  `json:"compareLastMonth"`
	PercentageChange float64 `json:"percentageChange"`
}

type TotalGMV struct {
	AllTime          float64 `json:"allTime"`
	CurrentMonth     float64 `json:"currentMonth"`
	CompareLastMonth string  `json:"compareLastMonth"`
	PercentageChange float64 `json:"percentageChange"`
}

type AverageTransactionAmount struct {
	AllTime          float64 `json:"allTime"`
	CurrentMonth     float64 `json:"currentMonth"`
	CompareLastMonth string  `json:"compareLastMonth"`
	PercentageChange float64 `json:"percentageChange"`
}

type TotalRevenue struct {
	AllTime          float64 `json:"allTime"`
	CurrentMonth     float64 `json:"currentMonth"`
	CompareLastMonth string  `json:"compareLastMonth"`
	PercentageChange float64 `json:"percentageChange"`
}

type StatisticsByMonth struct {
	TotalUsers           int     `json:"totalUsers"`
	ActiveUsers          int     `json:"activeUsers"`
	NumTransaction       int     `json:"numTransaction"`
	TotalGMV             float64 `json:"totalGMV"`
	AvgTransactionAmount float64 `json:"avgTransactionAmount"`
	TotalRevenue         float64 `json:"totalRevenue"`
}

type TransactionsDashboardRO struct {
	Value   float64 `json:"value"`
	Count   int     `json:"count"`
	Warning int     `json:"warning"`
}

type DashboardViewRO struct {
	Month                int     `json:"month"`
	Year                 int     `json:"year"`
	ActiveUsers          int     `json:"activeUsers"`
	NumTransactions      int     `json:"numTransactions"`
	TotalGMV             float64 `json:"totalGMV"`
	AvgTransactionAmount float64 `json:"avgTransactionAmount"`
	TotalRevenue         float64 `json:"totalRevenue"`
}

type CashbackDashboardRO struct {
	Status  string  `json:"status"`
	Warning int64   `json:"warning"`
	Value   float64 `json:"value"`
	Count   int64   `json:"count"`
}

type RewardByCurrencies struct {
	Currency     string `json:"currency"`
	State        string `json:"state"`
	TotalRewards int    `json:"totalRewards"`
}
