import {inject, Injectable} from '@angular/core';
import {MatIconRegistry} from '@angular/material/icon';
import {DomSanitizer} from '@angular/platform-browser';

@Injectable({providedIn: 'root'})
export class IconsService {
	/**
	 * Constructor
	 */
	constructor() {
		const domSanitizer = inject(DomSanitizer);
		const matIconRegistry = inject(MatIconRegistry);

		// Register icon sets
		/*matIconRegistry.addSvgIconSet(
            domSanitizer.bypassSecurityTrustResourceUrl(
                'icons/material-twotone.svg'
            )
        );
        matIconRegistry.addSvgIconSetInNamespace(
            'mat_outline',
            domSanitizer.bypassSecurityTrustResourceUrl(
                'icons/material-outline.svg'
            )
        );
        matIconRegistry.addSvgIconSetInNamespace(
            'mat_solid',
            domSanitizer.bypassSecurityTrustResourceUrl(
                'icons/material-solid.svg'
            )
        );
        matIconRegistry.addSvgIconSetInNamespace(
            'feather',
            domSanitizer.bypassSecurityTrustResourceUrl('icons/feather.svg')
        );
        matIconRegistry.addSvgIconSetInNamespace(
            'heroicons_outline',
            domSanitizer.bypassSecurityTrustResourceUrl(
                'icons/heroicons-outline.svg'
            )
        );
        matIconRegistry.addSvgIconSetInNamespace(
            'heroicons_solid',
            domSanitizer.bypassSecurityTrustResourceUrl(
                'icons/heroicons-solid.svg'
            )
        );
        matIconRegistry.addSvgIconSetInNamespace(
            'heroicons_mini',
            domSanitizer.bypassSecurityTrustResourceUrl(
                'icons/heroicons-mini.svg'
            )
        );*/

		const icons: {name: string; src: string}[] = [
			{name: 'mat_outline', src: 'assets/icons/material-outline.svg'},
			{name: 'mat_solid', src: 'assets/icons/material-solid.svg'},
			{name: 'feather', src: 'assets/icons/feather.svg'},
			{name: 'heroicons_outline', src: 'assets/icons/heroicons-outline.svg'},
			{name: 'heroicons_solid', src: 'assets/icons/heroicons-solid.svg'},
			{name: 'ellipsis-vertical', src: 'assets/icons/ellipsis-vertical.svg'},
			{name: 'arrow-left', src: 'assets/icons/arrow-left.svg'},
			{name: 'arrow-right', src: 'assets/icons/arrow-right.svg'},
			{name: 'external-link', src: 'assets/icons/arrow_external.svg'},
			{name: 'chart', src: 'assets/icons/chart.svg'},
			{name: 'visibility', src: 'assets/icons/visibility.svg'},
			{name: 'visibility_off', src: 'assets/icons/visibility_off.svg'},
			{name: 'error', src: 'assets/icons/error.svg'},
			{name: 'facebook', src: 'assets/icons/facebook.svg'},
			{name: 'google', src: 'assets/icons/google.svg'},
			{name: 'checked', src: 'assets/icons/checked.svg'},
			{name: 'USD', src: 'assets/icons/usa_flag.svg'},
			{name: 'EUR', src: 'assets/icons/euro_flag.svg'},
			{name: 'CAD', src: 'assets/icons/canada_flag.svg'},
			{name: 'AUD', src: 'assets/icons/australia_flag.svg'},
			{name: 'GBP', src: 'assets/icons/great_britain_flag.svg'},
			{name: 'BRL', src: 'assets/icons/brazil-flag.svg'},
			{name: 'MXN', src: 'assets/icons/mexico-flag.svg'},
			{name: 'envelope', src: 'assets/icons/envelope-regular.svg'},
			{name: 'pen', src: 'assets/icons/pen-solid.svg'},
			{name: 'tag-delete', src: 'assets/icons/tag-delete.svg'},
			{name: 'arrow-left-solid', src: 'assets/icons/arrow-left-solid.svg'},
			{name: 'duplicate', src: 'assets/icons/duplicate.svg'},
			{name: 'arrow_drop_down', src: 'assets/icons/arrow_drop_down.svg'},
			{name: 'chevron_backward', src: 'assets/icons/chevron_backward.svg'},
			{name: 'success', src: 'assets/icons/check_circle.svg'},
			{name: 'info', src: 'assets/icons/credit_score.svg'},
			{name: 'calender', src: 'assets/icons/calender.svg'},
			{name: 'close', src: 'assets/icons/close.svg'},
			{name: 'samsung', src: 'assets/icons/samsung.svg'},
			{name: 'table_chart_view', src: 'assets/icons/table-chart-view.svg'},
			{name: 'table_stop', src: 'assets/icons/table-stop.svg'},
			{name: 'stars', src: 'assets/icons/stars.svg'},
			{name: 'credit-score', src: 'assets/icons/credit-score.svg'},
			{name: 'no-rewards', src: 'assets/images/rewards/no-rewards.svg'},
			{name: 'close-modal', src: 'assets/icons/close-modal.svg'},
			{name: 'clock', src: 'assets/icons/clock.svg'},
			{name: 'sell', src: 'assets/icons/sell.svg'},
			{name: 'stop-reward', src: 'assets/images/rewards/stop-reward.svg'},
			{name: 'check', src: 'assets/icons/check.svg'},
			{name: 'budget', src: 'assets/icons/budget.svg'},
			{name: 'wallet', src: 'assets/icons/wallet.svg'},
			{name: 'cashback', src: 'assets/icons/cashback.svg'},
			{name: 'search', src: 'assets/icons/search.svg'},
			{name: 'pencil', src: 'assets/icons/pencil.svg'},
			{name: 'trash', src: 'assets/icons/trash.svg'},
			{name: 'open-in', src: 'assets/icons/open-in.svg'},
			{name: 'heart-handshake', src: 'assets/icons/heard-handshake.svg'},
			{name: 'globe', src: 'assets/icons/globe.svg'},
			{name: 'category', src: 'assets/icons/category.svg'},
			{name: 'arrow-down', src: 'assets/icons/arrow-down.svg'},
			{name: 'info-warning', src: 'assets/icons/info-warning.svg'},
			{name: 'menu', src: 'assets/icons/menu.svg'},
			{name: 'menu-black', src: 'assets/icons/menu-black.svg'},
			{name: 'tagpeak-logo', src: 'assets/images/logo/logo.svg'},
			{name: 'hyperlink-arrow', src: 'assets/icons/hyperlink-arrow.svg'},
			{name: 'linked-in', src: 'assets/icons/linked-in.svg'},
			{name: 'twitter', src: 'assets/icons/twitter.svg'},
			{name: 'instagram', src: 'assets/icons/instagram.svg'},
			{name: 'for-shoppers', src: 'assets/images/stores/for-shoppers.svg'},
			{name: 'wrong', src: 'assets/icons/x.svg'},
			{name: 'add-group', src: 'assets/icons/ad_group_off.svg'},
			{name: 'chip', src: 'assets/icons/chip_extraction.svg'},
			{name: 'cookie', src: 'assets/icons/cookie.svg'},
			{name: 'shopping', src: 'assets/icons/shopping_cart.svg'},
			{name: 'eye', src: 'assets/icons/eye-solid.svg'},
			{name: 'two_persons', src: 'assets/icons/two_persons.svg'},
			{name: 'trending_up', src: 'assets/icons/trending_up.svg'},
			{name: 'budget-withdrawal', src: 'assets/icons/budget-withdrawal.svg'},
			{name: 'upload', src: 'assets/icons/upload.svg'},
			{name: 'uploaded', src: 'assets/icons/uploaded.svg'},
			{name: 'arrow-next-right', src: 'assets/icons/arrow-next-right.svg'},
			{name: 'withdrawal', src: 'assets/icons/withdrawal.svg'},
			{name: 'settings', src: 'assets/icons/settings.svg'},
			{name: 'users', src: 'assets/icons/users.svg'},
			{name: 'warning-info', src: 'assets/icons/warning-info.svg'},
			{name: 'improvement', src: 'assets/icons/improvement.svg'},
			{name: 'downgrade', src: 'assets/icons/downgrade.svg'},
			{name: 'equal_compare', src: 'assets/icons/equal.svg'},
			{name: 'alarm', src: 'assets/icons/alarm.svg'},
			{name: 'alert-triangle', src: 'assets/icons/alert-triangle.svg'},
		];

		matIconRegistry.addSvgIconSet(
			domSanitizer.bypassSecurityTrustResourceUrl('assets/icons/material-twotone.svg'),
		);
		icons.forEach((icon) =>
			matIconRegistry.addSvgIcon(icon.name, domSanitizer.bypassSecurityTrustResourceUrl(icon.src)),
		);
		icons.forEach((icon) =>
			matIconRegistry.addSvgIconSetInNamespace(
				icon.name,
				domSanitizer.bypassSecurityTrustResourceUrl(icon.src),
			),
		);
	}
}
