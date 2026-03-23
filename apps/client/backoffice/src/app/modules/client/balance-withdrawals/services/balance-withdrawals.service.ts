import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {
	BalanceInterface,
	WithdrawalsInterface,
	WithdrawalsSearchParams,
} from '@app/modules/client/balance-withdrawals/models/withdrawals.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';

@Injectable({
	providedIn: 'root',
})
export class BalanceWithdrawalsService {
	endpoint: string = `${environment.host}${environment.tagpeak}/withdrawal`;

	private http: HttpClient = inject(HttpClient);

	getMyWithdrawals(
		page: number,
		size: number,
		sort: string = 'createdAt desc',
		filtersBody?: Partial<WithdrawalsSearchParams>,
	): Observable<PaginationInterface<WithdrawalsInterface>> {
		let params: HttpParams = new HttpParams({fromObject: filtersBody})
			.append('page', page)
			.append('limit', size)
			.append('sort', sort);
		return this.http.get<PaginationInterface<WithdrawalsInterface>>(`${this.endpoint}/me`, {
			params,
		});
	}

	getBalance(): Observable<BalanceInterface> {
		return this.http.get<BalanceInterface>(`${this.endpoint}/me/stats`);
	}

	createNewWithdrawal(): Observable<any> {
		return this.http.post<any>(this.endpoint, {});
	}
}
