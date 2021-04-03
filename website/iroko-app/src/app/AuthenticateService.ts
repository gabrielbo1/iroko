import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from "@angular/router";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";

@Injectable()
export class AuthenticateService implements CanActivate {
    constructor(private router: Router) {
    }

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
        if (localStorage.getItem('login') !== 'ok') {
            this.router.navigate(['login']);
            return false;
        }
        return true;
    }
}
