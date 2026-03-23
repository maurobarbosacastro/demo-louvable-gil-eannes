import type {ActionFunctionArgs, LoaderFunctionArgs} from '@remix-run/node';
import {redirect} from '@remix-run/node';
import {authenticate} from '../shopify.server';
import {useEffect, useMemo, useState} from 'react';
import {
	Banner,
	BlockStack,
	Box,
	Button,
	Card,
	Checkbox,
	FormLayout,
	InlineGrid,
	InlineStack,
	Link,
	Page,
	Tag,
	Text,
	TextField,
	useBreakpoints,
} from '@shopify/polaris';
import {Form, useLoaderData} from '@remix-run/react';
import {TitleBar, useAppBridge} from '@shopify/app-bridge-react';
import MultiAutocompleteCountries from '../components/MultiAutocompleteCountries';
import type {
	CountryInterface,
	PaginationInterface,
	CategoriesInterface,
} from '../interfaces/common.interface';
import MultiAutocompleteCategories from '../components/MultiAutocompleteCategories';

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
		termsUrl: process.env.TERMS_AND_CONDITION_URL,
	};
};

export const action = async ({request}: ActionFunctionArgs) => {
	try {
		const formBody = await request.formData();

		const body = {
			name: formBody.get('shopName') as string,
			percentage: Number(formBody.get('percentage')),
			returnPeriod: Number(formBody.get('returnPeriod')),
			url: formBody.get('shopUrl') as string,
			shopUuid: formBody.get('shopUuid') as string,
			description: formBody.get('shopDescription') as string,
			countries: JSON.parse(formBody.get('countries') as string),
			categories: JSON.parse(formBody.get('categories') as string),
		};

		const response = await fetch(`${process.env.BE_URL as string}/tagpeak/shopify`, {
			method: 'POST',
			body: JSON.stringify(body),
			headers: {
				'Content-Type': 'application/json',
				'Authorization': formBody.get('tpToken') as string,
			},
		});

		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}
		// const data = await response.json();

		const installationCall = await fetch(`${process.env.BE_URL as string}/shopify/install`, {
			method: 'POST',
			body: JSON.stringify({
				products: JSON.parse((formBody.get('products') as string) || '[]'),
			}),
			// @ts-ignore
			headers: {
				'Content-Type': 'application/json',
				'Authorization': formBody.get('tpToken') as string,
				'X-Shopify-Shop': formBody.get('shop') as string,
				'X-TP-Shopify-Token': formBody.get('accessToken') as string,
			},
		});
		// const installationData = await installationCall.json();
		if (!installationCall.ok) {
			throw new Error(`HTTP error! status: ${installationCall.status}`);
		}

		return redirect('/app');
	} catch (error) {
		console.error('Error in action:', error);
		throw new Response('Error processing request', {status: 500});
	}
};

export default function Welcome() {
	const data = useLoaderData<typeof loader>();
	const shopify = useAppBridge();
	const [isDirty, setIsDirty] = useState(false);
	const [isValid, setIsValid] = useState(true);

	const tpToken: string = useMemo(() => sessionStorage.getItem('tp_accessToken') as string, []);
	const shopUuid: string = useMemo(() => sessionStorage.getItem('shop_uuid') as string, []);

	const [percentage, setPercentage] = useState('10');
	const [days, setDays] = useState('30');
	const [shopName, setShopName] = useState(data.displayName);
	const [shopDescription, setShopDescription] = useState('');
	const [shopUrl, setShopUrl] = useState(data.url);
	const [selectedCountries, setSelectedCountries] = useState<string[]>([]);
	const [selectedCategories, setSelectedCategories] = useState<string[]>([]);
	const [selectedOptionsCountries, setSelectedOptionsCountries] = useState<string[]>([]);
	const [selectedOptionsCategories, setSelectedOptionsCategories] = useState<string[]>([]);
	const [acceptedTerms, setAcceptedTerms] = useState(false);
	const [selectedProducts, setSelectedProducts] = useState<Array<{id: string; title: string}>>([]);
	const {mdDown} = useBreakpoints();

	useEffect(() => {
		const res = data.countries.data.reduce((acc: CountryInterface[], curr: CountryInterface) => {
			if (selectedOptionsCountries.includes(curr.name.toLowerCase())) {
				acc.push(curr);
			}
			return acc;
		}, []);
		setSelectedCountries(res.map((c) => c.abbreviation));
	}, [data.countries.data, selectedOptionsCountries]);

	useEffect(() => {
		const res = data.categories.data.reduce(
			(acc: CategoriesInterface[], curr: CategoriesInterface) => {
				if (selectedOptionsCategories.includes(curr.name.toLowerCase())) {
					acc.push(curr);
				}
				return acc;
			},
			[],
		);
		setSelectedCategories(res.map((c) => c.code));
	}, [data.categories.data, selectedOptionsCategories]);

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

		const isTermsInvalid = !acceptedTerms;

		if (
			isPercentageInvalid ||
			isDaysInvalid ||
			isShopNameInvalid ||
			isShopDescriptionInvalid ||
			isShopUrlInvalid ||
			isCountriesInvalid ||
			isCategoriesInvalid ||
			isTermsInvalid
		) {
			setIsValid(false);
		} else {
			setIsDirty(false);
			setIsValid(true);
		}
	}, [
		percentage,
		days,
		shopName,
		shopDescription,
		shopUrl,
		selectedCountries,
		selectedCategories,
		acceptedTerms,
	]);

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

	const handleProductSelection = async () => {
		try {
			const selection = await shopify.resourcePicker({
				type: 'product',
				multiple: true,
				filter: {
					archived: false,
					draft: false,
					variants: false,
					hidden: false,
				},
			});

			if (selection) {
				setSelectedProducts(
					selection.map((product: any) => ({
						id: product.id,
						title: product.title,
					})),
				);
			}
		} catch (error) {
			// User cancelled selection
			console.log('Product selection cancelled');
		}
	};

	return (
		<Page fullWidth>
			<TitleBar title="TagPeak - Welcome!" />

			<Form method="post">
				<Box paddingBlockEnd="800">
					<BlockStack gap="800">
						{/* Section 1: Commission Settings */}
						<InlineGrid columns={{xs: '1fr', md: '2fr 5fr'}} gap="400">
							<Box>
								<BlockStack gap="200">
									<Text variant="headingMd" as="h2">
										Commission
									</Text>
									<Text variant="bodyMd" as="p" tone="subdued">
										Configure the commission percentage TagPeak receives per order
									</Text>
								</BlockStack>
							</Box>
							<Card>
								<BlockStack gap="400">
									<TextField
										label="Percentage"
										type="number"
										value={percentage}
										onChange={(perc) => setPercentage(perc)}
										prefix="%"
										autoComplete="off"
										min={0}
										max={100}
										step={1}
										name="percentage"
										helpText="Enter a value between 0 and 100"
										error={
											Number(percentage) < 0 || Number(percentage) > 100
												? 'Percentage must be between 0 and 100'
												: false
										}
									/>
								</BlockStack>
							</Card>
						</InlineGrid>

						{/* Section 2: Store Information */}
						<InlineGrid columns={{xs: '1fr', md: '2fr 5fr'}} gap="400">
							<Box>
								<BlockStack gap="200">
									<Text variant="headingMd" as="h2">
										Store details
									</Text>
									<Text variant="bodyMd" as="p" tone="subdued">
										This information will be used as a description of your store on Tagpeak’s
										partner stores directory
									</Text>
								</BlockStack>
							</Box>
							<Card>
								<FormLayout>
									<TextField
										label="Shop name"
										type="text"
										value={shopName}
										onChange={(name) => setShopName(name)}
										autoComplete="off"
										name="shopName"
										error={
											!shopName || shopName.trim().length === 0 ? 'Shop name is required' : false
										}
									/>

									<TextField
										label="Shop description"
										type="text"
										multiline={4}
										maxLength={255}
										value={shopDescription}
										onChange={(desc) => setShopDescription(desc)}
										autoComplete="off"
										name="shopDescription"
										showCharacterCount
										helpText="Maximum 255 characters"
										error={
											!shopDescription || shopDescription.trim().length === 0
												? 'Shop description is required'
												: false
										}
									/>

									<TextField
										label="Shop URL"
										type="text"
										value={shopUrl}
										onChange={(url) => setShopUrl(url)}
										autoComplete="off"
										name="shopUrl"
										error={isShopUrlInvalid(shopUrl)}
									/>
								</FormLayout>
							</Card>
						</InlineGrid>

						{/* Section 3: Returns Policy */}
						<InlineGrid columns={{xs: '1fr', md: '2fr 5fr'}} gap="400">
							<Box>
								<BlockStack gap="200">
									<Text variant="headingMd" as="h2">
										Returns policy
									</Text>
									<Text variant="bodyMd" as="p" tone="subdued">
										After this period, transactions will be automatically validated and the
										commission will be due. We always have a 2 week buffer period to account for any
										delivery delays. If your returns policy is 30 days, the system will only
										consider a transaction valid after 45 days
									</Text>
								</BlockStack>
							</Box>
							<Card>
								<BlockStack gap="400">
									<TextField
										label="Returns/Cancellation Period (days)"
										type="number"
										value={days}
										onChange={(daysVal) => setDays(daysVal)}
										autoComplete="off"
										min={0}
										step={1}
										name="returnPeriod"
										helpText="Specify the number of days for returns or cancellations"
										error={
											!days || Number(days) <= 0
												? 'Specify a valid number of days greater than 0'
												: false
										}
									/>
								</BlockStack>
							</Card>
						</InlineGrid>

						{/* Section 4: Store Classification */}
						<InlineGrid columns={{xs: '1fr', md: '2fr 5fr'}} gap="400">
							<Box>
								<BlockStack gap="200">
									<Text variant="headingMd" as="h2">
										Categories & Countries
									</Text>
									<Text variant="bodyMd" as="p" tone="subdued">
										This information will be used as a description of your store on Tagpeak’s
										partner stores directory.
									</Text>
								</BlockStack>
							</Box>
							<Card>
								<FormLayout>
									<Box paddingBlockStart={mdDown ? '400' : '200'}>
										<MultiAutocompleteCategories
											categories={data.categories.data}
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
											countries={data.countries.data}
											selectedOptions={selectedOptionsCountries}
											setSelectedOptions={setSelectedOptionsCountries}
											height={mdDown ? '300px' : '250px'}
											onChange={() => {
												handleFormChange();
											}}
											emptyError
										/>
									</Box>
								</FormLayout>
							</Card>
						</InlineGrid>

						{/* Section 5: Product Selection */}
						<InlineGrid columns={{xs: '1fr', md: '2fr 5fr'}} gap="400">
							<Box>
								<BlockStack gap="200">
									<Text variant="headingMd" as="h2">
										Products
									</Text>
									<Text variant="bodyMd" as="p" tone="subdued">
										You can select products that will be eligible for the Tagpeak integration. This
										step is optional and can be performed later at Products &gt; Collections
									</Text>
								</BlockStack>
							</Box>
							<Card>
								<BlockStack gap="400">
									<Text variant="bodyMd" as="p" tone="subdued">
										Select products that will be eligible for the Tagpeak integration
									</Text>

									<Button onClick={handleProductSelection}>Select products</Button>

									{selectedProducts.length > 0 && (
										<BlockStack gap="200">
											<Text variant="bodyMd" as="p">
												{selectedProducts.length} product{selectedProducts.length !== 1 ? 's' : ''}{' '}
												selected
											</Text>
											<InlineStack gap="200" wrap>
												{selectedProducts.map((product) => (
													<Tag
														key={product.id}
														onRemove={() => {
															setSelectedProducts((prev) =>
																prev.filter((p) => p.id !== product.id),
															);
														}}
													>
														{product.title}
													</Tag>
												))}
											</InlineStack>
										</BlockStack>
									)}
								</BlockStack>
							</Card>
						</InlineGrid>

						{/* Section 6: Installation Preview */}
						<InlineGrid columns={{xs: '1fr', md: '2fr 5fr'}} gap="400">
							<Box>
								<BlockStack gap="200">
									<Text variant="headingMd" as="h2">
										What happens next
									</Text>
									<Text variant="bodyMd" as="p" tone="subdued">
										Overview of the installation process
									</Text>
								</BlockStack>
							</Box>
							<Card>
								<Banner tone="info">
									<BlockStack gap="200">
										<Text variant="bodyMd" as="p">
											When you complete installation, TagPeak will:
										</Text>
										<BlockStack gap="100">
											<Text as="p">
												✓ Create a discount code for the integration to detect order it can process
											</Text>
											<Text as="p">
												✓ Create a collection named Tagpeak for you to associate products that will
												be affected by the integration
											</Text>
											{selectedProducts.length > 0 && (
												<Text as="p">
													✓ Associate {selectedProducts.length} selected product
													{selectedProducts.length !== 1 ? 's' : ''} to the collection
												</Text>
											)}
										</BlockStack>
									</BlockStack>
								</Banner>
							</Card>
						</InlineGrid>

						{/* Section 7: Agreement */}
						<InlineGrid columns={{xs: '1fr', md: '2fr 5fr'}} gap="400">
							<Box>
								<BlockStack gap="200">
									<Text variant="headingMd" as="h2">
										Terms and conditions
									</Text>
								</BlockStack>
							</Box>
							<Card>
								<BlockStack gap="400">
									<Checkbox
										label={
											<span>
												I accept the{' '}
												<Link url={data.termsUrl} target="_blank">
													Terms and Conditions
												</Link>
											</span>
										}
										checked={acceptedTerms}
										onChange={setAcceptedTerms}
									/>
								</BlockStack>
							</Card>
						</InlineGrid>

						{/* Bottom Action Bar */}
						<InlineGrid columns={{xs: '1fr', md: '2fr 5fr'}} gap="400">
							<Box />
							<Box>
								<Button submit disabled={isDirty || !isValid} variant="primary">
									Complete installation
								</Button>
							</Box>
						</InlineGrid>

						{/* Hidden form fields */}
						<input hidden name="tpToken" onChange={() => {}} value={tpToken} />
						<input hidden name="shopUuid" onChange={() => {}} value={shopUuid} />
						<input hidden name="shop" onChange={() => {}} value={data.shop} />
						<input hidden name="accessToken" onChange={() => {}} value={data.accessToken} />
						<input
							hidden
							name="countries"
							onChange={() => {}}
							value={JSON.stringify(selectedCountries)}
						/>
						<input
							hidden
							name="categories"
							onChange={() => {}}
							value={JSON.stringify(selectedCategories)}
						/>
						<input
							hidden
							name="products"
							onChange={() => {}}
							value={JSON.stringify(selectedProducts.map((p) => p.id))}
						/>
					</BlockStack>
				</Box>
			</Form>
		</Page>
	);
}
