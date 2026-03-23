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
import {finalize} from 'rxjs';
import {fuseAnimations} from '@fuse/animations';
import {FuseAlertComponent, FuseAlertType} from '@fuse/components/alert';
import {AuthService} from '@app/core/auth/auth.service';
import {Router, RouterModule} from '@angular/router';
import {MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatIconModule} from '@angular/material/icon';
import {MatInputModule} from '@angular/material/input';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {SharedModule} from '@app/shared/shared.module';
import {NgOptimizedImage} from '@angular/common';
import {TagpeakCarouselComponent} from '@app/shared/components/tagpeak-carousel/tagpeak-carousel.component';
import {TranslocoPipe} from '@ngneat/transloco';

@Component({
	selector: 'auth-forgot-password',
	templateUrl: './forgot-password.component.html',
	encapsulation: ViewEncapsulation.None,
	animations: fuseAnimations,
	standalone: true,
	imports: [
		RouterModule,
		MatButtonModule,
		MatFormFieldModule,
		MatIconModule,
		MatInputModule,
		MatProgressSpinnerModule,
		FuseAlertComponent,
		SharedModule,
		NgOptimizedImage,
		TagpeakCarouselComponent,
		TranslocoPipe,
	],
})
export class AuthForgotPasswordComponent implements OnInit {
	@ViewChild('forgotPasswordNgForm') forgotPasswordNgForm: NgForm;
	@ViewChild('carousel') carousel!: TagpeakCarouselComponent;

	alert: {type: FuseAlertType; message: string} = {
		type: 'success',
		message: '',
	};
	forgotPasswordForm: UntypedFormGroup;
	showAlert: boolean = false;

	$sent: WritableSignal<boolean> = signal(false);
	_authService = inject(AuthService);
	_formBuilder = inject(UntypedFormBuilder);
	_router = inject(Router);

	// -----------------------------------------------------------------------------------------------------
	// @ Lifecycle hooks
	// -----------------------------------------------------------------------------------------------------

	/**
	 * On init
	 */
	ngOnInit(): void {
		// Create the form
		this.forgotPasswordForm = this._formBuilder.group({
			email: ['', [Validators.required, Validators.email]],
		});
	}

	// -----------------------------------------------------------------------------------------------------
	// @ Public methods
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Send the reset link
	 */
	sendResetLink(): void {
		// Return if the form is invalid
		if (this.forgotPasswordForm.invalid) {
			this.forgotPasswordForm.markAllAsTouched();
			return;
		}

		// Disable the form
		this.forgotPasswordForm.disable();

		// Hide the alert
		this.showAlert = false;

		// Forgot password
		this._authService
			.forgotPassword(this.forgotPasswordForm.get('email').value)
			.pipe(
				finalize(() => {
					// Re-enable the form
					this.forgotPasswordForm.enable();

					// Reset the form
					this.forgotPasswordNgForm.resetForm();
				}),
			)
			.subscribe(
				(response) => {
					this.$sent.set(true);
				},
				(err) => {
					// Show the alert
					this.showAlert = true;

					if (err.status === 404) {
						this.alert = {
							type: 'error',
							message: 'Email not found! Are you sure you are already a member?',
						};
					} else {
						this.alert = {
							type: 'error',
							message: 'Something went wrong. Please, try again later!',
						};
					}
				},
			);
	}

	get emailControl(): FormControl {
		return this.forgotPasswordForm.get('email') as FormControl;
	}

	goBack() {
		this._router.navigate(['/stores']);
	}
}
