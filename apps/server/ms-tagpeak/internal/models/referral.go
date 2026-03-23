package models

import "github.com/google/uuid"

type Referral struct {
	Uuid                       uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	ReferrerUUID               *uuid.UUID `gorm:"type:uuid;" json:"referrerUUID"` // User who invites
	InviteeUUID                *uuid.UUID `gorm:"type:uuid;" json:"inviteeUUID"`  // User invited
	SuccessfulFirstTransaction bool       `gorm:"type:bool;" json:"successfulFirstTransaction"`

	BaseEntity
}

type ReferralClicks struct {
	Uuid         uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	ReferralUUID *uuid.UUID `gorm:"type:uuid;" json:"referralUUID"`
	Code         string     `gorm:"type:text" json:"code"`

	BaseEntity

	Referral *Referral `gorm:"foreignKey:ReferralUUID" json:"-"`
}

type ReferralRevenueHistory struct {
	Uuid            uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	ReferralUUID    uuid.UUID  `gorm:"type:uuid;" json:"referralUUID"`
	TransactionUUID *uuid.UUID `gorm:"type:uuid;" json:"transactionUUID"`
	RewardUUID      *uuid.UUID `gorm:"type:uuid;" json:"rewardUUID"`
	Amount          float64    `gorm:"type:double precision" json:"amount"`

	Transaction Transaction `gorm:"foreignKey:TransactionUUID" json:"-"`
	Reward      Reward      `gorm:"foreignKey:RewardUUID" json:"-"`

	BaseEntity
}
