import {
	Component,
	computed,
	DestroyRef,
	inject,
	input,
	OnInit,
	Signal,
	signal,
	WritableSignal,
} from '@angular/core';
import {CommonModule, NgOptimizedImage} from '@angular/common';
import {TranslocoPipe, TranslocoService} from '@ngneat/transloco';
import {SharedModule} from '@app/shared/shared.module';
import {FormControl, FormGroup} from '@angular/forms';
import {PublicStoreService} from '@app/modules/stores/services/public-store.service';
import {CountriesService} from '@app/modules/admin/countries/services/countries.service';
import {CountryInterface} from '@app/modules/admin/countries/interfaces/country.interface';
import {PaginationInterface} from '@app/shared/interfaces/pagination.interface';
import {MatIcon} from '@angular/material/icon';
import {CategoriesService} from '@app/modules/admin/categories/services/categories.service';
import {CategoriesInterface} from '@app/modules/admin/categories/models/categories.interface';
import {StoresInterface} from '@app/modules/stores/models/stores.interface';
import {SortListEnum} from '@app/utils/status.enum';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {
	BehaviorSubject,
	catchError,
	combineLatest,
	debounceTime,
	distinctUntilChanged,
	map,
	of,
	startWith,
	switchMap,
	take,
	tap,
} from 'rxjs';
import {CustomDropdownComponent} from '@app/shared/components/custom-dropdown/custom-dropdown.component';
import {Router} from '@angular/router';
import {DialogInfoComponent} from '@app/shared/components/dialog-info/dialog-info.component';
import {MatDialog} from '@angular/material/dialog';
import {
	CssDropdownInterface,
	IconDropdownInterface,
	OptionDropdownInterface,
} from '@app/shared/components/custom-dropdown/models/custom-dropdown.interface';
import {AppConstants} from '@app/app.constants';
import {AuthService} from '@app/core/auth/auth.service';
import {UrlQueryService} from '@app/shared/services/url-query.service';
import {environment} from '@environments/environment';

@Component({
	selector: 'tagpeak-stores',
	standalone: true,
	imports: [
		CommonModule,
		SharedModule,
		MatIcon,
		TranslocoPipe,
		CustomDropdownComponent,
		NgOptimizedImage,
	],
	templateUrl: './stores.component.html',
	styleUrl: './stores.component.scss',
})
export class StoresComponent implements OnInit {
	private _translocoService: TranslocoService = inject(TranslocoService);
	private publicStoreService: PublicStoreService = inject(PublicStoreService);
	private _countriesService: CountriesService = inject(CountriesService);
	private _categoriesService: CategoriesService = inject(CategoriesService);
	private _destroyRef: DestroyRef = inject(DestroyRef);
	private _router: Router = inject(Router);
	private dialog: MatDialog = inject(MatDialog);
	private authService = inject(AuthService);
	private urlQueryService = inject(UrlQueryService);

	page$: BehaviorSubject<number> = new BehaviorSubject(0);
	sort$: BehaviorSubject<string> = new BehaviorSubject(SortListEnum.MOST_POPULAR.toString());

	pageSize: number = 5;
	pageSizeStores: number = 100;
	totalSize: number = 1;
	currentPage: number = 0;
	totalPages: number;
	sort: string = 'name asc';

	filterForm = new FormGroup({
		country: new FormControl<CountryInterface>(null),
		category: new FormControl<CategoriesInterface>(null),
		sortValue: new FormControl<SortListEnum>(SortListEnum.MOST_POPULAR),
		email: new FormControl<string>(''),
	});
	filters$ = this.filterForm.valueChanges;
	searchFilterForm = new FormControl<string>('');
	searchText$: BehaviorSubject<string> = new BehaviorSubject<string>('');

	$stores: WritableSignal<StoresInterface[]> = signal([]);
	$countries: WritableSignal<CountryInterface[]> = signal([]);
	$mappedCountries: Signal<OptionDropdownInterface[]> = computed(() =>
		this.$countries()
			? this.$countries().map((value) => ({
					value: value.uuid,
					label: value.name,
				}))
			: [],
	);
	$categories: WritableSignal<CategoriesInterface[]> = signal([]);
	$mappedCategories: Signal<{value: string; label: string}[]> = computed(() =>
		this.$categories()
			? this.$categories().map((value) => ({
					value: value.uuid,
					label: value.name,
				}))
			: [],
	);
	$userCountry: WritableSignal<any> = signal(null);
	$authenticated: Signal<boolean> = toSignal(this.authService.check());

	sortList: {label: string; value: SortListEnum}[] = [
		{
			label: this._translocoService.translate('stores.sort-list.latest-stores'),
			value: SortListEnum.LATEST_STORES,
		},
		{
			label: this._translocoService.translate('stores.sort-list.title'),
			value: SortListEnum.BY_TITLE,
		},
		{
			label: this._translocoService.translate('stores.sort-list.most-popular'),
			value: SortListEnum.MOST_POPULAR,
		},
	];

	styleDropdown: CssDropdownInterface = {
		...AppConstants.DEFAULT_STYLE_DROPDOWN,
		styleHeader: 'bg-white rounded-full w-full h-13',
	};
	iconCheckOption: IconDropdownInterface = AppConstants.DEFAULT_ICON_CHECK_DROPDOWN;

	$preLoadedCountry = input<CountryInterface>();
	$preLoadedCategory = input<CategoriesInterface>();
	$preLoadedSearchText = input<string>();

	currentYear: number = new Date().getFullYear();

	ngOnInit(): void {
		this.publicStoreService.loadedStore = null;
		//Setup listeners that will be triggered by control changes
		this.setCountryListener();
		this.setCategoryListener();
		this.getStores();
		this.sortControl.valueChanges
			.pipe(takeUntilDestroyed(this._destroyRef))
			.subscribe((value: SortListEnum) => this.sort$.next(value));

		//Check if there is a country as a query param, if not, search by ip
		if (this.$preLoadedCountry()) {
			this.$userCountry.set(this.$preLoadedCountry());
			this.countryControl.setValue(this.$userCountry());
		} else {
			this.publicStoreService
				.getIpAddress()
				.pipe(
					switchMap((r) =>
						this._countriesService.getCountryByCode(r.code).pipe(catchError((_) => of(r))),
					),
					take(1),
				)
				.subscribe((value) => {
					this.$userCountry.set(value);
					this.setCountryFormValueByLocation();
				});
		}

		//Set category if there is a category as a query param
		if (this.$preLoadedCategory()) {
			this.categoryControl.setValue(this.$preLoadedCategory());
		}

		// Set search text if there is a search text as a query param
		// Since the searchText$ is a BehaviorSubject, it will trigger the store search.
		// So if not search text is present, we force it.
		if (this.$preLoadedSearchText()) {
			this.searchFilterForm.setValue(this.$preLoadedSearchText());
			this.searchText$.next(this.$preLoadedSearchText());
		} else {
			this.page$.next(0);
		}
	}

	setCountryFormValueByLocation() {
		this.dialog
			.open(DialogInfoComponent, {
				data: {
					text: this._translocoService.translate('public-store.country-find.message', {
						value: this.$userCountry().name,
					}),
				},
			})
			.afterClosed()
			.subscribe((_) => {
				this.countryControl.setValue(this.$userCountry());
				this.urlQueryService.addParam('country', this.$userCountry().abbreviation);
			});
	}

	setCountryListener() {
		this.countryControl.valueChanges
			.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				map((value: any) => value?.name ?? value),
				switchMap((name: string) => {
					return this._countriesService.getCountriesPublic(300, 0, this.sort, {
						name: name,
					});
				}),
				map((res: PaginationInterface<CountryInterface>) => res.data ?? []),
				takeUntilDestroyed(this._destroyRef),
			)
			.subscribe((value) => {
				this.$countries.set(value);
			});
	}

	setCategoryListener() {
		this.categoryControl.valueChanges
			.pipe(
				startWith(null),
				debounceTime(500),
				distinctUntilChanged(),
				map((value: any) => value?.name ?? value),
				switchMap((category: string) => {
					return this._categoriesService.getAllCategoriesPublic(100, 0, this.sort, {
						name: category,
					});
				}),
				map((res: PaginationInterface<CategoriesInterface>) => res.data ?? []),
			)
			.subscribe((value) => {
				this.$categories.set(value);
			});
	}

	getStores(): void {
		combineLatest([this.page$, this.sort$, this.filters$, this.searchText$])
			.pipe(
				debounceTime(100),
				switchMap(([page, sort, filters, searchText]: [number, string, any, string]) => {
					return this.publicStoreService.getStoresPublic(
						this.pageSizeStores,
						page,
						sort,
						filters.country?.abbreviation,
						filters.category?.code,
						searchText,
					);
				}),
				tap((res: PaginationInterface<StoresInterface>) => {
					this.currentPage = res.page;
					this.totalPages = res.totalPages;
					this.totalSize = res.totalRows;
				}),
				map((res: PaginationInterface<StoresInterface>) => res.data ?? []),
				takeUntilDestroyed(this._destroyRef),
			)
			.subscribe((stores: StoresInterface[]) => {
				this.$stores.set(stores);
			});
	}

	setCountry(countrySelected: {value: string; label: string}): void {
		if (!countrySelected) {
			this.countryControl.reset();
			this.urlQueryService.addParam('country', null);
			this.$stores.set([]);
			return;
		}

		let country: CountryInterface = this.$countries().find(
			(country: CountryInterface) => country.uuid === countrySelected?.value,
		);

		this.countryControl.reset();
		this.countryControl.setValue(country);

		this.urlQueryService.addParam('country', country.abbreviation);
		this.page$.next(0);
	}

	setCategory(categorySelected: {value: string; label: string}): void {
		if (!categorySelected) {
			this.urlQueryService.addParam('category', null);
		}
		let category: CategoriesInterface = this.$categories().find(
			(category: CategoriesInterface) => category.uuid === categorySelected?.value,
		);

		this.categoryControl.setValue(category);
		if (this.countryControl.value && this.countryControl.value.abbreviation) {
			this.urlQueryService.addParam('category', category.code);
			this.page$.next(0);
		}
	}

	searchName(): void {
		this.urlQueryService.addParam('name', this.nameControl.value);
		this.searchText$.next(this.nameControl.value);
		this.page$.next(0);
	}

	showPagination(): boolean {
		return this.totalPages > 1;
	}

	nextPage(): void {
		this.page$.next(this.currentPage + 1);
	}

	previousPage(): void {
		this.page$.next(this.currentPage - 1);
	}

	changePage(page: number): void {
		this.page$.next(page);
	}

	visiblePages(): number[] {
		const pages: number[] = [];

		if (this.totalPages <= 7) {
			return Array.from({length: this.totalPages}, (_, i) => i);
		}

		pages.push(0);

		const start = Math.max(1, this.currentPage - 1);
		const end = Math.min(this.totalPages - 2, this.currentPage + 1);

		for (let i = start; i <= end; i++) {
			pages.push(i);
		}

		pages.push(this.totalPages - 1);

		return pages;
	}

	// TODO: IMPLEMENT
	submitNewsletter(): void {
		console.log('Submit Newsletter');
	}

	get countryControl(): FormControl {
		return this.filterForm.get('country') as FormControl;
	}

	get categoryControl(): FormControl {
		return this.filterForm.get('category') as FormControl;
	}

	get nameControl(): FormControl {
		return this.searchFilterForm as FormControl;
	}

	get sortControl(): FormControl {
		return this.filterForm?.get('sortValue') as FormControl;
	}

	get emailControl(): FormControl {
		return this.filterForm.get('email') as FormControl;
	}

	openStore(id: string): void {
		this._router.navigate(['/stores', id]);
	}

	openLogin(): void {
		this._router.navigate(['/sign-in']);
	}

	openSignUp(): void {
		if (this.$authenticated()) {
			this._router.navigate(['client/dashboard']);
		} else {
			this._router.navigate(['/sign-up']);
		}
	}

	openTagpeak(): void {
		window.open(environment.commercial.base, '_blank');
	}

	protected readonly Number = Number;

	openLinkedin() {
		window.open(environment.commercial.linkedIn, '_blank');
	}

	openTwitter() {
		window.open(environment.commercial.x, '_blank');
	}

	openInstagram() {
		window.open(environment.commercial.instagram, '_blank');
	}

	protected readonly environment = environment;
}
