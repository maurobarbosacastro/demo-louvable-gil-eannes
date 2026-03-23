import {Component, computed, inject, signal, Signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {TableComponent} from '@app/shared/components/table/table.component';
import {TranslocoPipe} from '@ngneat/transloco';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {toSignal} from '@angular/core/rxjs-interop';
import {
	BehaviorSubject,
	combineLatest,
	debounceTime,
	distinctUntilChanged,
	map,
	startWith,
	switchMap,
	tap,
} from 'rxjs';
import {StoresService} from '@app/modules/admin/stores/services/stores.service';
import {AppConstants} from '@app/app.constants';
import {PageEvent} from '@angular/material/paginator';
import {StoreVisitsService} from '@app/modules/client/store-visits/services/store-visits.service';
import {
	StoreVisitsInterface,
	StoreVisitsSortingFields,
} from '@app/modules/client/store-visits/models/store-visits.interface';
import {StoreInterface} from '@app/modules/client/store-visits/models/store.interface';
import {AdminStoreInterface} from '@app/modules/admin/stores/interfaces/store.interface';
import {MatDatepicker, MatDatepickerInput} from '@angular/material/datepicker';
import {addDays} from 'date-fns';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-store-visits',
	standalone: true,
	imports: [
		CommonModule,
		MatIcon,
		ReactiveFormsModule,
		TableComponent,
		TranslocoPipe,
		MatDatepicker,
		MatDatepickerInput,
	],
	templateUrl: './store-visits.component.html',
	styleUrls: [
		'../../../../../styles/style-table.scss',
		'../../../../../styles/style-date-picker.scss',
	],
})
export class StoreVisitsComponent {
	private _storeService = inject(StoresService);
	private _storeVisitsService = inject(StoreVisitsService);
	private _screenService: ScreenService = inject(ScreenService);

	tableConfiguration: ITableConfiguration;
	page$: BehaviorSubject<number> = new BehaviorSubject<number>(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject<string>('created_at desc'); // Default sorting
	totalSize: number = 1;
	pageSize: number = AppConstants.PAGE_SIZE;
	page: number;
	total: number;

	searchForm = new FormGroup({
		name: new FormControl<string>(null),
		store: new FormControl<StoreInterface>(null),
		reference: new FormControl<string>(null),
		encodedReference: new FormControl<string>(null),
		startDate: new FormControl(null),
		endDate: new FormControl(null),
	});

	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	$stores: Signal<AdminStoreInterface[]> = toSignal(
		this._storeService.getAllStores().pipe(
			map((res) => res.data ?? []),
			tap((res) => {
				this.$allStores.set(res);
			}),
		),
	);
	$allStores: WritableSignal<AdminStoreInterface[]> = signal([]);

	$dataNew: Signal<StoreVisitsInterface[]> = toSignal(
		combineLatest([
			this.sort$,
			this.page$,
			this.nameControl.valueChanges.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				map((name) => name),
			),
			this.referenceControl.valueChanges.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				map((reference) => reference),
			),
			this.encodedReferenceControl.valueChanges.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				map((encodedReference) => encodedReference),
			),
			this.storeControl.valueChanges.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				map((store) => store),
			),
			this.startDateControl.valueChanges.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				map((startDate) => startDate),
			),
			this.endDateControl.valueChanges.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				map((endDate) => endDate),
			),
		]).pipe(
			switchMap(([sort, page, name, reference, encodedReference, store, startDate, endDate]) => {
				return this._storeVisitsService.getAllStoreVisits(page, this.pageSize, sort, {
					name: name,
					store: store ? store.uuid : '',
					startDate: startDate,
					endDate: endDate ? new Date(addDays(new Date(endDate), 1)).toISOString() : '',
					reference: encodedReference ? this.decodeStoreVisitRef(encodedReference) : reference,
				});
			}),
			tap((res) => {
				this.page = res.page;
				this.total = res.totalPages;
				this.totalSize = res.totalRows;
			}),
			map((res) => {
				res.data = res?.data?.map((sv) => {
					sv.encodedReference = this.encodeStoreVisitRef(sv?.reference, sv?.user.uuid);
					return sv;
				});
				return res.data ?? [];
			}),
		),
	);

	$tableConfig: Signal<ITableConfiguration> = computed(() => {
		return {
			dataSource: new MatTableDataSource(this.$dataNew()),
			css: 'rounded-t-[0.625rem] overflow-hidden shadow-table',
			pageSize: this.pageSize,
			styles: {
				header: 'font-plex text-project-waterloo text-md font-normal',
				content: 'font-plex text-project-licorice-900 text-md',
				paginator: 'font-plex text-black',
			},
			columns: [
				{
					id: 'user',
					name: 'Name',
					hasSort: false,
					hasTooltip: false,
					colWidth: '20%',
				},
				{
					id: 'reference',
					name: 'Reference',
					hasSort: true,
					hasTooltip: false,
					colWidth: '10%',
				},
				{
					id: 'encodedReference',
					name: 'Encoded Reference',
					hasSort: false,
					hasTooltip: false,
					colWidth: '15%',
				},
				{
					id: 'purchased',
					name: 'Purchase',
					hasSort: true,
					hasTooltip: false,
					colWidth: '10%',
				},
				{
					id: 'store',
					name: 'Store',
					hasSort: false,
					hasTooltip: false,
					colWidth: '20%',
				},
				{
					id: 'dateTime',
					name: 'Date & Time',
					hasSort: true,
					hasTooltip: false,
					colWidth: '20%',
				},
			],
		};
	});

	onSortAndPageChange(evt: {sort: string; page: PageEvent}): void {
		if (evt.sort) {
			const field = StoreVisitsSortingFields[evt.sort.split(',')[0]];
			const direction = evt.sort.split(',')[1];
			this.sort$.next(`${field} ${direction}`);
			return;
		}
		this.page$.next(evt.page.pageIndex);
	}

	searchByStoreName(storeName: string) {
		this.$allStores.set(this.$stores());
		if (storeName) {
			this.$allStores.set(
				this.$stores().filter((store: AdminStoreInterface) =>
					store.name.toLowerCase().includes(storeName.toLowerCase()),
				),
			);
		} else if (storeName === '') {
			this.resetFormControl('store');
		}
	}

	get nameControl(): FormControl {
		return this.searchForm.get('name') as FormControl;
	}

	get storeControl(): FormControl {
		return this.searchForm.get('store') as FormControl;
	}

	get startDateControl(): FormControl {
		return this.searchForm.get('startDate') as FormControl;
	}

	get endDateControl(): FormControl {
		return this.searchForm.get('endDate') as FormControl;
	}

	get referenceControl(): FormControl {
		return this.searchForm.get('reference') as FormControl;
	}

	get encodedReferenceControl(): FormControl {
		return this.searchForm.get('encodedReference') as FormControl;
	}

	setFormControlName(formControlName: string, value: StoreInterface | string): void {
		this.searchForm.get(formControlName).setValue(value);
	}

	resetFormControl(formControlName: string): void {
		this.searchForm.get(formControlName).reset();
		this.$allStores.set(this.$stores());
	}

	encodeStoreVisitRef(ref: number, userUuid: string): string {
		return btoa(ref + '_' + userUuid.split('-')[0]);
	}

	decodeStoreVisitRef(ref: string): string {
		try {
			const decodedRef = atob(ref);
			return decodedRef.split('_')[0] ?? '';
		} catch (error) {
			console.error('Invalid Base64 string:', ref);
			return '';
		}
	}
}
