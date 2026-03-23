import { Directive, ElementRef, OnInit } from '@angular/core';

@Directive({
    selector: '[tagpeakInputStyle]',
    standalone: true
})
export class TagpeakInputDirective implements OnInit {
    constructor(private el: ElementRef) {}

    ngOnInit() {
        const nativeElement = this.el.nativeElement;

        // Add the specified Tailwind CSS classes
        nativeElement.classList.add(
            'w-full',
            'pt-1',
            'placeholder:text-project-waterloo',
            'focus:outline-none',
            'text-project-licorice-500',
            'border-b',
            'border-project-licorice-900',
            'pb-2',
            'font-medium'
        );
    }
}
