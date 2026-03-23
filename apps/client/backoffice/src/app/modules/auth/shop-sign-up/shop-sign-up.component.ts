import { Component, OnInit, ViewChild, ViewEncapsulation } from '@angular/core';
import {
    AbstractControl,
    FormControl,
    NgForm,
    UntypedFormBuilder,
    UntypedFormGroup,
    Validators,
} from '@angular/forms';
import { ActivatedRoute, ParamMap, Router, RouterModule } from '@angular/router';
import { fuseAnimations } from '@fuse/animations';
import { AuthService, SocialProviders } from '@app/core/auth/auth.service';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { SharedModule } from '@app/shared/shared.module';
import { TranslocoPipe, TranslocoService } from '@ngneat/transloco';
import { createPasswordStrengthValidator } from '@app/utils/password';
import { of, switchMap } from 'rxjs';
import { environment } from '@environments/environment';
import { NgOptimizedImage, Location } from '@angular/common';
import { TagpeakCarouselComponent } from '@app/shared/components/tagpeak-carousel/tagpeak-carousel.component';
import { ShopService } from '@app/modules/shop/services/shop.service';
import { CustomDropdownComponent } from '@app/shared/components/custom-dropdown/custom-dropdown.component';
import { CssDropdownInterface, IconDropdownInterface, OptionDropdownInterface } from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import { AppConstants } from '@app/app.constants';
import { ToasterService } from '@app/shared/components/toaster/toaster.service';

@Component({
    selector: 'auth-shop-sign-up',
    templateUrl: './shop-sign-up.component.html',
    styleUrl: 'shop-sign-up.component.scss',
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
export class AuthShopSignUpComponent implements OnInit {
    @ViewChild('signUpNgForm') signUpNgForm: NgForm;

    signUpForm: UntypedFormGroup;
    currencies: OptionDropdownInterface[] = AppConstants.CURRENCIES.map((currency) => ({
        label: currency.label + ' (' + currency.symbol + ')',
        value: currency.key,
        iconName: currency.icon,
    }));

    styleDropdown: CssDropdownInterface = {
        ...AppConstants.DEFAULT_STYLE_DROPDOWN,
        styleHeader: 'bg-white border-b border-project-brand-500 pb-2',
        styleInput: 'text-project-licorice-500',
    };
    iconCheckOption: IconDropdownInterface = AppConstants.DEFAULT_ICON_CHECK_DROPDOWN;

    protected readonly environment = environment;

    /**
     * Constructor
     */
    constructor(
        private _shopService: ShopService,
        private _authService: AuthService,
        private _formBuilder: UntypedFormBuilder,
        private _router: Router,
        private _route: ActivatedRoute,
        private _location: Location,
        private _toaster: ToasterService,
        private _transloco: TranslocoService
    ) { }

    // -----------------------------------------------------------------------------------------------------
    // @ Lifecycle hooks
    // -----------------------------------------------------------------------------------------------------

    /**
     * On init
     */
    ngOnInit(): void {
        this._shopService.shopName = null;
        // Create the form
        this.signUpForm = this._formBuilder.group({
            name: ['', Validators.required, this.moreThanOneWordValidator()],
            email: ['', [Validators.required, Validators.email]],
            password: ['', Validators.required, this.passwordValidator()],
            currency: ['', Validators.required],
            shop: ['', Validators.required],
        });

        if (!environment.features.currency) {
            this.signUpForm.get('currency').removeValidators([Validators.required]);
            this.currencyControl.setValue('EUR');
        }

        this._route.queryParamMap.subscribe((params: ParamMap) => {
            this._shopService.shopName = params.get('shop');
            this.signUpForm.get('shop').setValue(params.get('shop'));
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

        const body = {
            firstName: this.signUpForm.value.name.split(' ')[0],
            lastName: this.signUpForm.value.name.slice(this.signUpForm.value.name.indexOf(' ') + 1),
            email: this.signUpForm.value.email,
            password: btoa(this.signUpForm.value.password),
            currency: this.signUpForm.value.currency.value,
            shop: this.signUpForm.value.shop,
        };

        this._shopService.signUp(body)
            .pipe(
                switchMap(_ => this._authService.signIn({
                    email: body.email,
                    password: this.signUpForm.value.password,
                }, true))
            )
            .subscribe({
                next: (_) => {
                    const shop = this._shopService.shopName.split('.')[0];
                    setTimeout(() => {
                        window.location.href = `https://admin.shopify.com/store/${shop}/apps/${environment.shopify.appName}/app?tok=${this._authService.accessToken}`;
                    }, 1000)

                    this.signUpNgForm.resetForm();
                    this.currencyControl.reset();
                    this.currencyValidator();
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
                            "error",
                            this._transloco.translate('auth.email-taken'),
                            'top',
                            null,
                            10000
                        );
                    }
                },
            });
    }

    loginWithSocial(social: SocialProviders) {
        this._authService.loginWithSocials(social);
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
            return of(!createPasswordStrengthValidator(control.value) ? { passwordStrength: true } : null);
        };
    }

    moreThanOneWordValidator() {
        return (control: AbstractControl) => {
            return of(!this.moreThanOneWord(control.value) ? { moreThanOneWord: true } : null);
        };
    }

    moreThanOneWord(value: string) {
        const wordCount = value.trim().split(/\s+/).length;
        return wordCount >= 2;
    }

    goBack() {
        this._location.back();
    }

    setCurrencyValue(value: OptionDropdownInterface): void {
        this.currencyControl.setValue(value);
    }

    currencyValidator(): void {
        if (this.currencyControl.invalid) {
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
