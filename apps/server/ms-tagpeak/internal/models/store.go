package models

import (
	"fmt"
	"ms-tagpeak/pkg/dotenv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	CashbackTypePercentage = "PERCENTAGE"
	CashbackTypeFixed      = "FIXED"
)

type Store struct {
	Uuid                        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name                        string    `gorm:"type:text;not null" json:"name"`
	Logo                        *string   `gorm:"type:text" json:"logo"`
	Banner                      *string   `gorm:"type:text" json:"banner"`
	ShortDescription            *string   `gorm:"type:text" json:"shortDescription"`
	Description                 *string   `gorm:"type:text" json:"description"`
	UrlSlug                     *string   `gorm:"type:text" json:"urlSlug"`
	InitialReward               *float64  `gorm:"type:float" json:"initialReward"`
	AverageRewardActivationTime *string   `gorm:"type:text" json:"averageRewardActivationTime"`
	State                       *string   `gorm:"type:varchar(50);check:state IN ('ACTIVE', 'INACTIVE', 'PENDING', 'BLOCKED')" json:"state"` // Assuming enum type will be a string
	Keywords                    *string   `gorm:"type:text" json:"keywords"`
	AffiliateLink               *string   `gorm:"type:text" json:"affiliateLink"`
	StoreUrl                    *string   `gorm:"type:text" json:"storeUrl"`
	TermsAndConditions          *string   `gorm:"type:text" json:"termsAndConditions"`
	CashbackType                *string   `gorm:"type:varchar(50)" json:"cashbackType"` // Assuming enum type will be a string
	CashbackValue               *float64  `gorm:"type:float" json:"cashbackValue"`
	PercentageCashout           *float64  `gorm:"type:float" json:"percentageCashout"`
	MetaTitle                   *string   `gorm:"type:text" json:"metaTitle"`
	MetaKeywords                *string   `gorm:"type:text" json:"metaKeywords"`
	MetaDescription             *string   `gorm:"type:text" json:"metaDescription"`
	PartnerIdentity             *string   `gorm:"type:text" json:"partnerIdentity"`
	LegacyId                    *string   `gorm:"type:text" json:"legacyId"`
	OverrideFee                 *float64  `gorm:"type:float" json:"overrideFee"`
	Position                    *int      `gorm:"type:int" json:"position"`

	BaseEntity

	// Foreign key relationships
	CategoriesCodes      []string    `gorm:"-" json:"-"`
	LanguageCODE         *string     `gorm:"type:text" json:"-"`
	AffiliatePartnerCODE *string     `gorm:"type:text" json:"-"`
	CountriesCodes       []string    `gorm:"-" json:"-"`
	Country              *[]Country  `gorm:"many2many:store_country;foreignKey:Uuid;joinForeignKey:StoreUUID;References:abbreviation;joinReferences:CountryCode" json:"country"`
	Category             *[]Category `gorm:"many2many:store_category;foreignKey:Uuid;joinForeignKey:StoreUUID;References:code;joinReferences:CategoryCode" json:"category"`

	// Optional: Define GORM relationships if needed
	Language         Language `gorm:"foreignKey:LanguageCODE; references:code" json:"language"`
	AffiliatePartner Partner  `gorm:"foreignKey:AffiliatePartnerCODE; references:code" json:"affiliatePartner"`
}

type StoreWeb struct {
	Uuid uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name string    `gorm:"type:text;not null" json:"name"`
	Logo *string   `gorm:"type:text" json:"logo"`
}

func (f *Store) AfterFind(tx *gorm.DB) (err error) {
	url := dotenv.GetEnv("MS_IMAGES_SERVER_PUBLIC_URL")
	if f.Logo != nil && *f.Logo != "" {
		f.Logo = stringPointer(fmt.Sprintf(url+"%s/logo.webp", *f.Logo))
	}
	if f.Banner != nil && *f.Banner != "" {
		f.Banner = stringPointer(fmt.Sprintf(url+"%s/resized.webp", *f.Banner))
	}
	return
}
func (f *Store) AfterCreate(tx *gorm.DB) (err error) {
	url := dotenv.GetEnv("MS_IMAGES_SERVER_PUBLIC_URL")
	if f.Logo != nil && *f.Logo != "" {
		f.Logo = stringPointer(fmt.Sprintf(url+"%s/logo.webp", *f.Logo))
	}
	if f.Banner != nil && *f.Banner != "" {
		f.Banner = stringPointer(fmt.Sprintf(url+"%s/resized.webp", *f.Banner))
	}
	return
}

func (f *Store) AfterSave(tx *gorm.DB) (err error) {
	url := dotenv.GetEnv("MS_IMAGES_SERVER_PUBLIC_URL")
	if f.Logo != nil && *f.Logo != "" {
		f.Logo = stringPointer(fmt.Sprintf(url+"%s/logo.webp", *f.Logo))
	}
	if f.Banner != nil && *f.Banner != "" {
		f.Banner = stringPointer(fmt.Sprintf(url+"%s/resized.webp", *f.Banner))
	}
	return
}

func stringPointer(s string) *string {
	return &s
}
