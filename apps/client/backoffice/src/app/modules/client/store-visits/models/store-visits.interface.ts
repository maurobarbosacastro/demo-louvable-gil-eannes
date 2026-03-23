import {StoreInterface} from '@app/modules/client/store-visits/models/store.interface';

export interface StoreVisitsInterface {
	uuid: string;
	store: StoreInterface;
	reference: number;
	purchased: boolean;
	dateVisited: string;
	encodedReference?: string;
    user: {
        uuid: string;
        firstName: string;
        lastName: string;
    }
}

export interface StoreVisitsSearchParams {
	store: string;
	startDate: string;
	endDate: string;
	name?: string;
	reference?: string;
}

export const StoreVisitsSortingFields = {
	user: 'user',
	reference: 'reference',
	refId: 'reference',
	purchased: 'purchase',
	store: 'store_uuid',
	storeName: 'store_uuid',
	dateTime: 'created_at',
};
