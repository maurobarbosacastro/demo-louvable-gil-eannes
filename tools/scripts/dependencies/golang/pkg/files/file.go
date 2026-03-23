package files

import "github.com/google/uuid"

type File struct {
	Uuid      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name      *string   `gorm:"type:text" json:"name"`
	Extension *string   `gorm:"type:text" json:"extension"`

	BaseEntity
}
