package auth

import (
	echo "github.com/labstack/echo/v4"
	"ms-shopify/pkg/dotenv"
	"net/http"
	"time"
)

type KeycloakMiddleware struct {
}

func (s *KeycloakMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		keycloak := KeycloakInstance
		now := time.Now()

		if now.After(keycloak.TokenExpireDate) {
			token, err := keycloak.Client.LoginAdmin(
				keycloak.Ctx,
				dotenv.GetEnv("TAGPEAK_ADMIN_USERNAME"),
				dotenv.GetEnv("TAGPEAK_ADMIN_PASSWORD"),
				"master",
			)

			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			KeycloakInstance.AdminToken = token
			KeycloakInstance.TokenExpireDate = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)

		}

		return next(c)
	}
}
