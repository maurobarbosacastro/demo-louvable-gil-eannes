package controllers

import (
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	shopifyService "ms-tagpeak/external/shopify"
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/http_client"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"net/http"
	"strings"
)

func CheckIfShopExists(c echo.Context) error {
	logster.StartFuncLogMsg(c.Param("url"))
	shop, err := shopifyService.GetShopByUrl(c.Param("url"))

	if err != nil || shop == nil {
		logster.Error(err, "Error getting shop or shop not found")
		return c.JSON(http.StatusNotFound, response_object.ShopifyShopExistRO{
			Uuid:   nil,
			Exists: false,
		})
	}

	logster.EndFuncLogMsg("Found shop")
	return c.JSON(http.StatusOK, response_object.ShopifyShopExistRO{
		Uuid:   utils.StringPointer(shop.Uuid.String()),
		Exists: true,
	})
}

func CreateUserShopifyShop(c echo.Context) error {
	logster.StartFuncLog()

	var body dto.CreateShopDTO

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// Validate the input (optional, requires validation setup)
	if err := c.Validate(body); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// 1 - Create user on keycloak
	keycloak := auth.KeycloakInstance

	bodyUserDTO := dto.CreateUserDto{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Country:   body.Country,
		Currency:  body.Currency,
		Password:  body.Password,
		IsShop:    gocloak.BoolP(true),
	}

	userKeycloak, err := service.CreateUser(bodyUserDTO, keycloak)
	if err != nil {
		if apiError, ok := err.(*gocloak.APIError); ok {
			switch apiError.Code {
			case http.StatusConflict:
				logster.Error(err, fmt.Sprintf("Error creating user: %v", apiError.Message))
				logster.EndFuncLog()
				return c.JSON(http.StatusConflict, apiError.Message)
			default:
				logster.Error(err, "Error creating user")
				logster.EndFuncLog()
				return c.JSON(apiError.Code, apiError.Message)
			}
		} else {
			logster.Error(err, "Error creating user")
			logster.EndFuncLog()
			// Generic error handling
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	logster.Info(fmt.Sprintf("User created: %v", userKeycloak.Uuid))

	// 2 - Create shop on ms-shopify and ms-tagpeak and link to the user
	shopifyShop, err := service.CreateShopifyShop(body, *userKeycloak)

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, shopifyShop)
}

func RelayMSShopify(c echo.Context) error {
	logster.StartFuncLog()
	// Get the original request
	req := c.Request()

	// Extract the path that should be forwarded
	pathToForward := strings.TrimPrefix(req.URL.Path, "/shopify/relay")

	// Create a new request to the Shopify service
	shopifyURL := dotenv.GetEnv("MS_SHOPIFY_URL") + pathToForward
	proxyReq, err := http.NewRequest(req.Method, shopifyURL, req.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create proxy request"})
	}

	// Copy headers
	for key, values := range req.Header {
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}
	proxyReq.Header.Del("Authorization")
	proxyReq.Header.Add("Authorization", http_client.GetInternalKeycloakToken())

	proxyReq.URL.RawQuery = req.URL.RawQuery

	// Send the proxied request
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "Shopify service unavailable"})
	}
	defer resp.Body.Close()

	// Copy the response from MS-shopify back to the client
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read response"})
	}

	// Copy response headers
	for key, values := range resp.Header {
		if key == "Access-Control-Allow-Origin" || key == "Access-Control-Allow-Methods" ||
			key == "Access-Control-Allow-Headers" || key == "Access-Control-Allow-Credentials" {
			continue
		}
		for _, value := range values {
			c.Response().Header().Add(key, value)
		}
	}

	logster.EndFuncLog()
	return c.Blob(resp.StatusCode, resp.Header.Get("Content-Type"), responseBody)
}

func CreateShopifyStore(c echo.Context) error {
	logster.StartFuncLog()

	var body dto.CreateShopifyStoreDTO

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate the input (optional, requires validation setup)
	if err := c.Validate(body); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	logster.Debug(fmt.Sprintf("shopify dto: %v", body))

	// create store entity
	store, err := service.CreateShopifyStore(body)
	if err != nil {
		logster.Error(err, "Error creating store")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Update shopify shop with store uuid
	updateShopifyShopDto := dto.UpdateShopifyShopDTO{
		StoreUuid: &store.Uuid,
	}

	shopifyShopUuid := uuid.MustParse(body.ShopUuid)
	err = service.UpdateShopifyShop(shopifyShopUuid, updateShopifyShopDto)
	if err != nil {
		logster.Error(err, "Error updating shopify shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, store)
}

func GetShopifyStore(c echo.Context) error {
	logster.StartFuncLogMsg(fmt.Sprintf("GetShopifyStore id: %s", c.Param("id")))
	uuidShop := uuid.MustParse(c.Param("id"))

	logster.Debug(fmt.Sprintf("GetShopifyStore id: %v", uuidShop))
	shop, err := service.GetShopifyShopByUuid(uuidShop)

	if err != nil {
		logster.Error(err, "Error getting shopify shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusNotFound, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, shop)
}

func GetShopifyStoreByUser(c echo.Context) error {
	logster.StartFuncLogMsg(fmt.Sprintf("Id: %s", c.Param("id")))
	uuidUser := uuid.MustParse(c.Param("id"))

	shop, err := service.GetShopifyShopByUserUuid(uuidUser)

	if err != nil {
		logster.Error(err, "Error getting shopify shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusNotFound, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, shop)
}

func GetShopifyStoreStats(c echo.Context) error {
	id := c.Param("id")
	idUuid := uuid.MustParse(id)

	res, err := service.GetShopifyStoreStats(idUuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res.Currency = auth.Principal.User.Currency

	return c.JSON(http.StatusOK, res)
}

func UpdateShopifyStore(c echo.Context) error {
	logster.StartFuncLog()
	shopUuid := uuid.MustParse(c.Param("id"))

	shopifyStore, err := service.GetShopifyShopByUuid(shopUuid)

	if err != nil {
		logster.Error(err, "Error getting shopify shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusNotFound, err)
	}

	updateShopifyShopDto := dto.UpdateShopifyShopDTO{
		StoreUuid: shopifyStore.StoreUuid,
	}

	err = service.UpdateShopifyShop(shopUuid, updateShopifyShopDto)
	if err != nil {
		logster.Error(err, "Error updating shopify shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()

	return c.JSON(http.StatusOK, nil)
}

func CheckIfUserIsOwnerOfShop(c echo.Context) error {
	logster.StartFuncLog()
	shopDomain := c.QueryParam("domain")

	shop, err := shopifyService.GetShopByUrl(shopDomain)

	if err != nil {
		logster.Error(err, "Error getting shop from ms-shopify")
		logster.EndFuncLog()
		return c.JSON(http.StatusNotFound, "Domain not found")
	}

	shopifyShop, err := service.GetShopifyShopByUuid(shop.Uuid)
	if err != nil {
		logster.Error(err, "Error getting shopify shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusNotFound, err)
	}

	if shopifyShop.UserUuid != auth.Principal.User.Uuid {
		logster.EndFuncLogMsg("User is not owner of this shop")
		return c.JSON(http.StatusConflict, "User is not owner of this shop")
	}

	logster.EndFuncLogMsg("User is owner of this shop")
	return c.JSON(http.StatusOK, "User is owner of this shop")
}
