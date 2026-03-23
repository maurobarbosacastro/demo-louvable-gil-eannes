package shopify

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy string          `json:"createdBy"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedBy *string         `json:"updatedBy"`
	DeletedAt *gorm.DeletedAt `gorm:"type:date" json:"deletedAt" swaggertype:"string" format:"date-time"`
	DeletedBy *string         `json:"deletedBy"`
}

type Shop struct {
	Uuid uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Url  string
	BaseEntity
}
