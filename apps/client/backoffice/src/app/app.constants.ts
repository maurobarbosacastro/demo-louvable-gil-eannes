import {PageEvent} from '@angular/material/paginator';
import {environment} from '@environments/environment';
import {
	CssDropdownInterface,
	IconDropdownInterface,
	IconInputDropdownInterface,
} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';

export enum UserTypesEnum {
	default = '/user_type/default',
	admin = '/user_type/admin',
	influencer = '/user_type/influencer',
    shop = '/user_type/shop',
}

export enum MembershipLevelsEnum {
	base = '/membership_levels/base',
	silver = '/membership_levels/silver',
	gold = '/membership_levels/gold',
	influencer = '/membership_levels/influencer',
}

export class AppConstants {
	public static readonly PAGE_SIZE: number = 10;
	public static readonly SORT_BY_UPDATE_AT = 'updatedAt, desc';
	public static readonly DEFAULT_PAGE_EVENT: PageEvent = {
		pageIndex: 1,
		pageSize: AppConstants.PAGE_SIZE,
		length: 0,
	};
	public static readonly ROUTES = {
		admin: '/admin',
		email: '/email',
		new: '/new',
	};

	public static readonly LOGIN_CHOICE = {
		EMAIL: 'email',
		SOCIALS: 'socials',
	};

	public static readonly STORAGE_KEYS = {
		UTM: 'utm',
		REMEMBER_ME: 'tpRememberMe',
		ACCESS_TOKEN: 'accessToken',
		REFRESH_TOKEN: 'refreshToken',
		STORE_INFO: 'storeInfoChecked',
		ID_TOKEN: 'idToken',
		LOGIN_CHOICE: 'loginChoice',
	};

	public static readonly DEFAULT_SUPPORTED_IMAGE_FORMATS = environment.imageSupportedFormats
		.split(',')
		.map((s) => s.trim());

	public static readonly MODAL_CONFIG = {
		panelClass: 'dialog-panel-details',
		width: '33%',
		disableClose: true,
	};

	public static DEFAULT_STYLE_DROPDOWN: CssDropdownInterface = {
		styleHeader: 'bg-white border rounded-lg h-12',
		styleInput: 'text-project-waterloo px-5',
		styleList: 'bg-white border-0 shadow-lg rounded-lg p-2',
	};

	public static DEFAULT_ICON_CHECK_DROPDOWN: IconDropdownInterface = {
		iconName: 'check',
		iconColor: '#4237DA',
	};

	public static DEFAULT_ICON_DROPDOWN: IconInputDropdownInterface = {
		iconName: 'arrow-down',
		position: 'right',
	};

	public static CURRENCIES = [
		{
			key: 'EUR',
			label: 'EUR',
			icon: 'EUR',
			symbol: '€',
		},
		{
			key: 'GBP',
			label: 'GBP',
			icon: 'GBP',
			symbol: '£',
		},
		{
			key: 'CAD',
			label: 'CAD',
			icon: 'CAD',
			symbol: 'CA$',
		},
		{
			key: 'AUD',
			label: 'AUD',
			icon: 'AUD',
			symbol: 'A$',
		},
		{
			key: 'USD',
			label: 'USD',
			icon: 'USD',
			symbol: 'US$',
		},
		{
			key: 'BRL',
			label: 'BRL',
			icon: 'BRL',
			symbol: 'R$',
		},
		{
			key: 'MXN',
			label: 'MXN',
			icon: 'MXN',
			symbol: 'MX$',
		},
	];
}
