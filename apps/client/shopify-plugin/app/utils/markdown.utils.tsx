import {Link as RemixLink} from '@remix-run/react';
import {Text} from '@shopify/polaris';
import type {ParsedSegment} from '../interfaces/faq.interface';

/**
 * Checks if a URL is external (contains http or www)
 */
function isExternalUrl(url: string): boolean {
	return url.includes('http') || url.includes('www');
}

/**
 * Checks if a URL belongs to the tagpeak domain
 */
function isTagPeakDomain(url: string): boolean {
	try {
		// Handle URLs without protocol
		const urlToCheck = url.startsWith('http') ? url : `https://${url}`;
		const urlObj = new URL(urlToCheck);
		const hostname = urlObj.hostname.toLowerCase();

		// Check if hostname contains tagpeak
		return hostname.includes('tagpeak');
	} catch {
		// If URL parsing fails, assume it's not a tagpeak domain
		return false;
	}
}

export function parseMarkdown(text: string): ParsedSegment[] {
	const segments: ParsedSegment[] = [];
	let remaining = text;

	while (remaining.length > 0) {
		// Find positions of both patterns
		const boldMatch = remaining.match(/\*\*(.+?)\*\*/);
		const linkMatch = remaining.match(/\[(.+?)\]\((.+?)\)/);

		const boldIndex = boldMatch ? remaining.indexOf(boldMatch[0]) : -1;
		const linkIndex = linkMatch ? remaining.indexOf(linkMatch[0]) : -1;

		// Determine which pattern comes first
		if (boldIndex !== -1 && (linkIndex === -1 || boldIndex < linkIndex)) {
			// Bold comes first
			if (boldIndex > 0) {
				segments.push({type: 'text', content: remaining.substring(0, boldIndex)});
			}
			segments.push({type: 'bold', content: boldMatch![1]});
			remaining = remaining.substring(boldIndex + boldMatch![0].length);
		} else if (linkIndex !== -1) {
			// Link comes first
			if (linkIndex > 0) {
				segments.push({type: 'text', content: remaining.substring(0, linkIndex)});
			}

			const linkText = linkMatch![1];
			const linkUrl = linkMatch![2];
			const isExternal = isExternalUrl(linkUrl);

			// Validate link according to rules
			if (isExternal) {
				// External links must be to tagpeak domain
				if (isTagPeakDomain(linkUrl)) {
					segments.push({
						type: 'link',
						content: linkText,
						url: linkUrl,
						isExternal: true,
					});
				} else {
					// External non-tagpeak links render as plain text
					segments.push({type: 'invalid-link', content: linkText});
				}
			} else {
				// Internal links (relative paths)
				segments.push({
					type: 'link',
					content: linkText,
					url: linkUrl,
					isExternal: false,
				});
			}

			remaining = remaining.substring(linkIndex + linkMatch![0].length);
		} else {
			// No more patterns found, add remaining as text
			segments.push({type: 'text', content: remaining});
			break;
		}
	}

	return segments;
}

export function renderSegments(segments: ParsedSegment[]) {
	return segments.map((segment, index) => {
		switch (segment.type) {
			case 'bold':
				return (
					<Text key={index} as="span" fontWeight="bold">
						{segment.content}
					</Text>
				);
			case 'link':
				// Check if it's an external link
				if (segment.isExternal) {
					// External tagpeak link - open in new tab
					return (
						<a
							key={index}
							href={segment.url!}
							target="_blank"
							rel="noopener noreferrer"
							style={{
								color: '#2c6ecb',
								textDecoration: 'none',
							}}
							onMouseEnter={(e) => {
								e.currentTarget.style.textDecoration = 'underline';
							}}
							onMouseLeave={(e) => {
								e.currentTarget.style.textDecoration = 'none';
							}}
						>
							{segment.content}
						</a>
					);
				} else {
					// Internal link - use RemixLink for Shopify app navigation
					return (
						<RemixLink
							key={index}
							to={segment.url!}
							style={{
								color: '#2c6ecb',
								textDecoration: 'none',
							}}
							onMouseEnter={(e) => {
								e.currentTarget.style.textDecoration = 'underline';
							}}
							onMouseLeave={(e) => {
								e.currentTarget.style.textDecoration = 'none';
							}}
						>
							{segment.content}
						</RemixLink>
					);
				}
			case 'invalid-link':
				// Invalid link - render as plain text
				return <span key={index}>{segment.content}</span>;
			default:
				return <span key={index}>{segment.content}</span>;
		}
	});
}

export function renderMarkdown(text: string) {
	const segments = parseMarkdown(text);
	return renderSegments(segments);
}
