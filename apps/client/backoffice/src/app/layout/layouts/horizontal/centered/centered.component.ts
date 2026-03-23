import {
	Component,
	computed,
	OnDestroy,
	OnInit,
	signal,
	Signal,
	ViewEncapsulation,
} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatIconModule} from '@angular/material/icon';
import {RouterOutlet} from '@angular/router';
import {FuseLoadingBarComponent} from '@fuse/components/loading-bar';
import {
	FuseHorizontalNavigationComponent,
	FuseNavigationService,
	FuseVerticalNavigationComponent,
} from '@fuse/components/navigation';
import {FuseMediaWatcherService} from '@fuse/services/media-watcher';
import {NavigationService} from '@app/core/navigation/navigation.service';
import {Navigation} from '@app/core/navigation/navigation.types';
import {UserComponent} from '@app/layout/common/user/user.component';
import {Subject, takeUntil} from 'rxjs';
import {UserService} from '@app/core/user/user.service';
import {MembershipLevelInterface, TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {toSignal} from '@angular/core/rxjs-interop';
import {TranslocoPipe} from '@ngneat/transloco';
import {MembershipLevelsEnum} from '@app/app.constants';
import {NgClass, NgIf} from '@angular/common';
import {MatDialog} from '@angular/material/dialog';
import {BugReportComponent} from '@app/modules/client/bug-report/components/bug-report.component';
import {environment} from '@environments/environment';
import {MenuClientComponent} from '@app/core/navigation/components/menu-client.component';

@Component({
	selector: 'centered-layout',
	templateUrl: './centered.component.html',
	encapsulation: ViewEncapsulation.None,
	standalone: true,
	imports: [
		FuseLoadingBarComponent,
		FuseVerticalNavigationComponent,
		FuseHorizontalNavigationComponent,
		MatButtonModule,
		MatIconModule,
		UserComponent,
		RouterOutlet,
		TranslocoPipe,
		NgClass,
		NgIf,
		MenuClientComponent,
	],
})
export class CenteredLayoutComponent implements OnInit, OnDestroy {
	navigation: Navigation;
	isScreenSmall: boolean;
	private _unsubscribeAll: Subject<any> = new Subject<any>();

	$user: Signal<TagpeakUser> = toSignal(this._userService.get());
	$userLevel: Signal<MembershipLevelInterface> = computed(
		toSignal(this._userService.getCurrentMembershipLevel(this.$user()?.uuid)),
	);

	/**
	 * Constructor
	 */
	constructor(
		private _navigationService: NavigationService,
		private _fuseMediaWatcherService: FuseMediaWatcherService,
		private _fuseNavigationService: FuseNavigationService,
		private _userService: UserService,
		private dialog: MatDialog,
	) {}

	// -----------------------------------------------------------------------------------------------------
	// @ Accessors
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Getter for current year
	 */
	get currentYear(): number {
		return new Date().getFullYear();
	}

	// -----------------------------------------------------------------------------------------------------
	// @ Lifecycle hooks
	// -----------------------------------------------------------------------------------------------------

	/**
	 * On init
	 */
	ngOnInit(): void {
		// Subscribe to navigation data
		this._navigationService.navigation$
			.pipe(takeUntil(this._unsubscribeAll))
			.subscribe((navigation: Navigation) => {
				this.navigation = navigation;
			});

		// Subscribe to media changes
		this._fuseMediaWatcherService.onMediaChange$
			.pipe(takeUntil(this._unsubscribeAll))
			.subscribe(({matchingAliases}) => {
				// Check if the screen is small
				this.isScreenSmall = !matchingAliases.includes('sm');
			});
	}

	/**
	 * On destroy
	 */
	ngOnDestroy(): void {
		// Unsubscribe from all subscriptions
		this._unsubscribeAll.next(null);
		this._unsubscribeAll.complete();
	}

	// -----------------------------------------------------------------------------------------------------
	// @ Public methods
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Toggle navigation
	 *
	 * @param name
	 */
	toggleNavigation(name: string): void {
		// Get the navigation
		const navigation = this._fuseNavigationService.getComponent<MenuClientComponent>(name);

		if (navigation) {
			// Toggle the opened status
			navigation.toggle();
		}
	}

	openReportBug(): void {
		if (!environment.features.bugReport) {
			return;
		}

		this.dialog.open(BugReportComponent, {
			panelClass: 'onboarding-dialog-panel.smaller',
			position: {right: '2rem', bottom: '4rem'},
			data: {
				email: this.$user().email,
				name: this.$user().firstName + ' ' + this.$user().lastName,
			},
		});
	}

	openTagpeak(): void {
		window.open(environment.commercial.base, '_blank');
	}

	protected readonly MembershipLevelsEnum = MembershipLevelsEnum;
	protected readonly environment = environment;
	protected readonly signal = signal;
}
