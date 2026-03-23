import { Component, OnDestroy, OnInit, ViewEncapsulation } from '@angular/core';
import {Router, RouterModule} from '@angular/router';
import { finalize, Subject, takeUntil, takeWhile, tap, timer } from 'rxjs';
import { AuthService } from '@app/core/auth/auth.service';
import {MatButtonModule} from "@angular/material/button";
import {SharedModule} from "@app/shared/shared.module";
import {AppConstants} from '@app/app.constants';

@Component({
    selector: 'auth-sign-out',
    templateUrl: './sign-out.component.html',
    encapsulation: ViewEncapsulation.None,
    standalone: true,
    imports: [
        RouterModule,
        MatButtonModule,
        SharedModule,
    ],
})
export class AuthSignOutComponent implements OnInit, OnDestroy {
    countdown: number = 5;
    countdownMapping: any = {
        '=1': '# second',
        other: '# seconds',
    };
    private _unsubscribeAll: Subject<any> = new Subject<any>();

    /**
     * Constructor
     */
    constructor(private _authService: AuthService, private _router: Router) {}

    // -----------------------------------------------------------------------------------------------------
    // @ Lifecycle hooks
    // -----------------------------------------------------------------------------------------------------

    /**
     * On init
     */
    ngOnInit(): void {
        // Sign out
        this._authService.signOut();
				if (this._authService.loginChoice && this._authService.loginChoice === AppConstants.LOGIN_CHOICE.SOCIALS){
					this._authService.logoutSocials();
				}

        // Redirect after the countdown
        timer(1000, 1000)
            .pipe(
                finalize(() => {
                    this._router.navigate(['sign-in']);
                }),
                takeWhile(() => this.countdown > 0),
                takeUntil(this._unsubscribeAll),
                tap(() => this.countdown--)
            )
            .subscribe();
    }

    /**
     * On destroy
     */
    ngOnDestroy(): void {
        // Unsubscribe from all subscriptions
        this._unsubscribeAll.next(null);
        this._unsubscribeAll.complete();
    }
}
