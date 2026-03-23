export interface PaginationInterface<T> {
  limit: number;
  page: number;
  sort: string;
  totalPages: number;
  totalRows: number;
  data: T[];
}

export interface PaginationOptions {
  limit: number;
  page: number;
  sort: string;
}

export interface CountryInterface {
  uuid: string;
  abbreviation: string;
  currency: string;
  flag: string;
  name: string;
  enabled: false;
}

export interface CategoriesInterface {
  uuid: string;
  name: string;
  code: string;
}

export interface StoreInterface {
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
}
