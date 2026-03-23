package config

import (
	"ms-cj/pkg/dotenv"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var server *echo.Echo

func InitServer() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	e.Server.ReadTimeout = 20 * time.Second
	e.Server.WriteTimeout = 30 * time.Second

	server = e
	initRoutes()
	port := dotenv.GetEnv("CJ_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

func initRoutes() {
	server.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}
