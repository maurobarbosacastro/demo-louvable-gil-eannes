package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"ms-tagpeak/internal/auth"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	_ "ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"
)

// GetTransaction godoc
// @Summary Get Transaction by ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction id"
// @Success 200 {object} response_object.CashbackDetailRO
// @Router /transaction/:id [get]
func GetTransaction(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetTransaction(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	resp := utils.MapTransactionToCashbackDetailRO(*res)

	return c.JSON(http.StatusOK, resp)
}

// GetRewardTransaction godoc
// @Summary Get Rewards transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param admin query string false "Admin"
// @Param id path string true "Transaction id"
// @Success 200 {object} response_object.RewardDetailRO
// @Router /transaction/:id/reward [get]
func GetRewardTransaction(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	admin := c.QueryParam("admin")
	user := c.Get("user").(*models.User).Uuid

	if admin == "true" {
		res, _ := service.GetTransaction(uuid)
		user = utils.ParseIDToUUID(res.User)
	}

	res, err := service.GetRewardByTransactionAndUser(uuid, user)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	returnResponse := response_object.RewardDetailRO{}.FromReward(*res)
	keycloak := auth.KeycloakInstance
	transaction, _ := service.GetTransaction(res.TransactionUUID)
	userTransaction, _ := service.GetUserById(transaction.User, keycloak)
	userMemberShipStatus := service.GetMembershipStatus(utils.StringPointer(userTransaction.Uuid.String()), userTransaction.Groups)
	percentOnTransaction := *userMemberShipStatus.PercentageOnTransaction

	amount := decimal.NewFromFloat(transaction.AmountUser)
	//Divided by 10000 because we get 50% so divide by 100 to get 0.5% then by 100 remove percentages
	minPercentage := decimal.NewFromFloat(percentOnTransaction / 10000)
	minAmount := amount.Mul(minPercentage)
	rewardAmount, _ := minAmount.Round(2).Float64()

	returnResponse.MinimumReward = rewardAmount

	return c.JSON(http.StatusOK, returnResponse)
}

// GetAllTransactions godoc
// @Summary Get all Transactions
// @Tags Transaction
// @Accept json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Param filters query dto.TransactionFiltersDTO true "Filters for transaction"
// @Produce json
// @Success 200 {object} pagination.PaginationResult{data=[]response_object.CashbackRO} "Array of Transactions"
// @Router /transaction [get]
func GetAllTransactions(c echo.Context) error {
	logster.StartFuncLog()

	var pag pagination.PaginationParams
	var filters dto.TransactionFiltersDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		logster.Error(errPag, "Error binding pagination")
		return c.JSON(http.StatusInternalServerError, errPag)
	}

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		logster.Error(errFilters, "Error binding filters")
		return errFilters
	}

	keycloak := auth.KeycloakInstance

	res, err := service.GetCashbackView(pag, filters, keycloak)
	if err != nil {
		logster.Error(err, "Error getting transactions")
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, res)
}

// CreateTransaction godoc
// @Summary Create Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param country body dto.CreateTransactionDTO true "Create Transaction dto"
// @Success 201 {object} response_object.CashbackRO "Transaction"
// @Router /transaction [post]
func CreateTransaction(c echo.Context) error {
	logster.StartFuncLog()
	var model dto.CreateTransactionDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&model); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	storeVisit, err := service.GetStoreVisit(*model.StoreVisitUUID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.CreateTransaction(model, storeVisit, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusCreated, res)

}

// PatchTransaction godoc
// @Summary Update Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction id"
// @Param Transaction body dto.UpdateTransactionDTO true "Update Transaction dto"
// @Success 200 {object} models.Transaction "Transaction"
// @Router /transaction/:id [patch]
func PatchTransaction(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var model dto.UpdateTransactionDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := service.UpdateTransaction(model, uuid, c.Get("user").(*models.User), true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, res)

}

// DeleteTransaction godoc
// @Summary Delete Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction id"
// @Success 204
// @Router /transaction/:id [delete]
func DeleteTransaction(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteTransaction(uuid, uuidUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

// GetStoreVisitByTransaction godoc
// @Summary Get AffiliatePartnerByTransaction by Transaction ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction id"
// @Success 200 {object} models.Partner "AffiliatePartnerByTransaction"
// @Router /transaction/:id/partner [get]
func GetStoreVisitByTransaction(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetStoreVisitByTransaction(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetStoreByTransaction godoc
// @Summary Get CountryByTransaction by Transaction ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction id"
// @Success 200 {object} dto.CountryDTO "CountryByTransaction"
// @Router /transaction/:id/country [get]
func GetStoreByTransaction(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetStoreByTransaction(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetCurrencyExchangeRateByTransaction godoc
// @Summary Get LanguageByTransaction by Transaction ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction id"
// @Success 200 {object} dto.LanguageDTO "LanguageByTransaction"
// @Router /transaction/:id/language [get]
func GetCurrencyExchangeRateByTransaction(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetCurrencyExchangeRateByTransaction(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetMyTransactions godoc
// @Summary Get User Transactions
// @Tags Transaction
// @Accept json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Param filters query dto.TransactionFiltersDTO true "Filters for transaction"
// @Produce json
// @Success 200 {array} response_object.CashbackRO "Array of Transactions"
// @Router /transaction/me [get]
func GetMyTransactions(c echo.Context) error {

	var pag pagination.PaginationParams
	var filters dto.TransactionFiltersDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return c.JSON(http.StatusInternalServerError, errPag)
	}

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		return errFilters
	}

	userUuid := c.Get("user").(*models.User).Uuid.String()

	filters.User = &userUuid

	keycloak := auth.KeycloakInstance

	res, err := service.GetCashbackView(pag, filters, keycloak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// BulkEditTransaction godoc
// @Summary Bulk Edit transaction
// @Description Bulk Edit transaction
// @Tags Transaction
// @Accept  json
// @Param Transaction body dto.UpdateTransactionDTO true "Update Transaction dto"
// @Success 200
// @Router /transaction/bulk/edit [patch]
func BulkEditTransaction(c echo.Context) error {

	var model dto.UpdateTransactionDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if model.Uuids == nil || len(*model.Uuids) == 0 {
		return utils.CustomErrorStruct{}.BadRequestError("uuids is required")
	}

	err := service.BulkUpdateTransaction(model, c.Get("user").(*models.User))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	//Set store visit purchase flag to true if transaction state is VALIDATED
	if model.State != nil && *model.State == "VALIDATED" {
		logster.Info("Setting store visits purchased to true")
		errFlag := service.BulkSetPurchasedStoreVisitsByTransactions(
			*model.Uuids,
		)

		if errFlag != nil {
			logster.Error(errFlag, "Error setting store visits purchased to true")
		}

	}

	for _, uuid := range *model.Uuids {
		err := service.CalculateAndSetCashback(uuid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, nil)
}

// GetAllShopTransactions godoc
// @Summary Get all Transactions for a shop
// @Tags Transaction
// @Accept json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Param filters query dto.TransactionFiltersDTO true "Filters for transaction"
// @Produce json
// @Success 200 {object} pagination.PaginationResult{data=[]response_object.CashbackViewRO} "Array of Transactions"
// @Router /:id/transaction [get]
func GetAllShopTransactions(c echo.Context) error {
	logster.StartFuncLog()

	var pag pagination.PaginationParams
	var filters dto.TransactionFiltersDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		logster.Error(errPag, "Error binding pagination")
		return c.JSON(http.StatusInternalServerError, errPag)
	}

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		logster.Error(errFilters, "Error binding filters")
		return errFilters
	}

	keycloak := auth.KeycloakInstance
	id := c.Param("id")

	filters.StoreUuid = &id

	res, err := service.GetCashbackView(pag, filters, keycloak)
	if err != nil {
		logster.Error(err, "Error getting transactions")
		return c.JSON(http.StatusInternalServerError, err)
	}
	res.Data = lo.Map(res.Data.([]response_object.CashbackRO), func(c response_object.CashbackRO, _ int) response_object.CashbackRO {
		return c.Shopify()
	})

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, res)
}
