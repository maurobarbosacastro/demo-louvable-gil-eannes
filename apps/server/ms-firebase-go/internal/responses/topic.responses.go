package responses

import (
	"ms-firebase-go/internal/models"

	"github.com/google/uuid"
)

type TopicWithoutTarget struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
	models.BaseEntity
}
