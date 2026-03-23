package controllers

import (
	"fmt"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	_ "ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetPartner godoc
// @Summary Get Partner by ID
// @Tags Partner
// @Accept json
// @Produce json
// @Param id path string true "Partner id"
// @Success 200 {object} models.Partner
// @Router /partner/:id [get]
func GetPartner(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetPartner(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetAllPartners godoc
// @Summary Get all Partners
// @Tags Partner
// @Accept json
// @Param pagination query pagination.PaginationParams true "Pagination params"
// @Produce json
// @Success 200 {array} models.Partner "Array of Partners"
// @Router /partner [get]
func GetAllPartners(c echo.Context) error {

	var pag pagination.PaginationParams
	var filters dto.PartnerFiltersDTO

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

	fmt.Printf("pagination: %v\n", pag)
	res, err := service.GetAllPartners(pag, filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)

}

// CreatePartner godoc
// @Summary Create Partner
// @Tags Partner
// @Accept json
// @Produce json
// @Param country body dto.CreatePartnerDTO true "Create partner dto"
// @Success 200 {object} models.Partner "Partner"
// @Router /partner [post]
func CreatePartner(c echo.Context) error {
	var model dto.CreatePartnerDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	exists, err := service.CodeAlreadyExists(*model.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if exists {
		return c.JSON(http.StatusConflict, utils.CustomErrorStruct{}.ConflictError("Partner code already exists"))
	}

	var uuidUser string

	if c.Get("user").(*models.User) == nil {
		uuidUser = c.Get("user").(*models.User).Uuid.String()
	} else {
		uuidUser = "system"
	}

	res, err := service.CreatePartner(model, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)

}

// PatchPartner godoc
// @Summary Update Partner
// @Tags Partner
// @Accept json
// @Produce json
// @Param id path string true "Partner id"
// @Param partner body dto.UpdatePartnerDTO true "Update partner dto"
// @Success 200 {object} models.Partner "Partner"
// @Router /partner/:id [patch]
func PatchPartner(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var model dto.UpdatePartnerDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.UpdatePartner(model, uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, res)

}

// DeletePartner godoc
// @Summary Delete Partner
// @Tags Partner
// @Accept json
// @Produce json
// @Param id path string true "Partner id"
// @Success 204
// @Router /partner/:id [delete]
func DeletePartner(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeletePartner(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)

}
