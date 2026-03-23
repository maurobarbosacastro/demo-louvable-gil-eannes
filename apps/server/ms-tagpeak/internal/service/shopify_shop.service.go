package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"ms-tagpeak/external/shopify"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/dto/webhooks"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/dotenv"
	keycloak_utils "ms-tagpeak/pkg/keycloak"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

func CreateShopifyShop(dto dto.CreateShopDTO, user models.User) (models.ShopifyShop, error) {
	logster.StartFuncLog()

	logster.Info("Creating shop on ms-shopify")
	// 1 - Create shop on ms-shopify
	body := shopify.CreateMSShopifyShopDTO{
		Shop: dto.Shop,
		User: user.Uuid.String(),
	}
	shop, err := shopify.CreateShop(body)

	if err != nil {
		logster.Error(err, "Error creating shop on ms-shopify")
		logster.EndFuncLog()
		return models.ShopifyShop{}, err
	}
	logster.Info(fmt.Sprintf("Shopify shop: %v+", shop))

	logster.Info("Creating shopify shop on ms-tagpeak")
	// 2 - Create shopify shop on ms-tagpeak linking in to the user
	shopifyShopBody := models.ShopifyShop{
		ShopUuid: shop.Uuid,
		UserUuid: user.Uuid,
	}

	shopifyShop, err := repository.CreateShopifyShop(shopifyShopBody)

	if err != nil {
		logster.Error(err, "Error creating shopify shop")
		logster.EndFuncLog()
		return models.ShopifyShop{}, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Shopify shop created: %v+", shopifyShop))
	return shopifyShop, nil
}

func GetShopifyShopByUuid(uuid uuid.UUID) (*models.ShopifyShop, error) {
	shopifyShop, err := repository.GetShopifyShopByUuid(uuid)

	if err != nil {
		logster.Error(err, "Error getting shopify shop")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Shopify shop: %v+", shopifyShop))
	return shopifyShop, nil
}

func GetShopifyShopByUserUuid(uuid uuid.UUID) (*models.ShopifyShop, error) {
	shopifyShop, err := repository.GetShopifyShopByUserUuid(uuid)

	if err != nil {
		logster.Error(err, "Error getting shopify shop")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Shopify shop: %v+", shopifyShop))
	return shopifyShop, nil
}

func GetShopifyShopByStoreUuid(uuid uuid.UUID) (*models.ShopifyShop, error) {
	shopifyShop, err := repository.GetShopifyShopByStoreUuid(uuid)

	if err != nil {
		logster.Error(err, "Error getting shopify shop")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Shopify shop: %v+", shopifyShop))
	return shopifyShop, nil
}

func CreateShopifyStore(body dto.CreateShopifyStoreDTO) (models.Store, error) {

	shop, _ := repository.GetShopifyShopByUuid(
		uuid.MustParse(body.ShopUuid),
	)

	if shop != nil && shop.StoreUuid != nil {
		return repository.GetStore(*shop.StoreUuid)
	}

	storeDto := models.Store{
		Name:                        body.Name,
		Logo:                        nil,
		ShortDescription:            body.Description,
		Description:                 utils.StringPointer(""),
		UrlSlug:                     utils.StringPointer(""),
		AverageRewardActivationTime: utils.StringPointer(fmt.Sprintf("%d days", body.ReturnPeriod)),
		State:                       utils.StringPointer("PENDING"),
		Keywords:                    utils.StringPointer(""),
		AffiliateLink:               utils.StringPointer(""),
		StoreUrl:                    utils.StringPointer(body.Url),
		TermsAndConditions:          utils.StringPointer(""),
		CashbackType:                utils.StringPointer("Percentage"),
		CashbackValue:               utils.FloatPointer(body.Percentage),
		PercentageCashout:           utils.FloatPointer(4),
		MetaTitle:                   utils.StringPointer(""),
		MetaKeywords:                utils.StringPointer(""),
		MetaDescription:             utils.StringPointer(""),
		Country:                     nil,
		Category:                    nil,
		OverrideFee:                 nil,
		PartnerIdentity:             utils.StringPointer(""),
		LanguageCODE:                nil,
		AffiliatePartnerCODE:        utils.StringPointer("shopify"),
		CountriesCodes:              body.Countries,
		CategoriesCodes:             body.Categories,
	}

	return repository.CreateStore(storeDto)
}

func UpdateShopifyShop(uuid uuid.UUID, data dto.UpdateShopifyShopDTO) error {
	return repository.SetStoreToShopifyShop(uuid, *data.StoreUuid)
}

func HandleWebhookCreateOrUpdate(data webhooks.ShopifyOrder, transaction *models.Transaction, keycloak *constants.Keycloak, storeDomain string) (*dto.CamundaCreateTransactionDTO, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Data: %+v | Transaction: %+v", data, transaction))
	keycloak_utils.RefreshKeycloakAdminToken(keycloak)

	response := dto.CamundaCreateTransactionDTO{
		SourceId:       strconv.FormatInt(data.ID, 10),
		State:          models.TransactionStateTracked,
		CurrencySource: data.Currency,
		Reference:      "",
	}

	//Get discount code index to get items applicable
	_, discountCodeIndex, _ := lo.FindIndexOf(data.DiscountApplications, func(code webhooks.DiscountApplication) bool {
		return code.Code == dotenv.GetEnv("SHOPIFY_TAGPEAK_CODE")
	})
	logster.Debug(fmt.Sprintf("Discount code index: %d", discountCodeIndex))

	//Filter applicable items to discount code
	items := lo.Filter(data.LineItems, func(item webhooks.LineItem, _ int) bool {
		_, found := lo.Find(item.DiscountAllocations, func(alloc webhooks.DiscountAllocation) bool {
			return alloc.DiscountApplicationIndex == discountCodeIndex
		})
		return found
	})
	logster.Debug(fmt.Sprintf("Items: %+v", items))

	//Transaction is null so it's a new transaction, no items applicable found. stop processing.
	if transaction == nil && len(items) == 0 {
		return nil, errors.New("No items found with discount code")
	}

	//Calculate total amount items
	response.AmountSource = lo.SumBy(items, func(item webhooks.LineItem) float64 {
		price, _ := strconv.ParseFloat(item.Price, 64)
		return price * float64(item.CurrentQuantity)
	})
	logster.Debug(fmt.Sprintf("Total amount items: %f", response.AmountSource))

	//Transaction is not null so it's an update if the amount is 0, items where removed
	//Set state to rejected
	if transaction != nil && response.AmountSource == 0 {
		response.State = "rejected"
	}

	//Get customer order email and check if user exists
	customerEmail := data.Customer.Email
	keycloakUser, err := GetUsers(keycloak, utils.StringPointer(customerEmail), 1)
	if err != nil {
		logster.Error(err, "Error getting user")
		return nil, err
	}
	logster.Debug(fmt.Sprintf("Keycloak user: %+v", keycloakUser))

	isNewUser := false

	if len(keycloakUser) == 0 {
		logster.Info("User not found, creating user")

		userDto := dto.CreateUserDto{
			FirstName:     data.Customer.FirstName,
			LastName:      data.Customer.LastName,
			Email:         data.Customer.Email,
			Password:      base64.StdEncoding.EncodeToString([]byte(utils.RandomWordsCode(16))),
			Currency:      "EUR",
			Country:       nil,
			UtmParams:     nil,
			ReferralCode:  nil,
			ReferralClick: nil,
			IsShop:        utils.BoolPointer(false),
		}

		userCreated, err := CreateUserFromShopifyOrder(userDto, keycloak)
		if err != nil {
			logster.Error(err, "Error creating user")
			logster.EndFuncLog()
			return nil, err
		}
		response.UserUUID = userCreated.Uuid
		isNewUser = true

	} else {
		logster.Info("User found")
		response.UserUUID = keycloakUser[0].Uuid
	}
	logster.Info(fmt.Sprintf("User %+v", response.UserUUID))

	//Find Store and create store visit
	shop, _ := shopify.GetShopByUrl(storeDomain)
	logster.Debug(fmt.Sprintf("Shop %+v", shop))
	shopifyShop, _ := repository.GetShopifyShopByUuid(shop.Uuid)
	logster.Debug(fmt.Sprintf("ShopifyShop %+v", shopifyShop))
	store, _ := GetStore(*shopifyShop.StoreUuid)

	storeOwner, errOwner := GetUserById(shopifyShop.UserUuid.String(), keycloak)
	if errOwner == nil {
		response.CurrencySource = storeOwner.Currency
	}

	if isNewUser {
		//Update in coroutine user to add email extra data and send verify email
		go func() {
			dtoUser := dto.UpdateUserDto{
				EmailExtras: utils.Ptr(map[string]string{"storeName": store.Name}),
			}
			_, errUpdateUser := UpdateUser(response.UserUUID, dtoUser, keycloak)
			if errUpdateUser != nil {
				logster.Error(errUpdateUser, "Error updating user")
				logster.EndFuncLog()
				return
			}

			errEmail := keycloak.Client.SendVerifyEmail(keycloak.Ctx, keycloak.AdminToken.AccessToken, response.UserUUID.String(), keycloak.Realm)

			if errEmail != nil {
				logster.Error(errEmail, "Error sending email")
			}
			logster.EndFuncLog()
		}()
	}

	//If new order, create store visit, if not, get store visit uuid from transaction.
	//In case of update, we check if the average reward time has passed, if so, we don't update the transaction.
	if transaction == nil {
		storeVisitDto := dto.CreateStoreVisitDTO{
			User:      response.UserUUID.String(),
			Reference: data.Name,
			Purchase:  true,
			StoreUUID: *shopifyShop.StoreUuid,
		}
		storeVisit, _ := CreateStoreVisit(storeVisitDto, response.UserUUID.String())
		response.StoreVisitUUID = storeVisit.Uuid
		logster.Debug(fmt.Sprintf("StoreVisit %+v", storeVisit.Uuid))
	} else {
		storeVisit, _ := repository.GetStoreVisit(*transaction.StoreVisitUUID)
		response.StoreVisitUUID = &storeVisit.Uuid
		logster.Debug(fmt.Sprintf("StoreVisit %+v", storeVisit.Uuid))

		//Check if store average_reward_activation_time has passed, if so, do not update the transaction.
		orderDate := transaction.OrderDate
		averageRewardActivationTime, _ := utils.ParseDaysString(store.AverageRewardActivationTime)
		if orderDate.Add(averageRewardActivationTime).Before(time.Now()) {
			logster.Error(nil, fmt.Sprintf("Store avereage reward activation time has passed"))
			logster.EndFuncLog()
			return nil, errors.New(fmt.Sprintf("Store avereage reward activation time has passed"))
		}
	}

	//Calculate values of transaction
	cashbackValue := decimal.NewFromFloat(store.CashbackValue / 100)
	decimalTotalAmountItems := decimal.NewFromFloat(response.AmountSource)
	commissionSource, _ := cashbackValue.Mul(decimalTotalAmountItems).Round(2).Float64()
	response.CommissionSource = commissionSource

	timeStandard, _ := time.Parse(time.RFC3339, data.CreatedAt)
	response.OrderDate = timeStandard.UTC()

	logster.EndFuncLogMsg(fmt.Sprintf("Response: %+v", response))
	return &response, nil
}

func StartCamundaProcessForShopifyOrder(transactionCamunda *dto.CamundaCreateTransactionDTO) {
	trx := map[string]interface{}{
		"loopCounter":  1,
		"transactions": [1]interface{}{transactionCamunda},
	}

	resp := camundaPKG.StartProcessInstance(
		camundaPKG.InjectEnvOnKey("sub-process-transaction"),
		*camundaPKG.GetCamundaClient(),
		trx,
	)
	logster.Info(fmt.Sprintf("Process started -Process Instance Key : %d", resp.ProcessInstanceKey))
}

func GetShopifyStoreStats(uuid uuid.UUID) (*models.ShopStats, error) {
	logster.StartFuncLog()

	res, err := repository.GetShopifyStoreStats(uuid.String())
	if err != nil {
		logster.Error(err, "Error getting shopify store stats")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return &res, nil
}
