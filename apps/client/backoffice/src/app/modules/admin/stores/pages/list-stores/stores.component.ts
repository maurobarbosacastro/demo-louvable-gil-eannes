import {Component, computed, DestroyRef, inject, OnInit, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TableComponent} from '@app/shared/components/table/table.component';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {MatTableDataSource} from '@angular/material/table';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {PageEvent} from '@angular/material/paginator';
import {StoresService} from '@app/modules/admin/stores/services/stores.service';
import {
	AdminStoreInterface,
	StoresSortingFields,
} from '@app/modules/admin/stores/interfaces/store.interface';
import {EditStoreComponent} from '@app/modules/admin/stores/components/edit-store/edit-store.component';
import {UploadCsvComponent} from '@app/modules/admin/stores/components/upload-csv/upload-csv.component';
import {AppConstants} from '@app/app.constants';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
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
import {MatTooltip} from '@angular/material/tooltip';
import {DeleteModalConfirmationComponent} from '@app/shared/components/delete-modal-confirmation/delete-modal-confirmation.component';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule} from '@angular/forms';
import {MatIcon} from '@angular/material/icon';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-stores',
	standalone: true,
	imports: [
		CommonModule,
		TableComponent,
		TranslocoPipe,
		MatTooltip,
		FormsModule,
		MatIcon,
		ReactiveFormsModule,
	],
	templateUrl: './stores.component.html',
	styleUrl: '../../../../../../styles/style-table.scss',
})
export class StoresComponent implements OnInit {
	private _storeService = inject(StoresService);
	private _dialog = inject(MatDialog);
	private _toaster = inject(ToasterService);
	private transloco = inject(TranslocoService);
	private destroyRef = inject(DestroyRef);
	private _screenService = inject(ScreenService);

	tableConfiguration: ITableConfiguration;
	page$: BehaviorSubject<number> = new BehaviorSubject<number>(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject<string>('name asc'); // Default sorting
	totalSize: number = 1;
	pageSize: number = AppConstants.PAGE_SIZE;
	page: number;
	total: number;
	searchText: string = '';

	searchForm = new FormGroup({
		search: new FormControl<string>(null),
	});

	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	$dataNew: Signal<AdminStoreInterface[]> = toSignal(
		combineLatest([this.sort$, this.page$]).pipe(
			switchMap(([sort, page]) => {
				return this._storeService.getAllStores(this.pageSize, page, sort, this.searchText);
			}),
			tap((res) => {
				this.page = res.page;
				this.total = res.totalPages;
				this.totalSize = res.totalRows;
			}),
			map((res) => res.data ?? []),
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
					id: 'name',
					name: 'Name',
					hasSort: true,
					hasTooltip: false,
					colWidth: '20%',
				},
				{
					id: 'partner',
					name: 'Source',
					hasSort: true,
					hasTooltip: false,
					colWidth: '20%',
				},
				{
					id: 'country',
					name: 'Country',
					hasSort: false,
					hasTooltip: false,
					colWidth: '20%',
				},
				{
					id: 'description',
					name: 'Description',
					hasSort: false,
					hasTooltip: false,
					colWidth: '20%',
				},
				{
					id: 'actions',
					name: 'Actions',
					hasSort: false,
					hasTooltip: false,
					colWidth: '15%',
					actions: [
						{
							type: 'edit',
							iconUrl: 'pen',
							css: 'scale-75',
						},
						{
							type: 'delete',
							iconUrl: 'tag-delete',
						},
					],
				},
			],
		};
	});

	ngOnInit(): void {
		this.searchControl.valueChanges
			.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				takeUntilDestroyed(this.destroyRef),
			)
			.subscribe((searchText) => {
				this.searchText = searchText;
				this.page$.next(1);
			});
	}

	openEditStoreDialog(uuid?: string): void {
		const modalConfig: MatDialogConfig = {
			...AppConstants.MODAL_CONFIG,
			data: {uuid: uuid},
		};

		if (this.$isMobile()) {
			modalConfig.width = '100%';
			modalConfig.height = '95%';
			modalConfig.panelClass = 'dialog-panel-details-mobile';
		}

		const dialogRef = this._dialog.open(EditStoreComponent, modalConfig);

		dialogRef.afterClosed().subscribe((result) => {
			if (!result.skipFetch) {
				this.page$.next(1);
			}
		});
	}

	openUploadExcelDialog() {
		const dialogRef = this._dialog.open(UploadCsvComponent, {
			panelClass: 'upload-csv-dialog-panel',
		});

		dialogRef.afterClosed().subscribe((result) => {
			if (result) {
				this.page$.next(1);
			}
		});
	}

	exportExcel(): void {
		const title = `stores_${new Date().getTime()}.csv`;
		this._storeService.exportStoresCSV(title, this.sort$.value).subscribe({
			next: (file: File) => {
				const url: string = window.URL.createObjectURL(file);
				const link: HTMLAnchorElement = document.createElement('a');
				link.href = url;
				link.setAttribute('download', file.name);
				document.body.appendChild(link);
				link.click();
				document.body.removeChild(link);
			},
			error: () => {
				this._toaster.showToast('error', this.transloco.translate('store.error-export'));
			},
		});
	}

	clickTableAction(event: any): void {
		switch (event.type) {
			case 'edit': {
				this.openEditStoreDialog(event.element.uuid);
				break;
			}
			case 'delete': {
				this.deleteStore(event.element.uuid);
				break;
			}
		}
	}

	deleteStore(id: string): void {
		this._dialog
			.open(DeleteModalConfirmationComponent, {
				panelClass: 'dialog-panel',
				height: 'fit-content',
				width: '26rem',
			})
			.afterClosed()
			.pipe(
				filter((result: {skipFetch: boolean}) => result && !result.skipFetch),
				switchMap(() => this._storeService.deleteStore(id)),
			)
			.subscribe({
				next: () => {
					this.page$.next(1);
					this._toaster.showToast(
						'success',
						this.transloco.translate('store.success-delete'),
						'bottom',
					);
				},
				error: (err) => {
					if (err) {
						this._toaster.showToast(
							'error',
							this.transloco.translate('store.error-delete'),
							'bottom',
						);
					}
				},
			});
	}

	onSortAndPageChange(evt: {sort: string; page: PageEvent}): void {
		if (evt.sort) {
			const field = StoresSortingFields[evt.sort.split(',')[0]];
			const direction = evt.sort.split(',')[1];
			this.sort$.next(`${field} ${direction}`);
			return;
		}
		this.page$.next(evt.page.pageIndex);
	}

	get searchControl(): FormControl {
		return this.searchForm.get('search') as FormControl;
	}
}
