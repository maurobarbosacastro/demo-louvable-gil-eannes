import {
	ChangeDetectionStrategy,
	ChangeDetectorRef,
	Component,
	Input,
	OnDestroy,
	OnInit,
	ViewEncapsulation,
	signal,
	WritableSignal,
	InputSignal,
	input,
} from '@angular/core';
import {Router} from '@angular/router';
import {BooleanInput} from '@angular/cdk/coercion';
import {Subject, takeUntil} from 'rxjs';
import {UserService} from '@app/core/user/user.service';
import {MatButtonModule} from '@angular/material/button';
import {MatDividerModule} from '@angular/material/divider';
import {MatIconModule} from '@angular/material/icon';
import {MatMenuModule} from '@angular/material/menu';
import {SharedModule} from '@app/shared/shared.module';
import {TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {environment} from '@environments/environment';

@Component({
	selector: 'user',
	templateUrl: './user.component.html',
	encapsulation: ViewEncapsulation.None,
	changeDetection: ChangeDetectionStrategy.OnPush,
	exportAs: 'user',
	standalone: true,
	imports: [MatButtonModule, MatDividerModule, MatIconModule, MatMenuModule, SharedModule],
})
export class UserComponent implements OnInit, OnDestroy {
	/* eslint-disable @typescript-eslint/naming-convention */
	static ngAcceptInputType_showAvatar: BooleanInput;
	/* eslint-enable @typescript-eslint/naming-convention */

	@Input() showAvatar: boolean = true;
	user: TagpeakUser;
	$avatar: WritableSignal<string> = signal('null');

	private _unsubscribeAll: Subject<any> = new Subject<any>();
    $reload: WritableSignal<boolean> = signal(false);

    $isShop: InputSignal<boolean> = input<boolean>(false);

    /**
	 * Constructor
	 */
	constructor(
		private _changeDetectorRef: ChangeDetectorRef,
		private _router: Router,
		private _userService: UserService,
	) {}

	// -----------------------------------------------------------------------------------------------------
	// @ Lifecycle hooks
	// -----------------------------------------------------------------------------------------------------

	/**
	 * On init
	 */
	ngOnInit(): void {
		// Subscribe to user changes
		this._userService.user$.pipe(takeUntil(this._unsubscribeAll)).subscribe((user: TagpeakUser) => {
            this.$reload.set(true);
			this.user = user;
			if (this.user.profilePicture) {
				const uuid = this._userService.getUUIDFromUrlImage(this.user.profilePicture);
				this.$avatar.set(
					`${environment.host}${environment.image.origin}${uuid}/profilePictureSmall.webp`,
				);
			}

			// Mark for check
			this._changeDetectorRef.markForCheck();
            //Necessary to force reload the image
            setTimeout(() => {
                this.$reload.set(false);
            }, 1000);
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
	 * Update the user status
	 *
	 * @param status
	 */
	updateUserStatus(status: string): void {
		// Return if user is not available
		if (!this.user) {
			return;
		}

		// Update the user
		this._userService
			.update(
				{
					...this.user,
				},
				this.user.uuid,
			)
			.subscribe();
	}

	/**
	 * Sign out
	 */
	signOut(): void {
		this._router.navigate(['/sign-out']);
	}

	goToProfile() {
		this._router.navigate(['/client/settings']);
	}
}
