package response_object

import "github.com/google/uuid"

type StoreVisitRO struct {
	Uuid      uuid.UUID `json:"uuid"`
	User      *string   `json:"user"`
	Reference *string   `json:"reference"`
	Purchase  bool      `json:"purchase"`

	StoreUUID *uuid.UUID `json:"storeUuid"`

	BaseEntityRO
}
