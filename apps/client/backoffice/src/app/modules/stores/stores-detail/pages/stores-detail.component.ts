import {
	Component,
	computed,
	inject,
	input,
	OnInit,
	Signal,
	signal,
	WritableSignal,
} from '@angular/core';
import {CommonModule, NgOptimizedImage} from '@angular/common';
import {StoreInfo, StoresInterface} from '@app/modules/stores/models/stores.interface';
import {MatIcon} from '@angular/material/icon';
import {FormControl, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {TranslocoPipe} from '@ngneat/transloco';
import {ActivatedRoute, Router} from '@angular/router';
import {TagpeakCarouselComponent} from '@app/shared/components/tagpeak-carousel/tagpeak-carousel.component';
import {StoresDetailDialogComponent} from '@app/modules/stores/stores-detail/components/stores-detail-dialog/stores-detail-dialog.component';
import {MatDialog} from '@angular/material/dialog';
import {StorageService} from '@app/shared/services/storage.service';
import {differenceInDays} from 'date-fns';
import {AuthService} from '@app/core/auth/auth.service';
import {AppConstants} from '@app/app.constants';
import {toSignal} from '@angular/core/rxjs-interop';
import {tap} from 'rxjs';
import {environment} from '@environments/environment';

@Component({
	selector: 'tagpeak-stores-detail',
	standalone: true,
	imports: [
		CommonModule,
		MatIcon,
		ReactiveFormsModule,
		TranslocoPipe,
		NgOptimizedImage,
		TagpeakCarouselComponent,
	],
	templateUrl: './stores-detail.component.html',
	styleUrl: './stores-detail.component.scss',
})
export class StoresDetailComponent implements OnInit {
	private _router = inject(Router);
	private _activatedRoute = inject(ActivatedRoute);
	private _dialog = inject(MatDialog);
	private _storageService = inject(StorageService);
	private authService = inject(AuthService);

	$storeInfo: WritableSignal<StoreInfo> = signal(null);
	$store = input<StoresInterface>(); //Store is fetched in a resolver in the app.routing.ts
	$storeVisitedLast30days: Signal<boolean> = computed(() => {
		if (!this.$storeInfo()) {
			return false;
		}
		const checkedDate = this.$storeInfo().checkedDate; // Assuming checkedDate is a Date object
		const daysDifference = differenceInDays(new Date(), new Date(checkedDate));

		return daysDifference <= 30;
	});

	emailForm = new FormGroup({
		email: new FormControl<string>(''),
	});

	$authenticated: Signal<boolean> = toSignal(this.authService.check().pipe(tap(console.log)));
	currentYear: number = new Date().getFullYear();

	ngOnInit(): void {
		this.getStoreInfoInStorage();
	}

	// Get store info from storage
	getStoreInfoInStorage(): void {
		const storeInfo: StoreInfo = this._storageService.get(AppConstants.STORAGE_KEYS.STORE_INFO);
		if (!!storeInfo) {
			this.$storeInfo.set({
				checked: this._storageService.storeInfo.checked,
				checkedDate: this._storageService.storeInfo.checkedDate,
			});
		}
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

	// TODO: IMPLEMENT
	submitNewsletter(): void {
		console.log('Submit Newsletter');
	}

	get emailControl(): FormControl {
		return this.emailForm.get('email') as FormControl;
	}

	goToStore(): void {
		if (!this.$authenticated()) {
			const redirectURL = this._router.routerState.snapshot.url;
			this._router.navigate(['sign-in'], {queryParams: {redirectURL: redirectURL}});
			return;
		}

		// If store was not visited on last 30 days, show dialog
		if (!this.$storeVisitedLast30days()) {
			const dialogRef = this._dialog.open(StoresDetailDialogComponent, {
				panelClass: 'dialog-panel-store-details',
				maxWidth: '100vw',
				maxHeight: '100vh',
				height: '100%',
				width: '100%',
				data: {url: this.$store().storeUrl, uuid: this.$store().uuid},
			});

			dialogRef.afterClosed().subscribe((result) => {
				this.getStoreInfoInStorage();
			});
		} else {
			const origin = new URL(window.location.href).origin;
			const internalRoute = this._router.createUrlTree(['redirect'], {
				relativeTo: this._activatedRoute,
			});
			const url = origin + '/#' + this._router.serializeUrl(internalRoute);

			window.open(url, '_blank');
		}
	}

	processTermsAndConditions(html: string): string {
		return html
			.replace(/<p>\s*<p>\s*\n\s*\n/g, '<br><br>')
			.replace(/<p>\s*<p>/g, '')
			.replace(/<p>\s*\n\s*\n/g, '<br>')
			.replace(/<\/p>/g, '<br>')
			.replace(/<p>/g, '');
	}

	openTagpeak(): void {
		window.open(environment.commercial.base, '_blank');
	}

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
