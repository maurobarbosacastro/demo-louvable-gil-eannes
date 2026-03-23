import {Component, DestroyRef, inject, OnInit, signal, WritableSignal} from '@angular/core';
import {CommonModule, NgOptimizedImage} from '@angular/common';
import {finalize, map, Subject, take, takeUntil, takeWhile, tap, timer} from 'rxjs';
import {PublicStoreService} from '@app/modules/stores/services/public-store.service';
import {StoresInterface} from '../../models/stores.interface';
import {environment} from '@environments/environment';
import {HttpClient} from '@angular/common/http';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {ImageInfo} from '@app/shared/interfaces/image.interface';

@Component({
	selector: 'tagpeak-redirect',
	standalone: true,
	imports: [CommonModule, NgOptimizedImage],
	templateUrl: './redirect.component.html',
	styleUrl: './redirect.component.scss',
})
export class RedirectComponent implements OnInit {
	private publicStoreService = inject(PublicStoreService);
	private http: HttpClient = inject(HttpClient);
	private destroyRef = inject(DestroyRef);

	secondsTimer: number = 5;

	$store: WritableSignal<StoresInterface> = signal(null);
	$logo: WritableSignal<string> = signal('');
	imageUuid: string;
	$redirectUrl: WritableSignal<string> = signal('');

	manualDestroy$: Subject<void> = new Subject();

	ngOnInit(): void {
		this.$store.set(this.publicStoreService.loadedStore);
		this.imageUuid = this.$store().logo ? this.$store().logo.split('/')[4] : null;
		if (this.imageUuid) {
			this.getImageInfo();
		}

		this.publicStoreService
			.getStoreRedirectUrl(this.publicStoreService.loadedStore.uuid)
			.pipe(
				take(1),
				map((url) => url.replace(/"/g, '')),
			)
			.subscribe((url) => this.$redirectUrl.set(this.addHttps(url)));

		//Redirect after the countdown
		timer(1000, 1000)
			.pipe(
				finalize(() => {
					this.redirectToStore();
				}),
				takeWhile(() => this.secondsTimer > 0),
				takeUntilDestroyed(this.destroyRef),
				takeUntil(this.manualDestroy$),
			)
			.subscribe((_) => this.secondsTimer--);
	}

	redirectToStore(): void {
		window.location.href = this.$redirectUrl();
	}

	getImageInfo(): void {
		this.http
			.get<ImageInfo>(
				`${environment.host + environment.image.host + environment.image.path}${this.imageUuid}`,
			)
			.pipe(takeUntilDestroyed(this.destroyRef))
			.subscribe((res: ImageInfo) => {
				this.$logo.set(
					`${environment.host + environment.image.host + environment.image.path}${res.ID}/original${res.Extension}`,
				);
			});
	}

	addHttps(url) {
		if (url === '') {
			return '';
		}

		if (!url.startsWith('https://') && !url.startsWith('http://')) {
			return 'https://' + url;
		}
		return url;
	}

	manualClick() {
		//We only need to destroy the timer because the finalize operator will take care of the redirect
		this.manualDestroy$.next();
	}
}
