import {Component, computed, inject, Signal, signal, ViewEncapsulation, WritableSignal} from '@angular/core';
import {fuseAnimations} from '@fuse/animations';
import {Router, RouterModule} from '@angular/router';
import {MatButtonModule} from '@angular/material/button';
import {SharedModule} from '@app/shared/shared.module';
import {MatIcon} from '@angular/material/icon';
import {TranslocoPipe} from '@ngneat/transloco';
import {NgOptimizedImage} from '@angular/common';
import {TagpeakCarouselComponent} from '@app/shared/components/tagpeak-carousel/tagpeak-carousel.component';
import {ShopService} from '@app/modules/shop/services/shop.service';
import {environment} from '@environments/environment';

@Component({
	selector: 'auth-confirmation-required',
	templateUrl: './confirmation-required.component.html',
	encapsulation: ViewEncapsulation.None,
	animations: fuseAnimations,
	standalone: true,
	imports: [
		RouterModule,
		MatButtonModule,
		SharedModule,
		MatIcon,
		TranslocoPipe,
		NgOptimizedImage,
		TagpeakCarouselComponent,
	],
})
export class AuthConfirmationRequiredComponent {
	_router: Router = inject(Router);
    _shopService: ShopService = inject(ShopService);

    $isShop: WritableSignal<boolean> = signal(!!this._shopService.shopName);
    $shopLink: Signal<string> = computed( () => {
        if (!this.$isShop()){
            return '';
        }
        const shop = this._shopService.shopName.split('.')[0];
        return `https://admin.shopify.com/store/${shop}/apps/${environment.shopify.appName}/app`
    });

	goBack() {
		this._router.navigate(['/stores']);
	}
}
