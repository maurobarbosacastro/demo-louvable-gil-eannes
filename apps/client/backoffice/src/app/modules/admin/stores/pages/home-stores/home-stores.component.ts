import {Component, computed, inject, Signal, signal, WritableSignal} from '@angular/core';
import {UrlQueryService} from '@app/shared/services/url-query.service';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {StoresComponent} from '@app/modules/admin/stores/pages/list-stores/stores.component';
import {NgClass} from '@angular/common';
import {ApprovalStoresComponent} from '@app/modules/admin/stores/pages/approval-stores/approval-stores.component';

enum Tabs {
	LIST_STORES = 'list',
	APPROVALS = 'approvals'
}

@Component({
	selector: 'tagpeak-home-stores',
	templateUrl: 'home-stores.component.html',
	standalone: true,
	imports: [TranslocoPipe, StoresComponent, NgClass, ApprovalStoresComponent],
})
export class HomeStoresComponent {
	private translocoService: TranslocoService = inject(TranslocoService);
	private queryService: UrlQueryService = inject(UrlQueryService);

	Tabs = Tabs;
	//Tab names
	tabs: Tabs[] = [Tabs.LIST_STORES, Tabs.APPROVALS];

	$selectedTab: WritableSignal<Tabs> = signal(this.tabs[0]);

	ngOnInit() {
		if (this.queryService.getParam('tab')) {
			this.$selectedTab.set(this.queryService.getParam('tab') as Tabs);
		} else {
			this.queryService.addParam('tab', this.tabs[0]);
		}
	}
	getTabName(tab: Tabs): string {
		return this.translocoService.translate(`admin-stores.tabs.${tab}`);
	}

	selectTab(tab: Tabs) {
		this.$selectedTab.update(() => tab);
		this.queryService.addParam('tab', tab);
	}
}
