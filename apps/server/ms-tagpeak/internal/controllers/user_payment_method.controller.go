package controllers

import (
	"fmt"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/files"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// CreateUserPaymentMethod godoc
// @Summary Create User Payment Method
// @Tags UserPaymentMethod
// @Accept json
// @Produce json
// @Success 200 {object} models.UserPaymentMethod "UserPaymentMethods"
// @Router /user/payment-method [post]
func CreateUserPaymentMethod(c echo.Context) error {
	var model dto.CreateUserPaymentMethodDTO
	uuidUser := c.Get("user").(*models.User).Uuid.String()

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&model); err != nil {
		fmt.Printf("Body not valid %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	paymentMethod, err := service.GetPaymentMethod(model.PaymentMethod)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	file, err := service.GetFileByUuid(*model.IbanStatement)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	userPaymentMethod, err := service.CreateUserPaymentMethod(model, paymentMethod, file, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, userPaymentMethod)
}

// GetUserPaymentMethodsByUserUuid godoc
// @Summary Get User Payment Method by User
// @Tags UserPaymentMethod
// @Accept json
// @Produce json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Success 200 {array} response_object.UserPaymentMethodRo "Array of UserPaymentMethods"
// @Router /user/payment-method [get]
func GetUserPaymentMethodsByUserUuid(c echo.Context) error {
	var pag pagination.PaginationParams
	uuidUser := c.Get("user").(*models.User).Uuid.String()

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}

	res, err := service.GetUserPaymentMethodsByUserUuidPaginated(uuidUser, pag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// GetUserPaymentMethodById godoc
// @Summary Get User Payment Method by User
// @Tags UserPaymentMethod
// @Accept json
// @Produce json
// @Param id path string true "UserPaymentMethod id"
// @Success 200 {object} response_object.UserPaymentMethodRo "UserPaymentMethods"
// @Router /user/payment-method/:id [get]
func GetUserPaymentMethodById(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetUserPaymentMethodById(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// DeleteUserPaymentMethod godoc
// @Summary Delete User Payment Method
// @Tags UserPaymentMethod
// @Accept json
// @Produce json
// @Param id path string true "UserPaymentMethod id"
// @Success 204
// @Router /user/payment-method/:id [delete]
func DeleteUserPaymentMethod(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteUserPaymentMethod(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

// UploadIbanStatement godoc
// @Summary Upload Iban Statement
// @Tags UserPaymentMethod
// @Accept json
// @Produce json
// @Success 200 {string} string "uuid of File"
// @Router /user/payment-method/file [post]
func UploadIbanStatement(c echo.Context) error {
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

	return c.JSON(http.StatusOK, fileUuid)
}

// DownloadFile godoc
// @Summary Download the file
// @Tags UserPaymentMethod
// @Accept json
// @Produce json
// @Success 200 {file} File "File"
// @Router /user/payment-method/file/:id [get]
func DownloadFile(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	file, err := service.GetFileByUuid(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	pathToFiles := dotenv.GetEnv("PATH_FILE")
	filePath := filepath.Join(pathToFiles, uuid.String(), *file.Name)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "File not found on disk"})
	}

	// Set response headers
	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%q", file.Name))
	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")

	return c.File(filePath)
}
