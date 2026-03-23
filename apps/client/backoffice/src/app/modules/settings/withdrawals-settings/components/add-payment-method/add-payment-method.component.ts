import {
	Component,
	computed,
	DestroyRef,
	OnInit,
	Signal,
	signal,
	WritableSignal,
} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {MatDialogRef} from '@angular/material/dialog';
import {TranslocoPipe} from '@ngneat/transloco';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {CountriesService} from '@app/modules/admin/countries/services/countries.service';
import {CountryInterface} from '@app/modules/admin/countries/interfaces/country.interface';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {PaymentMethodsService} from '@app/modules/settings/withdrawals-settings/services/payment-methods.service';
import {
	AvailablePaymentMethod,
	CreateUserPaymentMethod,
	UserPaymentMethodInterface,
} from '@app/modules/settings/withdrawals-settings/models/user-payment-method.interface';
import {debounceTime, distinctUntilChanged, filter, switchMap, tap} from 'rxjs';

@Component({
	selector: 'tagpeak-add-payment-method',
	standalone: true,
	imports: [CommonModule, MatIcon, TranslocoPipe, ReactiveFormsModule],
	templateUrl: './add-payment-method.component.html',
	styleUrl: './add-payment-method.component.scss',
})
export class AddPaymentMethodComponent implements OnInit {
	$countries: WritableSignal<CountryInterface[]> = signal([]);
	$paymentMethods: WritableSignal<AvailablePaymentMethod[]> = signal([]);

	paymentMethodForm = new FormGroup({
		paymentMethod: new FormControl(null, [Validators.required]),
		bankName: new FormControl('', [Validators.required]),
		bankAddress: new FormControl('', [Validators.required]),
		bankCountry: new FormControl('', [Validators.required]),
		country: new FormControl('', [Validators.required]),
		bankAccountTitle: new FormControl('', [Validators.required]),
		iban: new FormControl('', [
			Validators.required,
			Validators.minLength(15),
			Validators.maxLength(34),
			Validators.pattern(/^[A-Z0-9\s]+$/),
		]),
		ibanStatement: new FormControl(null, [Validators.required]),
		vat: new FormControl('', [Validators.required]),
	});

	ibanStatement: File;

	$validateVat: WritableSignal<{activated: boolean; done: boolean; valid: boolean}> = signal({
		activated: false,
		done: false,
		valid: false,
	});

	constructor(
		private readonly _myDialog: MatDialogRef<AddPaymentMethodComponent, {skipFetch: boolean}>,
		private readonly _countriesService: CountriesService,
		private readonly _paymentMethodsService: PaymentMethodsService,
		private readonly _destroyRef: DestroyRef,
		private readonly _toaster: ToasterService,
	) {}

	ngOnInit() {
		this.paymentMethodControl.valueChanges.subscribe((value: string) => {
			const method: AvailablePaymentMethod = this.$paymentMethods().find(
				(method: AvailablePaymentMethod) => method.uuid === value,
			);
			if (method) {
				this.validatePayment(method);
			}
		});

		this.fetchData();

		//Keep commented for now.
		/*this.vatNumberControl.valueChanges
            .pipe(
                tap( () => {
                    if (!!this.vatNumberControl.errors) {
                        this.$validateVat.update( d => ({...d, activated: false, valid: false}));
                    }
                }),
                debounceTime(1000),
                distinctUntilChanged(),
                takeUntilDestroyed(this._destroyRef),
                filter((res) => !!res.match(/^[A-Za-z]{2}\d+$/)),
                tap((res) => this.$validateVat.update((d) => ({...d, activated: true}))),
                switchMap((res) => this._paymentMethodsService.checkVatValidity(res)),
            )
            .subscribe((res) => {
                this.$validateVat.update( d => ({...d, done: true, valid: res}));
            });*/
	}

	handleSubmitFile(event: any): void {
		const file = event.target.files[0] as File;
		this.ibanStatement = file;
	}

	fetchData(): void {
		this._countriesService
			.getCountries(100, 0, 'name')
			.pipe(takeUntilDestroyed(this._destroyRef))
			.subscribe((counties: PaginationInterface<CountryInterface>) =>
				this.$countries.set(counties.data as CountryInterface[]),
			);

		this._paymentMethodsService
			.getPaymentMethodsAvailable()
			.subscribe((paymentMethod: AvailablePaymentMethod[]) =>
				this.$paymentMethods.set(paymentMethod),
			);
	}

	closeModal(): void {
		this._myDialog.close({skipFetch: true});
	}

	clearFile(): void {
		this.ibanStatementControl.reset();
		this.ibanStatement = null;
	}

	$isBankSelected(): Signal<boolean> {
		return computed(() =>
			this.$paymentMethods().some(
				(method: AvailablePaymentMethod) =>
					method.uuid === this.paymentMethodControl.value && method.code === 'bank',
			),
		);
	}

	validatePayment(value: AvailablePaymentMethod): void {
		if (value.code !== 'bank') {
			this.removeFormControlRequired(this.bankNameControl);
			this.removeFormControlRequired(this.bankAddressControl);
			this.removeFormControlRequired(this.bankCountryControl);
			this.removeFormControlRequired(this.bankAccountTitleControl);
			this.removeFormControlRequired(this.ibanControl);
			this.removeFormControlRequired(this.ibanStatementControl);
		}
	}

	removeFormControlRequired(formControl: FormControl): void {
		formControl.removeValidators(Validators.required);
		formControl.updateValueAndValidity();
	}

	finalize(): void {
		if (this.paymentMethodForm.invalid) {
			this.paymentMethodForm.markAllAsTouched();
			return;
		}

		let body: CreateUserPaymentMethod = {
			...this.paymentMethodForm.value,
			iban: this.ibanControl.value.replace(/\s/g, ''),
		} as CreateUserPaymentMethod;

		this._paymentMethodsService
			.uploadIbanStatement(this.ibanStatement)
			.pipe(
				switchMap((fileId: string) => {
					const updatedBody: CreateUserPaymentMethod = {
						...body,
						ibanStatement: fileId,
					};
					return this._paymentMethodsService.createPaymentMethod(updatedBody);
				}),
			)
			.subscribe({
				next: () => {
					this._myDialog.close({skipFetch: false});
					this._toaster.showToast('success', 'Payment method added successfully.');
				},
				error: (err) => {
					console.log(err);
					this._toaster.showToast('error', 'Error creating the payment method.');
				},
			});
	}

	onIbanInput(event: Event) {
		const input = event.target as HTMLInputElement;
		let value: string = input.value.replace(/\s/g, '').toUpperCase(); // Remove spaces and convert to uppercase

		// Keep only letters and numbers
		value = value.replace(/[^A-Z0-9]/g, '');

		// Format with spaces every 4 characters
		let formattedValue: string = '';
		for (let i: number = 0; i < value.length; i++) {
			if (i > 0 && i % 4 === 0) {
				formattedValue += ' ';
			}
			formattedValue += value[i];
		}

		this.ibanControl.setValue(formattedValue);
	}

	errorMessageIban(): string {
		if (this.ibanControl.hasError('required') && !this.ibanControl.untouched) {
			return 'store.field-required';
		}
		if (this.ibanControl.hasError('pattern') && !this.ibanControl.untouched) {
			return 'withdrawal-settings.iban-pattern-error';
		}
		if (this.ibanControl.hasError('minlength') && !this.ibanControl.untouched) {
			return 'withdrawal-settings.iban-error-min-length';
		}
		if (this.ibanControl.hasError('maxlength') && !this.ibanControl.untouched) {
			return 'withdrawal-settings.iban-error-max-length';
		}
		return '';
	}

	get paymentMethodControl(): FormControl {
		return this.paymentMethodForm.controls.paymentMethod as FormControl;
	}

	get bankNameControl(): FormControl {
		return this.paymentMethodForm.controls.bankName as FormControl;
	}

	get bankAddressControl(): FormControl {
		return this.paymentMethodForm.controls.bankAddress as FormControl;
	}

	get bankCountryControl(): FormControl {
		return this.paymentMethodForm.controls.bankCountry as FormControl;
	}

	get countyControl(): FormControl {
		return this.paymentMethodForm.controls.country as FormControl;
	}

	get bankAccountTitleControl(): FormControl {
		return this.paymentMethodForm.controls.bankAccountTitle as FormControl;
	}

	get ibanControl(): FormControl {
		return this.paymentMethodForm.controls.iban as FormControl;
	}

	get ibanStatementControl(): FormControl {
		return this.paymentMethodForm.controls.ibanStatement as FormControl;
	}

	get vatNumberControl(): FormControl {
		return this.paymentMethodForm.controls.vat as FormControl;
	}
}
