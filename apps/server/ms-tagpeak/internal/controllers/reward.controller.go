package controllers

import (
	"fmt"
	"ms-tagpeak/internal/auth"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	_ "ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/internal/service"
	camundaService "ms-tagpeak/internal/service/camunda_processes"
	"ms-tagpeak/pkg/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

// GetReward godoc
// @Summary Get Reward by ID
// @Tags Reward
// @Accept json
// @Produce json
// @Param id path string true "Reward id"
// @Success 200 {object} response_object.RewardDetailRO
// @Router /reward/:id [get]
func GetReward(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetReward(uuid)

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

// GetAllRewards godoc
// @Summary Get all Rewards
// @Tags Reward
// @Accept json
// @Produce json
// @Success 200 {array} models.Reward "Array of Rewards"
// @Router /reward [get]
func GetAllRewards(c echo.Context) error {

	res, err := service.GetAllRewards()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)

}

// CreateReward godoc
// @Summary Create Reward
// @Tags Reward
// @Accept json
// @Produce json
// @Param reward body dto.CreateRewardDTO true "Create Reward dto"
// @Success 201 {object} models.Reward "Reward"
// @Router /reward [post]
func CreateReward(c echo.Context) error {
	var model dto.CreateRewardDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	model.Origin = "PURCHASE"
	if err := c.Validate(&model); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	transaction, err := service.GetTransaction(model.TransactionUUID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	uuidUser := transaction.User
	currencyUser := c.Get("user").(*models.User).Currency
	keycloak := auth.KeycloakInstance

	res, err := service.CreateReward(model, transaction, uuidUser, currencyUser, keycloak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)

}

// PatchReward godoc
// @Summary Update Reward
// @Tags Reward
// @Accept json
// @Produce json
// @Param id path string true "Reward id"
// @Param Reward body dto.UpdateRewardDTO true "Update Reward dto"
// @Success 200 {object} models.Reward "Reward"
// @Router /reward/:id [patch]
func PatchReward(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var model dto.UpdateRewardDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()
	keycloak := auth.KeycloakInstance

	res, err := service.UpdateReward(model, uuid, uuidUser, keycloak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, res)

}

// DeleteReward godoc
// @Summary Delete Reward
// @Tags Reward
// @Accept json
// @Produce json
// @Param id path string true "Reward id"
// @Success 204
// @Router /reward/:id [delete]
func DeleteReward(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteReward(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

// GetTransactionByReward godoc
// @Summary Get GetTransactionByReward by Reward ID
// @Tags Reward
// @Accept json
// @Produce json
// @Param id path string true "Reward id"
// @Success 200 {object} models.Transaction "TransactionByReward"
// @Router /reward/:id/transaction [get]
func GetTransactionByReward(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetTransactionByReward(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetCurrencyExchangeRateByReward godoc
// @Summary Get GetCurrencyExchangeRateByReward by Reward ID
// @Tags Reward
// @Accept json
// @Produce json
// @Param id path string true "Reward id"
// @Success 200 {object} models.CurrencyExchangeRate "CurrencyExchangeRateByReward"
// @Router /reward/:id/currency-exchange-rate [get]
func GetCurrencyExchangeRateByReward(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetCurrencyExchangeRateByReward(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// StopReward godoc
// @Summary Stop a transaction
// @Description Stop a transaction
// @Tags Reward
// @Accept json
// @Produce json
// @Success 204
// @Router /reward/:id/stop [patch]
func StopReward(c echo.Context) error {
	uuidParam := utils.ParseIDToUUID(c.Param("id"))

	reward, err := service.GetReward(uuidParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	if reward.State != "LIVE" {
		return c.JSON(http.StatusConflict, "Reward must be LIVE to be set to STOPPED")
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err = service.EditStateReward(reward.Uuid, "STOPPED", uuidUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	rewardVariables := map[string]interface{}{
		"rewardUuid": uuidParam,
	}
	camundaService.EndInvestmentProcess(rewardVariables)

	return c.JSON(http.StatusNoContent, nil)
}

// FinishReward godoc
// @Summary Finish a transaction
// @Description Finish a transaction
// @Tags Reward
// @Accept json
// @Produce json
// @Success 204
// @Router /reward/:id/finish [patch]
func FinishReward(c echo.Context) error {
	uuidParam := utils.ParseIDToUUID(c.Param("id"))

	reward, err := service.GetReward(uuidParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	if reward.State != "LIVE" && reward.State != "EXPIRED" && reward.State != "STOPPED" {
		return c.JSON(http.StatusConflict, "Reward must be LIVE or EXPIRED to be set to FINISHED")
	}

	expiredConfig := service.GetLoadedConfig("expired")
	if expiredConfig.Id == 0 {
		return utils.CustomErrorStruct{ErrorMessage: "Configuration expired not found or not accessible"}
	}
	transactionExpiredFeature, err := strconv.ParseBool(expiredConfig.Value)

	if err != nil {
		return utils.CustomErrorStruct{}.InternalServerError("Error parsing env variable TRANSACTION_EXPIRED")
	}

	if transactionExpiredFeature {
		reward.State = "EXPIRED"
	} else {
		reward.State = "FINISHED"
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err = service.EditStateReward(reward.Uuid, reward.State, uuidUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

// VerifyReward godoc
// @Summary Verify reward
// @Description Verify reward
// @Tags Rewards
// @Param id path string true "Reward ID"
// @Success 200
// @Failure 400 {object} utils.CustomErrorStruct
// @Failure 404 {object} utils.CustomErrorStruct
// @Failure 409 {object} utils.CustomErrorStruct
// @Failure 500 {object} utils.CustomErrorStruct
// @Router /reward/:id/verify [post]
func VerifyReward(c echo.Context) error {
	uuidParam := utils.ParseIDToUUID(c.Param("id"))
	reward, err := service.GetReward(uuidParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	if reward.State == "STOPPED" || reward.State == "EXPIRED" || reward.State == "FINISHED" {
		err = service.VerifyReward(uuidParam)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

	} else if reward.State == "LIVE" {
		return c.JSON(http.StatusConflict, utils.CustomErrorStruct{}.BadRequestError(": Reward is Live, need to stop it first"))
	} else {
		return c.JSON(http.StatusConflict, utils.CustomErrorStruct{}.BadRequestError(": Reward is already finished"))
	}

	return c.JSON(http.StatusOK, nil)
}

// RewardBulkEdit godoc
// @Summary Bulk Edit reward
// @Description Bulk Edit reward
// @Tags Reward
// @Accept  json
// @Param Reward body dto.RewardBulkEditReq true "Update Reward dto"
// @Success 200
// @Router /reward/bulk/edit [patch]
func RewardBulkEdit(c echo.Context) error {

	var model dto.RewardBulkEditReq
	keycloak := auth.KeycloakInstance

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := service.BulkEditReward(model, c.Get("user").(*models.User), keycloak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

// CreateRewardBulk godoc
// @Summary Create reward bulk
// @Description Create reward bulk
// @Tags Reward
// @Accept  json
// @Param Reward body dto.RewardBulkEditReq true "Update Reward dto"
// @Success 200
// @Router /reward/bulk/create [post]
func CreateRewardBulk(c echo.Context) error {

	var model dto.CreateRewardBulkDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	model.Origin = "PURCHASE"
	err := service.BulkCreateReward(model, auth.KeycloakInstance, c.Get("user").(*models.User))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

// GetSumRewardsLiveByUserUuid godoc
// @Summary Get Sum Rewards Live By User Uuid
// @Tags Reward
// @Accept json
// @Produce json
// @Param userUUID path string true "User uuid"
// @Success 200 {float64} float64
// @Router /reward/:userUUID/sum-live [get]
func GetSumRewardsLiveByUserUuid(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("userUUID"))

	res, err := service.GetSumRewardsLiveByUserUuid(uuid.String())

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)
}

func RunRewardHistory(c echo.Context) error {
	_, found := lo.Find(c.Get("user").(*models.User).Groups, func(item string) bool { return item == "/user_type/admin" })

	if !found {
		return c.JSON(http.StatusUnauthorized, utils.CustomErrorStruct{ErrorMessage: "User must be admin"})
	}
	keycloak := auth.KeycloakInstance
	service.UpdateRewardsHistory(keycloak)

	return c.JSON(http.StatusOK, nil)
}
