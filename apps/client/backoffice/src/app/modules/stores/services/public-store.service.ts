import { inject, Injectable } from '@angular/core';
import { environment } from '@environments/environment';
import { HttpClient, HttpParams } from '@angular/common/http';
import { catchError, map, Observable, of, switchMap, tap, throwError } from 'rxjs';
import { PaginationInterface } from '@app/shared/interfaces/pagination.interface';
import { StoresInterface } from '@app/modules/stores/models/stores.interface';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
    providedIn: 'root',
})
export class PublicStoreService {
    publicEndpoint = `${environment.host}${environment.tagpeak}/public/store`;
    storeEndpoint = `${environment.host}${environment.tagpeak}/store`;
    auxEndpoint = `${environment.host}${environment.tagpeak}/public`;

    private httpClient = inject(HttpClient);
    cookieService = inject(CookieService);

    _loadedStore: StoresInterface = null;

    get loadedStore(): StoresInterface {
        return this._loadedStore;
    }

    set loadedStore(value: StoresInterface) {
        this._loadedStore = value;
    }

    getStoresPublic(
        size?: number,
        page?: number,
        sort?: string,
        countryCode?: string,
        categoryCode?: string,
        name?: string,
    ): Observable<PaginationInterface<StoresInterface>> {
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
        if (countryCode) {
            params = params.append('countryCode', countryCode);
        }
        if (categoryCode) {
            params = params.append('categoryCode', categoryCode);
        }
        if (name) {
            params = params.append('name', name);
        }

        return this.httpClient.get<PaginationInterface<StoresInterface>>(`${this.publicEndpoint}`, {
            params,
        });
    }

    getStorePublic(id: string): Observable<StoresInterface> {
        return this.httpClient.get<StoresInterface>(`${this.publicEndpoint}/${id}`)
            .pipe(
                tap(store => this.loadedStore = store)
            );
    }

    getIpAddress(): Observable<{ code: string; name: string }> {
        const sessionStorageValue = this.cookieService.get('tp_detected_country');

        if (!!sessionStorageValue) {
            return of(JSON.parse(sessionStorageValue))
        }

        return this.httpClient.get(`https://api.ipify.org`, { responseType: 'text' }).pipe(
            switchMap((ip) => this.httpClient.get(`${this.auxEndpoint}/ip`, { params: { ip } })),
            map((countryInfo: any) => ({ code: countryInfo.country_code, name: countryInfo.country_name })),
            tap(value => this.cookieService.set('tp_detected_country', JSON.stringify(value))),
        );
    }

    getStoreRedirectUrl(id: string): Observable<string> {
        return this.httpClient.get(`${this.storeEndpoint}/${id}/redirect`, {
            responseType: 'text',
        })
            .pipe(
                map(url => decodeURIComponent(JSON.parse(url))),
            );
    }
}
