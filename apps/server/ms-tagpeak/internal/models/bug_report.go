package models

type BugReportRequest struct {
	Name        string               `json:"name" validate:"required"`
	Type        string               `json:"type" validate:"required"`
	Email       string               `json:"email" validate:"required,email"`
	Description string               `json:"description" validate:"required"`
	Attachment  *BugReportAttachment `json:"attachment,omitempty"`
}

type BugReportAttachment struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
	MimeType string `json:"mimeType"`
}
