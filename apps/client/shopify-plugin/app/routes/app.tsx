import type {HeadersFunction, LoaderFunctionArgs} from '@remix-run/node';
import {Link, Outlet, useLoaderData, useRouteError, useLocation} from '@remix-run/react';
import {boundary} from '@shopify/shopify-app-remix/server';
import {AppProvider} from '@shopify/shopify-app-remix/react';
import {NavMenu} from '@shopify/app-bridge-react';
import polarisStyles from '@shopify/polaris/build/esm/styles.css?url';

import {authenticate} from '../shopify.server';
import {I18nContext, I18nManager} from '@shopify/react-i18n';

export const links = () => [{rel: 'stylesheet', href: polarisStyles}];

export const loader = async ({request}: LoaderFunctionArgs) => {
	await authenticate.admin(request);

	return {apiKey: process.env.SHOPIFY_API_KEY || ''};
};

const locale = 'en';
const i18nManager = new I18nManager({
	locale,
	onError(error) {},
});

export default function App() {
	const {apiKey} = useLoaderData<typeof loader>();
	const location = useLocation();
	const isWelcomePage = location.pathname === '/app/welcome';

	return (
		<AppProvider isEmbeddedApp apiKey={apiKey}>
			<I18nContext.Provider value={i18nManager}>
				{!isWelcomePage && (
					<NavMenu>
						<Link to="/app" rel="home">
							Home
						</Link>
						<Link to="/app/dashboard">Dashboard</Link>
						<Link to="/app/settings">Settings</Link>
						<Link to="/app/faq">FAQ</Link>
					</NavMenu>
				)}
				<Outlet />
			</I18nContext.Provider>
		</AppProvider>
	);
}

// Shopify needs Remix to catch some thrown responses, so that their headers are included in the response.
export function ErrorBoundary() {
	return boundary.error(useRouteError());
}

export const headers: HeadersFunction = (headersArgs) => {
	return boundary.headers(headersArgs);
};
