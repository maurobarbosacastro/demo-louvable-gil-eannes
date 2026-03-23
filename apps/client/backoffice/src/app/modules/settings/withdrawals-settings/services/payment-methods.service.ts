import {inject, Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {environment} from '@environments/environment';
import {map, Observable} from 'rxjs';
import {
    AvailablePaymentMethod, CreateUserPaymentMethod,
    UserPaymentMethodInterface,
} from '@app/modules/settings/withdrawals-settings/models/user-payment-method.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {AppConstants} from '@app/app.constants';

@Injectable({
	providedIn: 'root',
})
export class PaymentMethodsService {
    endpointRoot: string = `${environment.host}${environment.tagpeak}`;
	endpointPaymentMethod: string = `${this.endpointRoot}/payment-method`;
	endpoint: string = `${this.endpointRoot}/user/payment-method`;
	private httpClient = inject(HttpClient);

	getPaymentMethods(
        page: number = 1,
        size: number = AppConstants.PAGE_SIZE,
        sort: string = 'state desc',
    ): Observable<PaginationInterface<UserPaymentMethodInterface>> {
        let params = new HttpParams()
            .set('page', page.toString())
            .set('limit', size.toString())
            .set('sort', sort);

		return this.httpClient.get<PaginationInterface<UserPaymentMethodInterface>>(this.endpoint, {params});
	}

	getPaymentMethodById(id: string): Observable<UserPaymentMethodInterface> {
		return this.httpClient.get<UserPaymentMethodInterface>(`${this.endpoint}/${id}`);
	}

	createPaymentMethod(body: CreateUserPaymentMethod): Observable<void> {
		return this.httpClient.post<void>(this.endpoint, body);
	}

	deletePaymentMethod(id: string): Observable<void> {
		return this.httpClient.delete<void>(`${this.endpoint}/${id}`);
	}

	getPaymentMethodsAvailable(): Observable<AvailablePaymentMethod[]> {
		return this.httpClient.get<AvailablePaymentMethod[]>(this.endpointPaymentMethod);
	}

	uploadIbanStatement(file: File): Observable<string> {
		const formData = new FormData();
		formData.append('file', file);
		return this.httpClient.post<string>(`${this.endpoint}/file`, formData);
	}

	downloadFile(id: string): Observable<File> {
		return this.httpClient
			.get(`${this.endpoint}/file/${id}`, {responseType: 'blob'})
			.pipe(map((response: Blob) => new File([response], 'IBAN.pdf', {type: 'application/pdf'})));
	}

    checkVatValidity(vatNumber: string): Observable<boolean> {
        return this.httpClient.post<boolean>(`${this.endpointRoot}/aux/vat`, null, {params: {vat_number: vatNumber}})
            .pipe(
                map((res: any) => res.success)
            );

    }
}
