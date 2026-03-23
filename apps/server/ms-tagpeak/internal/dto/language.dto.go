package dto

import "github.com/google/uuid"

type LanguageDTO struct {
	Uuid *uuid.UUID `json:"uuid"`
	Name *string    `json:"name"`
	Code *string    `json:"code"`
}

type CreateLanguageDTO struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type UpdateLanguageDTO struct {
	Name *string `json:"name"`
	Code *string `json:"code"`
}
