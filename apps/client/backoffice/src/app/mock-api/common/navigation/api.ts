import {Injectable} from '@angular/core';
import {FuseNavigationItem} from '@fuse/components/navigation';
import {FuseMockApiService} from '@fuse/lib/mock-api';
import {cloneDeep} from 'lodash-es';
import {
    compactNavigation, compactNavigationShop,
    defaultNavigation,
    futuristicNavigation,
    horizontalNavigation,
} from '@app/mock-api/common/navigation/data';
import {BugReportComponent} from '@app/modules/client/bug-report/components/bug-report.component';
import {MatDialog} from '@angular/material/dialog';

@Injectable({providedIn: 'root'})
export class NavigationMockApi {
	private readonly _compactNavigation: FuseNavigationItem[] = compactNavigation;
	private readonly _compactNavigationShop: FuseNavigationItem[] = compactNavigationShop;
	private readonly _defaultNavigation: FuseNavigationItem[] = defaultNavigation;
	private readonly _futuristicNavigation: FuseNavigationItem[] = futuristicNavigation;
	private readonly _horizontalNavigation: FuseNavigationItem[] = horizontalNavigation;

	/**
	 * Constructor
	 */
	constructor(
		private _fuseMockApiService: FuseMockApiService,
		private dialog: MatDialog,
	) {
		// Register Mock API handlers
		this.registerHandlers();
	}

	// -----------------------------------------------------------------------------------------------------
	// @ Public methods
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Register Mock API handlers
	 */
	registerHandlers(): void {
		// -----------------------------------------------------------------------------------------------------
		// @ Navigation - GET
		// -----------------------------------------------------------------------------------------------------
		this._fuseMockApiService.onGet('api/common/navigation').reply(() => {
			// Fill compact navigation children using the default navigation
			this._compactNavigation.forEach((compactNavItem) => {
				this._defaultNavigation.forEach((defaultNavItem) => {
					if (defaultNavItem.id === compactNavItem.id) {
						compactNavItem.children = cloneDeep(defaultNavItem.children);
					}
				});
			});
			// Fill compact navigation children using the default navigation
			this._compactNavigationShop.forEach((compactNavItem) => {
				this._defaultNavigation.forEach((defaultNavItem) => {
					if (defaultNavItem.id === compactNavItem.id) {
						compactNavItem.children = cloneDeep(defaultNavItem.children);
					}
				});
			});

			// Fill futuristic navigation children using the default navigation
			this._futuristicNavigation.forEach((futuristicNavItem) => {
				this._defaultNavigation.forEach((defaultNavItem) => {
					if (defaultNavItem.id === futuristicNavItem.id) {
						futuristicNavItem.children = cloneDeep(defaultNavItem.children);
					}
				});
			});

			// Fill horizontal navigation children using the default navigation
			this._horizontalNavigation.forEach((horizontalNavItem) => {
				this._defaultNavigation.forEach((defaultNavItem) => {
					if (defaultNavItem.id === horizontalNavItem.id) {
						horizontalNavItem.children = cloneDeep(defaultNavItem.children);
					}
				});
				if (horizontalNavItem.id === 'help') {
					horizontalNavItem.function = () => {
						this.dialog.open(BugReportComponent, {
							panelClass: 'onboarding-dialog-panel.smaller',
							position: {right: '2rem', bottom: '4rem'},
						});
					};
				}
			});

			// Return the response
			return [
				200,
				{
					compact: cloneDeep(this._compactNavigation),
					compactShop: cloneDeep(this._compactNavigationShop),
					default: cloneDeep(this._defaultNavigation),
					futuristic: cloneDeep(this._futuristicNavigation),
					horizontal: cloneDeep(this._horizontalNavigation),
				},
			];
		});
	}
}
