package service

import (
	"errors"
	"fmt"
	"math"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/response_object"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func GetTransaction(uuid uuid.UUID) (*models.Transaction, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Transaction: %s", uuid))
	res, err := repository.GetTransaction(uuid)
	if err != nil {
		logster.Error(err, "Error getting transaction")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.CustomErrorStruct{}.NotFoundError("Transaction", uuid)
		}
		logster.EndFuncLog()
		return nil, err
	}
	logster.EndFuncLog()
	return &res, nil
}

func GetCashbackView(
	pag pagination.PaginationParams,
	filters dto.TransactionFiltersDTO,
	keycloak *constants.Keycloak,
) (*pagination.PaginationResult, error) {
	logster.StartFuncLog()

	res, resTotal, err := repository.GetCashbackView(pag, filters)
	if err != nil {
		return nil, err
	}

	for i := range res {
		userUuid := res[i].User
		user, err := GetUserById(userUuid, keycloak)
		if err != nil {
			logster.Error(err, "Error getting user")
			continue
		}

		res[i].Email = user.Email
		res[i].RefId = res[i].StoreVisitReference

		if res[i].StoreLogo != nil && *(res[i].StoreLogo) != "" {
			url := dotenv.GetEnv("MS_IMAGES_SERVER_PUBLIC_URL")
			res[i].StoreLogo = utils.StringPointer(fmt.Sprintf(url+"%s/logo.webp", *(res[i].StoreLogo)))
		}
	}

	if res == nil {
		logster.Info("Empty cashback view")
		res = []response_object.CashbackViewRO{}
	}

	response := lo.Map(res, func(item response_object.CashbackViewRO, index int) response_object.CashbackRO {
		var reward response_object.RewardCashbackRO

		if item.RewardUUID.String() != "00000000-0000-0000-0000-000000000000" {
			reward = response_object.RewardCashbackRO{
				Uuid:                &item.RewardUUID,
				Isin:                &item.ISIN,
				Conid:               &item.Conid,
				CurrentRewardSource: &item.CurrentRewardSource,
				CurrentRewardTarget: &item.CurrentRewardTarget,
				CurrentRewardUser:   &item.CurrentRewardUser,
				Status:              &item.Status,
				PriceDayZero:        &item.InitialPrice,
				Title:               &item.Title,
				StartDate:           &item.StartDate,
				EndDate:             &item.EndDate,
				Origin:              item.Origin,
				StoppedAt:           item.StoppedAt,
			}
		}

		var store *response_object.StoreCashbackRO
		if item.StoreUUID != nil {
			store = &response_object.StoreCashbackRO{
				Uuid:              *item.StoreUUID,
				Name:              *item.StoreName,
				Logo:              item.StoreLogo,
				PercentageCashout: &item.StorePercentageCashout,
				CashbackValue:     &item.StoreCashbackValue,
				CashbackType:      &item.StoreCashbackType,
			}
		}

		return response_object.CashbackRO{
			Uuid:              item.TransactionUUID,
			ExitId:            item.RefId,
			Store:             store,
			UserUuid:          item.User,
			Email:             item.Email,
			Date:              item.Date,
			AmountTarget:      item.AmountTarget,
			AmountSource:      item.AmountSource,
			AmountUser:        item.AmountUser,
			CurrencySource:    item.CurrencySource,
			CurrencyTarget:    item.CurrencyTarget,
			NetworkCommission: item.NetworkCommission,
			Status:            item.Status,
			Cashback:          item.Cashback,
			Reward:            &reward,
		}
	})

	pageableResponse := pagination.PaginationResult{
		Limit:      pag.Limit,
		Page:       pag.Page,
		Sort:       pag.Sort,
		TotalRows:  resTotal,
		TotalPages: int(math.Ceil(float64(resTotal) / float64(pag.Limit))),
		Data:       response,
	}

	logster.EndFuncLog()
	return &pageableResponse, nil
}

func CreateTransaction(dtoParam dto.CreateTransactionDTO, storeVisit *models.StoreVisit, uuidString string) (*response_object.CashbackRO, error) {
	logster.StartFuncLog()
	model := utils.TransactionDtoToModel(&dtoParam)
	model.User = uuidString
	model.CreatedBy = uuidString
	model.State = "TRACKED"
	model.IsProcessed = false

	if storeVisit != nil {
		model.StoreUUID = storeVisit.StoreUUID
	}

	currencyTargetConfig, _ := GetConfiguration("tagpeak_default_currency")
	model.CurrencyTarget = currencyTargetConfig.Value

	currencyRates, err := GetLatestCurrencyExchangeRateFromDatabase()
	if err != nil {
		return nil, err
	}
	model.CurrencyExchangeRateUUID = *currencyRates.Uuid

	res, err := repository.CreateTransaction(model)
	if err != nil {
		return nil, err
	}

	mapDto := MapTransactionToRO(res, nil)
	logster.EndFuncLog()
	return &mapDto, nil
}

func ConvertToTagpeakCurrency(transactionId uuid.UUID, userUUID string, keycloak *constants.Keycloak) {
	logster.StartFuncLog()

	currencyTargetConfig := GetLoadedConfig("tagpeak_default_currency")
	if currencyTargetConfig.Id == 0 {
		logster.Error(nil, "Configuration tagpeak_default_currency not found or not accessible")
		return
	}
	currencyTarget := currencyTargetConfig.Value

	userModel, err := GetUserById(userUUID, keycloak)
	if err != nil {
		logster.Error(err, "Error getting user with id:"+userUUID)
		fmt.Println(err)
		return
	}
	currencyUser := userModel.Currency

	transaction, err := repository.GetTransaction(transactionId)
	if err != nil {
		logster.Error(err, "Error getting transaction")
		return
	}

	currencyRates, err := GetCurrencyExchangeRate(transaction.CurrencyExchangeRateUUID)
	if err != nil {
		logster.Error(err, "Error getting currency rates")
		return
	}
	transaction.CurrencyExchangeRateUUID = *currencyRates.Uuid

	transaction.AmountTarget = utils.GetAmountByCurrencyRate(currencyRates.Rates[transaction.CurrencySource], transaction.AmountSource, currencyRates.Rates[currencyTarget])
	transaction.AmountUser = utils.GetAmountByCurrencyRate(currencyRates.Rates[transaction.CurrencySource], transaction.AmountSource, currencyRates.Rates[currencyUser])

	transaction.CommissionTarget = utils.GetAmountByCurrencyRate(currencyRates.Rates[transaction.CurrencySource], transaction.CommissionSource, currencyRates.Rates[currencyTarget])
	transaction.CommissionUser = utils.GetAmountByCurrencyRate(currencyRates.Rates[transaction.CurrencySource], transaction.CommissionSource, currencyRates.Rates[currencyUser])

	_, err = repository.UpdateTransaction(transaction)
	if err != nil {
		logster.Error(err, "Error updating transaction")
		return
	}

	logster.EndFuncLogMsg(fmt.Sprintf("AmountTarget: %f | AmountUser: %f || CommissionTarget: %f | CommissionUser: %f", transaction.AmountTarget, transaction.AmountUser, transaction.CommissionTarget, transaction.CommissionUser))
}

func UpdateTransaction(dtoParam dto.UpdateTransactionDTO, transactionUuid uuid.UUID, user *models.User, startCamundaProcess bool) (*models.Transaction, error) {
	logster.StartFuncLogMsg(transactionUuid)

	toUpdate, err := repository.GetTransaction(transactionUuid)
	if err != nil {
		logster.Error(err, "Error getting transaction")
		logster.EndFuncLog()
		return nil, err
	}

	if dtoParam.CurrencySource != nil {
		toUpdate.CurrencySource = *dtoParam.CurrencySource
	}

	if dtoParam.CommissionTarget != nil {
		// If CommissionTarget is updated, it means ManualCommission is updated
		// and the CommissionTarget stays the same
		toUpdate.ManualCommission = dtoParam.CommissionTarget

		toUpdate.Cashback = calculateCashback(transactionUuid, toUpdate.AmountTarget, dtoParam.CommissionTarget)
		logster.Info(fmt.Sprintf("Calculated cashback: %f", toUpdate.Cashback))
	}

	if dtoParam.OrderDate != nil {
		toUpdate.OrderDate = *dtoParam.OrderDate
	}

	if dtoParam.CommissionSource != nil {
		toUpdate.CommissionSource = *dtoParam.CommissionSource
	}

	if dtoParam.AmountSource != nil {
		toUpdate.AmountSource = *dtoParam.AmountSource
	}

	if user != nil {
		userUuid := user.Uuid.String()
		toUpdate.UpdatedBy = &userUuid
	} else {
		system := "camunda"
		toUpdate.UpdatedBy = &system
	}

	if dtoParam.State != nil {
		toUpdate.State = *dtoParam.State
		if *dtoParam.State == "VALIDATED" {

			cashbackCommissionSource := &toUpdate.CommissionTarget
			if toUpdate.ManualCommission != nil {
				cashbackCommissionSource = toUpdate.ManualCommission
			}

			toUpdate.Cashback = calculateCashback(transactionUuid, toUpdate.AmountTarget, cashbackCommissionSource)
			logster.Info(fmt.Sprintf("Calculated cashback: %f", toUpdate.Cashback))

			if toUpdate.StoreVisitUUID != nil {

				logster.Info(fmt.Sprintf("Setting store visit %s purchase flag to true", *toUpdate.StoreVisitUUID))
				_, err := UpdateStoreVisit(
					dto.UpdateStoreVisitDTO{Purchase: utils.BoolPointer(true)},
					*toUpdate.StoreVisitUUID,
					*toUpdate.UpdatedBy,
				)
				if err != nil {
					logster.Error(err, "Error updating purchase flag for store visit. Proceeding without it.")
				}
			}
		}
	}

	res, err := repository.UpdateTransaction(toUpdate)
	if err != nil {
		logster.Error(err, "Error updating transaction")
		return nil, err
	}

	if startCamundaProcess {
		logster.Info("Starting camunda process to handle other changes")
		transactionCamunda := []dto.CamundaCreateTransactionDTO{
			{
				SourceId:         res.SourceId,
				AmountSource:     res.AmountSource,
				CurrencySource:   res.CurrencySource,
				CommissionSource: res.CommissionSource,
				OrderDate:        res.OrderDate,
				StoreVisitUUID:   res.StoreVisitUUID,
				UserUUID:         uuid.MustParse(res.User),
				Reference:        *res.StoreVisit.Reference,
				State:            res.State,
			},
		}
		StartCamundaTransactionUpdateProcess(transactionCamunda)
	}

	logster.EndFuncLog()
	return &res, nil
}

func DeleteTransaction(uuid uuid.UUID, uuidUser string) error {
	err := repository.DeleteTransaction(uuid, uuidUser)
	if err != nil {
		return err
	}
	return nil
}

func GetStoreVisitByTransaction(uuid uuid.UUID) (models.StoreVisit, error) {
	res, err := repository.GetStoreVisitByTransaction(uuid)
	if err != nil {
		return models.StoreVisit{}, err
	}
	return res, nil
}

func GetStoreByTransaction(uuid uuid.UUID) (models.Store, error) {
	logster.StartFuncLog()
	res, err := repository.GetStoreByTransaction(uuid)
	if err != nil {
		logster.Error(err, "Error getting store")
		logster.EndFuncLog()
		return models.Store{}, err
	}
	logster.EndFuncLog()
	return res, nil
}

func GetStoreOverrideFeeByTransaction(uuid uuid.UUID) *float64 {
	res, err := repository.GetStoreByTransaction(uuid)
	if err != nil {
		return utils.FloatPointer(float64(0))
	}
	if res.OverrideFee == nil || *res.OverrideFee == 0 {
		return utils.FloatPointer(float64(0))
	}

	return res.OverrideFee
}

func GetCurrencyExchangeRateByTransaction(uuid uuid.UUID) (*models.CurrencyExchangeRate, error) {
	res, err := repository.GetCurrencyExchangeRateByTransaction(uuid)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetRewardByTransactionAndUser(uuid uuid.UUID, userUuid uuid.UUID) (*models.Reward, error) {
	res, err := repository.GetRewardByTransactionAndUser(uuid, userUuid)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func MapTransactionToRO(transaction models.Transaction, state *[]string) response_object.CashbackRO {
	var reward response_object.RewardCashbackRO

	// Fetch reward only if the transaction state is "VALIDATED"
	if strings.TrimSpace(transaction.State) == "VALIDATED" {
		rewardModel, e := repository.GetRewardByTransaction(transaction.Uuid, state)

		if e != nil {
			return utils.MapCashbackDto(&transaction, &(response_object.RewardCashbackRO{}))
		}

		rewardDto := utils.MapRewardRO(rewardModel)

		reward = rewardDto
	}

	return utils.MapCashbackDto(&transaction, &reward)
}

func GetTransactionBySourceId(sourceId string) (*models.Transaction, error) {
	logster.StartFuncLog()
	res, err := repository.GetTransactionBySourceId(sourceId)
	if err != nil {
		logster.Error(err, "Error getting transaction by sourceId")
		return nil, err
	}

	logster.EndFuncLog()
	return res, nil
}

func GetTransactionByStoreVisitUuid(storeVisitUuid string) (*models.Transaction, error) {
	logster.StartFuncLog()
	res, err := repository.GetTransactionByStoreVisitUuid(storeVisitUuid)
	if err != nil {
		logster.Error(err, "Error getting transaction by sourceId")
		return nil, err
	}

	logster.EndFuncLog()
	return res, nil
}

func IsTransactionProcessed(uuid uuid.UUID) (bool, error) {
	res, err := repository.IsTransactionProcessed(uuid)
	if err != nil {
		return false, err
	}
	return res, nil
}

func UpdateTransactionToProcessed(uuid uuid.UUID, processed bool) error {
	err := repository.UpdateTransactionToProcessed(uuid, processed)
	if err != nil {
		return err
	}
	return nil
}

func BulkUpdateTransaction(data dto.UpdateTransactionDTO, principal *models.User) error {
	logster.StartFuncLog()
	dataToUpdate := map[string]interface{}{
		"updated_by": principal.Uuid.String(),
	}

	if data.State != nil {
		dataToUpdate["state"] = *data.State
	}
	if data.CurrencySource != nil {
		dataToUpdate["currency_source"] = *data.CurrencySource
	}
	if data.CommissionTarget != nil {
		dataToUpdate["commission_target"] = *data.CommissionTarget
	}
	if data.OrderDate != nil {
		dataToUpdate["order_date"] = *data.OrderDate
	}
	if data.ExitClick != nil {
		dataToUpdate["exit_click"] = *data.ExitClick
	}

	err := repository.BulkUpdateTransactions(*data.Uuids, dataToUpdate)
	if err != nil {
		logster.Error(err, "Error updating transactions")
		return err
	}

	transactionsUpdated, err := repository.BulkGetTransactionsByUuids(*data.Uuids)
	transactionsCamunda := lo.Map(transactionsUpdated, func(item models.Transaction, index int) dto.CamundaCreateTransactionDTO {
		return dto.CamundaCreateTransactionDTO{
			SourceId:         item.SourceId,
			AmountSource:     item.AmountSource,
			CurrencySource:   item.CurrencySource,
			CommissionSource: item.CommissionSource,
			OrderDate:        item.OrderDate,
			StoreVisitUUID:   item.StoreVisitUUID,
			UserUUID:         uuid.MustParse(item.User),
			Reference:        *item.StoreVisit.Reference,
			State:            item.State,
		}
	})
	StartCamundaTransactionUpdateProcess(transactionsCamunda)

	if err != nil {
		logster.Error(err, "Error updating transactions")
		return err
	}

	logster.EndFuncLog()
	return nil
	// Update rewards state to "FINISHED" if isExpiredOn is false
}

func CalculateAndSetCashback(uuidTransactionParam string) error {
	logster.StartFuncLog()

	uuidTransaction, _ := uuid.Parse(uuidTransactionParam)
	store, err := GetStoreByTransaction(uuidTransaction)
	if err != nil {
		logster.Error(err, "Error while getting store by transaction")
		return utils.CustomErrorStruct{
			ErrorMessage: "Error while getting store by transaction",
		}
	}

	transaction, err := GetTransaction(uuidTransaction)
	if err != nil {
		logster.Error(err, "Error while getting transaction")
		return utils.CustomErrorStruct{
			ErrorMessage: "Error while getting transaction",
		}
	}

	fee := 0.0
	feeConfig := GetLoadedConfig("transaction_fixed_fee")
	if feeConfig.Id == 0 {
		fee = 0.0
	}
	fee, _ = strconv.ParseFloat(feeConfig.Value, 64)

	overrideFee := store.OverrideFee
	if overrideFee != nil && *overrideFee != 0 {
		fee = *overrideFee
	}

	cashbackValue := *store.CashbackValue
	if *store.CashbackType == models.CashbackTypeFixed {
		cashbackValue = transaction.CommissionTarget
		if transaction.ManualCommission != nil {
			cashbackValue = *transaction.ManualCommission
		}
	}

	cashback := CalculateTransactionCashback(cashbackValue, *store.CashbackType, fee, transaction.AmountTarget)

	err = repository.UpdateTransactionCashback(uuidTransaction, cashback)
	if err != nil {
		logster.Error(err, "Error while updating transaction cashback")
		return utils.CustomErrorStruct{
			ErrorMessage: "Error while updating transaction cashback",
		}
	}

	logster.EndFuncLog()
	return nil
}

func GetAmountUserTransactions(uuid uuid.UUID) float64 {
	amountTransaction, err := repository.GetUserAmountTransactions(uuid)
	if err != nil {
		return 0
	}

	return amountTransaction
}

func calculateCashback(transactionUuid uuid.UUID, amount float64, commission *float64) float64 {
	logster.StartFuncLog()
	store, _ := GetStoreByTransaction(transactionUuid)

	if store.Uuid.String() == "00000000-0000-0000-0000-000000000000" {
		logster.EndFuncLogMsg("Store is empty")
		return float64(0)
	}

	fee := 0.0
	feeConfig := GetLoadedConfig("transaction_fixed_fee")
	if feeConfig.Id == 0 {
		fee = 0.0
	}
	fee, _ = strconv.ParseFloat(feeConfig.Value, 64)

	overrideFee := GetStoreOverrideFeeByTransaction(transactionUuid)
	if overrideFee != nil && *overrideFee != 0 {
		fee = *overrideFee
	}

	// If we have a manual commission, we calculate the cashback as a fixed value instead of percentual
	if *store.CashbackType == models.CashbackTypeFixed && commission != nil {
		logster.EndFuncLog()
		return CalculateTransactionCashback(*commission, *store.CashbackType, fee, amount)
	}

	logster.EndFuncLog()
	return CalculateTransactionCashback(*store.CashbackValue, *store.CashbackType, fee, amount)
}

func ManageUserMembershipLevel(userId string, keycloak *constants.Keycloak) error {
	logster.StartFuncLogMsg(fmt.Sprintf("User uuid: %s", userId))

	userUuid, _ := uuid.Parse(userId)

	// Get referral to get referrer uuid
	referral, err := GetReferralByInvitee(userUuid)
	if err != nil {
		logster.Error(err, "GetReferralByInvitee")
		logster.EndFuncLog()
		return err
	}

	logster.Info(fmt.Sprintf("Get referral with referredUuid: %v\n", referral.ReferrerUUID))

	// VALIDATE IF THE USER IS AN INVITEE AND ALREADY MADE A TRANSACTION
	// Avoid nil pointer exception when referral is nil or with all zeros uuid
	if referral != nil && referral.Uuid != utils.ParseIDToUUID("00000000-0000-0000-0000-000000000000") {

		// 1. Check if is first purchase
		if !referral.SuccessfulFirstTransaction {
			referral.SuccessfulFirstTransaction = true
			_, err = repository.UpdateReferralFirstTransaction(*referral)
			if err != nil {
				logster.Error(err, "UpdateReferralFirstTransaction")
				logster.EndFuncLog()
				return err
			}

			logster.Info(fmt.Sprintf("Updated sucessful first transaction for user with userUuid: %s\n", userId))

			// Set referred user as silver since it's first transaction
			_ = UpdateUserGroups(userUuid, *Configuration.MembershipLevels.Silver, keycloak)
			logster.Info(fmt.Sprintf("Set user %s with Membership Level Silver", userId))

			// No need to set referrer to next level if meeting criteria
			// because the rest of the process already takes care of that
		}

		// 2. Get number of successful referrals to update REFERRER membership level

		// Get list of referrals for the referrer to get the total number of referrals
		numberOfReferrals, err := repository.GetNumberOfSuccessfulReferralsByReferrerUuid(*referral.ReferrerUUID)
		if err != nil {
			logster.Error(err, "GetNumberOfSuccessfulReferralsByReferrerUuid")
			logster.EndFuncLog()
			return err
		}

		logster.Info(fmt.Sprintf("Number of referrals of referrer uuid: %v -> %v successfull referrals \n", referral.ReferrerUUID, numberOfReferrals))

		err = UpdateMembershipByReferralsNumber(*referral.ReferrerUUID, keycloak, numberOfReferrals)
		if err != nil {
			logster.Error(err, "UpdateMembershipByReferralsNumber")
			logster.EndFuncLog()
			return err
		}

	}

	// 3. Get user total transactions amount
	totalAmount := GetAmountUserTransactions(userUuid)

	logster.Info(fmt.Sprintf("Total transaction's amount for user with uuid: %s -> %v  \n", userId, totalAmount))

	err = UpdateMembershipByTransactionAmount(userUuid, keycloak, totalAmount)
	if err != nil {
		logster.Error(err, "UpdateMembershipByTransactionAmount")
		logster.EndFuncLog()
		return err
	}

	logster.EndFuncLog()
	return nil
}

func ManageUserMembershipLevelMigration(userId string, keycloak *constants.Keycloak) error {
	fmt.Printf("\n\n---------------------------\n")
	fmt.Printf("Start ManageUserMembershipLevel with userId %s\n", userId)

	// 1 ver se é referral como invitee e se tiver 1 transaction como validated, meter como sucessfull transaction
	// 2 ver quantas successful referrals tem associadas a ele como referrer
	// 3 ver o total de transações do user

	userUuid, _ := uuid.Parse(userId)

	fmt.Printf("1-Check if user has referral as invitee and has, at least, 1 transaction as validated\n")
	// Get referral where he is invitee
	referral, err := GetReferralByInvitee(userUuid)
	if err != nil {
		fmt.Printf("Error -  %v\n", err)
		return err
	}
	fmt.Printf("Get referral: %v\n", referral.Uuid)

	// Avoid nil pointer exception when referral is nil or with all zeros uuid
	if referral != nil && referral.Uuid != utils.ParseIDToUUID("00000000-0000-0000-0000-000000000000") {
		// 1. Check if is first purchase
		// VALIDATE IF THE USER IS AN INVITEE AND ALREADY MADE A TRANSACTION
		if !referral.SuccessfulFirstTransaction {
			hasTransaction, _ := repository.UserHasMoreThan1Transaction(userUuid, "VALIDATED")
			if hasTransaction {
				referral.SuccessfulFirstTransaction = true
				_, err = repository.UpdateReferralFirstTransaction(*referral)
				if err != nil {
					fmt.Printf("Error -  %v\n", err)
					return err
				}
				fmt.Printf("Updated sucessful first transaction for referral: %s\n", referral.Uuid)
			}
		}
	}

	fmt.Printf("2-Check if user has referrals as referrer and how many with successful first transaction\n")
	// 2. Get number of successful referrals as a  REFERRER to update membership level
	// Get list of referrals for the referrer to get the total number of referrals
	numberOfReferrals, err := repository.GetNumberOfSuccessfulReferralsByReferrerUuid(userUuid)
	if err != nil {
		fmt.Printf("Error -  %v\n", err)
		return err
	}
	fmt.Printf("Number of successfull referrals for user: %v -> %v successfull referrals \n", userUuid, numberOfReferrals)

	err = UpdateMembershipByReferralsNumber(userUuid, keycloak, numberOfReferrals)
	if err != nil {
		return err
	}

	// Check if user membership level and cut the method early
	userInfo, err := GetUserById(userUuid.String(), keycloak)
	if err != nil {
		fmt.Printf("Error -  %v\n", err)
		return err
	}
	if lo.Contains(userInfo.Groups, *Configuration.MembershipLevels.Gold) {
		fmt.Printf("User membership level is Gold and can't upgrade any higher. Skipping step 3\n")

		fmt.Printf("End ManageUserMembershipLevel with userId %s\n", userId)
		fmt.Printf("---------------------------\n")
		return nil
	}

	fmt.Printf("3-Check if user amount of transactions to set membership_level\n")
	// 3. Get user total transactions amount
	totalAmount := GetAmountUserTransactions(userUuid)

	fmt.Printf("Total transaction's amount for user -> %v  \n", totalAmount)
	err = UpdateMembershipByTransactionAmount(userUuid, keycloak, totalAmount)
	if err != nil {
		fmt.Printf("Error -  %v\n", err)
		return err
	}

	fmt.Printf("End ManageUserMembershipLevel with userId %s\n", userId)
	fmt.Printf("---------------------------\n")

	return nil
}

// Update user membership level by number of referrals
func UpdateMembershipByReferralsNumber(userUuid uuid.UUID, keycloak *constants.Keycloak, numberReferrals int64) error {
	logster.StartFuncLogMsg(fmt.Sprintf("UserUuid: %s", userUuid))

	// Validate if user is influencer
	user, err := GetUserById(userUuid.String(), keycloak)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if slices.Contains(user.Groups, *Configuration.MembershipLevels.Influencer) {
		logster.Info("User is an Influencer. The membership level can't be updated.")
		return nil
	}

	// Get configs from DB
	numberReferralsSilver, _ := strconv.ParseInt(GetLoadedConfig("referral_silver_status_goal").Value, 10, 64)
	numberReferralsGold, _ := strconv.ParseInt(GetLoadedConfig("referral_gold_status_goal").Value, 10, 64)
	logster.Info(fmt.Sprintf("Treshold Silver %v -  Treshold Gold %v", numberReferralsSilver, numberReferralsGold))

	if numberReferrals >= numberReferralsGold {

		// Update user groups
		_ = UpdateUserGroups(userUuid, *Configuration.MembershipLevels.Gold, keycloak)

		logster.Info(fmt.Sprintf("Update membership of user with userId %v with level: GOLD due to number of referrals", userUuid))

	} else if numberReferrals >= numberReferralsSilver {

		// Update user groups
		_ = UpdateUserGroups(userUuid, *Configuration.MembershipLevels.Silver, keycloak)

		logster.Info(fmt.Sprintf("Update membership of user with userId %v with level: SILVER due to number of referrals", userUuid))

	}

	logster.EndFuncLog()
	return nil
}

// Update user membership level by total transactions amount
func UpdateMembershipByTransactionAmount(userUuid uuid.UUID, keycloak *constants.Keycloak, totalTransactionsAmount float64) error {
	logster.StartFuncLogMsg(fmt.Sprintf("UserUuid: %s", userUuid))

	// Validate if user is influencer
	user, err := GetUserById(userUuid.String(), keycloak)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if slices.Contains(user.Groups, *Configuration.MembershipLevels.Influencer) {
		logster.Info("User is an Influencer. The membership level can't be updated.")
		return nil
	}

	// Get configs from DB
	amountReferralsSilver, _ := strconv.ParseFloat(GetLoadedConfig("referral_silver_status_goal_amount").Value, 64)
	amountReferralsGold, _ := strconv.ParseFloat(GetLoadedConfig("referral_gold_status_goal_amount").Value, 64)
	logster.Info(fmt.Sprintf("Treshold Silver %v | Treshold Gold %v", amountReferralsSilver, amountReferralsGold))

	if totalTransactionsAmount >= amountReferralsGold {

		// Update user groups
		_ = UpdateUserGroups(userUuid, *Configuration.MembershipLevels.Gold, keycloak)

		logster.Info(fmt.Sprintf("Update membership of user with userId %v with level: GOLD due to total amount of transactions", userUuid))

	} else if totalTransactionsAmount >= amountReferralsSilver {

		// Update user groups
		_ = UpdateUserGroups(userUuid, *Configuration.MembershipLevels.Silver, keycloak)

		logster.Info(fmt.Sprintf("Update membership of user with userId %v with level: SILVER due to total amount of transactions", userUuid))

	}

	logster.EndFuncLog()
	return nil
}

func UpdateUserGroups(userUuid uuid.UUID, level string, keycloak *constants.Keycloak) error {
	logster.StartFuncLog()

	updateDto := dto.UpdateUserDto{
		Groups: &[]string{strings.Split(level, "/")[2], strings.Split(*Configuration.UserTypes.Default, "/")[2]},
	}

	_, err := UpdateUser(userUuid, updateDto, keycloak)
	if err != nil {
		logster.Error(err, "UpdateUserGroups")
		logster.EndFuncLog()
		return err
	}

	go UpdateUserTopic(userUuid.String(), level)

	logster.EndFuncLog()
	return nil
}

func ManageAllUsersMembershipLevel(keycloak *constants.Keycloak) {
	logster.StartFuncLog()

	users, err := GetMaxUsers(keycloak, nil)
	if err != nil {
		logster.Error(err, "Error GetMaxUsers")
	}
	logster.Info(fmt.Sprintf("Users length: %v\n", len(users)))
	for idx, user := range users {
		fmt.Printf("User index %v\n", idx)
		err := ManageUserMembershipLevelMigration(user.Uuid.String(), keycloak)
		if err != nil {
			logster.Error(err, "Error ManageUserMembershipLevelMigration")
		}
		getKeycloakAdminToken(keycloak)
	}

	logster.EndFuncLog()
	go SetMigratedUserBalance(keycloak)
}

func getKeycloakAdminToken(keycloak *constants.Keycloak) {
	logster.StartFuncLog()
	now := time.Now()

	if now.After(keycloak.TokenExpireDate) {
		logster.Info("Token expired, getting new token")
		token, err := keycloak.Client.LoginAdmin(
			keycloak.Ctx,
			dotenv.GetEnv("TAGPEAK_ADMIN_USERNAME"),
			dotenv.GetEnv("TAGPEAK_ADMIN_PASSWORD"),
			"master",
		)
		if err != nil {
			logster.Error(err, "Error getting Keycloak token")
		}

		keycloak.AdminToken = token
		keycloak.TokenExpireDate = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}
	logster.EndFuncLog()
}

func UserHaveTransactions(userUUID string) (bool, error) {
	res, err := repository.UserHaveTransactions(userUUID)
	if err != nil {
		return false, err
	}
	return res, nil
}

func StartCamundaTransactionUpdateProcess(transactionsCamunda []dto.CamundaCreateTransactionDTO) {
	trx := map[string]interface{}{
		"loopCounter":  len(transactionsCamunda),
		"transactions": transactionsCamunda,
	}

	resp := camundaPKG.StartProcessInstance(
		camundaPKG.InjectEnvOnKey("sub-process-transaction"),
		*camundaPKG.GetCamundaClient(),
		trx,
	)
	logster.Info(fmt.Sprintf("Process started -Process Instance Key : %d", resp.ProcessInstanceKey))
}

func UpdateShopifyActionableTransactionsState() error {
	logster.StartFuncLog()

	affectedTransactions, err := repository.UpdateShopifyActionableTransactionsState()
	if err != nil {
		logster.Error(err, "Error updating shopify actionable transactions state")
		return err
	}

	logster.Info(fmt.Sprintf("Updated %d shopify actionable transactions to VALIDATED state", len(affectedTransactions)))

	for _, affectedTransaction := range affectedTransactions {
		logster.Info(fmt.Sprintf("UUID: %s", affectedTransaction))

		transactionCamunda := &dto.CamundaCreateTransactionDTO{
			SourceId:         affectedTransaction.TrxSourceID,
			AmountSource:     affectedTransaction.TrxAmountSource,
			CurrencySource:   affectedTransaction.TrxCurrencySource,
			CommissionSource: affectedTransaction.TrxCommissionSource,
			OrderDate:        affectedTransaction.TrxOrderDate,
			StoreVisitUUID:   &affectedTransaction.TrxStoreVisitUUID,
			UserUUID:         uuid.MustParse(affectedTransaction.TrxUser),
			Reference:        "",
			State:            "approved",
		}
		StartCamundaProcessForShopifyOrder(transactionCamunda)
	}

	return nil
}

func UserHasMoreThanOneTransaction(userUuid string) (bool, error) {
	logster.StartFuncLogMsg(userUuid)
	hasMoreThanOne, errCount := repository.UserHasMoreThanOneTransaction(userUuid)
	if errCount != nil {
		logster.Error(errCount, "Error GetUserTransactionCount")
		logster.EndFuncLog()
		return true, errCount
	}

	logster.EndFuncLogMsg(fmt.Sprintf("User has more than one transaction: %v", hasMoreThanOne))
	return hasMoreThanOne, nil
}
