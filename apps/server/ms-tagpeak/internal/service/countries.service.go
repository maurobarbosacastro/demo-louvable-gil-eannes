package service

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	pagination "ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
)

func GetCountry(uuid uuid.UUID) (*models.Country, error) {
	country, err := repository.GetCountry(uuid)
	if err != nil {
		return nil, err // Return an empty Country and the error
	}
	return &country, nil
}

func GetCountryByCode(code string) (*models.Country, error) {
	country, err := repository.GetCountryByCode(code)
	if err != nil {
		return nil, err // Return an empty Country and the error
	}
	return &country, nil
}

func GetCountries(pag pagination.PaginationParams, filters dto.CountryFiltersDTO) (*pagination.PaginationResult, error) {
	log.Info("*** START service.GetCountries ***")

	res, err := repository.GetCountriesWithPagination(pag, filters)
	if err != nil {
		log.Errorf("*** Occurred an error getting countries: %v ***", err)
		return nil, err
	}

	log.Info("*** END service.GetCountries ***")
	return res, nil
}

func CreateCountry(countryDto dto.CreateCountryDTO, uuidUser string) (*models.Country, error) {
	model := utils.CountryDtoToModel(&countryDto)
	model.CreatedBy = uuidUser

	country, err := repository.CreateCountry(model)
	if err != nil {
		return nil, err // Return an empty Country and the error
	}
	return &country, nil
}

func UpdateCountry(countryDto dto.UpdateCountryDTO, uuid uuid.UUID, uuidUser string) (*models.Country, error) {

	countryToUpdate, err := repository.GetCountry(uuid)
	if err != nil {
		return nil, err
	}

	if countryDto.Name != nil {
		countryToUpdate.Name = countryDto.Name
	}
	if countryDto.Abbreviation != nil {
		countryToUpdate.Abbreviation = countryDto.Abbreviation
	}
	if countryDto.Flag != nil {
		countryToUpdate.Flag = countryDto.Flag
	}
	if countryDto.Currency != nil {
		countryToUpdate.Currency = countryDto.Currency
	}
	if countryDto.Enabled != nil {
		countryToUpdate.Enabled = countryDto.Enabled
	}

	countryToUpdate.UpdatedBy = &uuidUser

	country, err := repository.UpdateCountry(countryToUpdate)
	if err != nil {
		return nil, err // Return an empty Country and the error
	}
	return &country, nil
}

func SoftDeleteCountry(uuid uuid.UUID, uuidUser string) error {
	err := repository.SoftDeleteCountry(uuid, uuidUser)
	if err != nil {
		return err // Return an empty Country and the error
	}
	return nil
}
