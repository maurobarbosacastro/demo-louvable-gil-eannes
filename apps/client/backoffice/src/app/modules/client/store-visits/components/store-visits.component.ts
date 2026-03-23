import {Component, DestroyRef, inject, OnInit, Signal, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {StoreVisitsService} from '@app/modules/client/store-visits/services/store-visits.service';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {StoreInterface} from '@app/modules/client/store-visits/models/store.interface';
import {BehaviorSubject, combineLatest, map, switchMap, tap} from 'rxjs';
import {
	StoreVisitsInterface,
	StoreVisitsSearchParams,
	StoreVisitsSortingFields,
} from '@app/modules/client/store-visits/models/store-visits.interface';
import {MatDatepicker, MatDatepickerInput} from '@angular/material/datepicker';
import {TableComponent} from '@app/shared/components/table/table.component';
import {UserService} from '@app/core/user/user.service';
import {TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {addDays} from 'date-fns';
import {PageEvent} from '@angular/material/paginator';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {AppConstants} from '@app/app.constants';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-store-visits',
	standalone: true,
	imports: [
		CommonModule,
		ReactiveFormsModule,
		MatIcon,
		TranslocoPipe,
		MatDatepickerInput,
		MatDatepicker,
		TableComponent,
	],
	templateUrl: './store-visits.component.html',
	styleUrls: [
		'./store-visits.component.scss',
		'../../../../../styles/style-table.scss',
		'../../../../../styles/style-date-picker.scss',
	],
})
export class StoreVisitsComponent implements OnInit {
	private storeVisitsService: StoreVisitsService = inject(StoreVisitsService);
	private translocoService: TranslocoService = inject(TranslocoService);
	private userService = inject(UserService);
	private destroyRef = inject(DestroyRef);
	private _screenService: ScreenService = inject(ScreenService);

	currentPage: number = 1;
	pageSize: number = AppConstants.PAGE_SIZE;
	totalSize: number;
	totalPages: number;
	page: number;

	$stores: WritableSignal<StoreInterface[]> = signal([]);
	$user: WritableSignal<TagpeakUser> = signal(null);
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	filtersForm = new FormGroup({
		store: new FormControl<StoreInterface>(null),
		startDate: new FormControl(null),
		endDate: new FormControl(null),
	});

	tableConfiguration: ITableConfiguration;
	dataSource: MatTableDataSource<StoreVisitsInterface> =
		new MatTableDataSource<StoreVisitsInterface>();
	isSearched: boolean;

	page$: BehaviorSubject<number> = new BehaviorSubject(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject('created_at desc');
	filters$: BehaviorSubject<any> = new BehaviorSubject(null);

	ngOnInit() {
		this.userService.user$.subscribe((user: TagpeakUser) => this.$user.set(user));
		this.getAllStores();
		this.fetchData();
		this.setupTable();
	}

	getAllStores(): void {
		this.storeVisitsService
			.getStoresVisitedByUser(this.currentPage, this.pageSize, 'name asc', {})
			.subscribe((stores) => {
				this.$stores.set(stores.data);
				this.page$.next(1);
			});
	}

	fetchData(): void {
		combineLatest([this.sort$, this.page$, this.filters$])
			.pipe(
				takeUntilDestroyed(this.destroyRef),
				switchMap(([sort, page, filters]) =>
					this.storeVisitsService.getStoreVisitsByUser(page, this.pageSize, sort, filters),
				),
				tap((storeVisits) => {
					this.totalSize = storeVisits.totalRows;
					this.totalPages = storeVisits.totalPages;
					this.page = storeVisits.page;
				}),
				map((storeVisits) => storeVisits.data),
			)
			.subscribe((storeVisits) => {
				this.dataSource.data = storeVisits;
			});
	}

	setupTable(): void {
		this.tableConfiguration = {
			dataSource: this.dataSource,
			pageSize: this.pageSize,
			styles: {
				header: 'font-nunitoBlack text-project-waterloo',
				content: 'font-nunitoRegular text-project-licorice-900',
				paginator: 'font-nunitoRegular text-black',
			},
			columns: [
				{
					id: 'refId',
					name: this.translocoService.translate('store-visits.table.ref'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'storeName',
					name: this.translocoService.translate('store-visits.table.store-name'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'purchased',
					name: this.translocoService.translate('store-visits.table.purchased'),
					hasSort: true,
					hasTooltip: false,
				},
				{
					id: 'dateTime',
					name: this.translocoService.translate('store-visits.table.date-time'),
					hasSort: true,
					hasTooltip: false,
				},
			],
		};
	}

	setFormControlName(formControlName: string, value: StoreInterface | string): void {
		this.filtersForm.get(formControlName).setValue(value);
	}

	resetFormControl(formControlName: string): void {
		this.filtersForm.get(formControlName).reset();
		this.searchByFilter();
	}

	searchByFilter(): void {
		this.isSearched = true;
		const searchBody: Partial<StoreVisitsSearchParams> = {
			store: this.filtersForm.value.store?.uuid ?? null,
			startDate: this.filtersForm.value.startDate ?? null,
			endDate: this.filtersForm.value.endDate ?? null,
		};

		if (!searchBody.store && !searchBody.startDate && !searchBody.endDate) {
			this.isSearched = false;
		}

		if (searchBody?.endDate) {
			searchBody.endDate = new Date(addDays(new Date(searchBody.endDate), 1)).toISOString();
		}

		this.page$.next(1);
		this.filters$.next(searchBody);
	}

	onSortAndPageChange(evt: {sort: string; page: PageEvent}) {
		if (evt.sort) {
			const field = StoreVisitsSortingFields[evt.sort.split(',')[0]];
			const direction = evt.sort.split(',')[1];
			this.page$.next(1);
			this.sort$.next(`${field} ${direction}`);
			return;
		}
		this.page$.next(evt.page.pageIndex);
	}
}
