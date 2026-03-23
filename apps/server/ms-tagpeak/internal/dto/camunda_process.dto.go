package dto

import "github.com/google/uuid"

type CreateCamundaProcessDto struct {
	ProcessInstanceKey int64     `json:"processInstanceKey"`
	ProcessId          string    `json:"processId"`
	FieldUUID          uuid.UUID `json:"variableUuid"`
}

type UpdateCamundaProcessDto struct {
	ProcessInstanceKey *int64     `json:"processInstanceKey"`
	ProcessId          *string    `json:"processId"`
	FieldUUID          *uuid.UUID `json:"variableUuid"`
}

type CamundaProcessDto struct {
	UUID               uuid.UUID `json:"uuid"`
	ProcessInstanceKey int64     `json:"processInstanceKey"`
	ProcessId          string    `json:"processId"`
	FieldUUID          uuid.UUID `json:"variableUuid"`
}
