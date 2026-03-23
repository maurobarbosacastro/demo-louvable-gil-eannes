export interface PaymentMethodInterface {
	uuid: string;
	paymentMethod: string;
	bankName: string;
	bankAddress: string;
	country: string;
	bankAccountTitle: string;
	iban: string;
	ibanStatement: string;
}

export interface AvailablePaymentMethod {
	uuid: string;
	name: string;
	code: string;
}
