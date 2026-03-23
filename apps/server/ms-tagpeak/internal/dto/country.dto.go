package dto

import "github.com/google/uuid"

type CountryDTO struct {
	Uuid         *uuid.UUID `json:"uuid"`
	Abbreviation *string    `json:"abbreviation"`
	Currency     *string    `json:"currency"`
	Flag         *string    `json:"flag"`
	Name         *string    `json:"name"`
	Enabled      *bool      `json:"enabled"`
}

type CreateCountryDTO struct {
	Abbreviation string `json:"abbreviation" validate:"required"`
	Currency     string `json:"currency" validate:"required"`
	Flag         string `json:"flag"`
	Name         string `json:"name" validate:"required"`
	Enabled      bool   `json:"enabled"`
}

type UpdateCountryDTO struct {
	Abbreviation *string `json:"abbreviation"`
	Currency     *string `json:"currency"`
	Flag         *string `json:"flag"`
	Name         *string `json:"name"`
	Enabled      *bool   `json:"enabled"`
}

type CountryFiltersDTO struct {
	Currency *string `json:"currency" query:"currency"`
	Enabled  *bool   `json:"enabled" query:"enabled"`
	Name     *string `json:"name" query:"name"`
}
