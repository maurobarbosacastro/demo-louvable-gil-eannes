package models

type Example struct {
	Uuid  string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name  string
	Alias *string
	BaseEntity
}
