package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ms-images/controllers"
	"ms-images/middlewares"
	"ms-images/pkg/dotenv"
	"net/http"
	"time"
)

func NewRouter() *gin.Engine {

	router := gin.New()
	gin.SetMode(dotenv.GetEnv("SERVER.MODE"))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "content-type", "accept", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	//Private routes
	privateGroup := router.Group("/api/image", middlewares.Auth())
	{
		privateGroup.POST("/base64", controllers.HandleJsonRequest)
		privateGroup.POST("/", controllers.CreateImage)
		privateGroup.GET("/:id", func(c *gin.Context) {
			file := controllers.GetFile(c.Param("id"))
			c.IndentedJSON(http.StatusOK, file)
		})
		privateGroup.GET("/:id/free-transform", controllers.TransformImage)
		privateGroup.DELETE("/:id", func(c *gin.Context) {
			err := controllers.DeleteFile(c.Param("id"))

			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{
					"error": fmt.Sprintf("Error deleting image with id: %s", c.Param("id")),
				})
			} else {
				c.IndentedJSON(http.StatusNoContent, nil)
			}

		})
	}

	auxPrivateGroup := router.Group("/aux", middlewares.Auth())
	{
		auxPrivateGroup.GET("/ai-logo", controllers.GetLogo)
		auxPrivateGroup.POST("/url", controllers.GetImageFromUrl)
	}

	router.Static("/images", dotenv.GetEnv("SERVER_IMAGES"))

	return router

}
