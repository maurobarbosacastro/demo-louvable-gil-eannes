package controllers

import (
	"fmt"
	"ms-firebase-go/external/firebase"
	"ms-firebase-go/internal/dto"
	"ms-firebase-go/pkg/logster"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SendTokenMessage(c echo.Context) error {
	logster.StartFuncLog()
	token := c.Param("token")
	if token == "" {
		logster.Error(fmt.Errorf("token parameter is required"), "Error getting token")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "token parameter is required"})
	}

	messageDto := new(dto.MessageDTO)

	// Bind the incoming JSON request body to the messageDto
	if err := c.Bind(messageDto); err != nil {
		logster.Error(err, "Error binding messageDto")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, err)
	}

	response, err := firebase.SendNotification(c.Request().Context(), token, messageDto)
	if err != nil {
		logster.Error(err, "Error sending notification")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, response)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, response)
}

func SendTopicMessage(c echo.Context) error {
	logster.StartFuncLog()

	group := c.Param("group")
	if group == "" {
		logster.Error(fmt.Errorf("group parameter is required"), "Error getting group")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "group parameter is required"})
	}

	messageDto := new(dto.MessageDTO)

	// Bind the incoming JSON request body to the messageDto
	if err := c.Bind(messageDto); err != nil {
		logster.Error(err, "Error binding messageDto")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, err)
	}

	response, err := firebase.SendGroupMessage(c.Request().Context(), group, messageDto)
	if err != nil {
		logster.Error(err, "Error sending group message")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, response)
	}

	logster.EndFuncLogMsg(response)
	return c.JSON(http.StatusOK, response)
}
