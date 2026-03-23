package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ms-tagpeak/internal/constants"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
)

func GetWithdrawal(uuid uuid.UUID) (*models.Withdrawal, error) {
	logster.StartFuncLog()

	res, err := repository.GetWithdrawal(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.CustomErrorStruct{}.NotFoundError("Store", uuid)
		}
		return nil, err
	}

	logster.EndFuncLog()
	return &res, nil
}

func GetAllWithdrawals(
	pag pagination.PaginationParams,
	filters *dto.WithdrawalFiltersDTO) (*pagination.PaginationResult, error) {
	logster.StartFuncLog()

	res, err := repository.GetAllWithdrawalsWithPagination(pag, *filters)
	if err != nil {
		return nil, err
	}

	logster.EndFuncLog()
	return res, nil
}

func CreateWithdrawal(
	uuidUser *string,
	currencyUser string,
	amount float64,
	userMethods *[]models.UserPaymentMethod) (*response_object.WithdrawalRO, error) {
	logster.StartFuncLog()

	lastRate, err := GetLatestCurrencyExchangeRateFromDatabase()
	if err != nil {
		return nil, err
	}

	currencyTargetConfig := GetLoadedConfig("tagpeak_default_currency")
	if currencyTargetConfig.Id == 0 {
		return nil, utils.CustomErrorStruct{ErrorMessage: "Configuration tagpeak_default_currency not found or not accessible"}
	}
	targetCurrency := currencyTargetConfig.Value

	var methodUuid uuid.UUID

	// GET THE USER METHOD -> BANK
	for _, userMethod := range *userMethods {
		if *userMethod.PaymentMethod.Code == "bank" {
			methodUuid = userMethod.Uuid
		}
	}

	model := models.Withdrawal{
		UserMethod:               methodUuid,
		User:                     *uuidUser,
		AmountSource:             amount,
		AmountTarget:             utils.GetAmountByCurrencyRate(lastRate.Rates[currencyUser], amount, lastRate.Rates[targetCurrency]),
		State:                    "PENDING",
		CompletionDate:           nil,
		CurrencyExchangeRateUUID: *lastRate.Uuid,
		CurrencySource:           currencyUser,
		CurrencyTarget:           targetCurrency,
		Details:                  nil,
	}
	model.CreatedBy = *uuidUser
	model.UpdatedBy = uuidUser

	res, err := repository.CreateWithdrawal(model)
	if err != nil {
		return nil, err
	}

	response := response_object.WithdrawalRO{
		Uuid:         res.Uuid.String(),
		AmountSource: res.AmountSource,
		AmountTarget: res.AmountTarget,
		Details:      res.Details,
		State:        res.State,
		BaseEntityRO: response_object.BaseEntityRO{
			CreatedAt: res.CreatedAt,
			CreatedBy: res.CreatedBy,
			UpdatedAt: res.UpdatedAt,
			UpdatedBy: res.UpdatedBy,
			Deleted:   res.Deleted,
			DeletedAt: res.DeletedAt,
			DeletedBy: res.DeletedBy,
		},
	}

	if res.CompletionDate != nil {
		response.CompletionDate = res.CompletionDate
	}

	errorReward := GetFinishedRewardsAndSetAsRequested(*uuidUser, res.Uuid.String())

	if errorReward != nil {
		return nil, errorReward
	}

	logster.EndFuncLog()
	return &response, nil
}

func UpdateWithdrawal(dtoParam dto.UpdateWithdrawalDTO, uuid uuid.UUID, uuidUser string, keycloak *constants.Keycloak) (*models.Withdrawal, error) {
	logster.StartFuncLog()

	toUpdate, err := repository.GetWithdrawal(uuid)
	if err != nil {
		logster.Error(err, "Error getting withdrawal")
		return nil, err
	}

	if dtoParam.Details != nil {
		toUpdate.Details = dtoParam.Details
	}
	if dtoParam.State != nil && dtoParam.CompletionDate != nil {
		updatedModel, err := updateStateWithdrawal(&toUpdate, dtoParam, keycloak)
		if err != nil {
			logster.Error(err, "Error updating withdrawal state")
			return nil, err
		}
		toUpdate = *updatedModel
	}

	if dtoParam.CurrencySource != nil {
		toUpdate.CurrencySource = *dtoParam.CurrencySource
	}
	if dtoParam.CurrencyTarget != nil {
		toUpdate.CurrencyTarget = *dtoParam.CurrencyTarget
	}
	if dtoParam.UserMethodID != nil {
		toUpdate.UserMethod = *dtoParam.UserMethodID
	}
	if dtoParam.UserID != nil {
		toUpdate.User = *dtoParam.UserID
	}

	toUpdate.UpdatedBy = &uuidUser

	res, err := repository.UpdateWithdrawal(toUpdate)
	if err != nil {
		logster.Error(err, "Error updating withdrawal")
		return nil, err
	}

	if dtoParam.State != nil {
		if *dtoParam.State == "COMPLETED" {
			//set reward as paid
			errorReward := GetRequestedRewardsAndSetAsPaid(res.User, res.Uuid.String())

			if errorReward != nil {
				logster.Error(errorReward, "Error getting requested rewards and set as paid")
				return nil, errorReward
			}

			if toUpdate.UserPaymentMethod.State == "PENDING" {
				errorReward = UpdateUserPaymentMethodState(&res.UserMethod, "VALIDATED")
				if errorReward != nil {
					logster.Error(errorReward, "Error updating user payment method state")
					return nil, errorReward
				}
			}
		}

		if *dtoParam.State == "REJECTED" {
			//set reward as finished
			errorReward := GetRequestedRewardsAndSetAsFinished(res.User, res.Uuid.String())

			if errorReward != nil {
				logster.Error(errorReward, "Error getting requested reward and set as finished")
				return nil, errorReward
			}

			if toUpdate.UserPaymentMethod.State == "PENDING" {
				errorReward = UpdateUserPaymentMethodState(&res.UserMethod, "REJECTED")
				if errorReward != nil {
					logster.Error(errorReward, "Error updating user payment method state")
					return nil, errorReward
				}
			}
		}
	}

	logster.EndFuncLog()
	return &res, nil
}

func DeleteWithdrawal(uuid uuid.UUID, uuidUser string) error {
	logster.StartFuncLog()

	err := repository.DeleteWithdrawal(uuid, uuidUser)
	if err != nil {
		logster.Error(err, "Error deleting withdrawal")
		return err
	}

	logster.EndFuncLog()
	return nil
}

func GetPaidWithdrawalsAmount(uuidUser string) (float64, error) {
	logster.StartFuncLog()

	res, err := repository.GetPaidWithdrawalsAmount(uuidUser)
	if err != nil {
		logster.Error(err, "Error getting paid withdrawals amount")
		return 0, err
	}

	logster.EndFuncLog()
	return res, nil
}

func GetLatestPendingWithdrawal(uuidUser string) (bool, error) {
	logster.StartFuncLog()

	logster.EndFuncLog()
	return repository.GetLatestPendingWithdrawal(uuidUser)
}

func updateStateWithdrawal(model *models.Withdrawal, dtoParams dto.UpdateWithdrawalDTO, keycloak *constants.Keycloak) (*models.Withdrawal, error) {
	logster.StartFuncLog()

	if model.State != "PENDING" {
		return nil, fmt.Errorf("withdrawal can only be updated when in PENDING state, current state: %s", model.State)
	}

	model.State = *dtoParams.State
	model.CompletionDate = dtoParams.CompletionDate

	if *dtoParams.State == "COMPLETED" {
		user, err := GetUserById(model.User, keycloak)
		if err != nil {
			logster.Error(err, "Error getting user")
			return nil, err
		}

		user.Balance = 0

		balanceString := fmt.Sprintf("%.2f", user.Balance)
		user, err = UpdateUser(user.Uuid, dto.UpdateUserDto{Balance: &balanceString}, keycloak)
	}

	logster.EndFuncLog()
	return model, nil
}

func BulkUpdateStateWithdrawal(dtoParam dto.BulkUpdateStateRequestDTO, uuidUser string, keycloak *constants.Keycloak) error {
	logster.StartFuncLog()

	for _, withdrawal := range dtoParam.Uuids {
		model, err := repository.GetWithdrawal(utils.ParseIDToUUID(withdrawal))
		if err != nil {
			logster.Error(err, "Error getting withdrawal")
			return err
		}

		modelUpdated, err := updateStateWithdrawal(&model, dto.UpdateWithdrawalDTO{
			State:          dtoParam.State,
			CompletionDate: dtoParam.CompletionDate,
		}, keycloak)
		if err != nil {
			logster.Error(err, "Error updating withdrawal state")
			return err
		}

		if *dtoParam.State == "COMPLETED" {
			err = UpdateUserPaymentMethodState(&model.UserMethod, "VALIDATED")
			if err != nil {
				logster.Error(err, "Error updating user payment method state")
				return err
			}
		} else if *dtoParam.State == "REJECTED" {
			err = UpdateUserPaymentMethodState(&model.UserMethod, "REJECTED")
			if err != nil {
				logster.Error(err, "Error updating user payment method state")
				return err
			}
		}

		model = *modelUpdated
		model.UpdatedBy = &uuidUser

		_, err = repository.UpdateWithdrawal(model)
		if err != nil {
			logster.Error(err, "Error updating withdrawal")
			return err
		}
	}

	logster.EndFuncLog()
	return nil
}
