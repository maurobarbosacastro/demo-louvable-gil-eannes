import {CanActivateFn, Router} from '@angular/router';
import {AuthService} from '@app/core/auth/auth.service';
import {inject} from '@angular/core';
import {UserTypesEnum} from '@app/app.constants';

export const rolesGuard: CanActivateFn = (route, state) => {
	const authService = inject(AuthService);
	const router: Router = inject(Router);

	if (!authService.userGroups.includes(UserTypesEnum.admin)) {
		if (authService.userGroups.includes(UserTypesEnum.default)) {
			router.navigate(['/client']);
		}
		if (authService.userGroups.includes(UserTypesEnum.shop)) {
			router.navigate(['/shop']);
		}

		return false;
	}

	return true;
};

export const clientGuard: CanActivateFn = (route, state) => {
    const authService = inject(AuthService);
    const router: Router = inject(Router);

    if (!authService.userGroups.includes(UserTypesEnum.default)) {
        router.navigate(['/admin']);
        return false;
    }

    return true;
};


export const shopGuard: CanActivateFn = (route, state) => {
    const authService = inject(AuthService);
    const router: Router = inject(Router);

    if (!authService.userGroups.includes(UserTypesEnum.shop)) {
        router.navigate(['/client']);
        return false;
    }

    return true;
};
