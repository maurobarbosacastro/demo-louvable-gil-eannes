import {
	Component,
	computed,
	DestroyRef,
	inject,
	input,
	OnInit,
	Signal,
	signal,
	ViewChild,
	WritableSignal,
} from '@angular/core';
import {CommonModule, NgOptimizedImage} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {UserService} from '@app/core/user/user.service';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {MatDatepicker, MatDatepickerInput} from '@angular/material/datepicker';
import {format} from 'date-fns';
import {CountryInterface} from '@app/modules/admin/countries/interfaces/country.interface';
import {TagpeakUser} from '@app/modules/users/interfaces/user.interface';
import {CountriesService} from '@app/modules/admin/countries/services/countries.service';
import {BehaviorSubject, debounceTime, distinctUntilChanged, map, startWith, switchMap} from 'rxjs';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {CustomDropdownComponent} from '@app/shared/components/custom-dropdown/custom-dropdown.component';
import {environment} from '@environments/environment';
import {
	CssDropdownInterface,
	IconDropdownInterface,
	OptionDropdownInterface,
} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import {AppConstants} from '@app/app.constants';
import {ScreenService} from '@app/core/services/screen.service';

@Component({
	selector: 'tagpeak-profile',
	standalone: true,
	imports: [
		CommonModule,
		MatIcon,
		ReactiveFormsModule,
		TranslocoPipe,
		MatDatepickerInput,
		MatDatepicker,
		NgOptimizedImage,
		CustomDropdownComponent,
	],
	templateUrl: './profile.component.html',
})
export class ProfileComponent implements OnInit {
	@ViewChild(MatDatepicker) datepicker: MatDatepicker<Date>;

	profileForm = new FormGroup({
		name: new FormControl<string>('', Validators.required),
		email: new FormControl<string>('', [Validators.required, Validators.email]),
		password: new FormControl<string>('***********', Validators.required),
		country: new FormControl<Pick<CountryInterface, 'abbreviation' | 'name'>>(
			null,
			Validators.required,
		),
		displayName: new FormControl<string>(''),
		dateOfBirth: new FormControl<string | null>(null),
		newsletter: new FormControl<boolean>(true),
		filterCountry: new FormControl<string>(''),
	});

	private transloco = inject(TranslocoService);
	private userService = inject(UserService);
	private toasterService = inject(ToasterService);
	private countryService = inject(CountriesService);
	private destroyRef = inject(DestroyRef);
	private _screenService: ScreenService = inject(ScreenService);

	$user: WritableSignal<Partial<TagpeakUser>> = signal({});
	$profilePicture: Signal<string> = computed(() => this.$user().profilePicture);
	$dropdownOpen: WritableSignal<boolean> = signal(false);
	$countries: WritableSignal<CountryInterface[]> = signal([]);
	$reload: WritableSignal<boolean> = signal(false);
	$isMobile: Signal<boolean> = toSignal(this._screenService.isMobile$);

	sort: string = 'abbreviation asc';
	sort$: BehaviorSubject<string> = new BehaviorSubject('abbreviation asc');

	$preLoadedCountry = input<CountryInterface[]>();

	currencies: OptionDropdownInterface[] = AppConstants.CURRENCIES.map((currency) => ({
		label: currency.label + ' (' + currency.symbol + ')',
		value: currency.key,
		iconName: currency.icon,
	}));
	protected readonly environment = environment;
	currencyControl = new FormControl({value: 'EUR', disabled: true});
	iconCheckOption: IconDropdownInterface = AppConstants.DEFAULT_ICON_CHECK_DROPDOWN;
	styleDropdown: CssDropdownInterface = {
		...AppConstants.DEFAULT_STYLE_DROPDOWN,
		styleHeader: 'bg-white border-b border-project-brand-500 pb-2',
		styleInput: 'text-project-licorice-500',
	};

	ngOnInit(): void {
		this.profileForm.disable();
		this.newsletterControl.enable();

		this.$countries.set(this.$preLoadedCountry());

		this.userService.user$.subscribe((user) => {
			this.$user.set(user);
			this.profileForm.get('name').setValue(user.firstName + ' ' + user.lastName);
			this.profileForm.get('email').setValue(user.email);
			this.profileForm.get('country').setValue({
				abbreviation: user.country,
				name: this.getCountryName(user.country),
			});
			this.profileForm.get('displayName').setValue(user.displayName);
			this.profileForm.get('dateOfBirth').setValue(this.convertToISO(user.birthDate));
			this.profileForm.get('newsletter').setValue(user.newsletter);
			this.currencyControl.setValue(user.currency);
		});

		this.countryControl.valueChanges
			.pipe(
				startWith(null), // Ensures the observable starts immediately
				takeUntilDestroyed(this.destroyRef),
				debounceTime(500),
				distinctUntilChanged(),
				map((value) => {
					return value.name ?? value;
				}),
				switchMap((nameFilter) =>
					this.countryService.getCountries(10, 0, this.sort, {
						name: nameFilter || '',
					}),
				),
			)
			.subscribe((value) => {
				this.$countries.set(value.data);
			});
	}

	getCountryName(abbreviation: string): string {
		return this.$countries().find((country) => country.abbreviation === abbreviation)?.name ?? '';
	}

	editField(field: string): void {
		if (this.profileForm.get(field).disabled) {
			this.profileForm.get(field).enable();
		} else {
			this.profileForm.get(field).disable();
		}
	}

	saveField(field: string) {
		let body: {[key: string]: any} = {};
		const value = this.profileForm.get(field).value;

		//TODO: CHANGE THIS TO LINK OR LIKE THE DESIGN
		if (field === 'password') return;

		switch (field) {
			case 'name':
				const [firstName = '', ...lastNameParts] = value.split(' ');
				body = {
					firstName,
					lastName: lastNameParts.join(' ') || '',
				};
				break;
			case 'dateOfBirth':
				body = {birthDate: format(new Date(value), 'dd-MM-yyyy')};
				break;
			default:
				body = {[field]: value};
		}

		if (field !== 'newsletter') {
			this.profileForm.get(field).disable();
		}
		this.userService.update(body, this.$user().uuid).subscribe(
			(next) => {
				this.$user.set(next);
				this.toasterService.showToast('success', this.transloco.translate('settings.saved'));
			},
			(_) => {
				this.toasterService.showToast('error', this.transloco.translate('settings.error-saving'));
			},
		);
	}

	selectCountry(country: string): void {
		this.countryControl.setValue(country);
		this.saveField('country');
		this.$dropdownOpen.set(false);
	}

	toggleDropdown() {
		this.$dropdownOpen.set(!this.$dropdownOpen());
	}

	// Convert dd-MM-yyyy to YYYY-MM-DD (ISO 8601)
	convertToISO(dateStr: string): string {
		const [day, month, year] = dateStr.split('-');
		return `${year}-${month}-${day}`;
	}

	// Open the datepicker
	openCalendar(): void {
		this.datepicker.open();
	}

	get nameControl(): FormControl {
		return this.profileForm.get('name') as FormControl;
	}

	get emailControl(): FormControl {
		return this.profileForm.get('email') as FormControl;
	}

	get passwordControl(): FormControl {
		return this.profileForm.get('password') as FormControl;
	}

	get displayNameControl(): FormControl {
		return this.profileForm.get('displayName') as FormControl;
	}

	get dateOfBirthControl(): FormControl {
		return this.profileForm.get('dateOfBirth') as FormControl;
	}

	get countryControl(): FormControl {
		return this.profileForm.get('country') as FormControl;
	}

	get newsletterControl(): FormControl {
		return this.profileForm.get('newsletter') as FormControl;
	}

	updateProfilePicture(evt: any) {
		const file = (evt.target as HTMLInputElement).files[0];
		//check if file is larger than 5mb
		if (file.size > 5242880) {
			this.toasterService.showToast(
				'error',
				this.transloco.translate('misc.file-too-large'),
				'bottom',
			);
			return;
		}

		this.userService
			.updateProfilePicture(file, file.name, this.$user().uuid, this.$user().profilePicture)
			.subscribe({
				next: (val) => {
					this.$reload.set(true);
					this.$user.update((user) => ({...user, profilePicture: val}));
					this.userService.user = this.$user() as TagpeakUser;
					this.toasterService.showToast(
						'success',
						this.transloco.translate('users.success'),
						'bottom',
					);
					//Necessary to force reload the image
					setTimeout(() => {
						this.$reload.set(false);
					}, 1000);
				},
				error: (_) => {
					this.toasterService.showToast('error', this.transloco.translate('users.error'), 'bottom');
				},
			});
	}
}
