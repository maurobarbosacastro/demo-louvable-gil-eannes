package response_object

type UserStatsRO struct {
	Level      string  `json:"level"`
	ValueSpent float64 `json:"valueSpent"`
}
