package models

import "github.com/google/uuid"

type UserMigration struct {
	Id       int64     `gorm:"type:serial;primaryKey" json:"id"`
	UserUuid uuid.UUID `gorm:"type:uuid" json:"userUuid"`
	LegacyId int64     `gorm:"type:int" json:"legacyId"`
}
