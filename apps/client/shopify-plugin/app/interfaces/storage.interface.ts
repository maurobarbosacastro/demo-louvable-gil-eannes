export interface TpShopStorageInterface {
	uuid: string;
	storeUuid: string;
	shopUuid: string;
	userUuid: string;
}

export interface ShopStorageInterface {
	uuid: string
	url: string
	state: string
	installationDone: boolean
	createdAt: string
	createdBy: string
	updatedAt: string
	updatedBy: any
	deletedAt: any
	deletedBy: any
}
