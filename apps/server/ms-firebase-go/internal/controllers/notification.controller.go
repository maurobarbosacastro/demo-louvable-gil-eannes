package controllers

import (
	"ms-firebase-go/internal/dto"
	"ms-firebase-go/internal/models"
	"ms-firebase-go/internal/service"
	"ms-firebase-go/pkg/logster"
	"ms-firebase-go/pkg/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateNotification(c echo.Context) error {
	var body dto.CreateNotificationDto

	// Bind the incoming JSON request body to the messageDto
	if err := c.Bind(&body); err != nil {
		logster.Error(err, "Error while binding json")
		return c.JSON(http.StatusBadRequest, err)
	}

	notification, err := service.SaveNotifications(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, notification)
}

func UpdateNotification(c echo.Context) error {
	var body dto.NotificationUpdateFields

	notifUuid := uuid.MustParse(c.Param("uuid"))

	if err := c.Bind(&body); err != nil {
		logster.Error(err, "Error binding body")
		return c.JSON(http.StatusBadRequest, err)
	}

	notification, errUpdate := service.UpdateNotification(notifUuid, body)

	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, errUpdate)
	}

	return c.JSON(http.StatusOK, notification)
}

func SendNotification(c echo.Context) error {
	notifUUid := c.Param("uuid")

	notification, errGetNotification := service.GetNotificationByUuid(uuid.MustParse(notifUUid))

	if errGetNotification != nil {
		return c.JSON(http.StatusInternalServerError, errGetNotification)
	}

	if notification.Target != nil && len(notification.Target) > 0 {
		for _, target := range notification.Target {
			go service.SendNotificationToTarget(*notification, target)
		}
	}

	if notification.Topics != nil && len(notification.Topics) > 0 {
		for _, topic := range notification.Topics {
			go service.SendNotificationToTopic(*notification, topic)
		}
	}

	_, errUpdate := service.UpdateNotification(
		notification.UUID,
		dto.NotificationUpdateFields{
			State: utils.Ptr(models.NotificationStateProcessed),
		})

	if errUpdate != nil {
		logster.Error(errUpdate, "Error updating notification status")
	}

	return c.JSON(http.StatusNoContent, nil)
}
