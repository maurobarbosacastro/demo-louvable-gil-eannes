import { inject } from '@angular/core';
import { CanActivateChildFn, CanActivateFn, Router } from '@angular/router';
import { AuthService } from '@app/core/auth/auth.service';
import { of, switchMap } from 'rxjs';

export const NoAuthShopGuard: CanActivateFn | CanActivateChildFn = (
    route,
    state
) => {
    const router: Router = inject(Router);
    const authService = inject(AuthService);

    // Check the authentication status
    return authService
        .authValid()
        .pipe(
            switchMap((authenticated) => {
                // If the user is authenticated...
                if (authenticated) {
                    return authService.check()
                        .pipe(
                            switchMap( _ => of(router.parseUrl('/shop')))
                        )
                }

                // Allow the access
                return of(true);
            })
        );
};
