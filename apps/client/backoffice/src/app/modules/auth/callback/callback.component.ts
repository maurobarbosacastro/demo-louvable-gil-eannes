import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {AuthService} from '@app/core/auth/auth.service';
import {UserService} from '@app/core/user/user.service';
import {AppConstants} from '@app/app.constants';

@Component({
	selector: 'app-callback',
	template: `<h2>Processing login...</h2>`,
	standalone: true,
})
export class CallbackComponent implements OnInit {
	constructor(
		private authService: AuthService,
		private router: Router,
		private userService: UserService,
		private _route: ActivatedRoute,
	) {}

	ngOnInit() {
		this.authService.handleSocialsState();
		// First, check if the tokens already exist
		const accessToken = localStorage.getItem('access_token');
		const idToken = localStorage.getItem('id_token');

		// If tokens already exist, avoid reprocessing the authorization code
		if (accessToken && idToken) {
			console.log('Tokens already exist, skipping code exchange.');
			this.router.navigateByUrl('/');
			return; // Stop here if tokens are already available
		}

		// Otherwise, try to get the authorization code from the URL
		const code = this.authService.getAuthorizationCode();

		if (code) {
			// Exchange the authorization code for access and ID tokens
			this.authService.exchangeCodeForToken(code).subscribe(
				() => {
					this.userService.get().subscribe( _ =>{
						this.authService.loginChoice = AppConstants.LOGIN_CHOICE.SOCIALS
						this.router.navigateByUrl('/')
					});
				},
				(error) => {
					console.error('Error exchanging code for tokens', error);
				},
			);
		} else {
			console.error('Authorization code not found');
		}
	}
}
