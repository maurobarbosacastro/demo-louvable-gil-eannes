package dto

type CreateConfigurationDTO struct {
	Name     string `json:"name" validate:"required"`
	Code     string `json:"code" validate:"required"`
	Value    string `json:"value" validate:"required"`
	DataType string `json:"dataType" validate:"required"`
}

type UpdateConfigurationDTO struct {
	Name     *string `json:"name"`
	Value    *string `json:"value"`
	DataType *string `json:"dataType"`
}
