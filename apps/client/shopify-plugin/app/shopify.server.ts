import '@shopify/shopify-app-remix/adapters/node';
import {ApiVersion, AppDistribution, LogSeverity, shopifyApp} from '@shopify/shopify-app-remix/server';
import {PrismaSessionStorage} from '@shopify/shopify-app-session-storage-prisma';
import prisma from './db.server';

const shopify = shopifyApp({
	apiKey: process.env.SHOPIFY_API_KEY,
	apiSecretKey: process.env.SHOPIFY_API_SECRET || '',
	apiVersion: ApiVersion.April25,
	scopes: process.env.SCOPES?.split(','),
	appUrl: process.env.SHOPIFY_APP_URL || '',
	isEmbeddedApp: true,
	authPathPrefix: '/auth',
	sessionStorage: new PrismaSessionStorage(prisma),
	distribution: AppDistribution.AppStore,
	future: {
		unstable_newEmbeddedAuthStrategy: true,
		removeRest: true,
	},
	...(process.env.SHOP_CUSTOM_DOMAIN ? {customShopDomains: [process.env.SHOP_CUSTOM_DOMAIN]} : {}),
	useOnlineTokens: true,
	logger: {
		log: (severity, msg) => {
			switch (severity) {
				case LogSeverity.Debug:
					console.log(msg);
					break;
				case LogSeverity.Warning:
					console.warn(msg);
					break;
				case LogSeverity.Error:
					console.error(msg);
					break;
				case LogSeverity.Info:
					console.info(msg);
					break;
				default:
					console.log(msg);
			}
		},
		level: process.env.NODE_ENV === 'production' ? LogSeverity.Error : LogSeverity.Debug,
		httpRequests: true,
		timestamps: true
	},
});

export default shopify;
export const apiVersion = ApiVersion.April25;
export const addDocumentResponseHeaders = shopify.addDocumentResponseHeaders;
export const authenticate = shopify.authenticate;
export const unauthenticated = shopify.unauthenticated;
export const login = shopify.login;
export const registerWebhooks = shopify.registerWebhooks;
export const sessionStorage = shopify.sessionStorage;
