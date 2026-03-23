package service

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/utils"
)

func GetCamundaProcess(uuid uuid.UUID) (*models.CamundaProcess, error) {
	res, err := repository.GetUserCamundaProcess(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.CustomErrorStruct{}.NotFoundError("Payment Method", uuid)
		}
		return nil, err
	}
	return &res, nil
}

func GetCamundaProcessByVariableUUID(fieldUUID uuid.UUID) (*models.CamundaProcess, error) {
	res, err := repository.GetCamundaProcessByVariableUUID(fieldUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.CustomErrorStruct{}.NotFoundError("Camunda Process", fieldUUID)
		}
		return nil, err
	}
	return &res, nil
}

func GetCamundaProcessByVariableUUIDAndProcessId(userUUID uuid.UUID, processId string) (*models.CamundaProcess, error) {
	res, err := repository.GetCamundaProcessByVariableUUIDAndProcessId(userUUID, processId)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetAllCamundaProcesses() (*[]models.CamundaProcess, error) {
	res, err := repository.GetAllCamundaProcesses()
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func CreateCamundaProcess(dtoParam dto.CreateCamundaProcessDto) (*models.CamundaProcess, error) {

	model := utils.CamundaProcessDtoToModel(&dtoParam)

	res, err := repository.CreateUserCamundaProcess(model)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func UpdateCamundaProcess(dtoParam dto.UpdateCamundaProcessDto, uuid uuid.UUID) (*models.CamundaProcess, error) {

	toUpdate, err := repository.GetUserCamundaProcess(uuid)
	if err != nil {
		return nil, err
	}

	if dtoParam.FieldUUID != nil {
		toUpdate.FieldUUID = *dtoParam.FieldUUID
	}
	if dtoParam.ProcessInstanceKey != nil {
		toUpdate.ProcessInstanceKey = *dtoParam.ProcessInstanceKey
	}
	if dtoParam.ProcessId != nil {
		toUpdate.ProcessId = *dtoParam.ProcessId
	}

	res, err := repository.UpdateUserCamundaProcess(toUpdate)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func DeleteCamundaProcess(uuid uuid.UUID) error {

	_, err := repository.DeleteUserCamundaProcess(uuid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCamundaProcessByVariableUUID(userUUID uuid.UUID) error {

	_, err := repository.DeleteUserCamundaProcessByVariableUUID(userUUID)
	if err != nil {
		return err
	}
	return nil
}

func DynamicQueryCamunda(query string, userFilter *string, userFilters []string) (interface{}, error) {

	res, err := repository.DynamicQueryCamunda(query, userFilter, userFilters)
	if err != nil {
		return res, err
	}
	return res, nil
}
