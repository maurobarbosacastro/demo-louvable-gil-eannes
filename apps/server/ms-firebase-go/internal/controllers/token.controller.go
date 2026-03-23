package controllers

import (
	"fmt"
	"ms-firebase-go/internal/dto"
	"ms-firebase-go/internal/service"
	"ms-firebase-go/pkg/logster"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SaveToken(c echo.Context) error {
	logster.StartFuncLog()

	var body dto.TokenSave

	if err := c.Bind(&body); err != nil {
		logster.Error(err, "Error binding body")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, err)
	}

	logster.Debug(fmt.Sprintf("Body: %+v", body))

	token, err := service.CreateToken(body)
	if err != nil {
		logster.Error(err, "Error creating token")
		logster.EndFuncLog()
		return c.JSON(http.StatusConflict, err)
	}

	// Add connecting users to topics.
	topicAll, err := service.GetTopicByName("all")
	if err != nil {
		logster.Error(err, "Error get topic with name all")
	}

	errAddUserToTopic := service.AddUserToTopic(topicAll.UUID, body.UserUUID)

	if errAddUserToTopic != nil {
		logster.Error(errAddUserToTopic, "Error adding token to topic")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, errAddUserToTopic)
	}

	logster.EndFuncLogMsg(token)
	return c.JSON(http.StatusCreated, token)
}

func DeleteToken(c echo.Context) error {
	token := c.Param("token")

	err := service.RemoveToken(token)
	if err != nil {
		logster.Error(err, "Error removing token")
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func DeleteTokenByUserUuid(c echo.Context) error {
	userUuid := c.Param("uuid")

	err := service.RemoveTokenForUser(userUuid)
	if err != nil {
		logster.Error(err, "Error removing token")
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
