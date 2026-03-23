import { inject } from '@angular/core';
import { NavigationService } from '@app/core/navigation/navigation.service';
import {forkJoin, of, switchMap, tap} from 'rxjs';
import {StorageService} from '@app/shared/services/storage.service';
import {AppConstants} from '@app/app.constants';
import {UserService} from '@app/core/user/user.service';

export const initialDataResolver = () => {
    const navigationService = inject(NavigationService);

    // Fork join multiple API endpoint calls to wait all of them to finish
    return forkJoin([
        navigationService.get()
    ]);
};


export const manageStoredUtmParams = () => {
    const storageService = inject(StorageService);
    const userService = inject(UserService);
    const utmParams = storageService.get(AppConstants.STORAGE_KEYS.UTM);
    if (utmParams) {
        return userService.get()
            .pipe(
                switchMap((user) => userService.update({utmParams}, user.uuid)),
                tap(() => {
                    storageService.delete(AppConstants.STORAGE_KEYS.UTM);
                }),
            );
    }
    return of(true)
}
