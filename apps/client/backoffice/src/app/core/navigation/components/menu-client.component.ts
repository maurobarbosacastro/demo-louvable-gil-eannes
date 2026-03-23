import {
	ChangeDetectionStrategy,
	Component,
	computed,
	DestroyRef,
	effect,
	inject,
	input,
	InputSignal,
	OnInit,
	signal,
	WritableSignal,
} from '@angular/core';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {NgClass} from '@angular/common';
import {Router, NavigationEnd} from '@angular/router';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe} from '@ngneat/transloco';
import {filter} from 'rxjs/operators';

import {FuseNavigationService} from '@fuse/components/navigation';
import {fuseAnimations} from '@fuse/animations';

import {MENU_ITEMS} from './menu-client.config';
import {MenuItem} from './menu-client.types';

@Component({
	selector: 'tagpeak-menu-client',
	templateUrl: './menu-client.component.html',
	standalone: true,
	imports: [TranslocoPipe, MatIcon, NgClass],
	animations: fuseAnimations,
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MenuClientComponent implements OnInit {
	// Injected dependencies
	private readonly _router = inject(Router);
	private readonly _fuseNavigationService = inject(FuseNavigationService);
	private readonly _destroyRef = inject(DestroyRef);

	// Inputs using new signal-based API
	readonly opened: InputSignal<boolean> = input<boolean>(false);
	readonly name: InputSignal<string> = input.required<string>();

	// State signals
	protected readonly $isOpen: WritableSignal<boolean> = signal<boolean>(false);
	protected readonly $currentMenu: WritableSignal<MenuItem | null> = signal<MenuItem | null>(null);
	protected readonly $animationsEnabled: WritableSignal<boolean> = signal<boolean>(false);

	// Computed signals
	protected readonly $isMenuOpen = computed(() => this.$isOpen());

	// Public properties
	protected readonly menuList: MenuItem[] = MENU_ITEMS;

	constructor() {
		// Sync input with internal state
		effect(() => {
			this.$isOpen.set(this.opened());
		});

		// Track route changes to update current menu
		this._setupRouteTracking();
	}

	ngOnInit(): void {
		// Set initial current menu
		this._updateCurrentMenu(this._router.url);

		// Register component with Fuse navigation service
		this._fuseNavigationService.registerComponent(this.name(), this);
	}

	// ============================================
	// Public Methods
	// ============================================

	toggle(): void {
		this.$isOpen() ? this.close() : this.open();
	}

	open(): void {
		if (this.$isOpen()) return;
		this._setOpened(true);
	}

	close(): void {
		if (!this.$isOpen()) return;
		this._setOpened(false);
	}

	goToRoute(menu: MenuItem): void {
		this._router.navigate(menu.router);
	}

	// ============================================
	// Private Methods
	// ============================================

	private _setOpened(isOpen: boolean): void {
		this.$isOpen.set(isOpen);
		this._enableAnimations();
	}

	private _enableAnimations(): void {
		if (this.$animationsEnabled()) return;
		this.$animationsEnabled.set(true);
	}

	private _setupRouteTracking(): void {
		this._router.events
			.pipe(
				filter((event): event is NavigationEnd => event instanceof NavigationEnd),
				takeUntilDestroyed(this._destroyRef),
			)
			.subscribe((event) => {
				this._updateCurrentMenu(event.urlAfterRedirects);
			});
	}

	private _updateCurrentMenu(url: string): void {
		const currentMenu: MenuItem = this.menuList.find((menu: MenuItem): boolean =>
			url.includes(menu.value),
		);
		this.$currentMenu.set(currentMenu ?? null);
	}
}
