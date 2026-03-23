package service

import (
	"fmt"
	camundaPKG "ms-tagpeak/pkg/camunda"
	"ms-tagpeak/pkg/logster"

	"github.com/google/uuid"
)

func CamundaManualTask(fieldUUID uuid.UUID, taskAction func()) error {
	logster.StartFuncLogMsg(fmt.Sprintf("fieldUUID: %s", fieldUUID))
	//Get Camunda process by uuid (Can be any uuid - Rewards, Transactions etc, need to be unique)
	camundaProcess, err := GetCamundaProcessByVariableUUID(fieldUUID)
	if err != nil {
		logster.Error(err, "Error getting camunda process")
		return nil
	}

	// Get Task
	task, err := camundaPKG.GetTaskByStateAndProcessInstanceKey(camundaPKG.CREATED, camundaProcess.ProcessInstanceKey)
	if err != nil {
		logster.Error(err, "Error getting task")
		return nil
	}

	var taskCompleted bool

	if task.ID != "" {
		// Complete task and move to next stage
		if taskAction != nil {
			taskAction()
		}

		taskCompleted, _ = camundaPKG.CompleteTask(task.ID)
	} else {
		logster.Info(fmt.Sprintf("No task found for process instance key %d", camundaProcess.ProcessInstanceKey))
	}

	if taskCompleted {
		err = DeleteCamundaProcessByVariableUUID(camundaProcess.FieldUUID)
		if err != nil {
			return err
		}
		logster.Info(fmt.Sprintf("Task %s COMPLETED\n", task.ID))
	}

	logster.EndFuncLog()
	return nil
}
