import {Injectable} from '@angular/core';
import {BehaviorSubject} from 'rxjs';

export interface Toast {
	id: number;
	type: 'success' | 'error' | 'info' | 'warning';
	message: string;
	action?: {
		icon?: string;
		message?: string;
	};
	position: 'top' | 'bottom';
}

@Injectable({
	providedIn: 'root',
})
export class ToasterService {
	private toastsSubject = new BehaviorSubject<Toast[]>([]);
	toasts$ = this.toastsSubject.asObservable();
	private toastCounter = 0;

	showToast(
		type: 'success' | 'error' | 'info' | 'warning',
		message: string,
		position: 'top' | 'bottom' = 'bottom',
		action?: {
			icon?: string;
			message?: string;
		},
		time: number = 3000,
	) {
		const toast: Toast = {id: this.toastCounter++, type, message, position, action};
		this.toastsSubject.next([...this.toastsSubject.value, toast]);
		setTimeout(() => this.removeToast(toast.id), time); // Auto-remove after 3 seconds
	}

	removeToast(id: number) {
		this.toastsSubject.next(this.toastsSubject.value.filter((toast) => toast.id !== id));
	}
}
