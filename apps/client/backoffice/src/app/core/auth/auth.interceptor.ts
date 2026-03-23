import {HttpErrorResponse, HttpEvent, HttpHandlerFn, HttpRequest} from '@angular/common/http';
import {inject} from '@angular/core';
import {AuthService} from '@app/core/auth/auth.service';
import {AuthUtils} from '@app/core/auth/auth.utils';
import {catchError, map, Observable, of, switchMap, tap, throwError} from 'rxjs';
import {StorageService} from '@app/shared/services/storage.service';
import {Router} from '@angular/router';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {TranslocoService} from '@ngneat/transloco';
import {AuxService} from '@app/shared/services/aux.service';

const excludedUrls = [
	'/protocol/openid-connect/token',
    '/auth/valid'
	// Add more URLs or patterns you want to exclude
];

/**
 * Intercept
 *
 * @param req
 * @param next
 */
export const authInterceptor = (
	req: HttpRequest<unknown>,
	next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> => {
	const isExcluded = excludedUrls.some((url) => req.url.includes(url));

	if (isExcluded) {
		return next(req);
	}

	const authService = inject(AuthService);
	const storageService = inject(StorageService);
	const router = inject(Router);
	const toasterService = inject(ToasterService);
	const translocoService = inject(TranslocoService);
	const auxService = inject(AuxService);

	const ipifyUrl = 'api.ipify.org';

	// Clone the request object
	let newReq = req.clone();

	let loadedToastr = false;

	let renewToken: Observable<any>;
	if (
		authService.accessToken &&
		AuthUtils.isTokenExpired(authService.accessToken) &&
		storageService.rememberMe
	) {
		renewToken = authService.refreshTokenCall().pipe(
			catchError(() => {
				storageService.resetRememberMe();
				storageService.delete('accessToken');
				storageService.delete('refreshToken');
				// Return false
				return of(false);
			}),
			tap((response: any) => {
				authService.accessToken = response.access_token;
				authService.refreshToken = response.refresh_token;
			}),
			map((_) => true),
		);
	}

	// Request
	//
	// If the access token didn't expire, add the Authorization header.
	// We won't add the Authorization header if the access token expired.
	// This will force the server to return a "401 Unauthorized" response
	// for the protected API routes which our response interceptor will
	// catch and delete the access token from the local storage while logging
	// the user out from the app.
	if (
		authService.accessToken &&
		!AuthUtils.isTokenExpired(authService.accessToken) &&
		!newReq.url.includes(ipifyUrl)
	) {
		newReq = req.clone({
			headers: req.headers.set('Authorization', 'Bearer ' + authService.accessToken),
		});
	}


    return of({}).pipe(
		switchMap((_) =>
			!!renewToken ? renewToken.pipe(switchMap((_) => next(newReq))) : next(newReq),
		),
		catchError((error) => {
			// Catch "401 Unauthorized" responses
			if (error instanceof HttpErrorResponse && error.status === 401) {
				loadedToastr = auxService.getAuxToastrVar();
				if (!loadedToastr && !error.error.ErrorType) {
					toasterService.showToast(
						'warning',
						translocoService.translate('warning.lost-session'),
						'top',
						{},
						6000,
					);
					auxService.setAuxToastrVar(true);
				}

				// Sign out
				authService.signOut();

                ///shopify/shop-sign-in?shop=diogotag.myshopify.com
				//Redirect user to sign-in page and back to the previous page
				const redirectURL =
					router.routerState.snapshot.url === '/sign-out'
						? ''
						: `redirectURL=${router.routerState.snapshot.url}`;

				const urlTree = router.parseUrl(redirectURL === 'redirectURL=' ? 'sign-in' : `sign-in?${redirectURL}`);
				router.navigateByUrl(urlTree);
			}

			return throwError(error);
		}),
	);
};
