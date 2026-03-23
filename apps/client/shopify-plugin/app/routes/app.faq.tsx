import {useLoaderData} from '@remix-run/react';
import type {LoaderFunctionArgs} from '@remix-run/node';
import {authenticate} from '../shopify.server';
import {TitleBar} from '@shopify/app-bridge-react';
import {
	Page,
	Card,
	BlockStack,
	InlineStack,
	Text,
	Box,
	Icon,
} from '@shopify/polaris';
import {ChevronUpIcon, ChevronDownIcon} from '@shopify/polaris-icons';
import {useState, useCallback} from 'react';
import fs from 'fs/promises';
import path from 'path';
import type {FAQItem, FAQData} from '../interfaces/faq.interface';
import {renderMarkdown} from '../utils/markdown.utils';

export const loader = async ({request}: LoaderFunctionArgs) => {
	await authenticate.admin(request);

	try {
		const faqPath = path.join(process.cwd(), 'public', 'faq.json');
		const faqContent = await fs.readFile(faqPath, 'utf-8');
		const faqData: FAQData = JSON.parse(faqContent);
		return {faqs: faqData.faqs, error: null};
	} catch (error) {
		console.error('Error loading FAQ file:', error);
		return {faqs: [], error: 'Failed to load FAQ data'};
	}
};

export default function FAQPage() {
	const {faqs, error} = useLoaderData<typeof loader>();
	const [openItems, setOpenItems] = useState<Set<number>>(new Set());

	const toggleItem = useCallback((index: number) => {
		setOpenItems((prev) => {
			const newSet = new Set(prev);
			if (newSet.has(index)) {
				newSet.delete(index);
			} else {
				newSet.add(index);
			}
			return newSet;
		});
	}, []);

	if (error) {
		return (
			<Page fullWidth>
				<TitleBar title="FAQ" />
				<div
					style={{
						maxWidth: '800px',
						width: '100%',
						margin: '0 auto',
						padding: '0 16px',
					}}
				>
					<Card>
						<Box padding="400">
							<BlockStack gap="400">
								<Text as="p" tone="critical">
									{error}
								</Text>
							</BlockStack>
						</Box>
					</Card>
				</div>
			</Page>
		);
	}

	return (
		<Page fullWidth>
			<TitleBar title="Frequently Asked Questions" />
			<div
				style={{
					maxWidth: '800px',
					width: '100%',
					margin: '0 auto',
					padding: '0 16px',
				}}
			>
				<BlockStack gap="400">
					{faqs.length === 0 ? (
						<Card>
							<Box padding="400">
								<Text as="p" tone="subdued">
									No FAQs available at the moment.
								</Text>
							</Box>
						</Card>
					) : (
						faqs.map((faq, index) => (
							<Card key={index}>
								<BlockStack gap="200">
									<div
										onClick={() => toggleItem(index)}
										role="button"
										tabIndex={0}
										onKeyDown={(e) => {
											if (e.key === 'Enter' || e.key === ' ') {
												e.preventDefault();
												toggleItem(index);
											}
										}}
										style={{
											cursor: 'pointer',
											padding: '12px 16px',
											borderRadius: '8px',
										}}
									>
										<InlineStack align="space-between" blockAlign="center" gap="400">
											<div style={{flex: '1'}}>
												<Text as="h3" variant="headingSm" breakWord>
													{renderMarkdown((faq as FAQItem).question)}
												</Text>
											</div>
											<Icon
												source={openItems.has(index) ? ChevronUpIcon : ChevronDownIcon}
												tone="subdued"
											/>
										</InlineStack>
									</div>
									{openItems.has(index) && (
										<Box
											paddingInlineStart="400"
											paddingInlineEnd="400"
											paddingBlockEnd="400"
										>
											<Text as="p" variant="bodyMd" tone="subdued" breakWord>
												{renderMarkdown((faq as FAQItem).answer)}
											</Text>
										</Box>
									)}
								</BlockStack>
							</Card>
						))
					)}
				</BlockStack>
			</div>
		</Page>
	);
}
