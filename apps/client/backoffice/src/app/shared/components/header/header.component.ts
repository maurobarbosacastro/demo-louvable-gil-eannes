import { Component, EventEmitter, inject, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatIconModule } from '@angular/material/icon';
import { RouterLink } from '@angular/router';
import { UntilDestroy } from '@ngneat/until-destroy';
import { TranslocoModule } from '@ngneat/transloco';
import { HeaderService } from '@app/shared/components/header/header.service';

@UntilDestroy()
@Component({
    selector: 'app-header',
    standalone: true,
    imports: [CommonModule, MatIconModule, TranslocoModule, RouterLink],
    template: `
        <div class="relative hidden md:flex flex-auto justify-center p-1 lg:px-15 overflow-hidden dark:border-l bg-feamoney-blueGray">
            <ng-container *ngIf="headerService.getHeaderInfo$() | async as headerInfo">
                <div [ngClass]="{'z-10 relative w-full': true,
                    'mt-14 mb-[38px]': headerInfo.subTitle !== undefined,
                    ' my-14': headerInfo.subTitle === undefined}">
                    <div class="flex justify-between">
                        <div class="text-4xl font-black font-lato text-feamoney-darkBlue">{{ headerInfo.title | transloco | uppercase }}</div>
                        <button *ngIf="headerInfo.button" (click)="triggerAction()"
                                class="bg-feamoney-lightPurple  items-center flex justify-evenly text-lg font-black rounded-3xl font-lato dark w-36 h-10">
                            <mat-icon *ngIf="headerInfo.buttonIcon !== undefined"
                                      class="bg-feamoney-purple dark h-5 W-5 rounded-3xl "
                                      [svgIcon]="headerInfo.buttonIcon"
                            >
                            </mat-icon>
                            {{ headerInfo.buttonText | transloco }}
                        </button>
                    </div>
                    <div *ngIf="headerInfo.subTitle!== undefined" class="text-feamoney-darkGrey font-lato text-sm">
                        <span>{{ headerInfo.subTitle | transloco }}</span>
                    </div>
                </div>
            </ng-container>

        </div>

    `
})
export class HeaderComponent{
    @Input() title: string;
    @Input() subtitle: string;
    @Input() button: boolean;
    @Input() buttonText: string;
    @Input() buttonIcon: string;
    @Output() clickButtonEvent = new EventEmitter();

    route: string;

    headerService = inject(HeaderService);

    protected triggerAction(): void {
        this.headerService.actionListener().next();
    }

}
