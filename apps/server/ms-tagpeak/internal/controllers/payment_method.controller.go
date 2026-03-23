package controllers

import (
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	_ "ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetPaymentMethod godoc
// @Summary Get PaymentMethod by ID
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param id path string true "PaymentMethod id"
// @Success 200 {object} models.PaymentMethod
// @Router /payment-method/:id [get]
func GetPaymentMethod(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetPaymentMethod(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetAllPaymentMethods godoc
// @Summary Get all PaymentMethods
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Success 200 {array} models.PaymentMethod "Array of PaymentMethods"
// @Router /payment-method [get]
func GetAllPaymentMethods(c echo.Context) error {

	res, err := service.GetAllPaymentMethods()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)

}

// CreatePaymentMethod godoc
// @Summary Create PaymentMethod
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param country body dto.CreatePaymentMethodDTO true "Create PaymentMethod dto"
// @Success 201 {object} models.PaymentMethod "PaymentMethod"
// @Router /payment-method [post]
func CreatePaymentMethod(c echo.Context) error {
	var model dto.CreatePaymentMethodDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.CreatePaymentMethod(model, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)

}

// PatchPaymentMethod godoc
// @Summary Update PaymentMethod
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param id path string true "PaymentMethod id"
// @Param PaymentMethod body dto.UpdatePaymentMethodDTO true "Update PaymentMethod dto"
// @Success 200 {object} models.PaymentMethod "PaymentMethod"
// @Router /payment-method/:id [patch]
func PatchPaymentMethod(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var model dto.UpdatePaymentMethodDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.UpdatePaymentMethod(model, uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, res)

}

// DeletePaymentMethod godoc
// @Summary Delete PaymentMethod
// @Tags PaymentMethod
// @Accept json
// @Produce json
// @Param id path string true "PaymentMethod id"
// @Success 204
// @Router /payment-method/:id [delete]
func DeletePaymentMethod(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeletePaymentMethod(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)

}
