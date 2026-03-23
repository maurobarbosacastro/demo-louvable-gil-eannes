import {Component, OnInit} from '@angular/core';
import {UserService} from './services/user.service';
import {BehaviorSubject, catchError, map, of, tap} from 'rxjs';
import countriesList from '../../../assets/files/country_json.json';
import {SaveModalComponent} from '@app/shared/components/save-modal/save-modal.component';
import {MatDialog} from '@angular/material/dialog';
import {TranslocoModule, TranslocoService} from '@ngneat/transloco';
import {ActivatedRoute, Route, RouterModule} from '@angular/router';
import {User} from '@app/modules/users/interfaces/user.interface';
import {Image} from '@app/shared/interfaces/image.interface';
import {differenceInYears} from 'date-fns';
import {HeaderService} from '@app/layout/layouts/classy/components/header/header.service';
import {UntilDestroy} from '@ngneat/until-destroy';
import {CommonModule} from "@angular/common";
import {PaginationComponent} from "@app/shared/components/pagination/pagination.component";
import {MatIconModule} from "@angular/material/icon";
import {MatButtonModule} from "@angular/material/button";
import {MatMenuModule} from "@angular/material/menu";

interface AlphabeticalUserList {
    letter: string;
    list: User[];
}


@UntilDestroy()
@Component({
    selector: 'app-users',
    templateUrl: './users.component.html',
    styleUrls: ['./users.component.scss'],
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        PaginationComponent,
        TranslocoModule,
        MatIconModule,
        MatButtonModule,
        MatMenuModule,
    ],
})
export class UsersComponent implements OnInit {

    selectedUser: User;
    alphabeticalUserList: AlphabeticalUserList[];
    page: number;
    total: number;
    blockOpen: boolean = false;
    countries: any[] = countriesList;
    email: string = '';

    loading: boolean = true;

    pageChange$: BehaviorSubject<number>;

    constructor(
        private userService: UserService,
        public dialog: MatDialog,
        private activeRouter: ActivatedRoute,
        private translocoService: TranslocoService,
        private headerService: HeaderService) {
    }

    ngOnInit(): void {
        this.headerService.setHeaderInfo({title: 'users.users', button: true, buttonText: 'users.new'});
        this.headerService.actionListener()
            .subscribe(_ => {
                console.log('click');
            });
        this.pageChange$ = new BehaviorSubject<number>(0);

        this.activeRouter.params.subscribe((data) => {
            if (data.email) {
                this.email = data.email;
                this.getUsers(0, data.email);
            } else {
                this.getUsers(0, '');
            }
        });


    }

    getUsers(page: number, email: string): void {
        this.userService.getUsers(page, email)
            .pipe(
                map((data) => {
                    this.page = data.pageNumber;
                    this.total = data.totalPages;
                    let tempList: User[] = [];
                    try {
                        tempList = data.content.map(user => ({
                            ...user,
                            country: this.getCountryName(user) ?? '',
                            open: false,
                            age: differenceInYears(new Date(), new Date(user.birthDate))
                        }));
                    } catch (e) {
                        console.log({e});
                    }
                    return this.sortList(tempList);
                }),
                catchError(err => of([]))
            )
            .subscribe((res) => {
                this.alphabeticalUserList = res;
                this.selectUser(this.alphabeticalUserList[0].letter, this.alphabeticalUserList[0].list[0]);
                this.loading = false;
            });
    }

    getCountryName(user: User): string {
        return this.countries.find(elem => elem['Code'].toLowerCase() === user.country.toLowerCase())['Name'];
    }

    sortList(list: any[]): AlphabeticalUserList[] {
        const finalList: AlphabeticalUserList[] = [...Array(26)].map((_, i) => String.fromCharCode(i + 65)).map(letter => ({letter, list: []}));
        list.forEach((elem) => {
            const idx = finalList.findIndex(el => el.letter === elem.firstName.charAt(0).toUpperCase());
            finalList[idx].list.push(elem);
        });
        return finalList.filter(el => el.list.length > 0);
    }

    selectUser(letter: string, user: User): void {
        this.selectedUser = user;
        const userList = this.alphabeticalUserList.find(elem => letter.toUpperCase() === elem.letter.toUpperCase());

        this.alphabeticalUserList.forEach((item) => {
            item.list.forEach((elem) => {
                elem.open = false;
            });
        });

        const openUser = userList.list.find(item => item === user);
        openUser.open = true;
    }

    openBlocked(): void {
        this.blockOpen = true;
    }

    blockUser(user: User): void {
        let data: any = {
            positiveButton: this.translocoService.translate('misc.yes'),
            negativeButton: this.translocoService.translate('misc.no'),
            image: 'assets/images/gifs/blockUser.gif'
        };
        if (this.selectedUser.isBlocked) {
            data = {
                ...data,
                text: this.translocoService.translate('users.unblock-modal')
            };
        } else {
            data = {
                ...data,
                text: this.translocoService.translate('users.block-modal')
            };
        }
        const dialogRef = this.dialog.open(SaveModalComponent, {data: data});
        dialogRef.afterClosed().subscribe((result) => {
            if (result) {
                this.userService.blockUser(user.email).subscribe(() => {

                    this.getUsers(this.page - 1, '');
                });
            }
        });
    }

    nextPage(): void {
        this.getUsers(this.page, '');
    }

    previousPage(): void {
        this.getUsers(this.page - 2, '');
    }

    getUserImage(user: User): Image | string {
        if (!user) {
            return '';
        }
        return user.profilePicture.extension ? 'data:image/' + user.profilePicture.extension + '+xml;base64,' + user.profilePicture.base64 : user.profilePicture;
    }
}
