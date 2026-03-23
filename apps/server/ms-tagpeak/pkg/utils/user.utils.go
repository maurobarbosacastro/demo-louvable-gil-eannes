package utils

import (
	"fmt"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/dotenv"
	"strconv"

	"github.com/Nerzal/gocloak/v13"
	"github.com/samber/lo"
)

func GetBalance(user *gocloak.User) (float64, error) {
	if user.Attributes != nil {
		if balanceSlice, exists := (*user.Attributes)["balance"]; exists && len(balanceSlice) > 0 {
			return strconv.ParseFloat(balanceSlice[0], 64)
		}
	}

	return 0, fmt.Errorf("balance attribute not found")
}

func GetUserName(user *models.User) string {
	if user.DisplayName != "" {
		return user.DisplayName
	} else {
		return user.FirstName + " " + user.LastName
	}
}

func ConvertToBoolean(user *gocloak.User, attribute string) (bool, error) {
	if user.Attributes != nil {
		if attributeSlice, exists := (*user.Attributes)[attribute]; exists && len(attributeSlice) > 0 {
			return strconv.ParseBool(attributeSlice[0])
		}
	}

	return false, fmt.Errorf("%s attribute not found", attribute)
}

func ConvertToFloat(user *gocloak.User, attribute string) (float64, error) {
	if user.Attributes != nil {
		if attributeSlice, exists := (*user.Attributes)[attribute]; exists && len(attributeSlice) > 0 {
			return strconv.ParseFloat(attributeSlice[0], 64)
		}
	}
	return 0.0, nil
}

func GetFromAttributes(user *gocloak.User, name string) (string, error) {
	if user.Attributes != nil {
		if slice, exists := (*user.Attributes)[name]; exists && len(slice) > 0 {
			return slice[0], nil
		}
	}

	return "", nil
}

func CreateMapForUserDto(user *models.User, referral *models.Referral, inviteeUser *models.User) response_object.UserDto {
	var referralDto response_object.InviteeDto

	userDto := response_object.UserDto{
		Uuid:               user.Uuid,
		Email:              user.Email,
		CreatedAt:          user.CreatedAt,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		Balance:            user.Balance,
		Country:            user.Country,
		Currency:           user.Currency,
		ReferralCode:       user.ReferralCode,
		DisplayName:        user.DisplayName,
		BirthDate:          user.BirthDate,
		Groups:             user.Groups,
		IsVerified:         user.IsVerified,
		OnboardingFinished: user.OnboardingFinished,
		Newsletter:         user.Newsletter,
	}

	if user.ProfilePicture != "" {
		userDto.ProfilePicture = &user.ProfilePicture
	}

	if referral != nil && inviteeUser != nil {
		referralDto = CreateMapForInvitee(referral, inviteeUser)
		userDto.Referral = &referralDto
	}

	if user.TransactionPercentage != nil {
		userDto.TransactionPercentage = user.TransactionPercentage
	}

	if user.RewardPercentage != nil {
		userDto.RewardPercentage = user.RewardPercentage
	}

	return userDto
}

func KeycloakUserToAPIUser(user *gocloak.User, groups []*gocloak.Group) (*models.User, error) {
	balance, _ := GetBalance(user)
	country, _ := GetFromAttributes(user, "country")
	currency, _ := GetFromAttributes(user, "currency")
	referralCode, _ := GetFromAttributes(user, "referralCode")
	displayName, _ := GetFromAttributes(user, "displayName")
	birthDate, _ := GetFromAttributes(user, "birthDate")
	isVerified, _ := ConvertToBoolean(user, "isVerified")
	onboardingFinished, _ := ConvertToBoolean(user, "onboardingFinished")
	profilePictureAtt, _ := GetFromAttributes(user, "profilePicture")
	newsletter, _ := GetFromAttributes(user, "newsletter")
	currencySelected, _ := ConvertToBoolean(user, "currency_selected")
	influencerAmount, _ := ConvertToFloat(user, "influencer_amount")
	source, _ := GetFromAttributes(user, "source")

	url := dotenv.GetEnv("MS_IMAGES_SERVER_PUBLIC_URL")
	var profilePicture string
	if profilePictureAtt != "" {
		profilePicture = fmt.Sprintf(url+"%s/profilePicture.webp", profilePictureAtt)
	}

	var transactionPercentage float64
	transactionPercentageAlt, _ := GetFromAttributes(user, "user_percent")
	if transactionPercentageAlt != "" {
		conv, err := strconv.ParseFloat(transactionPercentageAlt, 64)
		if err == nil {
			transactionPercentage = conv
		}
	}

	var rewardPercentage float64
	rewardPercentageAlt, _ := GetFromAttributes(user, "ref_percent")
	if rewardPercentageAlt != "" {
		conv, err := strconv.ParseFloat(rewardPercentageAlt, 64)
		if err == nil {
			rewardPercentage = conv
		}
	}
	newsletterVal, _ := strconv.ParseBool(newsletter)

	var legacyId string
	legacyId, _ = GetFromAttributes(user, "legacyId")
	if legacyId == "" {
		legacyId = "NULL"
	}

	userResponse := models.User{
		Uuid:               ParseIDToUUID(gocloak.PString(user.ID)),
		Email:              gocloak.PString(user.Email),
		CreatedAt:          gocloak.PInt64(user.CreatedTimestamp),
		FirstName:          gocloak.PString(user.FirstName),
		LastName:           gocloak.PString(user.LastName),
		Balance:            balance,
		Country:            country,
		Currency:           currency,
		ReferralCode:       referralCode,
		DisplayName:        displayName,
		BirthDate:          birthDate,
		Groups:             lo.Map(groups, func(g *gocloak.Group, _ int) string { return gocloak.PString(g.Path) }),
		IsVerified:         isVerified,
		OnboardingFinished: onboardingFinished,
		ProfilePicture:     profilePicture,
		Newsletter:         newsletterVal,
		CurrencySelected:   currencySelected,
		LegacyId:           &legacyId,
		InfluencerAmount:   influencerAmount,
		Source:             source,
	}

	if transactionPercentage > 0 {
		userResponse.TransactionPercentage = FloatPointer(transactionPercentage)
	}
	if rewardPercentage > 0 {
		userResponse.RewardPercentage = FloatPointer(rewardPercentage)
	}

	return &userResponse, nil
}
