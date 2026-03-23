import { Component, ViewEncapsulation } from '@angular/core';
import {RouterModule} from "@angular/router";
import {MatButtonModule} from "@angular/material/button";
import {MatIconModule} from "@angular/material/icon";
import {SharedModule} from "@app/shared/shared.module";

@Component({
    selector: 'landing-home',
    templateUrl: './home.component.html',
    encapsulation: ViewEncapsulation.None,
    standalone: true,
    imports: [
        RouterModule,
        MatButtonModule,
        MatIconModule,
        SharedModule,
    ],
})
export class LandingHomeComponent {
    /**
     * Constructor
     */
    constructor() {}
}
