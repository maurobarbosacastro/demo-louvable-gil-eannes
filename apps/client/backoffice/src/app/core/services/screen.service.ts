
import { Injectable, NgZone } from '@angular/core';
import { BehaviorSubject, fromEvent, Observable } from 'rxjs';
import { debounceTime, map, startWith } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ScreenService {
  private isMobile = new BehaviorSubject<boolean>(window.innerWidth < 1280);
  public isMobile$ = this.isMobile.asObservable();

  constructor(private ngZone: NgZone) {
    fromEvent(window, 'resize')
      .pipe(
        debounceTime(100),
        map(() => window.innerWidth < 1280),
        startWith(window.innerWidth < 1280)
      )
      .subscribe(isMobile => {
        this.ngZone.run(() => {
          this.isMobile.next(isMobile);
        });
      });
  }
}
