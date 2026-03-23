import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable, of} from 'rxjs';
import {
	BulkEditRewardDto,
	CashbackHistory,
	CashbackResponse,
	CashbackTableResponse,
	CreateRewardDTO,
	ManageCashbackSearchParams,
	RewardResponse,
	RewardUpdateDTO,
	TransactionUpdateDTO,
} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import _ from 'lodash';
import {AppConstants} from '@app/app.constants';

@Injectable({
	providedIn: 'root',
})
export class CashbackService {
	endpoint: string = `${environment.host}${environment.tagpeak}/transaction`;
	endpointReward: string = `${environment.host}${environment.tagpeak}/reward`;

	private http: HttpClient = inject(HttpClient);

	getCashback(
		page: number = 1,
		size: number = AppConstants.PAGE_SIZE,
		sort: string,
		filtersBody: Partial<ManageCashbackSearchParams>,
	): Observable<PaginationInterface<CashbackTableResponse>> {
		let params: HttpParams = new HttpParams().append('page', page).append('limit', size);

		if (sort) {
			params = params.append('sort', sort);
		}

		if (filtersBody?.state) {
			params = params.append('state', filtersBody.state);
		}

		if (filtersBody?.storeUuid) {
			params = params.append('storeUuid', filtersBody.storeUuid);
		}

		if (filtersBody?.userUuid) {
			params = params.append('userUuid', filtersBody.userUuid);
		}

		if (filtersBody?.storeVisitUuid) {
			params = params.append('storeVisitUuid', filtersBody.storeVisitUuid);
		}

		if (filtersBody?.startDate) {
			params = params.append('startDate', filtersBody.startDate);
		}

		if (filtersBody?.endDate) {
			params = params.append('endDate', filtersBody.endDate);
		}

		return this.http.get<PaginationInterface<CashbackTableResponse>>(this.endpoint, {params});
	}

	getMyCashback(
		page: number = 1,
		size: number = AppConstants.PAGE_SIZE,
		sort: string,
		filtersBody: Partial<ManageCashbackSearchParams>,
	): Observable<PaginationInterface<CashbackTableResponse>> {
		let params: HttpParams = new HttpParams({fromObject: filtersBody})
			.append('page', page)
			.append('limit', size);

		if (sort) {
			params = params.append('sort', sort);
		}

		return this.http.get<PaginationInterface<CashbackTableResponse>>(`${this.endpoint}/me`, {
			params,
		});
	}

	getCashbackById(uuid: string): Observable<CashbackResponse> {
		return this.http.get<CashbackResponse>(`${this.endpoint}/${uuid}`);
	}

	getCashbackReward(uuid: string, admin: boolean = false): Observable<RewardResponse> {
		let url = `${this.endpoint}/${uuid}/reward`;
		if (admin) {
			url = `${this.endpoint}/${uuid}/reward?admin=true`;
		}
		return this.http.get<RewardResponse>(url);
	}

	updateCashback(uuid: string, body: Partial<TransactionUpdateDTO>): Observable<CashbackResponse> {
		if (_.isEmpty(body)) {
			return of({} as CashbackResponse);
		}
		return this.http.patch<CashbackResponse>(`${this.endpoint}/${uuid}`, body);
	}

	createCashbackReward(body: CreateRewardDTO): Observable<any> {
		return this.http.post<RewardResponse>(`${this.endpointReward}`, body);
	}

	updateCashbackReward(uuid: string, body: Partial<RewardUpdateDTO>): Observable<RewardResponse> {
		if (_.isEmpty(body)) {
			return of({} as RewardResponse);
		}
		return this.http.patch<RewardResponse>(`${this.endpointReward}/${uuid}`, body);
	}

	updateCashbackBulkTransaction(body: Partial<TransactionUpdateDTO>): Observable<void> {
		return this.http.patch<void>(`${this.endpoint}/bulk/edit`, body);
	}

	updateCashbackBulkReward(body: Partial<BulkEditRewardDto>): Observable<void> {
		return this.http.patch<void>(`${this.endpointReward}/bulk/edit`, body);
	}

	createCashbackRewardBulk(body: Partial<BulkEditRewardDto>): Observable<any> {
		return this.http.post<RewardResponse>(`${this.endpointReward}/bulk/create`, body);
	}

	deleteCashback(uuid: string): Observable<void> {
		return this.http.delete<void>(`${this.endpoint}/${uuid}`);
	}

	getCashbackHistory(rewardUuid: string): Observable<CashbackHistory[]> {
		return this.http.get<CashbackHistory[]>(`${this.endpointReward}/${rewardUuid}/history/graph`);
	}

	verifyReward(rewardUuid: string): Observable<any> {
		return this.http.post<any>(`${this.endpointReward}/${rewardUuid}/verify`, {});
	}
}
