package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"io"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"net/http"
	"strings"
)

func GetAidStore(advertiserId string, uuidUser string, keycloak *constants.Keycloak) (*string, error) {
	logster.StartFuncLog()

	user, err := GetUserById(uuidUser, keycloak)
	if err != nil {
		logster.Error(err, "Error getting user")
		return nil, err
	}

	msCJ := dotenv.GetEnv("MS_CJ_URL")
	cjToken := dotenv.GetEnv("CJ_TOKEN")
	publisherId := dotenv.GetEnv("CJ_TAGPEAK_CID")

	logster.Info(fmt.Sprintf("MS_CJ_URL: %s; CJ_TOKEN: %s, CJ_TAGPEAK_CID: %s", advertiserId, cjToken, publisherId))

	url := fmt.Sprintf("%scj-transaction/adid", msCJ)

	body, err := json.Marshal(map[string]interface{}{
		"publisherId":  publisherId,
		"advertiserId": advertiserId,
		"apiToken":     cjToken,
	})
	if err != nil {
		return nil, err
	}

	logster.Info("GetAidStore - Communicating with MS-CJ")
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	logster.Info("GetAidStore - Communication with MS-CJ ended")
	if err != nil {
		return nil, err
	}
	if res == nil {
		logster.Warn("No new transactions")
		return nil, nil
	}

	// Check response status code
	if res.StatusCode != http.StatusOK {
		logster.Error(nil, fmt.Sprintf("Error status code: %d", res.StatusCode))
		return nil, fmt.Errorf("error status code: %d", res.StatusCode)
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		logster.Error(err, "Error reading response body")
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var adIdRes response_object.AdIdResponse
	err = json.Unmarshal(responseBody, &adIdRes)
	if err != nil {
		logster.Error(err, "Error parsing JSON data")
	}

	feeds := adIdRes.ShoppingProductFeeds
	if feeds.TotalCount == 0 || len(feeds.ResultList) == 0 {
		logster.Error(nil, "No adId found in response")
		return nil, fmt.Errorf("no adId found")
	}
	logster.Info(fmt.Sprintf("Total adId count: %d", feeds.TotalCount))

	var adId string

	if feeds.TotalCount == 1 {
		adId = feeds.ResultList[0].AdId
	} else {
		userCountry := strings.ToLower(user.Country)
		userCurrency := strings.ToLower(user.Currency)

		filteredByCountry := lo.Filter(feeds.ResultList, func(item response_object.ShoppingProductFeed, _ int) bool {
			return strings.ToLower(item.Language) == userCountry
		})

		filteredByCurrency := lo.Filter(feeds.ResultList, func(item response_object.ShoppingProductFeed, _ int) bool {
			return strings.ToLower(item.Currency) == userCurrency
		})

		if len(filteredByCountry) == 0 && len(filteredByCurrency) == 0 {
			adId = feeds.ResultList[0].AdId
		} else {
			if len(filteredByCurrency) > 0 && len(filteredByCountry) == 0 {
				adId = filteredByCurrency[0].AdId
			} else {
				adId = filteredByCountry[0].AdId
			}
		}
	}

	logster.EndFuncLog()
	return &adId, nil
}
