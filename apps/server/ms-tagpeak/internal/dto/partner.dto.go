package dto

type CreatePartnerDTO struct {
	Name               string  `json:"name" validate:"required"`
	Code               *string `json:"code" validate:"required"`
	ECommercePlatform  string  `json:"eCommercePlatform" validate:"required"`
	CommissionRate     float64 `json:"commissionRate" validate:"required"`
	ValidationPeriod   int     `json:"validationPeriod" validate:"required"`
	DeepLink           string  `json:"deepLink" validate:"required"`
	DeepLinkIdentifier string  `json:"deepLinkIdentifier" validate:"required"`
	SubIdentifier      string  `json:"subIdentifier" validate:"required"`
	PercentageTagpeak  float64 `json:"percentageTagpeak" validate:"required"`
	PercentageInvested float64 `json:"percentageInvested" validate:"required"`
}

type UpdatePartnerDTO struct {
	Name               *string  `json:"name"`
	Code               *string  `json:"code"`
	ECommercePlatform  *string  `json:"eCommercePlatform"`
	CommissionRate     *float64 `json:"commissionRate"`
	ValidationPeriod   *int     `json:"validationPeriod"`
	DeepLink           *string  `json:"deepLink"`
	DeepLinkIdentifier *string  `json:"deepLinkIdentifier"`
	SubIdentifier      *string  `json:"subIdentifier"`
	PercentageTagpeak  *float64 `json:"percentageTagpeak"`
	PercentageInvested *float64 `json:"percentageInvested"`
}

type PartnerFiltersDTO struct {
	Name string `json:"name" query:"name"`
}
