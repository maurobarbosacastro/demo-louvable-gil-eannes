import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {StoreInterface} from '@app/modules/client/store-visits/models/store.interface';
import {
	StoreVisitsInterface,
	StoreVisitsSearchParams,
} from '@app/modules/client/store-visits/models/store-visits.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';

@Injectable({
	providedIn: 'root',
})
export class StoreVisitsService {
	endpoint: string = `${environment.host}${environment.tagpeak}/store-visit`;

	private http: HttpClient = inject(HttpClient);

	// Get all store visits of all users
	getAllStoreVisits(
		page: number,
		size: number,
		sort?: string,
		filterBody?: Partial<StoreVisitsSearchParams>,
	): Observable<PaginationInterface<StoreVisitsInterface>> {
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
		if (filterBody?.store) {
			params = params.append('storeUuid', filterBody.store);
		}
		if (filterBody?.name) {
			params = params.append('name', filterBody.name);
		}
		if (filterBody?.endDate) {
			params = params.append('dateTo', filterBody.endDate);
		}
		if (filterBody?.startDate) {
			params = params.append('dateFrom', filterBody.startDate);
		}
		if (filterBody?.reference) {
			params = params.append('reference', filterBody.reference);
		}
		return this.http.get<PaginationInterface<StoreVisitsInterface>>(`${this.endpoint}`, {
			params,
		});
	}

	// Get all store-visits of an user
	getStoreVisitsByUser(
		page: number,
		size: number,
		sort?: string,
		filterBody?: Partial<StoreVisitsSearchParams>,
	): Observable<PaginationInterface<StoreVisitsInterface>> {
		let params = new HttpParams().append('limit', size).append('page', page);
		if (sort) {
			params = params.append('sort', sort);
		}
		if (filterBody) {
			if (filterBody.store) {
				params = params.append('store', filterBody.store);
			}
			if (filterBody.endDate) {
				params = params.append('endDate', filterBody.endDate);
			}
			if (filterBody.startDate) {
				params = params.append('startDate', filterBody.startDate);
			}
		}

		return this.http.get<PaginationInterface<StoreVisitsInterface>>(`${this.endpoint}/user`, {
			params,
		});
	}

	validateReference(reference: string): Observable<boolean> {
		const url = this.endpoint + '/reference/' + reference;
		return this.http.post<boolean>(url, null);
	}

	// Get all stores that were visited by user
	getStoresVisitedByUser(
		page: number,
		size: number,
		sort: string,
		filterBody: Partial<StoreVisitsSearchParams>,
	): Observable<PaginationInterface<StoreInterface>> {
		let params = new HttpParams({fromObject: filterBody});
		if (size) {
			params = params.append('limit', size);
		}
		if (page) {
			params = params.append('page', page);
		}
		if (sort) {
			params = params.append('sort', sort);
		}
		return this.http.get<PaginationInterface<StoreInterface>>(`${this.endpoint}/stores`, {
			params,
		});
	}

	getStoreVisitsAdmin(
		page: number,
		size: number,
		sort: string,
		filtersBody: Partial<StoreVisitsSearchParams>,
	): Observable<PaginationInterface<StoreVisitsInterface>> {
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
		if (filtersBody?.reference) {
			params = params.append('reference', filtersBody.reference);
		}
		return this.http.get<PaginationInterface<StoreVisitsInterface>>(`${this.endpoint}/admin`, {
			params,
		});
	}
}
