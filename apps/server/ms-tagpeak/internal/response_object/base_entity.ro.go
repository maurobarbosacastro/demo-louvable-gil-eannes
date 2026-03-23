package response_object

import (
	"gorm.io/gorm"
	"time"
)

type BaseEntityRO struct {
	CreatedAt time.Time       `json:"createdAt,omitempty"`
	CreatedBy string          `json:"createdBy,omitempty"`
	UpdatedAt *time.Time      `json:"updatedAt,omitempty"`
	UpdatedBy *string         `json:"updatedBy,omitempty"`
	Deleted   bool            `json:"deleted,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty" swaggertype:"string" format:"date-time"`
	DeletedBy *string         `json:"deletedBy,omitempty"`
}
