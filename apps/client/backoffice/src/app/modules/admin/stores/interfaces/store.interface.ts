export interface AdminStoreInterface {
	uuid: string;
	name: string;
	logo: string;
	banner: string;
	shortDescription: string;
	description: string;
	urlSlug: string;
	initialReward: number;
	averageRewardActivationTime: string;
	state: string;
	keywords: string;
	affiliateLink: string;
	storeUrl: string;
	termsAndConditions: string;
	cashbackType: string;
	cashbackValue: number;
	percentageCashout: number;
	metaTitle: string;
	metaKeywords: string;
	metaDescription: string;
	country: string[];
	category: string[];
	language: string;
	affiliatePartnerCode: string;
	partner?: string;
	transactionFee?: number;
	position?: number;
}

export interface CreateStoreInterface {
	name: string;
	urlSlug: string;
	averageRewardActivationTime: string;
	state: string;
	cashbackType: string;
	cashbackValue: number;
	percentageCashout: number;
	logo?: string;
	banner?: string;
	shortDescription?: string;
	description?: string;
	keywords?: string;
	affiliateLink?: string;
	storeUrl?: string;
	termsAndConditions?: string;
	metaTitle?: string;
	metaKeywords?: string;
	metaDescription?: string;
	country: string[];
	category: string[];
	languageCode: string;
	affiliatePartner?: string;
	partnerIdentity?: string;
	transactionFee?: number;
	position?: number;
}

export type TableStoreAdminInterface = Pick<
	AdminStoreInterface,
	'uuid' | 'name' | 'description' | 'partner' | 'country'
>;

export interface LogoStore {
	id: string;
	logo: string;
	original: string;
}

export const StoresSortingFields = {
	name: 'name',
	partner: 'affiliate_partner_code',
	description: 'description',
	createdAt: 'created_at',
	status: 'state',
};

export interface StoreApprovalRequest {
	uuid: string;
	name: string;
	status: string;
	createdAt: string;
	user: {
		uuid: string;
		name: string;
		email: string;
	};
}
