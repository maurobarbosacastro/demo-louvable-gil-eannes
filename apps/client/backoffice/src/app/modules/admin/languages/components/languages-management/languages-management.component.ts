import {Component, computed, DestroyRef, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {BehaviorSubject, combineLatest, filter, map, Observable, switchMap, tap} from 'rxjs';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {LanguagesService} from '@app/modules/admin/languages/services/languages.service';
import {LanguagesInterface} from '@app/modules/admin/languages/models/languages.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {MatTableDataSource} from '@angular/material/table';
import {TableComponent} from '@app/shared/components/table/table.component';
import {MatDialog, MatDialogConfig} from '@angular/material/dialog';
import {LanguageDetailsComponent} from '@app/modules/admin/languages/components/language-details/language-details.component';
import {ModalActionResult} from '@app/utils/modal-action-result.enum';
import {PageEvent} from '@angular/material/paginator';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {HttpErrorResponse} from '@angular/common/http';
import {DeleteModalConfirmationComponent} from '@app/shared/components/delete-modal-confirmation/delete-modal-confirmation.component';
import {AppConstants} from '@app/app.constants';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-language-management',
	standalone: true,
	imports: [CommonModule, TranslocoPipe, TableComponent],
	templateUrl: './languages-management.component.html',
	styleUrl: '../../../../../../styles/style-table.scss',
})
export class LanguageManagementComponent {
	page$: BehaviorSubject<number> = new BehaviorSubject(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject('code');
	tableConfiguration: ITableConfiguration;
	totalSize: number = 1;
	pageSize: number = 10;
	page: number;
	total: number;

	$languages: Signal<LanguagesInterface[]> = toSignal(this.fetchData(this.pageSize));

	$tableConfiguration: Signal<ITableConfiguration> = computed(() => this.setupTableConfiguration());

	$isMobile: Signal<boolean> = toSignal(this.screenService.isMobile$);
	constructor(
		private readonly languagesService: LanguagesService,
		private readonly translocoService: TranslocoService,
		private readonly matDialog: MatDialog,
		private readonly toasterService: ToasterService,
		private readonly destroyRef: DestroyRef,
		private readonly screenService: ScreenService,
	) {}

	private fetchData(pageSize: number): Observable<LanguagesInterface[]> {
		return combineLatest([this.page$, this.sort$]).pipe(
			switchMap(([page, sort]: [number, string]) => {
				return this.languagesService.getAllLanguages(pageSize, page, sort);
			}),
			tap((res: PaginationInterface<LanguagesInterface>) => {
				this.page = res.page;
				this.total = res.totalPages;
				this.totalSize = res.totalRows;
			}),
			map((res: PaginationInterface<LanguagesInterface>) => res.data),
			takeUntilDestroyed(this.destroyRef),
		);
	}

	private setupTableConfiguration(): ITableConfiguration {
		return {
			dataSource: new MatTableDataSource(this.$languages()),
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
					name: this.translocoService.translate('languages.table.code'),
					hasSort: true,
					hasTooltip: false,
					sortActionDescription: 'code',
					colWidth: '30%',
				},
				{
					id: 'name',
					name: this.translocoService.translate('languages.table.name'),
					hasSort: true,
					hasTooltip: false,
					sortActionDescription: 'name',
					colWidth: '50%',
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
	}

	onSortAndPageChange(evt: {sort: string; page: PageEvent}): void {
		if (evt.sort) {
			this.sort$.next(evt.sort.split(',').join(' '));
		}
		this.page$.next(evt.page.pageIndex);
	}

	onActionClick(row: any): void {
		switch (row.type) {
			case 'edit': {
				this.openDetailsModal(row.element.uuid);
				break;
			}
			case 'delete': {
				this.deleteLanguage(row.element.uuid);
				break;
			}
		}
	}

	openDetailsModal(languageUuid?: string): void {
		const modalConfig: MatDialogConfig = {
			...AppConstants.MODAL_CONFIG,
			data: {id: languageUuid},
		};

		if (this.$isMobile()) {
			modalConfig.width = '100%';
			modalConfig.height = '95%';
			modalConfig.panelClass = 'dialog-panel-details-mobile';
		}

		this.matDialog
			.open(LanguageDetailsComponent, modalConfig)
			.afterClosed()
			.subscribe((result) => {
				if (!result.skipFetch) {
					this.page$.next(1);
				}
			});
	}

	deleteLanguage(languageUuid: string): void {
		this.matDialog
			.open(DeleteModalConfirmationComponent, {
				panelClass: 'dialog-panel',
				height: 'fit-content',
				width: '26rem',
			})
			.afterClosed()
			.pipe(
				filter((result: {skipFetch: boolean}) => result && !result.skipFetch),
				switchMap(() => this.languagesService.deleteLanguage(languageUuid)),
			)
			.subscribe({
				next: () => {
					this.page$.next(1);
					this.handleSuccessRequest('languages.success-delete');
				},
				error: (err) => {
					if (err) {
						this.handleErrorRequest('languages.error-delete');
					}
				},
			});
	}

	handleSuccessRequest(message: string): void {
		this.toasterService.showToast('success', this.translocoService.translate(message), 'bottom');
		this.fetchData(this.pageSize);
	}

	handleErrorRequest(message: string): void {
		this.toasterService.showToast('error', this.translocoService.translate(message), 'bottom');
	}
}
