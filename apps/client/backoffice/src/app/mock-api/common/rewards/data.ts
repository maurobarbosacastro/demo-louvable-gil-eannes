import {
	Reward,
	RewardDetails,
	RewardSummary,
} from '@app/modules/client/rewards/models/reward.interface';
import {format} from 'date-fns';
import {StatusEnum} from '@app/utils/status.enum';

export const rewards: Reward[] = [
	{
		uuid: '0',
		store: {
			uuid: '0',
			name: 'Samsung',
			logo: 'samsung',
			purchaseDate: '2024-10-01',
		},
		pricePaid: 107.49,
		status: StatusEnum.TRACKED,
	},
	{
		uuid: '1',
		store: {
			uuid: '1',
			name: 'Conforama',
			logo: '',
			purchaseDate: '2024-10-01',
		},
		pricePaid: 324.062,
		status: StatusEnum.TRACKED,
	},
	{
		uuid: '2',
		store: {
			uuid: '2',
			name: 'Auchan',
			logo: '',
			purchaseDate: '2024-08-08',
		},
		pricePaid: 89.1033,
		status: StatusEnum.TRACKED,
	},
	{
		uuid: '3',
		store: {
			uuid: '3',
			name: 'Sephora',
			logo: '',
			purchaseDate: '2024-08-08',
		},
		pricePaid: 102.98,
		stopDate: {
			date: '2025-03-21',
			daysLeft: 168,
		},
		currentReward: 2.38,
		status: StatusEnum.LIVE,
	},
	{
		uuid: '4',
		store: {
			uuid: '4',
			name: 'Booking',
			logo: '',
			purchaseDate: '2024-04-20',
		},
		pricePaid: 1104.51,
		currentReward: 30.37,
		status: StatusEnum.PAID,
	},
	{
		uuid: '5',
		store: {
			uuid: '4',
			name: 'Booking',
			logo: '',
			purchaseDate: '2024-04-20',
		},
		pricePaid: 362.96,
		currentReward: 5.91,
		status: StatusEnum.FINISHED,
	},
	{
		uuid: '6',
		store: {
			uuid: '0',
			name: 'Samsung',
			logo: 'samsung',
			purchaseDate: '2024-04-19',
		},
		pricePaid: 130.76,
		stopDate: {
			date: '2024-12-20',
			daysLeft: 77,
		},
		currentReward: 0.37,
		status: StatusEnum.LIVE,
	},
	{
		uuid: '7',
		store: {
			uuid: '7',
			name: 'Tiqets',
			logo: '',
			purchaseDate: '2024-07-19',
		},
		pricePaid: 150.02,
		stopDate: {
			date: '2024-12-20',
			daysLeft: 77,
		},
		currentReward: 3.45,
		status: StatusEnum.LIVE,
	},
	{
		uuid: '8',
		store: {
			uuid: '4',
			name: 'Booking',
			logo: '',
			purchaseDate: '2024-01-07',
		},
		pricePaid: 66.69,
		currentReward: 2.2,
		status: StatusEnum.STOPPED,
	},
	{
		uuid: '9',
		store: {
			uuid: '9',
			name: 'New Balance',
			logo: '',
			purchaseDate: '2024-01-07',
		},
		pricePaid: 74.1,
		currentReward: 2.06,
		status: StatusEnum.FINISHED,
	},
];

export const rewardSummary: RewardSummary = {
	liveRewards: 120.45,
	availableForWithdrawal: 41.8,
	valueSpent: 2090.83,
};

export const rewardDetails: RewardDetails[] = rewards.map((reward) => ({
	...reward,
	store: {
		...reward.store,
		link: 'https://www.samsung.com/pt/',
	},
	refId: 3063,
	investments: generateMockData(reward.store.purchaseDate, reward?.stopDate?.date),
}));

export function generateMockData(startDate: string, endDate?: string): {x: string; y: number}[] {
	const data: {x: string; y: number}[] = [];
	let date: Date = new Date(new Date(startDate).getTime());
	const endDateObj: Date = endDate
		? new Date(endDate) > new Date()
			? new Date()
			: new Date(endDate)
		: new Date();

	while (date < endDateObj) {
		if (endDateObj && date > endDateObj) {
			break;
		}

		const y: number = 1 + Math.random() * 3.5;
		data.push({
			x: format(date, 'MMM dd'),
			y: parseFloat(y.toFixed(2)),
		});

		date.setDate(date.getDate() + 10);
	}

	return data;
}
