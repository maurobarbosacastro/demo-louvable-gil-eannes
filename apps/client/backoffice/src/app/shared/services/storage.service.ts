import {Injectable} from '@angular/core';
import {AppConstants} from '@app/app.constants';
import {StoreInfo} from '@app/modules/stores/models/stores.interface';

export enum StorageMechanism {
	SessionStorage = 'sessionStorage',
	LocalStorage = 'localStorage',
}

@Injectable({
	providedIn: 'root',
})
export class StorageService {
	private _rememberMe: boolean = false;
	private _storeInfo: StoreInfo = {
		checked: false,
		checkedDate: new Date(),
	};

	constructor() {}

	set rememberMe(value: boolean) {
		this._rememberMe = value;
		this.saveInLocalStorage(AppConstants.STORAGE_KEYS.REMEMBER_ME, value);
	}

	get rememberMe(): boolean {
		this._rememberMe = this.getFromLocalStorage(AppConstants.STORAGE_KEYS.REMEMBER_ME);
		return this._rememberMe;
	}

	set storeInfo(storeInfo: StoreInfo) {
		this._storeInfo = {checked: storeInfo.checked, checkedDate: storeInfo.checkedDate};
		this.saveInLocalStorage(AppConstants.STORAGE_KEYS.STORE_INFO, {
			checked: storeInfo.checked,
			checkedDate: storeInfo.checkedDate,
		});
	}

	get storeInfo(): StoreInfo {
		this._storeInfo = this.getFromLocalStorage(AppConstants.STORAGE_KEYS.STORE_INFO);
		return this._storeInfo;
	}

	save(key: string, value: any, mechanism: StorageMechanism = StorageMechanism.LocalStorage) {
		switch (mechanism) {
			default:
			case StorageMechanism.LocalStorage:
				this.saveInLocalStorage(key, value);
				break;
			case StorageMechanism.SessionStorage:
				this.saveInSessionStorage(key, value);
				break;
		}
	}

	get(key: string, mechanism: StorageMechanism = StorageMechanism.LocalStorage) {
		switch (mechanism) {
			default:
			case StorageMechanism.LocalStorage:
				return this.getFromLocalStorage(key);
			case StorageMechanism.SessionStorage:
				return this.getFromSessionStorage(key);
		}
	}

	delete(key: string, mechanism: StorageMechanism = StorageMechanism.LocalStorage) {
		switch (mechanism) {
			default:
			case StorageMechanism.LocalStorage:
				this.deleteFromLocalStorage(key);
				break;
			case StorageMechanism.SessionStorage:
				this.deleteFromSessionStorage(key);
				break;
		}
	}

	resetRememberMe() {
		this._rememberMe = false;
		this.deleteFromLocalStorage('tpRememberMe');
	}

	private saveInSessionStorage(key: string, value: any) {
		sessionStorage.setItem(key, JSON.stringify(value));
	}

	private getFromSessionStorage(key: string) {
		const value = sessionStorage.getItem(key);
		return value ? JSON.parse(value) : null;
	}

	private saveInLocalStorage(key: string, value: any) {
		localStorage.setItem(key, JSON.stringify(value));
	}

	private getFromLocalStorage(key: string) {
		const value = localStorage.getItem(key);
		return value ? JSON.parse(value) : null;
	}

	private deleteFromLocalStorage(key: string) {
		localStorage.removeItem(key);
	}

	private deleteFromSessionStorage(key: string) {
		sessionStorage.removeItem(key);
	}
}
