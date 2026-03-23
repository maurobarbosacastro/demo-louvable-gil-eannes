import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {
	ReferralInfo,
	ReferralInterface,
	RevenueInfo,
	UserReferralInterface,
} from '@app/modules/client/referrals/models/referrals.interface';
import {MembershipLevelInfoInterface} from '@app/shared/interfaces/membership-level.interface';
import {MembershipLevelsEnum} from '@app/app.constants';
import {ConfigurationService} from '@app/modules/admin/configuration/services/configuration.service';

@Injectable({
	providedIn: 'root',
})
export class ReferralsService {
	endpoint: string = `${environment.host}${environment.tagpeak}/auth`;
	referralEndpoint: string = `${environment.host}${environment.tagpeak}/referral`;

	private http: HttpClient = inject(HttpClient);
	private _configService: ConfigurationService = inject(ConfigurationService);

	// Get all store visits of all users
	getAllReferralsByUserUuid(
		uuid: string,
		page: number,
		size: number,
		sort?: string,
	): Observable<PaginationInterface<ReferralInterface>> {
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
		return this.http.get<PaginationInterface<ReferralInterface>>(
			`${this.endpoint}/${uuid}/referral`,
			{params},
		);
	}

	getReferralInfo(userUuid: string): Observable<ReferralInfo> {
		return this.http.get<ReferralInfo>(`${this.endpoint}/${userUuid}/referral/clicks`);
	}

	getRevenueInfo(userUuid: string): Observable<RevenueInfo> {
		return this.http.get<RevenueInfo>(`${this.endpoint}/${userUuid}/referral/revenue`);
	}

	getUsersRevenueInfo(userUuid: string): Observable<UserReferralInterface[]> {
		return this.http.get<UserReferralInterface[]>(
			`${this.endpoint}/${userUuid}/referral/revenue/users-info`,
		);
	}

	getRevenueTotalValue(userUuid: string): Observable<number> {
		return this.http.get<number>(`${this.endpoint}/${userUuid}/referral/revenue/total-revenue`);
	}

	createReferralClick(code: string): Observable<any> {
		return this.http.post(`${this.referralEndpoint}/click`, {code});
	}

	getInfoBaseOnMembershipLevel(level: string): Partial<MembershipLevelInfoInterface> {
		switch (level) {
			case MembershipLevelsEnum.base:
				return {
					level: level,
					percentageOnTransaction: '0%',
					percentageOnReward:
						this._configService.configs['referral_member_friend_cash_reward'].value,
					maxValue: Number(this._configService.configs['referral_silver_status_goal_amount'].value),
				};
			case MembershipLevelsEnum.silver:
				return {
					level: level,
					percentageOnTransaction:
						this._configService.configs['referral_silver_transaction_cash_reward'].value,
					percentageOnReward: this._configService.configs['referral_gold_status_goal'].value,
					maxValue: Number(this._configService.configs['referral_gold_status_goal_amount'].value),
				};
			case MembershipLevelsEnum.gold:
				return {
					level: level,
					percentageOnTransaction:
						this._configService.configs['referral_gold_transaction_cash_reward'].value,
					percentageOnReward:
						this._configService.configs['referral_gold_friend_reward_share'].value,
					maxValue: Number(this._configService.configs['referral_gold_status_goal_amount'].value),
				};
		}
	}

	validateReferralCode(code: string): Observable<boolean> {
		let params: HttpParams = new HttpParams().append('code', code);

		return this.http.get<boolean>(`${this.referralEndpoint}/validate`, {params});
	}
}
