package keycloak_utils

import (
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"time"
)

func RefreshKeycloakAdminToken(keycloak *constants.Keycloak) {
	logster.StartFuncLog()

	now := time.Now()

	if now.After(keycloak.TokenExpireDate) {
		logster.Info("Admin token expired, generating new one")
		token, err := keycloak.Client.LoginAdmin(
			keycloak.Ctx,
			dotenv.GetEnv("TAGPEAK_ADMIN_USERNAME"),
			dotenv.GetEnv("TAGPEAK_ADMIN_PASSWORD"),
			"master",
		)

		if err != nil {
			logster.Error(err, "Error getting token")
		}

		keycloak.AdminToken = token
		keycloak.TokenExpireDate = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}

	logster.EndFuncLog()
}
