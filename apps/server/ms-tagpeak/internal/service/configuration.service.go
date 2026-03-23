package service

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"strconv"

	"github.com/samber/lo"
)

var Config *map[string]models.Configuration

var Configuration *Configs

type Configs struct {
	MembershipLevels      *MembershipLevels
	UserTypes             *UserTypes
	PercentageTransaction *PercentageTransaction
	PercentageReward      *PercentageReward
}

type MembershipLevels struct {
	Base       *string
	Silver     *string
	Gold       *string
	Influencer *string
}

type UserTypes struct {
	Default *string
	Shop    *string
}

type PercentageTransaction struct {
	Member     *float64
	Silver     *float64
	Gold       *float64
	Influencer *float64
}

type PercentageReward struct {
	Member     *float64
	Silver     *float64
	Gold       *float64
	Influencer *float64
}

func CreateConfiguration(c dto.CreateConfigurationDTO, uuidUser string) (models.Configuration, error) {

	exists, err := repository.ConfigurationExistsByCode(c.Code)
	if err != nil {
		return models.Configuration{}, err
	}

	if exists {
		return models.Configuration{}, utils.CustomErrorStruct{}.AlreadyExists("Configuration", c.Code)
	}

	model := models.Configuration{
		Name:     c.Name,
		Value:    c.Value,
		Code:     c.Code,
		DataType: c.DataType,
	}

	model.CreatedBy = uuidUser

	res, err := repository.CreateConfiguration(model)
	if err != nil {
		return models.Configuration{}, err
	}

	return res, nil
}

func GetConfigurations(pag pagination.PaginationParams) (*pagination.PaginationResult, error) {
	res, err := repository.GetConfigurations(pag)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetConfiguration(code string) (models.Configuration, error) {
	res, err := repository.GetConfigurationByCode(code)
	if err != nil {
		return models.Configuration{}, err
	}
	return res, nil
}

func UpdateConfiguration(id int, updateDto dto.UpdateConfigurationDTO, user string) (models.Configuration, error) {
	toUpdate, err := repository.GetConfiguration(id)
	if err != nil {
		return models.Configuration{}, err
	}

	if updateDto.Name != nil {
		toUpdate.Name = *updateDto.Name
	}
	if updateDto.Value != nil {
		toUpdate.Value = *updateDto.Value
	}
	if updateDto.DataType != nil {
		toUpdate.DataType = *updateDto.DataType
	}

	toUpdate.UpdatedBy = &user

	res, err := repository.UpdateConfiguration(toUpdate)
	if err != nil {
		return models.Configuration{}, err
	}
	return res, nil
}

func DeleteConfiguration(id int, user string) error {
	_, err := repository.DeleteConfiguration(id, user)
	if err != nil {
		return err
	}
	return nil
}

func GetAllConfigurations() ([]models.Configuration, error) {
	return repository.GetAllConfigurations()
}

func LoadConfigurations() {
	data, err := GetAllConfigurations()
	if err != nil {
		panic(err)
	}

	final := lo.KeyBy(data, func(c models.Configuration) string {
		return c.Code
	})

	Config = &final
}

func GetLoadedConfig(configName string) models.Configuration {
	cfg := *Config
	return cfg[configName]
}

func InitConfigs() {
	logster.StartFuncLog()

	fiftyFloat64, _ := strconv.ParseFloat("50", 64)
	zeroFloat64, _ := strconv.ParseFloat("0", 64)

	// Transaction percentage
	memberPercentageFloat64, _ := strconv.ParseFloat(GetLoadedConfig("referral_member_transaction_cash_reward").Value, 64)
	silverPercentageFloat64, _ := strconv.ParseFloat(GetLoadedConfig("referral_silver_transaction_cash_reward").Value, 64)
	goldPercentageFloat64, _ := strconv.ParseFloat(GetLoadedConfig("referral_gold_transaction_cash_reward").Value, 64)

	memberPercentageTransaction := fiftyFloat64 + memberPercentageFloat64
	silverPercentageTransaction := fiftyFloat64 + silverPercentageFloat64
	goldPercentageTransaction := fiftyFloat64 + goldPercentageFloat64

	// Reward percentage
	memberPercentageReward, _ := strconv.ParseFloat(GetLoadedConfig("referral_member_friend_cash_reward").Value, 64)
	silverPercentageReward, _ := strconv.ParseFloat(GetLoadedConfig("referral_silver_friend_reward_share").Value, 64)
	goldPercentageReward, _ := strconv.ParseFloat(GetLoadedConfig("referral_gold_friend_reward_share").Value, 64)

	config := &Configs{
		MembershipLevels: &MembershipLevels{
			Base:       utils.StringPointer("/membership_levels/base"),
			Silver:     utils.StringPointer("/membership_levels/silver"),
			Gold:       utils.StringPointer("/membership_levels/gold"),
			Influencer: utils.StringPointer("/membership_levels/influencer"),
		},
		UserTypes: &UserTypes{
			Default: utils.StringPointer("/user_type/default"),
			Shop:    utils.StringPointer("/user_type/shop"),
		},
		PercentageTransaction: &PercentageTransaction{
			Member:     utils.Float64Pointer(memberPercentageTransaction),
			Silver:     utils.Float64Pointer(silverPercentageTransaction),
			Gold:       utils.Float64Pointer(goldPercentageTransaction),
			Influencer: utils.Float64Pointer(fiftyFloat64), // This is just a default value because will be set on the user info on keycloack
		},
		PercentageReward: &PercentageReward{
			Member:     utils.Float64Pointer(memberPercentageReward),
			Silver:     utils.Float64Pointer(silverPercentageReward),
			Gold:       utils.Float64Pointer(goldPercentageReward),
			Influencer: utils.Float64Pointer(zeroFloat64), // This is just a default value because will be set on the user info on keycloack
		},
	}

	Configuration = config

	logster.EndFuncLog()
}
