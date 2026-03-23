import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ToasterComponent } from '@app/shared/components/toaster/toaster.component';
import { NotificationComponent } from './modules/notification/notification.component';


@Component({
    selector: 'app-root',
    standalone: true,
    imports: [RouterOutlet, ToasterComponent, NotificationComponent],
    template: `
        <tagpeak-toaster/>
        <tagpeak-notifications />
        <router-outlet></router-outlet>
    `,
    styles: [
        `
            :host {
                display: flex;
                flex: 1 1 auto;
                width: 100%;
                height: 100%;
            }
        `
    ]
})
export class AppComponent {
    /**
     * Constructor
     */
    constructor() { }
}
