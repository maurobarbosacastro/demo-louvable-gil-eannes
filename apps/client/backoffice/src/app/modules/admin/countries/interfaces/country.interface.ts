export interface CountryInterface {
	uuid: string;
	abbreviation: string;
	currency: string;
	flag: string;
	name: string;
	enabled: false;
}

export interface CountryFilters {
	name?: string;
	enabled?: boolean;
	currency?: string;
}
