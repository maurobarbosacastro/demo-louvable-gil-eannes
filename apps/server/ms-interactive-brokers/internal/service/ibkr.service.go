package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"ms-interactive-brokers/internal/models"
	"ms-interactive-brokers/internal/response_object"
	"ms-interactive-brokers/internal/service/ibkr"
	"ms-interactive-brokers/pkg/logster"
	"ms-interactive-brokers/pkg/utils"
	"strconv"
	"time"
	"unicode"

	"github.com/samber/lo"
)

func StartSession() {
	logster.StartFuncLog()

	// 1. Request the live session token and its expiration time
	lst, lstExpires, err := ibkr.GetLiveSessionToken(ibkr.Appconfig)
	if err != nil {
		logster.Panic(err, fmt.Sprintf("failed to get live session token: %w", err))
	}

	ibkr.Appconfig.LstToken = lst
	ibkr.Appconfig.LstTokenExpire = lstExpires

	// 2. Initialize the brokerage session
	brokerageResp, err := ibkr.InitBrokerageSession(ibkr.Appconfig, ibkr.Appconfig.AccessToken, lst)
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to initialize brokerage session: %w", err))
	}
	defer brokerageResp.Body.Close()

	body, _ := ioutil.ReadAll(brokerageResp.Body)
	if brokerageResp.StatusCode != 200 {
		logster.Error(err, fmt.Sprintf("failed to initialize brokerage session: %s", string(body)))
	}

	var brokerageSessionData models.BrokerageSession
	if err := json.Unmarshal(body, &brokerageSessionData); err != nil {
		logster.Error(err, fmt.Sprintf("failed to decode brokerage session response: %w", err))
	}

	logster.Info(fmt.Sprintf("Brokerage session: %+v\n", brokerageSessionData))
	logster.EndFuncLog()
}

func Logout() {
	logster.StartFuncLog()
	loggedOut, err := ibkr.Logout(ibkr.Appconfig, ibkr.Appconfig.AccessToken, ibkr.Appconfig.LstToken)
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to logout: %+v", err))
	}
	defer loggedOut.Body.Close()

	body, _ := ioutil.ReadAll(loggedOut.Body)
	if loggedOut.StatusCode != 200 {
		logster.Error(err, fmt.Sprintf("failed to logout: %s", string(body)))
	}

	var logout models.Logout
	if err := json.Unmarshal(body, &logout); err != nil {
		logster.Error(err, fmt.Sprintf("failed to decode logout response: %w", err))
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Logged out: %v", logout))
}

func getAccounts() {
	logster.StartFuncLog()
	response, err := ibkr.GetAccounts(ibkr.Appconfig, ibkr.Appconfig.AccessToken, ibkr.Appconfig.LstToken)
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to get accounts: %w", err))
	}
	defer response.Body.Close()

	if err != nil || response.StatusCode != 200 {
		body, _ := ioutil.ReadAll(response.Body)
		logster.Error(err, fmt.Sprintf("failed to get accounts: %s", string(body)))
		logster.EndFuncLog()
		return
	}

	logster.EndFuncLogMsg("Accounts successfully fetched")
}

func GetLastPriceBulk(conids []string) ([]*response_object.LastPriceBulkRO, error) {
	const maxRetries = 3
	const retryDelay = 200 * time.Millisecond

	var response []*response_object.LastPriceBulkRO
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			logster.Warn(fmt.Sprintf("Retry attempt %d after error: %v", attempt, lastErr))
			time.Sleep(retryDelay)
		}

		lastErr = nil

		StartSession()
		getAccounts()

		fields := []int{31, 6119, 6509}

		logster.Info("Requesting market data snapshot for conids")
		marketDataResp, err := ibkr.MarketDataSnapshot(conids, &fields)
		if err != nil {
			return nil, fmt.Errorf("failed to request market data snapshot: %w", err)
		}
		defer marketDataResp.Body.Close()

		if marketDataResp.StatusCode != 200 {
			body, _ := ioutil.ReadAll(marketDataResp.Body)
			return nil, fmt.Errorf("failed to request market data snapshot: %s", string(body))
		}

		var marketDataSnapshot []models.MarketDataSnapshot
		if err := json.NewDecoder(marketDataResp.Body).Decode(&marketDataSnapshot); err != nil {
			return nil, fmt.Errorf("failed to decode market data snapshot response: %w", err)
		}

		logster.Info(fmt.Sprintf("Market data snapshot size: %d | Content: %+v", len(marketDataSnapshot), marketDataSnapshot))

		func() {
			defer func() {
				if r := recover(); r != nil {
					lastErr = fmt.Errorf("panic during mapping operation: %v", r)
				}
			}()

			response = lo.Map(marketDataSnapshot, func(item models.MarketDataSnapshot, index int) *response_object.LastPriceBulkRO {
				logster.Info(fmt.Sprintf("Item for processing: %+v", item))
				if item.LastPrice == nil {
					logster.Info("Last price is nil, skipping")
					return nil
				}

				if item.LastPrice != nil && *item.LastPrice == "N/A" {
					logster.Info("Last price is N/A, skipping")
					return nil
				}

				logster.Info(fmt.Sprintf("Last price: %s", item.LastPrice))
				lastPriceHasLetter := unicode.IsLetter(rune((*item.LastPrice)[0]))
				stringConid := strconv.Itoa(item.ConID)

				if lastPriceHasLetter {
					return &response_object.LastPriceBulkRO{
						ConId:         stringConid,
						LastPrice:     (*item.LastPrice)[1:],
						LastPriceType: utils.GetLastPriceType(*item.LastPrice),
					}
				}

				return &response_object.LastPriceBulkRO{
					ConId:         stringConid,
					LastPrice:     *item.LastPrice,
					LastPriceType: "",
				}
			})
		}()

		// If we had a panic during mapping, continue to next retry
		if lastErr != nil {
			continue
		}

		go Logout()

		return response, nil
	}

	// If we exhausted all retries, return the last error
	return nil, fmt.Errorf("failed after %d retries. Last error: %v", maxRetries, lastErr)
}

func tickle() (*models.TickleResponseSuccess, error) {
	logster.StartFuncLog()

	respTickle, errTickle := ibkr.Tickle(ibkr.Appconfig, ibkr.Appconfig.AccessToken, ibkr.Appconfig.LstToken)

	if errTickle != nil {
		logster.Error(errTickle, "Error while tickling")
		return nil, errTickle
	}
	defer respTickle.Body.Close()

	if respTickle.StatusCode != 200 {
		body, errReadAll := ioutil.ReadAll(respTickle.Body)
		logster.Error(errReadAll, fmt.Sprintf("Fail to read body %s", string(body)))
		return nil, errReadAll
	}

	var tickleData models.TickleResponseSuccess
	if err := json.NewDecoder(respTickle.Body).Decode(&tickleData); err != nil {
		var tickleResponseError models.TickleResponseError
		if errE := json.NewDecoder(respTickle.Body).Decode(&tickleResponseError); errE != nil {
			logster.Error(err, fmt.Sprintf("Failed to parse data %w", err))
			return nil, errE
		}
	}

	logster.EndFuncLog()
	return &tickleData, nil
}

// Function not finished -> missing processing of message and conids received.
func getLastPriceBuildWS(conids []string) (*[]response_object.LastPriceBulkRO, error) {
	logster.StartFuncLogMsg(conids)

	StartSession()

	// Tickle endpoint call to get session value
	tickle, errorTickling := tickle()
	if errorTickling != nil {
		logster.Error(errorTickling, "Tickle error")
		logster.EndFuncLog()
		return nil, errorTickling
	}
	sessionToken := tickle.Session
	accessToken := ibkr.Appconfig.AccessToken

	// Start WebSocket client
	ibkrWsClient := ibkr.NewWebSocketClient(accessToken, sessionToken)
	if errConnectWs := ibkrWsClient.Connect(); errConnectWs != nil {
		return nil, errConnectWs
	}

	// Add a 5 seconds timer to give time for the session to start on the IBKR side.
	time.Sleep(5 * time.Second)

	// Read messages continuasly in a go routine so that we can still write messages and the main thread is not blocked.
	go ibkrWsClient.RunForever()

	lastPrices := lo.Map(conids, func(conid string, index int) response_object.LastPriceBulkRO {
		return response_object.LastPriceBulkRO{
			ConId: conid,
		}
	})

	formatWriteMessage := "smd+%s+{\"fields\":[\"31\", \"6509\", \"6119\"]}"
	formatCleanWriteMessage := "umd+%s+{}"
	messages := lo.Map(lastPrices, func(item response_object.LastPriceBulkRO, index int) string {
		return fmt.Sprintf(formatWriteMessage, item.ConId)
	})

	for _, message := range messages {
		ibkrWsClient.WriteTextMessage(message)
	}

	done := make(chan bool)

	go func() {
		for {
			marketData := <-ibkrWsClient.GetMarketDataChannel()
			logster.Debug(fmt.Sprintf("Has last price %s", marketData.Fields["31"]))
			_, idx, found := lo.FindIndexOf(lastPrices, func(item response_object.LastPriceBulkRO) bool {
				return string(marketData.ConID) == item.ConId
			})

			if found {
				lastPrices[idx].LastPrice = marketData.Fields["31"].(string)
			}

			select {
			case <-done:
				return
			default:
				// continue
			}
		}
	}()

	time.Sleep(1 * time.Minute)
	done <- true
	clearMessages := lo.Map(lastPrices, func(item response_object.LastPriceBulkRO, index int) string {
		return fmt.Sprintf(formatCleanWriteMessage, item.ConId)
	})

	for _, message := range clearMessages {
		ibkrWsClient.WriteTextMessage(message)
	}
	ibkrWsClient.Close()

	go Logout()
	logster.EndFuncLog()
	return utils.Ptr(lastPrices), nil
}

func GetLastPricesBasedOnAccountPositions(conids []string) (*[]response_object.LastPriceBulkRO, error) {
	lst, lstExpires, err := ibkr.GetLiveSessionToken(ibkr.Appconfig)
	if err != nil {
		logster.Panic(err, fmt.Sprintf("failed to get live session token: %v", err))
	}

	ibkr.Appconfig.LstToken = lst
	ibkr.Appconfig.LstTokenExpire = lstExpires

	portfolioAccountsResp, errPortfolioAccounts := ibkr.GetPortfolioAccounts(ibkr.Appconfig, ibkr.Appconfig.AccessToken, lst)

	if errPortfolioAccounts != nil {
		logster.Error(errPortfolioAccounts, "Get accounts error")
		return nil, errPortfolioAccounts
	}

	// PortfolioAccounts
	defer portfolioAccountsResp.Body.Close()

	if portfolioAccountsResp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(portfolioAccountsResp.Body)
		logster.Error(nil, fmt.Sprintf("Failed to get positions account: %s", string(body)))
		return nil, fmt.Errorf("%v", body)
	}

	var accounts []models.Account
	if err := json.NewDecoder(portfolioAccountsResp.Body).Decode(&accounts); err != nil {
		logster.Error(err, "Failed to decode")
		return nil, err
	}

	accountPositions, errAccountPositions := ibkr.GetPositionFromAccountId(accounts[0].AccountID, ibkr.Appconfig, ibkr.Appconfig.AccessToken, lst)
	if errAccountPositions != nil {
		logster.Error(errAccountPositions, "Error getting positions")
		return nil, errAccountPositions
	}
	defer accountPositions.Body.Close()

	if accountPositions.StatusCode != 200 {
		body, _ := io.ReadAll(accountPositions.Body)
		logster.Error(nil, fmt.Sprintf("Failed to get positions: %s", string(body)))
		return nil, fmt.Errorf("%v", body)
	}

	var positions []models.Position
	if err := json.NewDecoder(accountPositions.Body).Decode(&positions); err != nil {
		logster.Error(err, "Failed to decode")
	}

	lastPrices := lo.Map(conids, func(conid string, index int) response_object.LastPriceBulkRO {
		position, found := lo.Find(positions, func(item models.Position) bool {
			return strconv.Itoa(item.Conid) == conid
		})

		if !found {
			return response_object.LastPriceBulkRO{
				ConId: conid,
			}
		}

		return response_object.LastPriceBulkRO{
			ConId:     conid,
			LastPrice: strconv.FormatFloat((position.MktValue / position.Position), 'f', 2, 64),
		}
	})

	return &lastPrices, nil
}
