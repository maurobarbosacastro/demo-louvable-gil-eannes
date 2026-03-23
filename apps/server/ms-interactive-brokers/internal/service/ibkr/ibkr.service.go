package ibkr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ms-interactive-brokers/internal/models"
	"ms-interactive-brokers/pkg/logster"
	"ms-interactive-brokers/pkg/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type AppConfig struct {
	ConsumerKey       string
	SignatureKeyFP    string
	EncryptionKeyFP   string
	DHPrime           string
	DHGenerator       int
	Realm             string
	AccessToken       string
	AccessTokenSecret string
	LstToken          string
	LstTokenExpire    int64
}

var Appconfig *AppConfig

func LoadConfig() (*AppConfig, error) {
	dhGenerator, _ := strconv.Atoi(getEnvDefault("DH_GENERATOR", "2"))

	return &AppConfig{
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		SignatureKeyFP:    os.Getenv("SIGNATURE_KEY_FP"),
		EncryptionKeyFP:   os.Getenv("ENCRYPTION_KEY_FP"),
		DHPrime:           os.Getenv("DH_PRIME"),
		DHGenerator:       dhGenerator,
		Realm:             getEnvDefault("REALM", "limited_poa"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		LstToken:          "",
		LstTokenExpire:    0,
	}, nil
}

func getEnvDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func IsTokenExpired() bool {
	return Appconfig.LstTokenExpire < time.Now().Unix()
}

func sendOAuthRequest(
	config *AppConfig,
	method string,
	requestURL string,
	oauthToken string,
	liveSessionToken string,
	extraHeaders map[string]string,
	requestParams map[string]string,
	signatureMethod string,
	prepend string,
) (*http.Response, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Sending request: %s %+v", requestURL, requestParams))

	headers := map[string]string{
		"oauth_consumer_key":     config.ConsumerKey,
		"oauth_nonce":            utils.GenerateOAuthNonce(),
		"oauth_signature_method": signatureMethod,
		"oauth_timestamp":        utils.GenerateRequestTimestamp(),
	}

	if oauthToken != "" {
		headers["oauth_token"] = oauthToken
	}
	if extraHeaders != nil {
		for k, v := range extraHeaders {
			headers[k] = v
		}
	}

	baseString := utils.GenerateBaseString(method, requestURL, headers, requestParams, nil, nil, nil, prepend)

	var signature string
	var err error

	if signatureMethod == "HMAC-SHA256" {
		signature, err = utils.GenerateHMACSHA256Signature(baseString, liveSessionToken)
	} else {
		sigKey, err := utils.ReadPrivateKey(config.SignatureKeyFP)
		if err != nil {
			logster.Error(err, fmt.Sprintf("failed to read private key: %w", err))
			return nil, err
		}
		signature, err = utils.GenerateRSASHA256Signature(baseString, sigKey)
	}
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to generate signature: %w", err))
		return nil, err
	}

	headers["oauth_signature"] = signature

	authHeader := utils.GenerateAuthorizationHeaderString(headers, config.Realm)

	// Create request
	req, err := http.NewRequest(method, requestURL, nil)
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to create request: %w", err))
		return nil, err
	}

	req.Header.Set("Authorization", authHeader)

	// Add query parameters if any
	if requestParams != nil {
		q := req.URL.Query()
		for k, v := range requestParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{Timeout: 10 * time.Second}
	logster.EndFuncLog()
	return client.Do(req)
}

func GetLiveSessionToken(config *AppConfig) (string, int64, error) {
	logster.StartFuncLog()
	const (
		requestURL = "https://api.ibkr.com/v1/api/oauth/live_session_token"
		method     = "POST"
		sigMethod  = "RSA-SHA256"
	)

	dhRandom := utils.GenerateDHRandomBytes()
	dhChallenge := utils.GenerateDHChallenge(config.DHPrime, dhRandom, int64(config.DHGenerator))

	encKey, err := utils.ReadPrivateKey(config.EncryptionKeyFP)
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to read private key: %w", err))
		return "", 0, err
	}

	prepend, err := utils.CalculateLSTPrepend(config.AccessTokenSecret, encKey)
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to calculate LST prepend: %w", err))
		return "", 0, err
	}

	extraHeaders := map[string]string{
		"diffie_hellman_challenge": dhChallenge,
	}

	resp, err := sendOAuthRequest(
		config,
		method,
		requestURL,
		config.AccessToken,
		"",
		extraHeaders,
		nil,
		sigMethod,
		prepend,
	)
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to send request: %w", err))
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		logster.Error(err, fmt.Sprintf("request failed: %s", string(body)))
		return "", 0, fmt.Errorf("request failed: %s", string(body))
	}

	var response struct {
		DHResponse    string `json:"diffie_hellman_response"`
		LSTSignature  string `json:"live_session_token_signature"`
		LSTExpiration int64  `json:"live_session_token_expiration"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		logster.Error(err, fmt.Sprintf("failed to decode response: %w", err))
		return "", 0, err
	}

	lst, err := utils.CalculateLiveSessionToken(
		config.DHPrime,
		dhRandom,
		response.DHResponse,
		prepend,
	)
	if err != nil {
		logster.Error(err, fmt.Sprintf("failed to calculate live session token: %w", err))
		return "", 0, err
	}

	if !utils.ValidateLiveSessionToken(lst, response.LSTSignature, config.ConsumerKey) {
		logster.Error(err, fmt.Sprintf("LST validation failed"))
		return "", 0, fmt.Errorf("LST validation failed")
	}

	logster.EndFuncLog()
	return lst, response.LSTExpiration, nil
}

// Session management functions
func InitBrokerageSession(config *AppConfig, accessToken, lst string) (*http.Response, error) {
	logster.StartFuncLog()

	params := map[string]string{
		"compete": "true",
		"publish": "true",
	}

	return sendOAuthRequest(
		config,
		"POST",
		"https://api.ibkr.com/v1/api/iserver/auth/ssodh/init",
		accessToken,
		lst,
		nil,
		params,
		"HMAC-SHA256",
		"",
	)
}

func brokerageSessionStatus(config *AppConfig, accessToken, lst string) (*http.Response, error) {
	logster.StartFuncLog()
	// https://api.ibkr.com/iserver/auth/status

	return sendOAuthRequest(
		config,
		"POST",
		"https://api.ibkr.com/v1/api/iserver/auth/status",
		accessToken,
		lst,
		nil,
		nil,
		"HMAC-SHA256",
		"",
	)
}

func Logout(config *AppConfig, accessToken, lst string) (*http.Response, error) {
	logster.StartFuncLog()

	return sendOAuthRequest(
		config,
		"POST",
		"https://api.ibkr.com/v1/api/logout",
		accessToken,
		lst,
		nil,
		nil,
		"HMAC-SHA256",
		"",
	)
}

func TokenAndBridgeOpen() bool {
	if IsTokenExpired() {
		return false
	}

	brokerageSessionStatusResp, err := brokerageSessionStatus(Appconfig, Appconfig.AccessToken, Appconfig.LstToken)
	if err != nil {
		logster.Error(err, "failed to check brokerage session status")
		return false
	}

	body, _ := ioutil.ReadAll(brokerageSessionStatusResp.Body)
	if brokerageSessionStatusResp.StatusCode != 200 {
		logster.Error(err, fmt.Sprintf("failed to check brokerage session status: %s", string(body)))
		return false
	}
	var brokerageSessionStatusData models.BrokerageSession
	if err := json.Unmarshal(body, &brokerageSessionStatusData); err != nil {
		logster.Error(err, "failed to decode brokerage session status response")
		return false
	}

	if !brokerageSessionStatusData.Authenticated {
		logster.Info("Brokerage session not authenticated, initializing...")
		brokerageResp, err := InitBrokerageSession(Appconfig, Appconfig.AccessToken, Appconfig.LstToken)
		if err != nil {
			logster.Error(err, fmt.Sprintf("failed to initialize brokerage session: %w", err))
			return false
		}
		defer brokerageResp.Body.Close()

		body, _ := ioutil.ReadAll(brokerageResp.Body)
		if brokerageResp.StatusCode != 200 {
			logster.Error(err, fmt.Sprintf("failed to initialize brokerage session: %s", string(body)))
			return false
		}

		var brokerageSessionData models.BrokerageSession
		if err := json.Unmarshal(body, &brokerageSessionData); err != nil {
			logster.Error(err, fmt.Sprintf("failed to decode brokerage session response: %w", err))
			return false
		}
		if brokerageSessionData.Authenticated {
			return true
		} else {
			return false
		}
	}

	return true
}

func GetAccounts(config *AppConfig, accessToken, lst string) (*http.Response, error) {
	return sendOAuthRequest(
		config,
		"GET",
		"https://api.ibkr.com/v1/api/iserver/accounts",
		accessToken,
		lst,
		nil,
		nil,
		"HMAC-SHA256",
		"",
	)
}

func MarketDataSnapshot(conids []string, fields *[]int) (*http.Response, error) {
	config := Appconfig

	params := map[string]string{
		"conids": strings.Trim(strings.Join(strings.Fields(fmt.Sprint(conids)), ","), "[]"),
	}

	if fields != nil {
		params["fields"] = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(*fields)), ","), "[]")
	}

	return sendOAuthRequest(
		config,
		"GET",
		"https://api.ibkr.com/v1/api/iserver/marketdata/snapshot",
		config.AccessToken,
		config.LstToken,
		nil,
		params,
		"HMAC-SHA256",
		"",
	)
}

func CheckConidInfo(conid string) (*http.Response, error) {
	config := Appconfig
	params := map[string]string{
		"conid": conid,
	}

	return sendOAuthRequest(
		config,
		"GET",
		"https://api.ibkr.com/v1/api/md/regsnapshot",
		config.AccessToken,
		config.LstToken,
		nil,
		params,
		"HMAC-SHA256",
		"",
	)
}

func GetContractBySymbol(config *AppConfig, accessToken, lst string, isin string) (*http.Response, error) {
	params := map[string]string{
		"symbol": isin,
	}

	return sendOAuthRequest(
		config,
		"POST",
		"https://api.ibkr.com/v1/api/iserver/secdef/search",
		accessToken,
		lst,
		nil,
		params,
		"HMAC-SHA256",
		"",
	)
}

func GetPositionsFromInstrument(config *AppConfig, accessToken, lst string, conid string) (*http.Response, error) {
	return sendOAuthRequest(
		config,
		"GET",
		"https://api.ibkr.com/v1/api/portfolio/positions/"+conid,
		accessToken,
		lst,
		nil,
		nil,
		"HMAC-SHA256",
		"",
	)
}

func Tickle(config *AppConfig, accessToken, lst string) (*http.Response, error) {
	return sendOAuthRequest(
		config,
		"POST",
		"http://api.ibkr.com/v1/api/tickle",
		accessToken,
		lst,
		nil,
		nil,
		"HMAC-SHA256",
		"",
	)
}

func SubscribeWS(config *AppConfig, accessToken, lst string) (*http.Response, error) {
	return sendOAuthRequest(
		config,
		"GET",
		"https://api.ibkr.com/v1/api/ws",
		accessToken,
		lst,
		nil,
		nil,
		"HMAC-SHA256",
		"",
	)
}

// https://api.ibkr.com/v1/api/oauth/access_token
func GetAccessTokenFromCurrentSession(config *AppConfig, accessToken, lst string) (*http.Response, error) {
	return sendOAuthRequest(
		config,
		"GET",
		"https://api.ibkr.com/v1/api/oauth/access_token",
		accessToken,
		lst,
		nil,
		nil,
		"HMAC-SHA256",
		"",
	)
}

func GetPortfolioAccounts(config *AppConfig, accessToken, lst string) (*http.Response, error) {
	return sendOAuthRequest(
		config,
		"GET",
		"https://api.ibkr.com/v1/api/portfolio/accounts",
		accessToken,
		lst,
		nil,
		nil,
		"HMAC-SHA256",
		"",
	)
}

func GetPositionFromAccountId(accountId string, config *AppConfig, accessToken, lst string) (*http.Response, error) {
	// https://api.ibkr.com/v1/api/portfolio/{accountId}/positions/{pageId}

	return sendOAuthRequest(
		config,
		"GET",
		fmt.Sprintf("https://api.ibkr.com/v1/api/portfolio/%s/positions/0", accountId),
		accessToken,
		lst,
		nil,
		nil,
		"HMAC-SHA256",
		"",
	)
}
