
export interface StoresInterface {
    uuid: string;
    name: string;
    logo: string;
    banner: string;
    shortDescription: string;
    description: string;
    averageRewardActivationTime: string;
    storeUrl: string;
    termsAndConditions: string;
    percentageCashout: number;
    metaTitle: string;
    metaKeywords: string;
    metaDescription: string;
}

export interface StoreInfo {
	checked: boolean;
	checkedDate: Date;
}
