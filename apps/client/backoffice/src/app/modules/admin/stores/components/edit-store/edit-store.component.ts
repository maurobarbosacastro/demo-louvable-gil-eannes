import {Component, DestroyRef, inject, OnInit, Signal, signal, WritableSignal} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {SharedModule} from '@app/shared/shared.module';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {FormControl, UntypedFormBuilder, UntypedFormGroup, Validators} from '@angular/forms';
import {ToasterService} from '@app/shared/components/toaster/toaster.service';
import {StoresService} from '@app/modules/admin/stores/services/stores.service';
import {
	AdminStoreInterface,
	CreateStoreInterface,
	LogoStore,
} from '@app/modules/admin/stores/interfaces/store.interface';
import {
	CheckboxGroupComponent,
	CheckboxGroupInterface,
} from '@app/shared/components/checkbox-group/checkbox-group.component';

import {QuillEditorComponent} from 'ngx-quill';
import {cashbackTypeOptions} from '@app/utils/constants';
import {ImagePickerComponent} from '@app/shared/components/image-upload/image-picker.component';
import {PartnersService} from '@app/modules/admin/partners/service/partners.service';
import {CountriesService} from '@app/modules/admin/countries/services/countries.service';
import {PartnersInterface} from '@app/modules/admin/partners/interfaces/partners.interface';
import {LanguagesService} from '@app/modules/admin/languages/services/languages.service';
import {CategoriesService} from '@app/modules/admin/categories/services/categories.service';
import {LanguagesInterface} from '@app/modules/admin/languages/models/languages.interface';
import {CountryInterface} from '@app/modules/admin/countries/interfaces/country.interface';
import {CategoriesInterface} from '@app/modules/admin/categories/models/categories.interface';
import {StatusEnum} from '@app/utils/status.enum';
import {
	catchError,
	debounceTime,
	delay,
	distinctUntilChanged,
	filter,
	map,
	of,
	skip,
	switchMap,
	tap,
	throwError,
} from 'rxjs';
import {takeUntilDestroyed, toObservable, toSignal} from '@angular/core/rxjs-interop';
import {TagpeakInputDirective} from '@app/shared/directive/tagpeak-input.directive';
import {ConfigurationService} from '@app/modules/admin/configuration/services/configuration.service';

@Component({
	selector: 'tagpeak-edit-store',
	standalone: true,
	imports: [
		CommonModule,
		MatIcon,
		SharedModule,
		TranslocoPipe,
		CheckboxGroupComponent,
		QuillEditorComponent,
		ImagePickerComponent,
		TagpeakInputDirective,
	],
	templateUrl: './edit-store.component.html',
})
export class EditStoreComponent implements OnInit {
	private dialogData = inject(MAT_DIALOG_DATA);
	private dialogRef = inject(MatDialogRef);
	private _countriesService = inject(CountriesService);
	private _storesService = inject(StoresService);
	private _categoriesService = inject(CategoriesService);
	private _languagesService = inject(LanguagesService);
	private _partnersService = inject(PartnersService);
	private _formBuilder = inject(UntypedFormBuilder);
	private toaster = inject(ToasterService);
	private transloco = inject(TranslocoService);
	private destroyRef = inject(DestroyRef);
	private configService = inject(ConfigurationService);

	image: File;
	$errorLogo: WritableSignal<string> = signal(null);
	imageBanner: File;
	$errorBanner: WritableSignal<string> = signal(null);

	$logoNotAvailable: WritableSignal<boolean> = signal(false);

	$isEdit: WritableSignal<boolean> = signal(this.dialogData && !!this.dialogData.uuid);
	$isApproval: WritableSignal<boolean> = signal(this.dialogData && this.dialogData.approval);

	$store: Signal<AdminStoreInterface> = toSignal(
		toObservable(this.$isEdit).pipe(
			takeUntilDestroyed(this.destroyRef),
			filter((val) => !!val),
			switchMap((_) => this._storesService.getStore(this.dialogData.uuid)),
			map((val) => {
				this.storeForm.patchValue(val);
				this.stateControl.setValue(val.state.toLocaleLowerCase() == 'active');
				this.cashbackTypeControl.setValue(this.cashbackTypeControl.value.toUpperCase());
				return val;
			}),
		),
	);
	countries: CountryInterface[] = [];
	categories: CategoriesInterface[] = [];
	$countries: Signal<CheckboxGroupInterface[]> = toSignal(
		this._countriesService.getCountries(100, 0, 'name').pipe(
			delay(100),
			takeUntilDestroyed(this.destroyRef),
			map((val) => {
				return val.data.map((country) => {
					return this.createCheckboxGroupItem(
						country.abbreviation,
						country.name,
						this.$store() && this.$store()?.country?.includes(country.abbreviation),
					);
				});
			}),
		),
	);
	$categories: Signal<CheckboxGroupInterface[]> = toSignal(
		this._categoriesService.getAllCategories(100, 0, 'code').pipe(
			delay(100),
			takeUntilDestroyed(this.destroyRef),
			map((val) => {
				return val.data.map((categories) => {
					return this.createCheckboxGroupItem(
						categories.code,
						categories.name,
						this.$store() && this.$store()?.category?.includes(categories.code),
					);
				});
			}),
		),
	);
	$languages: Signal<LanguagesInterface[]> = toSignal(
		this._languagesService.getAllLanguages(100, 0, 'code').pipe(
			delay(100),
			takeUntilDestroyed(this.destroyRef),
			map((val) => val.data),
		),
	);
	$affiliatePartners: Signal<PartnersInterface[]> = toSignal(
		this._partnersService.getPartners(100, 0, 'code').pipe(
			delay(100),
			takeUntilDestroyed(this.destroyRef),
			map((val) => val.data),
		),
	);

	cashbackOptions: {label: string; value: string}[] = cashbackTypeOptions;

	storeForm: UntypedFormGroup;
	modules = {
		toolbar: [
			['bold', 'italic', 'underline', 'strike'], // toggled buttons
			['blockquote', 'code-block'],

			[{'list': 'ordered'}, {'list': 'bullet'}], // superscript/subscript

			[{'size': ['small', false, 'large', 'huge']}], // custom dropdown
			[{'header': [1, 2, 3, 4, 5, 6, false]}],

			[{'font': ['IBM Plex Sans']}], // fonts

			['link'], // link and image, video
		],
	};

	loadedLogoByDomainSearch: LogoStore = null;
	$overrideFeePlaceholder: WritableSignal<string> = signal('');
	$isLoading: WritableSignal<boolean> = signal(false);

	ngOnInit(): void {
		this.storeForm = this._formBuilder.group({
			name: ['', Validators.required],
			logo: [null],
			banner: [null],
			shortDescription: [null],
			description: [null],
			urlSlug: ['', Validators.required],
			averageRewardActivationTime: ['', Validators.required],
			state: [true, Validators.required],
			keywords: [null],
			affiliateLink: [null],
			storeUrl: [null],
			termsAndConditions: [null],
			cashbackType: ['', Validators.required],
			cashbackValue: [0, Validators.required],
			percentageCashout: [0, Validators.required],
			overrideFee: [null, [Validators.min(0), Validators.max(100)]],
			metaTitle: [null],
			metaKeywords: [null],
			metaDescription: [null],
			country: ['', Validators.required],
			category: ['', Validators.required],
			languageCode: [''],
			affiliatePartnerCode: ['', Validators.required],
			partnerIdentity: [''],
			position: [null],
		});

		const pattern = /^(?:https?:\/\/)?(?:www\.)?([a-zA-Z0-9-]+(?:\.[a-zA-Z]{2,})+)/;
		this.storeUrlControl.valueChanges
			.pipe(
				skip(1),
				takeUntilDestroyed(this.destroyRef),
				debounceTime(500),
				distinctUntilChanged(),
				filter(
					(val) =>
						val.trim().length > 0 &&
						val.trim().match(pattern) &&
						val.trim().match(pattern).length > 0,
				),
				tap(() => this.$isLoading.update(() => true)),
				map((val) => val.trim().match(pattern)[1]),
				switchMap((val) => this._storesService.getStoreLogoFromDomain(val)),
				catchError((err) => {
					this.$isLoading.update(() => false);
					this.toaster.showToast('error', this.transloco.translate('store.error-get-logo'));
					throw err;
				}),
			)
			.subscribe((res: LogoStore) => {
				if (res.id === '') this.$logoNotAvailable.set(true);
				else {
					this.$logoNotAvailable.set(false);
					this.$isLoading.update(() => false);
					if (!this.logoControl.value) {
						this.loadedLogoByDomainSearch = res;
					}
					this.logoControl.setValue(res.original);
				}
			});

		this.$overrideFeePlaceholder.set(
			this.transloco.translate('store.transaction-fee-placeholder', {
				value: this.configService.configs['transaction_fixed_fee'].value,
			}),
		);

		/*this.positionControl.valueChanges
			.pipe(
				takeUntilDestroyed(this.destroyRef),
				debounceTime(500),
				distinctUntilChanged(),
				switchMap((value) => this._storesService.positionIsUnique(value)),
			)
			.subscribe((res) => {
				if (!res) {
					this.positionControl.setErrors({position: true});
				}
			});*/
	}

	createCheckboxGroupItem(
		key: string,
		name: string,
		selected: boolean = false,
	): CheckboxGroupInterface {
		return {key, name, selected};
	}

	saveStore() {
		event.stopPropagation();

		if (this.storeForm.invalid) {
			this.storeForm.markAllAsTouched();
			return;
		}

		let dto: CreateStoreInterface = this.storeForm.value;
		dto.state =
			this.$store().state !== 'PENDING'
				? this.stateControl.value
					? 'ACTIVE'
					: 'INACTIVE'
				: this.$store().state;
		dto.languageCode = this.languageControl.value || null;
		dto.affiliatePartner = this.affiliatePartnerControl.value || null;
		delete dto.logo;
		delete dto.banner;

		if (this.loadedLogoByDomainSearch) {
			dto.logo = this.loadedLogoByDomainSearch.id;
		}

		const call$ = this.$isEdit()
			? this._storesService.updateStore(this.dialogData.uuid, dto)
			: this._storesService.createStore(dto);

		call$
			.pipe(
				switchMap((_) => {
					if (this.image) {
						return this._storesService.uploadLogo(this.dialogData.uuid, this.image);
					}

					if (!this.image && this.$store() && this.$store().logo && !this.logoControl.value) {
						return this._storesService.deleteLogo(this.dialogData.uuid);
					}

					return of(_);
				}),
				switchMap((_) => {
					if (this.imageBanner) {
						return this._storesService.uploadBanner(this.dialogData.uuid, this.imageBanner);
					}

					if (
						!this.imageBanner &&
						this.$store() &&
						this.$store().banner &&
						!this.bannerControl.value
					) {
						return this._storesService.deleteBanner(this.dialogData.uuid);
					}

					return of(_);
				}),
				catchError((err) => {
					if (this.$isEdit()) {
						this.toaster.showToast('error', this.transloco.translate('store.error-update'));
					} else {
						this.toaster.showToast('error', this.transloco.translate('store.error-create'));
					}
					return throwError(() => err);
				}),
			)
			.subscribe((_) => {
				if (this.$isEdit()) {
					this.toaster.showToast('success', this.transloco.translate('store.success-update'));
				} else {
					this.toaster.showToast('success', this.transloco.translate('store.success-create'));
				}
				this.dialogRef.close({skipFetch: false});
			});
	}

	get nameControl(): FormControl {
		return this.storeForm.get('name') as FormControl;
	}

	get logoControl(): FormControl {
		return this.storeForm.get('logo') as FormControl;
	}

	get bannerControl(): FormControl {
		return this.storeForm.get('banner') as FormControl;
	}

	get shortDescriptionControl(): FormControl {
		return this.storeForm.get('shortDescription') as FormControl;
	}

	get descriptionControl(): FormControl {
		return this.storeForm.get('description') as FormControl;
	}

	get urlSlugControl(): FormControl {
		return this.storeForm.get('urlSlug') as FormControl;
	}

	get averageRewardActivationTimeControl(): FormControl {
		return this.storeForm.get('averageRewardActivationTime') as FormControl;
	}

	get stateControl(): FormControl {
		return this.storeForm.get('state') as FormControl;
	}

	get keywordsControl(): FormControl {
		return this.storeForm.get('keywords') as FormControl;
	}

	get storeUrlControl(): FormControl {
		return this.storeForm.get('storeUrl') as FormControl;
	}

	get termsAndConditionsControl(): FormControl {
		return this.storeForm.get('termsAndConditions') as FormControl;
	}

	get cashbackTypeControl(): FormControl {
		return this.storeForm.get('cashbackType') as FormControl;
	}

	get cashbackValueControl(): FormControl {
		return this.storeForm.get('cashbackValue') as FormControl;
	}

	get percentageCashoutControl(): FormControl {
		return this.storeForm.get('percentageCashout') as FormControl;
	}

	get overrideFeeControl(): FormControl {
		return this.storeForm.get('overrideFee') as FormControl;
	}

	get metaTitleControl(): FormControl {
		return this.storeForm.get('metaTitle') as FormControl;
	}

	get metaKeywordsControl(): FormControl {
		return this.storeForm.get('metaKeywords') as FormControl;
	}

	get metaDescriptionControl(): FormControl {
		return this.storeForm.get('metaDescription') as FormControl;
	}

	get countryControl(): FormControl {
		return this.storeForm.get('country') as FormControl;
	}

	get categoryControl(): FormControl {
		return this.storeForm.get('category') as FormControl;
	}

	get languageControl(): FormControl {
		return this.storeForm.get('languageCode') as FormControl;
	}

	get affiliatePartnerControl(): FormControl {
		return this.storeForm.get('affiliatePartnerCode') as FormControl;
	}

	get affiliatePartnerIdentityControl(): FormControl {
		return this.storeForm.get('partnerIdentity') as FormControl;
	}

	get affiliateLinkControl(): FormControl {
		return this.storeForm.get('affiliateLink') as FormControl;
	}

	get positionControl(): FormControl {
		return this.storeForm.get('position') as FormControl;
	}

	closeModal() {
		this.dialogRef.close({skipFetch: true});
	}

	onCountrySelected(value: string[]) {
		this.countryControl.markAsTouched();
		this.countryControl.setValue(value);
	}

	onCategorySelected(value: string[]) {
		this.categoryControl.markAsTouched();
		this.categoryControl.setValue(value);
	}

	outputImage($event: any): void {
		if ($event) {
			this.image = $event;
		} else {
			this.image = null;
			this.logoControl.setValue(null);
		}
		this.$errorLogo.set(null);
	}

	outputImageBanner($event: any): void {
		if ($event) {
			this.imageBanner = $event;
		} else {
			this.imageBanner = null;
			this.bannerControl.setValue(null);
		}
		this.$errorBanner.set(null);
	}

	protected readonly StatusEnum = StatusEnum;

	reject() {
		let dto: CreateStoreInterface = this.storeForm.value;
		dto.state = 'INACTIVE';
		dto.languageCode = this.languageControl.value || null;
		dto.affiliatePartner = this.affiliatePartnerControl.value || null;
		delete dto.logo;
		delete dto.banner;

		this._storesService.updateStore(this.$store().uuid, dto).subscribe((_) => {
			this.toaster.showToast('success', this.transloco.translate('store.success-update'));
			this.dialogRef.close({skipFetch: false});
		});
	}

	accept() {
		let dto: CreateStoreInterface = this.storeForm.value;
		dto.state = 'ACTIVE';
		dto.languageCode = this.languageControl.value || null;
		dto.affiliatePartner = this.affiliatePartnerControl.value || null;
		delete dto.logo;
		delete dto.banner;

		this._storesService
			.updateStore(this.$store().uuid, dto)
			.pipe(
				switchMap((_) => {
					if (this.image) {
						return this._storesService.uploadLogo(this.dialogData.uuid, this.image);
					}
					return of(_);
				}),
				switchMap((_) => {
					if (this.imageBanner) {
						return this._storesService.uploadBanner(this.dialogData.uuid, this.imageBanner);
					}

					return of(_);
				}),
				catchError((err) => {
					this.toaster.showToast('error', this.transloco.translate('store.error-update'));
					return throwError(() => err);
				}),
			)
			.subscribe((_) => {
				this.toaster.showToast('success', this.transloco.translate('store.success-update'));
				this.dialogRef.close({skipFetch: false});
			});
	}
}
