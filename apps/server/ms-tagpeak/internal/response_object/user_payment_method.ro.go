package response_object

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"ms-tagpeak/internal/models"
)

type UserPaymentMethodRo struct {
	Uuid             uuid.UUID        `json:"uuid,omitempty"`
	PaymentMethod    *string          `json:"paymentMethod,omitempty"`
	BankName         *string          `json:"bankName,omitempty"`
	BankAddress      *string          `json:"bankAddress,omitempty"`
	BankCountry      *string          `json:"bankCountry,omitempty"`
	Country          *string          `json:"country,omitempty"`
	BankAccountTitle *string          `json:"bankAccountTitle,omitempty"`
	Iban             *string          `json:"iban,omitempty"`
	IbanStatement    *FileStatementRO `json:"ibanStatement,omitempty"`
	Vat              *string          `json:"vat,omitempty"`
	State            *string          `json:"state,omitempty"`
}

type FileStatementRO struct {
	Uuid     *uuid.UUID `json:"uuid,omitempty"`
	FileName *string    `json:"fileName,omitempty"`
}

type UserPaymentMethodRO struct {
	PaymentMethod    string `json:"paymentMethod"`
	BankName         string `json:"bankName"`
	BankAddress      string `json:"bankAddress"`
	BankCountry      string `json:"bankCountry"`
	Country          string `json:"country"`
	BankAccountTitle string `json:"bankAccountTitle"`
	Iban             string `json:"iban"`
	IbanStatement    string `json:"ibanStatement"`
	Vat              string `json:"vat"`
	State            string `json:"state"`
}

func MapUserPaymentMethod(model *models.UserPaymentMethod) *UserPaymentMethodRo {

	var details map[string]string
	if err := json.Unmarshal([]byte(*model.Information), &details); err != nil {
		fmt.Printf("failed to unmarshal JSON: %w", err)
		return nil
	}

	bankName := details["bankName"]
	bankAddress := details["bankAddress"]
	bankCountry := details["bankCountry"]
	country := details["country"]
	bankAccountTitle := details["bankAccountTitle"]
	iban := details["iban"]
	vat := details["vat"]

	ibanStatement := &FileStatementRO{
		Uuid:     &model.File.Uuid,
		FileName: model.File.Name,
	}

	return &UserPaymentMethodRo{
		Uuid:             model.Uuid,
		PaymentMethod:    model.PaymentMethod.Name,
		BankName:         &bankName,
		BankAddress:      &bankAddress,
		BankCountry:      &bankCountry,
		Country:          &country,
		BankAccountTitle: &bankAccountTitle,
		Iban:             &iban,
		IbanStatement:    ibanStatement,
		Vat:              &vat,
		State:            &model.State,
	}
}

func MapInformationRO(userMethod models.UserPaymentMethod, paymentMethod *models.PaymentMethod) *UserPaymentMethodRO {

	var details map[string]string
	if err := json.Unmarshal([]byte(*userMethod.Information), &details); err != nil {
		fmt.Printf("failed to unmarshal JSON: %w", err)
		return nil
	}

	bankName := details["bankName"]
	bankAddress := details["bankAddress"]
	bankCountry := details["bankCountry"]
	country := details["country"]
	bankAccountTitle := details["bankAccountTitle"]
	iban := details["iban"]
	vat := details["vat"]

	return &UserPaymentMethodRO{
		PaymentMethod:    *paymentMethod.Name,
		BankName:         bankName,
		BankAddress:      bankAddress,
		BankCountry:      bankCountry,
		Country:          country,
		BankAccountTitle: bankAccountTitle,
		Iban:             iban,
		IbanStatement:    userMethod.FileUUID.String(),
		Vat:              vat,
		State:            userMethod.State,
	}
}
