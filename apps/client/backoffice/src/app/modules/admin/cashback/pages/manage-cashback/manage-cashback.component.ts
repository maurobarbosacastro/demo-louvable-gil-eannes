import {
	Component,
	computed,
	DestroyRef,
	ElementRef,
	inject,
	OnInit,
	Signal,
	signal,
	viewChild,
	viewChildren,
	WritableSignal,
} from '@angular/core';
import {CommonModule} from '@angular/common';
import {statusFilters} from '@app/utils/constants';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {StatusContainerComponent} from '@app/shared/components/status-container/status-container.component';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {CashbackService} from '@app/modules/admin/cashback/services/cashback.service';
import {
	BehaviorSubject,
	combineLatest,
	debounceTime,
	distinctUntilChanged,
	filter,
	map,
	Observable,
	startWith,
	switchMap,
	tap,
} from 'rxjs';
import {
	BulkEditRewardDto,
	CashbackDto,
	CashbackTableResponse,
	ManageCashbackSearchParams,
} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {TableComponent} from '@app/shared/components/table/table.component';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {CashbackDetailsComponent} from '@app/modules/admin/cashback/components/cashback-details/cashback-details.component';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {DeleteModalConfirmationComponent} from '@app/shared/components/delete-modal-confirmation/delete-modal-confirmation.component';
import {TableModalComponent} from '@app/modules/admin/cashback/components/table-modal/table-modal.component';
import {BulkUpdateComponent} from '@app/modules/admin/cashback/components/bulk-update/bulk-update.component';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {StatusEnum} from '@app/utils/status.enum';
import {MatTooltip} from '@angular/material/tooltip';
import {CdkCopyToClipboard} from '@angular/cdk/clipboard';
import {showHistoryBtn} from '@app/modules/admin/cashback/utils/cashback.utils';
import {AppConstants} from '@app/app.constants';
import _ from 'lodash';
import {PageEvent} from '@angular/material/paginator';
import {CashbackSortingFields} from '@app/modules/admin/cashback/models/cashback.enum';
import {CustomDropdownComponent} from '@app/shared/components/custom-dropdown/custom-dropdown.component';
import {
	CssDropdownInterface,
	IconDropdownInterface,
	IconInputDropdownInterface,
	OptionDropdownInterface,
} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {StoresService} from '@app/modules/admin/stores/services/stores.service';
import {AdminStoreInterface} from '@app/modules/admin/stores/interfaces/store.interface';
import {StoreVisitsService} from '@app/modules/client/store-visits/services/store-visits.service';
import {StoreVisitsInterface} from '@app/modules/client/store-visits/models/store-visits.interface';
import {MatDatepicker, MatDatepickerInput} from '@angular/material/datepicker';
import {UserService} from '@app/core/user/user.service';
import {TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {UrlQueryService} from '@app/shared/services/url-query.service';
import {DashboardService} from '@app/modules/admin/dashboard/services/dashboard.service';
import {TransactionsStatusInterface} from '@app/modules/admin/dashboard/models/dashboard.interface';
import {MiniDashboardComponent} from '@app/modules/admin/dashboard/components/mini-dashboard/mini-dashboard.component';
import {environment} from '@environments/environment';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-manage-cashback',
	standalone: true,
	imports: [
		CommonModule,
		ReactiveFormsModule,
		MatIcon,
		TranslocoPipe,
		StatusContainerComponent,
		TableComponent,
		BulkUpdateComponent,
		MatTooltip,
		CustomDropdownComponent,
		MatDatepicker,
		MatDatepickerInput,
		MiniDashboardComponent,
	],
	templateUrl: './manage-cashback.component.html',
	styleUrls: [
		'./manage-cashback.component.scss',
		'../../../../../../styles/style-table.scss',
		'../../../../../../styles/style-date-picker.scss',
	],
})
export class ManageCashbackComponent implements OnInit {
	private _queryService: UrlQueryService = inject(UrlQueryService);
	private readonly _dashboardService: DashboardService = inject(DashboardService);
	private readonly _screenService: ScreenService = inject(ScreenService);

	filtersStatus: {label: string; value: string}[] = statusFilters
		.map((status) => {
			if (status.value === StatusEnum.REJECTED) {
				return {...status, label: 'status.CANCELLED', value: StatusEnum.REJECTED};
			}
			return status;
		})
		.sort((a, b) => a.label.localeCompare(b.label));

	filtersForm = new FormGroup({
		status: new FormControl<string[]>([]),
		searchStoreVisit: new FormControl(null),
		searchStore: new FormControl(null),
		searchUser: new FormControl<string>(null),
		startDate: new FormControl(null),
		endDate: new FormControl(null),
	});

	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);
	$dashboardTransactions: Signal<TransactionsStatusInterface> = toSignal(
		this._dashboardService.getTransactionsDashboard(),
	);

	$cashbacksSelected: WritableSignal<CashbackTableResponse[]> = signal([]);
	$cashbackSelectedAllSameStatus: Signal<boolean> = computed(
		() =>
			_.uniqBy(this.$cashbacksSelected(), (cb) => (cb.reward ? cb.reward.state : cb.status))
				.length === 1,
	);

	tableConfiguration: ITableConfiguration;
	dataSource: MatTableDataSource<CashbackTableResponse> =
		new MatTableDataSource<CashbackTableResponse>();
	page$: BehaviorSubject<number> = new BehaviorSubject<number>(1);
	currentPage: number = 1;
	sort: string = 'date desc';
	pageSize: number = 100; //AppConstants.PAGE_SIZE;
	totalSize: number = 0;
	protected readonly StatusEnum = StatusEnum;

	pageFilters$: BehaviorSubject<number> = new BehaviorSubject<number>(1);
	pageSizeFilters$: BehaviorSubject<number> = new BehaviorSubject<number>(30);

	showHistoryBtn = showHistoryBtn;

	$checkboxes = viewChildren<ElementRef>('checkbox');
	allCheckboxesSelected: boolean = false;
	$headerCheckbox = viewChild<ElementRef>('checkboxHeader');

	$storesAvailable: Signal<AdminStoreInterface[]> = toSignal(this.getStores());
	$storeVisits: Signal<StoreVisitsInterface[]> = toSignal(this.getStoreVisits());
	$users: Signal<TagpeakUser[]> = toSignal(this.getUsers());

	styleDropdown: CssDropdownInterface = AppConstants.DEFAULT_STYLE_DROPDOWN;
	iconCheckOption: IconDropdownInterface = AppConstants.DEFAULT_ICON_CHECK_DROPDOWN;
	iconInput: IconInputDropdownInterface = {
		iconName: 'search',
		position: 'left',
	};

	$lockedButtonBulk: WritableSignal<boolean> = signal(false);

	constructor(
		private readonly cashbackService: CashbackService,
		private readonly translocoService: TranslocoService,
		private readonly matDialog: MatDialog,
		private readonly toasterService: ToasterService,
		private readonly destroyRef: DestroyRef,
		private readonly _storeService: StoresService,
		private readonly _storeVisitService: StoreVisitsService,
		private readonly _userService: UserService,
	) {}

	ngOnInit(): void {
		if (this._queryService.getParam('status')) {
			this.setFilters(this._queryService.getParam('status'));
			this.fetchData(this.setBodyParams());
		} else {
			this.fetchData();
		}
		this.setupTable();

		this.startDateControl.valueChanges.subscribe(() => {
			this.fetchData(this.setBodyParams());
		});

		this.endDateControl.valueChanges.subscribe(() => {
			this.fetchData(this.setBodyParams());
		});
	}

	fetchData(filtersParams?: Partial<ManageCashbackSearchParams>): void {
		this.cashbackService
			.getCashback(this.currentPage - 1, this.pageSize, this.sort, filtersParams)
			.pipe(
				tap(
					(cashbacks: PaginationInterface<CashbackTableResponse>) =>
						(this.totalSize = cashbacks.totalRows),
				),
				map((cashbacks: PaginationInterface<CashbackTableResponse>) => cashbacks.data),
				tap((cashbacks: CashbackTableResponse[]) =>
					cashbacks.length === 0 ? this.$cashbacksSelected.update(() => []) : null,
				),
				takeUntilDestroyed(this.destroyRef),
			)
			.subscribe((cashbacks: CashbackTableResponse[]) => {
				this.dataSource.data = cashbacks.map((c) => ({
					...c,
					reward: Object.keys(c.reward).length > 0 ? c.reward : null,
				}));
			});
	}

	getStores(): Observable<AdminStoreInterface[]> {
		const sort = new BehaviorSubject<string>('name asc');
		return combineLatest([
			sort,
			this.pageSizeFilters$,
			this.pageFilters$,
			this.filtersForm.controls.searchStore.valueChanges.pipe(
				startWith(null),
				takeUntilDestroyed(this.destroyRef),
				debounceTime(500),
				distinctUntilChanged(),
				map((value: any) => value?.label ?? value),
			),
		]).pipe(
			switchMap(([sort, size, page, name]: [string, number, number, string]) =>
				this._storeService.getAllStores(size, page, sort, name),
			),
			map((res: PaginationInterface<AdminStoreInterface>) => res.data ?? []),
		);
	}

	getStoreVisits(): Observable<StoreVisitsInterface[]> {
		return combineLatest([
			this.pageSizeFilters$,
			this.pageFilters$,
			this.searchStoreVisitsForm.valueChanges.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				takeUntilDestroyed(this.destroyRef),
				map((value: any) => value?.label ?? value),
			),
		]).pipe(
			switchMap(([size, page, reference]: [number, number, string]) =>
				this._storeVisitService.getStoreVisitsAdmin(page, size, '', {
					reference: reference,
				}),
			),
			map((res: PaginationInterface<StoreVisitsInterface>) => res.data ?? []),
		);
	}

	getUsers(): Observable<TagpeakUser[]> {
		return combineLatest([
			this.pageSizeFilters$,
			this.searchUserForm.valueChanges.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				takeUntilDestroyed(this.destroyRef),
				map((value: any) => value?.label ?? value),
			),
		]).pipe(
			switchMap(([size, searchUser]: [number, string]) =>
				this._userService.getAllUsers(size, searchUser),
			),
			map((res: TagpeakUser[]) => res ?? []),
		);
	}

	setupTable(): void {
		this.tableConfiguration = {
			dataSource: this.dataSource,
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
					sticky: true,
				},
				{
					id: 'exitIdEncoded',
					name: this.translocoService.translate('cashback.table.source-id'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'exitId',
					name: this.translocoService.translate('cashback.table.exit-id'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'store',
					name: this.translocoService.translate('cashback.table.store-name'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'user',
					name: this.translocoService.translate('cashback.table.user-email'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'date',
					name: this.translocoService.translate('cashback.table.date'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'title',
					name: this.translocoService.translate('cashback.table.title'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'amountTarget',
					name: this.translocoService.translate('cashback.table.order-value'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'networkCommission',
					name: this.translocoService.translate('cashback.table.network-commission'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'currentRewardTarget',
					name: this.translocoService.translate('cashback.table.current-reward'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'priceDayZero',
					name: this.translocoService.translate('cashback.table.price-day-0'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'isin',
					name: this.translocoService.translate('cashback.table.isin'),
					hasSort: true,
					hasTooltip: false,
					colWidth: '8%',
				},
				{
					id: 'status',
					name: this.translocoService.translate('cashback.table.status'),
					hasSort: true,
					hasTooltip: false,
				},
				// The actions column has this name so it can be customized in the bodyCustomTemplate
				{
					id: 'customActions',
					name: this.translocoService.translate('cashback.table.actions'),
					hasSort: false,
					hasTooltip: false,
				},
			],
		};
	}

	$mappedDropdownStores: Signal<OptionDropdownInterface[]> = computed(() =>
		this.$storesAvailable()
			?.map((store: AdminStoreInterface) => ({
				label: store.name,
				value: store.uuid,
			}))
			.sort((a, b) => a.label.localeCompare(b.label)),
	);

	$mappedDropdownStoreVisits: Signal<OptionDropdownInterface[]> = computed(() =>
		this.$storeVisits()
			?.map((storeVisit: StoreVisitsInterface) => ({
				label: storeVisit.reference?.toString(),
				value: storeVisit.uuid,
			}))
			.sort((a, b) => a.label.toString().localeCompare(b.label.toString())),
	);

	$mappedDropdownUsers: Signal<OptionDropdownInterface[]> = computed(() =>
		this.$users()?.map((user: TagpeakUser) => ({
			label: `${user.firstName} ${user.lastName}`,
			value: user.uuid,
		})),
	);

	setFilters(filter: string, fromWarning: boolean = false): void {
		let filters: string[] = !fromWarning
			? Array.from(new Set([...this.filtersForm.controls.status.value, filter]))
			: [filter];

		this.filtersForm.controls.status.patchValue(filters);
		this.fetchData(this.setBodyParams());
	}

	removeFilter(filter: string): void {
		if (filter === 'CONFIRMED') {
			filter = 'VALIDATED';
		}
		const filters: string[] = this.filtersForm.controls.status.value.filter(
			(value: string) => value !== filter,
		);
		this.filtersForm.controls.status.patchValue(filters);
		this.fetchData(this.setBodyParams());
	}

	setBodyParams(): Partial<ManageCashbackSearchParams> {
		return {
			storeUuid: this.searchStoreForm.value?.value ?? null,
			userUuid: this.searchUserForm.value?.value ?? null,
			storeVisitUuid: this.searchStoreVisitsForm.value?.value ?? null,
			state: this.filtersForm.value.status.join(',') ?? null,
			startDate: this.startDateControl.value ?? null,
			endDate: this.endDateControl.value ?? null,
		};
	}

	historyCashback(element: any): void {
		this.matDialog
			.open(TableModalComponent, {
				panelClass: 'dialog-panel',
				height: 'fit-content',
				width: '35%',
				data: {uuid: element.row.reward.uuid},
			})
			.afterClosed();
	}

	editCashback(element: any): void {
		const modalConfig: MatDialogConfig = {
			...AppConstants.MODAL_CONFIG,
			data: {uuid: element.row.uuid},
		};

		if (this.$isMobile()) {
			modalConfig.width = '100%';
			modalConfig.height = '95%';
			modalConfig.panelClass = 'dialog-panel-details-mobile';
		}

		this.matDialog
			.open(CashbackDetailsComponent, modalConfig)
			.afterClosed()
			.subscribe({
				next: (result: any) => {
					if (!result.skipFetch) {
						this.fetchData();
					}
				},
			});
	}

	deleteCashback(element: any): void {
		this.matDialog
			.open(DeleteModalConfirmationComponent, {
				panelClass: 'dialog-panel',
				height: 'fit-content',
				width: '26rem',
				disableClose: true,
			})
			.afterClosed()
			.pipe(
				filter((result: {skipFetch: boolean}) => result && !result.skipFetch),
				switchMap(() => this.cashbackService.deleteCashback(element.row.uuid)),
			)
			.subscribe({
				next: () => {
					this.handleSuccessRequest('cashback.success-delete');
				},
			});
	}

	handleSuccessRequest(message: string): void {
		this.toasterService.showToast('success', this.translocoService.translate(message));
		this.fetchData();
		this.clearCashbacksSelected();
		this.$lockedButtonBulk.set(false)
	}

	handleErrorRequest(message: string): void {
		this.toasterService.showToast('error', this.translocoService.translate(message));
		this.$lockedButtonBulk.set(false)
	}

	selectCashback(element: CashbackTableResponse): void {
		if (this.$cashbacksSelected().some((cashback) => cashback.uuid === element.uuid)) {
			this.$cashbacksSelected.update((selected) => [
				...selected.filter((sel) => sel.uuid !== element.uuid),
			]);
			if (this.allCheckboxesSelected) {
				this.$headerCheckbox().nativeElement.checked = false;
			}
		} else {
			this.$cashbacksSelected.update((selected) => [...selected, element]);
			if (this.$cashbacksSelected().length === this.dataSource.data.length) {
				this.allCheckboxesSelected = true;
				this.$headerCheckbox().nativeElement.checked = true;
			}
		}
	}

	updateBulk(body: Partial<CashbackDto>): void {
		this.$lockedButtonBulk.set(true)

		// Transaction list to update (Only Status from Pending to Validated or Rejected)
		if (body.status === StatusEnum.VALIDATED || body.status === StatusEnum.REJECTED) {
			// update transaction status (Only status)
			this.cashbackService
				.updateCashbackBulkTransaction({
					uuids: this.$cashbacksSelected().map((selected) => selected.uuid),
					state: body.status,
				})
				.subscribe({
					next: () => {
						this.handleSuccessRequest('cashback.success-update');
					},
					error: (error) => {
						if (error) {
							this.handleErrorRequest('cashback.error-update');
						}
					},
				});
		} else {
			//Check if transaction have reward
			const rewardsToCreate = this.$cashbacksSelected()
				.filter((selected) => !selected?.reward?.uuid)
				.map((selected) => selected.uuid);

			const obj: Partial<BulkEditRewardDto> = {
				initialPrice: Number(body.priceDayZero) === 0 ? undefined : Number(body.priceDayZero),
				endDate: body.endDate === null || body.endDate === '' ? undefined : body.endDate,
				initialDate:
					body.initialDate === null || body.initialDate === '' ? undefined : body.initialDate,
				status: body.status === null || body.status === '' ? undefined : body.status,
				isin: body.isin === null || body.isin === '' ? undefined : body.isin,
				conid: body.conid === null || body.conid === '' ? undefined : body.conid,
				overridePrice: Number(body.overridePrice) === 0 ? undefined : Number(body.overridePrice),
			};

			//IF some transactions have reward create the reward with the fields
			if (rewardsToCreate.length > 0) {
				const dtoObj: Partial<BulkEditRewardDto> = {
					transactionUuids: rewardsToCreate,
					...obj,
				};

				this.cashbackService.createCashbackRewardBulk(dtoObj).subscribe({
					next: () => {
						this.handleSuccessRequest('cashback.success-update');
					},
					error: (error) => {
						if (error) {
							this.handleErrorRequest('cashback.error-update');
						}
					},
				});
			}
			const uuids: string[] = this.$cashbacksSelected().map((selected) => selected.reward.uuid);
			const dtoObj: Partial<BulkEditRewardDto> = {
				uuids,
				...obj,
			};

			this.cashbackService.updateCashbackBulkReward(dtoObj).subscribe({
				next: () => {
					this.handleSuccessRequest('cashback.success-update');
				},
				error: (error) => {
					if (error) {
						this.handleErrorRequest('cashback.error-update');
					}
				},
			});
		}
	}

	validateCheckboxValue(uuid: string): boolean {
		return this.$cashbacksSelected().some((selected) => selected.uuid === uuid);
	}

	clearCashbacksSelected(): void {
		this.$cashbacksSelected.update(() => []);
		this.$checkboxes().every((cb) => (cb.nativeElement.checked = false));
	}

	totalPages(): number {
		return Math.ceil(this.totalSize / this.pageSize);
	}

	informCopy() {
		this.toasterService.showToast('info', 'Copied to clipboard');
	}

	selectAllVisibleCashback() {
		if (this.$checkboxes().every((cb) => cb.nativeElement.checked)) {
			this.$cashbacksSelected.update(() => []);
			this.$checkboxes().forEach((checkbox) => (checkbox.nativeElement.checked = false));
			this.allCheckboxesSelected = false;
		} else {
			this.$cashbacksSelected.update(() => this.dataSource.data);
			this.$checkboxes().forEach((checkbox) => (checkbox.nativeElement.checked = true));
			this.allCheckboxesSelected = true;
		}
	}

	onSortAndPageChange(event: {sort: string; page: PageEvent}) {
		if (event.sort) {
			const field = CashbackSortingFields[event.sort.split(',')[0]];
			const direction = event.sort.split(',')[1];
			this.sort = `${field} ${direction}`;
		}
		this.currentPage = event.page.pageIndex;
		this.page$.next(event.page.pageIndex);
		this.fetchData(this.setBodyParams());
	}

	setFormControlName(formControlName: string, value: OptionDropdownInterface): void {
		if (value) {
			this.filtersForm.get(formControlName).setValue(value);
		} else {
			this.filtersForm.get(formControlName).setValue(null);
		}
		this.fetchData(this.setBodyParams());
	}

	resetFormControl(formControlName: string): void {
		this.filtersForm.get(formControlName).reset();
	}

	getCountryStore(storeId: string): string {
		let store: AdminStoreInterface = this.$storesAvailable().find(
			(store: AdminStoreInterface) => store.uuid === storeId,
		);
		return store.country.slice(0, store.country.length > 3 ? 3 : store.country.length).join(', ');
	}

	getEmailUser(userId: string): string {
		let user: TagpeakUser = this.$users().find((user: TagpeakUser) => user.uuid === userId);
		return user.email;
	}

	navigateToCashback(columnId: string): void {
		this._queryService.addParam('status', columnId);
		this.setFilters(columnId, true);
		this.fetchData(this.setBodyParams());
	}

	encodeStoreVisitRef(ref: string, userUuid: string): string {
		if (!ref || !userUuid) {
			return ''
		}

		if (ref.toUpperCase().includes('TP')) {
			return btoa(ref + '_' + userUuid.split('-')[0]);
		}

		return ref;
	}

	get statusControl(): FormControl {
		return this.filtersForm.get('status') as FormControl;
	}

	get searchStoreForm(): FormControl {
		return this.filtersForm.get('searchStore') as FormControl;
	}

	get searchUserForm(): FormControl {
		return this.filtersForm.get('searchUser') as FormControl;
	}

	get searchStoreVisitsForm(): FormControl {
		return this.filtersForm.get('searchStoreVisit') as FormControl;
	}

	get startDateControl(): FormControl {
		return this.filtersForm.get('startDate') as FormControl;
	}

	get endDateControl(): FormControl {
		return this.filtersForm.get('endDate') as FormControl;
	}

	protected readonly environment = environment;
}
