import {
	Component,
	computed,
	DestroyRef,
	inject,
	signal,
	Signal,
	WritableSignal,
} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {TableComponent} from '@app/shared/components/table/table.component';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {
	BehaviorSubject,
	combineLatest,
	debounceTime,
	distinctUntilChanged,
	filter,
	map,
	startWith,
	switchMap,
	tap,
} from 'rxjs';
import {StatusContainerComponent} from '@app/shared/components/status-container/status-container.component';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {AppConstants} from '@app/app.constants';
import {MatDatepicker, MatDatepickerInput} from '@angular/material/datepicker';
import {CustomDropdownComponent} from '@app/shared/components/custom-dropdown/custom-dropdown.component';
import {
	CssDropdownInterface,
	IconDropdownInterface,
	OptionDropdownInterface,
} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import {StatusEnum} from '@app/utils/status.enum';
import {WithdrawalsService} from '@app/modules/admin/cashout/services/withdrawals.service';
import {
	BulkUpdateStateRequestDTO,
	WithdrawalSortingFields,
	WithdrawalsRequestInterface,
} from '@app/modules/admin/cashout/models/withdrawals.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {PageEvent} from '@angular/material/paginator';
import {CashoutRequestDetailsComponent} from '@app/modules/admin/cashout/components/cashout-request-details/cashout-request-details.component';
import {MatTooltip} from '@angular/material/tooltip';
import {DeleteModalConfirmationComponent} from '@app/shared/components/delete-modal-confirmation/delete-modal-confirmation.component';
import {CdkCopyToClipboard} from '@angular/cdk/clipboard';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-cashout-request',
	standalone: true,
	imports: [
		CommonModule,
		ReactiveFormsModule,
		MatIcon,
		TranslocoPipe,
		TableComponent,
		StatusContainerComponent,
		MatDatepicker,
		CustomDropdownComponent,
		MatDatepickerInput,
		MatTooltip,
		CdkCopyToClipboard,
	],
	templateUrl: './cashout-request.component.html',
	styleUrls: [
		'../../../../../../styles/style-table.scss',
		'../../../../../../styles/style-date-picker.scss',
	],
})
export class CashoutRequestComponent {
	private readonly _withdrawalsService: WithdrawalsService = inject(WithdrawalsService);
	private readonly _transloco: TranslocoService = inject(TranslocoService);
	private readonly _matDialog: MatDialog = inject(MatDialog);
	private readonly _toaster: ToasterService = inject(ToasterService);
	private readonly _destroyRef: DestroyRef = inject(DestroyRef);
	private readonly _screenService: ScreenService = inject(ScreenService);

	filterForm = new FormGroup({
		status: new FormControl(null),
		startDate: new FormControl(null),
		endDate: new FormControl(null),
	});

	tableConfiguration: ITableConfiguration;
	page$: BehaviorSubject<number> = new BehaviorSubject<number>(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject<string>('created_at desc'); // Default sorting
	totalSize: number = 1;
	pageSize: number = AppConstants.PAGE_SIZE;
	page: number;
	total: number;

	statusAvailable: {label: string; value: string}[] = [
		{
			label: 'status.REJECTED',
			value: StatusEnum.REJECTED,
		},
		{
			label: 'Pending',
			value: StatusEnum.PENDING,
		},
		{
			label: 'status.PAID',
			value: StatusEnum.COMPLETED,
		},
	];

	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);
	$mappedStatus: Signal<OptionDropdownInterface[]> = computed(() =>
		this.statusAvailable
			.map((status: {label: string; value: string}) => ({
				label: this._transloco.translate(status.label),
				value: status.value,
			}))
			.sort((a, b) => a.value.localeCompare(b.value)),
	);

	styleDropdown: CssDropdownInterface = AppConstants.DEFAULT_STYLE_DROPDOWN;
	iconCheckOption: IconDropdownInterface = AppConstants.DEFAULT_ICON_CHECK_DROPDOWN;

	$withdrawals: Signal<Partial<WithdrawalsRequestInterface>[]> = toSignal(
		combineLatest([
			this.sort$,
			this.page$,
			this.statusControl.valueChanges.pipe(
				startWith(this.statusControl.value),
				debounceTime(500),
				distinctUntilChanged(),
			),
			this.startDateControl.valueChanges.pipe(
				startWith(this.startDateControl.value),
				debounceTime(500),
				distinctUntilChanged(),
			),
			this.endDateControl.valueChanges.pipe(
				startWith(this.endDateControl.value),
				debounceTime(500),
				distinctUntilChanged(),
			),
		]).pipe(
			switchMap(([sort, page, status, startDate, endDate]) => {
				return this._withdrawalsService.getAllWithdrawalsRequest(this.pageSize, page, sort, {
					state: status,
					startDate: startDate,
					endDate: endDate,
				});
			}),
			tap((res: PaginationInterface<Partial<WithdrawalsRequestInterface>>) => {
				this.page = res.page;
				this.total = res.totalPages;
				this.totalSize = res.totalRows;
			}),
			takeUntilDestroyed(this._destroyRef),
			map((res: PaginationInterface<Partial<WithdrawalsRequestInterface>>) => {
				const data: Partial<WithdrawalsRequestInterface>[] = res.data ?? [];
				return data.sort(
					(a: Partial<WithdrawalsRequestInterface>, b: Partial<WithdrawalsRequestInterface>) => {
						// Put PENDING states at the top
						if (a.state === 'PENDING' && b.state !== 'PENDING') return -1;
						if (a.state !== 'PENDING' && b.state === 'PENDING') return 1;
						return 0;
					},
				);
			}),
		),
	);

	$tableConfig: Signal<ITableConfiguration> = computed(() => this.setupTable());

	$withdrawalsSelected: WritableSignal<WithdrawalsRequestInterface[]> = signal([]);

	setupTable(): ITableConfiguration {
		return {
			dataSource: new MatTableDataSource(this.mapTable(this.$withdrawals())),
			css: 'rounded-t-[0.625rem] overflow-hidden shadow-table',
			pageSize: this.pageSize,
			styles: {
				header: 'font-plex text-project-waterloo text-md font-normal',
				content: 'font-plex text-project-licorice-900 text-md',
				paginator: 'font-plex text-black',
			},
			columns: [
				{
					id: 'checkbox',
					name: '',
					hasSort: false,
					hasTooltip: false,
					colWidth: '5%',
				},
				{
					id: 'uuid',
					name: 'ID',
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'user',
					name: 'User',
					hasSort: true,
					hasTooltip: false,
					colWidth: '25%',
				},
				{
					id: 'paymentMethod',
					name: 'Payment Method',
					hasSort: true,
					hasTooltip: false,
					colWidth: '15%',
				},
				{
					id: 'iban',
					name: 'IBAN',
					hasSort: true,
					hasTooltip: false,
					colWidth: '15%',
				},
				{
					id: 'completionDate',
					name: 'Payment Date',
					hasSort: true,
					hasTooltip: false,
					colWidth: '15%',
				},
				{
					id: 'note',
					name: 'Note',
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'state',
					name: 'Status',
					hasSort: true,
					hasTooltip: false,
					colWidth: '15%',
				},
				// The actions column has this name so it can be customized in the bodyCustomTemplate
				{
					id: 'customActions',
					name: 'Actions',
					hasSort: false,
					hasTooltip: false,
				},
			],
		};
	}

	mapTable(withdrawals: Partial<WithdrawalsRequestInterface>[]) {
		return withdrawals?.map((withdrawal: Partial<WithdrawalsRequestInterface>) => ({
			uuid: withdrawal.uuid,
			user: withdrawal.user.name,
			state: withdrawal.state,
			createdAt: withdrawal.createdAt,
			email: withdrawal.user.email,
			paymentMethod: withdrawal.paymentMethod.paymentMethod,
			iban: withdrawal.paymentMethod.iban,
			completionDate: withdrawal.completionDate,
			note: withdrawal.details,
			amountTarget: withdrawal.amountTarget,
		}));
	}

	onSortAndPageChange(evt: {sort: string; page: PageEvent}): void {
		if (evt.sort) {
			const field = WithdrawalSortingFields[evt.sort.split(',')[0]];
			const direction = evt.sort.split(',')[1];
			this.sort$.next(`${field} ${direction}`);
			return;
		}
		this.page$.next(evt.page.pageIndex);
	}

	editCashoutRequest(element: any): void {
		const modalConfig: MatDialogConfig = {
			...AppConstants.MODAL_CONFIG,
			data: {uuid: element.row.uuid},
		};

		if (this.$isMobile()) {
			modalConfig.width = '100%';
			modalConfig.height = '95%';
			modalConfig.panelClass = 'dialog-panel-details-mobile';
		}

		this._matDialog
			.open(CashoutRequestDetailsComponent, modalConfig)
			.afterClosed()
			.subscribe((result) => {
				if (!result.skipFetch) {
					this.page$.next(1);
				}
			});
	}

	deleteCashoutRequest(element: any): void {
		const data = {
			title: 'cashout-request.delete-cashout-request.title',
			body: 'cashout-request.delete-cashout-request.body',
		};

		this._matDialog
			.open(DeleteModalConfirmationComponent, {
				panelClass: 'dialog-panel',
				height: 'fit-content',
				width: '26rem',
				data: data,
			})
			.afterClosed()
			.pipe(
				filter((result: {skipFetch: boolean}) => result && !result.skipFetch),
				switchMap(() => this._withdrawalsService.deleteWithdrawalRequest(element.row.uuid)),
			)
			.subscribe({
				next: () => {
					this._toaster.showToast(
						'success',
						this._transloco.translate('cashout-request.success-delete'),
					);
					this.page$.next(1);
					this.clearWithdrawalsSelected();
				},
				error: (error) => {
					this._toaster.showToast(
						'error',
						this._transloco.translate('cashout-request.error-delete'),
					);
				},
			});
	}

	updateBulkWithdrawalRequest(state: StatusEnum): void {
		const body: BulkUpdateStateRequestDTO = {
			uuids: this.$withdrawalsSelected().map(
				(selected: WithdrawalsRequestInterface) => selected.uuid,
			),
			state: state,
			completionDate: new Date().toISOString(),
		};

		this._withdrawalsService.bulkUpdateStateRequest(body).subscribe({
			next: () => {
				this._toaster.showToast(
					'success',
					state === StatusEnum.COMPLETED
						? this._transloco.translate('withdrawals-request.accept-success')
						: this._transloco.translate('withdrawals-request.reject-success'),
				);
				this.page$.next(1);
				this.clearWithdrawalsSelected();
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

	selectWithdrawal(element: WithdrawalsRequestInterface): void {
		if (
			this.$withdrawalsSelected().some(
				(withdrawal: WithdrawalsRequestInterface) => withdrawal.uuid === element.uuid,
			)
		) {
			this.$withdrawalsSelected.update((selected: WithdrawalsRequestInterface[]) => [
				...selected.filter((sel: WithdrawalsRequestInterface) => sel.uuid !== element.uuid),
			]);
		} else {
			this.$withdrawalsSelected.update((selected: WithdrawalsRequestInterface[]) => [
				...selected,
				element,
			]);
		}
	}

	validateCheckboxValue(uuid: string): boolean {
		return this.$withdrawalsSelected().some(
			(selected: WithdrawalsRequestInterface) => selected.uuid === uuid,
		);
	}

	clearWithdrawalsSelected(): void {
		this.$withdrawalsSelected.update(() => []);
	}

	setFormControlName(formControlName: string, value: any): void {
		if (value) {
			this.filterForm.get(formControlName).setValue(value?.value);
		} else {
			this.filterForm.get(formControlName).setValue(null);
		}
	}

	resetFormControl(formControlName: string): void {
		this.filterForm.get(formControlName).reset();
	}

	getStatus(value: string): string {
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

	informCopy() {
		this._toaster.showToast('info', 'Copied to clipboard');
	}

	getAmountWithdrawal(): number {
		return this.$withdrawalsSelected().reduce(
			(acc: number, cur: Partial<WithdrawalsRequestInterface>) => acc + cur.amountTarget,
			0,
		);
	}

	get statusControl(): FormControl {
		return this.filterForm.controls.status as FormControl;
	}

	get startDateControl(): FormControl {
		return this.filterForm.controls.startDate as FormControl;
	}

	get endDateControl(): FormControl {
		return this.filterForm.controls.endDate as FormControl;
	}

	protected readonly StatusEnum = StatusEnum;
}
