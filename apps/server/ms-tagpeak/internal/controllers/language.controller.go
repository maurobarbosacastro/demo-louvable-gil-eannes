package controllers

import (
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	_ "ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetLanguage godoc
// @Summary Get Language by ID
// @Tags Language
// @Accept json
// @Produce json
// @Param id path string true "Language id"
// @Success 200 {object} models.Language
// @Router /language/:id [get]
func GetLanguage(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetLanguage(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetAllLanguages godoc
// @Summary Get all Languages
// @Tags Language
// @Accept json
// @Produce json
// @Success 200 {array} models.Language "Array of Languages"
// @Router /language [get]
func GetAllLanguages(c echo.Context) error {

	var pag pagination.PaginationParams

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}

	res, err := service.GetAllLanguages(pag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)

}

// CreateLanguage godoc
// @Summary Create Language
// @Tags Language
// @Accept json
// @Produce json
// @Param country body dto.CreateLanguageDTO true "Create Language dto"
// @Success 201 {object} models.Language "Language"
// @Router /language [post]
func CreateLanguage(c echo.Context) error {
	var model dto.CreateLanguageDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	exists, err := service.LanguageCodeExist(model.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if exists {
		return c.JSON(http.StatusConflict, utils.CustomErrorStruct{}.ConflictError("Language code already exists"))
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.CreateLanguage(model, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)

}

// PatchLanguage godoc
// @Summary Update Language
// @Tags Language
// @Accept json
// @Produce json
// @Param id path string true "Language id"
// @Param Language body dto.UpdateLanguageDTO true "Update Language dto"
// @Success 200 {object} models.Language "Language"
// @Router /language/:id [patch]
func PatchLanguage(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var model dto.UpdateLanguageDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.UpdateLanguage(model, uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, res)

}

// DeleteLanguage godoc
// @Summary Delete Language
// @Tags Language
// @Accept json
// @Produce json
// @Param id path string true "Language id"
// @Success 204
// @Router /language/:id [delete]
func DeleteLanguage(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteLanguage(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)

}
