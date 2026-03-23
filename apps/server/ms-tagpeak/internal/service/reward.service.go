package service

import (
	"errors"
	"fmt"
	interactive_brokers "ms-tagpeak/external/interactive-brokers"
	"ms-tagpeak/internal/constants"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/utils"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func GetReward(uuid uuid.UUID) (*models.Reward, error) {
	res, err := repository.GetReward(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.CustomErrorStruct{}.NotFoundError("Reward", uuid)
		}
		return nil, err
	}
	return &res, nil
}

func GetAllRewards() (*[]models.Reward, error) {
	res, err := repository.GetAllRewards()
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func CreateReward(dtoParam dto.CreateRewardDTO, transaction *models.Transaction, uuidUser string, currencyUser string, keycloak *constants.Keycloak) (*models.Reward, error) {
	currencyTargetConfig := GetLoadedConfig("tagpeak_default_currency")
	if currencyTargetConfig.Id == 0 {
		return nil, utils.CustomErrorStruct{ErrorMessage: "Configuration tagpeak_default_currency not found or not accessible"}
	}
	currencyTarget := currencyTargetConfig.Value

	model := utils.RewardDtoToModel(&dtoParam)
	model.User = uuidUser

	// Need to use transaction currency source.
	model.CurrencySource = transaction.CurrencySource

	if transaction.State != "VALIDATED" {
		return nil, utils.CustomErrorStruct{ErrorMessage: "Transaction is not Validated"}
	}

	currencyRates, errCurrency := GetLatestCurrencyExchangeRateFromDatabase()
	if errCurrency != nil {
		return nil, errCurrency
	}
	model.CreatedBy = uuidUser
	model.CurrencyExchangeRateUUID = *currencyRates.Uuid
	model.CurrencyTarget = currencyTarget
	model.CurrencyUser = currencyUser
	model.AssetUnits = RewardUnitInvested(transaction.Cashback, dtoParam.InitialPrice)

	currentRewardSource, currentRewardUser, currentRewardTarget, _ := CalculateCurrentRewards(transaction, keycloak, model.AssetUnits, dtoParam.InitialPrice, nil)
	model.CurrentRewardSource = currentRewardSource
	model.CurrentRewardTarget = currentRewardTarget
	model.CurrentRewardUser = currentRewardUser

	model.WithdrawalUuid = nil
	model.OverridePrice = nil

	res, err := repository.CreateReward(model)
	if err != nil {
		return nil, err
	}

	// Call func to start Investment Manual Task
	err = CamundaManualTask(transaction.Uuid, func() {
		err = EditStateReward(res.Uuid, "LIVE", uuidUser)
		if err != nil {
			fmt.Printf("Error Editing Reward with camunda: %v\n", err)
		}
	})
	if err != nil {
		fmt.Printf("Error closing manual task camunda for start investment: %v\n", err)
	}

	return &res, nil
}

func UpdateReward(dtoParam dto.UpdateRewardDTO, uuid uuid.UUID, uuidUser string, keycloak *constants.Keycloak) (*models.Reward, error) {
	toUpdate, err := repository.GetReward(uuid)
	if err != nil {
		return nil, err
	}

	if dtoParam.Isin != nil {
		toUpdate.Isin = *dtoParam.Isin
	}
	if dtoParam.Conid != nil {
		toUpdate.Conid = *dtoParam.Conid
	}
	if dtoParam.InitialReward != nil {
		toUpdate.InitialReward = *dtoParam.InitialReward
	}
	if dtoParam.CurrentRewardSource != nil {
		toUpdate.CurrentRewardSource = *dtoParam.CurrentRewardSource
	}
	if dtoParam.State != nil {
		toUpdate.State = *dtoParam.State
	}
	if dtoParam.InitialPrice != nil {
		toUpdate.InitialPrice = *dtoParam.InitialPrice
	}
	if dtoParam.EndDate != nil {
		toUpdate.EndDate = *dtoParam.EndDate
	}
	if dtoParam.Type != nil {
		toUpdate.Type = *dtoParam.Type
	}

	if dtoParam.Title != nil {
		toUpdate.Title = *dtoParam.Title
	}

	if dtoParam.Details != nil {
		toUpdate.Details = *dtoParam.Details
	}

	if dtoParam.OverridePrice != nil {
		toUpdate.OverridePrice = dtoParam.OverridePrice

		transaction, _ := GetTransaction(toUpdate.TransactionUUID)

		currentRewardSource, currentRewardUser, currentRewardTarget, _ := CalculateCurrentRewards(transaction, keycloak, toUpdate.AssetUnits, *dtoParam.OverridePrice, nil)
		toUpdate.CurrentRewardSource = currentRewardSource
		toUpdate.CurrentRewardUser = currentRewardUser
		toUpdate.CurrentRewardTarget = currentRewardTarget
	}

	toUpdate.UpdatedBy = &uuidUser

	res, err := repository.UpdateReward(toUpdate)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func DeleteReward(uuid uuid.UUID, uuidUser string) error {
	err := repository.DeleteReward(uuid, uuidUser)
	if err != nil {
		return err
	}
	return nil
}

func GetTransactionByReward(uuid uuid.UUID) (*models.Transaction, error) {
	res, err := repository.GetTransactionByReward(uuid)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetCurrencyExchangeRateByReward(uuid uuid.UUID) (*models.CurrencyExchangeRate, error) {
	res, err := repository.GetCurrencyExchangeRateByReward(uuid)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func EditStateReward(uuid uuid.UUID, state string, uuidUser string) error {
	toUpdate, err := repository.GetReward(uuid)
	if err != nil {
		return err
	}

	toUpdate.State = state
	if state == models.RewardStateStopped {
		now := time.Now()
		toUpdate.StoppedAt = &now
	}

	uuidString := uuidUser
	toUpdate.UpdatedBy = &uuidString

	_, err = repository.UpdateReward(toUpdate)
	if err != nil {
		return err
	}

	return nil
}

func GetExpiredRewards() (*[]string, error) {
	rewards, err := repository.GetExpiredRewards()
	if err != nil {
		return nil, err
	}
	return rewards, nil
}

func UpdateRewardsState(isExpiredOn string, rewards []string) error {
	dataToUpdate := map[string]interface{}{
		"updated_by": "camunda",
	}
	if isExpiredOn == "true" {
		dataToUpdate["state"] = "EXPIRED"
	} else {
		dataToUpdate["state"] = "FINISHED"
	}

	err := repository.BulkUpdateRewards(rewards, dataToUpdate)
	if err != nil {
		return err
	}
	return nil
	// Update rewards state to "FINISHED" if isExpiredOn is false
}

// IsRewardsWithTransactionSaved - Check if rewards with transaction exists
func IsRewardsWithTransactionSaved(transactionUUID uuid.UUID) bool {
	saved, err := repository.IsRewardsWithTransactionSaved(transactionUUID)
	if err != nil {
		fmt.Println("Error Getting rewards state: ", err)
		return false
	}

	return saved
}

func VerifyReward(uuid uuid.UUID) error {
	reward, err := GetReward(uuid)
	if err != nil {
		return err
	}

	if reward.State == "STOPPED" || reward.State == "EXPIRED" || reward.State == "FINISHED" {
		err = CamundaManualTask(reward.Uuid, nil)
		if err != nil {
			fmt.Println("Error verifying reward: ", err)
		}
		err = CamundaManualTask(reward.TransactionUUID, nil)
		if err != nil {
			fmt.Println("Error verifying reward's transaction: ", err)
		}

	} else if reward.State == "LIVE" {
		return utils.CustomErrorStruct{}.BadRequestError("Reward is Live, need to stop it first")
	} else {
		return utils.CustomErrorStruct{}.BadRequestError("Reward is already finished")
	}

	return nil
}

func BulkEditReward(req dto.RewardBulkEditReq, principal *models.User, keycloak *constants.Keycloak) error {
	var bulk []map[string]interface{}

	var uuidsForBulkEdit []string

	for _, reward := range req.Uuids {

		fieldsToUpdate := make(map[string]interface{})
		fieldsToUpdate["uuid"] = reward
		uuidsForBulkEdit = append(uuidsForBulkEdit, reward)
		if req.EndDate != nil {
			fieldsToUpdate["end_date"] = req.EndDate
		}
		if req.InitialPrice != nil {
			fieldsToUpdate["initial_price"] = req.InitialPrice
		}
		if req.InitialDate != nil {
			fieldsToUpdate["created_at"] = req.InitialDate
		}
		if req.Status != nil {
			fieldsToUpdate["state"] = req.Status
		}
		if req.Isin != nil {
			fieldsToUpdate["isin"] = req.Isin
		}
		if req.Conid != nil {
			fieldsToUpdate["conid"] = req.Conid
		}

		if req.OverridePrice != nil {
			r, err := GetReward(utils.ParseIDToUUID(reward))
			if err != nil {
				return err
			}
			fieldsToUpdate["override_price"] = req.OverridePrice

			transaction, err := GetTransaction(r.TransactionUUID)
			if err != nil {
				return err
			}

			currentRewardSource, currentRewardUser, currentRewardTarget, _ := CalculateCurrentRewards(transaction, keycloak, r.AssetUnits, *req.OverridePrice, nil)
			fieldsToUpdate["current_reward_source"] = currentRewardSource
			fieldsToUpdate["current_reward_user"] = currentRewardUser
			fieldsToUpdate["current_reward_target"] = currentRewardTarget

		}
		fieldsToUpdate["updated_by"] = principal.Uuid.String()

		bulk = append(bulk, fieldsToUpdate)
	}

	err := repository.BulkTransactionRewardUpdate(bulk)
	if err != nil {
		return err
	}

	for _, rewardUuid := range uuidsForBulkEdit {
		err = VerifyReward(utils.ParseIDToUUID(rewardUuid))
		if err != nil {
			return err
		}
	}

	return nil
}

func GetSumReferralRewards(uuidUser string) (float64, error) {
	res, err := repository.GetSumReferralRewards(uuidUser)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func BulkCreateReward(data dto.CreateRewardBulkDTO, keycloak *constants.Keycloak, principal *models.User) error {
	currencyTargetConfig := GetLoadedConfig("tagpeak_default_currency")
	if currencyTargetConfig.Id == 0 {
		return utils.CustomErrorStruct{ErrorMessage: "Configuration tagpeak_default_currency not found or not accessible"}
	}
	currencyTarget := currencyTargetConfig.Value

	model := utils.RewardBulkDtoToModel(&data)

	for i := 0; i < len(model); i++ {

		transaction, err := GetTransaction(model[i].TransactionUUID)
		if err != nil {
			return utils.CustomErrorStruct{ErrorMessage: "Transaction not found"}
		}
		if transaction.State != "VALIDATED" {
			return utils.CustomErrorStruct{ErrorMessage: "Transaction is not Validated"}
		}

		user, err := GetUserById(transaction.User, keycloak)
		if err != nil {
			return err
		}
		// Set the owner of the reward
		model[i].User = user.Uuid.String()
		currencyRates, errCurrency := GetLatestCurrencyExchangeRateFromDatabase()
		if errCurrency != nil {
			return errCurrency
		}

		model[i].CreatedBy = principal.Uuid.String()
		model[i].CurrencyExchangeRateUUID = *currencyRates.Uuid
		model[i].CurrencySource = transaction.CurrencySource
		model[i].CurrencyTarget = currencyTarget
		model[i].CurrencyUser = user.Currency
		model[i].AssetUnits = RewardUnitInvested(transaction.Cashback, model[i].InitialPrice)

		currentRewardSource, currentRewardUser, currentRewardTarget, _ := CalculateCurrentRewards(transaction, keycloak, model[i].AssetUnits, model[i].InitialPrice, nil)
		model[i].CurrentRewardSource = currentRewardSource
		model[i].CurrentRewardUser = currentRewardUser
		model[i].CurrentRewardTarget = currentRewardTarget

		model[i].Type = "INVESTMENT"
	}

	err := repository.CreateBulkRewards(model)
	if err != nil {
		return err
	}

	logster.Info("Executing camunda manual tasks")
	for i := 0; i < len(model); i++ {
		err = CamundaManualTask(model[i].TransactionUUID, func() {
			err = EditStateReward(model[i].Uuid, "LIVE", principal.Uuid.String())
			if err != nil {
				fmt.Printf("Error Editing Reward with camunda: %v\n", err)
			}
		})
		if err != nil {
			fmt.Printf("Error closing manual task camunda for start investment: %v\n", err)
		}
	}

	return nil
}

func CreateRewardFromReferral(dtoReward dto.CreateRewardDTO, user models.User) (*models.Reward, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("%+v", dtoReward))
	currencyTargetConfig := GetLoadedConfig("tagpeak_default_currency")
	if currencyTargetConfig.Id == 0 {
		return nil, utils.CustomErrorStruct{ErrorMessage: "Configuration tagpeak_default_currency not found or not accessible"}
	}
	currencyTarget := currencyTargetConfig.Value

	model := utils.RewardDtoToModel(&dtoReward)
	model.User = user.Uuid.String()

	currencyRates, errCurrency := GetLatestCurrencyExchangeRateFromDatabase()
	if errCurrency != nil {
		return nil, errCurrency
	}
	model.CreatedBy = "camunda"
	model.CurrencyExchangeRateUUID = *currencyRates.Uuid
	model.CurrencyTarget = currencyTarget
	model.CurrencyUser = user.Currency

	model.CurrentRewardTarget = utils.GetRewardByCurrencyRate(currencyRates.Rates[dtoReward.CurrencySource], *dtoReward.CurrentRewardUser, currencyRates.Rates[currencyTarget])
	model.CurrentRewardUser = utils.GetRewardByCurrencyRate(currencyRates.Rates[dtoReward.CurrencySource], *dtoReward.CurrentRewardUser, currencyRates.Rates[user.Currency])

	res, err := repository.CreateReward(model)
	if err != nil {
		return nil, err
	}

	logster.EndFuncLog()
	return &res, nil
}

func GetSumRewardsLiveByUserUuid(uuid string) (float64, error) {
	res, err := repository.GetUserLiveRewardSum(uuid)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func GetFinishedRewardsAndSetAsRequested(userUid string, withdrawalUuid string) error {
	rewards, err := repository.GetRewardsByStateAndUser("FINISHED", userUid)
	if err != nil {
		return err
	}

	errSet := repository.UpdateRewardsState("REQUESTED", rewards, withdrawalUuid)

	if errSet != nil {
		return errSet
	}

	return nil
}

func GetRequestedRewardsAndSetAsPaid(userUid string, withdrawalUuid string) error {
	rewards, err := repository.GetRewardsByStateAndUser("REQUESTED", userUid)
	if err != nil {
		logster.Error(err, "Error getting requested rewards")
		return err
	}

	errSet := repository.UpdateRewardsState("PAID", rewards, withdrawalUuid)

	if errSet != nil {
		logster.Error(errSet, "Error updating rewards state")
		return errSet
	}

	return nil
}

func GetRequestedRewardsAndSetAsFinished(userUid string, withdrawalUuid string) error {
	rewards, err := repository.GetRewardsByStateAndUser("REQUESTED", userUid)
	if err != nil {
		logster.Error(err, "Error getting requested rewards")
		return err
	}

	errSet := repository.UpdateRewardsState("FINISHED", rewards, withdrawalUuid)

	if errSet != nil {
		logster.Error(errSet, "Error updating rewards state")
		return errSet
	}

	return nil
}

func CalculateCurrentRewards(transaction *models.Transaction, keycloak *constants.Keycloak, units float64, dayPrice float64, user *string) (float64, float64, float64, *models.User) {
	logster.StartFuncLog()

	userId := transaction.User
	if user != nil {
		userId = *user
	}
	userTransaction, _ := GetUserById(userId, keycloak)
	userMemberShipStatus := GetMembershipStatus(utils.StringPointer(userTransaction.Uuid.String()), userTransaction.Groups)
	percentOnTransaction := *userMemberShipStatus.PercentageOnTransaction

	if userTransaction.TransactionPercentage != nil {
		percentOnTransaction = *userTransaction.TransactionPercentage
	}

	logster.Info("CalculateCurrentRewards - Source")
	currentRewardSource := CalculateCurrentCashRewardValue(
		transaction.AmountSource,
		units,
		dayPrice,
		percentOnTransaction,
	)

	logster.Info("CalculateCurrentRewards - User")
	currentRewardUser := CalculateCurrentCashRewardValue(
		transaction.AmountUser,
		units,
		dayPrice,
		percentOnTransaction,
	)

	logster.Info("CalculateCurrentRewards - Target")
	currentRewardTarget := CalculateCurrentCashRewardValue(
		transaction.AmountTarget,
		units,
		dayPrice,
		percentOnTransaction,
	)

	logster.EndFuncLog()
	return currentRewardSource, currentRewardUser, currentRewardTarget, userTransaction
}

func GetUsersWithTotalAmountReward() ([]response_object.UsersWithTotalAmountReward, error) {
	logster.StartFuncLog()
	list, err := repository.GetUsersWithTotalAmountReward()
	if err != nil {
		return nil, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("GetUsersWithTotalAmountReward - %v", len(list)))
	return list, nil
}

func GetRewardsByState(state string) ([]models.Reward, error) {
	list, err := repository.GetRewardsByState(state)
	if err != nil {
		return nil, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Length: %v", len(list)))
	return list, nil
}

func SaveCurrentReward(uuid uuid.UUID, currentRewardSource float64, currentRewardUser float64, currentRewardTarget float64) error {
	err := repository.UpdateCurrentRewardValue(uuid, currentRewardSource, currentRewardUser, currentRewardTarget)
	if err != nil {
		return err
	}

	return nil
}

func UpdateRewardsHistory(keycloak *constants.Keycloak) {
	logster.StartFuncLog()
	currentCurrencyRate, errCurrentCurrencyRate := GetLatestCurrencyExchangeRateFromDatabase()

	if errCurrentCurrencyRate != nil {
		logster.Error(errCurrentCurrencyRate, "Error while getting latest currency exchange rate")
		logster.EndFuncLog()
		return
	}

	// Get all live rewards
	rewards, err := GetRewardsByState("LIVE")
	if err != nil {
		logster.Error(err, "Error getting rewards")
		return
	}
	if len(rewards) == 0 {
		logster.Info("No rewards to update")
		return
	}
	rewards = lo.Filter(rewards, func(item models.Reward, index int) bool {
		return item.Conid != ""
	})

	// Get lastPrice for all with conid
	conids := lo.UniqMap(rewards, func(r models.Reward, _ int) string {
		return r.Conid
	})
	// iterate through conids so that stringgs like 722635126@SWB have the @
	conids = lo.Map(conids, func(conid string, _ int) string {
		if strings.Contains(conid, "@") {
			return strings.Split(conid, "@")[0]
		}
		return conid
	})

	logster.Info(fmt.Sprintf("Conids: %#v", conids))

	lastPrices, err := interactive_brokers.GetLastPriceBulk(interactive_brokers.GetLastPriceBulkDto{Conids: conids})
	if err != nil {
		logster.Error(err, "Error getting lastPrice from ms-interactive-brokers")
		lastPrices = nil
	}

	if lastPrices == nil {
		mappedRewardsUuids := lo.Map(rewards, func(reward models.Reward, _ int) string { return reward.Uuid.String() })
		errUpdating := RunCopyLatestRewards(mappedRewardsUuids)
		if errUpdating != nil {
			logster.Error(errUpdating, "Error updating failed rewards")
		}
		logster.EndFuncLog()
	} else {

		failedUpdates := make([]models.Reward, 0)

		// For each reward, update reward_history accordingly and the reward current_value
		for _, reward := range rewards {
			logster.Info(fmt.Sprintf("Updating reward: %s", reward.Uuid))

			rateReward, found := lo.Find(*lastPrices, func(item interactive_brokers.GetLastPriceBulkRO) bool {
				rewardConid := reward.Conid
				if strings.Contains(rewardConid, "@") {
					rewardConid = strings.Split(rewardConid, "@")[0]
				}
				return item.Conid == rewardConid
			})
			if !found {
				failedUpdates = append(failedUpdates, reward)
				logster.Error(nil, "No matching last price for this reward found")
				continue
			}
			rateRewardLastPrice, errParseFloat := strconv.ParseFloat(rateReward.LastPrice, 64)
			if errParseFloat != nil {
				logster.Error(errParseFloat, fmt.Sprintf("Error parsing float: %s", rateReward.LastPrice))
				failedUpdates = append(failedUpdates, reward)
				continue
			}

			// Update current reward value
			currentRewardSource, currentRewardUser, currentRewardTarget, userRewardInfo := CalculateCurrentRewards(&reward.Transaction, keycloak, reward.AssetUnits, rateRewardLastPrice, utils.StringPointer(reward.User))

			newRewardHistory := dto.RewardHistoryDTO{
				RewardUUID: &reward.Uuid,
				Rate:       &rateRewardLastPrice,
				Units:      &reward.AssetUnits,
				CashReward: &currentRewardTarget,
			}
			rewardHistory, errRH := CreateRewardHistory(newRewardHistory, reward.User)
			if errRH != nil {
				logster.Error(errRH, "Error creating reward history")
			}
			logster.Info(fmt.Sprintf("newRewardHistory: %s", rewardHistory.UUID.String()))

			valueTarget := currentRewardTarget
			valueUser := utils.GetRewardByCurrencyRate(currentCurrencyRate.Rates[reward.Transaction.CurrencyTarget], currentRewardUser, currentCurrencyRate.Rates[userRewardInfo.Currency])
			valueSource := utils.GetRewardByCurrencyRate(currentCurrencyRate.Rates[reward.Transaction.CurrencyTarget], currentRewardSource, currentCurrencyRate.Rates[reward.Transaction.CurrencySource])

			errSaveReward := SaveCurrentReward(reward.Uuid, valueSource, valueUser, valueTarget)
			if errSaveReward != nil {
				logster.Error(errSaveReward, "Error saving reward current values")
				continue
			}

			logster.Info(fmt.Sprintf("Reward %s updated", reward.Uuid))
		}

		if len(failedUpdates) > 0 {
			mappedRewardsUuids := lo.Map(failedUpdates, func(reward models.Reward, _ int) string { return reward.Uuid.String() })
			errUpdating := RunCopyLatestRewards(mappedRewardsUuids)
			if errUpdating != nil {
				logster.Error(errUpdating, "Error updating failed updated rewards")
			}
		}

	}
	logster.EndFuncLog()
}

func CreateRewardForInfluencer(userUuid string, transactionUuid string, keycloak *constants.Keycloak) (*models.Reward, error) {
	logster.StartFuncLogMsg(userUuid)
	conv, _ := strconv.ParseFloat(GetLoadedConfig("influencer_default_amount").Value, 64)

	amountToCredit := conv
	userReferral, err := GetUserById(userUuid, keycloak)
	if err != nil {
		logster.Error(err, "Error getting user")
		return nil, err
	}

	if userReferral.InfluencerAmount != 0 {
		amountToCredit = userReferral.InfluencerAmount
	}

	logster.Info(fmt.Sprintf("Amount to credit to user %s: %s", userUuid, amountToCredit))
	//The value to credit is in Tagpeak currency, so we have to convert it to the user currency

	currencyTargetConfig := GetLoadedConfig("tagpeak_default_currency")
	if currencyTargetConfig.Id == 0 {
		return nil, utils.CustomErrorStruct{ErrorMessage: "Configuration tagpeak_default_currency not found or not accessible"}
	}
	currencyTarget := currencyTargetConfig.Value

	currencyRates, errCurrency := GetLatestCurrencyExchangeRateFromDatabase()
	if errCurrency != nil {
		return nil, errCurrency
	}

	model := models.Reward{
		CurrencyExchangeRateUUID: *currencyRates.Uuid,
		User:                     userUuid,
		CurrencySource:           userReferral.Currency,
		CurrencyUser:             userReferral.Currency,
		CurrencyTarget:           currencyTarget,
		CurrentRewardTarget:      amountToCredit,
		Type:                     models.RewardTypeFixed,
		Origin:                   models.RewardOriginCommission,
		Title:                    "Influencer commission",
		TransactionUUID:          uuid.MustParse(transactionUuid),
		State:                    models.RewardStateFinished,
		EndDate:                  time.Now(),
		BaseEntity: models.BaseEntity{
			CreatedBy: "camunda",
		},
	}

	amountConverted := utils.GetRewardByCurrencyRate(currencyRates.Rates[userReferral.Currency], amountToCredit, currencyRates.Rates[currencyTarget])
	model.CurrentRewardSource = amountConverted
	model.CurrentRewardUser = amountConverted

	res, err := repository.CreateReward(model)
	if err != nil {
		return nil, err
	}

	logster.EndFuncLog()
	return &res, nil
}
