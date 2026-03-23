package controllers

import (
	"fmt"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	_ "ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCurrencyExchangeRate godoc
// @Summary Get CurrencyExchangeRate by ID
// @Tags CurrencyExchangeRate
// @Accept json
// @Produce json
// @Param id path string true "CurrencyExchangeRate id"
// @Success 200 {object} dto.CurrencyExchangeRateDTO
// @Router /currency-exchange-rate/:id [get]
func GetCurrencyExchangeRate(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetCurrencyExchangeRate(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetAllCurrencyExchangeRates godoc
// @Summary Get all CurrencyExchangeRates
// @Tags CurrencyExchangeRate
// @Accept json
// @Produce json
// @Success 200 {array} models.CurrencyExchangeRate "Array of CurrencyExchangeRates"
// @Router /currency-exchange-rate [get]
func GetAllCurrencyExchangeRates(c echo.Context) error {

	res, err := service.GetAllCurrencyExchangeRates()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)

}

// CreateCurrencyExchangeRate godoc
// @Summary Create CurrencyExchangeRate
// @Tags CurrencyExchangeRate
// @Accept json
// @Produce json
// @Param country body dto.CreateCurrencyExchangeRateDTO true "Create CurrencyExchangeRate dto"
// @Success 201 {object} models.CurrencyExchangeRate "CurrencyExchangeRate"
// @Router /currency-exchange-rate [post]
func CreateCurrencyExchangeRate(c echo.Context) error {
	var model dto.CreateCurrencyExchangeRateDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&model); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := service.CreateCurrencyExchangeRate(model)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)

}

// PatchCurrencyExchangeRate godoc
// @Summary Update CurrencyExchangeRate
// @Tags CurrencyExchangeRate
// @Accept json
// @Produce json
// @Param id path string true "CurrencyExchangeRate id"
// @Param CurrencyExchangeRate body dto.UpdateCurrencyExchangeRateDTO true "Update CurrencyExchangeRate dto"
// @Success 200 {object} models.CurrencyExchangeRate "CurrencyExchangeRate"
// @Router /currency-exchange-rate/:id [patch]
func PatchCurrencyExchangeRate(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var model dto.UpdateCurrencyExchangeRateDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.UpdateCurrencyExchangeRate(model, uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, res)

}

// DeleteCurrencyExchangeRate godoc
// @Summary Delete CurrencyExchangeRate
// @Tags CurrencyExchangeRate
// @Accept json
// @Produce json
// @Param id path string true "CurrencyExchangeRate id"
// @Success 204
// @Router /currency-exchange-rate/:id [delete]
func DeleteCurrencyExchangeRate(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteCurrencyExchangeRate(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)

}
