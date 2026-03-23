import {Component, ElementRef, inject, OnInit, ViewChild} from '@angular/core';
import {CommonModule} from '@angular/common';
import {Toast, ToasterService} from '@app/shared/components/toaster/toaster.service';
import {TranslocoService} from '@ngneat/transloco';
import {MatIcon} from '@angular/material/icon';

@Component({
	selector: 'tagpeak-toaster',
	standalone: true,
	imports: [CommonModule, MatIcon],
	templateUrl: './toaster.component.html',
})
export class ToasterComponent implements OnInit {
	toasterService = inject(ToasterService);
	transloco = inject(TranslocoService);
	toasts: Toast[] = [];
	toastAction: string;
	actionIcon: string;

	@ViewChild('toasterWrapper') toasterWrapper: ElementRef;

	ngOnInit() {
		this.toasterService.toasts$.subscribe((toasts) => {
			this.toasts = toasts;
		});
	}

	closeToast(id: number) {
		this.toasterService.removeToast(id);
	}

	getToastClass(type: string) {
		switch (type) {
			case 'success':
				this.actionIcon = 'success';
				this.toastAction = this.transloco.translate('toaster.success');
				return 'bg-[#d9d7f8] text-project-youngBlue';
			case 'error':
				this.actionIcon = 'error';
				this.toastAction = this.transloco.translate('toaster.error');
				return 'text-project-red bg-[#fcd4de]';
			case 'info':
				this.actionIcon = 'info';
				this.toastAction = this.transloco.translate('toaster.info');
				return 'text-project-waterloo bg-[#e4e5e9]';
			case 'warning':
				this.actionIcon = 'warning';
				this.toastAction = this.transloco.translate('toaster.info');
				return 'text-project-waterloo bg-[#FDCA65]';
		}
	}
}
