package controllers

import (
	"fmt"
	"ms-tagpeak/external/notifications"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/logster"
	"net/http"

	"github.com/samber/lo"

	"github.com/labstack/echo/v4"
)

func SaveUserFirebaseToken(c echo.Context) error {
	var body notifications.UserToken

	if err := c.Bind(&body); err != nil {
		logster.Error(err, "Error binding body")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err := notifications.SaveToken(notifications.TokenSave{
		UserUUID: c.Get("user").(*models.User).Uuid.String(),
		Token:    body.Token,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	membershipLevel := c.Get("membershipStatus").(*models.MembershipStatus).Level
	logster.Info(fmt.Sprintf("User membership level %s", *membershipLevel))

	topics, err := notifications.GetAllTopics()
	if err != nil {
		logster.Error(err, "Error getting topics")
		logster.EndFuncLog()
		return c.JSON(http.StatusCreated, nil)
	}
	topicMembershipLevel, _ := lo.Find(topics, func(item notifications.Topic) bool {
		return item.Name == *membershipLevel
	})
	logster.Info(fmt.Sprintf("Topic membership level %s", topicMembershipLevel))

	errAddTokenTopic := notifications.AddUserToTopic(topicMembershipLevel.UUID.String(), body.UserUUID)

	if errAddTokenTopic != nil {
		logster.Error(errAddTokenTopic, "Error adding token to topic")
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusCreated, nil)
}

func DeleteFirebaseToken(c echo.Context) error {
	logster.StartFuncLog()
	err := notifications.DeleteTokensByUser(c.Get("user").(*models.User).Uuid.String())
	if err != nil {
		logster.Error(err, "Error deleting tokens")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusNoContent, nil)
}

func GetTopics(c echo.Context) error {
	logster.StartFuncLog()
	topics, err := notifications.GetAllTopics()
	if err != nil {
		logster.Error(err, "Error getting topics")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, topics)
}
