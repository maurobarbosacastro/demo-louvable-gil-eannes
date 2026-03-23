import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {
	CountryFilters,
	CountryInterface,
} from '@app/modules/admin/countries/interfaces/country.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';

@Injectable({
	providedIn: 'root',
})
export class CountriesService {
	endpoint = `${environment.host}${environment.tagpeak}/countries`;
	publicEndpoint = `${environment.host}${environment.tagpeak}/public/countries`;

	private httpClient = inject(HttpClient);

	getCountriesPublic(
		size?: number,
		page?: number,
		sort?: string,
		filters?: CountryFilters,
	): Observable<PaginationInterface<CountryInterface>> {
		let params = new HttpParams();

		if (size) {
			params = params.append('limit', size);
		}
		if (page) {
			params = params.append('page', page);
		}
		if (sort) {
			params = params.append('sort', sort);
		}
		if (filters.name) {
			params = params.append('name', filters.name);
		}
		if (filters.currency) {
			params = params.append('currency', filters.currency);
		}
		if (filters.enabled) {
			params = params.append('enabled', filters.enabled);
		}

		return this.httpClient.get<PaginationInterface<CountryInterface>>(`${this.publicEndpoint}`, {
			params,
		});
	}

	getCountries(
		size?: number,
		page?: number,
		sort?: string,
		filters?: CountryFilters,
	): Observable<PaginationInterface<CountryInterface>> {
		let params = new HttpParams();

		if (size) {
			params = params.append('limit', size);
		}
		if (page) {
			params = params.append('page', page);
		}
		if (sort) {
			params = params.append('sort', sort);
		}
		if (filters?.name) {
			params = params.append('name', filters.name);
		}
		if (filters?.currency) {
			params = params.append('currency', filters.currency);
		}
		if (filters?.enabled) {
			params = params.append('enabled', filters.enabled);
		}

		return this.httpClient.get<PaginationInterface<CountryInterface>>(`${this.endpoint}`, {params});
	}

	getCountry(id: string): Observable<CountryInterface> {
		return this.httpClient.get<CountryInterface>(`${this.endpoint}/${id}`);
	}

	getCountryByCode(code: string): Observable<CountryInterface> {
		return this.httpClient.get<CountryInterface>(`${this.publicEndpoint}/code/${code}`);
	}

	createCountry(body: CountryInterface): Observable<CountryInterface> {
		return this.httpClient.post<CountryInterface>(`${this.endpoint}`, body);
	}

	updateCountry(id: string, body: CountryInterface): Observable<CountryInterface> {
		return this.httpClient.patch<CountryInterface>(`${this.endpoint}/${id}`, body);
	}

	deleteCountry(id: string): Observable<any> {
		return this.httpClient.delete<any>(`${this.endpoint}/${id}`);
	}
}
