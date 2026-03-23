import {Component, inject} from '@angular/core';
import {CommonModule} from '@angular/common';
import {SharedModule} from '@app/shared/shared.module';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {StoresService} from '@app/modules/admin/stores/services/stores.service';

@Component({
	selector: 'tagpeak-upload-csv',
	standalone: true,
	imports: [CommonModule, SharedModule, TranslocoPipe],
	templateUrl: './upload-csv.component.html',
})
export class UploadCsvComponent {
	public data = inject(MAT_DIALOG_DATA);
	private dialogRef = inject(MatDialogRef);
	private _storesService = inject(StoresService);
	private _toaster = inject(ToasterService);
	private _translocoService = inject(TranslocoService);
	private file: File;

	error: any;

	onCancelClick() {
		this.dialogRef.close(false);
	}

	onSaveClick() {
		if (this.file) {
			this._storesService.uploadCSV(this.file).subscribe(
				(value) => {
					this._toaster.showToast('success', value.message, 'bottom');
					this.dialogRef.close(true);
				},
				(error) => (this.error = error.error),
			);
		} else {
			this.error = {
				'message': this._translocoService.translate('store.insert-file'),
			};
		}
	}

	onFileChange($event: Event) {
		this.error = null;
		this.file = ($event.target as HTMLInputElement).files[0];
	}
}
