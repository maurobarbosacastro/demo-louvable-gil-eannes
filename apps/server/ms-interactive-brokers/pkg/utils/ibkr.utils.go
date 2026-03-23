package utils

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"ms-interactive-brokers/pkg/logster"
	"net/url"
	"sort"
	"strings"
	"time"
)

// GenerateRequestTimestamp generates the current timestamp in seconds
func GenerateRequestTimestamp() string {
	logster.Info(fmt.Sprintf("GenerateRequestTimestamp - %v\n", time.Now().Unix()))
	return fmt.Sprintf("%d", time.Now().Unix())
}

// ReadPrivateKey reads and parses a private key from file
func ReadPrivateKey(privateKeyPath string) (*rsa.PrivateKey, error) {

	pemBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		logster.Error(err, fmt.Sprintf("ReadPrivateKey: failed to read private key: %w", err))
		return nil, err
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		logster.Error(err, fmt.Sprintf("ReadPrivateKey: failed to decode PEM block: %w", err))
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		logster.Error(err, fmt.Sprintf("ReadPrivateKey: failed to parse PKCS8 private key: %w", err))
		return nil, err
	}

	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		logster.Error(err, fmt.Sprintf("ReadPrivateKey: not an RSA private key: %w", err))
		return nil, fmt.Errorf("not an RSA private key")
	}

	return rsaKey, nil
}

// GenerateOAuthNonce generates a random nonce value
func GenerateOAuthNonce() string {
	const nonceLength = 16
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, nonceLength)
	for i := range result {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[n.Int64()]
	}

	return string(result)
}

// GenerateBaseString generates the OAuth base string
func GenerateBaseString(requestMethod, requestURL string, requestHeaders map[string]string,
	requestParams, requestFormData, requestBody, extraHeaders map[string]string, prepend string) string {

	encodedURL := url.QueryEscape(requestURL)

	// Combine all parameters
	allParams := make(map[string]string)
	for k, v := range requestHeaders {
		allParams[k] = v
	}
	addParamsIfNotNil(allParams, requestParams)
	addParamsIfNotNil(allParams, requestFormData)
	addParamsIfNotNil(allParams, requestBody)
	addParamsIfNotNil(allParams, extraHeaders)

	// Create sorted parameter string
	var pairs []string
	for k, v := range allParams {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(pairs)
	paramsString := strings.Join(pairs, "&")
	encodedParams := url.QueryEscape(paramsString)

	baseString := fmt.Sprintf("%s&%s&%s", requestMethod, encodedURL, encodedParams)
	if prepend != "" {
		baseString = prepend + baseString
	}

	return baseString
}

// GenerateDHRandomBytes generates random bytes for DH challenge
func GenerateDHRandomBytes() string {
	lsh := new(big.Int).Lsh(big.NewInt(1), 256)
	n, _ := rand.Int(rand.Reader, lsh)

	return hex.EncodeToString(n.Bytes())
}

// GenerateDHChallenge generates the DH challenge
func GenerateDHChallenge(dhPrime, dhRandom string, dhGenerator int64) string {
	prime := new(big.Int)
	prime.SetString(dhPrime, 16)

	random := new(big.Int)
	random.SetString(dhRandom, 16)

	generator := big.NewInt(dhGenerator)

	challenge := new(big.Int).Exp(generator, random, prime)

	return hex.EncodeToString(challenge.Bytes())
}

// CalculateLSTPrepend calculates the LST prepend value
func CalculateLSTPrepend(accessTokenSecret string, privateKey *rsa.PrivateKey) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(accessTokenSecret)
	if err != nil {
		logster.Error(err, fmt.Sprintf("CalculateLSTPrepend: failed to decode access token secret: %w", err))
		return "", err
	}

	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		logster.Error(err, fmt.Sprintf("CalculateLSTPrepend: failed to decrypt access token secret: %w", err))
		return "", err
	}

	return hex.EncodeToString(decrypted), nil
}

// GenerateRSASHA256Signature generates RSA-SHA256 signature
func GenerateRSASHA256Signature(baseString string, privateKey *rsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write([]byte(baseString))
	digest := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest)
	if err != nil {
		logster.Error(err, fmt.Sprintf("GenerateRSASHA256Signature: failed to sign: %+v", err))
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(signature)
	return url.QueryEscape(encoded), nil
}

// GenerateHMACSHA256Signature generates HMAC-SHA256 signature
func GenerateHMACSHA256Signature(baseString, liveSessionToken string) (string, error) {

	key, err := base64.StdEncoding.DecodeString(liveSessionToken)
	if err != nil {
		logster.Error(err, fmt.Sprintf("GenerateHMACSHA256Signature: failed to decode live session token: %w", err))
		return "", err
	}

	h := hmac.New(sha256.New, key)
	h.Write([]byte(baseString))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return url.QueryEscape(signature), nil
}

// CalculateLiveSessionToken calculates the live session token
func CalculateLiveSessionToken(dhPrime, dhRandom, dhResponse, prepend string) (string, error) {
	prime := new(big.Int)
	prime.SetString(dhPrime, 16)

	random := new(big.Int)
	random.SetString(dhRandom, 16)

	response := new(big.Int)
	response.SetString(dhResponse, 16)

	K := new(big.Int).Exp(response, random, prime)
	KBytes := K.Bytes()
	if len(bin(K))%8 == 0 {
		KBytes = append([]byte{0}, KBytes...)
	}

	prependBytes, err := hex.DecodeString(prepend)
	if err != nil {
		logster.Error(err, fmt.Sprintf("CalculateLiveSessionToken: failed to decode prepend: %w", err))
		return "", err
	}

	h := hmac.New(sha1.New, KBytes)
	h.Write(prependBytes)

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

// ValidateLiveSessionToken validates the LST
func ValidateLiveSessionToken(lst, lstSignature, consumerKey string) bool {
	key, err := base64.StdEncoding.DecodeString(lst)
	if err != nil {
		logster.Error(err, fmt.Sprintf("CalculateLiveSessionToken: failed to decode lst: %w", err))
		return false
	}

	h := hmac.New(sha1.New, key)
	h.Write([]byte(consumerKey))

	return hex.EncodeToString(h.Sum(nil)) == lstSignature
}

// GenerateAuthorizationHeaderString generates the OAuth authorization header
func GenerateAuthorizationHeaderString(requestData map[string]string, realm string) string {

	var pairs []string
	for k, v := range requestData {
		pairs = append(pairs, fmt.Sprintf(`%s="%s"`, k, v))
	}
	sort.Strings(pairs)

	return fmt.Sprintf(`OAuth realm="%s", %s`, realm, strings.Join(pairs, ", "))
}

// Helper functions
func addParamsIfNotNil(target, source map[string]string) {
	if source != nil {
		for k, v := range source {
			target[k] = v
		}
	}
}

func bin(n *big.Int) string {
	return fmt.Sprintf("%b", n)
}

func GetLastPriceType(lastPrice string) string {
	switch lastPrice[:1] {
	case "C":
		return "Previous day's closing price"
	case "H":
		return "Trading has halted"
	default:
		return "Unknown"
	}
}
