package dto

import "github.com/google/uuid"

type StoreDTO struct {
	Uuid                        uuid.UUID `json:"uuid"`
	Name                        string    `json:"name"`
	Logo                        *string   `json:"logo,omitempty"`
	ShortDescription            *string   `json:"shortDescription,omitempty"`
	Description                 *string   `json:"description,omitempty"`
	UrlSlug                     *string   `json:"urlSlug,omitempty"`
	InitialReward               *float64  `json:"initialReward,omitempty"`
	AverageRewardActivationTime *float64  `json:"averageRewardActivationTime,omitempty"`
	State                       *string   `json:"state,omitempty"`
	Keywords                    *string   `json:"keywords,omitempty"`
	AffiliateLink               *string   `json:"affiliateLink,omitempty"`
	StoreUrl                    *string   `json:"storeUrl,omitempty"`
	TermsAndConditions          *string   `json:"termsAndConditions,omitempty"`
	CashbackType                *string   `json:"cashbackType,omitempty"`
	CashbackValue               *float64  `json:"cashbackValue,omitempty"`
	PercentageCashout           *float64  `json:"percentageCashout,omitempty"`
	MetaTitle                   *string   `json:"metaTitle,omitempty"`
	MetaKeywords                *string   `json:"metaKeywords,omitempty"`
	MetaDescription             *string   `json:"metaDescription,omitempty"`

	// Foreign keys with embedded related values
	Category         *uuid.UUID `json:"category,omitempty"`
	Country          *uuid.UUID `json:"country,omitempty"`
	Language         *uuid.UUID `json:"language,omitempty"`
	AffiliatePartner *uuid.UUID `json:"affiliatePartner,omitempty"`
}

type CreateStoreDTO struct {
	Name                        string   `json:"name"`
	Logo                        string   `json:"logo,omitempty"`
	ShortDescription            string   `json:"shortDescription,omitempty"`
	Description                 string   `json:"description,omitempty"`
	UrlSlug                     string   `json:"urlSlug,omitempty"`
	AverageRewardActivationTime string   `json:"averageRewardActivationTime,omitempty"`
	State                       string   `json:"state,omitempty"` // Assuming enum type
	Keywords                    string   `json:"keywords,omitempty"`
	AffiliateLink               string   `json:"affiliateLink,omitempty"`
	StoreUrl                    string   `json:"storeUrl,omitempty"`
	TermsAndConditions          string   `json:"termsAndConditions,omitempty"`
	CashbackType                string   `json:"cashbackType,omitempty"` // Assuming enum type
	CashbackValue               float64  `json:"cashbackValue,omitempty"`
	PercentageCashout           float64  `json:"percentageCashout,omitempty"`
	MetaTitle                   string   `json:"metaTitle,omitempty"`
	MetaKeywords                string   `json:"metaKeywords,omitempty"`
	MetaDescription             string   `json:"metaDescription,omitempty"`
	Country                     []string `json:"country"`
	Category                    []string `json:"category"`
	OverrideFee                 *float64 `json:"overrideFee,omitempty"`
	//InitialReward               float64  `json:"initialReward,omitempty"`
	PartnerIdentity *string `json:"partnerIdentity,omitempty"`
	Position        *int    `json:"position,omitempty"`

	// Foreign keys as UUIDs

	LanguageCODE         *string `json:"languageCode,omitempty"`
	AffiliatePartnerCODE *string `json:"affiliatePartnerCode,omitempty"`
}

type UpdateStoreDTO struct {
	Name                        *string  `json:"name"`
	Logo                        *string  `json:"logo,omitempty"`
	ShortDescription            *string  `json:"shortDescription,omitempty"`
	Description                 *string  `json:"description,omitempty"`
	UrlSlug                     *string  `json:"urlSlug,omitempty"`
	InitialReward               *float64 `json:"initialReward,omitempty"`
	AverageRewardActivationTime *string  `json:"averageRewardActivationTime,omitempty"`
	State                       *string  `json:"state,omitempty"` // Assuming enum type
	Keywords                    *string  `json:"keywords,omitempty"`
	AffiliateLink               *string  `json:"affiliateLink,omitempty"`
	StoreUrl                    *string  `json:"storeUrl,omitempty"`
	TermsAndConditions          *string  `json:"termsAndConditions,omitempty"`
	CashbackType                *string  `json:"cashbackType,omitempty"` // Assuming enum type
	CashbackValue               *float64 `json:"cashbackValue,omitempty"`
	PercentageCashout           *float64 `json:"percentageCashout,omitempty"`
	MetaTitle                   *string  `json:"metaTitle,omitempty"`
	MetaKeywords                *string  `json:"metaKeywords,omitempty"`
	MetaDescription             *string  `json:"metaDescription,omitempty"`
	Country                     []string `json:"country"`
	Category                    []string `json:"category"`
	PartnerIdentity             *string  `json:"partnerIdentity,omitempty"`
	OverrideFee                 *float64 `json:"overrideFee"`
	Position                    *int     `json:"position"`

	// Foreign keys as UUIDs
	LanguageCODE         *string `json:"languageCode,omitempty"`
	AffiliatePartnerCODE *string `json:"affiliatePartnerCode,omitempty"`
}

type StoreFiltersDTO struct {
	State        string `json:"state" query:"state"`
	CountryCode  string `json:"countryCode" query:"countryCode"`
	CategoryCode string `json:"categoryCode" query:"categoryCode"`
	Name         string `json:"name" query:"name"`
	Sort         string `json:"sort" query:"sort"`
}

type AdminStoreDTO struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        *string   `json:"name"`
	Description *string   `json:"description,omitempty"`
	Partner     *string   `json:"partner,omitempty"`
	Country     []string  `json:"country,omitempty"`
}

type ForStoreVisitDTO struct {
	Uuid uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
	Logo *string   `json:"logo"`
}
