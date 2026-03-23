import { Directive, ElementRef, OnInit } from '@angular/core';

@Directive({
    selector: '[tagpeakReadonlyInput]',
    standalone: true
})
export class ReadonlyInputDirective implements OnInit {
    constructor(private el: ElementRef) {}

    ngOnInit() {
        const nativeElement = this.el.nativeElement;

        // Set the input as readonly
        nativeElement.readOnly = true;

        // Add pointer-events-none class
        nativeElement.classList.add('pointer-events-none','text-project-waterloo', '!font-normal', 'opacity-50');
    }
}
