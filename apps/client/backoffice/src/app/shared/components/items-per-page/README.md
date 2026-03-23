# Atlanse Items-Per-Page @1.0.0

This is an Angular Component that provides a new feature to be used alongside the atl-table for displaying items-per-page [10,20,50,etc].

### How to implement:

#### Angular Code:

```
<atl-items-per-page
    class="justify-end"
    [totalItems]="totalItems"
    [minItemNumber]="minItemNumber"
    [maxItemNumber]="maxItemNumber"
    (onSelect)="onChangeItemsNumber($event)"
>
</atl-items-per-page>
```

#### Typescript:
```
    totalItems: string;
    minItemNumber: string;
    maxItemNumber: string;
```

```
 onChangeItemsNumber($event){
        this.pageSize = $event;
        this.getUsers(0, '');
    }
```

#### Example as below:

```
getUsers(page: number, name: string): void {
        combineLatest([
            this.textControl.valueChanges.pipe(
                startWith(name),
                debounceTime(500),
                distinctUntilChanged(),
                map((searchText) => searchText)
            ),
            this.typeControl.valueChanges.pipe(
                startWith(null),
                distinctUntilChanged(),
                map((code) => {
                    return code?.value?.toUpperCase()
                })
            )
        ]).pipe(
            switchMap(([text, status]) => {
                const filters = this.createUsersFilter(text, this.typeControl.value.value);
                return this.userService.getUsers(filters, page, this.pageSize, this.sort);
            }),
            map((data) => {
                this.page = data.pageNumber + 1;
                this.total = data.totalPages;
                
                this.totalItems = data.totalSize == null ? "0" : data.totalSize.toString();
                this.minItemNumber = (data.offset +1).toString() ;
                this.maxItemNumber = (this.pageSize * data.pageNumber + data.numberOfElements).toString();
                
                let tempList: User[] = [];
                try {
                    tempList = data.content.map(user => ({
                        ...user,
                        open: false
                    }));
                } catch (e) {
                    this.page = 0;
                }
                return this.sortList(tempList);
            }),
            catchError(err => of([]))
        )
            .subscribe(res => {
                this.alphabeticalUserList = res;
                this.selectUser(this.alphabeticalUserList[0]?.letter, this.alphabeticalUserList[0]?.list[0]);
                this.loading = false;
            });
    }
```
