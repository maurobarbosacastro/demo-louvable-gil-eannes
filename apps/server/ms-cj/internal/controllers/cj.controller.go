package controllers

import (
	"github.com/labstack/echo/v4"
	"ms-cj/internal/dto"
	"ms-cj/internal/services"
	"net/http"
)

func GetCjTransaction(c echo.Context) error {
	cjTransactionDTO := new(dto.CjTransactionDTO)

	// Bind the incoming JSON request body to the cjTransactionDTO
	if err := c.Bind(cjTransactionDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := services.GetCjTransaction(*cjTransactionDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

func GetADIDStore(c echo.Context) error {
	adIdDto := new(dto.AdIdDTO)

	if err := c.Bind(adIdDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := services.GetAdIdStore(*adIdDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
