export interface WithdrawalsRequestInterface {
	uuid: string;
	user: {
		uuid: string;
		name: string;
		email: string;
	};
	state: string;
	amountTarget: number;
	currencyTarget: string;
	createdAt: string;
	completionDate: string;
	details: string;
	paymentMethod: PaymentMethodInterface;
}

export interface PaymentMethodInterface {
	paymentMethod: string;
	bankName: string;
	bankAddress: string;
	bankCountry: string;
	country: string;
	bankAccountTitle: string;
	iban: string;
	ibanStatement: string;
    vat: string;
}

export interface UpdateStateRequestDTO {
	state: string;
	completionDate: string;
	details?: string;
}

export interface BulkUpdateStateRequestDTO {
	uuids: string[];
	state: string;
	completionDate: string;
}

export interface FiltersSearchDTO {
	state?: string;
	startDate?: string;
	endDate?: string;
}

export const WithdrawalSortingFields = {
	uuid: 'uuid',
	user: 'created_by',
	completionDate: 'completion_date',
	state: 'state',
	createdAt: 'created_at',
	paymentMethod: 'user_method',
	iban: 'user_method',
	note: 'details',
};
