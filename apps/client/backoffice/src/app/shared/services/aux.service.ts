import {Injectable} from '@angular/core';

@Injectable({
	providedIn: 'root',
})
export class AuxService {
	loadedToastr: boolean = false;

	constructor() {}

	setAuxToastrVar(value: boolean) {
		this.loadedToastr = value;
	}

	getAuxToastrVar(): boolean {
		return this.loadedToastr;
	}
}
