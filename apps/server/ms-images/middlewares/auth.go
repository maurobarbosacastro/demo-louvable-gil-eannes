package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"log"
	"ms-images/pkg/dotenv"
	"net/http"
	"strings"
)

type EndpointPublic struct {
	Path   string
	Method string
}

var publicURLs = []*EndpointPublic{
	{Path: "/api/image/:id", Method: "GET"},
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.FullPath() == "/health" {
			c.Next()
			return
		}
		fmt.Printf("c.FullPath: %s\n", c.FullPath())
		contained := lo.ContainsBy(publicURLs, func(e *EndpointPublic) bool {
			return e.Path == c.FullPath() && e.Method == c.Request.Method
		})
		if contained || strings.Contains(c.Request.URL.Path, "swagger") || strings.Contains(c.Request.URL.Path, "public") {
			// Skip authentication and move to the next handler
			c.Next()
			return
		}

		log.Print("Middleware - Auth Start")

		client := gocloak.NewClient(dotenv.GetEnv("KEYCLOAK_URL"))
		realm := dotenv.GetEnv("KEYCLOAK_REALM")
		clientID := dotenv.GetEnv("INTERNAL_KEYCLOAK_CLIENT_ID")
		clientSecret := dotenv.GetEnv("INTERNAL_KEYCLOAK_HOST_SECRET")

		fmt.Printf("keycloakURL: %s\n", dotenv.GetEnv("KEYCLOAK_URL"))
		fmt.Printf("clientID: %s\n", clientID)
		fmt.Printf("realm: %s\n", realm)

		ctx := context.Background()
		accessToken := c.Request.Header.Get("Authorization")

		if strings.TrimSpace(accessToken) == "" {
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("token Missing"))
			return
		}

		authHeaderParts := strings.Split(accessToken, " ")
		rptResult, err := client.RetrospectToken(ctx, authHeaderParts[1], clientID, clientSecret, realm)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("could not verify token"))
			return
		}

		if !(*rptResult.Active) {
			_ = c.AbortWithError(http.StatusUnauthorized, errors.New("token is not active"))
			return
		}

		c.Next()
		log.Print("Middleware - Auth End")
	}
}
