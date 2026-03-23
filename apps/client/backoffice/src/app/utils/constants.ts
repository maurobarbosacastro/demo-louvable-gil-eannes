import {CashbackTypeEnum, StatusEnum} from '@app/utils/status.enum';
import {MembershipLevelsEnum} from '@app/app.constants';
import {MembershipLevelInfoInterface} from '@app/shared/interfaces/membership-level.interface';

export const statusFilters: {label: string; value: string}[] = [
	{
		label: 'status.PENDING',
		value: StatusEnum.TRACKED,
	},
	{
		label: 'status.VALIDATED',
		value: StatusEnum.VALIDATED,
	},
	{
		label: 'status.REJECTED',
		value: StatusEnum.REJECTED,
	},
	{
		label: 'status.LIVE',
		value: StatusEnum.LIVE,
	},
	{
		label: 'status.STOPPED',
		value: StatusEnum.STOPPED,
	},
	{
		label: 'status.EXPIRED',
		value: StatusEnum.EXPIRED,
	},
	{
		label: 'status.FINISHED',
		value: StatusEnum.FINISHED,
	},
	{
		label: 'status.REQUESTED',
		value: StatusEnum.REQUESTED,
	},
	{
		label: 'status.PAID',
		value: StatusEnum.PAID,
	},
];

export const cashbackTypeOptions: {label: string; value: string}[] = [
	{
		label: 'store.cashback-type-options.percentage',
		value: CashbackTypeEnum.PERCENTAGE,
	},
	{
		label: 'store.cashback-type-options.fixed',
		value: CashbackTypeEnum.FIXED,
	},
];

export const membershipLevel: MembershipLevelInfoInterface[] = [
	{
		level: MembershipLevelsEnum.base,
	},
	{
		level: MembershipLevelsEnum.silver,
	},
	{
		level: MembershipLevelsEnum.gold,
	},
];
