package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
)

func GetStoreVisit(uuid uuid.UUID) (models.StoreVisit, error) {
	dbInstance := db.GetDB()

	var model models.StoreVisit
	// Execute the query and capture the error
	err := dbInstance.Preload("Store").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.StoreVisit{}, err
	}

	return model, nil
}

func GetStoreByStoreVisit(uuid uuid.UUID) (models.Store, error) {
	dbInstance := db.GetDB()

	var model models.StoreVisit
	// Execute the query and capture the error
	err := dbInstance.Preload("Store").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.Store{}, err
	}

	return model.Store, nil
}

func GetAllStoreVisitsWithPagination(pagDTO pagination.PaginationParams, filters dto.StoreVisitFiltersDTO) (*pagination.PaginationResult, error) {

	dbInstance := db.GetDB()
	var model []models.StoreVisit
	var res pagination.PaginationResult

	// Set pagination details
	res.Limit = pagDTO.Limit
	res.Page = pagDTO.Page
	res.Sort = pagDTO.Sort

	// Start building the query
	query := dbInstance.Model(&models.StoreVisit{}).Preload("Store")

	// Apply StoreUUID filter
	if filters.StoreUUID != nil {
		query = query.Where("store_uuid = ?", filters.StoreUUID)
	}
	// Apply DateFrom filter
	if filters.DateFrom != nil {
		query = query.Where("created_at >= ?", filters.DateFrom)
	}

	// Apply DateTo filter
	if filters.DateTo != nil {
		query = query.Where("created_at <= ?", filters.DateTo)
	}

	if filters.UserList != nil {
		query = query.Where("\"user\" in (?)", filters.UserList)
	}

	if filters.Reference != nil {
		query = query.Where("LOWER(reference) LIKE LOWER(?)", "%"+*filters.Reference+"%")
	}

	// Apply pagination and execute the query
	err := query.Scopes(pagination.Paginate(&model, &res, nil, query)).
		Find(&model).Error

	if err != nil {
		return &pagination.PaginationResult{}, err
	}

	// Transform the results into DTOs
	var modelDto []dto.StoreVisitDTO
	for _, storeVisit := range model {
		modelDto = append(modelDto, utils.BuildStoreVisitDTO(storeVisit))
	}
	res.Data = modelDto

	return &res, nil
}

func GetAllStoreVisitsByUserUUIDWithPagination(
	pagDTO pagination.PaginationParams,
	filters dto.StoreVisitFiltersUuidDTO,
) (*pagination.PaginationResult, error) {
	dbInstance := db.GetDB()
	var model []models.StoreVisit
	var res pagination.PaginationResult

	// Set pagination details
	res.Limit = pagDTO.Limit
	res.Page = pagDTO.Page
	res.Sort = pagDTO.Sort

	// Start building the query
	query := dbInstance.Model(&models.StoreVisit{}).Preload("Store")

	// Apply StoreUUID filter
	if filters.StoreUUID != nil {
		query = query.Where("store_uuid = ?", filters.StoreUUID)
	}

	// Apply UserUUID filter
	if filters.UserUUID != nil {
		query = query.Where("\"user\" = ?", filters.UserUUID)
	}

	// Apply DateFrom filter
	if filters.DateFrom != nil {
		query = query.Where("created_at >= ?", filters.DateFrom)
	}

	// Apply DateTo filter
	if filters.DateTo != nil {
		query = query.Where("created_at <= ?", filters.DateTo)
	}

	// Apply pagination and execute the query
	err := query.Scopes(pagination.Paginate(&model, &res, nil, query)).
		Find(&model).Error

	if err != nil {
		return &pagination.PaginationResult{}, err
	}

	// Transform the results into DTOs
	var modelDto []dto.StoreVisitDTO
	for _, storeVisit := range model {
		modelDto = append(modelDto, utils.BuildStoreVisitDTO(storeVisit))
	}
	res.Data = modelDto

	return &res, nil
}

func CreateStoreVisit(model models.StoreVisit) (dto.StoreVisitDTO, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return dto.StoreVisitDTO{}, err
	}
	modelDto := utils.BuildStoreVisitDTO(model)

	return modelDto, nil
}

func UpdateStoreVisit(model models.StoreVisit) (models.StoreVisit, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&model).Error
	if err != nil {
		return models.StoreVisit{}, err
	}

	return model, nil
}

func DeleteStoreVisit(uuid uuid.UUID, user string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Country{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.Country{}, "uuid = ?", uuid).Error

	if err != nil {
		return err
	}

	return nil
}

func GetStoreVisitByReference(reference string) (bool, error) {
	dbInstance := db.GetDB()

	count := int64(0)
	// Execute the query and capture the error
	err := dbInstance.Model(&models.StoreVisit{}).
		Scopes(utils.ActiveScope()).
		Where("reference = ?", reference).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func GetDistinctStoresVisitedByUserUUIDWithPagination(
	pagDTO pagination.PaginationParams,
	filters dto.StoreVisitFiltersUuidDTO,
) (*pagination.PaginationResult, error) {
	dbInstance := db.GetDB()
	var stores []models.Store
	var res pagination.PaginationResult

	// Set pagination details
	res.Limit = pagDTO.Limit
	res.Page = pagDTO.Page
	res.Sort = pagDTO.Sort

	// Build query to select DISTINCT stores visited by the user
	query := dbInstance.Model(&models.Store{}).
		Select("store.*").Distinct().
		Joins("JOIN store_visit sv ON store.uuid = sv.store_uuid")

	// Apply UserUUID filter
	if filters.UserUUID != nil {
		query = query.Where("sv.\"user\" = ?", filters.UserUUID)
	}

	// Apply pagination and execute query
	err := query.Scopes(pagination.Paginate(&stores, &res, nil, dbInstance)).
		Find(&stores).Error
	if err != nil {
		return &pagination.PaginationResult{}, err
	}

	// Set results into PaginationResult
	res.Data = stores

	return &res, nil
}

func GetStoreVisitByRef(ref string) (*models.StoreVisit, error) {
	dbInstance := db.GetDB()

	var storeVisit models.StoreVisit
	err := dbInstance.Model(&models.StoreVisit{}).Where("reference = ?", ref).First(&storeVisit).Error
	if err != nil {
		return nil, err
	}

	return &storeVisit, nil
}

func GetLatestRef() (*string, error) {
	dbInstance := db.GetDB()

	var ref string
	err := dbInstance.Model(&models.StoreVisit{}).
		Where("reference LIKE ?", "%TP%").
		Order("created_at desc").
		Limit(1).
		Select("reference").
		First(&ref).Error
	if err != nil {
		return nil, err
	}

	return &ref, nil
}

func BulkSetPurchasedStoreVisitsByTransactions(transactionUuids []string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Exec(
		`update store_visit
            set purchase = true
            where uuid in (
                select store_visit_uuid
                from transaction
                where uuid in (?))`,
		transactionUuids,
	).Error

	if err != nil {
		return err
	}

	return nil
}
