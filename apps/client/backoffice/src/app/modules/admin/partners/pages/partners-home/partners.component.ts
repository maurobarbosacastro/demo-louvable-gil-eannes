import {Component, computed, inject, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TableComponent} from '@app/shared/components/table/table.component';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {PartnersService} from '@app/modules/admin/partners/service/partners.service';
import {BehaviorSubject, combineLatest, filter, map, switchMap, tap} from 'rxjs';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {
	PartnersInterface,
	PartnerSortingFields,
} from '@app/modules/admin/partners/interfaces/partners.interface';
import {MatTableDataSource} from '@angular/material/table';
import {PageEvent} from '@angular/material/paginator';
import {toSignal} from '@angular/core/rxjs-interop';
import {EditPartnerComponent} from '@app/modules/admin/partners/components/edit-partner/edit-partner.component';
import {AppConstants} from '@app/app.constants';
import {DeleteModalConfirmationComponent} from '@app/shared/components/delete-modal-confirmation/delete-modal-confirmation.component';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-partners',
	standalone: true,
	imports: [CommonModule, TableComponent, TranslocoPipe],
	templateUrl: './partners.component.html',
	styleUrl: '../../../../../../styles/style-table.scss',
})
export class PartnersComponent {
	private partnersService = inject(PartnersService);
	private dialog = inject(MatDialog);
	private toaster = inject(ToasterService);
	private transloco = inject(TranslocoService);
	private screenService = inject(ScreenService);

	page$: BehaviorSubject<number> = new BehaviorSubject(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject('name');
	tableConfiguration: ITableConfiguration;
	totalSize: number = 1;
	pageSize = AppConstants.PAGE_SIZE;
	page: number;
	total: number;

	$isMobile: Signal<boolean> = toSignal(this.screenService.isMobile$);

	$dataNew: Signal<PartnersInterface[]> = toSignal(
		combineLatest([this.sort$, this.page$]).pipe(
			switchMap(([sort, page]) => {
				return this.partnersService.getPartners(this.pageSize, page, sort);
			}),
			tap((res) => {
				this.page = res.page;
				this.total = res.totalPages;
				this.totalSize = res.totalRows;
			}),
			map((res) => res.data),
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
					id: 'code',
					name: 'Code',
					hasSort: true,
					hasTooltip: false,
					sortActionDescription: 'Code',
					colWidth: '15%',
				},
				{
					id: 'name',
					name: 'Name',
					hasSort: true,
					hasTooltip: false,
					sortActionDescription: 'Country Name',
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

	openEditCountryDialog(partnerUuid?: string): void {
		const modalConfig: MatDialogConfig = {
			...AppConstants.MODAL_CONFIG,
			autoFocus: false,
			data: {partnerUuid: partnerUuid},
		};

		if (this.$isMobile()) {
			modalConfig.width = '100%';
			modalConfig.height = '95%';
			modalConfig.panelClass = 'dialog-panel-details-mobile';
		}

		const dialogRef = this.dialog.open(EditPartnerComponent, modalConfig);

		dialogRef.afterClosed().subscribe((result) => {
			if (!result.skipFetch) {
				this.page$.next(1);
			}
		});
	}

	clickTableAction(event: any): void {
		switch (event.type) {
			case 'edit': {
				this.openEditCountryDialog(event.element.uuid);
				break;
			}
			case 'delete': {
				this.deletePartner(event.element.uuid);
				break;
			}
		}
	}

	deletePartner(id: string): void {
		this.dialog
			.open(DeleteModalConfirmationComponent, {
				panelClass: 'dialog-panel',
				height: 'fit-content',
				width: '26rem',
			})
			.afterClosed()
			.pipe(
				filter((result: {skipFetch: boolean}) => result && !result.skipFetch),
				switchMap(() => this.partnersService.deletePartner(id)),
			)
			.subscribe({
				next: (_) => {
					this.page$.next(1);
					this.toaster.showToast('success', this.transloco.translate('partner.partner-deleted'));
				},
				error: (_) => {
					this.toaster.showToast('error', this.transloco.translate('partner.partner-delete-error'));
				},
			});
	}

	onSortAndPageChange(evt: {sort: string; page: PageEvent}): void {
		if (evt.sort) {
			const field = PartnerSortingFields[evt.sort.split(',')[0]];
			const direction = evt.sort.split(',')[1];
			this.sort$.next(`${field} ${direction}`);
			return;
		}
		this.page$.next(evt.page.pageIndex);
	}
}
