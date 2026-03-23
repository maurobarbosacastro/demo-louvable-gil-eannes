export interface UserPaymentMethodInterface {
	uuid: string;
	paymentMethod: string;
	bankName: string;
	bankAddress: string;
	bankCountry: string;
	country: string;
	vat: string;
	bankAccountTitle: string;
	iban: string;
	ibanStatement: {
		uuid: string;
		fileName: string;
	};
    state: 'PENDING' | 'VALIDATED' | 'REJECTED';
}

export type CreateUserPaymentMethod = Omit<UserPaymentMethodInterface, 'uuid' | 'ibanStatement'> & {
	ibanStatement: string;
};

export interface AvailablePaymentMethod {
	uuid: string;
	name: string;
	code: string;
}
