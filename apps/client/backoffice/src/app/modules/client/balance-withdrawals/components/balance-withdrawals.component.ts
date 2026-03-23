import {Component, computed, DestroyRef, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {MatDatepicker, MatDatepickerInput} from '@angular/material/datepicker';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {BalanceWithdrawalsService} from '@app/modules/client/balance-withdrawals/services/balance-withdrawals.service';
import {TableComponent} from '@app/shared/components/table/table.component';
import {StatusContainerComponent} from '@app/shared/components/status-container/status-container.component';
import {
	BalanceInterface,
	WithdrawalsInterface,
	WithdrawalsSearchParams,
	WithdrawalsSortingFields,
} from '@app/modules/client/balance-withdrawals/models/withdrawals.interface';
import {BehaviorSubject, combineLatest, map, switchMap, tap} from 'rxjs';
import {InfoCardComponent} from '@app/shared/components/info-card/info-card.component';
import {AppConstants} from '@app/app.constants';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {UserService} from '@app/core/user/user.service';
import {MatTooltip} from '@angular/material/tooltip';
import {Router} from '@angular/router';
import {formatISO} from 'date-fns';
import {DialogInfoComponent} from '@app/shared/components/dialog-info/dialog-info.component';
import {MatDialog} from '@angular/material/dialog';
import {CdkCopyToClipboard} from '@angular/cdk/clipboard';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {HttpErrorResponse} from '@angular/common/http';
import {ConfigurationService} from '@app/modules/admin/configuration/services/configuration.service';
import {PageEvent} from '@angular/material/paginator';
import {ConfirmationModalComponent} from '@app/shared/components/confirmation-modal/confirmation-modal.component';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-balance-withdrawals',
	standalone: true,
	imports: [
		CommonModule,
		ReactiveFormsModule,
		MatDatepicker,
		MatDatepickerInput,
		MatIcon,
		TranslocoPipe,
		TableComponent,
		StatusContainerComponent,
		InfoCardComponent,
		MatTooltip,
		CdkCopyToClipboard,
	],
	templateUrl: './balance-withdrawals.component.html',
	styleUrls: [
		'./balance-withdrawals.component.scss',
		'../../../../../styles/style-table.scss',
		'../../../../../styles/style-date-picker.scss',
	],
})
export class BalanceWithdrawalsComponent {
	filtersForm = new FormGroup({
		startDate: new FormControl(null),
		endDate: new FormControl(null),
	});

	$user: Signal<TagpeakUser> = toSignal(this.userService.get());
	$balance: Signal<BalanceInterface> = toSignal(this.withdrawalsService.getBalance());
	$canWithdraw: Signal<boolean> = computed(
		() =>
			this.$user() &&
			this.$user().balance > Number(this.configService.configs['withdrawal_balance_minimum'].value),
	);
	$valueToWithdraw: Signal<string> = computed(
		() => this.configService.configs['withdrawal_balance_minimum'].value,
	);
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	page$: BehaviorSubject<number> = new BehaviorSubject(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject('created_at desc');
	filters$: BehaviorSubject<Partial<WithdrawalsSearchParams>> = new BehaviorSubject({});
	pageSize: number = AppConstants.PAGE_SIZE;
	totalSize: number = 1;
	page: number;
	totalPages: number;

	$dataNew: Signal<WithdrawalsInterface[]> = toSignal(
		combineLatest([this.sort$, this.page$, this.filters$]).pipe(
			takeUntilDestroyed(this.destroyRef),
			switchMap(([sort, page, filters]) => {
				return this.withdrawalsService.getMyWithdrawals(page, this.pageSize, sort, filters);
			}),
			tap((res) => {
				this.page = res.page;
				this.totalPages = res.totalPages;
				this.totalSize = res.totalRows;
			}),
			map((res) => res.data ?? []),
		),
	);

	$tableConfig: Signal<ITableConfiguration> = computed(() => {
		return {
			dataSource: new MatTableDataSource(this.$dataNew()),
			pageSize: this.pageSize,
			styles: {
				header: 'font-nunitoBlack text-project-waterloo',
				content: 'font-nunitoRegular text-project-licorice-900',
				paginator: 'font-nunitoRegular text-black',
			},
			columns: [
				{
					id: 'uuid',
					name: this.translocoService.translate('balance-withdrawals.table.id'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'amountSource',
					name: this.translocoService.translate('balance-withdrawals.table.cash-out-amount'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'state',
					name: this.translocoService.translate('balance-withdrawals.table.status'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'completionDate',
					name: this.translocoService.translate('balance-withdrawals.table.payment-date'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'details',
					name: this.translocoService.translate('balance-withdrawals.table.payment-note'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'createdAt',
					name: this.translocoService.translate('balance-withdrawals.table.date-requested'),
					hasSort: true,
					hasTooltip: false,
				},
			],
		};
	});

	constructor(
		private readonly withdrawalsService: BalanceWithdrawalsService,
		private readonly translocoService: TranslocoService,
		private readonly userService: UserService,
		private readonly toasterService: ToasterService,
		private destroyRef: DestroyRef,
		private router: Router,
		private dialog: MatDialog,
		private readonly configService: ConfigurationService,
		private _screenService: ScreenService,
	) {}

	searchByFilter(): void {
		const bodyParams: Partial<WithdrawalsSearchParams> = {
			...(this.filtersForm.value.startDate && {
				startDate: formatISO(new Date(this.filtersForm.value.startDate)),
			}),
			...(this.filtersForm.value.endDate && {
				endDate: formatISO(new Date(this.filtersForm.value.endDate)),
			}),
		};
		this.filters$.next(bodyParams);
	}

	resetFormControl(formControlName: string): void {
		this.filtersForm.get(formControlName).reset();
		this.searchByFilter();
	}

	goToSettings(): void {
		this.router.navigateByUrl('/client/settings?tab=withdrawal-settings');
	}

	createNewWithdrawal() {
		this.dialog
			.open(ConfirmationModalComponent, {
				data: {
					title: this.translocoService.translate(
						'balance-withdrawals.confirmation-withdrawal.title',
					),
					message: this.translocoService.translate(
						'balance-withdrawals.confirmation-withdrawal.message',
					),
				},
			})
			.afterClosed()
			.subscribe(({next}) => {
				if (next) {
					this.withdrawalsService
						.createNewWithdrawal()
						.pipe(
							switchMap((_) => {
								return this.dialog
									.open(DialogInfoComponent, {
										data: {
											text: this.translocoService.translate('balance-withdrawals.success-create'),
										},
									})
									.afterClosed();
							}),
						)
						.subscribe({
							next: (_) => {
								this.page$.next(1);
							},
							error: (err: HttpErrorResponse) => {
								if (err.status === 422 || err.status === 409 || err.status === 404) {
									this.toasterService.showToast('error', err.error);
									return;
								}

								this.toasterService.showToast(
									'error',
									this.translocoService.translate('balance-withdrawals.error-create'),
								);
							},
						});
				}
			});
	}

	informCopy() {
		this.toasterService.showToast('info', 'Copied to clipboard');
	}

	onSortAndPageChange(evt: {sort: string; page: PageEvent}) {
		if (evt.sort) {
			const field = WithdrawalsSortingFields[evt.sort.split(',')[0]];
			const direction = evt.sort.split(',')[1];
			this.sort$.next(`${field} ${direction}`);
			return;
		}
		this.page$.next(evt.page.pageIndex);
	}
}
