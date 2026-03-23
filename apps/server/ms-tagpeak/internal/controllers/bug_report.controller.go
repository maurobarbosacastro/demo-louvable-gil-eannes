package controllers

import (
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/logster"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBugReport(c echo.Context) error {
	logster.StartFuncLog()

	var body models.BugReportRequest
	if err := c.Bind(&body); err != nil {
		logster.Error(err, "Error binding body")
		logster.EndFuncLog()
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := c.Validate(&body); err != nil {
		logster.Error(err, "Validation error")
		logster.EndFuncLog()
		return err
	}

	if err := service.SendBugReport(body); err != nil {
		logster.Error(err, "Error sending bug report")
		logster.EndFuncLog()
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send bug report"})
	}

	logster.EndFuncLog()
	return c.JSON(http.StatusOK, map[string]string{"message": "Bug report sent successfully"})
}
