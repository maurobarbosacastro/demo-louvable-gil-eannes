import { Component, OnInit, ViewChild, ViewEncapsulation } from '@angular/core';
import {
    AbstractControl,
    FormControl,
    NgForm,
    UntypedFormBuilder,
    UntypedFormGroup,
    Validators,
} from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { fuseAnimations } from '@fuse/animations';
import { AuthService } from '@app/core/auth/auth.service';
import { UserService } from '@app/core/user/user.service';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { SharedModule } from '@app/shared/shared.module';
import { TranslocoPipe } from '@ngneat/transloco';
import { of } from 'rxjs';
import { environment } from '@environments/environment';
import { NgOptimizedImage } from '@angular/common';
import { TagpeakCarouselComponent } from '@app/shared/components/tagpeak-carousel/tagpeak-carousel.component';
import { CustomDropdownComponent } from '@app/shared/components/custom-dropdown/custom-dropdown.component';
import {
    CssDropdownInterface,
    IconDropdownInterface,
    OptionDropdownInterface,
} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import { AppConstants } from '@app/app.constants';

@Component({
    selector: 'tagpeak-auth-profile-completion',
    templateUrl: './auth-profile-completion.component.html',
    styleUrl: 'auth-profile-completion.component.scss',
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
export class AuthProfileCompletionComponent implements OnInit {
    @ViewChild('profileNgForm') profileNgForm: NgForm;

    profileForm: UntypedFormGroup;
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

    /**
     * Constructor
     */
    constructor(
        private _authService: AuthService,
        private _userService: UserService,
        private _formBuilder: UntypedFormBuilder,
        private _router: Router,
    ) { }

    // -----------------------------------------------------------------------------------------------------
    // @ Lifecycle hooks
    // -----------------------------------------------------------------------------------------------------

    /**
     * On init
     */
    ngOnInit(): void {
        // Create the form
        this.profileForm = this._formBuilder.group({
            name: ['', Validators.required, this.moreThanOneWordValidator()],
            email: [{ value: '', disabled: true }, [Validators.required, Validators.email]],
            currency: ['', Validators.required],
        });

        if (!environment.features.currency) {
            this.profileForm.get('currency').removeValidators([Validators.required]);
            this.currencyControl.setValue('EUR');
        }

        // Get user data from auth service and populate name and email
        this._userService.get().subscribe((user) => {
            if (user) {
                const fullName = `${user.firstName} ${user.lastName}`.trim();
                this.nameControl.setValue(fullName);
                this.emailControl.setValue(user.email);
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
     * Complete profile
     */
    completeProfile(): void {
        if (this.profileForm.invalid) {
            this.profileForm.markAllAsTouched();
            this.currencyControl.markAsTouched();
            this.currencyControl.markAsDirty();
            this.currencyValidator();
            return;
        }

        const body = {
            firstName: this.nameControl.value.split(' ')[0],
            lastName: this.nameControl.value.slice(this.nameControl.value.indexOf(' ') + 1),
            currency: this.currencyControl.value.value,
            ...(this._authService.referralCodeSocialSignup && {
                referralCode: this._authService.referralCodeSocialSignup,
                referralClick: this._authService.referralClickSocialSignup,
            })
        };

        this._authService.finishSocialAuth(body).subscribe((res) => {
            this._userService.user = res;
            this._authService.referralCodeSocialSignup = null;
            this._authService.referralClickSocialSignup = null;

            this._router.navigate(['/']);
        });
    }

    get nameControl(): FormControl {
        return this.profileForm.get('name') as FormControl;
    }

    get emailControl(): FormControl {
        return this.profileForm.get('email') as FormControl;
    }

    get currencyControl(): FormControl {
        return this.profileForm.get('currency') as FormControl;
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
}
