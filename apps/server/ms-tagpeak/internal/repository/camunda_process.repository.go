package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/models"
	"strings"
)

func GetUserCamundaProcess(uuid uuid.UUID) (models.CamundaProcess, error) {
	dbInstance := db.GetDB()

	var model models.CamundaProcess
	// Execute the query and capture the error
	err := dbInstance.Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.CamundaProcess{}, err
	}

	return model, nil
}

func GetCamundaProcessByVariableUUID(userUUID uuid.UUID) (models.CamundaProcess, error) {
	dbInstance := db.GetDB()

	var userCamundaProcess models.CamundaProcess

	err := dbInstance.Where("field_uuid = ?", userUUID).First(&userCamundaProcess).Error
	if err != nil {
		return models.CamundaProcess{}, err
	}
	return userCamundaProcess, nil
}

func GetCamundaProcessByVariableUUIDAndProcessId(userUUID uuid.UUID, processId string) (models.CamundaProcess, error) {
	dbInstance := db.GetDB()

	var userCamundaProcess models.CamundaProcess

	err := dbInstance.Where("field_uuid = ? AND process_id = ?", userUUID, processId).First(&userCamundaProcess).Error
	if err != nil {
		return models.CamundaProcess{}, err
	}
	return userCamundaProcess, nil
}

func GetAllCamundaProcesses() ([]models.CamundaProcess, error) {

	dbInstance := db.GetDB()
	var model []models.CamundaProcess

	err := dbInstance.Find(&model).Error
	if err != nil {
		return []models.CamundaProcess{}, err
	}

	return model, nil
}

func CreateUserCamundaProcess(model models.CamundaProcess) (models.CamundaProcess, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.CamundaProcess{}, err
	}

	return model, nil
}

func UpdateUserCamundaProcess(model models.CamundaProcess) (models.CamundaProcess, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&model).Error
	if err != nil {
		return models.CamundaProcess{}, err
	}

	return model, nil
}

func DeleteUserCamundaProcess(uuid uuid.UUID) (models.CamundaProcess, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Delete(&models.CamundaProcess{UUID: uuid}).Error
	if err != nil {
		return models.CamundaProcess{}, err
	}

	return models.CamundaProcess{}, nil
}

func DeleteUserCamundaProcessByVariableUUID(userUUID uuid.UUID) (models.CamundaProcess, error) {
	dbInstance := db.GetDB()

	// Delete the record(s) where UserUUID matches
	err := dbInstance.Where("field_uuid = ?", userUUID).Delete(&models.CamundaProcess{}).Error
	if err != nil {
		return models.CamundaProcess{}, err
	}

	return models.CamundaProcess{}, nil
}

// DynamicQueryCamunda is a function that takes a query string and returns a list of CamundaProcesses that match the query
func DynamicQueryCamunda(query string, userFilter *string, userFilters []string) (interface{}, error) {

	dbInstance := db.GetDB()

	var res map[string]interface{}
	var args []interface{}

	if strings.Contains(query, "?") {
		if userFilter != nil && len(userFilters) == 0 {
			args = append(args, *userFilter)
		}

		if len(userFilters) > 0 {
			args = append(args, userFilters)
		}
	}

	err := dbInstance.Raw(query, args...).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
