import {Component, Inject, OnInit} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TranslocoPipe} from '@ngneat/transloco';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {MatIcon} from '@angular/material/icon';
import {environment} from '@environments/environment';

interface DialogContentInterface {
	title: string;
	body: string;
	cancelButton?: string;
	okButton?: string;
	showFAQ?: boolean;
}

@Component({
	selector: 'tagpeak-delete-modal-confirmation',
	standalone: true,
	imports: [CommonModule, TranslocoPipe, MatIcon],
	template: `
        <div class="flex flex-col gap-5">

            <div class="flex flex-col justify-center items-center">
                <span class="font-semibold text-xl">{{ modalContent.title | transloco }}</span>
            </div>

            <div class="flex flex-col justify-center items-center text-center">
                <span class="text-md text-project-waterloo">{{ modalContent.body | transloco }}</span>
            </div>

            <div class="flex flex-row justify-center items-center gap-2.5">
                <button class="btn border-project-licorice-900 bg-white text-project-licorice-900 text-base font-normal min-h-9 h-9 rounded" (click)="closeModal()">{{ modalContent.cancelButton | transloco }}</button>
                <button class="btn border-project-red bg-project-red text-white text-base font-normal min-h-9 h-9 rounded" (click)="deleteCashback()">{{ modalContent.okButton | transloco }}</button>
            </div>

        </div>
        @if (modalContent.showFAQ) {
            <div class="mt-8 border-t border-project-sideMenu">
                <div class="flex justify-between pt-4">
                    <span class="text-sm text-project-waterloo">{{ 'dashboard.stop-confirmation.footer-message' | transloco }}</span>
                    <mat-icon svgIcon="arrow-right" class="icon-size-4" (click)="goToFAQ()"></mat-icon>
                </div>
            </div>
        }
    `,
})
export class DeleteModalConfirmationComponent implements OnInit {
	constructor(
		@Inject(MAT_DIALOG_DATA)
		private readonly receivedData: DialogContentInterface,
		private readonly myDialog: MatDialogRef<DeleteModalConfirmationComponent, {skipFetch: boolean}>,
	) {}

	modalContent: DialogContentInterface = {
		title: 'modal-delete.title',
		body: 'modal-delete.body',
		cancelButton: 'misc.cancel',
		okButton: 'cashback.delete-cashback.confirm',
		showFAQ: false,
	};

	ngOnInit() {
		if (this.receivedData) {
			const {title, body, okButton, cancelButton, showFAQ} = this.receivedData;

			this.modalContent = {
				title: title || this.modalContent?.title,
				body: body || this.modalContent?.body,
				okButton: okButton || this.modalContent?.okButton,
				cancelButton: cancelButton || this.modalContent?.cancelButton,
				showFAQ: showFAQ || this.modalContent?.showFAQ,
			};
		}
	}

	goToFAQ(): void {
		window.open(environment.commercial.faqUrl, '_blank');
	}

	closeModal(): void {
		this.myDialog.close({skipFetch: true});
	}

	deleteCashback(): void {
		this.myDialog.close({skipFetch: false});
	}
}
