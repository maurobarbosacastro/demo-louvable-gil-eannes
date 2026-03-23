import {Injectable} from '@angular/core';
import {BehaviorSubject, Subject} from 'rxjs';

interface HeaderInfo {
    title: string;
    subTitle?: string;
    button: boolean;
    buttonIcon?: string;
    buttonText?: string;
}

@Injectable({
    providedIn: 'root'
})
export class HeaderService {

    private readonly headerInfo$: BehaviorSubject<HeaderInfo>;

    private readonly actionTriggered$: Subject<void>;

    constructor() {
        this.headerInfo$ = new BehaviorSubject<HeaderInfo>({} as HeaderInfo);
        this.actionTriggered$ = new Subject<void>();
    }

    getHeaderInfo$(): BehaviorSubject<HeaderInfo> {
        return this.headerInfo$;
    }

    setHeaderInfo(info: HeaderInfo): void {
        this.headerInfo$.next(info);
    }

    actionListener(): Subject<void> {
        return this.actionTriggered$;
    }


}
