import {inject, Injectable, signal, WritableSignal} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '@environments/environment';
import {CashbackShopTableResponse, DashboardStats, ShopifyShop, TagpeakShopUser} from '@app/modules/shop/interfaces/shop.interface';
import {AppConstants} from '@app/app.constants';
import {ManageCashbackSearchParams} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';

@Injectable({
	providedIn: 'root',
})
export class ShopService {
    private http = inject(HttpClient);
    shopName: string = null;

    private basePubicUrl = environment.host + environment.tagpeak + '/public/shopify';
    private baseUrl = environment.host + environment.tagpeak + '/shopify';

    $currentShop: WritableSignal<{userUuid: string; shopUuid: string; storeUuid: string; uuid: string}> = signal(null);

	/**
	 * Sign up
	 *
	 * @param user
	 */
	signUp(user: {
		firstName: string;
		lastName: string;
		email: string;
		password: string;
		currency: string;
		shop: string;
	}): Observable<TagpeakShopUser> {
		return this.http.post<TagpeakShopUser>(
            this.basePubicUrl + '/shop/auth',
			user,
		);
	}

    getShopByUserUuid(uuid: string): Observable<ShopifyShop> {
        return this.http.get<ShopifyShop>(`${this.baseUrl}/user/${uuid}`);
    }

    getShopStats(shopUuid: string): Observable<DashboardStats>{
        return this.http.get<DashboardStats>(`${this.baseUrl}/shop/${shopUuid}/stats`);
    }

    getCashback(
        page: number = 0,
        size: number = AppConstants.PAGE_SIZE,
        sort: string,
        filtersBody: Partial<ManageCashbackSearchParams>,
    ): Observable<PaginationInterface<CashbackShopTableResponse>> {
        let params: HttpParams = new HttpParams().append('page', page).append('limit', size);

        if (sort) {
            params = params.append('sort', sort);
        }

        if (filtersBody?.state) {
            params = params.append('state', filtersBody.state);
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
        const endpoint = this.baseUrl + `/store/${filtersBody.storeUuid}/transaction`;
        return this.http.get<PaginationInterface<CashbackShopTableResponse>>(endpoint, {params});
    }
}
