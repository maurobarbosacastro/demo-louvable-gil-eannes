import {shopifyApi} from '@shopify/shopify-api';
import {apiVersion} from './shopify.server';

export const shopifyApiClient = shopifyApi({
  apiKey: process.env.SHOPIFY_API_KEY,
  apiSecretKey: process.env.SHOPIFY_API_SECRET || '',
  apiVersion: apiVersion,
  appUrl: process.env.SHOPIFY_APP_URL || '',
  scopes: process.env.SCOPES?.split(','),
  hostScheme: 'https',
  hostName: 'ma-adjustments-scanner-holding.trycloudflare.com',
  isEmbeddedApp: true,
});
