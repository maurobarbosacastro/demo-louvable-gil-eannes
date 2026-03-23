package controllers

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

// CreateConfiguration godoc
// @Summary Create Configuration
// @Tags Configuration
// @Accept json
// @Produce json
// @Param configuration body dto.CreateConfigurationDTO true "Create Configuration dto"
// @Success 201 {object} models.Configuration
// @Router /configuration [post]
func CreateConfiguration(c echo.Context) error {
	var createDto dto.CreateConfigurationDTO

	if err := c.Bind(&createDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&createDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.CreateConfiguration(createDto, uuidUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

// GetConfigurations godoc
// @Summary Get Configurations
// @Tags Configuration
// @Accept json
// @Produce json
// @Param pagination query pagination.PaginationParams true "Pagination params"
// @Success 200 {object} pagination.PaginationResult{data=[]models.Configuration} "Array of Configurations"
// @Router /configuration [get]
func GetConfigurations(c echo.Context) error {
	var pag pagination.PaginationParams

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}

	configs, err := service.GetConfigurations(pag)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, configs)
}

// GetConfiguration godoc
// @Summary Get Configuration by code
// @Tags Configuration
// @Accept json
// @Produce json
// @Param code path string true "Configuration code"
// @Success 200 {object} models.Configuration
// @Router /configuration/:code [get]
func GetConfiguration(c echo.Context) error {
	code := c.Param("code")

	res, err := service.GetConfiguration(code)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)
}

// UpdateConfiguration godoc
// @Summary Update Configuration
// @Tags Configuration
// @Accept json
// @Produce json
// @Param id path string true "Configuration id"
// @Param configuration body dto.UpdateConfigurationDTO true "Update Configuration dto"
// @Success 200 {object} models.Configuration
// @Router /configuration/:id [patch]
func UpdateConfiguration(c echo.Context) error {
	id := utils.ParseIdToInt(c.Param("id"))
	var updateDto dto.UpdateConfigurationDTO

	if err := c.Bind(&updateDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&updateDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.UpdateConfiguration(id, updateDto, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	service.LoadConfigurations()

	return c.JSON(http.StatusOK, res)
}

// DeleteConfiguration godoc
// @Summary Delete Configuration
// @Tags Configuration
// @Accept json
// @Produce json
// @Param id path string true "Configuration id"
// @Success 204
// @Router /configuration/:id [delete]
func DeleteConfiguration(c echo.Context) error {
	uuidUser := c.Get("user").(*models.User).Uuid.String()
	id := utils.ParseIdToInt(c.Param("id"))

	err := service.DeleteConfiguration(id, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

// GetLatestConfigurations godoc
// @Summary GetLatestConfigurations
// @Tags Configuration
// @Accept json
// @Produce json
// @Success 200 {object} map[string]models.Configuration
// @Router /configuration/latest [get]
func GetLatestConfigurations(c echo.Context) error {
	config, err := service.GetAllConfigurations()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	final := lo.KeyBy(config, func(c models.Configuration) string {
		return c.Code
	})

	return c.JSON(http.StatusOK, final)
}
