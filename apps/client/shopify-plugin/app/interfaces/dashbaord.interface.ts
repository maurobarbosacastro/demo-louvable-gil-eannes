import type {RewardCashbackResponse} from '../../../backoffice/src/app/modules/admin/cashback/models/manage-cashback.interface';

export interface ServiceResponse<T> {
	status: number;
	message: string;
	data?: T | null;
}

export interface DashboardStats {
	totalOrders: number;
	totalAmount: number;
	currency: string;
}

export interface CashbackTableResponse {
	[key: string]: unknown;

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

	unvalidatedCurrentReward?: number;
}

export interface BulkEditTransaction {
	uuids: string[];
	state: string;
}
