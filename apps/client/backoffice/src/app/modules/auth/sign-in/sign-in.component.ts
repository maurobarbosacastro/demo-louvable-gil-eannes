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
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {NgOptimizedImage} from '@angular/common';
import {TagpeakCarouselComponent} from '@app/shared/components/tagpeak-carousel/tagpeak-carousel.component';
import {UserService} from '@app/core/user/user.service';
import {AccountVerifiedDialogComponent} from '@app/modules/auth/account-verified-dialog/account-verified-dialog.component';
import {MatDialog} from '@angular/material/dialog';
import {TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {environment} from '@environments/environment';
import {StorageService} from '@app/shared/services/storage.service';
import {catchError, switchMap, throwError} from 'rxjs';
import {UrlQueryService} from '@app/shared/services/url-query.service';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';

@Component({
	selector: 'auth-sign-in',
	templateUrl: './sign-in.component.html',
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
export class AuthSignInComponent implements OnInit {
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
	private _urlQueryService = inject(UrlQueryService);
	private _toastService = inject(ToasterService);
	private _transloco = inject(TranslocoService)

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

		if (this._urlQueryService.getParam("email")){
			this._toastService.showToast("success", this._transloco.translate('users.success-email-confirmation'),"top", null, 10000)
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
			)
			.subscribe(
				(user) => {
					this.user = user;
					// Set the redirect url.
					// The '/signed-in-redirect' is a dummy url to catch the request and redirect the user
					// to the correct page after a successful sign in. This way, that url can be set via
					// routing file and we don't have to touch here.
					const queryParams = this._activatedRoute.snapshot.queryParamMap;
					const redirectURL = queryParams.get('redirectURL');

					const finalRedirectURL =
						redirectURL && redirectURL.includes('sign-in')
							? '/signed-in-redirect'
							: redirectURL || '/signed-in-redirect';

					// Navigate to the redirect url
					this._router.navigateByUrl(finalRedirectURL);

					// Show modal on first login after account verification
					if (!user.isVerified) {
						this.openAccountVerifiedDialog();
					}
				},
				(response) => {
					// Re-enable the form
					this.signInForm.enable();

					if (response.status == 401 && response.error.error == 'invalid_grant') {
						this.wrongCredentials = true;
					}
				},
			);
	}

	openAccountVerifiedDialog(): void {
		const dialogRef = this.dialog.open(AccountVerifiedDialogComponent, {
			panelClass: 'verify-dialog-panel',
			data: this.user.uuid,
		});

		dialogRef.afterClosed().subscribe((data) => {
			this.userService.setEmailVerified(this.user.uuid, true).subscribe();

			//If user click on Lets go, skip onboarding
			if (data && data.skipOnboarding) {
				this.userService.update({onboardingFinished: 'true'}, this.user.uuid).subscribe();
			}
		});
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
		this._router.navigate(['/stores']);
	}

	protected readonly SocialProviders = SocialProviders;
}
