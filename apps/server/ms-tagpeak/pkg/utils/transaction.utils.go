package utils

import (
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/response_object"
)

func TransactionDtoToModel(c *dto.CreateTransactionDTO) models.Transaction {
	return models.Transaction{
		SourceId:         c.SourceId,
		AmountSource:     c.AmountSource,
		CurrencySource:   c.CurrencySource,
		State:            c.State,
		CommissionSource: c.CommissionSource,
		OrderDate:        c.OrderDate,
		StoreVisitUUID:   c.StoreVisitUUID,
	}
}

func MapCashbackDto(transaction *models.Transaction, reward *response_object.RewardCashbackRO) response_object.CashbackRO {
	if (response_object.RewardCashbackRO{}) == *reward {
		reward = nil
	}

	var exitId *string
	if transaction.StoreVisit != nil {
		exitId = transaction.StoreVisit.Reference
	}

	var store *response_object.StoreCashbackRO
	if transaction.Store != nil {
		store = &response_object.StoreCashbackRO{
			Uuid:              transaction.Store.Uuid,
			Name:              transaction.Store.Name,
			Logo:              transaction.Store.Logo,
			PercentageCashout: transaction.Store.PercentageCashout,
			CashbackValue:     transaction.Store.CashbackValue,
			CashbackType:      transaction.Store.CashbackType,
		}
	}

	return response_object.CashbackRO{
		Uuid:              transaction.Uuid,
		ExitId:            exitId,
		Store:             store,
		UserName:          transaction.User,
		Date:              transaction.OrderDate,
		AmountTarget:      transaction.AmountTarget,
		AmountSource:      transaction.AmountSource,
		AmountUser:        transaction.AmountUser,
		CurrencySource:    transaction.CurrencySource,
		CurrencyTarget:    transaction.CurrencyTarget,
		NetworkCommission: transaction.CommissionTarget,
		Status:            transaction.State,
		Reward:            reward,
		Cashback:          transaction.Cashback,
	}
}

func MapTransactionToCashbackDetailRO(transaction models.Transaction) response_object.CashbackDetailRO {
	var store *response_object.StoreCashbackRO
	if transaction.Store != nil {
		store = &response_object.StoreCashbackRO{
			Uuid:              transaction.Store.Uuid,
			Name:              transaction.Store.Name,
			Logo:              transaction.Store.Logo,
			PercentageCashout: transaction.Store.PercentageCashout,
			CashbackValue:     transaction.Store.CashbackValue,
			CashbackType:      transaction.Store.CashbackType,
		}
	}

	var storeVisit *response_object.StoreVisitRO
	if transaction.StoreVisitUUID != nil {
		storeVisit = &response_object.StoreVisitRO{
			Uuid:      transaction.StoreVisit.Uuid,
			User:      transaction.StoreVisit.User,
			Reference: transaction.StoreVisit.Reference,
			Purchase:  transaction.StoreVisit.Purchase,
			StoreUUID: transaction.StoreVisit.StoreUUID,
		}
	}

	return response_object.CashbackDetailRO{
		Uuid:                     transaction.Uuid,
		AmountSource:             transaction.AmountSource,
		AmountTarget:             transaction.AmountTarget,
		AmountUser:               transaction.AmountUser,
		CurrencySource:           transaction.CurrencySource,
		CurrencyTarget:           transaction.CurrencyTarget,
		State:                    transaction.State,
		CommissionSource:         transaction.CommissionSource,
		CommissionTarget:         transaction.CommissionTarget,
		CommissionUser:           transaction.CommissionUser,
		OrderDate:                transaction.OrderDate,
		Cashback:                 transaction.Cashback,
		User:                     transaction.User,
		Store:                    store,
		StoreVisitUUID:           transaction.StoreVisitUUID,
		CurrencyExchangeRateUUID: transaction.CurrencyExchangeRateUUID,
		StoreVisit:               storeVisit,
		BaseEntityRO: response_object.BaseEntityRO{
			CreatedAt: transaction.CreatedAt,
			CreatedBy: transaction.CreatedBy,
			UpdatedAt: transaction.UpdatedAt,
			UpdatedBy: transaction.UpdatedBy,
			Deleted:   transaction.Deleted,
			DeletedAt: transaction.DeletedAt,
			DeletedBy: transaction.DeletedBy,
		},
	}
}
