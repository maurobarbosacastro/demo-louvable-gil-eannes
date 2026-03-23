import {Component, DestroyRef, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ShopService} from '@app/modules/shop/services/shop.service';
import {CashbackShopTableResponse, DashboardStats} from '@app/modules/shop/interfaces/shop.interface';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {StatCardComponent} from '@app/modules/shop/components/stat-card.component';
import {TableComponent} from '@app/shared/components/table/table.component';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {ManageCashbackSearchParams} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {BehaviorSubject, map, tap} from 'rxjs';
import {AppConstants} from '@app/app.constants';
import {MatIcon} from '@angular/material/icon';
import {MatTooltip} from '@angular/material/tooltip';
import {StatusContainerComponent} from '@app/shared/components/status-container/status-container.component';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {PageEvent} from '@angular/material/paginator';
import {CashbackSortingFields} from '@app/modules/admin/cashback/models/cashback.enum';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {UrlQueryService} from '@app/shared/services/url-query.service';
import {statusFilters} from '@app/utils/constants';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {StatusEnum} from '@app/utils/status.enum';
import {environment} from '@environments/environment';
import {MatDatepicker, MatDatepickerInput} from '@angular/material/datepicker';
import {CdkCopyToClipboard} from '@angular/cdk/clipboard';

@Component({
	selector: 'tagpeak-empty',
	standalone: true,
	imports: [
		CommonModule,
		StatCardComponent,
		TableComponent,
		MatIcon,
		MatTooltip,
		StatusContainerComponent,
		TranslocoPipe,
		ReactiveFormsModule,
		MatDatepickerInput,
		MatDatepicker,
		CdkCopyToClipboard,
	],
	templateUrl: './shop-home.component.html',
	styleUrl: './shop-home.component.scss',
})
export class ShopHomeComponent {
	$stats: Signal<DashboardStats> = toSignal(
		this.shopService.getShopStats(this.shopService.$currentShop().shopUuid),
	);

	filtersStatus: {label: string; value: string}[] = statusFilters
		.map((status) => {
			if (status.value === StatusEnum.REJECTED) {
				return {...status, label: 'status.CANCELLED', value: StatusEnum.CANCELLED};
			}
			return status;
		})
		.sort((a, b) => a.label.localeCompare(b.label));

	filtersForm = new FormGroup({
		status: new FormControl<string[]>([]),
		searchStore: new FormControl(null),
		startDate: new FormControl(null),
		endDate: new FormControl(null),
	});


	tableConfiguration: ITableConfiguration;
	dataSource: MatTableDataSource<CashbackShopTableResponse> =
		new MatTableDataSource<CashbackShopTableResponse>();
	page$: BehaviorSubject<number> = new BehaviorSubject<number>(1);
	currentPage: number = 1;
	sort: string = 'date desc';
	pageSize: number = AppConstants.PAGE_SIZE;
	totalSize: number = 0;
	protected readonly StatusEnum = StatusEnum;

	constructor(
		private readonly translocoService: TranslocoService,
		private readonly toasterService: ToasterService,
		private readonly destroyRef: DestroyRef,
        private readonly shopService: ShopService,
        private _queryService: UrlQueryService
	) {}

	ngOnInit(): void {
		if (this._queryService.getParam('status')) {
			this.setFilters(this._queryService.getParam('status'));
			this.fetchData(this.getBodyParams());
		} else {
			this.fetchData(this.getBodyParams());
		}
		this.setupTable();

		this.startDateControl.valueChanges.subscribe(() => {
			this.fetchData(this.getBodyParams());
		});

		this.endDateControl.valueChanges.subscribe(() => {
			this.fetchData(this.getBodyParams());
		});
	}

	fetchData(filtersParams?: Partial<ManageCashbackSearchParams>): void {
		this.shopService
			.getCashback(this.currentPage - 1, this.pageSize, this.sort, filtersParams)
			.pipe(
				tap(
					(cashbacks: PaginationInterface<CashbackShopTableResponse>) =>
						(this.totalSize = cashbacks.totalRows),
				),
				map((cashbacks: PaginationInterface<CashbackShopTableResponse>) => cashbacks.data),
				takeUntilDestroyed(this.destroyRef),
			)
			.subscribe((cashbacks: CashbackShopTableResponse[]) => {
				this.dataSource.data = cashbacks.map((c) => ({
					...c,
					reward: null,
				}));
			});
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
					id: 'id',
					name: this.translocoService.translate('cashback.table.id'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'order_name',
					name: 'Order #',
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
					id: 'amountSource',
					name: this.translocoService.translate('cashback.table.order-value'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'status',
					name: this.translocoService.translate('cashback.table.status'),
					hasSort: true,
					hasTooltip: false,
				},
			],
		};
	}

	setFilters(filter: string, fromWarning: boolean = false): void {
		let filters: string[] = !fromWarning
			? Array.from(new Set([...this.filtersForm.controls.status.value, filter]))
			: [filter];

		this.filtersForm.controls.status.patchValue(filters);
		this.fetchData(this.getBodyParams());
	}

	removeFilter(filter: string): void {
		if (filter === 'CONFIRMED') {
			filter = 'VALIDATED';
		}
		const filters: string[] = this.filtersForm.controls.status.value.filter(
			(value: string) => value !== filter,
		);
		this.filtersForm.controls.status.patchValue(filters);
		this.fetchData(this.getBodyParams());
	}

	getBodyParams(): Partial<ManageCashbackSearchParams> {
		return {
			storeUuid: this.shopService.$currentShop().storeUuid,
			state: this.filtersForm.value.status.join(',') ?? null,
			startDate: this.startDateControl.value ?? null,
			endDate: this.endDateControl.value ?? null,
		};
	}

	totalPages(): number {
		return Math.ceil(this.totalSize / this.pageSize);
	}

	informCopy() {
		this.toasterService.showToast('info', 'Copied to clipboard');
	}

	onSortAndPageChange(event: {sort: string; page: PageEvent}) {
		if (event.sort) {
			const field = CashbackSortingFields[event.sort.split(',')[0]];
			const direction = event.sort.split(',')[1];
			this.sort = `${field} ${direction}`;
		}
		this.currentPage = event.page.pageIndex;
		this.page$.next(event.page.pageIndex);
		this.fetchData(this.getBodyParams());
	}

	resetFormControl(formControlName: string): void {
		this.filtersForm.get(formControlName).reset();
	}

	get searchStoreForm(): FormControl {
		return this.filtersForm.get('searchStore') as FormControl;
	}

	get startDateControl(): FormControl {
		return this.filtersForm.get('startDate') as FormControl;
	}

	get endDateControl(): FormControl {
		return this.filtersForm.get('endDate') as FormControl;
	}

	protected readonly environment = environment;
}
