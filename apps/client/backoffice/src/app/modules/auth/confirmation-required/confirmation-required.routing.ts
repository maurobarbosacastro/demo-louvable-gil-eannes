import { Routes } from '@angular/router';
import { AuthConfirmationRequiredComponent } from '@app/modules/auth/confirmation-required/confirmation-required.component';

export const authConfirmationRequiredRoutes: Routes = [
    {
        path: '',
        component: AuthConfirmationRequiredComponent,
    },
];
