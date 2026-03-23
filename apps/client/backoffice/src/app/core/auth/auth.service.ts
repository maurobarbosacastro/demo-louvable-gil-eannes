import {HttpClient, HttpHeaders, HttpParams} from '@angular/common/http';
import {inject, Injectable} from '@angular/core';
import {AuthUtils} from '@app/core/auth/auth.utils';
import {UserService} from '@app/core/user/user.service';
import {catchError, map, Observable, of, switchMap, tap, throwError} from 'rxjs';
import {environment} from '@environments/environment';
import {StorageService} from '@app/shared/services/storage.service';
import {MatDialog} from '@angular/material/dialog';
import {AppConstants, UserTypesEnum} from '@app/app.constants';
import {TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {ActivatedRoute} from '@angular/router';


export enum SocialProviders{
    GOOGLE = 'google',
    FACEBOOK = 'facebook'
}

@Injectable({providedIn: 'root'})
export class AuthService {
	private _authenticated: boolean = false;
	private _httpClient = inject(HttpClient);
	private _userService = inject(UserService);
	private _storageService = inject(StorageService);
	private _dialogRef = inject(MatDialog);
	private _route = inject(ActivatedRoute);

	private redirectUri = environment.host + environment.tagpeak + '/auth/callback';
	userGroups: UserTypesEnum[] = [];

	referralCodeSocialSignup: string;
	referralClickSocialSignup: string;

	// -----------------------------------------------------------------------------------------------------
	// @ Accessors
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Setter & getter for access token
	 */
	set accessToken(token: string) {
		if (token) {
			this.userGroups = this.parseJwt(token).groups;
			this._storageService.save(AppConstants.STORAGE_KEYS.ACCESS_TOKEN, token);
		}
	}

	get accessToken(): string {
		const token = this._storageService.get(AppConstants.STORAGE_KEYS.ACCESS_TOKEN);
		this.userGroups = this.parseJwt(token).groups;
		return token;
	}

	/**
	 * Setter & getter for refresh token
	 */
	set refreshToken(token: string) {
		if (token) {
			this._storageService.save(AppConstants.STORAGE_KEYS.REFRESH_TOKEN, token);
		}
	}

	get refreshToken(): string {
		return this._storageService.get(AppConstants.STORAGE_KEYS.REFRESH_TOKEN);
	}

	set idToken(token: string) {
		if (token){
			this._storageService.save(AppConstants.STORAGE_KEYS.ID_TOKEN, token);
		}
	}

	get idToken(): string {
		return this._storageService.get(AppConstants.STORAGE_KEYS.ID_TOKEN);
	}

	set loginChoice(choice: string) {
		if (choice){
			this._storageService.save(AppConstants.STORAGE_KEYS.LOGIN_CHOICE, choice);
		}
	}

	get loginChoice(): string {
		return this._storageService.get(AppConstants.STORAGE_KEYS.LOGIN_CHOICE);
	}

	// -----------------------------------------------------------------------------------------------------
	// @ Public methods
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Forgot password
	 *
	 * @param email
	 */
	forgotPassword(email: string): Observable<any> {
		const url = `${environment.host}${environment.tagpeak}/auth/reset`;
		return this._httpClient.post(url, {email});
	}

	/**
	 * Reset password
	 *
	 * @param password
	 */
	resetPassword(password: string): Observable<any> {
		return this._httpClient.post('api/auth/reset-password', password);
	}

	/**
	 * Sign in
	 *
	 * @param credentials
	 * @param rememberMe
	 */
	signIn(
		credentials: {email: string; password: string},
		rememberMe: boolean = false,
	): Observable<any> {
		// Throw error, if the user is already logged in
		if (this._authenticated) {
			return throwError('User is already logged in.');
		}

		const {email, password} = credentials;
		const headers = new HttpHeaders().set('Content-Type', 'application/x-www-form-urlencoded');
		let body = new HttpParams()
			.set('grant_type', 'password')
			.set('client_id', environment.keycloak.clientId)
			.set('client_secret', environment.keycloak.clientSecret)
			.set('response_type', 'code')
			.set('scope', 'email profile openid')
			.set('username', email)
			.set('password', password);

		return this._httpClient.post(environment.keycloak.tokenUrl, body, {headers}).pipe(
			switchMap((response: any) => {
				// Store the access token in the local storage
				this.accessToken = response.access_token;
				if (rememberMe) {
					this.refreshToken = response.refresh_token;
				}
				this.idToken = response.id_token;
				this.loginChoice = AppConstants.LOGIN_CHOICE.EMAIL

				// Set the authenticated flag to true
				this._authenticated = true;

				// Return a new observable with the response
				return this._userService.get();
			}),
			switchMap((user: TagpeakUser) => {
				const utmParams = this._storageService.get(AppConstants.STORAGE_KEYS.UTM);
				if (!utmParams) {
					return of(user);
				}
				return this._userService.update({utmParams}, user.uuid).pipe(
					map(() => {
						this._storageService.delete(AppConstants.STORAGE_KEYS.UTM);
						return user;
					}),
				);
			}),
		);
	}

	/**
	 * Sign out
	 */
	signOut(): Observable<any> {
		// Remove the access token from the local storage
		this._storageService.resetRememberMe();
		this._storageService.delete('accessToken');
		this._storageService.delete('refreshToken');

		this._dialogRef.closeAll();

		// Set the authenticated flag to false
		this._authenticated = false;

		this.referralCodeSocialSignup = null;
		this.referralClickSocialSignup = null;

		// Return the observable
		return of(true);
	}

	resetLoginAttempt() {
		this._authenticated = false;
	}

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
		referralCode: string;
	}): Observable<TagpeakUser> {
		const utmParams: string = this._storageService.get(AppConstants.STORAGE_KEYS.UTM);
		if (utmParams) {
			user = {...user, utmParams: utmParams} as any;
		}

		return this._httpClient.post<TagpeakUser>(
			environment.host + environment.tagpeak + '/auth/',
			user,
		);
	}

	/**
	 * Unlock session
	 *
	 * @param credentials
	 */
	unlockSession(credentials: {email: string; password: string}): Observable<any> {
		return this._httpClient.post('api/auth/unlock-session', credentials);
	}

	/**
	 * Check the authentication status
	 */
	check(): Observable<boolean> {
		if (this._authenticated) {
			return of(true);
		}

		// Check the access token availability
		if (!this.accessToken) {
			return of(false);
		}

		return of(AuthUtils.isTokenExpired(this.accessToken)).pipe(
			switchMap((expired) => {
				if (!expired) {
					return this._userService.get().pipe(
						map((_) => {
							return true;
						}),
					);
				}

				if (!this._storageService.rememberMe) {
					this._storageService.resetRememberMe();
					this._storageService.delete('accessToken');

					return of(false);
				}

				return this.refreshTokenCall().pipe(
					catchError(() => {
						this._storageService.resetRememberMe();
						this._storageService.delete('accessToken');
						this._storageService.delete('refreshToken');
						// Return false
						return of(false);
					}),
					tap((response: any) => {
						this.accessToken = response.access_token;
						this.refreshToken = response.refresh_token;
					}),
					switchMap((_) => this._userService.get()),
					map((_) => {
						return true;
					}),
				);
			}),
		);
	}

	loginWithSocials(provider: SocialProviders, state?: string, scope: string = 'openid') {
		const authUrl = environment.keycloak.authUrl;
		const params = new HttpParams({
			fromObject: {
				client_id: environment.keycloak.clientId,
				client_secret: environment.keycloak.clientSecret,
				redirect_uri: this.redirectUri,
				response_type: 'code',
				scope: scope,
				kc_idp_hint: provider, // This triggers Google login
				state: this.generateState(state),
			},
		});
		// Redirect the user to Keycloak's Google login page
		window.location.href = `${authUrl}?${params.toString()}`;
	}

	logoutSocials(){
		window.location.href =
			`${environment.keycloak.logoutUrl}?` +
			`id_token_hint=${encodeURIComponent(this.idToken)}` +
			`&post_logout_redirect_uri=${encodeURIComponent(environment.host)}/#/sign-out`;
	}

	getAuthorizationCode(): string | null {
		let code: string = '';
		this._route.queryParams.subscribe((params) => {
			code = params['code'];
		});
		return code;
	}

	private generateState(state ?: string): string {
        if (state){
            return btoa(state);
        }
		return Math.random().toString(36).substring(2);
	}

	exchangeCodeForToken(code: string): Observable<any> {
		const tokenUrl = environment.keycloak.tokenUrl;

		const body = new HttpParams()
			.set('grant_type', 'authorization_code')
			.set('code', code)
			.set('client_secret', environment.keycloak.clientSecret)
			.set('client_id', environment.keycloak.clientId)
			.set('redirect_uri', this.redirectUri);

		return this._httpClient
			.post(tokenUrl, body, {
				headers: {'Content-Type': 'application/x-www-form-urlencoded'},
			})
			.pipe(
				switchMap((response: any) => {
					// Store the access and refresh token in the local storage
					this.accessToken = response.access_token;
					this.refreshToken = response.refresh_token;
					this.idToken = response.id_token;

					// Set the authenticated flag to true
					this._authenticated = true;

					// Store the user on the user service
					return this._userService.get();
				}),
			);
	}

	refreshTokenCall() {
		const body = new HttpParams()
			.set('grant_type', 'refresh_token')
			.set('refresh_token', this.refreshToken.toString())
			.set('client_secret', environment.keycloak.clientSecret)
			.set('client_id', environment.keycloak.clientId);

		const headers = new HttpHeaders().set('Content-Type', 'application/x-www-form-urlencoded');

		return this._httpClient.post(environment.keycloak.tokenUrl, body, {headers});
	}

	parseJwt(token: string): any {
		if (!token || token.trim() === '') {
			return [];
		}
		const base64Url = token.split('.')[1];
		const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
		const jsonPayload = decodeURIComponent(
			window
				.atob(base64)
				.split('')
				.map(function (c) {
					return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
				})
				.join(''),
		);

		return JSON.parse(jsonPayload);
	}

	validateActionUser(email: string): Observable<void> {
		let params: HttpParams = new HttpParams();
		params = params.append('email', email);
		return this._httpClient.get<void>(
			environment.host + environment.tagpeak + '/public/auth/validate-action',
			{params},
		);
	}

	authValid(): Observable<boolean> {
		return this._httpClient.get<any>(environment.host + environment.tagpeak + '/auth/valid').pipe(
			map((_) => true),
			catchError((_) => of(false)),
		);
	}

	checkIfShopifyStoreOwner(domain: string): Observable<void> {
		return this._httpClient.get<void>(
			environment.host + environment.tagpeak + '/shopify/shop/owner',
			{params: {domain}},
		);
	}

	finishSocialAuth(user: {
		firstName?: string;
		lastName?: string;
		currency: string;
	}): Observable<TagpeakUser> {
		return this._httpClient.post<TagpeakUser>(
			environment.host + environment.tagpeak + '/auth/socials/finish-profile',
			user,
		);
	}

    handleSocialsState(){
        let state: string = '';
        this._route.queryParams.subscribe((params) => {
            state = params['state'];
        });
        if (state){
            const decodedState = atob(state);
            try {
                const referralInfo: {referralCode: string; referralClick: string;} = JSON.parse(decodedState);
                if (referralInfo.referralCode && referralInfo.referralClick){
                    this.referralCodeSocialSignup = referralInfo.referralCode;
                    this.referralClickSocialSignup = referralInfo.referralClick;
                }
            } catch (e){
                //Not an object, so no referral and we can ignore
            }
        }

    }
}
