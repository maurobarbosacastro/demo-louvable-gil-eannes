package constants

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"time"
)

type Keycloak struct {
	Client          *gocloak.GoCloak
	Ctx             context.Context
	AdminToken      *gocloak.JWT
	Realm           string
	TokenExpireDate time.Time
}
