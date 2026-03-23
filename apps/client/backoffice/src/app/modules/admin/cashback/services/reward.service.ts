import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {CashbackHistory} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {AppConstants} from '@app/app.constants';
import {Reward} from '@app/modules/client/rewards/models/reward.interface';

@Injectable({
	providedIn: 'root',
})
export class RewardService {
	endpoint: string = `${environment.host}${environment.tagpeak}/reward`;

	private http: HttpClient = inject(HttpClient);

	getRewardHistory(
		size: number = AppConstants.PAGE_SIZE,
		page: number = 1,
		sort: string,
		rewardUuid: string,
	): Observable<PaginationInterface<CashbackHistory>> {
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

		return this.http.get<PaginationInterface<CashbackHistory>>(
			`${this.endpoint}/${rewardUuid}/history`,
			{params},
		);
	}

	stopReward(body: {id: string; status: string}): Observable<Reward> {
		return this.http.patch<Reward>(`${this.endpoint}/${body.id}/stop`, body);
	}

	getLiveRewardsValue(userUuid: string): Observable<number> {
		return this.http.get<number>(`${this.endpoint}/${userUuid}/sum-live`);
	}
}
