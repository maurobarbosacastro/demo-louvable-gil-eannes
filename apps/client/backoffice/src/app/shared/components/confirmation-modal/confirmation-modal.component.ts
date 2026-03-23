import {Component, Inject} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MAT_DIALOG_DATA, MatDialogClose, MatDialogRef} from '@angular/material/dialog';
import {TranslocoPipe} from '@ngneat/transloco';

@Component({
	selector: 'tagpeak-confirmation-modal',
	standalone: true,
	imports: [CommonModule, TranslocoPipe],
	template: `
        <div class="flex flex-col items-center justify-center">
            <div class="flex flex-col items-center justify-center">
                <div class="flex flex-col items-center justify-center gap-5">
            <span class="text-2xl font-bold text-center text-project-licorice-500">
                @if (receivedData.title) {
                    {{ receivedData.title }}
                }
            </span>
                    <span class="text-center text-md text-project-waterloo">
               {{ receivedData.message }}
            </span>
                </div>
                <div class="flex flex-row items-center justify-center gap-4 mt-5">
                    <button class="btn border-project-licorice-900 bg-white text-project-licorice-900 text-base font-normal min-h-9 h-9 rounded"
                            (click)="onCancel()">
                        {{ 'misc.cancel' | transloco }}
                    </button>
                    <button class="btn border-project-licorice-900 bg-project-licorice-900  text-white text-base font-normal min-h-9 h-9 rounded"
                            (click)="onConfirm()">
                        {{ 'misc.confirm' | transloco }}
                    </button>
                </div>
            </div>
        </div>
    `,
})
export class ConfirmationModalComponent {
	constructor(
		@Inject(MAT_DIALOG_DATA)
		protected readonly receivedData: {title?: string; message: string},
		private readonly myDialog: MatDialogRef<ConfirmationModalComponent, {next: boolean}>,
	) {}

	onConfirm() {
		this.myDialog.close({next: true});
	}

	onCancel() {
		this.myDialog.close({next: false});
	}
}
