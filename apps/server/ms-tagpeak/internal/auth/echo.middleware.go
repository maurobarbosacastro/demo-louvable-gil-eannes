package auth

import (
	"fmt"
	gocloak "github.com/Nerzal/gocloak/v13"
	echo "github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

type EndpointPublic struct {
	Path   string
	Method string
}

var publicURLs = []*EndpointPublic{
	{Path: "/auth/", Method: "POST"},
	{Path: "/auth/callback", Method: "GET"},
	{Path: "/auth/reset", Method: "POST"},
	{Path: "/auth/migration", Method: "POST"},
	{Path: "/referral/click", Method: "POST"},
	{Path: "/health", Method: "GET"},
}

type PrincipalStruct struct {
	*gocloak.IntroSpectTokenResult
	Token            *string
	User             *models.User
	MembershipStatus *models.MembershipStatus
}

// Do not use globally, go through echo context
var Principal *PrincipalStruct = &PrincipalStruct{}

func (s *PrincipalStruct) Process(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		contained := lo.ContainsBy(publicURLs, func(e *EndpointPublic) bool {
			return e.Path == c.Request().URL.Path && e.Method == c.Request().Method
		})
		if contained || strings.Contains(c.Request().URL.Path, "swagger") || strings.Contains(c.Request().URL.Path, "public") {
			// Skip authentication and move to the next handler
			return next(c)
		}

		keycloak := KeycloakInstance
		headerToken := removeBearerFromToken(c.Request().Header.Get("Authorization"))
		clientId := dotenv.GetEnv("KEYCLOAK_CLIENT_ID")
		clientSecret := dotenv.GetEnv("KEYCLOAK_CLIENT_SECRET")
		realm := dotenv.GetEnv("KEYCLOAK_REALM")

		rptResult, err := keycloak.Client.RetrospectToken(keycloak.Ctx, headerToken, clientId, clientSecret, realm)
		if err != nil {
			c.Error(err)
			logster.Error(err, "echo:middleware:auth:Process:Error")
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Something went wrong with the auth client"})
		}

		if !*rptResult.Active {
			logster.Error(nil, "echo:middleware:auth:Process: Not Active")
			//return c.JSON(http.StatusUnauthorized, map[string]string{"error": "PrincipalStruct is not active"})
			//or - but without custom message
			return echo.ErrUnauthorized
		}

		s.IntroSpectTokenResult = rptResult

		token := removeBearerFromToken(c.Request().Header.Get("Authorization"))
		s.Token = &token
		basicInfo, err := keycloak.Client.GetUserInfo(keycloak.Ctx, token, keycloak.Realm)

		if err != nil {
			logster.Error(err, "echo:middleware:auth:GetUserInfo")
			return c.JSON(http.StatusInternalServerError, err)
		}

		users, err := keycloak.Client.GetUsers(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, gocloak.GetUsersParams{
			Exact:    gocloak.BoolP(true),
			Username: basicInfo.PreferredUsername,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		user := users[0]
		balance, _ := utils.GetBalance(user)
		country, _ := utils.GetFromAttributes(user, "country")
		currency, _ := utils.GetFromAttributes(user, "currency")
		referralCode, _ := utils.GetFromAttributes(user, "referralCode")
		birthDate, _ := utils.GetFromAttributes(user, "birthDate")
		displayName, _ := utils.GetFromAttributes(user, "displayName")
		isVerified, _ := utils.ConvertToBoolean(user, "isVerified")
		onboardingFinished, _ := utils.ConvertToBoolean(user, "onboardingFinished")
		newsletter, _ := utils.ConvertToBoolean(user, "newsletter")
		source, _ := utils.GetFromAttributes(user, "source")
		currencySelected, _ := utils.ConvertToBoolean(user, "currency_selected")

		profilePictureAtt, _ := utils.GetFromAttributes(user, "profilePicture")
		url := dotenv.GetEnv("MS_IMAGES_SERVER_PUBLIC_URL")
		var profilePicture string
		if profilePictureAtt != "" {
			profilePicture = fmt.Sprintf(url+"%s/profilePicture.webp", profilePictureAtt)
		}

		var transactionPercentage float64
		transactionPercentageAlt, _ := utils.GetFromAttributes(user, "user_percent")
		if transactionPercentageAlt != "" {
			conv, err := strconv.ParseFloat(transactionPercentageAlt, 64)
			if err == nil {
				transactionPercentage = conv
			}
		}

		var rewardPercentage float64
		rewardPercentageAlt, _ := utils.GetFromAttributes(user, "ref_percent")
		if rewardPercentageAlt != "" {
			conv, err := strconv.ParseFloat(rewardPercentageAlt, 64)
			if err == nil {
				rewardPercentage = conv
			}
		}

		groups, _ := keycloak.Client.GetUserGroups(
			keycloak.Ctx,
			keycloak.AdminToken.AccessToken,
			keycloak.Realm,
			gocloak.PString(user.ID),
			gocloak.GetGroupsParams{},
		)

		s.User = &models.User{
			Uuid:               utils.ParseIDToUUID(gocloak.PString(user.ID)),
			Email:              gocloak.PString(user.Username),
			CreatedAt:          gocloak.PInt64(user.CreatedTimestamp),
			FirstName:          gocloak.PString(user.FirstName),
			LastName:           gocloak.PString(user.LastName),
			Balance:            balance,
			Country:            country,
			Currency:           currency,
			ReferralCode:       referralCode,
			BirthDate:          birthDate,
			DisplayName:        displayName,
			Groups:             lo.Map(groups, func(g *gocloak.Group, _ int) string { return gocloak.PString(g.Path) }),
			IsVerified:         isVerified,
			OnboardingFinished: onboardingFinished,
			ProfilePicture:     profilePicture,
			Newsletter:         newsletter,
			Source:             source,
			CurrencySelected:   currencySelected,
		}

		if transactionPercentage > 0 {
			s.User.TransactionPercentage = utils.FloatPointer(transactionPercentage)
		}
		if rewardPercentage > 0 {
			s.User.RewardPercentage = utils.FloatPointer(rewardPercentage)
		}

		s.MembershipStatus = service.GetMembershipStatus(user.ID, s.User.Groups)

		c.Set("user", s.User)
		c.Set("membershipStatus", s.MembershipStatus)
		return next(c)
	}
}

func removeBearerFromToken(token string) string {
	// Check if the token contains "Bearer"
	if strings.HasPrefix(token, "Bearer ") {
		// Split by space and return the second part if available
		parts := strings.Split(token, " ")
		if len(parts) == 2 {
			return parts[1]
		}
	}
	// Return the original token if it does not contain "Bearer "
	return token
}
