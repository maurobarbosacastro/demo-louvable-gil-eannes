package service

import (
	"fmt"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/http_client"
	"ms-tagpeak/pkg/logster"
	"net/http"
	"strconv"
)

func CheckVatValidity(vatNumber string) (*response_object.EuVatApiResponse, error) {

	checkerUrl := dotenv.GetEnv("VAT_CHECKER_URL")
	checkerApiKey := dotenv.GetEnv("VAT_CHECKER_API_KEY")
	url := checkerUrl + "/validate?access_key=" + checkerApiKey + "&vat_number=" + vatNumber
	httpClient := &http_client.HttpClient{HttpClient: &http.Client{}}

	response := &response_object.EuVatApiResponse{}
	_, err := httpClient.Get(url, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func SetMigratedUserBalance(keycloak *constants.Keycloak) {
	logster.StartFuncLog()

	list, err := GetUsersWithTotalAmountReward()
	if err != nil {
		logster.Error(err, "Error")
	}

	for _, user := range list {
		formatedBalance := strconv.FormatFloat(user.Total, 'f', 2, 64)
		dtoUpdate := dto.UpdateUserDto{
			Balance: &formatedBalance,
		}
		updateUser, err := UpdateUser(user.User, dtoUpdate, keycloak)
		if err != nil {
			logster.Error(err, "Error")
		}

		if updateUser != nil {
			logster.Info(fmt.Sprintf("User with uuid %s updated", user.User))
		}
	}

	logster.EndFuncLog()
}
