package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
)

func CamundaProcessDtoToModel(c *dto.CreateCamundaProcessDto) models.CamundaProcess {
	return models.CamundaProcess{
		FieldUUID:          c.FieldUUID,
		ProcessInstanceKey: c.ProcessInstanceKey,
		ProcessId:          c.ProcessId,
	}
}
