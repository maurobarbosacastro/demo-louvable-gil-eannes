import {MenuItem} from './menu-client.types';

export const MENU_ITEMS: MenuItem[] = [
	{
		value: 'shop',
		label: 'Shop',
		icon: 'external-link',
		router: ['stores'],
		target: '_blank',
	},
	{
		value: 'dashboard',
		label: 'Dashboard',
		router: ['client', 'dashboard'],
	},
	{
		value: 'referrals',
		label: 'Refer & Earn',
		router: ['client', 'referrals'],
	},
];
