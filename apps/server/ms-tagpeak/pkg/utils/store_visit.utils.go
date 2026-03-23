package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
)

func StoreVisitDtoToModel(c *dto.CreateStoreVisitDTO) models.StoreVisit {
	return models.StoreVisit{
		User:      &c.User,
		Reference: &c.Reference,
		Purchase:  c.Purchase,
		StoreUUID: &c.StoreUUID,
	}
}

func BuildStoreVisitDTO(storeVisit models.StoreVisit) dto.StoreVisitDTO {

	return dto.StoreVisitDTO{
		Uuid:      &storeVisit.Uuid,
		User:      storeVisit.User,
		Reference: storeVisit.Reference,
		Purchased: &storeVisit.Purchase,
		Store:     ForStoreVisitsMap(storeVisit),
		DateTime:  &storeVisit.CreatedAt,
	}
}

func MapPaginationResultToStoreVisits(res pagination.PaginationResult, users []*models.User) *pagination.PaginationResult {
	var visits []dto.StoreVisitAdminDTO
	for _, storeVisit := range res.Data.([]dto.StoreVisitDTO) {
		visits = append(visits, MapStoreVisitToAdminDTO(storeVisit, users))
	}
	res.Data = visits
	return &res
}

// Auxiliary function to map the store visit to have the user details
func MapStoreVisitToAdminDTO(storeVisit dto.StoreVisitDTO, users []*models.User) dto.StoreVisitAdminDTO {

	var firstName string
	var lastName string
	var uuid string

	for _, user := range users {
		if user.Uuid.String() == ParseIDToUUID(*storeVisit.User).String() {
			firstName = user.FirstName
			lastName = user.LastName
			uuid = user.Uuid.String()
		}
	}
	return dto.StoreVisitAdminDTO{
		Uuid: storeVisit.Uuid,
		User: &dto.UserDto{
			Uuid:      uuid,
			FirstName: firstName,
			LastName:  lastName,
		},
		Reference: storeVisit.Reference,
		Purchased: storeVisit.Purchased,
		Store:     storeVisit.Store,
		DateTime:  storeVisit.DateTime,
	}
}
