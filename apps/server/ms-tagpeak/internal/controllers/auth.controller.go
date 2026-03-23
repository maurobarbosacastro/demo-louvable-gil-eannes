package controllers

import (
	"fmt"
	"ms-tagpeak/external/images"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateUser godoc
// @Summary CreateUser
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDto true "User info"
// @Success 201 {object} models.User "User created successfully"
// @Router /auth [post]
func CreateUser(c echo.Context) error {

	bodyUserDTO := new(dto.CreateUserDto)

	// Bind the incoming JSON request body to the UserDTO
	if err := c.Bind(bodyUserDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Validate the input (optional, requires validation setup)
	if err := c.Validate(bodyUserDTO); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	keycloak := auth.KeycloakInstance

	user, err := service.CreateUser(*bodyUserDTO, keycloak)

	if err != nil {
		if apiError, ok := err.(*gocloak.APIError); ok {
			switch apiError.Code {
			case http.StatusConflict:
				fmt.Printf("Error creating user: %v\n", apiError.Message)
				return c.JSON(http.StatusConflict, apiError.Message)
			default:
				return c.JSON(apiError.Code, apiError.Message)
			}
		} else {
			// Generic error handling
			return c.JSON(http.StatusInternalServerError, err)
		}

	}

	return c.JSON(http.StatusOK, user)
}

// CreateUserFromMigration godoc
// @Summary CreateUser
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDto true "User info"
// @Success 201 {object} models.User "User created successfully"
// @Router /auth [post]
func CreateUserFromMigration(c echo.Context) error {

	bodyUserDTO := new(dto.CreateUserFromMigrationDto)

	// Bind the incoming JSON request body to the UserDTO
	if err := c.Bind(bodyUserDTO); err != nil {
		fmt.Printf("Error Binding the body %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Validate the input (optional, requires validation setup)
	if err := c.Validate(bodyUserDTO); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	keycloak := auth.KeycloakInstance

	if *bodyUserDTO.ProfilePic != "NULL" {
		profilePic, _ := images.GetProfilePictureFromUrl(*bodyUserDTO.ProfilePic, bodyUserDTO.FirstName+"_profilePic")
		bodyUserDTO.ProfilePic = &profilePic.Id
	}

	user, err := service.CreateUserFromMigration(*bodyUserDTO, keycloak)

	if err != nil {
		if apiError, ok := err.(*gocloak.APIError); ok {
			switch apiError.Code {
			case http.StatusConflict:
				fmt.Printf("Error creating user: %v\n", apiError.Message)
				return c.JSON(http.StatusConflict, apiError.Message)
			default:
				return c.JSON(apiError.Code, apiError.Message)
			}
		} else {
			// Generic error handling
			return c.JSON(http.StatusInternalServerError, err)
		}

	}

	return c.JSON(http.StatusOK, user)
}

// GetOwnUser godoc
// @Summary GetOwnUser
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.User "Own user"
// @Router /auth/me [get]
func GetOwnUser(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Get("user").(*models.User))
}

// GetUserById godoc
// @Summary GetOwnUser
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} response_object.UserDto "Own user"
// @Router /auth/:id [get]
func GetUserById(c echo.Context) error {
	id := c.Param("id")

	keycloak := auth.KeycloakInstance

	user, err := service.GetUserById(id, keycloak)

	if err != nil {
		if apiError, ok := err.(*gocloak.APIError); ok {
			switch apiError.Code {
			case http.StatusConflict:
				fmt.Printf("Error creating user: %v\n", apiError.Message)
				return c.JSON(http.StatusConflict, apiError.Message)
			default:
				return c.JSON(apiError.Code, apiError.Message)
			}
		} else {
			// Generic error handling
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	userDto := service.GetInfoUserReferral(user, keycloak)

	return c.JSON(http.StatusOK, userDto)
}

// UpdateUser godoc
// @Summary UpdateUser
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Param user body dto.UpdateUserDto true "Update user dto"
// @Success 200 {object} response_object.UserDto "Own user"
// @Router /auth/:id [patch]
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	uuid := utils.ParseIDToUUID(id)

	if c.Get("user").(*models.User).Uuid != uuid {
		return c.JSON(http.StatusUnauthorized, "Cannot update other user information")
	}

	user := dto.UpdateUserDto{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := c.Validate(user); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	keycloak := auth.KeycloakInstance

	updateUser, err := service.UpdateUser(uuid, user, keycloak)

	if err != nil {
		if apiError, ok := err.(*gocloak.APIError); ok {
			switch apiError.Code {
			case http.StatusConflict:
				fmt.Printf("Error creating user: %v\n", apiError.Message)
				return c.JSON(http.StatusConflict, apiError.Message)
			default:
				return c.JSON(apiError.Code, apiError.Message)
			}
		} else {
			// Generic error handling
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	userDto := service.GetInfoUserReferral(updateUser, keycloak)

	return c.JSON(http.StatusOK, userDto)
}

// ResetPassword godoc
// @Summary ResetPassword
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.ResetPasswordDto true "Reset password dto"
// @Success 201
// @Failure 404
// @Router /auth/reset [post]
func ResetPassword(c echo.Context) error {
	dtoResetPassword := dto.ResetPasswordDto{}

	if err := c.Bind(&dtoResetPassword); err != nil {
		return c.JSON(http.StatusBadRequest, utils.CustomErrorStruct{}.BadRequestError("Reset Password Method"))
	}

	if err := c.Validate(dtoResetPassword); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	keycloak := auth.KeycloakInstance

	user, err := service.FindUserByEmail(dtoResetPassword.Email, keycloak)

	if err != nil || user.ID == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	err = service.ResetPassword(user.ID, keycloak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func EmailVerified(c echo.Context) error {
	bodyEmailVerifiedDTO := new(dto.EmailVerifiedDto)

	// Bind the incoming JSON request body to the UserDTO
	if err := c.Bind(bodyEmailVerifiedDTO); err != nil {
		return c.JSON(http.StatusBadRequest, utils.CustomErrorStruct{}.BadRequestError("Email Verified Method"))
	}

	// Validate the input (optional, requires validation setup)
	if err := c.Validate(bodyEmailVerifiedDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	keycloak := auth.KeycloakInstance

	val, err := service.EmailVerified(*bodyEmailVerifiedDTO, keycloak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, val)
}

// SendVerificationEmail godoc
// @Summary SendVerificationEmail
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} string "Email sent successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/send-verification-email [post]
func SendVerificationEmail(c echo.Context) error {
	logster.StartFuncLog()

	// Get authenticated user from context
	user := c.Get("user").(*models.User)

	keycloak := auth.KeycloakInstance

	// Send verification email via Keycloak
	err := keycloak.Client.SendVerifyEmail(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		user.Uuid.String(),
		keycloak.Realm,
	)

	if err != nil {
		logster.Error(err, "Error sending verification email")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send verification email"})
	}

	logster.Info(fmt.Sprintf("Verification email sent successfully to user: %s", user.Uuid.String()))
	logster.EndFuncLog()
	return c.JSON(http.StatusOK, map[string]string{"message": "Verification email sent successfully"})
}

// CheckEmailVerificationStatus godoc
// @Summary CheckEmailVerificationStatus
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]bool "Email verification status"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/email-verification-status [get]
func CheckEmailVerificationStatus(c echo.Context) error {
	logster.StartFuncLog()

	// Get authenticated user from context
	user := c.Get("user").(*models.User)
	userID := user.Uuid.String()

	keycloak := auth.KeycloakInstance

	// Get user from Keycloak to check built-in EmailVerified field
	keycloakUser, err := keycloak.Client.GetUserByID(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		keycloak.Realm,
		userID,
	)

	if err != nil {
		logster.Error(err, "Error getting user from Keycloak")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check email verification status"})
	}

	// Check built-in EmailVerified field
	isVerified := false
	if keycloakUser.EmailVerified != nil {
		isVerified = *keycloakUser.EmailVerified
	}

	// Get current custom isVerified attribute
	customIsVerified := false
	if user.IsVerified {
		customIsVerified = user.IsVerified
	}

	// If built-in is verified but custom is not, update the custom attribute
	if isVerified && !customIsVerified {
		updateDto := dto.UpdateUserDto{
			IsVerified: func() *string { s := "true"; return &s }(),
		}
		_, err := service.UpdateUser(user.Uuid, updateDto, keycloak)
		if err != nil {
			logster.Error(err, "Error updating custom isVerified attribute")
			// Continue to return the status even if update fails
		} else {
			logster.Info(fmt.Sprintf("Updated custom isVerified attribute for user: %s", userID))
		}
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, map[string]bool{"isVerified": isVerified})
}

// GetAllReferralByUserUuid godoc
// @Summary GetAllReferralByUserUuid
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Param pagination query pagination.PaginationParams true "Pagination params"
// @Success 200 {array} pagination.PaginationResult{data=[]dto.ReferralDTO} "Array of Referrals"
// @Router /auth/:id/referral [get]
func GetAllReferralByUserUuid(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	var pag pagination.PaginationParams

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return c.JSON(http.StatusInternalServerError, errPag)
	}

	referrals, err := service.GetAllReferralByUserUuid(pag, uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, referrals)
}

// GetAllUsers godoc
// @Summary GetAllUsers
// @Tags Auth
// @Accept json
// @Produce json
// @Param searchUser query string true "email or name"
// @Success 200 {array} models.User "List of all users"
// @Router /auth/users [get]
func GetAllUsers(c echo.Context) error {
	logster.StartFuncLog()

	searchUser := c.QueryParam("searchUser")
	limitUser, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		logster.Info("Limit not provided, defaulting to 10")
		limitUser = 10
	}

	keycloak := auth.KeycloakInstance
	users, err := service.GetUsers(keycloak, &searchUser, limitUser)

	sort.Slice(users, func(i, j int) bool {
		if strings.ToLower(users[i].FirstName) == strings.ToLower(users[j].FirstName) {
			return strings.ToLower(users[i].LastName) < strings.ToLower(users[j].LastName)
		}
		return strings.ToLower(users[i].FirstName) < strings.ToLower(users[j].FirstName)
	})

	if err != nil {
		logster.Error(err, "GetUsers")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, users)
}

// UploadUserProfilePicture godoc
// @Summary Upload user profile picture
// @Description Upload user profile picture
// @Tags Auth
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string true "User ID"
// @Param file formData file true "Logo file"
// @Success 200 {string} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/:id/profile-picture [post]
func UploadUserProfilePicture(c echo.Context) error {
	keycloak := auth.KeycloakInstance

	// Get the uploaded file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No file uploaded"})
	}

	fileUuid, errUpload := images.UploadImage(file, c.FormValue("fileName"), "profilePicture")
	if errUpload != nil {
		return c.JSON(http.StatusBadRequest, errUpload)
	}

	userUuid := utils.ParseIDToUUID(c.Param("id"))

	err = service.UpdateUseProfilePicture(userUuid, fileUuid, keycloak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	url := dotenv.GetEnv("MS_IMAGES_SERVER_PUBLIC_URL")

	return c.String(http.StatusOK, fmt.Sprintf(url+"%s/profilePicture.webp", *fileUuid))
}

// DeleteUserProfilePicture godoc
// @Summary Delete user profile picture
// @Description Delete user profile picture
// @Tags User
// @Produce  json
// @Param id path string true "User ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/:id/profile-picture/:pictureUuid [delete]
func DeleteUserProfilePicture(c echo.Context) error {
	userUuid := utils.ParseIDToUUID(c.Param("id"))
	fileUuid := utils.ParseIDToUUID(c.Param("pictureUuid"))
	keycloak := auth.KeycloakInstance

	err := images.DeleteImage(fileUuid.String())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = service.UpdateUseProfilePicture(userUuid, nil, keycloak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

// ValidateUserAction godoc
// @Summary Validate action
// @Description Validate if user has the required action to "UPDATE_PASSWORD" and communicate with the Camunda process
// @Tags Camunda
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /public/auth/validate-action [get]
func ValidateUserAction(c echo.Context) error {
	userEmail := c.QueryParam("email")
	keycloak := auth.KeycloakInstance

	userKeycloak, err := service.FindUserByEmail(userEmail, keycloak)
	if err != nil || userKeycloak.ID == nil {
		return c.JSON(http.StatusUnauthorized, utils.CustomErrorStruct{
			ErrorType: "invalid_grant",
		})
	}

	err = service.ValidateUserAction(userKeycloak, keycloak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

// GetReferralsInfoByUserUuid godoc
// @Summary GetReferralsInfoByUserUuid
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Success 200 {array} response_object.ReferralInfo
// @Router /auth/:id/referral/clicks [get]
func GetReferralsInfoByUserUuid(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	keycloak := auth.KeycloakInstance

	referrals, err := service.GetReferralsInfoByUserUuid(uuid, keycloak)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, referrals)
}

// GetReferralsRevenueInfoByUserUuid godoc
// @Summary GetReferralsRevenueInfoByUserUuid
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Success 200 {array} response_object.ReferralRevenueInfo
// @Router /auth/:id/referral/revenue [get]
func GetReferralsRevenueInfoByUserUuid(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	referrals, err := service.GetRevenuesInfoByUserUuid(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, referrals)
}

// GetRevenuesInfoByUserUuid godoc
// @Summary GetRevenuesInfoByUserUuid
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Success 200 {array} response_object.UserReferralRevenueInfoDto
// @Router /auth/:id/referral/revenue/info [get]
func GetRevenuesInfoByUserUuid(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	referrals, err := service.GetRevenuesInfoByUserUuid(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, referrals)
}

// GetUsersReferralsRevenueInfoByReferrerUuid godoc
// @Summary GetUsersReferralsRevenueInfoByReferrerUuid
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Success 200 {array} response_object.UserReferralRevenueInfoDto
// @Router /auth/:id/referral/revenue/user-info [get]
func GetUsersReferralsRevenueInfoByReferrerUuid(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	referrals, err := service.GetUsersReferralsRevenueInfoByReferrerUuid(uuid, auth.KeycloakInstance)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, referrals)
}

// GetAllReferralRevenueByReferrerUuid godoc
// @Summary GetAllReferralRevenueByReferrerUuid
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Success 200 {float64} float64
// @Router /auth/:id/referral/revenue/total-revenue [get]
func GetAllReferralRevenueByReferrerUuid(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	referrals, err := service.GetAllReferralRevenueByReferrerUuid(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, referrals)
}

// GetUserStats godoc
// @Summary Get User Stats
// @Description Get User Stats
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} response_object.UserStatsRO
// @Router /auth/:id/stats [get]
func GetUserStats(c echo.Context) error {
	statsRO := response_object.UserStatsRO{
		Level:      *(c.Get("membershipStatus").(*models.MembershipStatus)).Level,
		ValueSpent: service.GetAmountUserTransactions(c.Get("user").(*models.User).Uuid),
	}

	return c.JSON(http.StatusOK, statsRO)
}

func SetMigratedUserBalance(c echo.Context) error {

	list, err := service.GetUsersWithTotalAmountReward()
	if err != nil {
		return err
	}

	keycloak := auth.KeycloakInstance
	updated := make([]response_object.UsersWithTotalAmountReward, 0)
	errored := make([]response_object.UsersWithTotalAmountReward, 0)

	for _, user := range list {
		formatedBalance := strconv.FormatFloat(user.Total, 'f', 2, 64)
		dtoUpdate := dto.UpdateUserDto{
			Balance: &formatedBalance,
		}
		updateUser, err := service.UpdateUser(user.User, dtoUpdate, keycloak)
		if err != nil {
			errored = append(errored, user)
		}

		if updateUser != nil {
			updated = append(updated, user)
		}
	}

	return c.JSON(http.StatusOK, response_object.UsersWithTotalAmountRewardFinal{Updated: updated, Error: errored})
}

func CheckUserState(c echo.Context) error {
	principal := c.Get("user").(*models.User)

	today := time.Now()
	to := today.AddDate(0, 0, 1)
	from := today.AddDate(0, 0, -10)

	fmt.Printf("today: %s, from: %s\n", to.Format("2006-01-02"), from.Format("2006-01-02"))

	events, err := getEvents(gocloak.GetEventsParams{
		UserID:   utils.StringPointer(principal.Uuid.String()),
		Client:   utils.StringPointer("tagpeak-client"),
		DateFrom: utils.StringPointer(from.Format("2006-01-02")),
		DateTo:   utils.StringPointer(to.Format("2006-01-02")),
	}, []string{"LOGIN"})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, events)
}

func getEvents(params gocloak.GetEventsParams, paramTypes []string) (res []*gocloak.EventRepresentation, err error) {
	keycloak := auth.KeycloakInstance

	if len(paramTypes) > 0 {
		keycloak.Client.RestyClient().OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
			r.SetQueryParamsFromValues(url.Values{
				"type": paramTypes,
			})
			return nil
		})
	}

	return keycloak.Client.GetEvents(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, params)
}

func SetCurrencyUser(c echo.Context) error {
	logster.StartFuncLog()

	userUuid := uuid.MustParse(c.Param("id"))
	bodyUserDTO := new(dto.CurrencySetDto)
	keycloak := auth.KeycloakInstance

	// Bind the incoming JSON request body to the UserDTO
	if err := c.Bind(bodyUserDTO); err != nil {
		fmt.Printf("Error Binding the body %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updateDto := dto.UpdateUserDto{
		CurrencySelected: utils.BoolPointer(true),
		Currency:         utils.StringPointer(bodyUserDTO.CurrencyCode),
	}

	_, err := service.UpdateUser(userUuid, updateDto, keycloak)

	if err != nil {
		logster.Error(err, "Error Updating user")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	//update and recalculate transaction to correct user currency
	service.UpdateUserTransactionAndRewardsByCurrency(userUuid.String(), bodyUserDTO.CurrencyCode)

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, nil)
}

func AuthCallback(c echo.Context) error {
	logster.StartFuncLog()
	q := c.QueryParams()
	logster.Info(fmt.Sprintf("QueryParams: %+v", q))

	env := dotenv.GetEnv("ENV")

	redirectURL := "/#/auth/callback?" + q.Encode()
	if strings.Contains(strings.ToLower(env), "dev") {
		redirectURL = "http://localhost:4200" + redirectURL
	} else {
		redirectURL = "https://" + c.Request().Host + redirectURL
	}

	logster.EndFuncLogMsg(redirectURL)
	return c.Redirect(http.StatusFound, redirectURL)
}

func SocialsFinishProfile(c echo.Context) error {
	logster.StartFuncLog()

	bodyUserDTO := new(dto.SocialProfileFinish)

	// Bind the incoming JSON request body to the UserDTO
	if err := c.Bind(bodyUserDTO); err != nil {
		logster.Error(err, "Error Binding the body")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Validate the input (optional, requires validation setup)
	if err := c.Validate(bodyUserDTO); err != nil {
		logster.Error(err, "Error validating the body")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, err)
	}

	keycloak := auth.KeycloakInstance
	loggedUser := c.Get("user").(*models.User)

	user, err := service.SocialsFinishProfile(*loggedUser, bodyUserDTO, keycloak)

	if err != nil {
		logster.Error(err, "Error finishing social auth")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, user)
}
