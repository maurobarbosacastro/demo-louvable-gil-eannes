import {Component, HostListener, Inject} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TranslocoPipe} from '@ngneat/transloco';
import {MAT_DIALOG_DATA, MatDialogClose, MatDialogRef} from '@angular/material/dialog';
import {MatIcon} from '@angular/material/icon';

@Component({
	selector: 'tagpeak-dialog-info',
	standalone: true,
	imports: [CommonModule, TranslocoPipe, MatDialogClose, MatIcon],
	template: `
        <div class="flex flex-col gap-8 my-2">

            <div class="flex justify-center my-4">
                <mat-icon svgIcon="info-warning" class="scale-[3] text-blue-200"></mat-icon>
            </div>

            <div class="flex flex-col justify-center items-center text-center">
                <span class="text-project-waterloo text-xl">{{ receivedData.text}}</span>
            </div>

            <div class="flex flex-row justify-center items-center gap-2.5">
                <button class="btn border-project-licorice-900 bg-white text-project-licorice-900 text-base font-normal min-h-9 h-9 rounded" mat-dialog-close>
                    @if (receivedData.buttonText) {
                        {{ receivedData.buttonText }}
                    } @else {
                        {{ 'misc.got-it' | transloco }}
                    }
                </button>
            </div>
        </div>
    `,
})
export class DialogInfoComponent {
	constructor(
		@Inject(MAT_DIALOG_DATA)
		protected readonly receivedData: {text: string; buttonText?: string},
        private readonly myDialog: MatDialogRef<DialogInfoComponent>,
	) {}

    @HostListener('keydown.enter', ['$event'])
    onEnterPress(event: KeyboardEvent) {
        event.preventDefault();
        // Your logic here
        this.myDialog.close();
    }
}
