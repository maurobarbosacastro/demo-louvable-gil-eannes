package auth

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/labstack/echo/v4"
	"ms-interactive-brokers/pkg/dotenv"
	"ms-interactive-brokers/pkg/logster"
	"net/http"
)

type InternalKeycloak struct {
	*gocloak.IntroSpectTokenResult
	Token *string
}

func (s *InternalKeycloak) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headerToken := c.Request().Header.Get("Authorization")
		clientId := dotenv.GetEnv("INTERNAL_KEYCLOAK_CLIENT_ID")
		clientSecret := dotenv.GetEnv("INTERNAL_KEYCLOAK_CLIENT_SECRET")
		realm := dotenv.GetEnv("KEYCLOAK_REALM")

		rptResult, err := KeycloakInstance.Client.RetrospectToken(KeycloakInstance.Ctx, headerToken, clientId, clientSecret, realm)
		if err != nil {
			c.Error(err)
			logster.Error(err, "echo:middleware:auth:Process:Error")
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Something went wrong with the auth client"})
		}

		if !*rptResult.Active {
			logster.Error(nil, "echo:middleware:auth:Process: Not Active")
			//return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token is not active"})
			//or - but without custom message
			return echo.ErrUnauthorized
		}

		s.IntroSpectTokenResult = rptResult
		return next(c)
	}
}
