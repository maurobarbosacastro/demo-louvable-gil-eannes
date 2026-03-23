import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {HttpClient, HttpParams} from '@angular/common/http';
import {map, Observable} from 'rxjs';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {
	AdminStoreInterface,
	CreateStoreInterface,
	LogoStore,
	StoreApprovalRequest,
} from '@app/modules/admin/stores/interfaces/store.interface';
import {AppConstants} from '@app/app.constants';

@Injectable({
	providedIn: 'root',
})
export class StoresService {
	endpoint = `${environment.host}${environment.tagpeak}/store`;
	endpointAux = `${environment.host}${environment.tagpeak}/aux`;

	private httpClient = inject(HttpClient);

	getAllStores(
		size?: number,
		page?: number,
		sort: string = 'name asc',
		name?: string,
	): Observable<PaginationInterface<AdminStoreInterface>> {
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
		if (name) {
			params = params.append('name', name);
		}

		return this.httpClient.get<PaginationInterface<AdminStoreInterface>>(`${this.endpoint}/admin`, {
			params,
		});
	}

	getStore(id: string): Observable<AdminStoreInterface> {
		return this.httpClient.get<AdminStoreInterface>(`${this.endpoint}/${id}`);
	}

	createStore(body: CreateStoreInterface): Observable<AdminStoreInterface> {
		return this.httpClient.post<AdminStoreInterface>(`${this.endpoint}`, body);
	}

	updateStore(id: string, body: Partial<CreateStoreInterface>): Observable<AdminStoreInterface> {
		return this.httpClient.patch<AdminStoreInterface>(`${this.endpoint}/${id}`, body);
	}

	deleteStore(id: string): Observable<any> {
		return this.httpClient.delete<any>(`${this.endpoint}/${id}`);
	}

	uploadCSV(file: File): Observable<any> {
		const formData = new FormData();
		formData.append('file', file);

		return this.httpClient.post<any>(`${this.endpoint}/upload-excel`, formData);
	}

	uploadLogo(id: string, file: File): Observable<AdminStoreInterface> {
		const formData = new FormData();
		formData.append('file', file);
		return this.httpClient.post<AdminStoreInterface>(`${this.endpoint}/${id}/logo`, formData);
	}

	deleteLogo(id: string): Observable<void> {
		return this.httpClient.delete<void>(`${this.endpoint}/${id}/logo`);
	}

	uploadBanner(id: string, file: File): Observable<AdminStoreInterface> {
		const formData = new FormData();
		formData.append('file', file);
		return this.httpClient.post<AdminStoreInterface>(`${this.endpoint}/${id}/banner`, formData);
	}

	deleteBanner(id: string): Observable<void> {
		return this.httpClient.delete<void>(`${this.endpoint}/${id}/banner`);
	}

	getStoreLogoFromDomain(domain: string): Observable<LogoStore> {
		const param = new HttpParams().set('name', domain);
		return this.httpClient.get<LogoStore>(`${this.endpointAux}/logo`, {params: param});
	}

	exportStoresCSV(fileName: string, sort: string): Observable<File> {
		let params: HttpParams = new HttpParams().append('fileName', fileName);
		params = params.append('sort', sort);
		return this.httpClient
			.get(`${this.endpoint}/export-csv`, {params: params, responseType: 'blob'})
			.pipe(map((response: Blob) => new File([response], fileName, {type: response.type})));
	}

	getStoresForApprovals(
		size: number = AppConstants.PAGE_SIZE,
		page: number = 0,
		sort: string = 'createdAt asc',
	): Observable<PaginationInterface<StoreApprovalRequest>> {
		let params = new HttpParams().append('limit', size).append('page', page).append('sort', sort);

		return this.httpClient.get<PaginationInterface<StoreApprovalRequest>>(
			`${this.endpoint}/approvals`,
			{params},
		);
	}

	positionIsUnique(position: number): Observable<boolean> {
		return this.httpClient.get<boolean>(`${this.endpoint}/position/${position}`);
	}
}
