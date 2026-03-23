export interface BugReportRequest {
	name: string;
	type: string;
	email: string;
	description: string;
	attachment?: BugReportAttachment;
}

export interface BugReportAttachment {
	filename: string;
	data: string;
	mimeType: string;
}
