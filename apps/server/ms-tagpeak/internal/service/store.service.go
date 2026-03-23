package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"ms-tagpeak/external/images"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"gorm.io/gorm"
)

func GetStore(uuid uuid.UUID) (*response_object.GetStoreRO, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Searching store with uuid: %s", uuid.String()))
	res, err := repository.GetStore(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logster.Error(err, "Store not found")
			logster.EndFuncLog()
			return nil, utils.CustomErrorStruct{}.NotFoundError("Store", uuid)
		}
		logster.Error(err, "Error getting store")
		logster.EndFuncLog()
		return nil, err
	}

	var model = utils.ModelStoreToSimpleDTO(&res)
	logster.EndFuncLogMsg("Store found")
	return &model, nil
}

func GetAffiliatePartnerByStore(uuid uuid.UUID) (models.Partner, error) {
	logster.StartFuncLog()

	res, err := repository.GetAffiliatePartnerByStore(uuid)
	if err != nil {
		logster.Error(err, "Error getting affiliate partner by store")
		logster.EndFuncLog()
		return models.Partner{}, err
	}

	logster.EndFuncLog()
	return res, nil
}

func GetCategoryByStore(uuid uuid.UUID) ([]models.Category, error) {
	logster.StartFuncLog()
	res, err := repository.GetCategoryByStore(uuid)
	if err != nil {
		logster.Error(err, "Error getting category by store")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return res, nil
}

func GetCountryByStore(uuid uuid.UUID) (*[]models.Country, error) {
	logster.StartFuncLog()

	res, err := repository.GetCountryByStore(uuid)
	if err != nil {
		logster.Error(err, "Error getting country by store")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return &res, nil
}

func GetLanguageByStore(uuid uuid.UUID) (*models.Language, error) {
	logster.StartFuncLog()

	res, err := repository.GetLanguageByStore(uuid)
	if err != nil {
		logster.Error(err, "Error getting language by store")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return &res, nil
}

func GetAllStores(pag pagination.PaginationParams, filters dto.StoreFiltersDTO) (*pagination.PaginationResult, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("*** START service.GetAllStores ***"))

	res, err := repository.GetAllStoresWithPagination(pag, filters)
	if err != nil {
		logster.Error(err, "Error getting all stores")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return res, nil
}

func GetAllStoresForAdmin(pag pagination.PaginationParams, filters dto.StoreFiltersDTO) (*pagination.PaginationResult, error) {
	logster.StartFuncLog()

	res, err := repository.GetAllStores(pag, filters)

	var aux []dto.AdminStoreDTO

	if dataSlice, ok := res.Data.([]models.Store); ok {
		for _, item := range dataSlice {
			aux = append(aux, utils.ModelStoreToStoreAdminDTO(&item)) // Process each item
		}

		res.Data = aux
	} else {
		return nil, errors.New("res.Data is not of type []models.Store")
	}

	if err != nil {
		logster.Error(err, "Error getting all stores")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return res, nil
}

func CreateStore(dtoParam dto.CreateStoreDTO, uuidUser string) (*models.Store, error) {
	logster.StartFuncLog()

	var countries []string
	var categories []string

	//Check if country exists
	for _, country := range dtoParam.Country {
		_, err := GetCountryByCode(country)
		if err == nil {
			countries = append(countries, country)
		}
	}

	//Check if category exists
	for _, category := range dtoParam.Category {
		_, err := repository.GetCategoryByCode(category)
		if err == nil {
			categories = append(categories, category)
		}
	}

	if len(categories) == 0 {
		categories = make([]string, 0)
	}
	if len(countries) == 0 {
		countries = make([]string, 0)
	}

	model := utils.CreateStoreDtoToModel(&dtoParam)
	model.CreatedBy = uuidUser
	model.CountriesCodes = countries
	model.CategoriesCodes = categories

	res, err := repository.CreateStore(model)
	if err != nil {
		logster.Error(err, "Error creating store")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return &res, nil
}

func UpdateStore(dtoParam dto.UpdateStoreDTO, uuid uuid.UUID, uuidUser string) (*models.Store, error) {
	logster.StartFuncLog()

	// Retrieve the existing store record from the repository
	toUpdate, err := repository.GetStore(uuid)
	if err != nil {
		logster.Error(err, "Error getting store")
		return nil, err
	}

	var countries []string
	var categories []string

	//Check if country exists
	for _, country := range dtoParam.Country {
		_, err := GetCountryByCode(country)
		if err == nil {
			countries = append(countries, country)
		}
	}

	//Check if country exists
	for _, category := range dtoParam.Category {
		_, err := repository.GetCategoryByCode(category)
		if err == nil {
			categories = append(categories, category)
		}
	}

	if len(categories) == 0 {
		categories = make([]string, 0)
	}
	if len(countries) == 0 {
		countries = make([]string, 0)
	}

	// Update fields in `toUpdate` based on non-nil fields in `dtoParam`
	if dtoParam.Name != nil {
		toUpdate.Name = *dtoParam.Name
	}
	if dtoParam.ShortDescription != nil {
		toUpdate.ShortDescription = dtoParam.ShortDescription
	}
	if dtoParam.Description != nil {
		toUpdate.Description = dtoParam.Description
	}
	if dtoParam.UrlSlug != nil {
		toUpdate.UrlSlug = dtoParam.UrlSlug
	}
	if dtoParam.InitialReward != nil {
		toUpdate.InitialReward = dtoParam.InitialReward
	}
	if dtoParam.AverageRewardActivationTime != nil {
		toUpdate.AverageRewardActivationTime = dtoParam.AverageRewardActivationTime
	}
	if dtoParam.State != nil {
		toUpdate.State = dtoParam.State
	}
	if dtoParam.Keywords != nil {
		toUpdate.Keywords = dtoParam.Keywords
	}
	if dtoParam.AffiliateLink != nil {
		toUpdate.AffiliateLink = dtoParam.AffiliateLink
	}
	if dtoParam.StoreUrl != nil {
		toUpdate.StoreUrl = dtoParam.StoreUrl
	}
	if dtoParam.TermsAndConditions != nil {
		toUpdate.TermsAndConditions = dtoParam.TermsAndConditions
	}
	if dtoParam.CashbackType != nil {
		toUpdate.CashbackType = dtoParam.CashbackType
	}
	if dtoParam.CashbackValue != nil {
		toUpdate.CashbackValue = dtoParam.CashbackValue
	}
	if dtoParam.PercentageCashout != nil {
		toUpdate.PercentageCashout = dtoParam.PercentageCashout
	}
	if dtoParam.MetaTitle != nil {
		toUpdate.MetaTitle = dtoParam.MetaTitle
	}
	if dtoParam.MetaKeywords != nil {
		toUpdate.MetaKeywords = dtoParam.MetaKeywords
	}
	if dtoParam.MetaDescription != nil {
		toUpdate.MetaDescription = dtoParam.MetaDescription
	}
	if dtoParam.PartnerIdentity != nil {
		toUpdate.PartnerIdentity = dtoParam.PartnerIdentity
	}

	if dtoParam.Logo != nil {
		toUpdate.Logo = dtoParam.Logo
	} else {
		if toUpdate.Logo != nil && *toUpdate.Logo != "" {
			toUpdate.Logo = utils.GetUuidFromUrl(*toUpdate.Logo)
		}

	}

	if toUpdate.Banner != nil && *toUpdate.Banner != "" {
		toUpdate.Banner = utils.GetUuidFromUrl(*toUpdate.Banner)
	}

	// Update foreign keys if provided

	if dtoParam.LanguageCODE != nil {
		toUpdate.LanguageCODE = dtoParam.LanguageCODE
	}
	if dtoParam.AffiliatePartnerCODE != nil {
		toUpdate.AffiliatePartnerCODE = dtoParam.AffiliatePartnerCODE
	}
	if dtoParam.OverrideFee != nil {
		toUpdate.OverrideFee = dtoParam.OverrideFee
	} else {
		toUpdate.OverrideFee = nil
	}

	if dtoParam.Position != nil {
		toUpdate.Position = dtoParam.Position
	}

	toUpdate.UpdatedBy = &uuidUser
	toUpdate.CountriesCodes = countries
	toUpdate.CategoriesCodes = categories

	// Save the updated store to the repository
	res, err := repository.UpdateStore(toUpdate)
	if err != nil {
		logster.Error(err, "Error updating store")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return &res, nil
}

func DeleteStore(uuid uuid.UUID, uuidUser string) error {
	logster.StartFuncLog()

	err := repository.DeleteStore(uuid, uuidUser)
	if err != nil {
		logster.Error(err, "Error deleting store")
		logster.EndFuncLog()
		return err
	}

	logster.EndFuncLog()
	return nil
}

func AddCountryToStore(storeUuid uuid.UUID, countries dto.StoreCountryDTO) error {
	logster.StartFuncLog()

	var model []models.StoreCountry
	for _, country := range countries.Countries {
		model = append(model, models.StoreCountry{
			StoreUUID:   storeUuid,
			CountryCode: country,
		})
	}

	err := repository.AddCountryToStore(model)
	if err != nil {
		logster.Error(err, "Error adding country to store")
		logster.EndFuncLog()
		return err
	}

	logster.EndFuncLog()
	return nil
}

func RemoveCountryFromStore(storeUuid uuid.UUID, countries dto.StoreCountryDTO) error {
	logster.StartFuncLog()

	var model []models.StoreCountry
	for _, country := range countries.Countries {
		model = append(model, models.StoreCountry{
			StoreUUID:   storeUuid,
			CountryCode: country,
		})
	}

	err := repository.RemoveCountryFromStore(model)
	if err != nil {
		logster.Error(err, "Error removing country from store")
		logster.EndFuncLog()
		return err
	}

	logster.EndFuncLog()
	return nil
}

func AddCategoryToStore(storeUuid uuid.UUID, countries dto.StoreCategoryDTO) error {

	var model []models.StoreCategory
	for _, cat := range countries.Categories {
		model = append(model, models.StoreCategory{
			StoreUUID:    storeUuid,
			CategoryCode: cat,
		})
	}

	err := repository.AddCategoryToStore(model)
	if err != nil {
		logster.Error(err, "Error adding category to store")
		logster.EndFuncLog()
		return err
	}
	return nil
}

func RemoveCategoryFromStore(storeUuid uuid.UUID, countries dto.StoreCategoryDTO) error {

	var model []models.StoreCategory
	for _, cat := range countries.Categories {
		model = append(model, models.StoreCategory{
			StoreUUID:    storeUuid,
			CategoryCode: cat,
		})
	}

	err := repository.RemoveCategoryFromStore(model)
	if err != nil {
		logster.Error(err, "Error removing category from store")
		logster.EndFuncLog()
		return err
	}
	return nil
}

func UploadStoreByExcel(file *multipart.FileHeader) (*dto.StoreUploadResponseDTO, error) {
	logster.StartFuncLog()

	// Open the file using the temporary filename
	reader, err := file.Open()
	if err != nil {
		logster.Error(err, "Failed to open uploaded file")
		logster.EndFuncLog()
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer func(reader multipart.File) {
		err := reader.Close()
		if err != nil {
			logster.Error(err, "Error closing the reader")
		}
	}(reader) // Close the reader after use

	stores, er := parseCSVFile(reader)
	if er != nil {
		logster.Error(err, "Failed parsing CSV file")
		logster.EndFuncLog()
		return nil, er
	}

	// Insert the stores into the database
	_, err = repository.BulkInsertStores(stores)

	if err != nil {
		logster.Error(err, "Error inserting stores into database")
		logster.EndFuncLog()
		return nil, err
	}

	//Create the response
	var uuidList []uuid.UUID
	for _, store := range stores {
		uuidList = append(uuidList, store.Uuid) // Append each UUID to the uuidList slice
	}

	response := dto.StoreUploadResponseDTO{
		Message: "Stores inserted successfully, number of Rows added: " + strconv.Itoa(len(stores)),
		Stores:  uuidList,
	}

	logster.EndFuncLog()
	return &response, nil
}

// parseCSVFile parses a CSV file and returns a slice of Stores
func parseCSVFile(reader io.Reader) ([]models.Store, error) {
	logster.StartFuncLog()
	decoder := charmap.ISO8859_1.NewDecoder()
	newReader := transform.NewReader(reader, decoder)

	f := csv.NewReader(newReader)

	records, err := f.ReadAll()
	if err != nil {
		logster.Error(err, "Failed to read CSV file")
		logster.EndFuncLog()
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	var isUUID bool
	stores := make([]models.Store, len(records)-1)
	for i, record := range records {
		var uuidStore string
		var legacyId *string = nil

		if record[0] == "uuid" {
			isUUID = true
		}

		// Skip the header row
		if i == 0 {
			continue
		}

		if isUUID {
			uuidStore = record[0]
		} else {
			legacyId = &record[0]
		}

		// Convert the reward value to float64
		cbValue, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			logster.Error(err, "Error converting reward value to float64")
			logster.EndFuncLog()
			return nil, utils.CustomErrorStruct{}.ErrorParsingValue(
				record[5],
				"string",
				"float64",
			)
		}

		// Convert the percentage cashout to float64
		record[7] = strings.Replace(record[7], ",", ".", 1)
		perCashout, err := strconv.ParseFloat(record[7], 64)
		if err != nil {
			logster.Error(err, "Error converting percentage cashout to float64")
			logster.EndFuncLog()
			return nil, utils.CustomErrorStruct{}.ErrorParsingValue(
				record[7],
				"string",
				"float64",
			)
		}

		// Split the country and category codes to arrays
		countryCodes := strings.Split(record[1], ",")
		categoryCodes := strings.Split(record[2], ",")

		var logoUuid string
		//Check if logo (record[10]) is empty, if so, use storeUrl (record[14]) to get logo. If both are empty, set null
		if record[10] == "" {
			if record[14] != "" {
				logo, errLogoDoman := images.GetLogoFromDomain(record[14])
				if errLogoDoman != nil {
					logster.Error(errLogoDoman, "Error getting logo from domain")
				}
				if logo != nil {
					logoUuid = logo.Id
				}
				logster.Debug(fmt.Sprintf("Logo Uuid: %s\n", logoUuid))
			}
		} else {
			storeLogo := record[10]

			if !strings.Contains(storeLogo, "http") {
				logoUuid = storeLogo
			} else {
				logo, errLogoUrl := images.GetLogoFromUrl(storeLogo, record[3]+"_logo")

				if errLogoUrl != nil {
					logster.Error(errLogoUrl, "Error getting logo from url")
				}
				if logo != nil {
					logoUuid = logo.Id
				}
			}

		}

		var bannerUuid string
		//Check if banner (record[11]) is empty, if not, get banner image from url
		if record[11] != "" {
			storeBanner := record[11]

			if !strings.Contains(storeBanner, "http") {
				bannerUuid = storeBanner

			} else {
				banner, errBannerUrl := images.GetBannerFromUrl(storeBanner, record[3]+"_banner")

				if errBannerUrl != nil {
					logster.Error(errBannerUrl, "Error getting banner from url")
				}

				if banner != nil {
					bannerUuid = banner.Id
				}
			}
		}

		var overrideFee *float64 = nil
		if record[24] != "" {
			// Convert the transaction fee to float64
			txFee, err := strconv.ParseFloat(record[24], 64)
			if err != nil {
				return nil, utils.CustomErrorStruct{}.ErrorParsingValue(
					record[24],
					"string",
					"float64",
				)
			}
			overrideFee = &txFee
		}

		var languageCode *string = nil
		if record[25] != "" {
			languageCode = &record[25]
		}

		state := strings.ToUpper(record[20])

		stores[i-1] = models.Store{
			LegacyId:                    legacyId,
			CountriesCodes:              countryCodes,  //COUNTRY CODES
			CategoriesCodes:             categoryCodes, //CATEGORY CODE
			Name:                        record[3],
			CashbackType:                &record[4],
			CashbackValue:               &cbValue,
			AverageRewardActivationTime: &record[6],
			PercentageCashout:           &perCashout,
			ShortDescription:            &record[8],
			Description:                 &record[9],
			Logo:                        &logoUuid,
			Banner:                      &bannerUuid,
			AffiliatePartnerCODE:        &record[12], //PARTNER CODE
			AffiliateLink:               &record[13],
			StoreUrl:                    &record[14],
			TermsAndConditions:          &record[15],
			Keywords:                    &record[16],
			MetaTitle:                   &record[17],
			MetaKeywords:                &record[18],
			MetaDescription:             &record[19],
			State:                       &state,
			UrlSlug:                     &record[22],
			PartnerIdentity:             &record[23],
			OverrideFee:                 overrideFee,
			//Commented out because it is not needed YET
			LanguageCODE: languageCode,
		}

		if isUUID {
			stores[i-1].Uuid = utils.ParseIDToUUID(uuidStore)
		}
	}

	logster.EndFuncLog()
	return stores, nil
}

func UpdateStoreLogo(uuid uuid.UUID, fileUuid *string, uuidUser string) (*models.Store, error) {
	logster.StartFuncLog()

	// Save the updated store to the repository
	res, err := repository.UpdateStoreLogo(uuid, *fileUuid, uuidUser)
	if err != nil {
		logster.Error(err, "Error updating store logo")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return res, nil
}

func DeleteStoreLogo(uuidStore uuid.UUID, uuidUser string) error {
	logster.StartFuncLog()

	toUpdate, err := repository.GetStore(uuidStore)
	if err != nil {
		logster.Error(err, "Error getting store")
		logster.EndFuncLog()
		return err
	}
	imageUuid := utils.GetUuidFromUrl(*toUpdate.Logo)

	_, err = repository.DeleteStoreLogo(uuidStore, uuidUser)
	if err != nil {
		logster.Error(err, "Error deleting store logo")
		logster.EndFuncLog()
		return err
	}

	err = images.DeleteImage(*imageUuid)

	if err != nil {
		logster.Error(err, "Error deleting image from ms-images")
		logster.EndFuncLog()
		return err
	}

	logster.EndFuncLog()
	return nil
}

func UpdateStoreBanner(uuid uuid.UUID, fileUuid *string, uuidUser string) (*models.Store, error) {
	logster.StartFuncLog()

	// Save the updated store to the repository
	res, err := repository.UpdateStoreBanner(uuid, *fileUuid, uuidUser)
	if err != nil {
		logster.Error(err, "Error updating store banner")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return res, nil
}

func DeleteStoreBanner(uuidStore uuid.UUID, uuidUser string) error {
	logster.StartFuncLog()

	toUpdate, err := repository.GetStore(uuidStore)
	if err != nil {
		logster.Error(err, "Error getting store")
		logster.EndFuncLog()
		return err
	}
	imageUuid := utils.GetUuidFromUrl(*toUpdate.Banner)

	_, err = repository.DeleteStoreBanner(uuidStore, uuidUser)
	if err != nil {
		logster.Error(err, "Error deleting store banner")
		logster.EndFuncLog()
		return err
	}

	err = images.DeleteImage(*imageUuid)

	if err != nil {
		logster.Error(err, "Error deleting image from ms-images")
		logster.EndFuncLog()
		return err
	}

	logster.EndFuncLog()
	return nil
}

func GenerateRefAndGetStoreLink(store *response_object.GetStoreRO, uuidUser string, keycloak *constants.Keycloak) (*string, error) {
	logster.StartFuncLog()

	var newRef string
	var encodedRef string
	now := time.Now()

	switch *store.AffiliatePartnerCODE {
	case "awin_pt", "awin_uk":
		{
			//Generate reference and store visit
			awinRefPrefix := dotenv.GetEnv("AWIN_REF_PREFIX")
			lastRef, err := repository.GetLatestRef()
			if err != nil {
				logster.Error(err, "Error getting latest ref")
				logster.EndFuncLog()
				return nil, err
			}
			lastRefSplit := strings.Split(*lastRef, "-")

			if len(lastRefSplit) > 1 {
				lastRefNumberInt, _ := strconv.Atoi(lastRefSplit[1])
				lastRefNumberInt = lastRefNumberInt + 1
				newRef = awinRefPrefix + "-" + strconv.Itoa(lastRefNumberInt)
			} else {
				newRef = awinRefPrefix + "-1"
			}
		}
	case "cj_pt", "cj_uk":
		{
			cjRefPrefix := dotenv.GetEnv("CJ_REF_PREFIX")
			if cjRefPrefix != "" {
				newRef = fmt.Sprintf("%s_%d", cjRefPrefix, now.UnixMilli())
			} else {
				newRef = fmt.Sprintf("%d", now.UnixMilli())
			}
		}
	}

	//Get partner and generate url
	partner, err := GetPartnerByCode(store.AffiliatePartnerCODE)

	if err != nil {
		logster.Error(err, "Error getting partner")
		logster.EndFuncLog()
		return nil, err
	}
	logster.Info(fmt.Sprintf("Partner found: %s", partner.Uuid.String()))

	if newRef != "" {
		storeVisit := models.StoreVisit{
			StoreUUID: &store.Uuid,
			User:      &uuidUser,
			Reference: utils.StringPointer(newRef),
			BaseEntity: models.BaseEntity{
				CreatedBy: uuidUser,
				UpdatedBy: utils.StringPointer(uuidUser),
				CreatedAt: now,
			},
		}
		_, err = repository.CreateStoreVisit(storeVisit)
		if err != nil {
			logster.Error(err, "Error creating store visit")
			logster.EndFuncLog()
			return nil, err
		}
		logster.Info(fmt.Sprintf("Created store visit %s", *storeVisit.Reference))

		encodedRef = utils.EncodeAwinStoreVisitRef(newRef, uuidUser)
	}

	switch *partner.Code {
	case "awin_pt", "awin_uk":
		{
			if store.AffiliateLink != "" {
				url := store.AffiliateLink
				url = url + *partner.SubIdentifier + encodedRef
				return &url, nil
			}

			if store.PartnerIdentity == nil || *store.PartnerIdentity == "" {
				return &store.StoreUrl, nil
			}
			urlTemplate := dotenv.GetEnv("AWIN_TRACKING_LINK_TEMPLATE")
			tagpeakPublisherId := dotenv.GetEnv("AWIN_TAGPEAK_PUBLISHER_ID")
			url := strings.Replace(urlTemplate, "{storeIdentity}", *store.PartnerIdentity, 1)
			url = strings.Replace(url, "{tagpeakPublisherId}", tagpeakPublisherId, 1)

			url = url + *partner.SubIdentifier + encodedRef

			logster.Info("REDIRECT URL: " + url)

			return &url, nil
		}
	case "cj_pt", "cj_uk":
		{
			sid := newRef

			if store.AffiliateLink != "" {
				url := store.AffiliateLink

				if !strings.Contains(url, "?url=") {
					storeUrl := store.StoreUrl
					if storeUrl == "" {
						storeUrl = "https://www." + store.UrlSlug + ".com"
					}
					url += "?url=" + storeUrl
				}

				if strings.Contains(url, "&sid=") {
					re := regexp.MustCompile(`&sid=[^&]*`)
					url = re.ReplaceAllString(url, "")
				}

				url += "&sid=" + sid

				return &url, nil
			}

			if store.PartnerIdentity == nil {
				logster.Error(nil, "Partner identity is empty")
				return nil, errors.New("partner identity is empty")
			}

			urlTemplate := dotenv.GetEnv("CJ_TRACKING_LINK_TEMPLATE")
			pid := dotenv.GetEnv("CJ_PID")
			url := strings.Replace(urlTemplate, "{pid}", pid, 1)

			url = strings.Replace(url, "{aid}", fmt.Sprintf("%s", *store.PartnerIdentity), 1)

			if store.StoreUrl != "" {
				url = strings.Replace(url, "{storeUrl}", store.StoreUrl, 1)
			} else {
				url = strings.Replace(url, "{storeUrl}", "https://www."+store.UrlSlug+".com", 1)
			}

			url = strings.Replace(url, "{sid}", sid, 1)

			logster.Info("REDIRECT URL: " + url)

			return &url, nil
		}
	default:
		{
			//If there is an affiliate link, return it
			if store.AffiliateLink != "" {
				return &store.AffiliateLink, nil
			}

			if store.StoreUrl != "" {
				return &store.StoreUrl, nil
			}
		}
	}

	logster.EndFuncLog()
	return nil, nil
}

func ExportStoresCSV(dtoPag pagination.PaginationParams, filters dto.StoreFiltersDTO) ([]response_object.ExportStoresRO, error) {
	logster.StartFuncLog()

	res, err := repository.GetAllStores(dtoPag, filters)
	if err != nil {
		logster.Error(err, "Error getting all stores")
		logster.EndFuncLog()
		return []response_object.ExportStoresRO{}, err
	}

	modelDto := lo.Map(res.Data.([]models.Store), func(item models.Store, _ int) response_object.ExportStoresRO {
		var countryCode []string
		if item.Country != nil {
			for _, country := range *item.Country {
				countryCode = append(countryCode, *country.Abbreviation)
			}
		}

		var categoryCode []string
		if item.Category != nil {
			for _, category := range *item.Category {
				categoryCode = append(categoryCode, *category.Code)
			}
		}

		var logo *string
		var banner *string

		if item.Logo != nil && *item.Logo != "" {
			logo = utils.GetUuidFromUrl(*item.Logo)
		}

		if item.Banner != nil && *item.Banner != "" {
			banner = utils.GetUuidFromUrl(*item.Banner)
		}

		return response_object.ExportStoresRO{
			Uuid:             item.Uuid,
			CountryCode:      &countryCode,
			Language:         item.Language.Code,
			CategoryCode:     &categoryCode,
			Name:             item.Name,
			UrlSlug:          item.UrlSlug,
			PartnerCode:      item.PartnerIdentity,
			ShortDescription: item.ShortDescription,
			Description:      item.Description,
			AveragePayout:    item.AverageRewardActivationTime,
			CashbackType:     item.CashbackType,
			CashbackValue:    item.CashbackValue,
			AverageCashout:   item.PercentageCashout,
			NetworkId:        item.AffiliatePartnerCODE,
			DeepLink:         item.AffiliateLink,
			StoreUrl:         item.StoreUrl,
			TermAndCondition: item.TermsAndConditions,
			Keywords:         item.Keywords,
			MetaTitle:        item.MetaTitle,
			MetaKeywords:     item.MetaKeywords,
			MetaDescription:  item.MetaDescription,
			Status:           item.State,
			AveragePayoutNum: item.AverageRewardActivationTime,
			PartnerIdentity:  &item.AffiliatePartner.Uuid,
			OverrideFee:      item.OverrideFee,
			LegacyId:         item.LegacyId,
			StoreIcon:        logo,
			StoreBanner:      banner,
		}
	})

	logster.EndFuncLog()
	return modelDto, nil
}

func GetStoresForApproval(pag pagination.PaginationParams, filters dto.StoreFiltersDTO, keycloak *constants.Keycloak) (*pagination.PaginationResult, error) {
	logster.StartFuncLog()

	res, err := repository.GetAllStores(pag, filters)

	if err != nil {
		logster.Error(err, "Error getting all stores")
		logster.EndFuncLog()
		return nil, err
	}

	var aux []response_object.ApprovalRequestRo

	for _, store := range res.Data.([]models.Store) {
		data := response_object.ApprovalRequestRo{
			Uuid:      store.Uuid,
			Status:    *store.State,
			Name:      store.Name,
			CreatedAt: store.CreatedAt,
		}

		if *(store.AffiliatePartnerCODE) == "shopify" {
			shopifyShop, err := GetShopifyShopByStoreUuid(store.Uuid)
			if err != nil {
				continue
			}
			user, err := GetUserById(shopifyShop.UserUuid.String(), keycloak)
			if err != nil {
				continue
			}
			data.User.Uuid = user.Uuid
			data.User.Name = user.FirstName + " " + user.LastName
			data.User.Email = user.Email
		}

		aux = append(aux, data)
	}

	res.Data = aux

	return res, nil
}

func IsPositionUnique(position int) bool {
	logster.StartFuncLog()
	isUnique := repository.IsPositionUnique(position)

	logster.StartFuncLogMsg(fmt.Sprintf("Position %d is unique: %t", position, isUnique))
	return isUnique
}
