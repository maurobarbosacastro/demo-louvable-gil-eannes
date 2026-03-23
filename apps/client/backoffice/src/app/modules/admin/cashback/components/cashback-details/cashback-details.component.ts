import {Component, DestroyRef, Inject, OnInit, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {statusFilters} from '@app/utils/constants';
import {CashbackService} from '@app/modules/admin/cashback/services/cashback.service';
import {
	CashbackResponse,
	CreateRewardDTO,
	RewardResponse,
	RewardUpdateDTO,
	TransactionUpdateDTO,
} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {MatDatepicker, MatDatepickerInput} from '@angular/material/datepicker';
import {catchError, map, of, switchMap, throwError} from 'rxjs';
import {StoreVisitsService} from '@app/modules/client/store-visits/services/store-visits.service';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {StatusEnum, transactionStatuses} from '@app/utils/status.enum';
import _ from 'lodash';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {TagpeakInputDirective} from '@app/shared/directive/tagpeak-input.directive';
import {ReadonlyInputDirective} from '@app/shared/directive/readonly-input.directive';
import {HttpErrorResponse} from '@angular/common/http';

@Component({
	selector: 'tagpeak-cashback-details',
	standalone: true,
	imports: [
		CommonModule,
		MatIcon,
		ReactiveFormsModule,
		TranslocoPipe,
		MatDatepicker,
		MatDatepickerInput,
		TagpeakInputDirective,
		ReadonlyInputDirective,
	],
	templateUrl: './cashback-details.component.html',
	styleUrls: [
		'./cashback-details.component.scss',
		'../../../../../../styles/style-date-picker.scss',
	],
})
export class CashbackDetailsComponent implements OnInit {
	cashbackForm = new FormGroup({
		exitClick: new FormControl(null, [Validators.required]),
		currency: new FormControl('EUR'),
		priceDayZero: new FormControl(null),
		initialDate: new FormControl(''),
		endDate: new FormControl(''),
		isin: new FormControl(''),
		conid: new FormControl(''),
		title: new FormControl('', [Validators.required]),
		details: new FormControl(''),
		orderValue: new FormControl(null, [Validators.required]),
		networkCommission: new FormControl(null, [Validators.required]),
		cashback: new FormControl(null, [Validators.required]),
		orderDate: new FormControl('', [Validators.required]),
		status: new FormControl(''),
		overridePrice: new FormControl(null),
		unit: new FormControl(null),
		currentCashReward: new FormControl(null),
	});

	status: {label: string; value: string}[] = statusFilters
		.map((status) => {
			if (status.value === StatusEnum.REJECTED) {
				return {...status, label: 'status.CANCELLED', value: StatusEnum.CANCELLED};
			}
			return status;
		})
		.sort((a, b) => a.label.localeCompare(b.label));
	$cashback: WritableSignal<CashbackResponse> = signal(null);
	$reward: WritableSignal<RewardResponse> = signal(null);
	$clickSubmit: WritableSignal<boolean> = signal(true);

	$validate: WritableSignal<{activated: boolean; done: boolean; valid: boolean}> = signal({
		activated: false,
		done: false,
		valid: false,
	});
	transactionUuid: string;

	constructor(
		private readonly myDialog: MatDialogRef<CashbackDetailsComponent, {skipFetch: boolean}>,
		@Inject(MAT_DIALOG_DATA) private readonly receivedData: {uuid: string},
		private readonly cashbackService: CashbackService,
		private readonly storeVisitService: StoreVisitsService,
		private destroyRef: DestroyRef,
		private readonly translocoService: TranslocoService,
		private readonly toasterService: ToasterService,
	) {}

	ngOnInit(): void {
		this.transactionUuid = this.receivedData.uuid;
		this.cashbackService
			.getCashbackById(this.transactionUuid)
			.pipe(
				takeUntilDestroyed(this.destroyRef),
				switchMap((cashback) =>
					this.cashbackService.getCashbackReward(this.transactionUuid, true).pipe(
						map((reward) => ({cashback, reward})),
						catchError((_) => {
							return of({cashback, reward: null});
						}),
					),
				),
			)
			.subscribe((resp: {cashback: CashbackResponse; reward: RewardResponse}) => {
				const {cashback, reward} = resp;

				this.$cashback.set(cashback);
				this.$reward.set(reward);

				this.cashbackForm.patchValue({
					exitClick: cashback.storeVisit ? cashback.storeVisit.reference : null,
					currency: cashback.currencySource,
					orderValue: cashback.amountTarget,
					networkCommission: cashback.commissionTarget,
					orderDate: cashback.orderDate,
					status: this.getStatusValue(cashback.state),
					cashback: cashback.cashback,
				});

				if (reward) {
					this.cashbackForm.patchValue({
						priceDayZero: reward?.initialPrice,
						initialDate: reward?.createdAt,
						endDate: reward?.endDate,
						isin: reward?.isin,
						conid: reward?.conid,
						title: reward?.title,
						details: reward?.details,
						overridePrice: reward?.overridePrice,
						unit: reward?.assetUnits,
						currentCashReward: reward?.currentRewardTarget,
						status: reward?.state,
					});
				}
			});
	}

	getStatusValue(status: string): string {
		if (status === StatusEnum.VALIDATED) {
			return StatusEnum.CONFIRMED;
		} else if (status === StatusEnum.REJECTED) {
			return StatusEnum.CANCELLED;
		}
		return status;
	}

	closeModal(): void {
		this.myDialog.close({skipFetch: true});
	}

	setFormControlValue(formControlName: string, value: string | any): void {
		this.cashbackForm.get(formControlName).patchValue(value);
		this.cashbackForm.get(formControlName).markAsDirty();
		this.cashbackForm.get(formControlName).markAsTouched();
	}

	editCashback(): void {
		if (!this.$clickSubmit()) {
			return;
		}

		this.$clickSubmit.set(false);
		const {transaction, reward} = this.setupBodyDto();

		this.cashbackService
			.updateCashback(this.transactionUuid, transaction)
			.pipe(
				switchMap((_) => {
					if (this.$reward()) {
						return this.cashbackService.updateCashbackReward(this.$reward().uuid, reward).pipe(
							switchMap((_) => this.cashbackService.verifyReward(this.$reward().uuid)),
							catchError((err: HttpErrorResponse) => {
								if (err.status === 409) {
									return of(null);
								}
								return throwError(() => err);
							}),
						);
					}

					if (!this.$reward() && this.cashbackForm.value.status === StatusEnum.LIVE) {
						return this.cashbackService.createCashbackReward(reward as CreateRewardDTO);
					}

					return of(null);
				}),
			)
			.subscribe({
				next: (_) => {
					this.handleSuccessRequest('cashback.success-update');
					this.myDialog.close({skipFetch: false});
					this.$clickSubmit.set(true);
				},
				error: (_) => {
					this.handleErrorRequest('cashback.error-update');
					this.$clickSubmit.set(true);
				},
			});
	}

	validateIt(): void {
		this.$validate.set({activated: true, done: false, valid: false});
		if (
			!this.cashbackForm.get('exitClick').value ||
			this.cashbackForm.get('exitClick').value === ''
		) {
			return;
		}

		this.storeVisitService
			.validateReference(this.cashbackForm.get('exitClick').value)
			.pipe(
				takeUntilDestroyed(this.destroyRef),
				catchError((_) => of(false)),
			)
			.subscribe((res) => {
				this.$validate.set({activated: true, done: true, valid: res});
			});
	}

	setupBodyDto(): {
		transaction: Partial<TransactionUpdateDTO>;
		reward: Partial<RewardUpdateDTO> | CreateRewardDTO;
	} {
		const transaction: Partial<TransactionUpdateDTO> = {
			orderDate: this.isControlEdited('orderDate') ? this.cashbackForm.value.orderDate : null,
			commissionTarget: this.isControlEdited('networkCommission')
				? Number(this.cashbackForm.value.networkCommission)
				: null,
			currencySource: this.isControlEdited('currency') ? this.cashbackForm.value.currency : null,
			exitClick: this.isControlEdited('exitClick') ? this.cashbackForm.value.exitClick : null,
		};

		let reward: Partial<RewardUpdateDTO> = {};
		if (this.$reward()) {
			reward = {
				currentRewardSource: this.isControlEdited('overridePrice')
					? Number(this.cashbackForm.value.overridePrice)
					: null,
				endDate: this.isControlEdited('endDate') ? this.cashbackForm.value.endDate : null,
				initialPrice: this.cashbackForm.get('priceDayZero').dirty
					? Number(this.cashbackForm.value.priceDayZero)
					: null,
				isin: this.isControlEdited('isin') ? this.cashbackForm.value.isin : null,
				conid: this.isControlEdited('conid') ? this.cashbackForm.value.conid : null,
				initialDate: this.isControlEdited('initialDate')
					? this.cashbackForm.value.initialDate
					: null,
				title: this.isControlEdited('title') ? this.cashbackForm.value.title : null,
				details: this.isControlEdited('details') ? this.cashbackForm.value.details : null,
				overridePrice: this.isControlEdited('overridePrice')
					? Number(this.cashbackForm.value.overridePrice)
					: null,
			} as Partial<RewardUpdateDTO>;
		} else {
			if (this.cashbackForm.value.status === StatusEnum.LIVE) {
				reward = {
					currency: 'EUR',
					details: this.cashbackForm.get('details').value,
					endDate: this.cashbackForm.get('endDate').value,
					initialDate: this.cashbackForm.get('initialDate').value,
					initialPrice: Number(this.cashbackForm.get('priceDayZero').value),
					isin: this.cashbackForm.get('isin').value,
					conid: this.cashbackForm.get('conid').value,
					title: this.cashbackForm.get('title').value,
					transactionUuid: this.$cashback().uuid,
					type: 'INVESTMENT',
				} as CreateRewardDTO;
			}
		}

		if (this.cashbackForm.get('status').dirty) {
			if (transactionStatuses.includes(this.cashbackForm.value.status as StatusEnum)) {
				transaction.state = this.cashbackForm.value.status;
			} else {
				reward.state = this.cashbackForm.value.status;
			}
		} else {
			transaction.state = null;
			reward.state = null;
		}

		return {
			transaction: _.omitBy(transaction, _.isNull),
			reward: _.omitBy(reward, _.isNull),
		};
	}

	handleSuccessRequest(message: string): void {
		this.toasterService.showToast('success', this.translocoService.translate(message), 'bottom');
	}

	handleErrorRequest(message: string): void {
		this.toasterService.showToast('error', this.translocoService.translate(message), 'bottom');
	}

	isControlEdited(name: string) {
		return this.cashbackForm.get(name).dirty && this.cashbackForm.get(name).touched;
	}

	get exitClickControl(): FormControl {
		return this.cashbackForm.get('exitClick') as FormControl;
	}

	get currencyControl(): FormControl {
		return this.cashbackForm.get('currency') as FormControl;
	}

	get priceDayZeroControl(): FormControl {
		return this.cashbackForm.get('priceDayZero') as FormControl;
	}

	get initialDateControl(): FormControl {
		return this.cashbackForm.get('initialDate') as FormControl;
	}

	get endDateControl(): FormControl {
		return this.cashbackForm.get('endDate') as FormControl;
	}

	get isinControl(): FormControl {
		return this.cashbackForm.get('isin') as FormControl;
	}

	get conidControl(): FormControl {
		return this.cashbackForm.get('conid') as FormControl;
	}

	get titleControl(): FormControl {
		return this.cashbackForm.get('title') as FormControl;
	}

	get detailsControl(): FormControl {
		return this.cashbackForm.get('details') as FormControl;
	}

	get orderValueControl(): FormControl {
		return this.cashbackForm.get('orderValue') as FormControl;
	}

	get networkCommissionControl(): FormControl {
		return this.cashbackForm.get('networkCommission') as FormControl;
	}

	get cashbackControl(): FormControl {
		return this.cashbackForm.get('cashback') as FormControl;
	}

	get orderDateControl(): FormControl {
		return this.cashbackForm.get('orderDate') as FormControl;
	}

	get statusControl(): FormControl {
		return this.cashbackForm.get('status') as FormControl;
	}

	get overridePriceControl(): FormControl {
		return this.cashbackForm.get('overridePrice') as FormControl;
	}

	get unitControl(): FormControl {
		return this.cashbackForm.get('unit') as FormControl;
	}

	get currentCashRewardControl(): FormControl {
		return this.cashbackForm.get('currentCashReward') as FormControl;
	}
}
