package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"ms-shopify/internal/dto"
	"ms-shopify/internal/dto/webhooks"
	"ms-shopify/internal/response_object"
	"ms-shopify/internal/service"
	"ms-shopify/internal/service/redis"
	"ms-shopify/pkg/dotenv"
	"ms-shopify/pkg/http_client"
	"ms-shopify/pkg/logster"
	"ms-shopify/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

// GetShopByUrl godoc
// @Summary Get GetShopByUrl
// @Description Check if a shop exists by it's url
// @Tags Users -> Groups in swagger
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Found!"
// @Failure 404 {string} string "Shop not found"
// @Router /shop [get]
func GetShopByUrl(c echo.Context) error {
	logster.StartFuncLogMsg(fmt.Sprintf("Shop URL: %s", c.QueryParam("url")))

	shop, err := service.GetShopByUrl(c.QueryParam("url"))

	if err != nil {
		logster.Error(err, "404 - Error checking if shop exists")
		logster.EndFuncLog()
		return c.JSON(http.StatusNotFound, err)
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Exists: %t", shop))
	return c.JSON(http.StatusOK, shop)
}

func CreateShop(c echo.Context) error {
	logster.StartFuncLog()

	body := dto.CreateShopifyShopDTO{}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	logster.Info(fmt.Sprintf("Creating shop entity for: %s", body.Shop))

	shop, err := service.CreateShop(body)

	if err != nil {
		logster.Error(err, "Error creating shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusCreated, shop)
}

func UpdateShop(c echo.Context) error {
	logster.StartFuncLog()
	uuid := utils.ParseIDToUUID(c.Param("id"))

	//get token from header
	token := c.Request().Header.Get("X-TP-Shopify-Token")

	body := dto.UpdateShopifyShopDTO{
		AccessToken: &token,
	}

	updatedShop, err := service.UpdateShop(uuid, body)

	if err != nil {
		logster.Error(err, "Error updating shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, updatedShop)
}

func SetupShopify(c echo.Context) error {
	shop := c.Request().Header.Get("X-Shopify-Shop")
	logster.StartFuncLogMsg(fmt.Sprintf("Shop: %v", shop))
	shopToken := c.Request().Header.Get("X-TP-Shopify-Token")

	productsIds := dto.InstallSetupDto{}

	if err := c.Bind(&productsIds); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//Updating token is optional, so we do it in a go routine
	go service.UpdateTokenByShopUrl(shop, shopToken)

	shopModel, err := service.GetShopByUrl(shop)
	if err != nil {
		logster.Error(err, "Error getting shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := response_object.SetupRO{
		AlreadyExists:  false,
		DiscountCodeId: "",
		Shop:           *shopModel,
	}

	service.ShopifyClient.CreateClient(shop, shopToken)

	code, errGetCode := service.ShopifyClient.GetDiscountCode(dotenv.GetEnv("SHOPIFY_TAGPEAK_CODE"))

	if errGetCode != nil {
		logster.Error(errGetCode, "Error getting discount code")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, errGetCode)
	}

	if code.CodeDiscountNodeByCode.ID == "" {

		collectionId, errCollectionId := service.GetCollectionId()
		if errCollectionId != nil {
			return c.JSON(http.StatusInternalServerError, errCollectionId)
		}
		logster.Info(fmt.Sprintf("Collection ID: %s", *collectionId))

		//Publish collection to all publishables
		errPublishing := service.PublishCollection(*collectionId)
		if errPublishing != nil {
			return c.JSON(http.StatusInternalServerError, errPublishing)
		}

		//Associate products to the collection if needed
		if len(*productsIds.Products) > 0 {
			errProductsAdded := service.AddProductsToCollection(*collectionId, *productsIds.Products)
			if errProductsAdded != nil {
				logster.Error(errProductsAdded, "Error adding products to collection")
				logster.EndFuncLog()
				//No need to end the flow here
			}
		}

		//Check if code exists, if not create it
		tagpeakDiscountCode, err := service.ShopifyClient.GetDiscountCode(dotenv.GetEnv("SHOPIFY_TAGPEAK_CODE"))

		if err != nil {
			logster.Error(err, "Error getting discount code")
			logster.EndFuncLog()
			return c.JSON(http.StatusInternalServerError, err)
		}

		if tagpeakDiscountCode.CodeDiscountNodeByCode.ID != "" {
			return c.JSON(http.StatusCreated, tagpeakDiscountCode)
		}

		codeCreated, errCodeCreation := service.HandleDiscountCodeCreation(*collectionId)

		if errCodeCreation != nil {
			return c.JSON(http.StatusInternalServerError, errCodeCreation)
		}

		response.AlreadyExists = false
		response.DiscountCodeId = codeCreated.DiscountCodeBasicCreate.CodeDiscountNode.ID

		go func() {
			logster.Info("Running go routine to update shop installation flag")
			_, err := service.UpdateShop(shopModel.Uuid, dto.UpdateShopifyShopDTO{InstallationDone: utils.BoolPointer(true)})
			if err != nil {
				logster.Error(err, "Error setting installation done for shop "+shopModel.Url)
			}
		}()

		return c.JSON(http.StatusCreated, response)
	} else {
		response.AlreadyExists = true
		response.DiscountCodeId = code.CodeDiscountNodeByCode.ID
		logster.EndFuncLog()

		go func() {
			logster.Info("Running go routine to update shop installation flag")
			_, err := service.UpdateShop(shopModel.Uuid, dto.UpdateShopifyShopDTO{InstallationDone: utils.BoolPointer(true)})
			if err != nil {
				logster.Error(err, "Error setting installation done for shop "+shopModel.Url)
			}
		}()

		return c.JSON(http.StatusOK, response)
	}
}

// ActivateShop godoc
// @Summary Activate a shop
// @Description Change the shop state to ACTIVE
// @Tags Shop
// @Accept  json
// @Produce  json
// @Param id path string true "Shop UUID"
// @Success 200 {object} models.Shop
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Shop not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/{id}/activate [post]
func ActivateShop(c echo.Context) error {
	logster.StartFuncLog()
	uuid := utils.ParseIDToUUID(c.Param("id"))

	// Set the state to ACTIVE
	activeState := "ACTIVE"
	body := dto.UpdateShopifyShopDTO{
		State: &activeState,
	}

	updatedShop, err := service.UpdateShop(uuid, body)

	if err != nil {
		logster.Error(err, "Error activating shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, updatedShop)
}

// DeactivateShop godoc
// @Summary Deactivate a shop
// @Description Change the shop state to CLOSED
// @Tags Shop
// @Accept  json
// @Produce  json
// @Param id path string true "Shop UUID"
// @Success 200 {object} models.Shop
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Shop not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/{id}/deactivate [post]
func DeactivateShop(c echo.Context) error {
	logster.StartFuncLog()
	uuid := utils.ParseIDToUUID(c.Param("id"))

	// Set the state to CLOSED
	closedState := "CLOSED"
	body := dto.UpdateShopifyShopDTO{
		State: &closedState,
	}

	updatedShop, err := service.UpdateShop(uuid, body)

	if err != nil {
		logster.Error(err, "Error deactivating shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, updatedShop)
}

func HandleWebhook(c echo.Context) error {
	logster.StartFuncLog()
	topic := c.Request().Header.Get("x-shopify-topic")
	domain := c.Request().Header.Get("x-shopify-shop-domain")
	logster.Info(fmt.Sprintf("Webhook received - Topic: %s - Domain: %s", topic, domain))

	shop, err := service.GetShopByUrl(domain)
	if err != nil {
		logster.Error(err, "Error getting shop")
		logster.EndFuncLog()
		//Sending response error to be able to receive later on
		return c.JSON(http.StatusInternalServerError, err)
	}
	if shop.State != "ACTIVE" {
		logster.EndFuncLogMsg(fmt.Sprintf("Shop is not active, skipping webhook for shop %s", domain))
		return c.JSON(http.StatusOK, true)
	}

	shop, err = service.GetShopByUrl(domain)
	if err != nil {
		logster.Error(err, "Error getting shop")
		logster.EndFuncLog()
		//Sending response error to be able to receive later on
		return c.JSON(http.StatusInternalServerError, err)
	}
	if shop.State != "ACTIVE" {
		logster.EndFuncLogMsg(fmt.Sprintf("Shop is not active, skipping webhook for shop %s", domain))
		return c.JSON(http.StatusOK, true)
	}

	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to read body")
	}

	if topic == "orders/create" || topic == "orders/updated" || topic == "orders/cancelled" {
		c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		var shopifyOrder webhooks.ShopifyOrder
		if err := c.Bind(&shopifyOrder); err != nil {
			logster.Error(err, "Error binding shopify order")
			return c.JSON(http.StatusBadRequest, err) //Only return error if we want to process the order later.
		}
		logster.Debug(fmt.Sprintf("ShopifyOrder: %+v", shopifyOrder))

		_, found := lo.Find(shopifyOrder.DiscountCodes, func(item webhooks.DiscountCode) bool {
			return item.Code == dotenv.GetEnv("SHOPIFY_TAGPEAK_CODE")
		})

		if !found {
			logster.EndFuncLogMsg(fmt.Sprintf("No discount code found for order id %d (%s)", shopifyOrder.ID, shopifyOrder.Name))
			return c.JSON(http.StatusOK, false)
		}

		//check if customer.email exists, if not, check graphql with customer.AdminGraphqlAPIID
		if shopifyOrder.Customer.Email == "" {
			logster.Info(fmt.Sprintf("Customer email is empty, getting it from graphql for customer %s", shopifyOrder.Customer.AdminGraphqlAPIID))

			offlineToken := service.GetShopOfflineToken(domain)
			if offlineToken == nil || *offlineToken == "" {
				logster.EndFuncLog()
				return c.JSON(http.StatusInternalServerError, "Error getting offline token for shop")
			}

			service.ShopifyClient.CreateClient(domain, *offlineToken)
			email, err := service.GetCustomerEmail(shopifyOrder.Customer.AdminGraphqlAPIID)
			if err != nil {
				logster.Error(err, "Error getting customer email")
				logster.EndFuncLog()
				return c.JSON(http.StatusInternalServerError, "Error getting customer email")
			} else {
				logster.Info(fmt.Sprintf("Customer email found: %s", email))
				shopifyOrder.Customer.Email = email
			}
		}

		newBytes, _ := json.Marshal(shopifyOrder)

		// Convert to string (Redis stores strings)
		bodyStr := string(newBytes)
		redis.PushToRedisQueue(redis.ShopifyRedisQueueKey, topic, domain, bodyStr)
	}

	if topic == "orders/delete" {
		bodyStr := string(bodyBytes)
		redis.PushToRedisQueue(redis.DeleteShopifyRedisQueueKey, topic, domain, bodyStr)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, true)
}

func HandleUninstallWebhook(c echo.Context) error {
	logster.StartFuncLog()
	topic := c.Request().Header.Get("x-shopify-topic")
	domain := c.Request().Header.Get("x-shopify-shop-domain")
	logster.Info(fmt.Sprintf("Webhook Uninstall received - Topic: %s - Domain: %s", topic, domain))

	shop, err := service.GetShopByUrl(domain)
	if err != nil {
		logster.Error(err, "Error getting shop")
		logster.EndFuncLog()
		//Sending response error to be able to receive later on
		return c.JSON(http.StatusInternalServerError, err)
	}

	_, err = service.UpdateShop(shop.Uuid, dto.UpdateShopifyShopDTO{State: utils.StringPointer("CLOSED")})
	if err != nil {
		logster.Error(err, "Error updating shop")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	httpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}
	msTagpeak := dotenv.GetEnv("MS_TAGPEAK_URL")
	msTagpeakUrl := fmt.Sprintf("%s/shopify/shop/%s", msTagpeak, shop.Uuid)

	_, err = httpClient.Post(msTagpeakUrl, nil, nil, nil)

	if err != nil {
		logster.Error(err, "Error calling ms-tagpeak")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	//Remove/disable discount code
	//Remove/disable collection

	return c.JSON(http.StatusOK, true)
}
