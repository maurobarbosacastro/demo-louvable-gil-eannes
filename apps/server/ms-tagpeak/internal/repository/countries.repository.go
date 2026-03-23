package service

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"strings"
)

func GetCountry(uuid uuid.UUID) (models.Country, error) {
	dbInstance := db.GetDB()

	var country models.Country
	// Execute the query and capture the error
	err := dbInstance.Scopes(utils.ActiveScope()).Where("uuid = ?", uuid).First(&country).Error
	if err != nil {
		return models.Country{}, err // Return an empty Country and the error if not found or other error occurs
	}

	return country, nil // Return the found country and nil error if successful
}

func GetCountryByCode(code string) (models.Country, error) {
	dbInstance := db.GetDB()

	var country models.Country
	// Execute the query and capture the error
	err := dbInstance.Scopes(utils.ActiveScope()).Where("abbreviation = ?", code).First(&country).Error
	if err != nil {
		return models.Country{}, err // Return an empty Country and the error if not found or other error occurs
	}

	return country, nil // Return the found country and nil error if successful
}

func GetCountriesWithPagination(pagDTO pagination.PaginationParams, filters dto.CountryFiltersDTO) (*pagination.PaginationResult, error) {
	log.Info("*** START repository.GetCountriesWithPagination ***")
	dbInstance := db.GetDB()
	var countries []models.Country
	var res pagination.PaginationResult

	res.Limit = pagDTO.Limit
	res.Page = pagDTO.Page
	res.Sort = pagDTO.Sort

	if filters.Name != nil {
		dbInstance = dbInstance.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(*filters.Name)+"%")
		filters.Name = nil
	}

	if filters.Currency != nil {
		dbInstance = dbInstance.Where("currency = ?", filters.Currency)
	}

	if filters.Enabled != nil {
		dbInstance = dbInstance.Where("enabled = ?", filters.Enabled)
	}

	err := dbInstance.Scopes(pagination.Paginate(&countries, &res, filters, dbInstance)).Find(&countries).Error
	res.Data = countries
	log.Infof("*** Found %v countries ***", len(countries))

	if err != nil {
		log.Errorf("*** Occurred an error getting countries: %v ***", err)
		return nil, err
	}

	log.Info("*** END repository.GetCountriesWithPagination ***")
	return &res, nil
}

func CreateCountry(country models.Country) (models.Country, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&country).Error
	if err != nil {
		return models.Country{}, err // Return an empty Country and the error if not found or other error occurs
	}

	return country, nil
}

func UpdateCountry(country models.Country) (models.Country, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&country).Error
	if err != nil {
		return models.Country{}, err // Return an empty Country and the error if not found or other error occurs
	}

	return country, nil
}

func SoftDeleteCountry(uuid uuid.UUID, user string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Country{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.Country{}, "uuid = ?", uuid).Error

	if err != nil {
		return err
	}

	return nil

}
