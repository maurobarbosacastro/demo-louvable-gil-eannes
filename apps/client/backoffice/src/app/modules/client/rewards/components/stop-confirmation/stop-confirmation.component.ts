import {Component} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe} from '@ngneat/transloco';
import {MatDialogClose, MatDialogRef} from '@angular/material/dialog';
import {ModalActionResult} from '@app/utils/modal-action-result.enum';
import {environment} from '@environments/environment';

@Component({
	selector: 'tagpeak-stop-confirmation',
	standalone: true,
	imports: [CommonModule, MatIcon, TranslocoPipe, MatDialogClose],
	templateUrl: './stop-confirmation.component.html',
})
export class StopConfirmationComponent {
	constructor(private readonly myDialog: MatDialogRef<{action: ModalActionResult}>) {}

	redirectToSomewhere(): void {
		window.open(environment.commercial.faqUrl, '_blank');
	}

	stopReward(): void {
		this.myDialog.close({action: ModalActionResult.OK});
	}
}
