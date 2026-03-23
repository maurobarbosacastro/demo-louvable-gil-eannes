package dto

import (
	"github.com/google/uuid"
	"time"
)

type RewardHistoryDTO struct {
	UUID       *uuid.UUID `json:"uuid,omitempty"`
	RewardUUID *uuid.UUID `json:"rewardUuid,omitempty"`
	Rate       *float64   `json:"rate,omitempty"`
	Units      *float64   `json:"units,omitempty"`
	CashReward *float64   `json:"cashReward,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
}
