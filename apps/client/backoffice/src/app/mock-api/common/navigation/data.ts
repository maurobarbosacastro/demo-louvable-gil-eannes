/* eslint-disable */
import {FuseNavigationItem} from '@fuse/components/navigation';
import {environment} from '@environments/environment';

export const defaultNavigation: FuseNavigationItem[] = [
	{
		id: 'example',
		title: 'Example',
		type: 'basic',
		icon: 'heroicons_outline:chart-pie',
		link: '/example',
	},
	{
		id: 'email',
		title: 'Email',
		type: 'basic',
		icon: 'heroicons_outline:envelope',
		link: '/email',
	},
];
export const futuristicNavigation: FuseNavigationItem[] = [
	{
		id: 'example',
		title: 'Example',
		type: 'basic',
		icon: 'heroicons_outline:chart-pie',
		link: '/example',
	},
];

//Admin routes
export const compactNavigation: FuseNavigationItem[] = [
	{
		id: 'dashboard',
		title: 'Dashboard',
		tooltip: 'Dashboard',
		type: 'basic',
		icon: 'chart',
		link: '/admin/dashboard',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'cashback',
		title: 'Cashback',
		tooltip: 'Cashback',
		type: 'basic',
		icon: 'cashback',
		link: '/admin/cashbacks',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'stores',
		title: 'Stores',
		tooltip: 'Stores',
		type: 'basic',
		icon: 'store',
		link: '/admin/stores',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'store-visits',
		title: 'Store Visits',
		tooltip: 'Store Visits',
		type: 'basic',
		icon: 'eye',
		link: '/admin/store-visits',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'countries',
		title: 'Countries',
		tooltip: 'Countries',
		type: 'basic',
		icon: 'flag',
		link: '/admin/countries',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'partners',
		title: 'Sources',
		tooltip: 'Sources',
		type: 'basic',
		icon: 'heart-handshake',
		link: '/admin/sources',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'languages',
		title: 'Languages',
		tooltip: 'Languages',
		type: 'basic',
		icon: 'globe',
		link: '/admin/languages',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'categories',
		title: 'Categories',
		tooltip: 'Categories',
		type: 'basic',
		icon: 'category',
		link: '/admin/categories',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'email',
		title: 'Email',
		tooltip: 'Email',
		type: 'basic',
		icon: 'envelope',
		link: '/admin/email',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'configurations',
		title: 'Configurations',
		tooltip: 'Configurations',
		type: 'basic',
		icon: 'settings',
		link: '/admin/configurations',
		classes: {
			title: 'text-base font-plex',
		},
	},
];

//Client routes
export const horizontalNavigation: FuseNavigationItem[] = [
	{
		id: 'shop',
		title: 'Shop',
		link: '#/stores',
		type: 'basic',
		externalLink: true,
		target: '_blank',
		icon: 'external-link',
		classes: {
			title: 'font-plex',
			icon: 'w-2 min-w-2',
		},
	},
	{
		id: 'dashboard',
		title: 'Dashboard',
		type: 'basic',
		link: '/client/dashboard',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'referrals',
		title: 'Refer & Earn',
		type: 'basic',
		link: '/client/referrals',
		classes: {
			title: 'text-base font-plex',
		},
	},
	{
		id: 'faq',
		title: 'FAQ',
		link: environment.commercial.faqUrl,
		externalLink: true,
		type: 'basic',
		target: '_blank'
	},
	{
		id: 'help',
		title: 'Help',
		type: 'basic',
		externalLink: true,
		target: '_blank',
	},
];

//Shop routes
export const compactNavigationShop: FuseNavigationItem[] = [
    {
        id: 'home',
        title: 'Home',
        tooltip: 'Home',
        type: 'basic',
        icon: 'chart',
        link: '/shop/home',
        classes: {
            title: 'text-base font-plex',
        },
    },
];
