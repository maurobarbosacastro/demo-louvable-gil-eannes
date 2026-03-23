import {HttpClient, HttpParams} from '@angular/common/http';
import {computed, inject, Injectable, Signal} from '@angular/core';
import {catchError, Observable, ReplaySubject, switchMap, tap, throwError} from 'rxjs';
import {environment} from '@environments/environment';
import {MembershipLevelInterface, TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {toSignal} from '@angular/core/rxjs-interop';

@Injectable({providedIn: 'root'})
export class UserService {
	private _httpClient = inject(HttpClient);
	private _user: ReplaySubject<TagpeakUser> = new ReplaySubject<TagpeakUser>(1);

	// -----------------------------------------------------------------------------------------------------
	// @ Accessors
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Setter & getter for user
	 *
	 * @param value
	 */
	set user(value: TagpeakUser) {
		// Store the value
		this._user.next(value);
	}

	get user$(): Observable<TagpeakUser> {
		return this._user.asObservable();
	}

	$user: Signal<TagpeakUser> = toSignal(this.user$);
	$isProfileCompleted: Signal<boolean> = computed( () => {
		console.log('computed', this.$user());
		if (!this.$user()) return false;

		return this.$user().displayName !== "" && this.$user().currencySelected && this.$user().referralCode !== ""
	})

	// -----------------------------------------------------------------------------------------------------
	// @ Public methods
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Get the current signed-in user data
	 */
	get(): Observable<TagpeakUser> {
		return this._httpClient
			.get<TagpeakUser>(environment.host + environment.tagpeak + '/auth/me')
			.pipe(
				tap((user) => {
					this._user.next(user);
				}),
			);
	}

	/**
	 * Update the user
	 *
	 * @param user
	 * @param userId
	 */
	update(user: Partial<TagpeakUser>, userId: string): Observable<any> {
		return this._httpClient
			.patch<TagpeakUser>(environment.host + environment.tagpeak + `/auth/${userId}`, user)
			.pipe(
				tap((user) => {
					this._user.next(user);
				}),
				catchError((error) => {
					// Show the error message in the toaster
					return throwError(() => error);
				}),
			);
	}

	setEmailVerified(userId: string, isVerified: boolean): Observable<boolean> {
		const body = {
			userId: userId,
			isVerified: isVerified ? 'true' : 'false',
		};

		return this._httpClient.post<boolean>(
			environment.host + environment.tagpeak + `/auth/email-verified`,
			body,
		);
	}

	getUUIDFromUrlImage(url: string): string {
		const parts = url.split('/');
		const index = parts.lastIndexOf('images') + 1;
		return parts[index] || null;
	}

	updateProfilePicture(
		file: File,
		fileName: string,
		userUuid: string,
		previousPictureUuid?: string,
	): Observable<string> {
		const formData = new FormData();
		formData.append('file', file);
		formData.append('fileName', fileName);

		if (previousPictureUuid) {
			return this._httpClient
				.delete<void>(
					`${environment.host + environment.tagpeak}/auth/${userUuid}/profile-picture/${this.getUUIDFromUrlImage(previousPictureUuid)}`,
				)
				.pipe(
					switchMap(() =>
						this._httpClient.post(
							`${environment.host + environment.tagpeak}/auth/${userUuid}/profile-picture`,
							formData,
							{responseType: 'text'},
						),
					),
				);
		} else {
			return this._httpClient.post(
				`${environment.host + environment.tagpeak}/auth/${userUuid}/profile-picture`,
				formData,
				{responseType: 'text'},
			);
		}
	}

	getCurrentMembershipLevel(userUuid: string): Observable<MembershipLevelInterface> {
		return this._httpClient.get<MembershipLevelInterface>(
			`${environment.host + environment.tagpeak}/auth/${userUuid}/stats`,
		);
	}

	getAllUsers(limit: number, user: string): Observable<TagpeakUser[]> {
		let params = new HttpParams().append('limit', limit);
		if (user) {
			params = params.append('searchUser', user);
		}
		return this._httpClient.get<any>(environment.host + environment.tagpeak + '/auth/users', {
			params,
		});
	}

	setCurrency(userUuid: string, currencyCode: string): Observable<any> {
		return this._httpClient.patch(environment.host + environment.tagpeak + `/auth/${userUuid}/currency`, {currencyCode})
	}
}
