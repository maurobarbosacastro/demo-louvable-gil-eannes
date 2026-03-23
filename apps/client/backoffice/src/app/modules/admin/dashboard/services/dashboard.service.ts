import {HttpClient, HttpParams} from '@angular/common/http';
import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {Observable, of} from 'rxjs';
import {
	DashboardInterface,
	RewardByCurrenciesInterface,
	StatisticsByMonthInterface,
	TransactionsStatusInterface,
} from '@app/modules/admin/dashboard/models/dashboard.interface';

@Injectable({
	providedIn: 'root',
})
export class DashboardService {
	endpoint: string = `${environment.host}${environment.tagpeak}/dashboard`;

	private readonly _http: HttpClient = inject(HttpClient);

	getValuesDashboard(): Observable<DashboardInterface> {
		if (!environment.features.dashboard) {
			return of(null);
		}
		return this._http.get<DashboardInterface>(`${this.endpoint}`);
	}

	getStatisticsDashboard(filter: number): Observable<StatisticsByMonthInterface> {
		if (!environment.features.dashboard) {
			return of(null);
		}
		const params: HttpParams = new HttpParams().append('year', filter);
		return this._http.get<StatisticsByMonthInterface>(`${this.endpoint}/statistics`, {params});
	}

	getTransactionsDashboard(): Observable<TransactionsStatusInterface> {
		if (!environment.features.dashboard) {
			return of(null);
		}
		return this._http.get<TransactionsStatusInterface>(`${this.endpoint}/transactions`);
	}

	getRewardCountByCurrencies(): Observable<RewardByCurrenciesInterface> {
		return this._http.get<any>(`${this.endpoint}/rewards/currencies/count`);
	}
}
