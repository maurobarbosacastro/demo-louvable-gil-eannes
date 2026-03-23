package dto

type CreateCategoryDTO struct {
	Name string  `json:"name" validate:"required"`
	Code *string `json:"code" validate:"required"`
}

type UpdateCategoryDTO struct {
	Name *string `json:"name"`
	Code *string `json:"code"`
}

type CategoryFiltersDTO struct {
	Name string `json:"name" query:"name"`
}
