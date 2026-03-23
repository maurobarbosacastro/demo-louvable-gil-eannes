package controllers

import (
	"github.com/labstack/echo/v4"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/service"
	"net/http"
)

// GetValuesDashboard godoc
// @Summary Get Values Dashboard
// @Tags Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} response_object.DashboardRO
// @Router /dashboard [get]
func GetValuesDashboard(c echo.Context) error {
	keycloak := auth.KeycloakInstance
	res, err := service.GetValuesDashboard(keycloak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// GetStatisticsByMonth godoc
// @Summary Get Statistics By Month
// @Tags Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} map[string]response_object.StatisticsByMonth
// @Router /dashboard/statistics [get]
func GetStatisticsByMonth(c echo.Context) error {
	var filters dto.DashboardFiltersDTO

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		return errFilters
	}

	keycloak := auth.KeycloakInstance

	res, err := service.GetDashboardStatisticsByMonth(keycloak, filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// GetTransactionsDashboard godoc
// @Summary Get Transaction dashboard
// @Tags Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} map[string]response_object.TransactionsDashboardRO
// @Router /dashboard/transactions [get]
func GetTransactionsDashboard(c echo.Context) error {
	res, err := service.GetTransactionsDashboard()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// GetRewardCountByCurrencies godoc
// @Summary Get Reward Count By Currencies
// @Tags Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} map[string]response_object.RewardByCurrencies
// @Router /dashboard/rewards/currencies/count [get]
func GetRewardCountByCurrencies(c echo.Context) error {
	res, err := service.GetRewardsByCurrencies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
