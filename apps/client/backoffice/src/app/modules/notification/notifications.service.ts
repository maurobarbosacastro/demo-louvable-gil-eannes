import {HttpClient} from '@angular/common/http';
import {Injectable, Signal, WritableSignal, inject, signal} from '@angular/core';
import {Messaging, getToken, onMessage} from '@angular/fire/messaging';
import {createApplication} from '@angular/platform-browser';
import {UserService} from '@app/core/user/user.service';
import {user} from '@app/mock-api/common/user/data';
import {environment} from '@environments/environment';
import {
	catchError,
	EMPTY,
	from,
	map,
	of,
	Subject,
	switchMap,
	BehaviorSubject,
	timer,
	takeUntil,
	fromEvent,
	startWith,
} from 'rxjs';
import {toSignal} from '@angular/core/rxjs-interop';
import {NotificationMessage} from './notification.interface';

@Injectable({
	providedIn: 'root',
})
export class NotificationService {
	private readonly messaging = inject(Messaging, {optional: true});
	private readonly http = inject(HttpClient);
	private readonly userService = inject(UserService);

	private readonly url = environment.host + environment.tagpeak + '/notifications';

	$notificationPermitted: WritableSignal<boolean> = signal(false);
	notificationList$: Subject<NotificationMessage> = new Subject();

	private readonly notificationStack$ = new BehaviorSubject<NotificationMessage[]>([]);
	private readonly defaultDuration = 500000;
	private readonly maxNotifications = 5;

	readonly notificationStack: Signal<NotificationMessage[]> = toSignal(this.notificationStack$, {
		initialValue: [],
	});
	readonly notificationCount: Signal<number> = toSignal(
		this.notificationStack$.pipe(map((stack) => stack.length)),
		{initialValue: 0},
	);

	requestPermission() {
		console.log(this.messaging);
		if (!environment.features.notifications) {
			console.warn('Notification feature is off.');
			return;
		}
		console.log('Requesting permission...');
		from(Notification.requestPermission())
			.pipe(
				switchMap((permission) => {
					if (permission === 'granted') {
						console.log('Notification permission granted.');
						return from(getToken(this.messaging, {vapidKey: environment.firebaseConfig.vapidKey}));
					}

					if (permission === 'denied') {
						console.log('Notification permission denied');
						// return this.clearUserToken()
					}
					return EMPTY;
				}),
				switchMap((token) => {
					if (token) {
						return this.saveToken(token as string);
					}

					console.warn('No registration token available. Request permission to generate one.');
					return EMPTY;
				}),
			)
			.subscribe((token) => {
				this.$notificationPermitted.set(token);
				if (token) {
					this.onMessage();
					return;
				}
				console.warn('No registration token available. Request permission to generate one.');
			});
	}

	private saveToken(token: string) {
		if (!environment.features.notifications) {
			console.warn('Notification feature is off.');
			return;
		}

		return this.http
			.post(`${this.url}/token`, {
				userUuid: this.userService.$user().uuid,
				token,
			})
			.pipe(
				catchError((_) => {
					return of(false);
				}),
				map((_) => true),
			);
	}

	clearUserToken() {
		if (!environment.features.notifications) {
			console.warn('Notification feature is off.');
			return;
		}

		return this.http.delete(`${this.url}/token`).pipe(map((_) => null));
	}

	private onMessage() {
		onMessage(this.messaging, (payload) => {
			console.info({
				payload,
				visibility: document.visibilityState,
			});
			if (payload.notification) {
				if (document.visibilityState === 'hidden') {
					new Notification(payload.notification.title, {body: payload.notification.body});
				}
				const notification: NotificationMessage = {
					id: crypto.randomUUID(),
					title: payload.notification.title,
					body: payload.notification.body,
					duration: this.defaultDuration,
					dismissible: true,
				};
				this.notificationList$.next(notification);
				this.addNotification(notification);
			}
		});
	}

	private addNotification(notification: NotificationMessage): void {
		const currentStack = this.notificationStack$.value;
		let newStack = [...currentStack, notification];

		if (newStack.length > this.maxNotifications) {
			newStack = newStack.slice(-this.maxNotifications);
		}

		this.notificationStack$.next(newStack);

		const duration = notification.duration ?? this.defaultDuration;
		if (duration > 0) {
			timer(duration).subscribe(() => {
				this.removeNotification(notification.id);
			});
		}
	}

	removeNotification(id: string): void {
		const currentStack = this.notificationStack$.value;
		const newStack = currentStack.filter((notification) => notification.id !== id);
		this.notificationStack$.next(newStack);
	}

	clearNotifications(): void {
		this.notificationStack$.next([]);
	}
}
