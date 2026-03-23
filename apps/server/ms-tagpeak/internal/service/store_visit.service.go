package service

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ms-tagpeak/internal/constants"
	dto "ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	repository "ms-tagpeak/internal/repository"
	"ms-tagpeak/pkg/logster"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
)

func GetStoreVisit(uuid uuid.UUID) (*models.StoreVisit, error) {
	logster.StartFuncLog()
	res, err := repository.GetStoreVisit(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.CustomErrorStruct{}.NotFoundError("Store Visit", uuid)
		}
		return nil, err
	}

	logster.EndFuncLog()
	return &res, nil
}

func GetStoreByStoreVisit(uuid uuid.UUID) (*models.Store, error) {
	res, err := repository.GetStoreByStoreVisit(uuid)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetAllStoreVisits(pag pagination.PaginationParams, filters dto.StoreVisitFiltersDTO, keycloakInstance *constants.Keycloak) (*pagination.PaginationResult, error) {
	logster.StartFuncLog()

	//Get all users filtered by name provided
	users, err := GetMaxUsers(keycloakInstance, filters.Name)
	if err != nil {
		logster.Error(err, "Error getting users")
		return nil, err
	}

	if len(users) == 0 {
		return &pagination.PaginationResult{
			TotalPages: 0,
			TotalRows:  0,
			Page:       0,
			Data:       nil,
		}, nil
	}

	var userUUIDs []string
	var userDetails []dto.UserDto

	for _, user := range users {
		userUUIDs = append(userUUIDs, user.Uuid.String())
		userDetails = append(userDetails, dto.UserDto{
			Uuid:      user.Uuid.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}

	// Setting the user list to the filters to be used in the query
	filters.UserList = userUUIDs

	// Fetch store visits using the repository layer
	res, err := repository.GetAllStoreVisitsWithPagination(pag, filters)
	if err != nil {
		logster.Error(err, "Error getting store visits")
		return nil, err
	}

	// Map result to have the user details
	result := utils.MapPaginationResultToStoreVisits(*res, users)

	logster.EndFuncLog()
	return result, nil
}

func CreateStoreVisit(dtoParam dto.CreateStoreVisitDTO, uuidUser string) (*dto.StoreVisitDTO, error) {

	model := utils.StoreVisitDtoToModel(&dtoParam)
	model.CreatedBy = uuidUser

	res, err := repository.CreateStoreVisit(model)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetAllStoreVisitsByUserUuid(pag pagination.PaginationParams, filters dto.StoreVisitFiltersUuidDTO) (*pagination.PaginationResult, error) {
	res, err := repository.GetAllStoreVisitsByUserUUIDWithPagination(pag, filters)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateStoreVisit(dtoParam dto.UpdateStoreVisitDTO, uuid uuid.UUID, uuidUser string) (*models.StoreVisit, error) {
	logster.StartFuncLog()

	toUpdate, err := repository.GetStoreVisit(uuid)
	if err != nil {
		logster.Error(err, "Error getting store visit")
		logster.EndFuncLog()
		return nil, err
	}

	if dtoParam.User != nil {
		toUpdate.User = dtoParam.User
	}
	if dtoParam.Reference != nil {
		toUpdate.Reference = dtoParam.Reference
	}
	if dtoParam.Purchase != nil {
		toUpdate.Purchase = *dtoParam.Purchase
	}
	if dtoParam.StoreUUID != nil {
		toUpdate.StoreUUID = dtoParam.StoreUUID
	}

	toUpdate.UpdatedBy = &uuidUser

	res, err := repository.UpdateStoreVisit(toUpdate)
	if err != nil {
		logster.Error(err, "Error updating store visit")
		logster.EndFuncLog()
		return nil, err
	}
	logster.EndFuncLog()
	return &res, nil
}

func DeleteStoreVisit(uuid uuid.UUID, uuidUser string) error {
	err := repository.DeleteStoreVisit(uuid, uuidUser)
	if err != nil {
		return err
	}
	return nil
}

func ValidateReference(reference string) (bool, error) {
	exists, err := repository.GetStoreVisitByReference(reference)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func GetDistinctStoresVisitedByUserUuid(
	pag pagination.PaginationParams,
	filters dto.StoreVisitFiltersUuidDTO,
) (*pagination.PaginationResult, error) {
	res, err := repository.GetDistinctStoresVisitedByUserUUIDWithPagination(pag, filters)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetStoreVisitByRef(ref string) (*models.StoreVisit, error) {
	res, err := repository.GetStoreVisitByRef(ref)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetStoreVisitsAdmin(pag pagination.PaginationParams, filters dto.StoreVisitFiltersDTO) (*pagination.PaginationResult, error) {
	res, err := repository.GetAllStoreVisitsWithPagination(pag, filters)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func BulkSetPurchasedStoreVisitsByTransactions(transactionUuids []string) error {
	return repository.BulkSetPurchasedStoreVisitsByTransactions(transactionUuids)
}
