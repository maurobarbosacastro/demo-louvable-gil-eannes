import {Component, DestroyRef, effect, inject, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {PartnersService} from '@app/modules/admin/partners/service/partners.service';
import {filter, map, of, Subject, switchMap, takeUntil} from 'rxjs';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {PartnersInterface} from '@app/modules/admin/partners/interfaces/partners.interface';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {HttpErrorResponse} from '@angular/common/http';

@Component({
	selector: 'tagpeak-edit-partner',
	standalone: true,
	imports: [CommonModule, FormsModule, MatIcon, TranslocoPipe, ReactiveFormsModule],
	templateUrl: './edit-partner.component.html',
})
export class EditPartnerComponent {
	private dialogData = inject(MAT_DIALOG_DATA);
	private dialogRef = inject(MatDialogRef);
	private partnerService = inject(PartnersService);
	private toaster = inject(ToasterService);
	private transloco = inject(TranslocoService);
	private destroyRef = inject(DestroyRef);

	partnerForm: FormGroup = new FormGroup({
		name: new FormControl('', [Validators.required]),
		code: new FormControl('', [Validators.required]),
		eCommercePlatform: new FormControl('', [Validators.required]),
		commissionRate: new FormControl(0, [
			Validators.min(0),
			Validators.max(100),
			Validators.required,
		]),
		validationPeriod: new FormControl('', [Validators.min(0), Validators.required]),
		deepLink: new FormControl('', [Validators.required]),
		deepLinkIdentifier: new FormControl('', [Validators.required]),
		subIdentifier: new FormControl('', [Validators.required]),
		percentageTagpeak: new FormControl(0, [
			Validators.min(0),
			Validators.max(100),
			Validators.required,
		]),
		percentageInvested: new FormControl(0, [
			Validators.min(0),
			Validators.max(100),
			Validators.required,
		]),
	});

	$editPartner: Signal<boolean> = toSignal(
		of(this.dialogData).pipe(
			filter((data: any) => !!data.partnerUuid),
			map((data: any) => data.partnerUuid),
			switchMap((partnerUuid: string) => this.partnerService.getPartner(partnerUuid)),
			map((partner) => {
				this.partnerForm.patchValue(partner);
				Object.keys(this.partnerForm.controls).forEach((key) => {
					this.partnerForm.controls[key].removeValidators(Validators.required);
					this.partnerForm.controls[key].updateValueAndValidity();
				});
				return true;
			}),
		),
	);

	_ = effect(() => {
		if (!this.$editPartner()) {
			this.partnerForm
				.get('name')
				.valueChanges.pipe(takeUntil(this.killIt$), takeUntilDestroyed(this.destroyRef))
				.subscribe((value) => {
					if (value) {
						this.partnerForm.get('code').setValue(value.toLowerCase().replace(/ /g, '_'));
					}
				});
		} else {
			this.killIt$.next();
		}
	});

	killIt$: Subject<void> = new Subject();

	closeModal() {
		this.dialogRef.close({skipFetch: true});
	}

	savePartner() {
		if (this.partnerForm.invalid) {
			this.partnerForm.markAllAsTouched();
			return;
		}

		let partnerDto: PartnersInterface = this.partnerForm.value;
		if (this.$editPartner()) {
			this.partnerService
				.updatePartner(this.dialogData.partnerUuid, partnerDto)
				.subscribe((partner) => {
					if (partner) {
						this.toaster.showToast('success', this.transloco.translate('partner.partner-updated'));
						this.dialogRef.close({skipFetch: false});
					} else {
						this.toaster.showToast(
							'error',
							this.transloco.translate('partner.partner-update-error'),
						);
					}
				}, (error: HttpErrorResponse) => {
                    this.toaster.showToast('error', error.error.ErrorMessage);
                });
		} else {
			this.partnerService.createPartner(partnerDto).subscribe((partner) => {
				if (partner) {
					this.toaster.showToast('success', this.transloco.translate('partner.partner-created'));
					this.dialogRef.close({skipFetch: false});
				} else {
					this.toaster.showToast('error', this.transloco.translate('partner.partner-create-error'));
				}
			}, (error: HttpErrorResponse) => {
                if (error.status === 409) {
                    this.toaster.showToast('error', error.error.ErrorMessage);
                }

                this.toaster.showToast('error', error.error.ErrorMessage);
            });
		}
	}

	get nameControl(): FormControl {
		return this.partnerForm.get('name') as FormControl;
	}

	get codeControl(): FormControl {
		return this.partnerForm.get('code') as FormControl;
	}

	get eCommercePlatformControl(): FormControl {
		return this.partnerForm.get('eCommercePlatform') as FormControl;
	}

	get commissionRateControl(): FormControl {
		return this.partnerForm.get('commissionRate') as FormControl;
	}

	get validationPeriodControl(): FormControl {
		return this.partnerForm.get('validationPeriod') as FormControl;
	}

	get deepLinkControl(): FormControl {
		return this.partnerForm.get('deepLink') as FormControl;
	}

	get deepLinkIdentifierControl(): FormControl {
		return this.partnerForm.get('deepLinkIdentifier') as FormControl;
	}

	get subIdentifierControl(): FormControl {
		return this.partnerForm.get('subIdentifier') as FormControl;
	}

	get percentageTagpeakControl(): FormControl {
		return this.partnerForm.get('percentageTagpeak') as FormControl;
	}

	get percentageInvestedControl(): FormControl {
		return this.partnerForm.get('percentageInvested') as FormControl;
	}

	getCommissionRateErrorMessage(): string {
		if (!this.commissionRateControl.errors) {
			return '';
		}

		if (this.commissionRateControl.hasError('required')) {
			return 'partner.commissionRate-required';
		}

		if (this.commissionRateControl.hasError('min')) {
			return 'partner.commissionRate-min';
		}
		if (this.commissionRateControl.hasError('max')) {
			return 'partner.commissionRate-max';
		}

		return 'partner.commissionRate-invalid';
	}

	getPercentageTagpeakErrorMessage(): string {
		if (!this.percentageTagpeakControl.errors) {
			return '';
		}

		if (this.percentageTagpeakControl.hasError('required')) {
			return 'partner.percentageTagpeak-required';
		}

		if (this.percentageTagpeakControl.hasError('min')) {
			return 'partner.percentageTagpeak-min';
		}
		if (this.percentageTagpeakControl.hasError('max')) {
			return 'partner.percentageTagpeak-max';
		}

		return 'partner.percentageTagpeak-invalid';
	}

	getPercentageInvestedErrorMessage(): string {
		if (!this.percentageInvestedControl.errors) {
			return '';
		}

		if (this.percentageInvestedControl.hasError('required')) {
			return 'partner.percentageInvested-required';
		}

		if (this.percentageInvestedControl.hasError('min')) {
			return 'partner.percentageInvested-min';
		}
		if (this.percentageInvestedControl.hasError('max')) {
			return 'partner.percentageInvested-max';
		}

		return 'partner.percentageInvested-invalid';
	}
}
