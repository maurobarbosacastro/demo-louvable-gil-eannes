package controllers

import (
	"fmt"
	"ms-interactive-brokers/internal/dto"
	"ms-interactive-brokers/internal/response_object"
	"ms-interactive-brokers/internal/service"
	"ms-interactive-brokers/pkg/logster"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func GetLastPriceBulk(c echo.Context) error {
	logster.StartFuncLog()
	var lastPriceBulkDTO dto.LastPriceBulkDTO

	if err := c.Bind(&lastPriceBulkDTO); err != nil {
		logster.Error(err, fmt.Sprintf("failed to bind request: %v", err))
		return c.JSON(http.StatusBadRequest, err)
	}

	response, err := service.GetLastPricesBasedOnAccountPositions(lastPriceBulkDTO.Conids)
	if err != nil {
		logster.Error(err, "failed to get last price bulk")
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := lo.Filter((*response), func(item response_object.LastPriceBulkRO, index int) bool {
		return item.LastPrice != ""
	})
	logster.EndFuncLogMsg(fmt.Sprintf("Response: %+v", response))
	return c.JSON(http.StatusOK, res)
}
