import {Component, computed, OnInit, Signal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TableComponent} from '@app/shared/components/table/table.component';
import {ITableConfiguration} from '@app/shared/components/table/table.interface';
import {PageEvent} from '@angular/material/paginator';
import {Router} from '@angular/router';
import {BehaviorSubject, combineLatest, map, switchMap, tap} from 'rxjs';
import {MatTableDataSource} from '@angular/material/table';
import {TemplateInterface} from '@app/shared/interfaces/template.interface';
import {AppConstants} from '@app/app.constants';
import {EmailService} from '@app/modules/admin/email/email.service';
import {TranslocoPipe} from '@ngneat/transloco';
import {toSignal} from '@angular/core/rxjs-interop';

@Component({
	selector: 'tagpeak-ms-listing',
	standalone: true,
	imports: [CommonModule, TableComponent, TranslocoPipe],
	templateUrl: './listing.component.html',
	styleUrls: ['./listing.component.scss', '../../../../../styles/style-table.scss'],
})
export class ListingComponent implements OnInit {
	page$: BehaviorSubject<number> = new BehaviorSubject(1);
	sort$: BehaviorSubject<string> = new BehaviorSubject<string>('name');
	tableConfiguration: ITableConfiguration;
	page: number;
	pageSize = AppConstants.PAGE_SIZE;
	sort = AppConstants.SORT_BY_UPDATE_AT;
	total: number;
	totalSize: number = 1;

	$dataNew: Signal<TemplateInterface[]> = toSignal(
		combineLatest([this.sort$, this.page$]).pipe(
			switchMap(([sort, page]) => {
				return this.emailService.getTemplates(this.pageSize, page - 1, sort);
			}),
			tap((res) => {
				this.page = res['pageNumber'];
				this.total = res['totalPages'];
				this.totalSize = res['totalSize'];
			}),
			map((res) => res['content']),
		),
	);

	$tableConfig: Signal<ITableConfiguration> = computed(() => {
		return {
			dataSource: new MatTableDataSource(this.$dataNew()),
			css: 'rounded-t-[0.625rem] overflow-hidden shadow-table',
			pageSize: this.pageSize,
			styles: {
				header: 'font-nunitoBlack text-black bg-project-tableGray',
				content: 'text-black',
				paginator: 'text-black',
			},
			columns: [
				{
					id: 'code',
					name: 'Code',
					hasSort: false,
					hasTooltip: false,
					sortActionDescription: 'id',
					colWidth: '20%',
				},
				{
					id: 'name',
					name: 'Template Name',
					hasSort: false,
					hasTooltip: false,
					sortActionDescription: 'Nome do template',
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
							type: 'edit',
							iconUrl: 'pen',
							css: 'scale-75',
						},
						{
							type: 'duplicate',
							iconUrl: 'duplicate',
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

	constructor(
		private readonly router: Router,
		private readonly emailService: EmailService,
	) {}

	ngOnInit(): void {}

	clickTableAction(event: any): void {
		switch (event.type) {
			case 'edit': {
				void this.router.navigate([
					AppConstants.ROUTES.admin + AppConstants.ROUTES.email,
					event.element.id,
				]);
				break;
			}
			case 'duplicate': {
				this.duplicateTemplate(event);
				break;
			}
			case 'delete': {
				void this.emailService.deleteTemplate(event.element.id).subscribe((_) => {});
				break;
			}
		}
	}

	duplicateTemplate(event: any): void {
		this.emailService.templateInfoDuplicate(event);
		this.redirectToNewTemplate();
	}

	onSortAndPageChange(evt: {sort: string; page: PageEvent}): void {
		if (evt.sort) {
			this.sort$.next(evt.sort.split(',').join(' '));
		}
		this.page$.next(evt.page.pageIndex);
	}

	redirectToNewTemplate(): void {
		this.router.navigate([
			AppConstants.ROUTES.admin + AppConstants.ROUTES.email + AppConstants.ROUTES.new,
		]);
	}
}
