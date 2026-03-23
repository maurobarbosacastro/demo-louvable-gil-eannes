package controllers

import (
	"fmt"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	_ "ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	pagination "ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCountry godoc
// @Summary Get Country by ID
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "Country id"
// @Success 200 {object} models.Country
// @Router /country/:id [get]
func GetCountry(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	countryDTO, err := service.GetCountry(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, countryDTO)

}

// GetCountryByCode godoc
// @Summary Get Country by ID
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "Country id"
// @Success 200 {object} models.Country
// @Router /public/country/code/:code [get]
func GetCountryByCode(c echo.Context) error {
	code := c.Param("code")

	countryDTO, err := service.GetCountryByCode(code)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, countryDTO)

}

// GetCountries godoc
// @Summary Get all Countries
// @Tags Country
// @Accept json
// @Param pagination query pagination.PaginationParams true "Pagination params"
// @Param filters query dto.CountryFiltersDTO true "Filters for countries"
// @Produce json
// @Success 200 {array} models.Country "Array of Countries"
// @Router /country [get]
func GetCountries(c echo.Context) error {

	var pag pagination.PaginationParams
	var filters dto.CountryFiltersDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		return errFilters
	}

	res, err := service.GetCountries(pag, filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// CreateCountry godoc
// @Summary Create Country
// @Tags Country
// @Accept json
// @Produce json
// @Param country body dto.CreateCountryDTO true "Create country dto"
// @Success 201 {object} models.Country "Country"
// @Router /country [post]
func CreateCountry(c echo.Context) error {
	var country dto.CreateCountryDTO

	if err := c.Bind(&country); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&country); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	countryDTO, err := service.CreateCountry(country, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, countryDTO)

}

// PatchCountry godoc
// @Summary Update Country
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "Country id"
// @Param country body dto.UpdateCountryDTO true "Update country dto"
// @Success 200 {object} models.Country "Country"
// @Router /country/:id [patch]
func PatchCountry(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var country dto.UpdateCountryDTO

	if err := c.Bind(&country); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	countryDto, err := service.UpdateCountry(country, uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, countryDto)

}

// DeleteCountry godoc (Soft delete)
// @Summary Delete Country
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "Country id"
// @Success 204
// @Router /country/:id [delete]
func DeleteCountry(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.SoftDeleteCountry(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)

}
