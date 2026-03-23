package response_object

import "github.com/google/uuid"

type ReferralDto struct {
	Uuid         uuid.UUID  `json:"uuid,omitempty"`
	ReferrerUUID *uuid.UUID `json:"referrerUUID,omitempty"`
	ReferrerName string     `json:"referrerName,omitempty"`
	InviteeUUID  *uuid.UUID `json:"inviteeUUID,omitempty"`
}

type InviteeDto struct {
	Uuid        uuid.UUID  `json:"uuid,omitempty"`
	InviteeUUID *uuid.UUID `json:"inviteeUUID,omitempty"`
	InviteeName string     `json:"inviteeName,omitempty"`
}

type ReferralInfo struct {
	TotalClicks          int64       `json:"totalClicks"`
	TotalUserRegistered  int64       `json:"totalUserRegistered"`
	TotalFirstPurchase   int64       `json:"totalFirstPurchase"`
	ClicksByMonth        []MonthData `json:"clicksByMonth"`
	RegisteredByMonth    []MonthData `json:"registeredByMonth"`
	FirstPurchaseByMonth []MonthData `json:"firstPurchaseByMonth"`
}

type ReferralRevenueInfo struct {
	TotalRevenue   *float64    `json:"totalRevenue"`
	RevenueByMonth []MonthData `json:"revenueByMonth"`
}

type MonthData struct {
	Month string  `json:"month"`
	Value float64 `json:"value"`
}
