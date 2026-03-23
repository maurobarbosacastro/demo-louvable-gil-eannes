import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {Configuration, PatchConfiguration} from '../interfaces/configuration.interface';
import {Observable, take} from 'rxjs';
import {HttpClient, HttpParams} from '@angular/common/http';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {AppConstants} from '@app/app.constants';

@Injectable({
	providedIn: 'root',
})
export class ConfigurationService {
	private http = inject(HttpClient);
	endpoint = `${environment.host}${environment.tagpeak}/configuration`;

	configs: {[key: string]: Configuration} = {};

	loadConfigurations(): void {
		this.http
			.get(`${this.endpoint}/latest`)
			.pipe(take(1))
			.subscribe((res: any) => {
				this.configs = res;
			});
	}

	getConfigurations(
		page: number = 1,
		pageSize: number = AppConstants.PAGE_SIZE,
		sort: string = 'id asc',
	): Observable<PaginationInterface<Configuration>> {
		const params = new HttpParams()
			.set('page', page.toString())
			.set('limit', pageSize.toString())
			.set('sort', sort);
		return this.http.get<PaginationInterface<Configuration>>(`${this.endpoint}`, {params});
	}

	updateConfiguration(id: number, body: PatchConfiguration): Observable<Configuration> {
		return this.http.patch<Configuration>(`${this.endpoint}/${id}`, body);
	}
}
