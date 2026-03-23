import {Component, computed, inject, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {LineChartComponent} from '@app/shared/components/line-chart/line-chart.component';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {CashbackHistory} from '@app/modules/admin/cashback/models/manage-cashback.interface';
import {TableComponent} from '@app/shared/components/table/table.component';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {MatTableDataSource} from '@angular/material/table';
import {MatTooltip} from '@angular/material/tooltip';
import {BehaviorSubject, combineLatest, map, switchMap, tap} from 'rxjs';
import {PageEvent} from '@angular/material/paginator';
import {toSignal} from '@angular/core/rxjs-interop';
import {RewardService} from '@app/modules/admin/cashback/services/reward.service';

@Component({
	selector: 'tagpeak-chart-modal',
	standalone: true,
	imports: [CommonModule, MatIcon, LineChartComponent, TableComponent, MatTooltip],
	styleUrl: '../../../../../../styles/style-table.scss',
	templateUrl: './table-modal.component.html',
})
export class TableModalComponent {
	private dialogData = inject(MAT_DIALOG_DATA);
	private dialogRef = inject(MatDialogRef);
	private _rewardService: RewardService = inject(RewardService);

	tableConfiguration: ITableConfiguration;
	page$: BehaviorSubject<number> = new BehaviorSubject<number>(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject<string>('created_at'); // Default sorting
	totalSize: number = 1;
	pageSize: number = 5;
	page: number;
	total: number;

	$dataNew: Signal<CashbackHistory[]> = toSignal(
		combineLatest([this.sort$, this.page$]).pipe(
			switchMap(([sort, page]) => {
				return this._rewardService.getRewardHistory(
					this.pageSize,
					page,
					sort,
					this.dialogData.uuid,
				);
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
					id: 'rate',
					name: 'Rate',
					hasSort: false,
					hasTooltip: false,
					colWidth: '10%',
				},
				{
					id: 'units',
					name: 'Units',
					hasSort: false,
					hasTooltip: false,
					colWidth: '10%',
				},
				{
					id: 'cashReward',
					name: 'Cash Reward',
					hasSort: false,
					hasTooltip: false,
					colWidth: '20%',
				},
				{
					id: 'createdAt',
					name: 'Date',
					hasSort: true,
					hasTooltip: false,
					colWidth: '15%',
				},
			],
		};
	});

	onSortAndPageChange(evt: {sort: string; page: PageEvent}): void {
		if (evt.sort) {
			if (evt.sort.split(',')[0] == 'createdAt') {
				this.sort$.next('created_at' + ' ' + evt.sort.split(',')[1]);
			} else {
				this.sort$.next(evt.sort.split(',').join(' '));
			}
		}
		this.page$.next(evt.page.pageIndex);
	}

	closeModal(): void {
		this.dialogRef.close();
	}
}
