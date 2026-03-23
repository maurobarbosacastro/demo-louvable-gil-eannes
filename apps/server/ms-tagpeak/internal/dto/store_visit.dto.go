package dto

import (
	"github.com/google/uuid"
	"time"
)

type StoreVisitDTO struct {
	Uuid      *uuid.UUID        `json:"uuid"`
	User      *string           `json:"user"`
	Reference *string           `json:"reference"`
	Purchased *bool             `json:"purchased"`
	Store     *ForStoreVisitDTO `json:"store"`
	DateTime  *time.Time        `json:"dateTime"`
}

type StoreVisitAdminDTO struct {
	Uuid      *uuid.UUID        `json:"uuid"`
	User      *UserDto          `json:"user"`
	Reference *string           `json:"reference"`
	Purchased *bool             `json:"purchased"`
	Store     *ForStoreVisitDTO `json:"store"`
	DateTime  *time.Time        `json:"dateTime"`
}

type CreateStoreVisitDTO struct {
	User      string    `json:"user" validate:"required"`
	Reference string    `json:"reference" validate:"required"`
	Purchase  bool      `json:"purchase" validate:"required"`
	StoreUUID uuid.UUID `json:"store" validate:"required"`
}

type UpdateStoreVisitDTO struct {
	User      *string    `json:"user"`
	Reference *string    `json:"reference"`
	Purchase  *bool      `json:"purchase"`
	StoreUUID *uuid.UUID `json:"store"`
}

type StoreVisitFiltersDTO struct {
	StoreUUID *uuid.UUID `json:"storeUuid" query:"storeUuid"`
	DateFrom  *time.Time `json:"dateFrom" query:"dateFrom"`
	DateTo    *time.Time `json:"dateTo" query:"dateTo"`
	Name      *string    `json:"name" query:"name"`
	UserList  []string   `json:"userList"`
	Reference *string    `json:"reference,omitempty" query:"reference"`
}

type StoreVisitFiltersUuidDTO struct {
	StoreUUID *uuid.UUID `json:"storeUuid" query:"store"`
	UserUUID  *uuid.UUID `json:"userUuid" query:"userUuid"`
	DateFrom  *time.Time `json:"dateFrom" query:"startDate"`
	DateTo    *time.Time `json:"dateTo" query:"endDate"`
}
