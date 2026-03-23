import { Injectable } from '@angular/core';
import { environment } from '@environments/environment';
import { HttpClient, HttpParams } from '@angular/common/http';
import { TemplateCreation, TemplateInterface } from '@app/shared/interfaces/template.interface';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class EmailService {

    endpoint = `${environment.host}${environment.email}/api`;

    jsonTemplateValue: string = '';

    constructor(private httpClient: HttpClient) {
    }

    getTemplates(size: number = 10, page: number = 0, sort: string = 'createdAt,desc'): Observable<TemplateInterface[]> {

        let params = new HttpParams()
            .append('size', size)
            .append('page', page)
            .append('sort', sort);

        return this.httpClient.get<TemplateInterface[]>(`${this.endpoint}/templates`, { params });
    }

    getTemplate(id: string): Observable<TemplateInterface> {
        return this.httpClient.get<TemplateInterface>(`${this.endpoint}/templates/${id}`);
    }

    createTemplate(body: TemplateCreation): Observable<TemplateInterface> {
        return this.httpClient.post<TemplateInterface>(`${this.endpoint}/templates`, body);
    }

    deleteTemplate(id: string): Observable<any> {
        return this.httpClient.delete(`${this.endpoint}/templates/${id}`);
    }

    updateTemplate(id: string, body: TemplateCreation): Observable<TemplateInterface> {
        return this.httpClient.patch<TemplateInterface>(`${this.endpoint}/templates/${id}`, body);
    }

    templateInfoDuplicate(event: any): void {
        this.jsonTemplateValue = event.element.templateJson;
    }

}
