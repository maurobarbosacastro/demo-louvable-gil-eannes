export interface MenuItem {
	value: string;
	label: string;
	router: string[];
	target?: '_self' | '_blank';
	icon?: string;
}
