package models

import (
	"time"
)

type BaseEntity struct {
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}
