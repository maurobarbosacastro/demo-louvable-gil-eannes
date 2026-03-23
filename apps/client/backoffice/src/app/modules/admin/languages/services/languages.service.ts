import {inject, Injectable} from '@angular/core';
import {HttpClient, HttpParams} from "@angular/common/http";
import {Observable} from "rxjs";
import {environment} from "@environments/environment";
import {PaginationInterface} from "@app/shared/interfaces/pagination.interface";
import {LanguagesInterface} from "@app/modules/admin/languages/models/languages.interface";

@Injectable({
  providedIn: 'root'
})
export class LanguagesService {

    private http: HttpClient = inject(HttpClient);
    endpoint: string = `${environment.host}${environment.tagpeak}/language`;

    getAllLanguages(size?: number, page?: number, sort?: string): Observable<PaginationInterface<LanguagesInterface>> {
        let params: HttpParams = new HttpParams()

        if (size) {
            params = params.append('limit', size);
        }
        if (page) {
            params = params.append('page', page);
        }
        if (sort) {
            params = params.append('sort', sort);
        }

        return this.http.get<PaginationInterface<LanguagesInterface>>(this.endpoint, { params });
    }

    getLanguageById(id: string): Observable<LanguagesInterface> {
        return this.http.get<LanguagesInterface>(`${this.endpoint}/${id}`);
    }

    createLanguage(body: Partial<LanguagesInterface>): Observable<void> {
        return this.http.post<void>(this.endpoint, body);
    }

    updateLanguage(body: Partial<LanguagesInterface>, id: string): Observable<void> {
        return this.http.patch<void>(`${this.endpoint}/${id}`, body);
    }

    deleteLanguage(id: string): Observable<void> {
        return this.http.delete<void>(`${this.endpoint}/${id}`);
    }
}
