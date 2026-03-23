package models

import "github.com/google/uuid"

type Partner struct {
	Uuid               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name               *string   `gorm:"type:text;not null;" json:"name"`      // Country name, not null
	Code               *string   `gorm:"type:text; unique" json:"code"`        // Country code, unique
	ECommercePlatform  *string   `gorm:"type:text" json:"eCommercePlatform"`   // eCommerce platform name, nullable
	ValidationPeriod   *int      `gorm:"type:int" json:"validationPeriod"`     // Validation period, nullable
	DeepLink           *string   `gorm:"type:text" json:"deepLink"`            // Deep link, nullable
	DeepLinkIdentifier *string   `gorm:"type:text" json:"deepLinkIdentifier"`  // Deep link identifier, nullable
	SubIdentifier      *string   `gorm:"type:text" json:"subIdentifier"`       // Sub identifier, nullable
	PercentageTagpeak  *float64  `gorm:"type:float" json:"percentageTagpeak"`  // Percentage tagpeak, nullable
	PercentageInvested *float64  `gorm:"type:float" json:"percentageInvested"` // Percentage invested, nullable
	BaseEntity
}
