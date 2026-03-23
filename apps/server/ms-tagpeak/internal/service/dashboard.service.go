package service

import (
	"fmt"
	"math"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"strconv"
	"time"
)

func GetValuesDashboard(keycloak *constants.Keycloak) (response_object.DashboardRO, error) {
	logster.StartFuncLog()

	cashbackSection, err := getCashbackSection()
	if err != nil {
		return response_object.DashboardRO{}, err
	}

	indicatorsSection, err := getIndicatorsSection(keycloak)
	if err != nil {
		return response_object.DashboardRO{}, err
	}

	logster.EndFuncLog()
	return response_object.DashboardRO{
		CashbackSection:   cashbackSection,
		IndicatorsSection: indicatorsSection,
	}, nil
}

func getCashbackSection() (response_object.CashbackSection, error) {
	logster.StartFuncLog()
	states := []string{"VALIDATED", "STOPPED", "REQUESTED", "PAID"}
	totalsCashbacks := make(map[string]int64)

	for _, state := range states {
		filters := dto.TransactionFiltersDTO{
			State: utils.StringPointer(state),
		}
		_, total, err := repository.GetCashbackView(pagination.PaginationParams{}, filters)
		if err != nil {
			return response_object.CashbackSection{}, err
		}
		totalsCashbacks[state] = total
	}

	logster.EndFuncLog()
	return response_object.CashbackSection{
		TotalValidatedCashbacks: totalsCashbacks["VALIDATED"],
		TotalStoppedCashbacks:   totalsCashbacks["STOPPED"],
		TotalPaidCashbacks:      totalsCashbacks["PAID"],
		TotalRequestedCashbacks: totalsCashbacks["REQUESTED"],
	}, nil
}

func getIndicatorsSection(keycloak *constants.Keycloak) (response_object.IndicatorsSection, error) {
	logster.StartFuncLog()
	totalUsers, err := getTotalUsers(keycloak)
	if err != nil {
		return response_object.IndicatorsSection{}, err
	}

	activeUsers, err := getActiveUsers()
	if err != nil {
		return response_object.IndicatorsSection{}, err
	}

	numTransactions, err := getNumTransactions()
	if err != nil {
		return response_object.IndicatorsSection{}, err
	}

	totalGMV, err := getTotalGMV()
	if err != nil {
		return response_object.IndicatorsSection{}, err
	}

	avgTransAmount, err := getAverageTransactionAmount()
	if err != nil {
		return response_object.IndicatorsSection{}, err
	}

	totalRevenue, err := getTotalRevenue()
	if err != nil {
		return response_object.IndicatorsSection{}, err
	}

	logster.EndFuncLog()
	return response_object.IndicatorsSection{
		TotalUsers:               totalUsers,
		ActiveUsers:              activeUsers,
		NumTransactions:          numTransactions,
		TotalGMV:                 totalGMV,
		AverageTransactionAmount: avgTransAmount,
		TotalRevenue:             totalRevenue,
	}, nil
}

func getTotalUsers(keycloak *constants.Keycloak) (response_object.TotalUsers, error) {
	logster.StartFuncLog()
	users, err := GetMaxUsers(keycloak, nil)
	if err != nil {
		return response_object.TotalUsers{}, err
	}
	logster.Info(fmt.Sprintf("Found %v users", len(users)))

	timeRanges := utils.GetTimeRanges()
	lastMonth := timeRanges["lastMonth"]
	currentMonth := timeRanges["currentMonth"]

	var (
		lastMonthUsers    = 0
		currentMonthUsers = 0
	)
	for _, user := range users {
		if user.CreatedAt >= lastMonth.Start.UnixMilli() && user.CreatedAt <= lastMonth.End.UnixMilli() {
			lastMonthUsers++
		}
		if user.CreatedAt >= currentMonth.Start.UnixMilli() && user.CreatedAt <= currentMonth.End.UnixMilli() {
			currentMonthUsers++
		}
	}

	compareLastMonth, percentageChange := compareMonths(float64(currentMonthUsers), float64(lastMonthUsers))

	logster.EndFuncLog()
	return response_object.TotalUsers{
		AllTime:          len(users),
		LastMonth:        lastMonthUsers,
		CompareLastMonth: compareLastMonth,
		PercentageChange: percentageChange,
	}, nil
}

func getActiveUsers() (response_object.ActiveUsers, error) {
	logster.StartFuncLog()

	activeUsersAllTime, err := repository.GetActiveUsersAllTime()
	logster.Info(fmt.Sprintf("Active users all time: %v", activeUsersAllTime))

	if err != nil {
		logster.Error(err, "Error getting active users all time")
		return response_object.ActiveUsers{}, err
	}

	activeUsers12months, err := repository.GetActiveUsersLast12Months()
	logster.Info(fmt.Sprintf("Active users last 12 months: %v", activeUsers12months))

	if err != nil {
		logster.Error(err, "Error getting active users last 12 months")
		return response_object.ActiveUsers{}, err
	}

	activeUserCurrentMonth, err := repository.GetActiveUserCurrentMonth()
	logster.Info(fmt.Sprintf("Active user current month: %v", activeUserCurrentMonth))

	if err != nil {
		logster.Error(err, "Error getting active user current month")
		return response_object.ActiveUsers{}, err
	}

	activeUserLastMonth, err := repository.GetActiveUserLastMonth()
	logster.Info(fmt.Sprintf("Active user last month: %v", activeUserLastMonth))

	if err != nil {
		logster.Error(err, "Error getting active user last month")
		return response_object.ActiveUsers{}, err
	}

	compareLastMonth, percentageChange := compareMonths(float64(activeUserCurrentMonth), float64(activeUserLastMonth))

	logster.EndFuncLog()
	return response_object.ActiveUsers{
		AllTime:          int(activeUsersAllTime),
		Last12Months:     int(activeUsers12months),
		CompareLastMonth: compareLastMonth,
		PercentageChange: percentageChange,
	}, nil
}

func getNumTransactions() (response_object.NumTransactions, error) {
	logster.StartFuncLog()

	currentTime := time.Now()
	lastMonth := currentTime.AddDate(0, -1, 0).Month()

	var month = make([]string, 0)
	month = append(month, strconv.Itoa(int(currentTime.Month())))
	month = append(month, strconv.Itoa(int(lastMonth)))
	dashboardInfo, err := repository.GetDashboardInfo(currentTime.Year(), month)
	if err != nil {
		return response_object.NumTransactions{}, err
	}
	logster.Info(fmt.Sprintf("Found %v dashboard info on year %v and between months %v - %v", len(dashboardInfo), currentTime.Year(), currentTime.Month(), lastMonth))

	totalTransactions, err := repository.GetTotalTransactions()
	if err != nil {
		return response_object.NumTransactions{}, err
	}
	logster.Info(fmt.Sprintf("Found %v total transactions", totalTransactions))

	var (
		currentMonthCount = 0
		lastMonthCount    = 0
	)
	for _, row := range dashboardInfo {
		if row.Month == int(currentTime.Month()) {
			currentMonthCount = row.NumTransactions
		}
		if row.Month == int(lastMonth) {
			lastMonthCount = row.NumTransactions
		}
	}

	compareLastMonth, percentageChange := compareMonths(float64(currentMonthCount), float64(lastMonthCount))

	logster.EndFuncLog()
	return response_object.NumTransactions{
		AllTime:          totalTransactions,
		CurrentMonth:     currentMonthCount,
		CompareLastMonth: compareLastMonth,
		PercentageChange: percentageChange,
	}, nil
}

func getTotalGMV() (response_object.TotalGMV, error) {
	logster.StartFuncLog()

	currentTime := time.Now()
	lastMonth := currentTime.AddDate(0, -1, 0).Month()

	var month = make([]string, 0)
	month = append(month, strconv.Itoa(int(currentTime.Month())))
	month = append(month, strconv.Itoa(int(lastMonth)))
	dashboardInfo, err := repository.GetDashboardInfo(currentTime.Year(), month)
	if err != nil {
		return response_object.TotalGMV{}, err
	}
	logster.Info(fmt.Sprintf("Found %v dashboard info on year %v and between months %v - %v", len(dashboardInfo), currentTime.Year(), currentTime.Month(), lastMonth))

	allTimeTotal, err := repository.GetTotalGMVAllTime()
	if err != nil {
		return response_object.TotalGMV{}, err
	}
	logster.Info(fmt.Sprintf("Found %v total GMV on year %v", allTimeTotal, currentTime.Year()))

	var (
		currentMonthTotal = 0.0
		lastMonthTotal    = 0.0
	)
	for _, row := range dashboardInfo {
		if row.Month == int(currentTime.Month()) {
			currentMonthTotal = row.TotalGMV
		}
		if row.Month == int(lastMonth) {
			lastMonthTotal = row.TotalGMV
		}
	}

	compareLastMonth, percentageChange := compareMonths(currentMonthTotal, lastMonthTotal)

	logster.EndFuncLog()
	return response_object.TotalGMV{
		AllTime:          utils.Round2Digits(allTimeTotal),
		CurrentMonth:     utils.Round2Digits(currentMonthTotal),
		CompareLastMonth: compareLastMonth,
		PercentageChange: percentageChange,
	}, nil
}

func getAverageTransactionAmount() (response_object.AverageTransactionAmount, error) {
	logster.StartFuncLog()

	currentTime := time.Now()
	lastMonth := currentTime.AddDate(0, -1, 0).Month()

	var month = make([]string, 0)
	month = append(month, strconv.Itoa(int(currentTime.Month())))
	month = append(month, strconv.Itoa(int(lastMonth)))
	dashboardInfo, err := repository.GetDashboardInfo(currentTime.Year(), month)
	if err != nil {
		return response_object.AverageTransactionAmount{}, err
	}
	logster.Info(fmt.Sprintf("Found %v dashboard info on year %v and between months %v - %v", len(dashboardInfo), currentTime.Year(), currentTime.Month(), lastMonth))

	avgAllTime, err := repository.GetAvgTransactionAmount()
	if err != nil {
		return response_object.AverageTransactionAmount{}, err
	}
	logster.Info(fmt.Sprintf("Found %v total GMV on year %v", avgAllTime, currentTime.Year()))

	var (
		avgCurrentMonth = 0.0
		avgLastMonth    = 0.0
	)
	for _, row := range dashboardInfo {
		if row.Month == int(currentTime.Month()) {
			avgCurrentMonth = row.AvgTransactionAmount
		}
		if row.Month == int(lastMonth) {
			avgLastMonth = row.AvgTransactionAmount
		}
	}

	compareLastMonth, percentageChange := compareMonths(avgCurrentMonth, avgLastMonth)

	logster.EndFuncLog()
	return response_object.AverageTransactionAmount{
		AllTime:          utils.Round2Digits(avgAllTime),
		CurrentMonth:     utils.Round2Digits(avgCurrentMonth),
		CompareLastMonth: compareLastMonth,
		PercentageChange: percentageChange,
	}, nil
}

func getTotalRevenue() (response_object.TotalRevenue, error) {
	logster.StartFuncLog()

	currentTime := time.Now()
	lastMonth := currentTime.AddDate(0, -1, 0).Month()

	var month = make([]string, 0)
	month = append(month, strconv.Itoa(int(currentTime.Month())))
	month = append(month, strconv.Itoa(int(lastMonth)))
	dashboardInfo, err := repository.GetDashboardInfo(currentTime.Year(), month)
	if err != nil {
		return response_object.TotalRevenue{}, err
	}
	logster.Info(fmt.Sprintf("Found %v dashboard info on year %v and between months %v - %v", len(dashboardInfo), currentTime.Year(), currentTime.Month(), lastMonth))

	allTimeRevenue, err := repository.GetTotalRevenue()
	if err != nil {
		return response_object.TotalRevenue{}, err
	}
	logster.Info(fmt.Sprintf("Found %v total GMV on year %v", allTimeRevenue, currentTime.Year()))

	var (
		currentMonthRevenue = 0.0
		lastMonthRevenue    = 0.0
	)
	for _, row := range dashboardInfo {
		if row.Month == int(currentTime.Month()) {
			currentMonthRevenue = row.AvgTransactionAmount
		}
		if row.Month == int(lastMonth) {
			lastMonthRevenue = row.AvgTransactionAmount
		}
	}

	compareLastMonth, percentageChange := compareMonths(currentMonthRevenue, lastMonthRevenue)

	logster.EndFuncLog()
	return response_object.TotalRevenue{
		AllTime:          utils.Round2Digits(allTimeRevenue),
		CurrentMonth:     utils.Round2Digits(currentMonthRevenue),
		CompareLastMonth: compareLastMonth,
		PercentageChange: percentageChange,
	}, nil
}

func GetDashboardStatisticsByMonth(keycloak *constants.Keycloak, filters dto.DashboardFiltersDTO) (map[string]response_object.StatisticsByMonth, error) {
	logster.StartFuncLog()
	statistics := utils.InitializeMonthlyStats(filters.Year)

	users, err := GetMaxUsers(keycloak, nil)
	if err != nil {
		return nil, err
	}

	statistics, err = statisticsDashboard(users, statistics, filters)
	if err != nil {
		return nil, err
	}

	logster.EndFuncLog()
	return statistics, nil
}

func statisticsDashboard(users []*models.User, statistics map[string]response_object.StatisticsByMonth, filters dto.DashboardFiltersDTO) (map[string]response_object.StatisticsByMonth, error) {
	logster.StartFuncLog()
	activeUsersByMonth := make(map[string]map[string]bool)

	// Process total users first.
	for _, user := range users {
		createdAt := time.UnixMilli(user.CreatedAt)
		if createdAt.Year() == filters.Year {
			monthKey := utils.MonthKey(createdAt)
			if stats, exists := statistics[monthKey]; exists {
				stats.TotalUsers++
				statistics[monthKey] = stats
			}
		}
	}

	var month []string
	dashboardInfo, err := repository.GetDashboardInfo(filters.Year, month)
	if err != nil {
		return nil, err
	}
	logster.Info(fmt.Sprintf("Found %v dashboard info on year %v", len(dashboardInfo), filters.Year))

	for _, row := range dashboardInfo {
		date := time.Date(row.Year, time.Month(row.Month), 1, 0, 0, 0, 0, time.UTC)
		monthKey := utils.MonthKey(date)

		// Initialize month map if it doesn't exist
		if _, exists := activeUsersByMonth[monthKey]; !exists {
			activeUsersByMonth[monthKey] = make(map[string]bool)
		}

		if stats, exists := statistics[monthKey]; exists {
			stats.ActiveUsers = row.ActiveUsers
			stats.NumTransaction = row.NumTransactions
			stats.AvgTransactionAmount = utils.Round2Digits(row.AvgTransactionAmount)
			stats.TotalGMV = utils.Round2Digits(row.TotalGMV)
			stats.TotalRevenue = utils.Round2Digits(row.TotalRevenue)
			statistics[monthKey] = stats
		}
	}

	logster.EndFuncLog()
	return statistics, nil
}

func GetTransactionsDashboard() (map[string]response_object.TransactionsDashboardRO, error) {
	logster.StartFuncLog()

	cashbackRes := make(map[string]response_object.TransactionsDashboardRO)

	cashbackDashboard, err := repository.GetCashbackDashboard()
	if err != nil {
		return nil, err
	}
	logster.Info(fmt.Sprintf("Found %v cashbacks", cashbackDashboard))

	for _, cashback := range cashbackDashboard {
		cashbackStatus := cashback.Status
		status := cashbackRes[cashbackStatus]
		status.Count = int(cashback.Count)
		status.Warning = int(cashback.Warning)
		status.Value = float64(int(cashback.Value))
		cashbackRes[cashbackStatus] = status
	}

	logster.EndFuncLog()
	return cashbackRes, nil
}

func compareMonths(currentMonth, lastMonth float64) (string, float64) {
	// Handle division by zero case
	if lastMonth == 0 {
		if currentMonth == 0 {
			return "EQUAL", 0
		}
		return "IMPROVEMENT", 100 // Special case for growth from zero
	}

	diff := currentMonth - lastMonth
	percentageChange := (diff / lastMonth) * 100

	// Avoid floating point precision issues by using a small epsilon
	const epsilon = 0.000001
	switch {
	case diff > epsilon:
		return "IMPROVEMENT", utils.Round2Digits(percentageChange)
	case diff < -epsilon:
		return "DOWNGRADE", utils.Round2Digits(math.Abs(percentageChange))
	default:
		return "EQUAL", 0
	}
}

func GetRewardsByCurrencies() (map[string]map[string]response_object.RewardByCurrencies, error) {
	logster.StartFuncLog()

	// Structure: map[state]map[currency]RewardByCurrencies
	rewardsCurrencies := make(map[string]map[string]response_object.RewardByCurrencies)

	rewardsCurrencyRes, err := repository.GetRewardsByCurrencies()
	if err != nil {
		return nil, err
	}
	logster.Info(fmt.Sprintf("Found %v rewards by currencies", len(rewardsCurrencyRes)))

	for _, rewardCurrency := range rewardsCurrencyRes {
		state := rewardCurrency.State
		currency := rewardCurrency.Currency

		// Initialize the state map if it doesn't exist
		if rewardsCurrencies[state] == nil {
			rewardsCurrencies[state] = make(map[string]response_object.RewardByCurrencies)
		}

		// Set the currency data within the state
		rewardsCurrencies[state][currency] = response_object.RewardByCurrencies{
			State:        rewardCurrency.State,
			Currency:     rewardCurrency.Currency,
			TotalRewards: rewardCurrency.TotalRewards,
		}
	}

	logster.EndFuncLog()
	return rewardsCurrencies, nil
}
