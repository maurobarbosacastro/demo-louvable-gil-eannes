package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"

	"github.com/google/uuid"
)

func CreateMapForReferral(model *models.Referral, user *models.User) response_object.ReferralDto {
	return response_object.ReferralDto{
		Uuid:         model.Uuid,
		ReferrerUUID: model.ReferrerUUID,
		ReferrerName: user.FirstName + " " + user.LastName,
		InviteeUUID:  model.InviteeUUID,
	}
}

func CreateMapForInvitee(model *models.Referral, inviteeUser *models.User) response_object.InviteeDto {
	return response_object.InviteeDto{
		Uuid:        model.Uuid,
		InviteeUUID: model.InviteeUUID,
		InviteeName: inviteeUser.FirstName + " " + inviteeUser.LastName,
	}
}

func CreateReferralDto(referrerUuid uuid.UUID, inviteeUuid uuid.UUID) models.Referral {
	return models.Referral{
		ReferrerUUID: &referrerUuid,
		InviteeUUID:  &inviteeUuid,
		BaseEntity: models.BaseEntity{
			CreatedBy: inviteeUuid.String(),
		},
	}
}

func BuildReferralDTO(referral models.Referral) dto.ReferralDTO {

	return dto.ReferralDTO{
		Uuid:                       &referral.Uuid,
		ReferrerUUID:               referral.ReferrerUUID,
		InviteeUUID:                referral.InviteeUUID,
		SuccessfulFirstTransaction: &referral.SuccessfulFirstTransaction,
	}
}

func CreateReferralClickDto(referralUUID *uuid.UUID, code string) models.ReferralClicks {
	return models.ReferralClicks{
		ReferralUUID: referralUUID,
		Code:         code,
	}
}

func CreateReferralRevenueDto(dto dto.CreateReferralRevenueDTO) models.ReferralRevenueHistory {
	return models.ReferralRevenueHistory{
		ReferralUUID:    dto.ReferralUUID,
		TransactionUUID: dto.TransactionUUID,
		RewardUUID:      dto.RewardUUID,
		Amount:          dto.Amount,
		BaseEntity: models.BaseEntity{
			CreatedBy: dto.CreatedBy,
		},
	}
}
