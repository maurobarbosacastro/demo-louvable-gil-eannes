package controllers

import (
	"fmt"
	"ms-tagpeak/external/images"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/files"
	"ms-tagpeak/pkg/http_client"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetLogo godoc
// @Summary Get a logo based on domain
// @Description Get a logo based on domain. Ex: worten.pt / aliexpress.com / booking.com
// @Tags Auxiliary
// @Produce  json
// @Param name query string true "Store Name"
// @Success 200 {object} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /aux/logo [get]
func GetLogo(c echo.Context) error {
	nameSearch := c.QueryParam("name")

	logo, err := images.GetLogoFromDomain(nameSearch)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, logo)
}

// CheckVatValidity godoc
// @Summary Get vat validity
// @Description Get vat validity
// @Tags Auxiliary
// @Produce  json
// @Param name query string true "vat"
// @Success 200 {object} response_object.VatValidityResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /aux/vat [get]
func CheckVatValidity(c echo.Context) error {
	fmt.Printf("Start CheckVatValidity - %v\n", c.QueryParam("vat_number"))
	res, err := service.CheckVatValidity(c.QueryParam("vat_number"))
	if err != nil {
		fmt.Printf("Error - %v\n", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	fmt.Printf("End CheckVatValidity - %#v\n", res)
	return c.JSON(http.StatusOK, response_object.VatValidityResponse{
		Success: res.Success,
	})
}

func UploadFile(c echo.Context) error {
	fmt.Printf("Start UploadFile \n")

	// Get the uploaded fileHeader from the request
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No fileHeader uploaded"})
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()
	file, err := service.SaveFile(fileHeader, uuidUser)

	fileUuid, errUpload := files.HandleFile(fileHeader, file)
	if errUpload != nil {
		return c.JSON(http.StatusBadRequest, errUpload)
	}

	fmt.Printf("End UploadFile - %#v\n", fileUuid)
	return c.JSON(http.StatusOK, fileUuid)

}

// GetInfoFromIp godoc
// @Summary Get IP information
// @Description Get geographical and network information for a given IP address
// @Tags Auxiliary
// @Produce json
// @Param ip query string true "IP Address"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /aux/ip [get]
func GetInfoFromIp(c echo.Context) error {
	ip := c.QueryParam("ip")

	if ip == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "IP parameter is required"})
	}

	// Create HTTP client
	client := &http_client.HttpClient{
		HttpClient: &http.Client{},
	}

	// Build the URL
	url := fmt.Sprintf("https://ipapi.co/%s/json/", ip)

	// Set headers required by ipapi.co
	headers := map[string]string{
		"User-Agent": "ipapi.co/#go-v1.5",
	}

	// Make the request
	var result map[string]interface{}
	_, err := client.Get(url, headers, &result)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
