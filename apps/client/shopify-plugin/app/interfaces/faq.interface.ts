export interface FAQItem {
	question: string;
	answer: string;
}

export interface FAQData {
	faqs: FAQItem[];
}

export interface ParsedSegment {
	type: 'text' | 'bold' | 'link' | 'invalid-link';
	content: string;
	url?: string;
	isExternal?: boolean;
}
