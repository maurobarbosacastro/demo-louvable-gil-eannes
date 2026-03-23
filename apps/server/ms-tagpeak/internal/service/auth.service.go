package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
	"github.com/samber/lo"
	interactive_brokers "ms-tagpeak/external/interactive-brokers"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/camunda"
	keycloak_utils "ms-tagpeak/pkg/keycloak"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"strconv"
	"strings"
)

func CreateUser(bodyUserDTO dto.CreateUserDto, keycloak *constants.Keycloak) (createdUser *models.User, err error) {
	logster.StartFuncLogMsg(fmt.Sprintf("body user dto: %+v", bodyUserDTO))

	password, errPass := base64.StdEncoding.DecodeString(bodyUserDTO.Password)

	if errPass != nil {
		logster.Error(errPass, "Error decoding password")
		logster.EndFuncLog()
		return nil, errPass
	}

	attributes := &map[string][]string{
		"balance":           {"0"},
		"referralCode":      {utils.RandomWordsCode(7)},
		"currency":          {bodyUserDTO.Currency},
		"newsletter":        {"true"},
		"displayName":       {utils.GenerateDisplayName()},
		"currency_selected": {"true"},
	}

	if bodyUserDTO.Country != nil {
		(*attributes)["country"] = []string{*bodyUserDTO.Country}
	}

	if bodyUserDTO.UtmParams != nil {
		(*attributes)["utm"] = []string{*bodyUserDTO.UtmParams}
	}

	groups := []string{*Configuration.MembershipLevels.Base, *Configuration.UserTypes.Default}
	requiredActions := []string{"VERIFY_EMAIL"}
	emailVerified := gocloak.BoolP(false)

	if bodyUserDTO.IsShop != nil && *bodyUserDTO.IsShop {
		groups = []string{*Configuration.UserTypes.Shop}
		requiredActions = []string{}
	}

	user := gocloak.User{
		Enabled:         gocloak.BoolP(true),
		EmailVerified:   emailVerified,
		FirstName:       gocloak.StringP(bodyUserDTO.FirstName),
		LastName:        gocloak.StringP(bodyUserDTO.LastName),
		Email:           gocloak.StringP(bodyUserDTO.Email),
		Groups:          &groups,
		Attributes:      attributes,
		RequiredActions: &requiredActions,
	}

	userId, errCreation := keycloak.Client.CreateUser(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		keycloak.Realm,
		user,
	)

	if errCreation != nil {
		logster.Error(errCreation, "Error creating user")
		logster.EndFuncLog()
		return nil, errCreation
	}

	errSetPassword := keycloak.Client.SetPassword(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		userId,
		keycloak.Realm,
		string(password),
		false,
	)

	if errSetPassword != nil {
		logster.Error(errSetPassword, "Error setting password")
		logster.EndFuncLog()
		return nil, errSetPassword
	}

	errEmail := keycloak.Client.SendVerifyEmail(keycloak.Ctx, keycloak.AdminToken.AccessToken, userId, keycloak.Realm)

	if errEmail != nil {
		logster.Error(errEmail, "Error sending email")
		logster.EndFuncLog()
		return nil, errEmail
	}

	userResponse, _ := GetUserById(userId, keycloak)

	if bodyUserDTO.ReferralCode != nil {
		errReferral := handleReferral(
			*bodyUserDTO.ReferralCode,
			bodyUserDTO.ReferralClick,
			keycloak,
			userResponse.Uuid,
		)

		if errReferral != nil {
			return userResponse, errReferral
		}
	}

	userVar := map[string]interface{}{
		"userUuid": userId,
	}

	// Start camunda registration process
	process := camunda.StartProcessInstance(camunda.InjectEnvOnKey(camunda.RegistrationFlow), *camunda.GetCamundaClient(), userVar)

	createUserCamundaProcessDto := dto.CreateCamundaProcessDto{
		FieldUUID:          userResponse.Uuid,
		ProcessInstanceKey: process.ProcessInstanceKey,
		ProcessId:          camunda.RegistrationFlow,
	}

	_, err = CreateCamundaProcess(createUserCamundaProcessDto)

	if err != nil {
		logster.Error(err, "Error creating camunda process")
		logster.EndFuncLog()
	}

	logster.EndFuncLogMsg(fmt.Sprintf("User created: %v", userResponse.Uuid))
	return userResponse, nil
}

func CreateUserFromMigration(bodyUserDTO dto.CreateUserFromMigrationDto, keycloak *constants.Keycloak) (*models.User, error) {
	logster.StartFuncLog()

	refCode := utils.RandomWordsCode(7)
	if bodyUserDTO.RefCode != "NULL" && bodyUserDTO.RefCode != "" && &bodyUserDTO.RefCode != nil {
		refCode = bodyUserDTO.RefCode
	}

	attributes := &map[string][]string{
		"balance":            {"0"},
		"referralCode":       {refCode},
		"currency":           {bodyUserDTO.Currency},
		"isVerified":         {"true"},
		"onboardingFinished": {"false"},
		"legacyId":           {strconv.FormatInt(bodyUserDTO.LegacyId, 10)},
		"newsletter":         {strings.ToLower(bodyUserDTO.Newsletter)},
	}

	if bodyUserDTO.Country != "N/A" && bodyUserDTO.Country != "" && &bodyUserDTO.Country != nil {
		(*attributes)["country"] = []string{bodyUserDTO.Country}
	}

	if bodyUserDTO.DisplayName != "NULL" && bodyUserDTO.DisplayName != "" && &bodyUserDTO.DisplayName != nil {
		(*attributes)["displayName"] = []string{bodyUserDTO.DisplayName}
	} else {
		(*attributes)["displayName"] = []string{utils.GenerateDisplayName()}
	}

	if *bodyUserDTO.BirthDate != "NULL" && *bodyUserDTO.BirthDate != "" && bodyUserDTO.BirthDate != nil {
		(*attributes)["birthDate"] = []string{*bodyUserDTO.BirthDate}
	}

	if bodyUserDTO.ProfilePic != nil && *bodyUserDTO.ProfilePic != "NULL" && *bodyUserDTO.ProfilePic != "" {
		(*attributes)["profilePicture"] = []string{*bodyUserDTO.ProfilePic}
	}

	if bodyUserDTO.UserPercent != "NULL" && bodyUserDTO.UserPercent != "" && &bodyUserDTO.UserPercent != nil {
		(*attributes)["user_percent"] = []string{bodyUserDTO.UserPercent}
	}

	if bodyUserDTO.RefPercent != "NULL" && bodyUserDTO.RefPercent != "" && &bodyUserDTO.RefPercent != nil {
		(*attributes)["ref_percent"] = []string{bodyUserDTO.RefPercent}
	}

	var groups []string
	if bodyUserDTO.BadgeType == "auto" {
		groups = []string{*Configuration.MembershipLevels.Base, *Configuration.UserTypes.Default}
	} else if bodyUserDTO.BadgeType == "influencer" {
		groups = []string{*Configuration.MembershipLevels.Influencer, *Configuration.UserTypes.Default}
	}

	enabled := gocloak.BoolP(true)
	emailVerified := gocloak.BoolP(true)
	firstName := gocloak.StringP(bodyUserDTO.FirstName)
	lastName := gocloak.StringP(bodyUserDTO.LastName)
	email := gocloak.StringP(bodyUserDTO.Email)
	var requiredActions []string
	if bodyUserDTO.IdentifyProvider == nil || *bodyUserDTO.IdentifyProvider == "" {
		requiredActions = []string{"UPDATE_PASSWORD"}
	}

	user := gocloak.User{
		Enabled:         enabled,
		EmailVerified:   emailVerified,
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		Groups:          &groups,
		Attributes:      attributes,
		RequiredActions: &requiredActions,
	}

	userId, errCreation := keycloak.Client.CreateUser(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		keycloak.Realm,
		user,
	)
	if errCreation != nil {
		return nil, errCreation
	}

	userResponse, _ := GetUserById(userId, keycloak)

	err := CreateUserMigration(models.UserMigration{
		UserUuid: userResponse.Uuid,
		LegacyId: bodyUserDTO.LegacyId,
	})
	if err != nil {
		return nil, err
	}

	if bodyUserDTO.IdentifyProvider == nil || *bodyUserDTO.IdentifyProvider == "" {
		err = keycloak.Client.CreateUserFederatedIdentity(
			keycloak.Ctx,
			keycloak.AdminToken.AccessToken,
			keycloak.Realm,
			userId,
			*bodyUserDTO.IdentifyProvider,
			gocloak.FederatedIdentityRepresentation{
				UserID:           bodyUserDTO.FederatedUserId,
				IdentityProvider: bodyUserDTO.IdentifyProvider,
				UserName:         &userResponse.Email,
			},
		)
		if err != nil {
			return nil, err
		}
	}

	userVar := map[string]interface{}{
		"userUuid": userId,
	}

	// Start camunda registration process
	process := camunda.StartProcessInstance(camunda.InjectEnvOnKey(camunda.UserMigrationFlow), *camunda.GetCamundaClient(), userVar)

	createUserCamundaProcessDto := dto.CreateCamundaProcessDto{
		FieldUUID:          userResponse.Uuid,
		ProcessInstanceKey: process.ProcessInstanceKey,
		ProcessId:          camunda.UserMigrationFlow,
	}

	_, err = CreateCamundaProcess(createUserCamundaProcessDto)

	logster.EndFuncLog()
	return userResponse, nil
}

func GetUserById(id string, keycloak *constants.Keycloak) (user *models.User, err error) {
	logster.StartFuncLogMsg(fmt.Sprintf("User Uuid: %s", id))

	userKeycloak, err := keycloak.Client.GetUserByID(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, id)
	if err != nil {
		logster.Error(err, "Error getting user with id: "+id)
		return nil, err
	}

	groups, _ := keycloak.Client.GetUserGroups(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		keycloak.Realm,
		gocloak.PString(userKeycloak.ID),
		gocloak.GetGroupsParams{},
	)

	userReturn, _ := utils.KeycloakUserToAPIUser(userKeycloak, groups)

	logster.EndFuncLog()
	return userReturn, nil
}

func GetUsers(keycloak *constants.Keycloak, searchUser *string, limit int) ([]*models.User, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Search user: %+v - Limit: %d", searchUser, limit))

	filters := gocloak.GetUsersParams{Max: gocloak.IntP(limit)}

	if searchUser != nil && *searchUser != "" {
		filters.Search = searchUser // Search users by email/firstName/lastName/username
		filters.Exact = gocloak.BoolP(false)
		filters.Max = gocloak.IntP(limit)
	}

	// Fetch users from Keycloak
	keycloakUsers, err := keycloak.Client.GetUsers(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		keycloak.Realm,
		filters,
	)
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to fetch users from Keycloak: %s", err))
		return nil, err
	}

	// Transform Keycloak users to API users
	userReturn := make([]*models.User, 0, len(keycloakUsers))
	for _, user := range keycloakUsers {
		transformedUser, _ := utils.KeycloakUserToAPIUser(user, nil)
		userReturn = append(userReturn, transformedUser)
	}

	logster.EndFuncLog()
	return userReturn, nil
}

func UpdateUser(userUuid uuid.UUID, dto dto.UpdateUserDto, keycloak *constants.Keycloak) (user *models.User, err error) {
	logster.StartFuncLogMsg(fmt.Sprintf("User Uuid %s", userUuid))

	userKeycloak, err := keycloak.Client.GetUserByID(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, userUuid.String())
	if err != nil {
		logster.Error(err, fmt.Sprintf("Error -  %v", err))
		return nil, err
	}

	if dto.FirstName != nil {
		userKeycloak.FirstName = dto.FirstName
	}

	if dto.LastName != nil {
		userKeycloak.LastName = dto.LastName
	}

	if dto.Country != nil {
		(*userKeycloak.Attributes)["country"] = []string{*dto.Country}
	}

	if dto.Currency != nil {
		(*userKeycloak.Attributes)["currency"] = []string{*dto.Currency}
	}

	if dto.Balance != nil {
		(*userKeycloak.Attributes)["balance"] = []string{*dto.Balance}
	}

	if dto.DisplayName != nil {
		(*userKeycloak.Attributes)["displayName"] = []string{*dto.DisplayName}
	}

	if dto.BirthDate != nil {
		(*userKeycloak.Attributes)["birthDate"] = []string{*dto.BirthDate}
	}

	if dto.EmailExtras != nil {
		emailExtrasJson, err := json.Marshal(dto.EmailExtras)
		if err != nil {
			logster.Error(err, "Error marshaling email extras")
			return nil, err
		}
		(*userKeycloak.Attributes)["email_extras"] = []string{string(emailExtrasJson)}
	}

	if dto.Groups != nil {
		groups, _ := keycloak.Client.GetUserGroups(
			keycloak.Ctx,
			keycloak.AdminToken.AccessToken,
			keycloak.Realm,
			gocloak.PString(userKeycloak.ID),
			gocloak.GetGroupsParams{},
		)

		for _, group := range groups {
			err := keycloak.Client.DeleteUserFromGroup(
				keycloak.Ctx,
				keycloak.AdminToken.AccessToken,
				keycloak.Realm,
				gocloak.PString(userKeycloak.ID),
				*group.ID,
			)
			if err != nil {
				logster.Error(err, fmt.Sprintf("Error -  %v\n", err))
				return nil, err
			}
		}

		for _, group := range *dto.Groups {
			// Find Group in 2 steps. Get parent group and then find subgroup

			groupInfo, err := keycloak.Client.GetGroups(
				keycloak.Ctx,
				keycloak.AdminToken.AccessToken,
				keycloak.Realm,
				gocloak.GetGroupsParams{
					Search: gocloak.StringP(group),
				},
			)
			subGroup, found := lo.Find(*groupInfo[0].SubGroups, func(g gocloak.Group) bool { return *g.Name == group })

			if err == nil && found {
				err = keycloak.Client.AddUserToGroup(
					keycloak.Ctx,
					keycloak.AdminToken.AccessToken,
					keycloak.Realm,
					gocloak.PString(userKeycloak.ID),
					gocloak.PString(subGroup.ID),
				)
				if err != nil {
					logster.Error(err, fmt.Sprintf("Error -  %v\n", err))
					return nil, err
				}
			}

		}

	}

	if dto.IsVerified != nil {
		(*userKeycloak.Attributes)["isVerified"] = []string{*dto.IsVerified}
	}

	if dto.OnboardingFinished != nil {
		(*userKeycloak.Attributes)["onboardingFinished"] = []string{*dto.OnboardingFinished}
	}

	if dto.UtmParams != nil {
		(*userKeycloak.Attributes)["utm"] = []string{*dto.UtmParams}
	}

	if dto.Newsletter != nil {
		(*userKeycloak.Attributes)["newsletter"] = []string{strconv.FormatBool(*dto.Newsletter)}
	}

	if dto.Source != nil {
		(*userKeycloak.Attributes)["source"] = []string{*dto.Source}
	}

	if dto.CurrencySelected != nil {
		(*userKeycloak.Attributes)["currency_selected"] = []string{strconv.FormatBool(*dto.CurrencySelected)}
	}

	err = keycloak.Client.UpdateUser(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, *userKeycloak)
	if err != nil {
		logster.Error(err, fmt.Sprintf("Error -  %v\n", err))
		return nil, err
	}

	groups, _ := keycloak.Client.GetUserGroups(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		keycloak.Realm,
		gocloak.PString(userKeycloak.ID),
		gocloak.GetGroupsParams{},
	)

	userReturn, err := utils.KeycloakUserToAPIUser(userKeycloak, groups)
	if err != nil {
		logster.Error(err, fmt.Sprintf("Error -  %v\n", err))
		return nil, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("End Update user with userId %v\n", userUuid))
	return userReturn, nil
}

func FindUserByEmail(email string, keycloak *constants.Keycloak) (*gocloak.User, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Email: %s", email))

	users, errFind := keycloak.Client.GetUsers(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, gocloak.GetUsersParams{
		Exact:    gocloak.BoolP(true),
		Username: gocloak.StringP(email),
		Email:    gocloak.StringP(email),
	})

	if errFind != nil {
		return nil, errFind
	}

	var user gocloak.User
	if len(users) > 0 {
		user = *users[0]
	}

	logster.EndFuncLog()
	return &user, nil
}

func ResetPassword(userId *string, keycloak *constants.Keycloak) error {
	logster.StartFuncLogMsg(fmt.Sprintf("user UUID: %s", *userId))

	actions := []string{"UPDATE_PASSWORD"}
	err := keycloak.Client.ExecuteActionsEmail(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, gocloak.ExecuteActionsEmail{
		UserID:  userId,
		Actions: &actions,
	})
	if err != nil {
		return err
	}

	logster.EndFuncLog()
	return nil
}

func EmailVerified(bodyEmailVerifiedDTO dto.EmailVerifiedDto, keycloak *constants.Keycloak) (bool, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("body email verified dto: %+v", bodyEmailVerifiedDTO))

	userDto := dto.UpdateUserDto{
		IsVerified: &bodyEmailVerifiedDTO.IsVerified,
	}
	userId := utils.ParseIDToUUID(bodyEmailVerifiedDTO.UserId)
	_, err := UpdateUser(userId, userDto, keycloak)

	userCamundaProcess, err := repository.GetCamundaProcessByVariableUUID(userId)
	if err != nil {
		return false, err
	}

	// Call tasklist to get the task of the user
	task, err := camunda.GetTaskByStateAndProcessInstanceKey(camunda.CREATED, userCamundaProcess.ProcessInstanceKey)
	if err != nil {
		return false, err
	}

	var taskCompleted bool

	if task.ID != "" {
		// Complete task and move to next stage
		taskCompleted, _ = camunda.CompleteTask(task.ID)
	} else {
		logster.Info(fmt.Sprintf("No task found for process instance key %d", userCamundaProcess.ProcessInstanceKey))
	}

	if taskCompleted {
		// Remove user register from db since task is completed
		_, err = repository.DeleteUserCamundaProcessByVariableUUID(userCamundaProcess.FieldUUID)
		if err != nil {
			return false, err
		}
		logster.Info(fmt.Sprintf("Task %s COMPLETED\n", task.ID))
	}

	logster.EndFuncLog()
	return true, nil
}

func GetInfoUserReferral(user *models.User, keycloak *constants.Keycloak) *response_object.UserDto {
	logster.StartFuncLogMsg(fmt.Sprintf("User UUID: %s", user.Uuid))
	// TODO: IMPROVE THIS FUNCTION
	referral, errReferral := repository.GetReferral(user.Uuid)
	if errReferral != nil {
		userDto := utils.CreateMapForUserDto(user, nil, nil)
		return &userDto
	}

	var inviteeUser *models.User
	var errInvitee error

	if referral.InviteeUUID != nil {
		inviteeUser, errInvitee = GetUserById(referral.InviteeUUID.String(), keycloak)
		if errInvitee != nil {
			userDto := utils.CreateMapForUserDto(user, &referral, nil)
			return &userDto
		}
	}

	userDto := utils.CreateMapForUserDto(user, &referral, inviteeUser)

	logster.EndFuncLog()
	return &userDto
}

func GetUserByReferralCode(referralCode string, keycloak *constants.Keycloak) (*models.User, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Referral code: %s", referralCode))

	users, errFind := keycloak.Client.GetUsers(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, gocloak.GetUsersParams{
		Exact: gocloak.BoolP(true),
		Q:     gocloak.StringP("referralCode:" + referralCode),
	})

	if errFind != nil {
		logster.Error(errFind, "Error getting users by referral code")
		logster.EndFuncLog()
		return nil, errFind
	}

	userFound := users[0]

	if userFound == nil {
		logster.Error(nil, "User not found by referral code")
		logster.EndFuncLog()
		return &models.User{}, nil
	}

	groups, _ := keycloak.Client.GetUserGroups(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		keycloak.Realm,
		gocloak.PString(userFound.ID),
		gocloak.GetGroupsParams{},
	)

	userReturn, err := utils.KeycloakUserToAPIUser(userFound, groups)
	if err != nil {
		logster.Error(err, "Error getting user by referral code")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return userReturn, nil
}

func UpdateUseProfilePicture(userUuid uuid.UUID, fileUuid *string, keycloak *constants.Keycloak) error {
	logster.StartFuncLogMsg(fmt.Sprintf("User UUID %s", userUuid.String()))
	userKeycloak, err := keycloak.Client.GetUserByID(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, userUuid.String())
	if err != nil {
		return err
	}

	if fileUuid != nil {
		(*userKeycloak.Attributes)["profilePicture"] = []string{*fileUuid}
	} else {
		(*userKeycloak.Attributes)["profilePicture"] = []string{}
	}

	err = keycloak.Client.UpdateUser(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, *userKeycloak)
	if err != nil {
		return err
	}

	logster.EndFuncLog()
	return nil
}

// ValidateUserAction This function will validate if the user has the required action to "UPDATE_PASSWORD"
// If the user doesn't have the required action, the function will call the CamundaManualTask function, that will validate the process and remove the entry on  camunda_process table
func ValidateUserAction(userKeycloak *gocloak.User, keycloak *constants.Keycloak) error {
	logster.StartFuncLogMsg(fmt.Sprintf("User UUID %+v", userKeycloak.ID))
	camundaProcess, _ := GetCamundaProcessByVariableUUIDAndProcessId(utils.ParseIDToUUID(*userKeycloak.ID), camunda.UserMigrationFlow)
	if lo.Contains(*userKeycloak.RequiredActions, "UPDATE_PASSWORD") && camundaProcess != nil {
		userUuid := userKeycloak.ID

		err := keycloak.Client.ExecuteActionsEmail(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, gocloak.ExecuteActionsEmail{
			UserID:  userUuid,
			Actions: userKeycloak.RequiredActions,
		})
		if err != nil {
			return err
		}

		return utils.CustomErrorStruct{ErrorType: "need_reset_password"}
	} else if !lo.Contains(*userKeycloak.RequiredActions, "UPDATE_PASSWORD") && camundaProcess != nil {
		return CamundaManualTask(utils.ParseIDToUUID(*userKeycloak.ID), nil)
	}

	logster.EndFuncLog()
	return nil
}

func GetMaxUsers(keycloak *constants.Keycloak, searchName *string) ([]*models.User, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Search name %+v", searchName))

	filters := gocloak.GetUsersParams{Max: gocloak.IntP(10000)}

	if searchName != nil && *searchName != "" {
		filters.Search = searchName // Search users by email/firstName/lastName/username
		filters.Exact = gocloak.BoolP(false)
	}

	// Fetch users from Keycloak
	keycloakUsers, err := keycloak.Client.GetUsers(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		keycloak.Realm,
		filters,
	)
	if err != nil {
		logster.Error(err, "failed to fetch users from Keycloak")
		return nil, fmt.Errorf("failed to fetch users from Keycloak: %w", err)
	}

	// Transform Keycloak users to API users
	userReturn := make([]*models.User, 0, len(keycloakUsers))
	for _, user := range keycloakUsers {
		transformedUser, _ := utils.KeycloakUserToAPIUser(user, nil)
		userReturn = append(userReturn, transformedUser)
	}

	logster.EndFuncLog()
	return userReturn, nil
}

// CAMUNDA
func GetAllUsersUuids(keycloak *constants.Keycloak, levelFilter *string) ([]*string, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Level filter %+v", levelFilter))

	var allUsers []*gocloak.User // To store all users
	first := 0                   // Starting index
	maxValue := 1

	for {
		// Fetch users in batches

		keycloakUsers, err := keycloak.Client.GetUsers(
			keycloak.Ctx,
			keycloak.AdminToken.AccessToken,
			keycloak.Realm,
			gocloak.GetUsersParams{
				First: gocloak.IntP(first),
				Max:   gocloak.IntP(maxValue),
			},
		)
		if err != nil {
			return nil, err
		}

		// Append fetched users to the list
		allUsers = append(allUsers, keycloakUsers...)

		// Break the loop if no more users are returned
		if len(keycloakUsers) < maxValue {
			break
		}

		// Increment the starting index for the next batch
		first += maxValue
	}

	// UUIDs of users List
	var userUuidsReturn []*string
	// Loop through users
	for _, user := range allUsers {
		// If levelFilter is not nil, then we need to check if the user is in the correct group
		if levelFilter != nil {
			// Get the groups of the user
			groups, _ := keycloak.Client.GetUserGroups(
				keycloak.Ctx,
				keycloak.AdminToken.AccessToken,
				keycloak.Realm,
				gocloak.PString(user.ID),
				gocloak.GetGroupsParams{},
			)

			for _, group := range groups {
				// Check if the group path contains the levelFilter
				if strings.Contains(*group.Path, *levelFilter) {
					userUuidsReturn = append(userUuidsReturn, user.ID)
					break
				}
			}
		} else {
			userUuidsReturn = append(userUuidsReturn, user.ID)
		}
	}

	logster.EndFuncLog()
	return userUuidsReturn, nil
}

func GetMemberRewardsByUsers(users []string, keycloak *constants.Keycloak) ([]interactive_brokers.UserMemberReward, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Users %+v", users))
	keycloak_utils.RefreshKeycloakAdminToken(keycloak)

	var result []interactive_brokers.UserMemberReward

	for _, user := range users {
		_, found := lo.Find(result, func(r interactive_brokers.UserMemberReward) bool {
			return r.UserUuid == user
		})
		if found {
			continue
		}

		userKeycloak, err := GetUserById(user, keycloak)
		if err != nil {
			logster.Error(err, fmt.Sprintf("Error getting user with id %s", user))
			continue
		}

		// parse group.path : "/membership_levels/base" | "/membership_levels/silver" | "/membership_levels/gold"
		membershipGroup, _ := lo.Find(userKeycloak.Groups, func(g string) bool {
			return strings.Contains(g, "/membership_levels/")
		})

		membershipRewardsValue := PercentageBaseOnMembership(&membershipGroup)

		result = append(result, interactive_brokers.UserMemberReward{
			UserUuid: user,
			Group:    membershipGroup,
			Amount:   membershipRewardsValue.PercentageOnTransaction,
			Currency: userKeycloak.Currency,
		})
	}

	logster.EndFuncLog()
	return result, nil
}

func CreateUserFromShopifyOrder(bodyUserDTO dto.CreateUserDto, keycloak *constants.Keycloak) (createdUser *models.User, err error) {
	logster.StartFuncLog()

	attributes := &map[string][]string{
		"balance":           {"0"},
		"referralCode":      {utils.RandomWordsCode(7)},
		"currency":          {bodyUserDTO.Currency},
		"newsletter":        {"true"},
		"displayName":       {utils.GenerateDisplayName()},
		"source":            {models.SOURCE_SHOPIFY},
		"currency_selected": {"false"},
	}

	if bodyUserDTO.Country != nil {
		(*attributes)["country"] = []string{*bodyUserDTO.Country}
	}

	if bodyUserDTO.UtmParams != nil {
		(*attributes)["utm"] = []string{*bodyUserDTO.UtmParams}
	}

	groups := []string{*Configuration.MembershipLevels.Base, *Configuration.UserTypes.Default}

	user := gocloak.User{
		Enabled:         gocloak.BoolP(true),
		EmailVerified:   gocloak.BoolP(false),
		FirstName:       gocloak.StringP(bodyUserDTO.FirstName),
		LastName:        gocloak.StringP(bodyUserDTO.LastName),
		Email:           gocloak.StringP(bodyUserDTO.Email),
		Groups:          &groups,
		Attributes:      attributes,
		RequiredActions: &[]string{"VERIFY_EMAIL", "UPDATE_PASSWORD"},
	}

	userId, errCreation := keycloak.Client.CreateUser(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		keycloak.Realm,
		user,
	)

	if errCreation != nil {
		logster.Error(errCreation, "Error creating user")
		logster.EndFuncLog()
		return nil, errCreation
	}

	//Email sending was moved to another part of the flow so that the store info is set for the email sending

	userResponse, _ := GetUserById(userId, keycloak)

	userVar := map[string]interface{}{
		"userUuid": userId,
	}

	// Start camunda registration process
	process := camunda.StartProcessInstance(camunda.InjectEnvOnKey(camunda.RegistrationFlow), *camunda.GetCamundaClient(), userVar)

	createUserCamundaProcessDto := dto.CreateCamundaProcessDto{
		FieldUUID:          userResponse.Uuid,
		ProcessInstanceKey: process.ProcessInstanceKey,
		ProcessId:          camunda.RegistrationFlow,
	}

	_, err = CreateCamundaProcess(createUserCamundaProcessDto)
	if err != nil {
		logster.Error(err, "Error creating camunda process")
		logster.EndFuncLog()
	}

	logster.EndFuncLogMsg(fmt.Sprintf("User created: %v", userResponse.Uuid))
	return userResponse, nil
}

func UpdateUserTransactionAndRewardsByCurrency(userUuid string, currencyCode string) {
	logster.StartFuncLog()

	transctionsUpdated, err := repository.UpdateUserTransactionAndRewardsByCurrency(userUuid, currencyCode)
	logster.Info(fmt.Sprintf("Transactions updated: %+v", transctionsUpdated))

	if err != nil {
		logster.Error(err, "Error updating transactions")
	}

	logster.EndFuncLog()
}

func SocialsFinishProfile(loggedUser models.User, userDTO *dto.SocialProfileFinish, keycloak *constants.Keycloak) (*models.User, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("user dto: %+v", userDTO))

	userKeycloak, err := keycloak.Client.GetUserByID(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, loggedUser.Uuid.String())
	if err != nil {
		logster.Error(err, fmt.Sprintf("Error -  %v", err))
		return nil, err
	}

	(*userKeycloak.Attributes)["currency"] = []string{userDTO.Currency}
	(*userKeycloak.Attributes)["referralCode"] = []string{utils.RandomWordsCode(7)}
	(*userKeycloak.Attributes)["displayName"] = []string{utils.GenerateDisplayName()}
	(*userKeycloak.Attributes)["currency_selected"] = []string{"true"}
	(*userKeycloak.Attributes)["balance"] = []string{"0"}
	userKeycloak.EmailVerified = gocloak.BoolP(true)

	if userDTO.FirstName != nil {
		userKeycloak.FirstName = userDTO.FirstName
	}

	if userDTO.LastName != nil {
		userKeycloak.LastName = userDTO.LastName
	}

	err = keycloak.Client.UpdateUser(keycloak.Ctx, keycloak.AdminToken.AccessToken, keycloak.Realm, *userKeycloak)
	if err != nil {
		logster.Error(err, fmt.Sprintf("Error -  %v\n", err))
		return nil, err
	}

	groups, _ := keycloak.Client.GetUserGroups(
		keycloak.Ctx,
		keycloak.AdminToken.AccessToken,
		keycloak.Realm,
		gocloak.PString(userKeycloak.ID),
		gocloak.GetGroupsParams{},
	)

	userReturn, err := utils.KeycloakUserToAPIUser(userKeycloak, groups)
	if err != nil {
		logster.Error(err, fmt.Sprintf("Error -  %v\n", err))
		return nil, err
	}

	// handle utm params
	if userDTO.UtmParams != nil {
		(*userKeycloak.Attributes)["utm"] = []string{*userDTO.UtmParams}
	}

	//Welcome email trigger
	welcomeVar := map[string]interface{}{
		"userUuid":     loggedUser.Uuid.String(),
		"isNewsletter": true,
	}
	welcomeProcess := camunda.StartProcessInstance(camunda.InjectEnvOnKey(camunda.WelcomeEmailFlow), *camunda.GetCamundaClient(), welcomeVar)
	createUserCamundaProcessDto := dto.CreateCamundaProcessDto{
		FieldUUID:          loggedUser.Uuid,
		ProcessInstanceKey: welcomeProcess.ProcessInstanceKey,
		ProcessId:          camunda.WelcomeEmailFlow,
	}

	_, err = CreateCamundaProcess(createUserCamundaProcessDto)
	if err != nil {
		logster.Error(err, "Error creating camunda process for welcome email flow")
	}

	//Never Purchased email trigger
	neverPurchaseVar := map[string]interface{}{
		"userUuid":         loggedUser.Uuid.String(),
		"isNewsletter":     true,
		"haveTransactions": false,
	}
	neverPurchaseFlow := camunda.StartProcessInstance(camunda.InjectEnvOnKey(camunda.NeverPurchased), *camunda.GetCamundaClient(), neverPurchaseVar)
	createUserCamundaProcessDto = dto.CreateCamundaProcessDto{
		FieldUUID:          loggedUser.Uuid,
		ProcessInstanceKey: neverPurchaseFlow.ProcessInstanceKey,
		ProcessId:          camunda.NeverPurchased,
	}

	_, err = CreateCamundaProcess(createUserCamundaProcessDto)
	if err != nil {
		logster.Error(err, "Error creating camunda process for welcome email flow")
	}

	// Handle referral
	if userDTO.ReferralCode != nil {
		errReferral := handleReferral(*userDTO.ReferralCode, userDTO.ReferralClick, keycloak, loggedUser.Uuid)

		if errReferral != nil {
			logster.Error(errReferral, "Error creating camunda process for welcome email flow")
			logster.EndFuncLog()
			return userReturn, errReferral
		}
	}

	logster.EndFuncLog()
	return userReturn, nil
}

func handleReferral(referralCode string, referralClick *string, keycloak *constants.Keycloak, userInviteeUuid uuid.UUID) error {
	logster.StartFuncLogMsg(fmt.Sprintf("Code: %s | ClickUuid: %s | inviteeUuid: %s", referralCode, referralClick, userInviteeUuid))

	userReferrer, errReferrer := GetUserByReferralCode(referralCode, keycloak)

	if errReferrer != nil {
		logster.Error(errReferrer, "Error getting user by referral code")
		logster.EndFuncLog()
		return errReferrer
	}

	modelReferral := utils.CreateReferralDto(userReferrer.Uuid, userInviteeUuid)

	ref, errRef := repository.CreateReferral(modelReferral)

	if referralClick != nil {
		// Update referral click with referral uuid
		referralClickUuid := utils.ParseIDToUUID(*referralClick)
		err := repository.UpdateReferralClickReferral(referralClickUuid, ref.Uuid)
		if err != nil {
			logster.Error(err, "Error updating referral click referral")
			logster.EndFuncLog()
			return err
		}

	}

	if errRef != nil {
		logster.Error(errRef, "Error creating referral")
		logster.EndFuncLog()
		return errRef
	}

	logster.EndFuncLog()
	return nil
}
