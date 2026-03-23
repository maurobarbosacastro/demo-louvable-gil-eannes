package config

import (
	"fmt"
	"ms-firebase-go/internal/auth"
	"ms-firebase-go/internal/controllers"
	"ms-firebase-go/pkg/dotenv"
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

	authInstance := &auth.Token{}
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
	port := dotenv.GetEnv("FIREBASE_PORT") // Change to MS port env
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

	legacyGroup := server.Group("/legacy")
	messagesGroup := legacyGroup.Group("/api/messages")
	messagesGroup.POST("/token/:token", controllers.SendTokenMessage)
	messagesGroup.POST("/topic/:group", controllers.SendTopicMessage)

	tokenGroup := server.Group("/token")
	tokenGroup.POST("", controllers.SaveToken)
	tokenGroup.DELETE("/:token", controllers.DeleteToken)
	tokenGroup.DELETE("/user/:uuid", controllers.DeleteTokenByUserUuid)

	topicsGroup := server.Group("/topic")
	topicsGroup.GET("", controllers.GetAllTopics)
	topicsGroup.POST("", controllers.CreateTopic)
	topicsGroup.POST("/:uuid/token", controllers.AddTokenToTopic)
	topicsGroup.DELETE("/:uuid/token", controllers.RemoveTokenToTopic)
	topicsGroup.POST("/:uuid/user", controllers.AddUserToTopic)
	topicsGroup.DELETE("/:uuid/user/:user", controllers.RemoveUserFromTopic)

	notification := server.Group("/notification")
	notification.POST("", controllers.CreateNotification)
	notification.PATCH("/:uuid", controllers.UpdateNotification)
	notification.POST("/:uuid/send", controllers.SendNotification)
}
