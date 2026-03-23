package auth

import (
	"fmt"
	"ms-firebase-go/pkg/dotenv"
	"net/http"
	"strings"

	gocloak "github.com/Nerzal/gocloak/v13"
	echo "github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type EndpointPublic struct {
	Path   string
	Method string
}

var publicURLs = []*EndpointPublic{}

type Token struct {
	*gocloak.IntroSpectTokenResult
}

func (s *Token) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("c.Request().URL.Path: %s\n", c.Request().URL.Path)
		contained := lo.ContainsBy(publicURLs, func(e *EndpointPublic) bool {
			return e.Path == c.Request().URL.Path && e.Method == c.Request().Method
		})
		if contained || strings.Contains(c.Request().URL.Path, "swagger") || strings.Contains(c.Request().URL.Path, "public") {
			// Skip authentication and move to the next handler
			return next(c)
		}

		headerToken := c.Request().Header.Get("Authorization")
		clientId := dotenv.GetEnv("INTERNAL_KEYCLOAK_CLIENT_ID")
		clientSecret := dotenv.GetEnv("INTERNAL_KEYCLOAK_HOST_SECRET")
		realm := dotenv.GetEnv("KEYCLOAK_REALM")

		rptResult, err := KeycloakInstance.Client.RetrospectToken(KeycloakInstance.Ctx, headerToken, clientId, clientSecret, realm)
		if err != nil {
			c.Error(err)
			c.Logger().Errorf("echo:middleware:auth:Process:Error: %v\n", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Something went wrong with the auth client"})
		}

		/*if !*rptResult.Active {
			c.Logger().Error("echo:middleware:auth:Process: Not Active\n")
			// return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token is not active"})
			// or - but without custom message
			return echo.ErrUnauthorized
		}*/

		s.IntroSpectTokenResult = rptResult
		return next(c)
	}
}
