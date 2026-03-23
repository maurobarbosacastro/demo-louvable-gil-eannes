package utils

import (
	"fmt"
	"ms-tagpeak/internal/response_object"
	"time"
)

// MonthKey Helper function to prepare month key
func MonthKey(t time.Time) string {
	return fmt.Sprintf("%d/%d", t.Month(), t.Year())
}

// InitializeMonthlyStats Helper function to initialize statistics map
func InitializeMonthlyStats(year int) map[string]response_object.StatisticsByMonth {
	statistics := make(map[string]response_object.StatisticsByMonth)
	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year, time.December, 31, 23, 59, 59, 999999999, time.UTC)

	current := startDate
	for current.Before(endDate) {
		monthKey := MonthKey(current)
		statistics[monthKey] = NewStatisticsByMonth()
		current = current.AddDate(0, 1, 0)
	}

	return statistics
}

// NewStatisticsByMonth initializes a month's statistics with default values
func NewStatisticsByMonth() response_object.StatisticsByMonth {
	return response_object.StatisticsByMonth{
		TotalUsers:           0,
		ActiveUsers:          0,
		NumTransaction:       0,
		TotalGMV:             0,
		AvgTransactionAmount: 0,
		TotalRevenue:         0,
	}
}
