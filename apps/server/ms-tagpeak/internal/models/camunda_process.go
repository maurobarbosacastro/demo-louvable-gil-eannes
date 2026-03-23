package models

import "github.com/google/uuid"

type CamundaProcess struct {
	UUID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" `
	ProcessInstanceKey int64     `gorm:"type:int64"`
	ProcessId          string    `gorm:"type:text"`
	FieldUUID          uuid.UUID `gorm:"type:uuid"`
}
