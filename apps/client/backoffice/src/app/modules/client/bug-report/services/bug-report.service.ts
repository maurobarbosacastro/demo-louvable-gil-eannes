import {inject, Injectable} from '@angular/core';
import {environment} from '@environments/environment';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {BugReportRequest} from '@app/modules/client/bug-report/models/bug-report.interface';

@Injectable({
	providedIn: 'root',
})
export class BugReportService {
	private httpClient = inject(HttpClient);

	submitBugReport(body: BugReportRequest): Observable<void> {
		return this.httpClient.post<void>(
			`${environment.host}${environment.tagpeak}/bug-report`,
			body,
		);
	}
}
