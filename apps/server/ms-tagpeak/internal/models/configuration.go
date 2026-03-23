package models

type Configuration struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Code     string `gorm:"type:text;not null" json:"code"`
	Name     string `gorm:"type:text;not null" json:"name"`
	Value    string `gorm:"type:text;not null" json:"value"`
	DataType string `gorm:"type:text;not null" json:"dataType"`
	Editable bool   `gorm:"default:true" json:"editable"`
	BaseEntity
}
