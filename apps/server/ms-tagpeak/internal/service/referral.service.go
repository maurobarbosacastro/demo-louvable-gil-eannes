package service

import (
	"errors"
	"fmt"
	"math"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetReferralById(uuid uuid.UUID, keycloak *constants.Keycloak) (response_object.ReferralDto, error) {
	// TODO: IMPROVE THIS FUNCTION
	referral, err := repository.GetReferral(uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response_object.ReferralDto{}, utils.CustomErrorStruct{}.NotFoundError("Referral", uuid)
		}
		return response_object.ReferralDto{}, err
	}

	var user *models.User
	var errUser error

	if referral.ReferrerUUID != nil {
		user, errUser = GetUserById(referral.ReferrerUUID.String(), keycloak)

		if errUser != nil {
			return response_object.ReferralDto{}, err
		}

		return utils.CreateMapForReferral(&referral, user), nil
	}

	return response_object.ReferralDto{}, nil
}

func GetAllReferralByUserUuid(pag pagination.PaginationParams, uuid uuid.UUID) (*pagination.PaginationResult, error) {

	referrals, err := repository.GetAllReferralByUserUuidWithPagination(pag, uuid)

	if err != nil {
		return nil, err
	}

	return referrals, nil
}

// ReferredByWho - Check if user was referred by another user
// Returns the Referral owner if the user was referred by another user
func ReferredByWho(invitedUserUUID uuid.UUID) (*string, error) {

	referral, err := repository.GetReferralByInvitee(invitedUserUUID)
	if err != nil {
		return nil, err
	}
	if referral.ReferrerUUID == nil {
		return nil, nil
	} else {
		u := referral.ReferrerUUID.String()
		return &u, nil
	}
}

func CreateReferralClick(referralUUID *uuid.UUID, code string) (*models.ReferralClicks, error) {

	model := utils.CreateReferralClickDto(referralUUID, code)

	res, err := repository.CreateReferralClick(model)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetReferralsInfoByUserUuid(referrerUuid uuid.UUID, keycloak *constants.Keycloak) (*response_object.ReferralInfo, error) {

	var response response_object.ReferralInfo

	user, err := GetUserById(referrerUuid.String(), keycloak)
	if err != nil {
		return nil, err
	}

	// GET NUMBER OF TOTAL CLICKS
	response.TotalClicks, _ = repository.CountReferralClickByCode(user.ReferralCode)

	// GET NUMBER OF USER REGISTERED
	response.TotalUserRegistered, _ = repository.CountReferralByUserRegistered(referrerUuid)

	// GET NUMBER OF FIRST PURCHASE
	response.TotalFirstPurchase, _ = repository.CountReferralByFirstTransaction(referrerUuid, true)

	startDate := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())

	// GET NUMBER OF TOTAL CLICKS ON LAST 4 MONTHS, BY MONTH
	value, _ := repository.GetClicksByMonth(user.ReferralCode, startDate, time.Now().AddDate(0, 0, 1))
	appendClicksMonthData(&response, startDate, value)

	for i := 1; i < 4; i++ {
		startDate, endDate := handleDate(i)
		value, _ := repository.GetClicksByMonth(user.ReferralCode, startDate, endDate)
		appendClicksMonthData(&response, startDate, value)
	}

	// GET NUMBER OF USER REGISTERED ON LAST 4 MONTHS, BY MONTH
	value, _ = repository.GetRegisteredByMonth(referrerUuid, startDate, time.Now().AddDate(0, 0, 1))
	appendRegisteredMonthData(&response, startDate, value)

	for i := 1; i < 4; i++ {
		startDate, endDate := handleDate(i)
		value, _ := repository.GetRegisteredByMonth(referrerUuid, startDate, endDate)
		appendRegisteredMonthData(&response, startDate, value)
	}

	// GET NUMBER OF USER THAT MAID FIRST PURCHASE ON LAST 4 MONTHS, BY MONTH
	value, _ = repository.GetFirstPurchaseByMonth(referrerUuid, startDate, time.Now().AddDate(0, 0, 1), true)
	appendFirstPurchaseMonthData(&response, startDate, value)

	for i := 1; i < 4; i++ {
		startDate, endDate := handleDate(i)
		value, _ := repository.GetFirstPurchaseByMonth(referrerUuid, startDate, endDate, true)
		appendFirstPurchaseMonthData(&response, startDate, value)
	}

	return &response, nil
}

func GetRevenuesInfoByUserUuid(referrerUuid uuid.UUID) (*response_object.ReferralRevenueInfo, error) {

	var response response_object.ReferralRevenueInfo

	// GET SUM OF REVENUE
	amount, _ := repository.GetAllReferralRevenueByReferrerUuid(referrerUuid)

	roundedAmount := math.Round(*amount*100) / 100
	response.TotalRevenue = &roundedAmount

	startDate := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())

	// GET SUM OF REVENUE BY MONTH
	value, _ := repository.GetRevenueByMonth(referrerUuid, startDate, time.Now().AddDate(0, 0, 1))
	roundedValue := math.Round(value*100) / 100
	appendRevenueMonthData(&response, startDate, roundedValue)

	for i := 1; i < 11; i++ {
		startDate, endDate := handleDate(i)
		value, _ := repository.GetRevenueByMonth(referrerUuid, startDate, endDate)
		appendRevenueMonthData(&response, startDate, value)
	}

	return &response, nil
}

func GetUsersReferralsRevenueInfoByReferrerUuid(referrerUuid uuid.UUID, keycloak *constants.Keycloak) (*[]response_object.UserReferralRevenueInfoDto, error) {

	// Get list of usersUUID
	referrals, err := repository.GetReferralsByReferrerUuid(referrerUuid)
	if err != nil {
		return nil, err
	}

	var response = make([]response_object.UserReferralRevenueInfoDto, len(referrals))

	// Get UserInfo for each usersUUID
	for index, referral := range referrals {

		userInfo, err := GetUserById(referral.InviteeUUID.String(), keycloak)
		if err != nil {
			return nil, err
		}

		amount, err := repository.GetAmountByReferralUuid(referral.Uuid)
		if err != nil {
			return nil, err
		}

		response[index] = response_object.UserReferralRevenueInfoDto{
			Uuid:                       userInfo.Uuid,
			FirstName:                  userInfo.FirstName,
			LastName:                   userInfo.LastName,
			ProfilePicture:             userInfo.ProfilePicture,
			ReferredValue:              amount,
			FirstTransactionSuccessful: referral.SuccessfulFirstTransaction,
			DisplayName:                userInfo.DisplayName,
		}
	}
	return &response, nil
}

func GetAllReferralRevenueByReferrerUuid(referrerUuid uuid.UUID) (*float64, error) {
	amount, err := repository.GetAllReferralRevenueByReferrerUuid(referrerUuid)
	if err != nil {
		return nil, err
	}
	// Ensure the amount has 2 decimal places
	roundedAmount := math.Round(*amount*100) / 100

	return &roundedAmount, nil
}

func GetReferralByInvitee(inviteeUUID uuid.UUID) (*models.Referral, error) {
	logster.StartFuncLog()

	referral, err := repository.GetReferralByInvitee(inviteeUUID)
	if err != nil {
		logster.Error(err, "GetReferralByInvitee")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Referral found %s", referral.Uuid))
	return referral, nil
}

func ValidateReferralCode(code string, keycloak *constants.Keycloak) (bool, error) {
	logster.StartFuncLog()

	user, err := GetUserByReferralCode(code, keycloak)
	if err != nil && user != nil {
		return false, err
	}

	logster.EndFuncLog()
	return user != nil, nil
}

func handleDate(subtractIndexMonth int) (time.Time, time.Time) {
	currentTime := time.Now()
	targetMonth := currentTime.AddDate(0, -subtractIndexMonth, 0)

	firstDayOfTargetMonth := time.Date(targetMonth.Year(), targetMonth.Month(), 1, 0, 0, 0, 0, targetMonth.Location())

	firstDayOfNextMonth := firstDayOfTargetMonth.AddDate(0, 1, 0)
	lastDayOfTargetMonth := firstDayOfNextMonth

	return firstDayOfTargetMonth, lastDayOfTargetMonth
}

func appendRevenueMonthData(response *response_object.ReferralRevenueInfo, startDate time.Time, value float64) {
	monthData := response_object.MonthData{
		Month: startDate.Format("Jan 06"),
		Value: value,
	}

	response.RevenueByMonth = append(response.RevenueByMonth, monthData)
}

func appendRegisteredMonthData(response *response_object.ReferralInfo, startDate time.Time, value float64) {
	monthData := response_object.MonthData{
		Month: startDate.Format("Jan 06"),
		Value: value,
	}
	response.RegisteredByMonth = append(response.RegisteredByMonth, monthData)

}

func appendFirstPurchaseMonthData(response *response_object.ReferralInfo, startDate time.Time, value float64) {
	monthData := response_object.MonthData{
		Month: startDate.Format("Jan 06"),
		Value: value,
	}
	response.FirstPurchaseByMonth = append(response.FirstPurchaseByMonth, monthData)
}

func appendClicksMonthData(response *response_object.ReferralInfo, startDate time.Time, value float64) {
	monthData := response_object.MonthData{
		Month: startDate.Format("Jan 06"),
		Value: value,
	}
	response.ClicksByMonth = append(response.ClicksByMonth, monthData)
}
