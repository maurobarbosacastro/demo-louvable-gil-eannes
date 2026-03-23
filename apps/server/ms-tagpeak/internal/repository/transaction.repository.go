package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"ms-tagpeak/internal/db"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/pagination"
	"ms-tagpeak/pkg/utils"
	"strconv"
	"strings"
	"time"
)

func GetTransaction(uuid uuid.UUID) (models.Transaction, error) {
	dbInstance := db.GetDB()

	var model models.Transaction
	// Execute the query and capture the error
	err := dbInstance.
		Preload("StoreVisit", func(db *gorm.DB) *gorm.DB {
			return db.Select("Uuid", "Reference", "StoreUUID", "User")
		}).
		Preload("Store").
		Where("uuid = ?", uuid).First(&model).Error

	if err != nil {
		return models.Transaction{}, err
	}

	return model, nil
}

func GetUserAmountTransactions(userUuid uuid.UUID) (float64, error) {
	dbInstance := db.GetDB()

	var amountSource float64
	// Execute the query and capture the error
	err := dbInstance.Model(&models.Transaction{}).
		Select("COALESCE(SUM(amount_user), 0)").
		Where("state = ?", "VALIDATED").
		Where("\"user\" = ?", userUuid).
		Scan(&amountSource).Error

	if err != nil {
		return 0, err
	}

	return amountSource, nil
}

func GetAllTransactionsForDashboard(filters dto.DashboardFiltersDTO) ([]*models.Transaction, error) {
	dbInstance := db.GetDB()
	var transactions []*models.Transaction
	startDate := time.Date(filters.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(filters.Year, time.December, 31, 23, 59, 59, 999999999, time.UTC)

	err := dbInstance.
		Where("created_at >= ?", startDate).
		Where("created_at <= ?", endDate).
		//Where("state = ?", "VALIDATED").
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func CreateTransaction(model models.Transaction) (models.Transaction, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Create(&model).Error
	if err != nil {
		return models.Transaction{}, err
	}

	return model, nil
}

func UpdateTransaction(model models.Transaction) (models.Transaction, error) {
	dbInstance := db.GetDB()
	err := dbInstance.Updates(&model).Error
	if err != nil {
		return models.Transaction{}, err
	}

	return model, nil
}

func DeleteTransaction(uuid uuid.UUID, user string) error {
	dbInstance := db.GetDB()
	err := dbInstance.Model(&models.Transaction{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"deleted":    true,
			"deleted_by": user,
		}).Error

	err = dbInstance.Delete(&models.Transaction{}, "uuid = ?", uuid).Error

	if err != nil {
		return err
	}

	return nil
}

func GetStoreVisitByTransaction(uuid uuid.UUID) (models.StoreVisit, error) {
	dbInstance := db.GetDB()

	var model models.Transaction
	// Execute the query and capture the error
	err := dbInstance.Preload("StoreVisit").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.StoreVisit{}, err
	}

	return *model.StoreVisit, nil
}

func GetStoreByTransaction(uuid uuid.UUID) (models.Store, error) {
	dbInstance := db.GetDB()

	var model models.Transaction
	// Execute the query and capture the error
	err := dbInstance.Preload("Store").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.Store{}, err
	}
	if model.Store == nil {
		return models.Store{}, nil
	}

	return *model.Store, nil
}

func GetCurrencyExchangeRateByTransaction(uuid uuid.UUID) (models.CurrencyExchangeRate, error) {
	dbInstance := db.GetDB()

	var model models.Transaction
	// Execute the query and capture the error
	err := dbInstance.Preload("CurrencyExchangeRate").Where("uuid = ?", uuid).First(&model).Error
	if err != nil {
		return models.CurrencyExchangeRate{}, err
	}

	return model.CurrencyExchangeRate, nil
}

func GetRewardByTransactionAndUser(uuid uuid.UUID, userUuid uuid.UUID) (models.Reward, error) {
	dbInstance := db.GetDB()

	var reward models.Reward

	err := dbInstance.
		Where("transaction_uuid = ?", uuid).
		Where("\"user\" = ?", userUuid).
		First(&reward).Error

	if err != nil {
		return models.Reward{}, err
	}

	return reward, nil
}

func GetRewardByTransaction(uuid uuid.UUID, state *[]string) (models.Reward, error) {
	dbInstance := db.GetDB()

	var reward models.Reward

	if state != nil && len(*state) > 0 {
		dbInstance = dbInstance.Scopes(utils.StateScope(state))
	}

	err := dbInstance.
		Where("transaction_uuid = ?", uuid).
		First(&reward).Error

	if err != nil {
		return models.Reward{}, err
	}

	return reward, nil
}

func GetTransactionBySourceId(sourceId string) (*models.Transaction, error) {
	dbInstance := db.GetDB()

	var transaction *models.Transaction // Use a pointer to handle nil values

	err := dbInstance.Preload("StoreVisit").
		Model(&models.Transaction{}).
		Where("source_id = ?", sourceId).
		First(&transaction).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			transaction = nil
		} else {
			fmt.Printf("Error executing query: %v", err)
			return nil, err // Return an error for unexpected failures
		}
	}

	return transaction, nil
}

func GetTransactionByStoreVisitUuid(storeVisitUuid string) (*models.Transaction, error) {
	dbInstance := db.GetDB()

	var transaction *models.Transaction // Use a pointer to handle nil values

	err := dbInstance.Model(&models.Transaction{}).
		Where("store_visit_uuid = ?", storeVisitUuid).
		Scan(&transaction).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			transaction = nil
		} else {
			fmt.Printf("Error executing query: %v", err)
			return nil, err // Return an error for unexpected failures
		}
	}

	return transaction, nil
}

func IsTransactionProcessed(uuid uuid.UUID) (bool, error) {
	dbInstance := db.GetDB()

	var isProcessed bool
	// Execute the query and capture the error
	err := dbInstance.Model(&models.Transaction{}).
		Where("uuid = ?", uuid).
		Select("is_processed").
		Scan(&isProcessed).
		Error

	if err != nil {
		return false, err
	}

	return isProcessed, nil
}

func UpdateTransactionToProcessed(uuid uuid.UUID, processed bool) error {

	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Transaction{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"is_processed": processed,
		}).Error

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func BulkUpdateTransactions(uuids []string, data map[string]interface{}) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Transaction{}).
		Where("uuid IN ?", uuids). // Filter rows by UUIDs
		Updates(data).             // Provide fields to update
		Error

	if err != nil {
		fmt.Printf("Error updating records: %v\n", err)
		return err
	} else {
		fmt.Println("Records updated successfully")
		return nil
	}
}

func UpdateTransactionCashback(uuid uuid.UUID, cashback float64) error {
	dbInstance := db.GetDB()

	err := dbInstance.Model(&models.Transaction{}).
		Where("uuid = ?", uuid).
		Update("cashback", cashback).
		Error

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetCashbackView(pag pagination.PaginationParams, filters dto.TransactionFiltersDTO) (
	[]response_object.CashbackViewRO,
	int64,
	error) {
	dbInstance := db.GetDB()
	var cashbackViewRO []response_object.CashbackViewRO
	totalSize := int64(0)

	sql := `SELECT * FROM cashback`
	sqlCountItems := `SELECT count(*) FROM cashback`
	instructions := []string{}

	if filters.State != nil && *filters.State != "" {
		statuses := strings.Split(*filters.State, ",")
		instructions = append(instructions, ` status IN ('`+strings.Join(statuses, "','")+`')`)
	}

	if filters.User != nil && *filters.User != "" {
		instructions = append(instructions, ` "user" = '`+*filters.User+`'`)
	}

	if filters.StoreUuid != nil && *filters.StoreUuid != "" {
		instructions = append(instructions, ` store_uuid = '`+*filters.StoreUuid+`'`)
	}

	if filters.StoreVisitUuid != nil && *filters.StoreVisitUuid != "" {
		instructions = append(instructions, ` store_visit_uuid = '`+*filters.StoreVisitUuid+`'`)
	}

	if filters.StartDate != nil && filters.EndDate == nil {
		date := filters.StartDate.Format("2006-01-02")
		instructions = append(instructions, ` date >= '`+date+`'`)
	}

	if filters.EndDate != nil && filters.StartDate == nil {
		date := filters.EndDate.Format("2006-01-02")
		instructions = append(instructions, ` date <= '`+date+`'`)
	}

	if filters.EndDate != nil && filters.StartDate != nil {
		date := filters.StartDate.Format("2006-01-02")
		endDate := filters.EndDate.Format("2006-01-02")
		instructions = append(instructions, ` date >= '`+date+`' AND date <= '`+endDate+`'`)
	}

	if len(instructions) > 0 {
		sql = sql + ` WHERE ` + strings.Join(instructions, " AND ")
		sqlCountItems = sqlCountItems + ` WHERE ` + strings.Join(instructions, " AND ")
	}

	if pag.Sort != "" {
		sql = sql + ` order by ` + pag.Sort
	}

	sql = sql + ` limit ` + strconv.Itoa(pag.Limit)
	sql = sql + ` offset ` + strconv.Itoa(pag.Page*pag.Limit)

	err := dbInstance.Raw(sql).Scan(&cashbackViewRO).Error
	errCountItems := dbInstance.Raw(sqlCountItems).Scan(&totalSize).Error

	if err != nil {
		log.Println(err)
		return nil, totalSize, err
	}

	if errCountItems != nil {
		log.Println(errCountItems)
		return nil, totalSize, errCountItems
	}

	return cashbackViewRO, totalSize, nil
}

func UserHasMoreThan1Transaction(userUuid uuid.UUID, transactionState string) (bool, error) {
	dbInstance := db.GetDB()

	var count int64
	// Execute the query and capture the error
	err := dbInstance.Model(&models.Transaction{}).
		Where("user = ?", userUuid).
		Where("state = ?", transactionState).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count >= 1, nil
}

func UserHaveTransactions(userUUID string) (bool, error) {
	dbInstance := db.GetDB()

	var haveTransaction bool
	err := dbInstance.
		Model(&models.Transaction{}).
		Select("COUNT(*) > 0").
		Where("\"user\" = ?", userUUID).
		Scan(&haveTransaction).Error

	if err != nil {
		return false, err
	}

	return haveTransaction, nil
}

func GetActiveUsersAllTime() (int64, error) {
	dbInstance := db.GetDB()
	var count int64

	err := dbInstance.Raw(`select count(distinct("user")) from transaction t`).
		Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetActiveUsersLast12Months() (int64, error) {
	dbInstance := db.GetDB()
	var count int64

	err := dbInstance.Raw(`select count(distinct("user")) from transaction t where t.order_date < NOW() - INTERVAL '12 months'`).
		Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetActiveUserCurrentMonth() (int64, error) {
	dbInstance := db.GetDB()
	var count int64

	err := dbInstance.Raw(`select count(distinct("user")) from transaction t  WHERE t.order_date  >= DATE_TRUNC('month', NOW()) AND t.order_date  < DATE_TRUNC('month', NOW()) + INTERVAL '1 month'`).
		Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetActiveUserLastMonth() (int64, error) {
	dbInstance := db.GetDB()
	var count int64

	err := dbInstance.Raw(`select count(distinct("user")) from transaction t  WHERE t.order_date >= DATE_TRUNC('month', NOW()) - INTERVAL '1 month' AND t.order_date < DATE_TRUNC('month', NOW())`).
		Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetDashboardInfo(year int, month []string) ([]response_object.DashboardViewRO, error) {
	dbInstance := db.GetDB()

	var result []response_object.DashboardViewRO

	sql := `SELECT * FROM dashboard`

	if year != 0 {
		sql += fmt.Sprintf(" WHERE year = %v", year)
	}

	if len(month) > 0 {
		sql += fmt.Sprintf(" AND month IN (%v)", strings.Join(month, ","))
	}

	err := dbInstance.Raw(sql).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTotalTransactions() (int64, error) {
	dbInstance := db.GetDB()

	var count int64
	err := dbInstance.Model(&models.Transaction{}).
		Select("COALESCE(COUNT(uuid), 0)").
		Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetTotalGMVAllTime() (float64, error) {
	dbInstance := db.GetDB()

	var count float64
	err := dbInstance.Model(&models.Transaction{}).
		Select("COALESCE(SUM(amount_target), 0)").
		Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetAvgTransactionAmount() (float64, error) {
	dbInstance := db.GetDB()

	var count float64
	err := dbInstance.Model(&models.Transaction{}).
		Select("COALESCE(AVG(amount_target), 0)").
		Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetTotalRevenue() (float64, error) {
	dbInstance := db.GetDB()

	var count float64
	err := dbInstance.Raw(`
        SELECT COALESCE(SUM(
                        CASE
                            WHEN t.manual_commission IS NOT NULL THEN t.manual_commission
                            ELSE t.commission_target
                            END
                ), 0) FROM transaction t`).
		Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetCashbackDashboard() ([]response_object.CashbackDashboardRO, error) {
	dbInstance := db.GetDB()

	var res []response_object.CashbackDashboardRO

	sql := `SELECT * FROM cashback_dashboard`
	err := dbInstance.Raw(sql).Scan(&res).Error
	if err != nil {
		return []response_object.CashbackDashboardRO{}, err
	}
	return res, nil
}

func BulkGetTransactionsByUuids(uuids []string) ([]models.Transaction, error) {
	dbInstance := db.GetDB()

	var transactions []models.Transaction
	err := dbInstance.
		Preload("StoreVisit", func(db *gorm.DB) *gorm.DB {
			return db.Select("Uuid", "Reference", "StoreUUID", "User")
		}).
		Where("uuid IN ?", uuids).Find(&transactions).Error
	if err != nil {
		fmt.Printf("Error fetching updated records: %v\n", err)
		return nil, err
	}

	return transactions, nil
}

func UpdateShopifyActionableTransactionsState() ([]models.TransactionStateUpdate, error) {
	dbInstance := db.GetDB()

	var res []models.TransactionStateUpdate

	sql := `SELECT * FROM update_shopify_actionable_transactions_state()`
	err := dbInstance.Raw(sql).Scan(&res).Error
	if err != nil {
		return []models.TransactionStateUpdate{}, err
	}
	return res, nil
}

func UpdateUserTransactionAndRewardsByCurrency(userUuid string, newCurrencyCode string) (models.TransactionsCurrencyUserUpdate, error) {
	dbInstance := db.GetDB()
	var model models.TransactionsCurrencyUserUpdate

	sql := `select * from convert_user_transactions(?, ?)`
	err := dbInstance.Raw(sql, userUuid, newCurrencyCode).Scan(&model).Error
	if err != nil {
		return models.TransactionsCurrencyUserUpdate{}, err
	}
	return model, nil
}

func UserHasMoreThanOneTransaction(userUuid string) (bool, error) {
	dbInstance := db.GetDB()
	var transactions []models.Transaction

	// Use LIMIT 2 for efficiency - database stops scanning after finding 2 rows
	// Only select uuid column to minimize data transfer
	err := dbInstance.Model(&models.Transaction{}).
		Select("uuid").
		Where("\"user\" = ?", userUuid).
		Limit(2).
		Find(&transactions).Error

	if err != nil {
		return false, err
	}

	return len(transactions) > 1, nil
}
