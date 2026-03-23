import {
	Badge,
	Button,
	InlineStack,
	useBreakpoints,
	Text,
	Box,
	BlockStack,
	Card,
} from '@shopify/polaris';
import {useCallback, useState, useEffect, useMemo} from 'react';
import {activateShop, deactivateShop} from '../service/common.service';
import type {ShopStorageInterface, TpShopStorageInterface} from '../interfaces/storage.interface';
import {logout} from '../utils/auth.utils';

type ShopStateProps = {

};


export default function ShopState(props: ShopStateProps) {
  const shop: ShopStorageInterface = useMemo( () => JSON.parse(sessionStorage.getItem('shop')!), []);
  const tpShop: TpShopStorageInterface = useMemo( () => JSON.parse(sessionStorage.getItem('tp_shop')!), []);

	const [enabled, setEnabled] = useState(shop.state === 'ACTIVE');
	const [isLoading, setIsLoading] = useState(false);

	useEffect(() => {
		if (!tpShop) {
			logout(window.location.hostname);
			return;
		}

	}, []);

	const handleToggle = useCallback(() => {
		setIsLoading(true);

		if (enabled) {
			// Deactivate shop
			deactivateShop(shop.uuid)
				.then(response => {
					if (response.status === 401) {
						logout(shop.url);
						return;
					}

					if (response.status !== 200) {
						setIsLoading(false);
						return;
					}

					// Update session storage
					const storeData = JSON.parse(
						sessionStorage.getItem('tp_store')!
					);
					if (storeData) {
						storeData.state = 'CLOSED';
						sessionStorage.setItem('tp_store', JSON.stringify(storeData));
					}

					setEnabled(false);
					setIsLoading(false);
          sessionStorage.setItem('shop', JSON.stringify({...shop, state: 'CLOSED'}));
				})
				.catch(error => {
					console.error('Error deactivating shop:', error);
					setIsLoading(false);
				});
		} else {
			// Activate shop
			activateShop(shop.uuid)
				.then(response => {
					if (response.status === 401) {
						logout(shop.url);
						return;
					}

					if (response.status !== 200) {
						setIsLoading(false);
						return;
					}

					// Update session storage
					const storeData = JSON.parse(
						sessionStorage.getItem('tp_store')!
					);
					if (storeData) {
						storeData.state = 'ACTIVE';
						sessionStorage.setItem('tp_store', JSON.stringify(storeData));
					}

					setEnabled(true);
					setIsLoading(false);
          sessionStorage.setItem('shop', JSON.stringify({...shop, state: 'ACTIVE'}));
				})
				.catch(error => {
					console.error('Error activating shop:', error);
					setIsLoading(false);
				});
		}
	}, [enabled, shop]);


	const contentStatus = enabled ? 'Disable' : 'Enable';

	const toggleId = 'setting-toggle-uuid';
	const descriptionId = 'setting-toggle-description-uuid';

	const {mdDown} = useBreakpoints();

	const badgeStatus = enabled ? 'success' : 'critical';

	const badgeContent = enabled ? 'Enabled' : 'Disabled';

	const title = 'Plugin state';
	const description =
		'Enable or Disable the plugin in your store. Disabling the plugin will stop processing orders and prevent user from receiving rewards on future orders';

	const settingStatusMarkup = (
		<Badge
			tone={badgeStatus}
			toneAndProgressLabelOverride={`Setting is ${badgeContent}`}
		>
			{badgeContent}
		</Badge>
	);

	const settingTitle = title ? (
		<InlineStack gap="200" wrap={false}>
			<InlineStack gap="200" align="start" blockAlign="baseline">
				<label htmlFor={toggleId}>
					<Text variant="headingMd" as="h6">
						{title}
					</Text>
				</label>
				<InlineStack gap="200" align="center" blockAlign="center">
					{settingStatusMarkup}
				</InlineStack>
			</InlineStack>
		</InlineStack>
	) : null;

	const actionMarkup = (
		<Button
			role="switch"
			id={toggleId}
			ariaChecked={enabled ? 'true' : 'false'}
			onClick={handleToggle}
			size="slim"
			loading={isLoading}
			disabled={isLoading}
		>
			{contentStatus}
		</Button>
	);

	const headerMarkup = (
		<Box width="100%">
			<InlineStack
				gap="1200"
				align="space-between"
				blockAlign="start"
				wrap={false}
			>
				{settingTitle}
				{!mdDown ? (
					<Box minWidth="fit-content">
						<InlineStack align="end">{actionMarkup}</InlineStack>
					</Box>
				) : null}
			</InlineStack>
		</Box>
	);

	const descriptionMarkup = (
		<BlockStack gap="400">
			<Text id={descriptionId} variant="bodyMd" as="p" tone="subdued">
				{description}
			</Text>
			{mdDown ? (
				<Box width="100%">
					<InlineStack align="start">{actionMarkup}</InlineStack>
				</Box>
			) : null}
		</BlockStack>
	);

	return (
			<Card>
				<BlockStack gap={{xs: '400', sm: '500'}}>
					<Box width="100%">
						<BlockStack gap={{xs: '200', sm: '400'}}>
							{headerMarkup}
							{descriptionMarkup}
						</BlockStack>
					</Box>
				</BlockStack>
			</Card>
	);
}
