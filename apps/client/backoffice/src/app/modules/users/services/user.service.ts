import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {APIUsersList} from '@app/modules/users/interfaces/user.interface';

@Injectable({
    providedIn: 'root'
})
export class UserService {

    constructor(private httpClient: HttpClient,) {
    }


    getUsers(page: number, email: string): Observable<APIUsersList> {
        let params = new HttpParams().append('size', 9).append('page', page);
        params = params.append('params', email);
        const apiRoot = 'api/common/users'; //environment.backendAPIHost + environment.userManagement + '/api/users';
        return this.httpClient.get<APIUsersList>(apiRoot, {params});
    }

    blockUser(email: string): Observable<any> {
        const apiRoot = ''; //environment.backendAPIHost + environment.userManagement + '/api/users/revoke-comment';
        return this.httpClient.patch<any>(apiRoot, {email: email});
    }

}
