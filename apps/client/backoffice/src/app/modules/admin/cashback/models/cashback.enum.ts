export enum CashbackTabEnum {
	MANAGE_CASHBACK = 'manage-cashback',
	CASHOUT_REQUEST = 'cashout-request',
}

export enum CurrencyEnum {
	EUR = 'EUR',
	USD = 'USD',
}

export const CashbackSortingFields = {
	id: 'transaction_uuid',
	store: 'store_name',
	exitStoreUser: 'store_visit_uuid',
	date: 'date',
	amountTarget: 'amount_target',
	networkCommission: 'network_commission',
	stopDate: 'end_date',
	currentRewardTarget: 'current_reward_target',
	status: 'status',
	title: 'title',
	priceDayZero: 'initial_price',
	isin: 'isin',
};
