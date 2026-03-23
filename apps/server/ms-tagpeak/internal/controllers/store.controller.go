package controllers

import (
	"fmt"
	"ms-tagpeak/external/images"
	"ms-tagpeak/internal/auth"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	_ "ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/internal/service"
	"ms-tagpeak/pkg/files"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetStore godoc
// @Summary Get Store by ID
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Success 200 {object} response_object.GetStoreRO
// @Router /store/:id [get]
func GetStore(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetStore(uuid)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)

}

// GetAllStores godoc
// @Summary Get all Stores
// @Tags Store
// @Accept json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Param filters body dto.StoreFiltersDTO true "Filters for stores"
// @Produce json
// @Success 200 {object} pagination.PaginationResult{data=[]models.Store} "Array of Stores"
// @Router /store [get]
func GetAllStores(c echo.Context) error {

	var pag pagination.PaginationParams
	var filters dto.StoreFiltersDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		return errFilters
	}

	res, err := service.GetAllStores(pag, filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetAllStoresAdmin godoc
// @Summary Get all Stores for admin
// @Tags Store
// @Accept json
// @Param pagination body pagination.PaginationParams true "Pagination params"
// @Param filters body dto.StoreFiltersDTO true "Filters for stores"
// @Produce json
// @Success 200 {array} pagination.PaginationResult{data=[]dto.AdminStoreDTO} "Array of Stores"
// @Router /store/admin [get]
func GetAllStoresAdmin(c echo.Context) error {

	var pag pagination.PaginationParams
	var filters dto.StoreFiltersDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		return errFilters
	}

	res, err := service.GetAllStoresForAdmin(pag, filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)

}

// CreateStore godoc
// @Summary Create Store
// @Tags Store
// @Accept json
// @Produce json
// @Param country body dto.CreateStoreDTO true "Create Store dto"
// @Success 201 {object} models.Store "Store"
// @Router /store [post]
func CreateStore(c echo.Context) error {
	var model dto.CreateStoreDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.CreateStore(model, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)

}

// PatchStore godoc
// @Summary Update Store
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Param Store body dto.UpdateStoreDTO true "Update Store dto"
// @Success 200 {object} models.Store "Store"
// @Router /store/:id [patch]
func PatchStore(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	var model dto.UpdateStoreDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	res, err := service.UpdateStore(model, uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, res)

}

// DeleteStore godoc
// @Summary Delete Store
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Success 204
// @Router /store/:id [delete]
func DeleteStore(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteStore(uuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

// GetAffiliatePartnerByStore godoc
// @Summary Get AffiliatePartnerByStore by Store ID
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Success 200 {object} models.Partner "AffiliatePartnerByStore"
// @Router /store/:id/partner [get]
func GetAffiliatePartnerByStore(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetAffiliatePartnerByStore(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetCountryByStore godoc
// @Summary Get CountryByStore by Store ID
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Success 200 {object} models.Country "CountryByStore"
// @Router /store/:id/country [get]
func GetCountryByStore(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetCountryByStore(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetLanguageByStore godoc
// @Summary Get LanguageByStore by Store ID
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Success 200 {object} models.Language "LanguageByStore"
// @Router /store/:id/language [get]
func GetLanguageByStore(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetLanguageByStore(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetCategoryByStore godoc
// @Summary Get CategoryByStore by Store ID
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Success 200 {object} models.Category "CategoryByStore"
// @Router /store/:id/category [get]
func GetCategoryByStore(c echo.Context) error {

	uuid := utils.ParseIDToUUID(c.Param("id"))

	res, err := service.GetCategoryByStore(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

// UploadStoreByExcel godoc
// @Summary Get UploadStoreByExcel
// @Tags Store
// @Accept Multipart/form-data
// @Produce json
// @Success 201 {object} models.Store
// @Router /store/upload-excel [post]
func UploadStoreByExcel(c echo.Context) error {

	formFile, errForm := c.FormFile("file")
	if errForm != nil {
		return c.JSON(http.StatusBadRequest, errForm)
	}

	stores, err := service.UploadStoreByExcel(formFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, stores)

}

// AddCountryToStore godoc
// @Summary Get AddCountryToStore
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Param country body models.StoreCountry true "StoreCountry"
// @Success 201 {object} models.StoreCountry
// @Router /store/country [post]
func AddCountryToStore(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	var model dto.StoreCountryDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := service.AddCountryToStore(uuid, model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, model)

}

// RemoveCountryFromStore godoc
// @Summary Delete RemoveCountryFromStore
// @Description Delete RemoveCountryFromStore
// @Tags StoreCountry
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Param country body models.StoreCountry true "StoreCountry"
// @Success 204
// @Router /store/country [delete]
func RemoveCountryFromStore(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	var model dto.StoreCountryDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := service.RemoveCountryFromStore(uuid, model)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, model)

}

// AddCategoryToStore godoc
// @Summary Get AddCategoryToStore
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Param country body models.StoreCountry true "StoreCountry"
// @Success 201 {object} models.StoreCountry
// @Router /store/category [post]
func AddCategoryToStore(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	var model dto.StoreCategoryDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := service.AddCategoryToStore(uuid, model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, model)

}

// RemoveCategoryFromStore godoc
// @Summary Delete RemoveCategoryFromStore
// @Description Delete RemoveCategoryFromStore
// @Tags StoreCategory
// @Accept json
// @Produce json
// @Param id path string true "Store id"
// @Param country body models.StoreCountry true "StoreCountry"
// @Success 204
// @Router /store/category [delete]
func RemoveCategoryFromStore(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))

	var model dto.StoreCategoryDTO

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := service.RemoveCategoryFromStore(uuid, model)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, model)

}

// UploadStoreLogo godoc
// @Summary Upload store logo
// @Description Upload store logo
// @Tags Store
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string true "Store ID"
// @Param file formData file true "Logo file"
// @Success 200 {object} models.Store
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /store/:id/logo [post]
func UploadStoreLogo(c echo.Context) error {
	// Get the uploaded file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No file uploaded"})
	}

	fileUuid, errUpload := images.UploadImage(file, c.FormValue("fileName"), "logo")
	if errUpload != nil {
		return c.JSON(http.StatusBadRequest, errUpload)
	}

	uuid := utils.ParseIDToUUID(c.Param("id"))

	uuidUser := c.Get("user").(*models.User).Uuid.String()

	store, err := service.UpdateStoreLogo(uuid, fileUuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Return the destination service's response
	return c.JSON(http.StatusOK, store)
}

// UploadStoreBanner godoc
// @Summary Upload store banner
// @Description Upload store banner
// @Tags Store
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string true "Store ID"
// @Param file formData file true "Logo file"
// @Success 200 {object} models.Store
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /store/:id/banner [post]
func UploadStoreBanner(c echo.Context) error {
	// Get the uploaded file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No file uploaded"})
	}

	fileUuid, errUpload := images.UploadImage(file, c.FormValue("fileName"), "resized")
	if errUpload != nil {
		return c.JSON(http.StatusBadRequest, errUpload)
	}

	uuid := utils.ParseIDToUUID(c.Param("id"))
	uuidUser := c.Get("user").(*models.User).Uuid.String()

	store, err := service.UpdateStoreBanner(uuid, fileUuid, uuidUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Return the destination service's response
	return c.JSON(http.StatusOK, store)
}

// DeleteStoreLogo godoc
// @Summary Delete store logo
// @Description Delete store logo
// @Tags Store
// @Produce  json
// @Param id path string true "Store ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /store/:id/logo [delete]
func DeleteStoreLogo(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteStoreLogo(uuid, uuidUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

// DeleteStoreBanner godoc
// @Summary Delete store banner
// @Description Delete store banner
// @Tags Store
// @Produce  json
// @Param id path string true "Store ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /store/:id/banner [delete]
func DeleteStoreBanner(c echo.Context) error {
	uuid := utils.ParseIDToUUID(c.Param("id"))
	uuidUser := c.Get("user").(*models.User).Uuid.String()

	err := service.DeleteStoreBanner(uuid, uuidUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

// ExportStoreCSV godoc
// @Summary Export Stores CSV
// @Description Export Stores CSV
// @Tags Stores
// @Accept  json
// @Produce  json
// @Param  pagination query pagination.PaginationParams false "Request pagination parameters"
// @Param  filters query dto.StoreFiltersDTO false "Request filters"
// @Success 200 {file} File "File"
// @Router /store/export-csv [get]
func ExportStoreCSV(c echo.Context) error {
	fmt.Printf("START controller.ExportStoreCSV - %v\n", c.QueryParam("fileName"))
	fileName := c.QueryParam("fileName")
	var pag pagination.PaginationParams
	var filters dto.StoreFiltersDTO

	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		fmt.Println("Error binding pagination")
		return errPag
	}

	// Bind query params for filters
	errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters)
	if errFilters != nil {
		fmt.Println("Error binding filters")
		return errFilters
	}

	res, err := service.ExportStoresCSV(pag, filters)
	if err != nil {
		fmt.Println("Error exporting stores")
		return c.JSON(http.StatusInternalServerError, err)
	}

	headers := []string{
		"uuid",
		"country_id",
		"store_category_id",
		"title",
		"cashback_type",
		"cashback_value",
		"average_payout",
		"average_cashout",
		"short_description",
		"description",
		"store_icon",
		"store_banner",
		"network_id",
		"deeplink",
		"store_url",
		"term_and_condition",
		"keywords",
		"meta_title",
		"meta_keyword",
		"meta_description",
		"status",
		"average_payout_num",
		"slug",
		"partner_identity",
		"override_fee",
		"language_code",
	}

	var data [][]string
	for _, store := range res {
		row := response_object.MapExportStoresRO(store)
		data = append(data, row)
	}

	buf, errFile := files.CreateCSVFile(headers, data, fileName)
	if errFile != nil {
		fmt.Printf("Error creating CSV file - %v\n", errFile)
		return c.JSON(http.StatusInternalServerError, errFile)
	}

	// Set headers for CSV download
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

	_, err = c.Response().Write(buf.Bytes())
	if err != nil {
		fmt.Printf("Error writing the file response %v\n", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	fmt.Println("END controller.ExportStoreCSV")
	// return nil to avoid adding “null / {}” to the last line of the file
	return nil
}

// GetStoresForApproval godoc
// @Summary Get stores pending approval
// @Description Get all stores that need admin approval
// @Tags Store
// @Success 200 {object} pagination.PaginationResult{data=[]response_object.ApprovalRequestRo} "Array of Stores"
// @Param pagination query pagination.PaginationParams false "Request pagination parameters"
// @Success 200 {array} models.Store "Array of stores pending approval"
// @Router /store/approvals [get]
func GetStoresForApproval(c echo.Context) error {
	var pag pagination.PaginationParams
	// Bind query params for pagination
	errPag := (&echo.DefaultBinder{}).BindQueryParams(c, &pag)
	if errPag != nil {
		return errPag
	}

	filters := dto.StoreFiltersDTO{
		State: "PENDING",
	}

	keycloak := auth.KeycloakInstance
	stores, err := service.GetStoresForApproval(pag, filters, keycloak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, stores)
}

// IsUniquePosition godoc
// @Summary Check if position is unique
// @Tags Store Check if position is unique
// @Success 200 {bool} bool
// @Param position path string true "Store position"
// @Router /store/position/:position [get]
func IsUniquePosition(c echo.Context) error {
	position := c.Param("position")

	pos, err := strconv.Atoi(position)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	isUnique := service.IsPositionUnique(pos)

	return c.JSON(http.StatusOK, isUnique)
}
