package controllers

import (
	echo "github.com/labstack/echo/v4"
	"ms-changeme/internal/service"
	"net/http"
)

// GetHello godoc
// @Summary Get Hello
// @Description Get Hello
// @Tags Users -> Groups in swagger
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hello, World!"
// @Router /hello [get]
func GetHello(c echo.Context) error {

	example := service.GetExample()
	c.Logger().Infof("This is an example %s \n", example.Name)

	return c.JSON(http.StatusOK, "Hello, World!")
}

// GetHellos godoc
// @Summary Get Hellos
// @Description Get Hellos
// @Tags
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hello, World!"
// @Router /hellos [get]
func GetHellos(c echo.Context) error {

	example := service.GetExample()
	c.Logger().Infof("This is an example %s \n", example.Name)

	return c.JSON(http.StatusOK, "Hello, World!")
}
