package controllers

import (
	"fmt"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetPublicAllStores godoc
// @Summary Get all Stores
// @Tags PublicStore
// @Accept json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Param filters body dto.StoreFiltersDTO true "Filters for stores"
// @Produce json
// @Success 200 {object} pagination.PaginationResult{data=[]response_object.PublicStore} "Array of Stores"
// @Router /public/store [get]
func GetPublicAllStores(c echo.Context) error {

	var pag pagination.PaginationParams
	var filters dto.StoreFiltersDTO

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

	res, err := service.GetAllStores(pag, filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetPublicStore godoc
// @Summary Get Store by ID
// @Tags PublicStore
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Success 200 {object} response_object.GetStorePublicRO
// @Router /public/store/:id [get]
func GetPublicStore(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetStore(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	response := response_object.GetStorePublicRO{
		Uuid:                        res.Uuid,
		Name:                        res.Name,
		Logo:                        res.Logo,
		Banner:                      res.Banner,
		ShortDescription:            res.ShortDescription,
		Description:                 res.Description,
		AverageRewardActivationTime: res.AverageRewardActivationTime,
		Keywords:                    res.Keywords,
		StoreUrl:                    res.StoreUrl,
		TermsAndConditions:          res.TermsAndConditions,
		PercentageCashout:           res.PercentageCashout,
		MetaTitle:                   res.MetaTitle,
		MetaKeywords:                res.MetaKeywords,
		MetaDescription:             res.MetaDescription,
	}

	return c.JSON(http.StatusOK, response)

}

// GetPublicCountries godoc
// @Summary Get all Countries
// @Tags PublicStore
// @Accept json
// @Param pagination query pagination.PaginationParams true "Pagination params"
// @Param filters query dto.CountryFiltersDTO true "Filters for countries"
// @Produce json
// @Success 200 {array} models.Country "Array of Countries"
// @Router /public/country [get]
func GetPublicCountries(c echo.Context) error {

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

// GetPublicCategories godoc
// @Summary Get all Categories
// @Tags PublicStore
// @Accept json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Produce json
// @Success 200 {array} models.Category "Array of Categories"
// @Router /public/category [get]
func GetPublicCategories(c echo.Context) error {

	var pag pagination.PaginationParams
	var filters dto.CategoryFiltersDTO

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

	res, err := service.GetAllCategories(pag, &filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetStoreRedirectUrl godoc
// @Summary Get store redirect url
// @Description Get store redirect url
// @Tags PublicStore
// @Produce json
// @Param id path string true "Store ID"
// @Success 200 {object} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /store/:id/redirect [get]
func GetStoreRedirectUrl(c echo.Context) error {
	logster.StartFuncLog()

	uuid := utils.ParseIDToUUID(c.Param("id"))
	uuidUser := c.Get("user").(*models.User).Uuid.String()

	//Get store
	store, err := service.GetStore(uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	url, err := service.GenerateRefAndGetStoreLink(store, uuidUser, auth.KeycloakInstance)
	if err != nil {
		logster.Error(err, "Error GenerateRefAndGetStoreLink")
		return c.JSON(http.StatusInternalServerError, err)
	}
	if url == nil || *url == "" {
		logster.Warn(fmt.Sprintf("GenerateRefAndGetStoreLink - Empty"))
		return c.JSON(http.StatusNoContent, nil)
	}

	logster.EndFuncLogMsg(fmt.Sprintf("GetStoreRedirectUrl - End: %s", *url))
	return c.JSON(http.StatusOK, url)
}
