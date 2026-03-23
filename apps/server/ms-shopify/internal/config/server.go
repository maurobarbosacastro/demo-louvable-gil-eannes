package config

import (
	"fmt"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"ms-shopify/internal/auth"
	"ms-shopify/internal/controllers"
	"ms-shopify/pkg/dotenv"
	"net/http"
	"strings"
	"time"
)

var server *echo.Echo

func InitServer() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.BodyLimit(dotenv.GetEnv("BODYLIMIT")))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	authInstance := &auth.PrincipalStruct{}
	e.Use(authInstance.Process)

	fmt.Printf("\n\nMiddlewares: CORST, BodyLimit - %s, Logger, Recover, Secure (XSS), Custome auth keycloak", dotenv.GetEnv("BODYLIMIT"))

	e.Debug = dotenv.GetEnv("ENV") == "dev"
	fmt.Printf("\n - Debug mode: %t", e.Debug)

	e.Logger.SetLevel(getLogLevel())
	fmt.Printf("\n - Log level: %s", dotenv.GetEnv("LOGLEVEL"))

	e.Server.ReadTimeout = 20 * time.Second
	e.Server.WriteTimeout = 30 * time.Second

	server = e
	initRoutes()
	port := dotenv.GetEnv("SHOPIFY_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

func getLogLevel() log.Lvl {
	logLevel := dotenv.GetEnv("LOGLEVEL") //DEBUG | INFO | WARN | ERROR | OFF
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

	server.GET("/shop", controllers.GetShopByUrl)
	server.POST("/shop", controllers.CreateShop)
	server.PATCH("/shop/:id", controllers.UpdateShop)
	server.POST("/shop/:id/activate", controllers.ActivateShop)
	server.POST("/shop/:id/deactivate", controllers.DeactivateShop)
	server.POST("/install", controllers.SetupShopify)

	publicGroup := server.Group("/public")
	publicGroup.POST("/webhook", controllers.HandleWebhook)
	publicGroup.POST("/webhook-uninstall", controllers.HandleUninstallWebhook)
}
