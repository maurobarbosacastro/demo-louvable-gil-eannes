import {
	Component,
	computed,
	DestroyRef,
	inject,
	Signal,
	signal,
	WritableSignal,
} from '@angular/core';
import {CommonModule, CurrencyPipe, DatePipe} from '@angular/common';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {CashoutRequestForm} from '@app/modules/admin/cashout/models/cashout-request.enum';
import {MatDatepicker, MatDatepickerInput} from '@angular/material/datepicker';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {WithdrawalsService} from '@app/modules/admin/cashout/services/withdrawals.service';
import {PaymentMethodsService} from '@app/modules/settings/withdrawals-settings/services/payment-methods.service';
import {
	UpdateStateRequestDTO,
	WithdrawalsRequestInterface,
} from '@app/modules/admin/cashout/models/withdrawals.interface';
import {takeUntilDestroyed, toObservable, toSignal} from '@angular/core/rxjs-interop';
import {filter, map, switchMap, tap} from 'rxjs';
import {StatusEnum} from '@app/utils/status.enum';
import {CustomDropdownComponent} from '@app/shared/components/custom-dropdown/custom-dropdown.component';
import {
	CssDropdownInterface,
	IconDropdownInterface,
	OptionDropdownInterface,
} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import {AppConstants} from '@app/app.constants';

@Component({
	selector: 'tagpeak-cashout-request-details',
	standalone: true,
	imports: [
		CommonModule,
		MatIcon,
		TranslocoPipe,
		ReactiveFormsModule,
		MatDatepicker,
		MatDatepickerInput,
		CustomDropdownComponent,
	],
	templateUrl: './cashout-request-details.component.html',
	styleUrl: '../../../../../../styles/style-date-picker.scss',
	providers: [DatePipe, CurrencyPipe],
})
export class CashoutRequestDetailsComponent {
	private readonly _myDialog: MatDialogRef<CashoutRequestDetailsComponent, {skipFetch: boolean}> =
		inject(MatDialogRef);
	private readonly _dialogData = inject(MAT_DIALOG_DATA);
	private readonly _destroyRef: DestroyRef = inject(DestroyRef);
	private readonly _withdrawalService: WithdrawalsService = inject(WithdrawalsService);
	private readonly _toaster: ToasterService = inject(ToasterService);
	private readonly _transloco: TranslocoService = inject(TranslocoService);
	private readonly _paymentMethodService: PaymentMethodsService = inject(PaymentMethodsService);
	private readonly _datePipe: DatePipe = inject(DatePipe);
	private readonly _currencyPipe: CurrencyPipe = inject(CurrencyPipe);

	status: {label: string; value: string}[] = [
		{
			label: 'status.STOPPED',
			value: StatusEnum.REJECTED,
		},
		{
			label: 'status.PENDING',
			value: StatusEnum.PENDING,
		},
		{
			label: 'status.PAID',
			value: StatusEnum.COMPLETED,
		},
	];

	cashoutForm: FormGroup[] = [
		new FormGroup({
			paymentMethod: new FormControl(''),
			bankName: new FormControl(''),
			bankAddress: new FormControl(''),
			bankAccountTitle: new FormControl(''),
			iban: new FormControl(''),
			country: new FormControl(''),
			bankCountry: new FormControl(''),
			vat: new FormControl(''),
		}),
		new FormGroup({
			cashoutValue: new FormControl(null),
			status: new FormControl(''),
			requestDate: new FormControl('', [Validators.required]),
			paymentDate: new FormControl('', [Validators.required]),
			note: new FormControl(''),
		}),
	];

	$haveRequest: WritableSignal<boolean> = signal(this._dialogData && !!this._dialogData.uuid);
	$withdrawalRequest: Signal<WithdrawalsRequestInterface> = toSignal(
		toObservable(this.$haveRequest).pipe(
			takeUntilDestroyed(this._destroyRef),
			filter((val: boolean) => !!val),
			switchMap(() => this._withdrawalService.getWithdrawalRequest(this._dialogData.uuid)),
			map((withdrawalRequest: WithdrawalsRequestInterface) => withdrawalRequest),
			tap((withdrawalRequest: WithdrawalsRequestInterface) => this.setupForm(withdrawalRequest)),
		),
	);

	styleDropdown: CssDropdownInterface = {
		...AppConstants.DEFAULT_STYLE_DROPDOWN,
		styleHeader: 'bg-white border-b border-project-licorice-900 pb-2',
		styleInput: 'text-project-licorice-500',
	};

	iconCheckOption: IconDropdownInterface = AppConstants.DEFAULT_ICON_CHECK_DROPDOWN;

	$mappedStatus: Signal<OptionDropdownInterface[]> = computed(() =>
		this.status
			.map((status: {label: string; value: string}) => ({
				label: this._transloco.translate(status.label),
				value: status.value,
			}))
			.sort((a, b) => a.value.localeCompare(b.value)),
	);

	setupForm(withdrawalRequest: WithdrawalsRequestInterface): void {
		this.cashoutForm[CashoutRequestForm.PAYMENT_DETAILS].setValue({
			paymentMethod: withdrawalRequest.paymentMethod.paymentMethod,
			bankName: withdrawalRequest.paymentMethod.bankName,
			bankAddress: withdrawalRequest.paymentMethod.bankAddress,
			bankAccountTitle: withdrawalRequest.paymentMethod.bankAccountTitle,
			iban: this.ibanFormat(withdrawalRequest.paymentMethod.iban),
			bankCountry: withdrawalRequest.paymentMethod.bankCountry,
			country: withdrawalRequest.paymentMethod.country,
			vat: withdrawalRequest.paymentMethod.vat,
		});

		this.cashoutForm[CashoutRequestForm.CASHOUT_REQUEST_INFO].patchValue({
			cashoutValue: this._currencyPipe.transform(
				withdrawalRequest.amountTarget,
				withdrawalRequest.currencyTarget,
				'symbol-narrow',
				'1.2-2',
			),
			status: this.getStatus(withdrawalRequest.state),
			requestDate: this._datePipe.transform(withdrawalRequest.createdAt, 'dd/MM/yyyy HH:mm'),
			paymentDate: withdrawalRequest?.completionDate
				? withdrawalRequest?.completionDate
				: new Date(),
			note: withdrawalRequest.details,
		});
	}

	setFormControlValue(formPosition: number, formControlName: string, value: string | any): void {
		this.cashoutForm[formPosition].get(formControlName).setValue(value);
	}

	requestAction(state: StatusEnum): void {
		const body: UpdateStateRequestDTO = this.setupCashoutRequestDto(state);

		this._withdrawalService.updateStateRequest(body, this.$withdrawalRequest().uuid).subscribe({
			next: () => {
				this._toaster.showToast(
					'success',
					state === StatusEnum.COMPLETED
						? this._transloco.translate('withdrawals-request.accept-success')
						: this._transloco.translate('withdrawals-request.reject-success'),
				);
				this._myDialog.close({skipFetch: false});
			},
			error: () => {
				this._toaster.showToast(
					'error',
					state === StatusEnum.COMPLETED
						? this._transloco.translate('withdrawals-request.error-accept')
						: this._transloco.translate('withdrawals-request.error-reject'),
				);
			},
		});
	}

	setupCashoutRequestDto(state: StatusEnum): UpdateStateRequestDTO {
		return {
			state: state ? state : this.statusControl.value,
			completionDate: new Date(this.paymentDateControl.value).toISOString(),
			details: this.noteControl.value,
		};
	}

	ibanFormat(iban: string): string {
		let formattedValue: string = '';
		for (let i: number = 0; i < iban?.length; i++) {
			if (i > 0 && i % 4 === 0) {
				formattedValue += ' ';
			}
			formattedValue += iban[i];
		}
		return formattedValue;
	}

	closeModal(): void {
		this._myDialog.close({skipFetch: true});
	}

	downloadFile(): void {
		this._paymentMethodService
			.downloadFile(this.$withdrawalRequest().paymentMethod.ibanStatement['uuid'])
			.subscribe((file: File) => {
				window.open(URL.createObjectURL(file));
			});
	}

	getStatus(value: string): string {
		value &&= value.toUpperCase();

		if (value === StatusEnum.COMPLETED) {
			return this._transloco.translate('status.' + StatusEnum.PAID);
		}

		if (value === StatusEnum.REJECTED) {
			return this._transloco.translate('status.' + StatusEnum.STOPPED);
		}

		if (value === StatusEnum.PENDING) {
			return 'PENDING';
		}
		return value ? this._transloco.translate('status.' + value) : null;
	}

	isPendingState(): boolean {
		return this.$withdrawalRequest().state === StatusEnum.PENDING;
	}

	get paymentMethodControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.PAYMENT_DETAILS].get('paymentMethod') as FormControl;
	}

	get bankNameControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.PAYMENT_DETAILS].get('bankName') as FormControl;
	}

	get bankAddressControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.PAYMENT_DETAILS].get('bankAddress') as FormControl;
	}

	get bankAccountTitleControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.PAYMENT_DETAILS].get(
			'bankAccountTitle',
		) as FormControl;
	}

	get ibanControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.PAYMENT_DETAILS].get('iban') as FormControl;
	}

	get bankCountryControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.PAYMENT_DETAILS].get('bankCountry') as FormControl;
	}

	get countryControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.PAYMENT_DETAILS].get('country') as FormControl;
	}

	get vatControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.PAYMENT_DETAILS].get('vat') as FormControl;
	}

	get cashoutValueControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.CASHOUT_REQUEST_INFO].get(
			'cashoutValue',
		) as FormControl;
	}

	get statusControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.CASHOUT_REQUEST_INFO].get('status') as FormControl;
	}

	get requestDateControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.CASHOUT_REQUEST_INFO].get(
			'requestDate',
		) as FormControl;
	}

	get paymentDateControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.CASHOUT_REQUEST_INFO].get(
			'paymentDate',
		) as FormControl;
	}

	get noteControl(): FormControl {
		return this.cashoutForm[CashoutRequestForm.CASHOUT_REQUEST_INFO].get('note') as FormControl;
	}

	protected readonly CashoutRequestForm = CashoutRequestForm;
	protected readonly StatusEnum = StatusEnum;
}
