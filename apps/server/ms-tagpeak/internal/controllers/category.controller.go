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

// GetCategory godoc
// @Summary Get Category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} models.Category
// @Router /category/:id [get]
func GetCategory(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetCategory(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetCategoryByCode godoc
// @Summary Get Category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} models.Category
// @Router /category/code/:code [get]
func GetCategoryByCode(c echo.Context) error {
	code := c.Param("code")

	res, err := service.GetCategoryByCode(code)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetAllCategories godoc
// @Summary Get all Categories
// @Tags Category
// @Accept json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Produce json
// @Success 200 {array} models.Category "Array of Categories"
// @Router /category [get]
func GetAllCategories(c echo.Context) error {

	var pag pagination.PaginationParams
	var filters dto.CategoryFiltersDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}
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

// CreateCategory godoc
// @Summary Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param country body dto.CreateCategoryDTO true "Create category dto"
// @Success 200 {object} models.Category "Category"
// @Router /category [post]
func CreateCategory(c echo.Context) error {
	var category dto.CreateCategoryDTO

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	exists, err := service.CategoryCodeExist(*category.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if exists {
		return c.JSON(http.StatusConflict, utils.CustomErrorStruct{}.ConflictError("Category code already exists"))
	}

	var uuidUser string

	if c.Get("user").(*models.User) == nil {
		uuidUser = c.Get("user").(*models.User).Uuid.String()
	} else {
		uuidUser = "system"
	}

	res, err := service.CreateCategory(category, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)

}

// PatchCategory godoc
// @Summary Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category id"
// @Param country body dto.UpdateCategoryDTO true "Update category dto"
// @Success 200 {object} models.Category "Category"
// @Router /category/:id [patch]
func PatchCategory(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var category dto.UpdateCategoryDTO

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.UpdateCategory(category, uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, res)

}

// DeleteCategory godoc
// @Summary Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} models.Category "Category"
// @Router /category/:id [delete]
func DeleteCategory(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteCategory(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)

}
