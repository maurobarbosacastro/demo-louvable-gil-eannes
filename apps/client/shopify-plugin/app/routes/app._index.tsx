import type {LoaderFunctionArgs} from '@remix-run/node';
import {Card, Grid, Page, SkeletonBodyText, SkeletonPage} from '@shopify/polaris';
import {TitleBar, useAppBridge} from '@shopify/app-bridge-react';
import {useLoaderData, useNavigate, useSearchParams} from '@remix-run/react';
import {authenticate} from '../shopify.server';
import {useEffect, useState} from 'react';

export const loader = async ({request, context}: LoaderFunctionArgs) => {
	const {session} = await authenticate.admin(request);
	const beUrl = process.env.BE_URL;
	const boUrl = process.env.BO_URL;

	let isRegistered = false;
	let shopifyShopUuid = null;
	// @ts-ignore
	let response: {uuid: string; exists: boolean} = null;

	try {
		const call = await fetch(`${beUrl}/tagpeak/public/shopify/shop/${session.shop}/exists`, {
			method: 'GET',
		});

		response = await call.json();
	} catch (e) {
		console.log(e);
	}

	if (response && response.exists) {
		isRegistered = true;
		shopifyShopUuid = response.uuid;
	}

	return {
		shop: session.shop,
		accessToken: session.accessToken,
		beUrl,
		boUrl,
		isRegistered,
		shopifyShopUuid,
	};
};

export default function Index() {
	const data = useLoaderData<typeof loader>();
	const navigate = useNavigate();
	const appBridge = useAppBridge();
	const [searchParams, setSearchParams] = useSearchParams();
	const [loggedIn, setLoggedIn] = useState(false);
	const [isLoading, setIsLoading] = useState(true);
	const [isSetup, setIsSetup] = useState(false);

	// Refactor this to be done in the server (loader) instead of client.
	// store tagpeak token app storage and retrieve it by loader whenever needed
	useEffect(() => {
		appBridge.loading(true);
		setIsLoading(true);

		//Check if shop is already registered by the shop url/name
		if (!data.isRegistered) {
			appBridge.loading(false);
			//Redirect user to shop sign up page
			window.open(`${data.boUrl}/#/shopify/shop-sign-up?shop=${data.shop}`, '_top');
		} else {
			//Shop is already registered, check if he is signed in or coming from the sign in page

			//If there is a query param tok, it means that the user is coming from the sign in page
			if (searchParams.get('tok')) {
				// Save it locally and remove it from the url
				sessionStorage.setItem('tp_accessToken', searchParams.get('tok') as string);
			}

			//If there is no query param tok and no token in the local storage, redirect to sign in page
			if (!searchParams.get('tok') && !sessionStorage.getItem('tp_accessToken')) {
				appBridge.loading(false);
				window.open(`${data.boUrl}/#/shopify/shop-sign-in?shop=${data.shop}`, '_top');
				return;
			}

			//If there is a token in the local storage, check validity against ms-tagpeak
			if (sessionStorage.getItem('tp_accessToken')) {
				fetch(`${data.beUrl}/tagpeak/auth/valid`, {
					method: 'GET',
					headers: {
						'Content-Type': 'application/json',
						'Authorization': sessionStorage.getItem('tp_accessToken') as string,
					},
				})
					.then((res) => {
						if (!res.ok) {
							throw new Error('Invalid token');
						}
						setLoggedIn(true);

						fetch(`${data.beUrl}/shopify/shop?url=${data.shop}`, {
							method: 'GET',
							// @ts-ignore
							headers: {
								'Content-Type': 'application/json',
								'Authorization': sessionStorage.getItem('tp_accessToken') as string,
								'X-Shopify-Shop': data.shop,
								'X-TP-Shopify-Token': data.accessToken,
							},
						})
							.then((res) => res.json())
							.then((shop: {uuid: string; installationDone: boolean; url: string}) => {
								sessionStorage.setItem('shop', JSON.stringify(shop));
								sessionStorage.setItem('shop_uuid', shop.uuid);

								if (shop.installationDone) {
									setIsSetup(true);
									fetch(`${data.beUrl}/tagpeak/shopify/shop/${shop.uuid}`, {
										method: 'GET',
										// @ts-ignore
										headers: {
											'Content-Type': 'application/json',
											'Authorization': sessionStorage.getItem('tp_accessToken') as string,
										},
									})
										.then((res) => {
											return res.json();
										})
										.then((shop) => {
											sessionStorage.setItem('tp_shop', JSON.stringify(shop));
											appBridge.loading(false);
											setIsLoading(false);
										});
								} else {
                  setIsSetup(false);
                  appBridge.loading(false);
                  setIsLoading(false);
                }
							});
					})
					.catch((_) => {
						appBridge.loading(false);
						window.open(`${data.boUrl}/#/shopify/shop-sign-in?shop=${data.shop}`, '_top');

						//Clear tp_token from session storage and logout from keycloak
						sessionStorage.removeItem('tp_accessToken');
            sessionStorage.removeItem('shop_uuid');
            sessionStorage.removeItem('tp_shop');
            sessionStorage.removeItem('shop');
						setLoggedIn(false);
					});
			}
		}
	}, []);

	useEffect(() => {
		if (isLoading) {
			setSearchParams((oldParams) => {
				oldParams.delete('tok');
				return oldParams;
			});
		} else {
			if (loggedIn) {
				if (isSetup) {
					navigate('dashboard');
				} else {
					navigate('welcome');
				}
			}
		}
	}, [isLoading, isSetup, loggedIn, navigate, setSearchParams]);

	return (
		<Page fullWidth>
			<TitleBar title="TagPeak - Welcome!"></TitleBar>
			<SkeletonPage fullWidth>
				<Grid>
					<Grid.Cell columnSpan={{xs: 6, sm: 3, md: 3, lg: 6, xl: 6}}>
						<Card>
							<SkeletonBodyText />
						</Card>
					</Grid.Cell>
					<Grid.Cell columnSpan={{xs: 6, sm: 3, md: 3, lg: 6, xl: 6}}>
						<Card>
							<SkeletonBodyText />
						</Card>
					</Grid.Cell>

					<Grid.Cell columnSpan={{xs: 6, sm: 6, md: 6, lg: 12, xl: 12}}>
						<Card>
							<SkeletonBodyText lines={10} />
						</Card>
					</Grid.Cell>

				</Grid>
			</SkeletonPage>
		</Page>
	);
}
