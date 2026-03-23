package controllers

import (
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	_ "ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetRewardHistory godoc
// @Summary Get RewardHistory by Reward ID
// @Tags RewardHistory
// @Accept json
// @Produce json
// @Param id path string true "Reward id"
// @Success 200 {object} models.RewardHistory "RewardHistoryByReward"
// @Router /reward/:id/history [get]
func GetRewardHistory(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var pag pagination.PaginationParams

	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}

	res, err := service.GetRewardHistory(uuid, pag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// CreateRewardHistory godoc
// @Summary Create Reward History
// @Tags RewardHistory
// @Accept json
// @Produce json
// @Param country body dto.RewardHistoryDTO true "Create RewardHistory dto"
// @Success 201 {object} models.RewardHistory "RewardHistory"
// @Router /reward/:id/history [post]
func CreateRewardHistory(c echo.Context) error {
	var model dto.RewardHistoryDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.CreateRewardHistory(model, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)

}

// GetRewardHistoryGraph godoc
// @Summary Get RewardHistory by Reward ID
// @Tags RewardHistory
// @Accept json
// @Produce json
// @Param id path string true "Reward id"
// @Success 200 {object} dto.RewardHistoryDTO "RewardHistoryByReward"
// @Router /reward/:id/history/graph [get]
func GetRewardHistoryGraph(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetRewardHistoryGraph(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}
