import {
	Component,
	OnInit,
	signal,
	ViewChild,
	ViewEncapsulation,
	WritableSignal,
} from '@angular/core';
import {
	AbstractControl,
	FormControl,
	NgForm,
	UntypedFormBuilder,
	UntypedFormGroup,
	Validators,
} from '@angular/forms';
import {ActivatedRoute, ParamMap, Router, RouterModule} from '@angular/router';
import {fuseAnimations} from '@fuse/animations';
import {AuthService, SocialProviders} from '@app/core/auth/auth.service';
import {MatButtonModule} from '@angular/material/button';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatIconModule} from '@angular/material/icon';
import {MatInputModule} from '@angular/material/input';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {SharedModule} from '@app/shared/shared.module';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {createPasswordStrengthValidator} from '@app/utils/password';
import {filter, map, of, switchMap, tap} from 'rxjs';
import {environment} from '@environments/environment';
import {NgOptimizedImage} from '@angular/common';
import {TagpeakCarouselComponent} from '@app/shared/components/tagpeak-carousel/tagpeak-carousel.component';
import {ReferralsService} from '@app/modules/client/referrals/services/referrals.service';
import {CustomDropdownComponent} from '@app/shared/components/custom-dropdown/custom-dropdown.component';
import {Location} from '@angular/common';
import {
	CssDropdownInterface,
	IconDropdownInterface,
	OptionDropdownInterface,
} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import {AppConstants} from '@app/app.constants';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';

@Component({
	selector: 'auth-sign-up',
	templateUrl: './sign-up.component.html',
	styleUrl: 'sign-up.component.scss',
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
		CustomDropdownComponent,
	],
})
export class AuthSignUpComponent implements OnInit {
	@ViewChild('signUpNgForm') signUpNgForm: NgForm;

	signUpForm: UntypedFormGroup;
	currencies: OptionDropdownInterface[] = AppConstants.CURRENCIES.map((currency) => ({
		label: currency.label + ' (' + currency.symbol + ')',
		value: currency.key,
		iconName: currency.icon,
	}));

	protected readonly environment = environment;

	styleDropdown: CssDropdownInterface = {
		...AppConstants.DEFAULT_STYLE_DROPDOWN,
		styleHeader: 'bg-white border-b border-project-brand-500 pb-2',
		styleInput: 'text-project-licorice-500',
	};

	iconCheckOption: IconDropdownInterface = AppConstants.DEFAULT_ICON_CHECK_DROPDOWN;
	$referralCodeValid: WritableSignal<boolean> = signal(true);

	/**
	 * Constructor
	 */
	constructor(
		private _authService: AuthService,
		private _formBuilder: UntypedFormBuilder,
		private _router: Router,
		private _route: ActivatedRoute,
		private _referralService: ReferralsService,
		private _toaster: ToasterService,
		private _transloco: TranslocoService,
		private _location: Location,
	) {}

	// -----------------------------------------------------------------------------------------------------
	// @ Lifecycle hooks
	// -----------------------------------------------------------------------------------------------------

	/**
	 * On init
	 */
	ngOnInit(): void {
		// Create the form
		this.signUpForm = this._formBuilder.group({
			name: ['', Validators.required, this.moreThanOneWordValidator()],
			email: ['', [Validators.required, Validators.email]],
			password: ['', Validators.required, this.passwordValidator()],
			currency: ['', Validators.required],
			referralCode: [null],
			referralClick: [null],
		});

		if (!environment.features.currency) {
			this.signUpForm.get('currency').removeValidators([Validators.required]);
			this.currencyControl.setValue('EUR');
		}

		this._route.paramMap.subscribe((params: ParamMap): void => {
			this.signUpForm.get('referralCode').setValue(params.get('referralCode'));

			if (params.get('referralCode')) {
				this._referralService
					.validateReferralCode(params.get('referralCode'))
					.pipe(
						tap((valid: boolean): void => {
							this.$referralCodeValid.set(valid);
							if (!valid) {
								this.signUpForm.get('referralCode').setValue(null);
								this._location.replaceState('/sign-up');
							}
						}),
						filter((valid: boolean): boolean => valid),
						switchMap(() => this._referralService.createReferralClick(params.get('referralCode'))),
					)
					.subscribe((click) => {
						this.signUpForm.get('referralClick').setValue(click.uuid);
					});
			}
		});

		this.currencyControl.valueChanges.subscribe((): void => {
			this.currencyValidator();
		});
	}

	// -----------------------------------------------------------------------------------------------------
	// @ Public methods
	// -----------------------------------------------------------------------------------------------------

	/**
	 * Sign up
	 */
	signUp(): void {
		if (this.signUpForm.invalid) {
			this.signUpForm.markAllAsTouched();
			this.currencyControl.markAsTouched();
			this.currencyControl.markAsDirty();
			this.currencyValidator();
			return;
		}

		// Disable the form
		this.signUpForm.disable();
		this.currencyControl.disable();

		const body = {
			firstName: this.signUpForm.value.name.split(' ')[0],
			lastName: this.signUpForm.value.name.slice(this.signUpForm.value.name.indexOf(' ') + 1),
			email: this.signUpForm.value.email,
			password: btoa(this.signUpForm.value.password),
			currency: this.signUpForm.value.currency.value,
			referralCode: this.signUpForm.value.referralCode,
			referralClick: this.signUpForm.value.referralClick,
		};

		this._authService.signUp(body).subscribe({
			next: (_) => {
				this.signUpNgForm.resetForm();
				this.currencyControl.reset();
				this.currencyValidator();
				this._router.navigateByUrl('/confirmation-required');
				return;
			},
			error: (error) => {
				// Re-enable the form
				this.signUpForm.enable();
				this.currencyControl.enable();

				// Reset the form
				this.signUpNgForm.resetForm();
				this.currencyControl.reset();
				this.currencyValidator();
				if (error.status == 409) {
					this._toaster.showToast(
						'error',
						this._transloco.translate('auth.email-taken'),
						'top',
						null,
						10000,
					);
				}
			},
		});
	}

	loginWithSocial(social: SocialProviders) {
		this._authService.loginWithSocials(
			social,
			JSON.stringify({
				referralCode: this.signUpForm.get('referralCode').value,
				referralClick: this.signUpForm.get('referralClick').value,
			}),
		);
	}

	get nameControl(): FormControl {
		return this.signUpForm.get('name') as FormControl;
	}

	get emailControl(): FormControl {
		return this.signUpForm.get('email') as FormControl;
	}

	get passwordControl(): FormControl {
		return this.signUpForm.get('password') as FormControl;
	}

	get currencyControl(): FormControl {
		return this.signUpForm.get('currency') as FormControl;
	}

	passwordValidator() {
		return function (control: AbstractControl) {
			return of(!createPasswordStrengthValidator(control.value) ? {passwordStrength: true} : null);
		};
	}

	moreThanOneWordValidator() {
		return (control: AbstractControl) => {
			return of(!this.moreThanOneWord(control.value) ? {moreThanOneWord: true} : null);
		};
	}

	moreThanOneWord(value: string) {
		const wordCount = value.trim().split(/\s+/).length;
		return wordCount >= 2;
	}

	goBack() {
		this._router.navigate(['/stores']);
	}

	setCurrencyValue(value: OptionDropdownInterface): void {
		this.currencyControl.setValue(value);
	}

	currencyValidator(): void {
		if (this.currencyControl.dirty && this.currencyControl.invalid) {
			this.styleDropdown.styleHeader = 'bg-white border-b border-project-red pb-2';
		} else {
			this.styleDropdown.styleHeader = 'bg-white border-b border-project-brand-500 pb-2';
		}
	}

	getCurrencyValue(value: OptionDropdownInterface): string {
		if (!value) return null;

		const currency: OptionDropdownInterface = this.currencies.find(
			(currency: OptionDropdownInterface): boolean => currency.value === value.value,
		);
		return currency.label;
	}

	protected readonly SocialProviders = SocialProviders;
}
