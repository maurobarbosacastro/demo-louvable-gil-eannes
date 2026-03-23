package controllers

import (
	"errors"
	"fmt"
	"ms-firebase-go/internal/dto"
	"ms-firebase-go/internal/service"
	"ms-firebase-go/pkg/logster"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
)

func GetAllTopics(c echo.Context) error {
	topics, err := service.GetAllTopics()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, topics)
}

func CreateTopic(c echo.Context) error {
	var data dto.CreateTopicDto

	// Bind the incoming JSON request data to the messageDto
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err := service.CreateTopic(data)
	if err != nil {
		return c.JSON(http.StatusCreated, "")
	}

	return c.JSON(http.StatusCreated, "")
}

func AddTokenToTopic(c echo.Context) error {
	topic := c.Param("uuid")

	var tokenBody dto.AddRemoveTokenToTopicDto

	// Bind the incoming JSON request data to the messageDto
	if err := c.Bind(&tokenBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := service.AddTokenToTopic(uuid.MustParse(topic), tokenBody.Token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func RemoveTokenToTopic(c echo.Context) error {
	topic := c.Param("uuid")

	var tokenBody dto.AddRemoveTokenToTopicDto

	// Bind the incoming JSON request data to the messageDto
	if err := c.Bind(&tokenBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	logster.Debug(fmt.Sprintf("%+v", tokenBody))
	err := service.RemoveTokenFromTopic(uuid.MustParse(topic), tokenBody.Token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func AddUserToTopic(c echo.Context) error {
	logster.StartFuncLog()

	topic := c.Param("uuid")

	var tokenBody dto.AddRemoveUserToTopicDto

	if err := c.Bind(&tokenBody); err != nil {
		logster.Error(err, "Error binding tokenBody")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, err)
	}

	err := service.AddUserToTopic(uuid.MustParse(topic), tokenBody.UserUUID)

	if err != nil {
		var pge *pgconn.PgError
		if errors.As(err, &pge) && pge.Code == "23505" {
			// duplicate key -> handle gracefully
			return c.JSON(http.StatusConflict, err)
		}

		return c.JSON(http.StatusInternalServerError, err)
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusCreated, nil)
}

func RemoveUserFromTopic(c echo.Context) error {
	topic := c.Param("uuid")
	user := c.Param("user")

	err := service.RemoveUserFromTopic(uuid.MustParse(topic), user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}
