import {useLoaderData} from '@remix-run/react';
import {useEffect, useState} from 'react';
import {authenticate} from '../shopify.server';
import {TitleBar, useAppBridge} from '@shopify/app-bridge-react';
import type {LoaderFunctionArgs} from '@remix-run/node';
import {getCashbacks, getDashboardStats} from '../service/dashboard.service';
import {
	BlockStack,
	Card,
	Grid,
	Text,
	Page,
	SkeletonBodyText,
	Badge,
	InlineStack,
	Button,
	ButtonGroup,
	Spinner,
} from '@shopify/polaris';
import {useI18n} from '@shopify/react-i18n';
import type {CashbackTableResponse, DashboardStats, ServiceResponse} from '../interfaces/dashbaord.interface';
import type {PaginationInterface} from '../interfaces/common.interface';
import Cashback from '../components/Cashback';
import {logout} from 'app/utils/auth.utils';
import type {ShopStorageInterface} from '../interfaces/storage.interface';

export const loader = async ({request}: LoaderFunctionArgs) => {
	const {session} = await authenticate.admin(request);
	const beUrl = process.env.BE_URL;
	const boUrl = process.env.BO_URL;

	return {
		boUrl,
		beUrl,
		shop: session.shop,
		accessToken: session.accessToken,
	};
};

export default function App() {
	const data = useLoaderData<typeof loader>();
	const appBridge = useAppBridge();
	const [isLoading, setIsLoading] = useState(true);
  const [dashboardStats, setDashboardStats] = useState<DashboardStats>({totalOrders: 0, totalAmount: 0, currency: 'USD'});
  const [cashbacks, setCashbacks] = useState<PaginationInterface<CashbackTableResponse>>();
  const [i18n] = useI18n();
  const [currentShop, setCurrentShop] = useState<{uuid: string; shopUuid: string; storeUuid: string; userUuid: string}>();

	const [shop, setShop] = useState<ShopStorageInterface>();
	const [showCallout, setShowCallout] = useState(false);
	const [refresh, setRefresh] = useState<number>(0);
	const [isRefreshing, setIsRefreshing] = useState<boolean>(false);

  useEffect(() => {
    setTimeout(() => {
	    appBridge.loading(isLoading);
    },500);
  }, [isLoading]);

	useEffect(() => {
		setIsLoading(true);

    const currentShopTmp: {uuid: string; shopUuid: string; storeUuid: string; userUuid: string} =
      JSON.parse(sessionStorage.getItem('tp_shop')!);
    if (!currentShopTmp){
	    logout(data.shop);
      setIsLoading(false);
      return;
    }
    setCurrentShop(currentShopTmp)

		const shopTmp: ShopStorageInterface = JSON.parse(sessionStorage.getItem('shop')!);
		setShop(shopTmp);
		if (!shopTmp){
			logout(data.shop);
			setIsLoading(false);
			return;
		}
		setShowCallout(shopTmp.state !== 'ACTIVE')

    Promise.allSettled([
      getDashboardStats(currentShopTmp.shopUuid),
      getCashbacks(currentShopTmp.storeUuid)
    ]).then( ([statsPromise, cashbacksPromise]) => {
      const stats = (statsPromise as PromiseFulfilledResult<ServiceResponse<DashboardStats>>);
      const cashbacks = (cashbacksPromise as PromiseFulfilledResult<ServiceResponse<PaginationInterface<CashbackTableResponse>>>);
      if (stats.value.status === 401 || cashbacks.value.status === 401) {logout(data.shop);
        setIsLoading(false);
        return;
      }

      if (stats.value.status === 500 || cashbacks.value.status === 500) {
        //Handle error
        setIsLoading(false);
        return;
      }

      if (stats.value.data){
        setDashboardStats(stats.value.data);
      }
      if (cashbacks.value.data){
        setCashbacks(cashbacks.value.data);
      }
      setIsLoading(false);
    })

	}, []);

	useEffect(() => {
		if (currentShop){
			setIsRefreshing(true);
			Promise.allSettled([
				getDashboardStats(currentShop!.shopUuid),
				getCashbacks(currentShop!.storeUuid)
			]).then(([statsPromise, cashbacksPromise]) => {
				const stats = (statsPromise as PromiseFulfilledResult<ServiceResponse<DashboardStats>>);
				const cashbacks = (cashbacksPromise as PromiseFulfilledResult<ServiceResponse<PaginationInterface<CashbackTableResponse>>>);
				if (stats.value.status === 401 || cashbacks.value.status === 401) {
					logout(data.shop);
					setIsRefreshing(false);
					return;
				}

				if (stats.value.status === 500 || cashbacks.value.status === 500) {
					//Handle error
					setIsRefreshing(false);
					return;
				}

				if (stats.value.data){
					setDashboardStats(stats.value.data);
				}
				if (cashbacks.value.data){
					setCashbacks(cashbacks.value.data);
				}
				setIsRefreshing(false);
			})
		}
	}, [currentShop, refresh]);

	return (
		<Page fullWidth>
			<TitleBar title="TagPeak - Dashboard"></TitleBar>
			<BlockStack gap="800">
				{
					showCallout &&
					<Card roundedAbove="sm">
						<BlockStack gap="200">
							<Text as="h2" variant="headingSm">
								Plugin <Badge tone="critical">Disabled</Badge>
							</Text>
							<BlockStack gap="200">
								<Text as="h3" variant="headingSm" fontWeight="medium">
									When disabled, the plugin does not process orders. This will prevent user from receiving rewards on future orders.
								</Text>
							</BlockStack>
							<InlineStack align="end">
								<ButtonGroup>
									<Button
										variant="primary"
										onClick={() => {}}
										accessibilityLabel="Go to Settings"
									>
										Go to Settings
									</Button>
								</ButtonGroup>
							</InlineStack>
						</BlockStack>
					</Card>
				}
				{
					isLoading &&
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
				}
				{
					!isLoading && currentShop &&
					<Grid>
						<Grid.Cell columnSpan={{xs: 6, sm: 3, md: 3, lg: 6, xl: 6}}>
							<Card>
								<BlockStack gap="200">
									<Text as="h3" variant="headingLg" fontWeight="medium">
										Total Orders
									</Text>
									<Text as={'h1'} variant="headingXl" fontWeight="bold" alignment={'end'}>
										{
											isRefreshing ?
												<Spinner accessibilityLabel="Refreshing" size="small" />:
												dashboardStats.totalOrders
										}
									</Text>
								</BlockStack>
							</Card>
						</Grid.Cell>
						<Grid.Cell columnSpan={{xs: 6, sm: 3, md: 3, lg: 6, xl: 6}}>
							<Card>
								<BlockStack gap="200">
									<Text as="h3" variant="headingLg" fontWeight="medium">
										Total Amount
									</Text>
									<Text as={'h1'} variant="headingXl" fontWeight="bold" alignment={'end'}>
										{
											isRefreshing ?
												<Spinner accessibilityLabel="Refreshing" size="small" />:
												i18n.formatCurrency(dashboardStats.totalAmount, {
													currency: dashboardStats.currency,
													form: 'explicit'
												})
										}
									</Text>
								</BlockStack>
							</Card>
						</Grid.Cell>

						<Grid.Cell columnSpan={{xs: 6, sm: 6, md: 6, lg: 12, xl: 12}}>
							<Cashback
								cashback={cashbacks!}
								beUrl={data.beUrl!}
								storeUuid={currentShop!.storeUuid}
								boUrl={data.boUrl!}
								shop={data.shop}
								setRefreshParent={setRefresh}
							></Cashback>
						</Grid.Cell>

					</Grid>
				}
			</BlockStack>

    </Page>
	);
}
