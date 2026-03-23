package models

import (
	"github.com/google/uuid"
)

type RewardHistory struct {
	UUID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	RewardUUID uuid.UUID `gorm:"type:uuid" json:"-"`
	Rate       float64   `gorm:"type:double precision" json:"rate"`
	Units      float64   `gorm:"type:double precision" json:"units"`
	CashReward float64   `gorm:"type:double precision" json:"cashReward"`

	BaseEntity

	Reward Reward `gorm:"foreignKey:RewardUUID" json:"reward"`
}
