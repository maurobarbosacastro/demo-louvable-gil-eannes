package config

import (
	"fmt"
	"ms-interactive-brokers/internal/controllers"
	"ms-interactive-brokers/pkg/dotenv"
	"net/http"
	"strings"
	"time"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var server *echo.Echo

func InitServer() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.BodyLimit(dotenv.GetEnv("BODYLIMIT")))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	// internalKeycloakMiddleware := &auth.InternalKeycloak{}
	// e.Pre(internalKeycloakMiddleware.Process)

	fmt.Printf("\n\nMiddlewares: CORST, BodyLimit - %s, Logger, Recover, Secure (XSS), Custome auth keycloak", dotenv.GetEnv("BODYLIMIT"))

	e.Debug = dotenv.GetEnv("ENV") == "dev"
	fmt.Printf("\n - Debug mode: %t", e.Debug)

	e.Logger.SetLevel(getLogLevel())
	fmt.Printf("\n - Log level: %s", dotenv.GetEnv("LOGLEVEL"))

	e.Server.ReadTimeout = 20 * time.Second
	e.Server.WriteTimeout = 1 * time.Minute

	server = e
	initRoutes()
	port := dotenv.GetEnv("INTERACTIVE_BROKERS_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

func getLogLevel() log.Lvl {
	logLevel := dotenv.GetEnv("LOGLEVEL") // DEBUG | INFO | WARN | ERROR | OFF
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	case "OFF":
		return log.OFF
	default:
		return log.INFO // Default log level
	}
}

/**
 * Define all API routes here
 */
func initRoutes() {
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	server.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	server.POST("/lastPrice", controllers.GetLastPriceBulk)
}
