package auth

import (
	"context"
	"fmt"
	gocloak "github.com/Nerzal/gocloak/v13"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"time"
)

var KeycloakInstance *constants.Keycloak

func InitAuth() {
	logster.StartFuncLog()
	authKeycloak := &constants.Keycloak{}

	authKeycloak.Client = gocloak.NewClient(dotenv.GetEnv("KEYCLOAK_URL"))
	authKeycloak.Ctx = context.Background()
	authKeycloak.Realm = dotenv.GetEnv("KEYCLOAK_REALM")

	logster.Info(fmt.Sprintf("Keycloak connection - user=%s realm=%s",
		dotenv.GetEnv("TAGPEAK_ADMIN_USERNAME"),
		"master",
	))

	token, err := authKeycloak.Client.LoginAdmin(
		authKeycloak.Ctx,
		dotenv.GetEnv("TAGPEAK_ADMIN_USERNAME"),
		dotenv.GetEnv("TAGPEAK_ADMIN_PASSWORD"),
		"master",
	)

	if err != nil {
		logster.Panic(err, "Something wrong with the credentials or url")
	}

	authKeycloak.AdminToken = token
	authKeycloak.TokenExpireDate = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	KeycloakInstance = authKeycloak
	logster.EndFuncLog()
}
