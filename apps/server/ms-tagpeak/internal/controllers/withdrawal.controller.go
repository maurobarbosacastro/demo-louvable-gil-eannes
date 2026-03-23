package controllers

import (
	"ms-tagpeak/internal/auth"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	_ "ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

// GetWithdrawal godoc
// @Summary Get Withdrawal by ID
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Param id path string true "Withdrawal id"
// @Success 200 {object} response_object.WithdrawalRO
// @Router /withdrawal/:id [get]
func GetWithdrawal(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	keycloak := auth.KeycloakInstance

	res, err := service.GetWithdrawal(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user, err := service.GetUserById(res.User, keycloak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	userPaymentMethod, err := service.GetUserPaymentMethodByIdUnscoped(res.UserMethod)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	modelDto := response_object.WithdrawalRO{
		Uuid:           res.Uuid.String(),
		AmountTarget:   res.AmountTarget,
		CurrencyTarget: &res.CurrencyTarget,
		State:          res.State,
		CreatedAt:      &res.CreatedAt,
		CompletionDate: res.CompletionDate,
		User: response_object.SimpleUserDto{
			Uuid: user.Uuid,
			Name: utils.GetUserName(user),
		},
		PaymentMethod: userPaymentMethod,
	}

	return c.JSON(http.StatusOK, modelDto)

}

// GetAllWithdrawals godoc
// @Summary Get all Withdrawals
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Success 200 {array} response_object.WithdrawalRO "Array of Withdrawals"
// @Router /withdrawal [get]
func GetAllWithdrawals(c echo.Context) error {
	var pag pagination.PaginationParams
	var filters dto.WithdrawalFiltersDTO
	keycloak := auth.KeycloakInstance

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}

	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		return errFilters
	}

	res, err := service.GetAllWithdrawals(pag, &filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	modelDto := lo.Map(res.Data.([]models.Withdrawal), func(item models.Withdrawal, _ int) response_object.WithdrawalRO {
		user, err := service.GetUserById(item.User, keycloak)
		if err != nil {
			return response_object.WithdrawalRO{}
		}

		userPaymentMethod, err := service.GetUserPaymentMethodByIdUnscoped(item.UserMethod)
		if err != nil {
			return response_object.WithdrawalRO{}
		}

		return response_object.WithdrawalRO{
			Uuid:           item.Uuid.String(),
			CompletionDate: item.CompletionDate,
			State:          item.State,
			CreatedAt:      &item.CreatedAt,
			Details:        item.Details,
			AmountTarget:   item.AmountTarget,
			User: response_object.SimpleUserDto{
				Uuid:  user.Uuid,
				Name:  utils.GetUserName(user),
				Email: user.Email,
			},
			PaymentMethod: userPaymentMethod,
		}
	})

	res.Data = modelDto

	return c.JSON(http.StatusOK, res)
}

// GetMyWithdrawals godoc
// @Summary Get my Withdrawals
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Success 200 {array} response_object.WithdrawalRO "Array of Withdrawals"
// @Router /withdrawal/me [get]
func GetMyWithdrawals(c echo.Context) error {
	var pag pagination.PaginationParams
	var filters dto.WithdrawalFiltersDTO
	userUuid := c.Get("user").(*models.User).Uuid.String()

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}

	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		return errFilters
	}

	filters.User = &userUuid
	res, err := service.GetAllWithdrawals(pag, &filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res.Data = lo.Map(res.Data.([]models.Withdrawal), func(item models.Withdrawal, _ int) response_object.WithdrawalRO {
		return response_object.WithdrawalRO{
			Uuid:           item.Uuid.String(),
			AmountSource:   item.AmountSource,
			AmountTarget:   item.AmountTarget,
			Details:        item.Details,
			State:          item.State,
			CompletionDate: item.CompletionDate,
			User: response_object.SimpleUserDto{
				Uuid:  c.Get("user").(*models.User).Uuid,
				Email: c.Get("user").(*models.User).Email,
			},
			CreatedAt: &item.CreatedAt,
			BaseEntityRO: response_object.BaseEntityRO{
				CreatedAt: item.CreatedAt,
				CreatedBy: item.CreatedBy,
				UpdatedAt: item.UpdatedAt,
				UpdatedBy: item.UpdatedBy,
				Deleted:   item.Deleted,
				DeletedAt: item.DeletedAt,
				DeletedBy: item.DeletedBy,
			},
		}
	})

	return c.JSON(http.StatusOK, res)
}

// CreateWithdrawal godoc
// @Summary Create Withdrawal
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Success 201 {object} response_object.WithdrawalRO "Withdrawal"
// @Router /withdrawal [post]
func CreateWithdrawal(c echo.Context) error {
	user := c.Get("user").(*models.User)

	hasPendingWithdrawal, err := service.GetLatestPendingWithdrawal(user.Uuid.String())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if hasPendingWithdrawal {
		return c.JSON(http.StatusUnprocessableEntity, "User has pending withdrawal")
	}

	minimumBalanceConfig := service.GetLoadedConfig("withdrawal_balance_minimum")
	if minimumBalanceConfig.Id == 0 {
		return utils.CustomErrorStruct{ErrorMessage: "Configuration withdrawal_balance_minimum not found or not accessible"}
	}

	minimumBalance, err := strconv.ParseFloat(minimumBalanceConfig.Value, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if c.Get("user").(*models.User).Balance < minimumBalance {
		return c.JSON(http.StatusUnprocessableEntity, "User balance is not enough")
	}

	userMethods, err := repository.GetUserPaymentMethodsByUserUuid(user.Uuid.String())
	if err != nil || len(userMethods) == 0 {
		return c.JSON(http.StatusNotFound, "No User Payment Methods! Go to withdrawals settings to setup it up!")
	}

	res, err := service.CreateWithdrawal(
		utils.StringPointer(user.Uuid.String()),
		user.Currency,
		user.Balance,
		&userMethods,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res.User = response_object.SimpleUserDto{
		Uuid:  c.Get("user").(*models.User).Uuid,
		Email: c.Get("user").(*models.User).Email,
	}

	return c.JSON(http.StatusCreated, res)

}

// PatchWithdrawal godoc
// @Summary Update Withdrawal
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Param id path string true "Withdrawal id"
// @Param Withdrawal body dto.UpdateWithdrawalDTO true "Update Withdrawal dto"
// @Success 200 {object} models.Withdrawal "Withdrawal"
// @Router /withdrawal/:id [patch]
func PatchWithdrawal(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var model dto.UpdateWithdrawalDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if model.UserMethodID != nil {
		_, err := service.GetUserPaymentMethodByIdUnscoped(*model.UserMethodID)
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()
	keycloak := auth.KeycloakInstance

	res, err := service.UpdateWithdrawal(model, uuid, uuidUser, keycloak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)

}

// DeleteWithdrawal godoc
// @Summary Delete Withdrawal
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Param id path string true "Withdrawal id"
// @Success 204
// @Router /withdrawal/:id [delete]
func DeleteWithdrawal(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))
	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteWithdrawal(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

// GetMyWithdrawalsStats godoc
// @Summary Get my Withdrawals stats
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Success 204
// @Router /withdrawal/me/stats [get]
func GetMyWithdrawalsStats(c echo.Context) error {
	user := c.Get("user").(*models.User)

	paidWithdrawals, err := service.GetPaidWithdrawalsAmount(user.Uuid.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	amountReferralRewards, err := service.GetSumReferralRewards(user.Uuid.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	amountReferrals := amountReferralRewards
	amountRewards := user.Balance - amountReferralRewards

	if amountReferrals < 0 {
		amountReferrals = 0
	}
	if amountRewards < 0 {
		amountRewards = 0
	}

	res := response_object.WithdrawalStatsRO{
		PaidWithdrawals: paidWithdrawals,
		AmountRewards:   amountRewards,
		AmountReferrals: amountReferrals,
	}

	return c.JSON(http.StatusOK, res)
}

// BulkUpdateStateWithdrawal godoc
// @Summary Bulk update state request Withdrawal
// @Tags Withdrawal
// @Accept json
// @Produce json
// @Param Withdrawal body dto.BulkUpdateStateRequestDTO true "Update Withdrawal dto"
// @Success 200
// @Router /withdrawal/bulk/state-request [patch]
func BulkUpdateStateWithdrawal(c echo.Context) error {
	var model dto.BulkUpdateStateRequestDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()
	keycloak := auth.KeycloakInstance

	err := service.BulkUpdateStateWithdrawal(model, uuidUser, keycloak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
