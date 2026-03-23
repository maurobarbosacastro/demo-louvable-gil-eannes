import {Image} from '@app/shared/interfaces/image.interface';

export interface User {
	email: string;
	firstName: string;
	lastName: string;
	profilePicture: Image;
	birthDate: string;
	age?: string;
	genre: string;
	country: string;
	open: boolean;
	isBlocked?: boolean;
}

export interface APIUsersList {
	pageNumber: number;
	totalPages: number;
	content: User[];
}

export interface TagpeakUser {
	uuid?: string;
	referralCode?: string;
	lastName?: string;
	groups?: string[];
	firstName?: string;
	email?: string;
	displayName?: string;
	country?: string;
	currency?: string;
	balance?: number;
	birthDate?: string;
	avatar?: string;
	isVerified: string;
	onboardingFinished: string;
	utmParams?: string;
	profilePicture?: string;
	newsletter?: boolean;
	source?: string;
	currencySelected?: boolean;
}

export interface MembershipLevelInterface {
	level: string;
	valueSpent: number;
}
