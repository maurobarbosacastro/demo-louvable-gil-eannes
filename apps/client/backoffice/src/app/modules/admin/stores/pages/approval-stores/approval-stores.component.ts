import {Component, computed, inject, Signal} from '@angular/core';
import {StoresService} from '@app/modules/admin/stores/services/stores.service';
import {toSignal} from '@angular/core/rxjs-interop';
import {AsyncPipe, DatePipe, NgIf, TitleCasePipe} from '@angular/common';
import {TableComponent} from '@app/shared/components/table/table.component';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {BehaviorSubject, combineLatest, map, switchMap} from 'rxjs';
import {AppConstants} from '@app/app.constants';
import {
	StoreApprovalRequest,
	StoresSortingFields,
} from '@app/modules/admin/stores/interfaces/store.interface';
import {PageEvent} from '@angular/material/paginator';
import {MatTooltip} from '@angular/material/tooltip';
import {EditStoreComponent} from '@app/modules/admin/stores/components/edit-store/edit-store.component';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-approval-stores',
	templateUrl: 'approval-stores-component.component.html',
	standalone: true,
	imports: [AsyncPipe, NgIf, TableComponent, DatePipe, MatTooltip, TitleCasePipe],
})
export class ApprovalStoresComponent {
	private storeService = inject(StoresService);
	private _dialog = inject(MatDialog);
	private _screenService = inject(ScreenService);

	tableConfiguration: ITableConfiguration;
	page$: BehaviorSubject<number> = new BehaviorSubject<number>(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject<string>('created_at asc'); // Default sorting
	totalSize: number = 1;
	pageSize: number = AppConstants.PAGE_SIZE;
	page: number;
	total: number;
	searchText: string = '';

	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);
	$stores: Signal<StoreApprovalRequest[]> = toSignal(
		combineLatest([this.sort$, this.page$]).pipe(
			switchMap(([sort, page]) => {
				return this.storeService.getStoresForApprovals(this.pageSize, page, sort);
			}),
			map((res) => {
				this.page = res.page;
				this.total = res.totalPages;
				this.totalSize = res.totalRows;
				return res.data ?? [];
			}),
		),
	);

	$tableConfig: Signal<ITableConfiguration> = computed(() => {
		return {
			dataSource: new MatTableDataSource(this.$stores()),
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
					colWidth: '25%',
				},
				{
					id: 'status',
					name: 'Status',
					hasSort: true,
					hasTooltip: false,
					colWidth: '10%',
				},
				{
					id: 'createdAt',
					name: 'Date',
					hasSort: true,
					hasTooltip: false,
					colWidth: '10%',
				},
				{
					id: 'user',
					name: 'User',
					hasSort: false,
					hasTooltip: false,
					colWidth: '35%',
				},
				{
					id: 'actions',
					name: 'Actions',
					hasSort: false,
					hasTooltip: false,
					colWidth: '15%',
					actions: [
						{
							type: 'view',
							iconUrl: 'eye',
							css: 'scale-75',
						},
					],
				},
			],
		};
	});

	clickTableAction(event: any): void {
		switch (event.type) {
			case 'view': {
				this.openEditStoreDialog(event.element.uuid);
				break;
			}
		}
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

	openEditStoreDialog(uuid?: string): void {
		const modalConfig: MatDialogConfig = {
			...AppConstants.MODAL_CONFIG,
			data: {
				uuid: uuid,
				approval: true,
			},
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
}
