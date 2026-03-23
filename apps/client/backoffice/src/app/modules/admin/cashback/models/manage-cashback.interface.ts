import {StatusEnum} from '@app/utils/status.enum';

/**
 * @deprecated Use `CashbackResponse` instead.
 */
export interface ManageCashbackInterface {
	id: string;
	exitStoreUser: string;
	date: string;
	title: string;
	orderValue: number;
	networkCommission: number;
	currentReward: number;
	priceDayZero: number;
	isin: string;
	status: StatusEnum;
}

export interface ManageCashbackSearchParams {
	state: string;
	storeUuid: string;
	storeVisitUuid: string;
	startDate: string;
	endDate: string;
	userUuid: string;
}

/**
 * @deprecated Use `CashbackResponse` instead.
 */
export interface ManageCashbackDetails extends ManageCashbackInterface {
	refId: number;
	currency: string;
	initialDate: string;
	endDate: string;
	details: string;
	cashback: number;
	orderDate: string;
	overridePrice: number;
	unit: number;
	currentCashReward: number;
}

export interface CashbackDto {
	exitClick: string;
	currency: string;
	priceDayZero: number;
	initialDate: string;
	endDate: string;
	isin: string;
    conid: string;
	title: string;
	details: string;
	orderValue: number;
	networkCommission: number;
	cashback: number;
	orderDate: string;
	status: string;
	overridePrice: number;
	uuids?: string[];
}

export interface BulkEditRewardDto {
	uuids: string[];
	transactionUuids: string[];
	initialPrice: number;
	endDate: string;
	initialDate: string;
	status: string;
	isin: string;
    conid: string;
	overridePrice: number;
}

export interface CashbackHistory {
	uuid: string;
	rate: number;
	units: number;
	cashReward: number;
	createdAt: string;
}

export interface CashbackTableResponse {
	uuid: string;
	exitId: string;
	store: {
		uuid: string;
		name: string;
		logo: string;
		percentageCashout?: number;
		cashbackValue?: number;
		cashbackType?: string;
	};
	userName: string;
	userUuid: string;
	email: string;
	date: string;
	amountSource: number;
	amountTarget: number;
	amountUser: number;
	currencySource: string;
	currencyTarget: string;
	networkCommission: number;
	status: string;
	reward: RewardCashbackResponse;
	cashback: number;

    unvalidatedCurrentReward?: number
}

export interface RewardCashbackResponse {
	uuid: string;
	isin: string;
    conid: string;
	currentRewardSource: number;
	currentRewardTarget: number;
	currentRewardUser: number;
	state: string;
	initialPrice: number;
	title: string;
	endDate: string;
	origin: string;
}

export interface CashbackResponse {
	uuid: string;
	amountSource: number;
	amountTarget: number;
	amountUser: number;
	currencySource: string;
	currencyTarget: string;
	state: string;
	commissionSource: number;
	commissionTarget: number;
	commissionUser: number;
	orderDate: string;
	user: string;
    store: {
        uuid: string;
        name: string;
        logo: string
        percentageCashout?: number;
        cashbackValue?: number;
        cashbackType?: string;
    };
	storeVisit: StoreVisitResponse;
	currencyExchangeRateUuid: string;
	cashback: number;
	createdAt: string;
	createdBy: string;
	updatedAt: any;
	updatedBy: any;
	deleted: boolean;
	deletedAt: any;
	deletedBy: any;
}

export interface StoreVisitResponse {
	uuid: string;
	user: string;
	reference: string;
	purchase: boolean;
	storeUUID: string;
	createdAt: string;
}

export interface RewardResponse {
	uuid: string;
	user: string;
	transactionUuid: string;
	isin: string;
    conid: string;
	initialReward: number;
	currentRewardSource: number;
	currentRewardTarget: number;
	currentRewardUser: number;
	currencyExchangeRateUuid: string;
	currencySource: string;
	currencyTarget: string;
	currencyUser: string;
	state: string;
	initialPrice: number;
	endDate: string;
	assetUnits: number;
	history: any;
	type: string;
	title: string;
	details: string;
	createdAt: string;
	createdBy: string;
	updatedAt: string;
	updatedBy: any;
	deleted: boolean;
	deletedAt: any;
	deletedBy: any;
	overridePrice: number;
	withdrawalUuid: string;
	origin: string;
    stoppedAt?: string;
    minimumReward?: number;

    unvalidatedCurrentReward?: number
}

export interface TransactionUpdateDTO {
	amountTarget: number;
	commissionTarget: number;
	currencySource: string;
	orderDate: string;
	state: string;
	exitClick: string;
	uuids?: string[];
}

export interface CreateRewardDTO {
	currency: string;
	details: string;
	initialDate: string;
	endDate: string;
	initialPrice: number;
	isin: string;
    conid: string;
	state: string;
	title: string;
	transactionUuid: string;
	type: string;
}

export interface RewardUpdateDTO {
	currentRewardSource: number;
	endDate: string;
	initialPrice: number;
	initialReward: number;
	isin: string;
    conid: string;
	state: string;
	initialDate: string;
	title: string;
	details: string;
	overridePrice: number;
}

export const CashbackSortingFields = {
	store: 'store_name',
	exitId: 'store_visit_uuid',
	amountUser: 'amount_user',
	stopDate: 'end_date',
	currentReward: 'current_reward_user',
	status: 'status',
};

export enum RewardOrigins {
	PURCHASE = 'PURCHASE',
	REFERRAL = 'REFERRAL',
	COMMISSION = 'COMMISSION',
	DUPLICATE = 'DUPLICATE',
}
