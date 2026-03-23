package camunda

import "time"

// Task represents a single task in the response
type TaskDTO struct {
	ID                   string     `json:"id"`
	Name                 string     `json:"name"`
	TaskDefinitionID     string     `json:"taskDefinitionId"`
	ProcessName          string     `json:"processName"`
	CreationDate         string     `json:"creationDate"`
	CompletionDate       string     `json:"completionDate"`
	Assignee             string     `json:"assignee"`
	TaskState            string     `json:"taskState"`
	SortValues           []string   `json:"sortValues"`
	IsFirst              bool       `json:"isFirst"`
	FormKey              string     `json:"formKey"`
	FormID               string     `json:"formId"`
	FormVersion          int        `json:"formVersion"`
	IsFormEmbedded       bool       `json:"isFormEmbedded"`
	ProcessDefinitionKey string     `json:"processDefinitionKey"`
	ProcessInstanceKey   string     `json:"processInstanceKey"`
	TenantID             string     `json:"tenantId"`
	DueDate              time.Time  `json:"dueDate"`
	FollowUpDate         time.Time  `json:"followUpDate"`
	CandidateGroups      []string   `json:"candidateGroups"`
	CandidateUsers       []string   `json:"candidateUsers"`
	Variables            []Variable `json:"variables"`
	Implementation       string     `json:"implementation"`
	Priority             int        `json:"priority"`
}

// Variable represents a variable in a task
type Variable struct {
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Value            string        `json:"value"`
	IsValueTruncated bool          `json:"isValueTruncated"`
	PreviewValue     string        `json:"previewValue"`
	Draft            VariableDraft `json:"draft"`
}

// VariableDraft represents a draft of a variable
type VariableDraft struct {
	Value            string `json:"value"`
	IsValueTruncated bool   `json:"isValueTruncated"`
	PreviewValue     string `json:"previewValue"`
}
