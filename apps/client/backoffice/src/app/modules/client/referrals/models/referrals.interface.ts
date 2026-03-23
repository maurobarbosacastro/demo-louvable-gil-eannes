export enum ReferralsTabEnum {
	YOUR_REFERRALS = 'your-referrals',
	HOW_IT_WORKS = 'how-it-works',
}

export interface ReferralInterface {
	uuid: string;
	referrerUUID: string;
	inviteeUUID: string;
	successfulFirstTransaction: boolean;
}

export interface UserReferralInterface {
	uuid: string;
	firstName: string;
	lastName: string;
	profilePicture: string;
	referredValue: number;
	firstTransactionSuccessful: boolean;
	displayName?: string;
}

export interface RevenueInfo {
	totalRevenue: number;
	revenueByMonth: MonthData[];
}

export interface ReferralInfo {
	totalClicks: number;
	totalUserRegistered: number;
	totalFirstPurchase: number;
	clicksByMonth: MonthData[];
	registeredByMonth: MonthData[];
	firstPurchaseByMonth: MonthData[];
}

export interface MonthData {
	month: string;
	value: number;
}
