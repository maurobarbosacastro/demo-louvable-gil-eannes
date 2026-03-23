import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {HttpClient, HttpParams} from '@angular/common/http';
import {AppConstants} from '@app/app.constants';
import {Observable} from 'rxjs';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {PartnersInterface} from '@app/modules/admin/partners/interfaces/partners.interface';

@Injectable({
	providedIn: 'root',
})
export class PartnersService {
	private httpClient = inject(HttpClient);
	endpoint = `${environment.host}${environment.tagpeak}/partner`;

	getPartners(
		size: number = AppConstants.PAGE_SIZE,
		page: number = 1,
		sort: string,
		filters?: string,
	): Observable<PaginationInterface<PartnersInterface>> {
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
		if (filters) {
			params = params.append('filters', filters);
		}

		return this.httpClient.get<PaginationInterface<PartnersInterface>>(`${this.endpoint}`, {
			params,
		});
	}

	createPartner(body: PartnersInterface): Observable<PartnersInterface> {
		return this.httpClient.post<PartnersInterface>(`${this.endpoint}`, body);
	}

	getPartner(id: string): Observable<PartnersInterface> {
		return this.httpClient.get<PartnersInterface>(`${this.endpoint}/${id}`);
	}

	updatePartner(id: string, body: PartnersInterface): Observable<PartnersInterface> {
		return this.httpClient.patch<PartnersInterface>(`${this.endpoint}/${id}`, body);
	}

	deletePartner(id: string): Observable<void> {
		return this.httpClient.delete<any>(`${this.endpoint}/${id}`);
	}
}
