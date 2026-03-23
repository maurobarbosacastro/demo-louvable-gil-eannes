package response_object

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type PublicStore struct {
	StoreUuid uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Logo      *string   `json:"logo"`
	Position  *int      `json:"position"`
}

type GetStoreRO struct {
	Uuid                        uuid.UUID `json:"uuid"`
	Name                        string    `json:"name"`
	Logo                        *string   `json:"logo,omitempty"`
	Banner                      *string   `json:"banner,omitempty"`
	ShortDescription            string    `json:"shortDescription,omitempty"`
	Description                 string    `json:"description,omitempty"`
	UrlSlug                     string    `json:"urlSlug,omitempty"`
	AverageRewardActivationTime string    `json:"averageRewardActivationTime,omitempty"`
	State                       string    `json:"state,omitempty"` // Assuming enum type
	Keywords                    string    `json:"keywords,omitempty"`
	AffiliateLink               string    `json:"affiliateLink,omitempty"`
	StoreUrl                    string    `json:"storeUrl,omitempty"`
	TermsAndConditions          string    `json:"termsAndConditions,omitempty"`
	CashbackType                string    `json:"cashbackType,omitempty"` // Assuming enum type
	CashbackValue               float64   `json:"cashbackValue,omitempty"`
	PercentageCashout           float64   `json:"percentageCashout,omitempty"`
	MetaTitle                   string    `json:"metaTitle,omitempty"`
	MetaKeywords                string    `json:"metaKeywords,omitempty"`
	MetaDescription             string    `json:"metaDescription,omitempty"`
	Country                     []string  `json:"country"`
	Category                    []string  `json:"category"`
	PartnerIdentity             *string   `json:"partnerIdentity,omitempty"`
	OverrideFee                 *float64  `json:"overrideFee,omitempty"`
	Position                    *int      `json:"position,omitempty"`
	//InitialReward               float64  `json:"initialReward,omitempty"`

	// Foreign keys as UUIDs

	LanguageCODE         *string `json:"languageCode,omitempty"`
	AffiliatePartnerCODE *string `json:"affiliatePartnerCode,omitempty"`
}

type GetStorePublicRO struct {
	Uuid                        uuid.UUID `json:"uuid"`
	Name                        string    `json:"name"`
	Logo                        *string   `json:"logo,omitempty"`
	Banner                      *string   `json:"banner,omitempty"`
	ShortDescription            string    `json:"shortDescription,omitempty"`
	Description                 string    `json:"description,omitempty"`
	AverageRewardActivationTime string    `json:"averageRewardActivationTime,omitempty"`
	Keywords                    string    `json:"keywords,omitempty"`
	StoreUrl                    string    `json:"storeUrl,omitempty"`
	TermsAndConditions          string    `json:"termsAndConditions,omitempty"`
	PercentageCashout           float64   `json:"percentageCashout,omitempty"`
	MetaTitle                   string    `json:"metaTitle,omitempty"`
	MetaKeywords                string    `json:"metaKeywords,omitempty"`
	MetaDescription             string    `json:"metaDescription,omitempty"`
}

type ExportStoresRO struct {
	Uuid             uuid.UUID  `json:"uuid,omitempty"`
	CountryCode      *[]string  `json:"countryCode,omitempty"`
	Language         *string    `json:"language,omitempty"`
	CategoryCode     *[]string  `json:"categoryCode,omitempty"`
	Name             string     `json:"name,omitempty"`
	UrlSlug          *string    `json:"urlSlug,omitempty"`
	PartnerCode      *string    `json:"partnerCode,omitempty"`
	ShortDescription *string    `json:"shortDescription,omitempty"`
	Description      *string    `json:"description,omitempty"`
	AveragePayout    *string    `json:"averagePayout,omitempty"`
	CashbackType     *string    `json:"cashbackType,omitempty"`
	CashbackValue    *float64   `json:"cashbackValue,omitempty"`
	AverageCashout   *float64   `json:"averageCashout,omitempty"`
	NetworkId        *string    `json:"networkId,omitempty"`
	DeepLink         *string    `json:"deepLink,omitempty"`
	StoreUrl         *string    `json:"storeUrl,omitempty"`
	TermAndCondition *string    `json:"termAndCondition,omitempty"`
	Keywords         *string    `json:"keywords,omitempty"`
	MetaTitle        *string    `json:"metaTitle,omitempty"`
	MetaKeywords     *string    `json:"metaKeywords,omitempty"`
	MetaDescription  *string    `json:"metaDescription,omitempty"`
	Status           *string    `json:"status,omitempty"`
	AveragePayoutNum *string    `json:"averagePayoutNum,omitempty"`
	PartnerIdentity  *uuid.UUID `json:"partnerIdentity,omitempty"`
	OverrideFee      *float64   `json:"overrideFee,omitempty"`
	LegacyId         *string    `json:"legacyId,omitempty"`
	StoreIcon        *string    `json:"storeIcon,omitempty"`
	StoreBanner      *string    `json:"storeBanner,omitempty"`
}

func MapExportStoresRO(store ExportStoresRO) []string {
	var countryCode, categoryCode, cashbackValue, averageCashout, overrideFee string
	if store.CountryCode != nil {
		countryCode = strings.Join(*store.CountryCode, ",")
	}
	if store.CategoryCode != nil {
		categoryCode = strings.Join(*store.CategoryCode, ",")
	}
	if store.CashbackValue != nil {
		cashbackValue = fmt.Sprintf("%.2f", *store.CashbackValue)
	}
	if store.AverageCashout != nil {
		averageCashout = fmt.Sprintf("%.2f", *store.AverageCashout)
	}
	if store.OverrideFee != nil {
		overrideFee = fmt.Sprintf("%.2f", *store.OverrideFee)
	}

	language := stringOrEmpty(store.Language)
	urlSlug := stringOrEmpty(store.UrlSlug)
	partnerCode := stringOrEmpty(store.PartnerCode)
	shortDesc := stringOrEmpty(store.ShortDescription)
	desc := stringOrEmpty(store.Description)
	avgRewardTime := stringOrEmpty(store.AveragePayout)
	storeIcon := stringOrEmpty(store.StoreIcon)
	storeBanner := stringOrEmpty(store.StoreBanner)
	networkId := stringOrEmpty(store.NetworkId)
	deepLink := stringOrEmpty(store.DeepLink)
	storeUrl := stringOrEmpty(store.StoreUrl)
	termAndCondition := stringOrEmpty(store.TermAndCondition)
	keywords := stringOrEmpty(store.Keywords)
	metaTitle := stringOrEmpty(store.MetaTitle)
	metaKeywords := stringOrEmpty(store.MetaKeywords)
	metaDescription := stringOrEmpty(store.MetaDescription)

	row := []string{
		store.Uuid.String(),
		countryCode,
		categoryCode,
		store.Name,
		*store.CashbackType,
		cashbackValue,
		avgRewardTime,
		averageCashout,
		shortDesc,
		desc,
		storeIcon,
		storeBanner,
		networkId,
		deepLink,
		storeUrl,
		termAndCondition,
		keywords,
		metaTitle,
		metaKeywords,
		metaDescription,
		*store.Status,
		*store.AveragePayoutNum,
		urlSlug,
		partnerCode,
		overrideFee,
		language,
	}

	return row
}

func stringOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

type ApprovalRequestRo struct {
	Uuid      uuid.UUID     `json:"uuid"`
	Status    string        `json:"status"`
	Name      string        `json:"name"`
	CreatedAt time.Time     `json:"createdAt"`
	User      SimpleUserDto `json:"user"`
}
