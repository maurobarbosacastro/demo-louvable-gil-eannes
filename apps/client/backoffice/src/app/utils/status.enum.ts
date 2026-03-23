export enum StatusEnum {
	LIVE = 'LIVE',
	TRACKED = 'TRACKED',
	PAID = 'PAID',
	FINISHED = 'FINISHED',
	STOPPED = 'STOPPED',
	REJECTED = 'REJECTED',
	VALIDATED = 'VALIDATED',
	EXPIRED = 'EXPIRED',
	PENDING = 'PENDING',
	COMPLETED = 'COMPLETED',
	CONFIRMED = 'CONFIRMED',
	CANCELLED = 'CANCELLED',
	REQUESTED = 'REQUESTED',
}

export const transactionStatuses: StatusEnum[] = [
	StatusEnum.TRACKED,
	StatusEnum.VALIDATED,
	StatusEnum.REJECTED,
];

export const rewardStatuses: StatusEnum[] = [
	StatusEnum.TRACKED,
	StatusEnum.CONFIRMED,
	StatusEnum.REJECTED,
	StatusEnum.LIVE,
	StatusEnum.EXPIRED,
	StatusEnum.STOPPED,
	StatusEnum.PAID,
	StatusEnum.FINISHED,
];

export enum CashbackTypeEnum {
	FIXED = 'FIXED',
	PERCENTAGE = 'PERCENTAGE',
}

export enum SortListEnum {
	LATEST_STORES = 'created_at desc',
	BY_TITLE = 'name',
	MOST_POPULAR = 'most-popular',
}
