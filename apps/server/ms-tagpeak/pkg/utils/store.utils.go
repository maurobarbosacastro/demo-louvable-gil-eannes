package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/logster"
	"strings"
)

func CreateStoreDtoToModel(c *dto.CreateStoreDTO) models.Store {
	return models.Store{
		Name:                        c.Name,
		Logo:                        &c.Logo,
		ShortDescription:            &c.ShortDescription,
		Description:                 &c.Description,
		UrlSlug:                     &c.UrlSlug,
		AverageRewardActivationTime: &c.AverageRewardActivationTime,
		State:                       &c.State,
		Keywords:                    &c.Keywords,
		AffiliateLink:               &c.AffiliateLink,
		StoreUrl:                    &c.StoreUrl,
		TermsAndConditions:          &c.TermsAndConditions,
		CashbackType:                &c.CashbackType,
		CashbackValue:               &c.CashbackValue,
		PercentageCashout:           &c.PercentageCashout,
		MetaTitle:                   &c.MetaTitle,
		MetaKeywords:                &c.MetaKeywords,
		MetaDescription:             &c.MetaDescription,
		LanguageCODE:                c.LanguageCODE,
		AffiliatePartnerCODE:        c.AffiliatePartnerCODE,
		PartnerIdentity:             c.PartnerIdentity,
		OverrideFee:                 c.OverrideFee,
		//InitialReward:               &c.InitialReward,

	}
}

func ModelStoreToSimpleDTO(c *models.Store) response_object.GetStoreRO {

	var countries []string
	var categories []string

	if c.Country != nil {
		for _, country := range *c.Country {
			countries = append(countries, *country.Abbreviation)
		}
	}
	if c.Category != nil {
		for _, category := range *c.Category {
			categories = append(categories, *category.Code)
		}
	}

	return response_object.GetStoreRO{
		Uuid:                        c.Uuid,
		Name:                        c.Name,
		Logo:                        c.Logo,
		Banner:                      c.Banner,
		ShortDescription:            *c.ShortDescription,
		Description:                 *c.Description,
		UrlSlug:                     *c.UrlSlug,
		AverageRewardActivationTime: *c.AverageRewardActivationTime,
		State:                       *c.State,
		Keywords:                    *c.Keywords,
		AffiliateLink:               *c.AffiliateLink,
		StoreUrl:                    *c.StoreUrl,
		TermsAndConditions:          *c.TermsAndConditions,
		CashbackType:                *c.CashbackType,
		CashbackValue:               *c.CashbackValue,
		PercentageCashout:           *c.PercentageCashout,
		MetaTitle:                   *c.MetaTitle,
		MetaKeywords:                *c.MetaKeywords,
		MetaDescription:             *c.MetaDescription,
		LanguageCODE:                c.Language.Code,
		AffiliatePartnerCODE:        c.AffiliatePartner.Code,
		Country:                     countries,
		Category:                    categories,
		PartnerIdentity:             c.PartnerIdentity,
		OverrideFee:                 c.OverrideFee,
		Position:                    c.Position,
	}
}

func ForStoreVisitsMap(storeVisit models.StoreVisit) *dto.ForStoreVisitDTO {
	return &dto.ForStoreVisitDTO{
		Uuid: storeVisit.Store.Uuid,
		Name: storeVisit.Store.Name,
		Logo: storeVisit.Store.Logo,
	}
}

func ModelStoreToStoreAdminDTO(c *models.Store) dto.AdminStoreDTO {

	var countries []string

	for _, country := range *c.Country {
		countries = append(countries, *country.Abbreviation)
	}

	return dto.AdminStoreDTO{
		Uuid:        c.Uuid,
		Name:        &c.Name,
		Description: c.Description,
		Partner:     c.AffiliatePartner.Name,
		Country:     countries,
	}
}

func GetUuidFromUrl(url string) *string {
	logster.StartFuncLog()
	parts := strings.Split(url, "/")
	uuidString := parts[len(parts)-2]

	logster.EndFuncLog()
	return &uuidString
}
