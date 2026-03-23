import {Page, Layout, SkeletonBodyText, Card, BlockStack} from '@shopify/polaris';
import type {LoaderFunctionArgs} from '@remix-run/node';
import {authenticate} from '../shopify.server';
import type {
	CategoriesInterface,
	CountryInterface,
	PaginationInterface,
	StoreInterface,
} from '../interfaces/common.interface';
import ShopDetails from '../components/ShopDetails.client';
import EmailVerification from '../components/EmailVerification.client';
import {useLoaderData} from '@remix-run/react';
import {useEffect, useState} from 'react';
import {getStoreDetails, updateStoreDetails} from '../service/common.service';
import type {TpShopStorageInterface} from '../interfaces/storage.interface';
import {logout} from '../utils/auth.utils';
import ShopState from '../components/ShopState.client';
import type {StoreUpdateDTO} from '../interfaces/dtos.interface';
import {TitleBar, useAppBridge} from '@shopify/app-bridge-react';

export const loader = async ({request}: LoaderFunctionArgs) => {
	const {admin, session} = await authenticate.admin(request);
	const beUrl = process.env.BE_URL;

	const responsegql = await admin.graphql(
		`#graphql
  query ShopInformation {
    shop {
      name
      id
      url
    }
  }`,
	);

	const datagql = await responsegql.json();

	const countriesCall = await fetch(`${beUrl}/tagpeak/public/countries`);
	const countriesJson: PaginationInterface<CountryInterface> = await countriesCall.json();
	const categoriesCall = await fetch(`${beUrl}/tagpeak/public/category`);
	const categoriesJson: PaginationInterface<CategoriesInterface> = await categoriesCall.json();

	return {
		displayName: datagql.data.shop.name,
		url: datagql.data.shop.url,
		shop: session.shop,
		accessToken: session.accessToken,
		countries: countriesJson,
		categories: categoriesJson,
	};
};

export default function Configuration() {
	const data = useLoaderData<typeof loader>();
	const [store, setStore] = useState<StoreInterface | null>(null);
	const [isLoading, setIsLoading] = useState(true);
	const appBridge = useAppBridge();

	const handleSave = (dto: StoreUpdateDTO) => {
		setIsLoading(true);
		updateStoreDetails(store!.uuid, dto).then((updateResponse) => {
			if (updateResponse.status === 401) {
				logout(data.shop);
				setIsLoading(false);
				return;
			}

			if (updateResponse.status === 500) {
				//Handle error
				setIsLoading(false);
				return;
			}
		});
	};

	useEffect(() => {
		setIsLoading(true);
		const tpShop = JSON.parse(sessionStorage.getItem('tp_shop')!) as TpShopStorageInterface;

		if (!tpShop) {
			setIsLoading(false);
			logout(data.shop);
			return;
		}

		getStoreDetails(tpShop.storeUuid).then((storeResponse) => {
			if (storeResponse.status === 401) {
				setIsLoading(false);
				logout(data.shop);
				return;
			}

			if (storeResponse.status === 500) {
				setIsLoading(false);
				//Handle error
				setIsLoading(false);
				return;
			}

			setStore(storeResponse.data!);
			setIsLoading(false);
		});
	}, []);

	useEffect(() => {
		appBridge.loading(isLoading);
	}, [isLoading]);

	return (
		<Page fullWidth>
			<TitleBar title="TagPeak - Settings"></TitleBar>
			{isLoading || !store ? (
				<BlockStack gap="800">
					<Card>
						<SkeletonBodyText lines={3} />
					</Card>
					<Card>
						<SkeletonBodyText lines={10} />
					</Card>
				</BlockStack>
			) : (
				<Layout>
					<Layout.Section>
						<BlockStack gap="800">
							<ShopState />

							<EmailVerification />

							<ShopDetails
								shop={data.shop}
								url={data.url}
								displayName={data.displayName}
								countries={data.countries}
								categories={data.categories}
								accessToken={data.accessToken!}
								currentStore={store}
								onSave={handleSave}
							/>
						</BlockStack>
					</Layout.Section>
				</Layout>
			)}
		</Page>
	);
}
