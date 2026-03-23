import {Component, Inject, OnInit, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {SharedModule} from '@app/shared/shared.module';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {ModalActionResult} from '@app/utils/modal-action-result.enum';
import {LanguagesService} from '@app/modules/admin/languages/services/languages.service';
import {LanguagesInterface} from '@app/modules/admin/languages/models/languages.interface';
import {MatIcon} from '@angular/material/icon';
import {toSignal} from '@angular/core/rxjs-interop';
import {filter, map, of, switchMap} from 'rxjs';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {HttpErrorResponse} from '@angular/common/http';

@Component({
	selector: 'tagpeak-language-details',
	standalone: true,
	imports: [CommonModule, ReactiveFormsModule, SharedModule, TranslocoPipe, MatIcon],
	templateUrl: './language-details.component.html',
})
export class LanguageDetailsComponent {
	languageForm = new FormGroup({
		code: new FormControl('', [Validators.required]),
		name: new FormControl('', [Validators.required]),
	});

	$editLanguage: Signal<boolean> = toSignal(
		of(this.receivedData).pipe(
			filter((data: {id: string}) => !!data?.id),
			map((data: {id: string}) => data?.id),
			switchMap((id: string) => this.languageService.getLanguageById(id)),
			map((language: LanguagesInterface) => {
				this.setupForm(language);
				return true;
			}),
		),
	);

	constructor(
		private readonly myDialog: MatDialogRef<LanguageDetailsComponent, {skipFetch: boolean}>,
		@Inject(MAT_DIALOG_DATA)
		private readonly receivedData: {id: string},
		private readonly languageService: LanguagesService,
		private readonly toasterService: ToasterService,
		private readonly translocoService: TranslocoService,
	) {}

	private setupForm(language: LanguagesInterface): void {
		this.languageForm.patchValue({
			code: language.code,
			name: language.name,
		});
	}

	closeModal(): void {
		this.myDialog.close({skipFetch: true});
	}

	finalize(): void {
		if (this.languageForm.invalid) {
			this.languageForm.markAllAsTouched();
			return;
		}

		const body: Partial<LanguagesInterface> = {
			code: this.languageForm.value.code,
			name: this.languageForm.value.name,
		};

		if (this.$editLanguage()) {
			this.languageService.updateLanguage(body, this.receivedData.id).subscribe({
				next: () => {
					this.toasterService.showToast(
						'success',
						this.translocoService.translate('languages.success-update'),
					);
					this.myDialog.close({skipFetch: false});
				},
				error: (error: HttpErrorResponse) => {
					this.toasterService.showToast('error', error.error.ErrorMessage);
				},
			});
		} else {
			this.languageService.createLanguage(body).subscribe({
				next: () => {
					this.toasterService.showToast(
						'success',
						this.translocoService.translate('languages.success-create'),
					);
					this.myDialog.close({skipFetch: false});
				},
				error: (error: HttpErrorResponse) => {
					this.toasterService.showToast('error', error.error.ErrorMessage);
				},
			});
		}
	}

	get nameControl(): FormControl {
		return this.languageForm.get('name') as FormControl;
	}

	get codeControl(): FormControl {
		return this.languageForm.get('code') as FormControl;
	}
}
