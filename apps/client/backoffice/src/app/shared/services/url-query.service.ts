import {DestroyRef, Injectable} from '@angular/core';
import {delay, Subject} from 'rxjs';
import {ActivatedRoute, ParamMap, QueryParamsHandling, Router} from '@angular/router';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';

type QueryItemRequest = {
    name: string,
    value: string,
    paramHandling?: QueryParamsHandling,
    replaceUrl?: boolean
}

@Injectable({
  providedIn: 'root'
})
export class UrlQueryService {
    private paramsQueue: Subject<QueryItemRequest> = new Subject()
    currentParams: ParamMap = null;

    constructor(
        private destroyRef: DestroyRef,
        private router: Router,
        private route: ActivatedRoute
    ) {
        this.paramsQueue.asObservable()
            .pipe(
                delay(100),
                takeUntilDestroyed(this.destroyRef)
            )
            .subscribe( paramRequest => {
                this.router.navigate([], {
                    relativeTo: this.route,
                    queryParams: {[paramRequest.name]: paramRequest.value},
                    queryParamsHandling: paramRequest.paramHandling,
                    replaceUrl: paramRequest.replaceUrl
                })
            });
        this.route.queryParamMap.pipe(takeUntilDestroyed(this.destroyRef)).subscribe( paramsMap => this.currentParams = paramsMap)
    }

    addParam(
        name: string,
        value: string,
        replaceUrl: boolean = false,
        paramHandling: QueryParamsHandling = "merge"): void{
        this.paramsQueue.next({name, value, paramHandling, replaceUrl})
    }

    resetParams(): void{
        this.router.navigate([], {
            queryParams: {},
            queryParamsHandling: "merge"
        })
    }

    getParam(name: string): string{
        return this.currentParams.get(name)
    }

    get params$(){
        return this.route.queryParamMap;
    }

}
