package responses

import "time"

type NotificationResponse struct {
	Success   bool      `json:"success"`
	MessageID string    `json:"message_id,omitempty"`
	Error     string    `json:"error,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	RequestID string    `json:"request_id,omitempty"`
}
