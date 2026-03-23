package controllers

import (
	"fmt"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetStoreVisit godoc
// @Summary Get StoreVisit by ID
// @Tags StoreVisit
// @Accept json
// @Produce json
// @Param id path string true "StoreVisit id"
// @Success 200 {object} models.StoreVisit
// @Router /store-visit/:id [get]
func GetStoreVisit(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetStoreVisit(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetAllStoreVisits godoc
// @Summary Get all StoreVisits
// @Tags StoreVisit
// @Accept json
// @Param pagination query pagination.PaginationParams true "Pagination params"
// @Param filters query dto.StoreVisitFiltersDTO true "Filters for storeVisits"
// @Produce json
// @Success 200 {array} dto.StoreVisitDTO "Array of StoreVisits"
// @Router /store-visit [get]
func GetAllStoreVisits(c echo.Context) error {

	var pag pagination.PaginationParams
	var filters dto.StoreVisitFiltersDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return c.JSON(http.StatusInternalServerError, errPag)
	}

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		return errFilters
	}

	res, err := service.GetAllStoreVisits(pag, filters, auth.KeycloakInstance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetAllStoreVisitsByUserUUID godoc
// @Summary Get all StoreVisits by user UUID
// @Tags StoreVisit
// @Accept json
// @Param pagination query pagination.PaginationParams true "Pagination params"
// @Param filters query dto.StoreVisitFiltersDTO true "Filters for storeVisits"
// @Produce json
// @Success 200 {array} dto.StoreVisitDTO "Array of StoreVisits"
// @Router /store-visit/user [get]
func GetAllStoreVisitsByUserUUID(c echo.Context) error {

	var principal = c.Get("user").(*models.User)
	var pag pagination.PaginationParams
	var filters dto.StoreVisitFiltersUuidDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return c.JSON(http.StatusInternalServerError, errPag)
	}

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		return errFilters
	}

	filters.UserUUID = &principal.Uuid

	res, err := service.GetAllStoreVisitsByUserUuid(pag, filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)

}

// CreateStoreVisit godoc
// @Summary Create StoreVisit
// @Tags StoreVisit
// @Accept json
// @Produce json
// @Param country body dto.CreateStoreVisitDTO true "Create StoreVisit dto"
// @Success 201 {object} dto.StoreVisitDTO "StoreVisit"
// @Router /store-visit [post]
func CreateStoreVisit(c echo.Context) error {
	var model dto.CreateStoreVisitDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&model); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.CreateStoreVisit(model, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)

}

// PatchStoreVisit godoc
// @Summary Update StoreVisit
// @Tags StoreVisit
// @Accept json
// @Produce json
// @Param id path string true "StoreVisit id"
// @Param StoreVisit body dto.UpdateStoreVisitDTO true "Update StoreVisit dto"
// @Success 200 {object} models.StoreVisit "StoreVisit"
// @Router /store-visit/:id [patch]
func PatchStoreVisit(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var model dto.UpdateStoreVisitDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.UpdateStoreVisit(model, uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, res)

}

// DeleteStoreVisit godoc
// @Summary Delete StoreVisit
// @Tags StoreVisit
// @Accept json
// @Produce json
// @Param id path string true "StoreVisit id"
// @Success 204
// @Router /store-visit/:id [delete]
func DeleteStoreVisit(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteStoreVisit(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)

}

// GetStoreByStoreVisit godoc
// @Summary Get StoreByStoreVisit by StoreVisit ID
// @Tags StoreVisit
// @Accept json
// @Produce json
// @Param id path string true "StoreVisit id"
// @Success 200 {object} models.Store "StoreByStoreVisit"
// @Router /store-visit/:id/store [get]
func GetStoreByStoreVisit(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetStoreByStoreVisit(uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetDistinctStoresVisitedByUserUUID godoc
// @Summary Get Stores visited by user ID
// @Tags StoreVisit
// @Accept json
// @Produce json
// @Param pagination query pagination.PaginationParams true "Pagination params"
// @Param filters query dto.StoreVisitFiltersDTO true "Filters for storeVisits"
// @Success 200 {array} models.Store
// @Router /store-visit/stores [get]
func GetDistinctStoresVisitedByUserUUID(c echo.Context) error {
	var principal = c.Get("user").(*models.User)
	var pag pagination.PaginationParams
	var filters dto.StoreVisitFiltersUuidDTO

	// Bind query params for pagination
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &pag); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Bind query params for filters
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &filters); err != nil {
		return err
	}

	// Set UserUUID from the logged-in user
	filters.UserUUID = &principal.Uuid

	// Call the service
	res, err := service.GetDistinctStoresVisitedByUserUuid(pag, filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// ValidateReference godoc
// @Summary Validate Reference by reference Id
// @Tags StoreVisit
// @Accept json
// @Produce json
// @Param id path string true "StoreVisit id"
// @Success 200
// @Failure 404
// @Router /store-visit/reference/{reference} [post]
func ValidateReference(c echo.Context) error {
	reference := c.Param("reference")

	res, err := service.ValidateReference(reference)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetStoreVisitsAdmin godoc
// @Summary Get StoreVisits admin
// @Tags StoreVisit
// @Accept json
// @Param pagination query pagination.PaginationParams true "Pagination params"
// @Param filters query dto.StoreVisitFiltersDTO true "Filters for storeVisits"
// @Produce json
// @Success 200 {array} dto.StoreVisitDTO "Array of StoreVisits"
// @Router /store-visit/admin [get]
func GetStoreVisitsAdmin(c echo.Context) error {
	var pag pagination.PaginationParams
	var filters dto.StoreVisitFiltersDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return c.JSON(http.StatusInternalServerError, errPag)
	}

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		return errFilters
	}

	res, err := service.GetStoreVisitsAdmin(pag, filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
