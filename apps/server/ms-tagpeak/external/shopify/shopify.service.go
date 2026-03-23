package shopify

import (
	"fmt"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/http_client"
	"ms-tagpeak/pkg/logster"
	"net/http"
)

func Test(body dto.GetInstallShopifyDto) *string {
	msShopifyUrl := dotenv.GetEnv("MS_SHOPIFY_URL")
	url := msShopifyUrl + "hello"
	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}
	var response string

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	logster.Info(fmt.Sprintf("shopify dto: %v", body))
	_, err := internalHttpClient.Post(url, headers, body, &response)

	if err != nil {
		return nil
	}

	return &response
}

func GetShopByUrl(urlShop string) (*Shop, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("urlShop: %s", urlShop))
	msShopifyUrl := dotenv.GetEnv("MS_SHOPIFY_URL")
	url := msShopifyUrl + "shop?url=" + urlShop
	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}

	var response Shop
	httpResponse, err := internalHttpClient.Get(url, nil, &response)

	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode == 404 {
		logster.EndFuncLogMsg("Shop does not exist")
		return nil, nil
	}
	logster.EndFuncLogMsg(fmt.Sprintf("Shop does exist %+v", response))
	return &response, nil
}

func CreateShop(shopDto CreateMSShopifyShopDTO) (*Shop, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("urlShop: %s", shopDto.Shop))
	msShopifyUrl := dotenv.GetEnv("MS_SHOPIFY_URL")
	url := msShopifyUrl + "shop"
	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}

	var response Shop
	_, err := internalHttpClient.PostJson(url, nil, &shopDto, &response)

	if err != nil {
		return nil, err
	}

	logster.EndFuncLog()
	return &response, nil
}
