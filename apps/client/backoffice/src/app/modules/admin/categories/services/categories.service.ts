import {inject, Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {environment} from '@environments/environment';
import {Observable} from 'rxjs';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {CategoriesInterface} from '@app/modules/admin/categories/models/categories.interface';

@Injectable({
	providedIn: 'root',
})
export class CategoriesService {
	private http: HttpClient = inject(HttpClient);
	endpoint: string = `${environment.host}${environment.tagpeak}/category`;
	publicEndpoint: string = `${environment.host}${environment.tagpeak}/public/category`;

	getAllCategories(
		size?: number,
		page?: number,
		sort?: string,
	): Observable<PaginationInterface<CategoriesInterface>> {
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

		return this.http.get<PaginationInterface<CategoriesInterface>>(this.endpoint, {params});
	}

	getAllCategoriesPublic(
		size?: number,
		page?: number,
		sort?: string,
		filters?: {name?: string},
	): Observable<PaginationInterface<CategoriesInterface>> {
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
		if (filters.name) {
			params = params.append('name', filters.name);
		}

		return this.http.get<PaginationInterface<CategoriesInterface>>(this.publicEndpoint, {params});
	}

	getCategoryById(id: string): Observable<CategoriesInterface> {
		return this.http.get<CategoriesInterface>(`${this.endpoint}/${id}`);
	}

	getCategoryByCode(code: string): Observable<CategoriesInterface> {
		return this.http.get<CategoriesInterface>(`${this.publicEndpoint}/code/${code}`);
	}

	createCategory(body: Partial<CategoriesInterface>): Observable<void> {
		return this.http.post<void>(this.endpoint, body);
	}

	updateCategory(body: Partial<CategoriesInterface>, id: string): Observable<void> {
		return this.http.patch<void>(`${this.endpoint}/${id}`, body);
	}

	deleteCategory(id: string): Observable<void> {
		return this.http.delete<void>(`${this.endpoint}/${id}`);
	}
}
