import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MAT_DIALOG_DATA, MatDialog, MatDialogClose, MatDialogRef } from '@angular/material/dialog';
import { MatIcon } from '@angular/material/icon';
import { TranslocoPipe } from '@ngneat/transloco';
import { OnboardingDialogComponent } from '@app/modules/onboarding/onboarding-dialog/onboarding-dialog.component';
import { UserService } from '@app/core/user/user.service';

@Component({
    selector: 'tagpeak-account-verified-dialog',
    standalone: true,
    imports: [CommonModule, MatDialogClose, MatIcon, TranslocoPipe],
    templateUrl: './account-verified-dialog.component.html',
    styleUrl: './account-verified-dialog.component.scss'
})
export class AccountVerifiedDialogComponent {
    dialog = inject(MatDialog);
    dialogData = inject(MAT_DIALOG_DATA);
    dialogRef = inject(MatDialogRef);
    userService = inject(UserService);

    constructor() {
    }

    closeAndSkipOnboarding(): void {
        this.dialogRef.close({ skipOnboarding: true });
    }

    closeAndGoToOnboarding(): void {

        // Close account verification dialog and open onboarding dialog
        this.dialogRef.close({ skipOnboarding: false });

        const dialogRef = this.dialog.open(OnboardingDialogComponent, {
            panelClass: 'onboarding-dialog-panel',
            position: { right: '2rem' }
        });

        dialogRef.afterClosed().subscribe(_ => {
            this.userService.update({ onboardingFinished: 'true' }, this.dialogData).subscribe();
        });
    }
}
