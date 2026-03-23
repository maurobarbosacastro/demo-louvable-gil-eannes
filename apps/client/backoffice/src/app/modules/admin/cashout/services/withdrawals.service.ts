import {inject, Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '@environments/environment';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {
	FiltersSearchDTO,
	UpdateStateRequestDTO,
	WithdrawalsRequestInterface,
} from '@app/modules/admin/cashout/models/withdrawals.interface';

@Injectable({
	providedIn: 'root',
})
export class WithdrawalsService {
	private readonly _endpoint: string = `${environment.host}${environment.tagpeak}/withdrawal`;
	private readonly _http: HttpClient = inject(HttpClient);

	getAllWithdrawalsRequest(
		size?: number,
		page?: number,
		sort?: string,
		filters?: FiltersSearchDTO,
	): Observable<PaginationInterface<Partial<WithdrawalsRequestInterface>>> {
		let params: HttpParams = new HttpParams();
		if (size) {
			params = params.append('limit', size);
		}
		if (page) {
			params = params.append('page', page);
		}
		if (sort) {
			params = params.append('sort', sort);
		}
		if (filters.state) {
			params = params.append('state', filters.state);
		}
		if (filters.endDate) {
			params = params.append('endDate', filters.endDate);
		}
		if (filters.startDate) {
			params = params.append('startDate', filters.startDate);
		}

		return this._http.get<PaginationInterface<Partial<WithdrawalsRequestInterface>>>(
			`${this._endpoint}`,
			{params},
		);
	}

	getWithdrawalRequest(uuid: string): Observable<WithdrawalsRequestInterface> {
		return this._http.get<WithdrawalsRequestInterface>(`${this._endpoint}/${uuid}`);
	}

	updateStateRequest(
		body: UpdateStateRequestDTO,
		id: string,
	): Observable<WithdrawalsRequestInterface> {
		return this._http.patch<WithdrawalsRequestInterface>(`${this._endpoint}/${id}`, body);
	}

	bulkUpdateStateRequest(body: any): Observable<WithdrawalsRequestInterface> {
		return this._http.patch<WithdrawalsRequestInterface>(`${this._endpoint}/bulk`, body);
	}

	deleteWithdrawalRequest(id: string): Observable<void> {
		return this._http.delete<void>(`${this._endpoint}/${id}`);
	}
}
