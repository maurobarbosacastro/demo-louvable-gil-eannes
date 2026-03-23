import {
	Component,
	inject,
	OnInit,
	signal,
	ViewChild,
	ViewEncapsulation,
	WritableSignal,
} from '@angular/core';
import {
	FormControl,
	NgForm,
	UntypedFormBuilder,
	UntypedFormGroup,
	Validators,
} from '@angular/forms';
import {ActivatedRoute, Router, RouterModule} from '@angular/router';
import {fuseAnimations} from '@fuse/animations';
import {FuseAlertType} from '@fuse/components/alert';
import {AuthService, SocialProviders} from '@app/core/auth/auth.service';
import {MatButtonModule} from '@angular/material/button';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatIconModule} from '@angular/material/icon';
import {MatInputModule} from '@angular/material/input';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {SharedModule} from '@app/shared/shared.module';
import {TranslocoPipe} from '@ngneat/transloco';
import {Location, NgOptimizedImage} from '@angular/common';
import {TagpeakCarouselComponent} from '@app/shared/components/tagpeak-carousel/tagpeak-carousel.component';
import {UserService} from '@app/core/user/user.service';
import {MatDialog} from '@angular/material/dialog';
import {TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {environment} from '@environments/environment';
import {StorageService} from '@app/shared/services/storage.service';
import {catchError, map, switchMap, throwError} from 'rxjs';
import {ShopService} from '@app/modules/shop/services/shop.service';
import {UrlQueryService} from '@app/shared/services/url-query.service';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';

@Component({
	selector: 'auth-shop-sign-in',
	templateUrl: './shop-sign-in.component.html',
	encapsulation: ViewEncapsulation.None,
	animations: fuseAnimations,
	standalone: true,
	imports: [
		RouterModule,
		MatButtonModule,
		MatCheckboxModule,
		MatFormFieldModule,
		MatIconModule,
		MatInputModule,
		MatProgressSpinnerModule,
		SharedModule,
		TranslocoPipe,
		NgOptimizedImage,
		TagpeakCarouselComponent,
	],
})
export class AuthShopSignInComponent implements OnInit {
	@ViewChild('signInNgForm') signInNgForm: NgForm;

	@ViewChild('carousel') carousel!: TagpeakCarouselComponent;

	alert: {type: FuseAlertType; message: string};
	signInForm: UntypedFormGroup;
	wrongCredentials: boolean = false;
	$needResetPassword: WritableSignal<boolean> = signal<boolean>(false);
	userService = inject(UserService);
	dialog = inject(MatDialog);
	user: TagpeakUser;

	protected readonly environment = environment;

	private _activatedRoute = inject(ActivatedRoute);
	private _authService = inject(AuthService);
	private _formBuilder = inject(UntypedFormBuilder);
	private _router = inject(Router);
	private _storageService = inject(StorageService);
	private _location = inject(Location);
	private _shopService = inject(ShopService);
	private _urlQueryService = inject(UrlQueryService);
	private _toastService = inject(ToasterService);

	isFromShopify: boolean = false;

	// -----------------------------------------------------------------------------------------------------
	// @ Lifecycle hooks
	// -----------------------------------------------------------------------------------------------------

	/**
	 * On init
	 */
	ngOnInit(): void {
		// Create the form
		this.signInForm = this._formBuilder.group({
			email: ['', [Validators.required, Validators.email]],
			password: ['', Validators.required],
			rememberMe: [false],
		});

		if (!!this._urlQueryService.getParam('shop')) {
			this._shopService.shopName = this._urlQueryService.getParam('shop');
			this.isFromShopify = true;
		}
	}

	// -----------------------------------------------------------------------------------------------------
	// @ Public methods
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Sign in
	 */
	signIn(event: Event): void {
		event.preventDefault();
		// Return if the form is invalid
		if (this.signInForm.invalid) {
			this.signInForm.markAllAsTouched();
			return;
		}

		// Disable the form
		this.signInForm.disable();

		// Hide the alert
		this.wrongCredentials = false;

		this._storageService.rememberMe = this.rememberMeControl.value;

		// Sign in with validation action user
		this._authService
			.validateActionUser(this.emailControl.value)
			.pipe(
				catchError((error) => {
					if (error.status === 401 && error.error.ErrorType === 'invalid_grant') {
						this.wrongCredentials = true;
					}
					if (error.error.ErrorType === 'need_reset_password') {
						this.$needResetPassword.set(true);
					}
					return throwError(() => error);
				}),
				switchMap(() =>
					this._authService.signIn(this.signInForm.value, this.rememberMeControl.value),
				),
				switchMap( (user) => {
					return this._authService.checkIfShopifyStoreOwner(this._shopService.shopName)
						.pipe(
							map( _ => user),
							catchError( (error) => {
								if (error.status === 409) {
									//log info
									this._toastService.showToast(
										"error",
										`The user ${this.emailControl.value} is not the owner of the shop ${this._shopService.shopName.split('.')[0]}. Please contact the store owner to grant access.`,
										"top",
										null,
										10000
									)
								}
								return throwError(() => error);
							})
						)
				})
			)
			.subscribe(
				(user) => {
					this.user = user;
					const shop = this._shopService.shopName.split('.')[0];
					window.location.href = `https://admin.shopify.com/store/${shop}/apps/${environment.shopify.appName}/app?tok=${this._authService.accessToken}`;
				},
				(response) => {
					this._authService.resetLoginAttempt()
					// Re-enable the form
					this.signInForm.enable();

					if (response.status == 401 && response.error.error == 'invalid_grant') {
						this.wrongCredentials = true;
					}
				},
			);
	}

	get emailControl(): FormControl {
		return this.signInForm.get('email') as FormControl;
	}

	get passwordControl(): FormControl {
		return this.signInForm.get('password') as FormControl;
	}

	get rememberMeControl(): FormControl {
		return this.signInForm.get('rememberMe') as FormControl;
	}

	loginWithSocial(social: SocialProviders) {
		this._authService.loginWithSocials(social);
	}

	goBack() {
		this._location.back();
	}

	protected readonly SocialProviders = SocialProviders;
}
