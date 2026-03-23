import {useEffect, useState} from 'react';
import {
	BlockStack,
	Box,
	Card,
	Divider,
	FormLayout,
	Text,
	TextField,
	useBreakpoints,
} from '@shopify/polaris';
import MultiAutocompleteCountries from '../components/MultiAutocompleteCountries';
import MultiAutocompleteCategories from '../components/MultiAutocompleteCategories';
import type {
	CategoriesInterface,
	CountryInterface,
	PaginationInterface,
	StoreInterface,
} from '../interfaces/common.interface';
import {Form} from '@remix-run/react';
import type {StoreUpdateDTO} from '../interfaces/dtos.interface';

type ShopDetailsProps = {
	displayName: string;
	url: string;
	countries: PaginationInterface<CountryInterface>;
	categories: PaginationInterface<CategoriesInterface>;
	shop: string;
	accessToken: string;
	currentStore: StoreInterface;
	onSave?: (data: StoreUpdateDTO) => void;
	isSaving?: boolean;
};

export default function ShopDetails(props: ShopDetailsProps) {
	// At the beginning of your ShopDetails component, add these state variables
	const [isDirty, setIsDirty] = useState(false);
	const [isValid, setIsValid] = useState(true);
	const {smDown, mdDown} = useBreakpoints();

	const [percentage, setPercentage] = useState(props.currentStore.cashbackValue.toString());
	const [days, setDays] = useState(props.currentStore.averageRewardActivationTime.split(' ')[0]);
	const [shopName, setShopName] = useState(props.currentStore.name);
	const [shopDescription, setShopDescription] = useState(props.currentStore.shortDescription);
	const [shopUrl, setShopUrl] = useState(props.currentStore.storeUrl);
	const [selectedCountries, setSelectedCountries] = useState<string[]>(props.currentStore.country);
	const [selectedCategories, setSelectedCategories] = useState<string[]>(
		props.currentStore.category,
	);
	const [selectedOptionsCountries, setSelectedOptionsCountries] = useState<string[]>([]);
	const [selectedOptionsCategories, setSelectedOptionsCategories] = useState<string[]>([]);

	useEffect(() => {
		const setInitialData = () => {
			//Set initial data
			setSelectedCategories(props.currentStore.category);
			setSelectedCountries(props.currentStore.country);
			setShopDescription(props.currentStore.shortDescription);
			setShopName(props.currentStore.name);
			setShopUrl(props.currentStore.storeUrl);
			setPercentage(props.currentStore.cashbackValue.toString());
			setDays(props.currentStore.averageRewardActivationTime.split(' ')[0]);

			//Set country as preselected
			const preSelectedCountries = props.countries.data.reduce(
				(acc: CountryInterface[], curr: CountryInterface) => {
					if (props.currentStore.country.includes(curr.abbreviation)) {
						acc.push(curr);
					}
					return acc;
				},
				[],
			);
			setSelectedOptionsCountries(preSelectedCountries.map((r) => r.name.toLowerCase()));

			//Set categories as preselected
			const preSelectedCategories = props.categories.data.reduce(
				(acc: CategoriesInterface[], curr: CategoriesInterface) => {
					if (props.currentStore.category.includes(curr.code.toLowerCase())) {
						acc.push(curr);
					}
					return acc;
				},
				[],
			);
			setSelectedOptionsCategories(preSelectedCategories.map((r) => r.name.toLowerCase()));
		};
		setInitialData();

		// Handle the save bar setup
		const saveBarElement = document.getElementById('shop-details-save-bar');
		if (saveBarElement) {
			// @ts-ignore
			saveBarElement.hide();
		}

		if (saveBarElement) {
			const discardButton = document.getElementById('shop-details-discard-button');
			const handleDiscard = () => {
				setInitialData();
				setIsDirty(false);
				// @ts-ignore
				saveBarElement.hide();
			};

			discardButton!.addEventListener('click', handleDiscard);
			return () => {
				discardButton!.removeEventListener('click', handleDiscard);
			};
		}
	}, []);
	useEffect(() => {
		const res = props.countries.data.reduce((acc: CountryInterface[], curr: CountryInterface) => {
			if (selectedOptionsCountries.includes(curr.name.toLowerCase())) {
				acc.push(curr);
			}
			return acc;
		}, []);
		setSelectedCountries(res.map((c) => c.abbreviation));
	}, [props.countries.data, selectedOptionsCountries]);
	useEffect(() => {
		const res = props.categories.data.reduce(
			(acc: CategoriesInterface[], curr: CategoriesInterface) => {
				if (selectedOptionsCategories.includes(curr.name.toLowerCase())) {
					acc.push(curr);
				}
				return acc;
			},
			[],
		);
		setSelectedCategories(res.map((c) => c.code));
	}, [props.categories.data, selectedOptionsCategories]);
	useEffect(() => {
		const saveBarElement = document.getElementById('shop-details-save-bar');
		if (saveBarElement) {
			if (isDirty) {
				// @ts-ignore
				saveBarElement.show();
			} else {
				// @ts-ignore
				saveBarElement.hide();
			}
		}
	}, [isDirty]);

	useEffect(() => {
		const isPercentageInvalid =
			!percentage ||
			Number(percentage) < 0 ||
			Number(percentage) > 100 ||
			percentage.trim().length === 0;

		const isDaysInvalid = !days || Number(days) <= 0;

		const isShopNameInvalid = !shopName || shopName.trim().length === 0;

		const isShopDescriptionInvalid = !shopDescription || shopDescription.trim().length === 0;

		const isShopUrlInvalid = !shopUrl || shopUrl.trim().length === 0 || !isURLValid(shopUrl);

		const isCountriesInvalid = !selectedCountries || selectedCountries.length === 0;

		const isCategoriesInvalid = !selectedCategories || selectedCategories.length === 0;

		if (
			isPercentageInvalid ||
			isDaysInvalid ||
			isShopNameInvalid ||
			isShopDescriptionInvalid ||
			isShopUrlInvalid ||
			isCountriesInvalid ||
			isCategoriesInvalid
		) {
			setIsValid(false);
		} else {
			setIsValid(true);
		}
	}, [percentage, days, shopName, shopDescription, shopUrl, selectedCountries, selectedCategories]);

	// Modify your form change handlers to set isDirty
	const handleFormChange = () => {
		setIsDirty(true);
	};

	const isURLValid = (url: string) => {
		const urlRegex = /^(https?:\/\/)([\w-]+\.)+[\w-]+(\/[\w-./?%&=]*)?$/;
		return urlRegex.test(url);
	};

	const isShopUrlInvalid = (url: string) => {
		if (!url || url.trim().length === 0) {
			return 'Shop URL is required';
		}

		if (!isURLValid(url)) {
			return 'Add a valid URL';
		}

		return false;
	};

	return (
		<Card>
			<BlockStack gap="500">
				<BlockStack gap="200">
					<Text variant="headingMd" as="h2">
						Store details
					</Text>
					<Text variant="bodyMd" as="p" tone="subdued">
						Tagpeak will use this information to process rewards and advertise your store.
					</Text>
				</BlockStack>

				<Divider />

				<ui-save-bar id="shop-details-save-bar">
					<button
						variant="primary"
						id="shop-details-save-button"
						disabled={!isValid || props.isSaving}
						onClick={() => {
							if (props.onSave) {
								props.onSave({
									name: shopName,
									cashbackValue: Number(percentage),
									averageRewardActivationTime: days,
									storeUrl: shopUrl,
									shortDescription: shopDescription,
									country: selectedCountries,
									category: selectedCategories,
								});
							}
							setIsDirty(false);
							const saveBarElement = document.getElementById('shop-details-save-bar');
							if (saveBarElement) {
								// @ts-ignore
								saveBarElement.hide();
							}
						}}
					>
						Save
					</button>
					<button id="shop-details-discard-button">Discard</button>
				</ui-save-bar>

				<Form
					onChange={handleFormChange}
					onSubmit={(event) => {
						console.log('onSubmit', event);
					}}
				>
					<FormLayout>
						<BlockStack gap={smDown ? '400' : '500'}>
							<TextField
								label="What percentage of commission will Tagpeak receive for each order?"
								type="number"
								value={percentage}
								onChange={(perc) => setPercentage(perc)}
								prefix="%"
								autoComplete="off"
								min={0}
								max={100}
								step={1}
								name={'percentage'}
								error={
									Number(percentage) < 0 || Number(percentage) > 100
										? 'Percentage must be between 0 and 100'
										: false
								}
							/>

							<TextField
								label="Shop name"
								type="text"
								value={shopName}
								onChange={(name) => {
									setShopName(name);
								}}
								autoComplete="off"
								name={'shopName'}
								error={!shopName || shopName.trim().length === 0 ? 'Shop name is required' : false}
							/>

							<TextField
								label="Shop description (max 255 characters)"
								type="text"
								multiline={4}
								maxLength={255}
								value={shopDescription}
								onChange={(desc) => setShopDescription(desc)}
								autoComplete="off"
								name={'shopDescription'}
								showCharacterCount
								error={
									!shopDescription || shopDescription.trim().length === 0
										? 'Shop description is required'
										: false
								}
								helpText={
									<span>
										This information will be used as a description of your store on Tagpeak’s
										partner stores directory.
									</span>
								}
							/>

							<TextField
								label="Shop URL"
								type="text"
								value={shopUrl}
								onChange={(url) => setShopUrl(url)}
								autoComplete="off"
								name={'shopUrl'}
								error={isShopUrlInvalid(shopUrl)}
							/>

							<TextField
								label="Returns/Cancellation Period (No. of Days)"
								type="number"
								value={days}
								onChange={(days) => setDays(days)}
								autoComplete="off"
								min={0}
								step={1}
								name={'returnPeriod'}
								error={
									!days || Number(days) <= 0
										? 'Specify a valid number of days greater than 0'
										: false
								}
								helpText={
									<span>
										After this period, transactions will be automatically validated and the
										commission will be due. We always have a 2 week buffer period to account for any
										delivery delays. If your returns policy is 30 days, the system will only
										consider a transaction valid after 45 days
									</span>
								}
							/>

							<Box paddingBlockStart={mdDown ? '400' : '200'}>
								<MultiAutocompleteCategories
									categories={props.categories.data}
									selectedOptions={selectedOptionsCategories}
									setSelectedOptions={setSelectedOptionsCategories}
									height={mdDown ? '140px' : '100px'}
									onChange={() => {
										handleFormChange();
									}}
                  emptyError
								/>
							</Box>

							<Box paddingBlockStart={mdDown ? '400' : '200'}>
								<MultiAutocompleteCountries
									countries={props.countries.data}
									selectedOptions={selectedOptionsCountries}
									setSelectedOptions={setSelectedOptionsCountries}
									height={mdDown ? '300px' : '250px'}
									onChange={() => {
										handleFormChange();
									}}
									emptyError
								/>
							</Box>
						</BlockStack>
					</FormLayout>
				</Form>
			</BlockStack>
		</Card>
	);
}
