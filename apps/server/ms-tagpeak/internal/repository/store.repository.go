package service

import (
	"errors"
	"fmt"
	"math"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	lo "github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetStore(uuid uuid.UUID) (models.Store, error) {
	dbInstance := db.GetDB()

	var model models.Store
	// Execute the query and capture the error
	err := dbInstance.Preload("Country").Preload("Language").Preload("Category").Preload("AffiliatePartner").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.Store{}, err
	}

	return model, nil
}

func GetAffiliatePartnerByStore(uuid uuid.UUID) (models.Partner, error) {

	dbInstance := db.GetDB()
	var model models.Store
	// Execute the query and capture the error
	err := dbInstance.Preload("AffiliatePartner").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.Partner{}, err
	}

	return model.AffiliatePartner, nil
}

func GetCategoryByStore(uuid uuid.UUID) ([]models.Category, error) {

	dbInstance := db.GetDB()
	var model models.Store
	// Execute the query and capture the error
	err := dbInstance.Preload("Category").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return nil, err
	}

	return *model.Category, nil
}

func GetCountryByStore(uuid uuid.UUID) ([]models.Country, error) {

	dbInstance := db.GetDB()
	var model models.Store
	// Execute the query and capture the error
	err := dbInstance.Preload("Country").Where("uuid = ?", uuid).First(&model).Error

	if err != nil {
		return []models.Country{}, err
	}

	return *model.Country, nil
}

func GetLanguageByStore(uuid uuid.UUID) (models.Language, error) {

	dbInstance := db.GetDB()
	var model models.Store
	// Execute the query and capture the error\
	err := dbInstance.Preload("Language").Where("uuid = ?", uuid).First(&model).Error

	if err != nil {
		return models.Language{}, err
	}

	return model.Language, nil
}

func GetAllStores(pagDTO pagination.PaginationParams, filters dto.StoreFiltersDTO) (*pagination.PaginationResult, error) {

	dbInstance := db.GetDB()
	var model []models.Store
	var res pagination.PaginationResult

	res.Limit = pagDTO.Limit
	res.Page = pagDTO.Page
	res.Sort = pagDTO.Sort

	dbWithFilters := dbInstance

	if filters.CategoryCode != "" {
		dbWithFilters = dbWithFilters.Joins("inner join store_category on store_category.category_code = ? and store.uuid = store_category.store_uuid", filters.CategoryCode)
	}
	if filters.CountryCode != "" {
		dbWithFilters = dbWithFilters.Joins("inner join store_country on store_country.country_code = ? and store.uuid = store_country.store_uuid", filters.CountryCode)
	}
	if filters.Name != "" {
		dbWithFilters = dbWithFilters.Where("LOWER(store.name) LIKE LOWER(?)", "%"+filters.Name+"%")
	}
	if filters.State != "" {
		dbWithFilters = dbWithFilters.Where("store.state = ?", filters.State)
	}

	if pagDTO.Sort != "" {
		dbWithFilters = dbWithFilters.Order(pagDTO.Sort)
	}

	err := dbWithFilters.Model(&models.Store{}).Select("*").Scopes(pagination.Paginate(&model, &res, nil, dbWithFilters)).
		Preload("Country").Preload("Language").Preload("Category").Preload("AffiliatePartner").
		Find(&model).Error

	res.Data = model
	if err != nil {
		return &pagination.PaginationResult{}, err
	}

	return &res, nil
}

func GetAllStoresWithPagination(pagDTO pagination.PaginationParams, filters dto.StoreFiltersDTO) (*pagination.PaginationResult, error) {
	log.Info("*** START repository.GetAllStoresWithPagination ***")
	dbInstance := db.GetDB()
	var model []response_object.PublicStore
	var res pagination.PaginationResult
	totalSize := int64(0)

	// Populate basic pagination metadata
	res.Limit = pagDTO.Limit
	res.Page = pagDTO.Page
	res.Sort = pagDTO.Sort

	query := `SELECT public_stores.* FROM public_stores`
	queryCount := `SELECT COUNT(*) FROM public_stores`
	instructions := []string{}
	// Country code is a required filter
	// We are using the "ANY" operator to match any of the values in the array of countries in the store view
	instructions = append(instructions, fmt.Sprintf(` '%s' = ANY(countries_code) `, filters.CountryCode))

	// We are using the "ANY" operator to match any of the values in the array of categories in the store view
	if filters.CategoryCode != "" {
		instructions = append(instructions, fmt.Sprintf(` '%s' = ANY(categories_code) `, filters.CategoryCode))
	}

	if filters.Name != "" {
		instructions = append(instructions, fmt.Sprintf(" LOWER(name) LIKE LOWER('%s')", "%"+filters.Name+"%"))
	}

	if len(instructions) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(instructions, " AND "))
		queryCount += fmt.Sprintf(" WHERE %s", strings.Join(instructions, " AND "))
	}

	if strings.Contains(pagDTO.Sort, "most-popular") {
		query += ` ORDER BY total_visits DESC, total_cashbacks DESC`
	}

	if pagDTO.Sort != "" && !strings.Contains(pagDTO.Sort, "most-popular") {
		query += fmt.Sprintf(" ORDER BY %s", pagDTO.Sort)
	}

	query += ` limit ` + strconv.Itoa(pagDTO.Limit)
	query += ` offset ` + strconv.Itoa(pagDTO.Page*pagDTO.Limit)

	err := dbInstance.Raw(query).Scan(&model).Error
	errCount := dbInstance.Raw(queryCount).Scan(&totalSize).Error

	if err != nil {
		log.Errorf("Occurred an error getting the stores %v", err)
		return nil, err
	}
	if errCount != nil {
		log.Errorf("Occurred an error getting the stores %v", err)
		return nil, err
	}
	log.Infof("*** Found %v stores ***", len(model))

	for i := range model {
		model[i].Logo = InjectUrlImage(model[i])
	}

	// Position sorting logic
	if strings.Contains(pagDTO.Sort, "most-popular") {
		// Separate stores with and without positions using lo.Filter
		storesWithPosition := lo.Filter(model, func(store response_object.PublicStore, _ int) bool {
			return store.Position != nil
		})
		storesWithoutPosition := lo.Filter(model, func(store response_object.PublicStore, _ int) bool {
			return store.Position == nil
		})

		// Sort stores with positions by their position value
		sort.Slice(storesWithPosition, func(i, j int) bool {
			return *storesWithPosition[i].Position < *storesWithPosition[j].Position
		})

		// Create a map of positioned stores by their position number for quick lookup
		// Changed to map to slice to support multiple stores at same position
		positionedStoreMap := make(map[int][]response_object.PublicStore)
		for _, store := range storesWithPosition {
			pos := *store.Position
			positionedStoreMap[pos] = append(positionedStoreMap[pos], store)
		}

		// Insert stores with positions at their designated positions
		result := make([]response_object.PublicStore, 0, len(model))
		withoutPosIdx := 0

		for i := 0; i < len(model); i++ {
			// Check if any store(s) should be at this position (positions are 1-indexed, so compare with i+1)
			if storesAtPos, exists := positionedStoreMap[i+1]; exists {
				// Add all stores at this position
				result = append(result, storesAtPos...)
				delete(positionedStoreMap, i+1) // Mark as used by removing from map
			} else if withoutPosIdx < len(storesWithoutPosition) {
				// If no positioned store goes here, add next store without position
				result = append(result, storesWithoutPosition[withoutPosIdx])
				withoutPosIdx++
			}
		}

		// Append remaining stores without positions
		result = append(result, storesWithoutPosition[withoutPosIdx:]...)

		// Append any positioned stores that weren't placed (position beyond result size)
		// Flatten remaining map values and append
		for _, stores := range positionedStoreMap {
			result = append(result, stores...)
		}

		model = result
	}

	res.Data = model
	res.TotalRows = totalSize
	res.TotalPages = int(math.Ceil(float64(totalSize) / float64(pagDTO.Limit)))

	log.Info("*** END repository.GetAllStoresWithPagination ***")
	return &res, nil
}

func CreateStore(model models.Store) (models.Store, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Transaction(func(tx *gorm.DB) error {
		// Save the store
		if err := tx.Create(&model).Error; err != nil {
			return err
		}

		var storeCountries []models.StoreCountry
		var storeCategories []models.StoreCategory

		for _, country := range model.CountriesCodes {
			storeCountries = append(storeCountries, models.StoreCountry{
				StoreUUID:   model.Uuid,
				CountryCode: country,
			})
		}

		for _, category := range model.CategoriesCodes {
			storeCategories = append(storeCategories, models.StoreCategory{
				StoreUUID:    model.Uuid,
				CategoryCode: category,
			})
		}

		if len(model.CountriesCodes) > 0 {
			if err := tx.Create(&storeCountries).Error; err != nil {
				return err
			}
		}

		if len(model.CategoriesCodes) > 0 {

			if err := tx.Create(&storeCategories).Error; err != nil {
				return err
			}

		}

		return nil
	})
	if err != nil {
		return models.Store{}, err
	}
	return model, nil
}

func UpdateStoreLogo(uuid uuid.UUID, fileUuid string, updatedBy string) (*models.Store, error) {
	dbInstance := db.GetDB()

	var model models.Store
	err := dbInstance.Model(&model).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"logo":       fileUuid,
			"updated_by": updatedBy,
		}).Error

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func DeleteStoreLogo(uuid uuid.UUID, updatedBy string) (*models.Store, error) {
	dbInstance := db.GetDB()

	var model models.Store
	err := dbInstance.Model(&model).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"logo":       nil,
			"updated_by": updatedBy,
		}).Error

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func UpdateStoreBanner(uuid uuid.UUID, fileUuid string, updatedBy string) (*models.Store, error) {
	dbInstance := db.GetDB()

	var model models.Store
	err := dbInstance.Model(&model).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"banner":     fileUuid,
			"updated_by": updatedBy,
		}).Error

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func DeleteStoreBanner(uuid uuid.UUID, updatedBy string) (*models.Store, error) {
	dbInstance := db.GetDB()

	var model models.Store
	err := dbInstance.Model(&model).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"banner":     nil,
			"updated_by": updatedBy,
		}).Error

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func UpdateStore(model models.Store) (models.Store, error) {
	dbInstance := db.GetDB()

	// Transaction to rollback if some error happens on update
	err := dbInstance.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&model).Error; err != nil {
			return err
		}

		var storeCountries []models.StoreCountry
		var storeCategories []models.StoreCategory

		for _, country := range model.CountriesCodes {
			storeCountries = append(storeCountries, models.StoreCountry{
				StoreUUID:   model.Uuid,
				CountryCode: country,
			})
		}

		for _, category := range model.CategoriesCodes {
			storeCategories = append(storeCategories, models.StoreCategory{
				StoreUUID:    model.Uuid,
				CategoryCode: category,
			})
		}

		if err := dbInstance.Where("store_uuid = ?", model.Uuid).Delete(&models.StoreCountry{}).Error; err != nil {
			return err
		}

		if len(model.CountriesCodes) > 0 {
			if err := dbInstance.Create(&storeCountries).Error; err != nil {
				return err
			}
		}

		if err := dbInstance.Where("store_uuid = ?", model.Uuid).Delete(&models.StoreCategory{}).Error; err != nil {
			return err
		}

		if len(model.CategoriesCodes) > 0 {
			if err := dbInstance.Create(&storeCategories).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return models.Store{}, err
	}
	return model, nil
}

func DeleteStore(uuid uuid.UUID, user string) error {
	dbInstance := db.GetDB()
	err := dbInstance.Model(&models.Store{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.Store{}, "uuid = ?", uuid).Error

	if err != nil {
		return err
	}

	return nil
}

func AddCountryToStore(model []models.StoreCountry) error {
	dbInstance := db.GetDB()

	if err := dbInstance.Create(&model).Error; err != nil {
		return err
	}
	return nil
}

func RemoveCountryFromStore(model []models.StoreCountry) error {
	dbInstance := db.GetDB()

	if err := dbInstance.Delete(&model).Error; err != nil {
		return err
	}
	return nil
}

func AddCategoryToStore(model []models.StoreCategory) error {
	dbInstance := db.GetDB()

	if err := dbInstance.Create(&model).Error; err != nil {
		return err
	}
	return nil
}

func RemoveCategoryFromStore(model []models.StoreCategory) error {
	dbInstance := db.GetDB()

	if err := dbInstance.Delete(&model).Error; err != nil {
		return err
	}
	return nil
}

func BulkInsertStores(stores []models.Store) ([]models.Store, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Transaction(func(tx *gorm.DB) error {
		// If store already exists, update it, otherwise create it
		// In this case, we use DoUpdates instead of UpdateAll to nominate the columns to update
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "uuid"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "short_description", "description", "url_slug", "initial_reward", "average_reward_activation_time", "state", "keywords", "affiliate_link", "store_url", "terms_and_conditions", "cashback_type", "cashback_value", "percentage_cashout", "meta_title", "meta_keywords", "meta_description", "partner_identity", "override_fee", "language_code", "affiliate_partner_code", "logo", "banner"}),
		}).Create(&stores).Error; err != nil {
			return err
		}

		var storeCountries []models.StoreCountry
		var storeCategories []models.StoreCategory

		// Loop through each store and prepare StoreCountry entries
		for _, model := range stores {
			for _, country := range model.CountriesCodes {
				storeCountries = append(storeCountries, models.StoreCountry{
					StoreUUID:   model.Uuid,
					CountryCode: country,
				})
			}
		}

		// Loop through each store and prepare StoreCategory entries
		for _, model := range stores {
			for _, category := range model.CategoriesCodes {
				storeCategories = append(storeCategories, models.StoreCategory{
					StoreUUID:    model.Uuid,
					CategoryCode: category,
				})
			}
		}

		// Save all StoreCountry entries in bulk
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&storeCountries).Error; err != nil {
			//Error handling for foreign key violations and duplicate key errors
			if err != nil && (errors.Is(err, gorm.ErrForeignKeyViolated) ||
				strings.Contains(err.Error(), "duplicate key value violates")) || errors.Is(err, gorm.ErrDuplicatedKey) {

				//Loop through each store country and check where the error occurred
				//In the future, this could be improved by using a more efficient data structure
				for _, storeCountry := range storeCountries {
					if strings.Contains(err.Error(), storeCountry.CountryCode) {

						// Return a custom error with relevant information
						return utils.CustomErrorStruct{}.ErrorUploadCSVDuplicatedFK(
							"StoreCountry",
							storeCountry.CountryCode,
							findStoreByUUID(stores, storeCountry.StoreUUID).Name,
						)
					}
				}
			}
			return err
		}

		// Save all StoreCategory entries in bulk
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&storeCategories).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return stores, nil
}

// findStoreByUUID returns the first store in the slice with the given UUID for error handling
func findStoreByUUID(stores []models.Store, uuid uuid.UUID) *models.Store {
	for _, store := range stores {
		fmt.Println(store.Uuid)
		if store.Uuid == uuid {
			return &store
		}
	}
	return nil // Return nil if no store with the given UUID is found
}

func InjectUrlImage(model response_object.PublicStore) *string {
	url := dotenv.GetEnv("MS_IMAGES_SERVER_PUBLIC_URL")
	if model.Logo != nil && *model.Logo != "" {
		return utils.StringPointer(fmt.Sprintf(url+"%s/logo.webp", *model.Logo))
	}
	return nil
}

func IsPositionUnique(position int) bool {
	dbInstance := db.GetDB()
	var store models.Store

	err := dbInstance.Where("position = ?", position).First(&store).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	if store.Uuid != uuid.Nil {
		return false
	}

	return true
}
