package service

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/utils"
	"strings"
)

func GetMembershipStatus(uuid *string, groups []string) *models.MembershipStatus {

	_, err := repository.GetReferral(utils.ParseIDToUUID(gocloak.PString(uuid)))

	if err == nil {
		level := GetMembershipLevel(groups)

		return &models.MembershipStatus{
			Level:                   level,
			PercentageOnTransaction: PercentageBaseOnMembership(level).PercentageOnTransaction,
			PercentageOnReward:      PercentageBaseOnMembership(level).PercentageOnReward,
		}
	}

	return nil
}

func GetMembershipLevel(groups []string) *string {

	membership := Configuration.MembershipLevels.Base

	for _, group := range groups {
		if strings.Contains(group, *Configuration.MembershipLevels.Silver) {
			membership = Configuration.MembershipLevels.Silver
		}

		if strings.Contains(group, *Configuration.MembershipLevels.Gold) {
			membership = Configuration.MembershipLevels.Gold
		}
	}

	return membership

}

func ValidateMembership(uuid uuid.UUID) *string {
	countReferral := CountReferralSuccessfulTransactions(uuid)
	checkSuccessfulTransaction := GetAmountUserTransactions(uuid)

	// Gold tier: 5+ referrals OR transactions >= 5000
	if countReferral >= 5 || checkSuccessfulTransaction >= 5000 {
		return Configuration.MembershipLevels.Gold
	}

	// Silver tier: 1+ referrals OR transactions >= 500
	if countReferral >= 1 || checkSuccessfulTransaction >= 500 {
		return Configuration.MembershipLevels.Silver
	}

	// Base tier: Default membership
	return Configuration.MembershipLevels.Base
}

func PercentageBaseOnMembership(membership *string) models.MembershipStatus {
	if membership == Configuration.MembershipLevels.Silver {
		return models.MembershipStatus{
			PercentageOnTransaction: Configuration.PercentageTransaction.Silver,
			PercentageOnReward:      Configuration.PercentageReward.Silver,
		}
	}

	if membership == Configuration.MembershipLevels.Gold {
		return models.MembershipStatus{
			PercentageOnTransaction: Configuration.PercentageTransaction.Gold,
			PercentageOnReward:      Configuration.PercentageReward.Gold,
		}
	}

	return models.MembershipStatus{
		PercentageOnTransaction: Configuration.PercentageTransaction.Member,
		PercentageOnReward:      Configuration.PercentageReward.Member,
	}
}

func CountReferralSuccessfulTransactions(uuid uuid.UUID) int {
	userReferrals, err := repository.GetAllReferralByUserUuid(uuid)

	if err != nil {
		return 0
	}

	var countSuccessfulTransaction int

	for _, userReferral := range userReferrals {
		if userReferral.SuccessfulFirstTransaction {
			countSuccessfulTransaction += countSuccessfulTransaction
		}
	}

	return countSuccessfulTransaction
}
