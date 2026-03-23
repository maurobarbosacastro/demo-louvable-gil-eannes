import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core'
import { NotificationService } from './notifications.service';
import { NotificationMessage } from './notification.interface';


@Component({
    selector: 'tagpeak-notifications',
    standalone: true,
    imports: [
        CommonModule,
    ],
    template: `
        <div class="fixed top-8 right-8 z-50 space-y-2 max-w-sm">
            @for (notification of notificationStack(); track notification.id) {
                <div class="bg-[#EEEDF8] border border-[#E2E1F5] shadow-lg rounded-lg p-4 relative cursor-pointer transition-all duration-300 hover:shadow-xl"
                     (click)="onNotificationClick(notification)">

                    @if (notification.dismissible) {
                        <button class="absolute top-2 right-2 text-gray-400 hover:text-gray-600 transition-colors duration-200"
                                (click)="onDismissClick($event, notification.id)"
                                aria-label="Dismiss notification">
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                            </svg>
                        </button>
                    }

                    <div class="pr-6">
                        <h6 class="font-semibold text-gray-900 text-sm mb-1">{{ notification.title }}</h6>
                        <p class="text-sm text-gray-600">{{ notification.body }}</p>
                    </div>
                </div>
            }
        </div>
    `
})
export class NotificationComponent {

    private readonly notifService = inject(NotificationService);

    readonly notificationStack = this.notifService.notificationStack;

    onNotificationClick(notification: NotificationMessage): void {
        if (notification.dismissible) {
            this.notifService.removeNotification(notification.id);
        }
    }

    onDismissClick(event: Event, notificationId: string): void {
        event.stopPropagation();
        this.notifService.removeNotification(notificationId);
    }

}
