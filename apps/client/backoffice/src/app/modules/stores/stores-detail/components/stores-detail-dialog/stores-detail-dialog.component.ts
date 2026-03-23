import {Component, inject, OnInit} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe} from '@ngneat/transloco';
import {FormControl, UntypedFormBuilder, UntypedFormGroup} from '@angular/forms';
import {SharedModule} from '@app/shared/shared.module';
import {ActivatedRoute, Router} from '@angular/router';
import {StorageService} from '@app/shared/services/storage.service';
import {AppConstants} from '@app/app.constants';

@Component({
	selector: 'tagpeak-stores-detail-dialog',
	standalone: true,
	imports: [CommonModule, MatIcon, TranslocoPipe, SharedModule],
	templateUrl: './stores-detail-dialog.component.html',
	styleUrl: './stores-detail-dialog.component.scss',
})
export class StoresDetailDialogComponent implements OnInit {
	private _dialogData = inject(MAT_DIALOG_DATA);
	private _dialogRef = inject(MatDialogRef);
	private _formBuilder = inject(UntypedFormBuilder);
	private _router = inject(Router);
	private _activatedRoute = inject(ActivatedRoute);
    private _storageService = inject(StorageService);

	dialogForm: UntypedFormGroup;

	ngOnInit(): void {
		this.dialogForm = this._formBuilder.group({
			checked: [false],
		});

        this.checkedControl.valueChanges
            .subscribe( res => {
                if (res){
                    this._storageService.save(AppConstants.STORAGE_KEYS.STORE_INFO, {
                        checked: true,
                        checkedDate: new Date(),
                    });
                } else {
                    this._storageService.delete(AppConstants.STORAGE_KEYS.STORE_INFO);
                }
            })
	}

	get checkedControl(): FormControl {
		return this.dialogForm.get('checked') as FormControl;
	}

	closeModal() {
        this._storageService.delete(AppConstants.STORAGE_KEYS.STORE_INFO);
		this._dialogRef.close();
	}

	goToStore(): void {
        const origin = new URL(window.location.href).origin;
        const internalRoute =  this._router.createUrlTree(['stores', this._dialogData.uuid, 'redirect'], {
            relativeTo: this._activatedRoute,
        })
        const url = origin + '/#' + this._router.serializeUrl(internalRoute);

        window.open(url, '_blank');

		this._dialogRef.close();
	}
}
