package service

import (
	"github.com/google/uuid"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/pagination"
)

func GetWithdrawal(uuid uuid.UUID) (models.Withdrawal, error) {
	dbInstance := db.GetDB()

	var model models.Withdrawal
	// Execute the query and capture the error
	err := dbInstance.Preload("UserPaymentMethod").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.Withdrawal{}, err
	}

	return model, nil
}

func GetAllWithdrawalsWithPagination(
	pageDTO pagination.PaginationParams,
	filters dto.WithdrawalFiltersDTO) (*pagination.PaginationResult, error) {

	dbInstance := db.GetDB()
	var model []models.Withdrawal
	var res pagination.PaginationResult

	res.Limit = pageDTO.Limit
	res.Page = pageDTO.Page
	res.Sort = pageDTO.Sort

	query := dbInstance.Preload("UserPaymentMethod")

	if filters.User != nil {
		query = query.Where("\"user\" = ?", filters.User)
	}

	if filters.State != "" {
		query = query.Where("state = ?", filters.State)
	}

	if filters.StartDate != nil {
		query = query.Where("created_at >= ?", filters.StartDate)
	}

	// Apply DateTo filter
	if filters.EndDate != nil {
		query = query.Where("created_at <= ?", filters.EndDate)
	}

	err := query.Scopes(pagination.Paginate(&model, &res, nil, dbInstance)).
		Find(&model).Error

	if err != nil {
		return &pagination.PaginationResult{}, err
	}
	res.Data = model

	return &res, nil
}

func CreateWithdrawal(model models.Withdrawal) (models.Withdrawal, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.Withdrawal{}, err
	}

	return model, nil
}

func UpdateWithdrawal(model models.Withdrawal) (models.Withdrawal, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Save(&model).Error
	if err != nil {
		return models.Withdrawal{}, err
	}

	return model, nil
}

func DeleteWithdrawal(uuid uuid.UUID, user string) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Withdrawal{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.Withdrawal{}, "uuid = ?", uuid).Error

	if err != nil {
		return err
	}

	return nil
}

func GetPaidWithdrawalsAmount(uuidUser string) (float64, error) {
	dbInstance := db.GetDB()

	var res float64
	err := dbInstance.Model(&models.Withdrawal{}).
		Where(" \"user\" = ?", uuidUser).
		Where("state = ?", "COMPLETED").
		Select("COALESCE(SUM(amount_source), 0)").
		Scan(&res).Error

	if err != nil {
		return 0, err
	}

	return res, nil
}

func GetLatestPendingWithdrawal(uuidUser string) (bool, error) {
	dbInstance := db.GetDB()

	var saved bool
	err := dbInstance.Model(&models.Withdrawal{}).
		Select("1").
		Where("\"user\" = ?", uuidUser).
		Where("state = 'PENDING'").
		Scan(&saved).
		Error

	if err != nil {
		return false, err
	}

	return saved, nil
}
