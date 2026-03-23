package controllers

import (
	"ms-tagpeak/internal/auth"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateReferralClick godoc
// @Summary Create Referral Click
// @Tags Referral
// @Produce json
// @Param referralClick body dto.CreateReferralClickDTO true "Create Referral Click dto"
// @Success 200 {object} models.ReferralClicks "ReferralClick"
// @Router /referral/click [POST]
func CreateReferralClick(c echo.Context) error {

	var referralClick dto.CreateReferralClickDTO

	if err := c.Bind(&referralClick); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := service.CreateReferralClick(nil, referralClick.Code)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)

}

// ValidateReferralCode godoc
// @Summary Validate Referral Code
// @Tags Referral
// @Produce json
// @Param referralCode body string true "Referral Code"
// @Success 200 {bool} bool
// @Router /referral/validate [GET]
func ValidateReferralCode(c echo.Context) error {
	referralCode := c.QueryParam("code")

	keycloak := auth.KeycloakInstance

	res, err := service.ValidateReferralCode(referralCode, keycloak)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}
