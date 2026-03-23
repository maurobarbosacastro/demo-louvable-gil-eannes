package models

import "github.com/google/uuid"

type Shop struct {
	Uuid             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Url              string    `gorm:"type:text" json:"url"`
	State            string    `gorm:"type:shop_states;default:'ACTIVE'" json:"state"`
	AccessToken      *string   `gorm:"type:text" json:"-"`
	InstallationDone bool      `gorm:"type:boolean" json:"installationDone"`
	BaseEntity
}

type OfflineTokenShopStruct struct {
	OfflineTokenShop string `json:"offlineToken"`
}
