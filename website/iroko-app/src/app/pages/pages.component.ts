import {Component, Inject} from '@angular/core';
import {Router} from "@angular/router";
import { DOCUMENT } from "@angular/common";

@Component({
  selector: 'pages',
  styleUrls: ['./pages.component.scss'],
  template: `

<mat-drawer-container [class.drawer-opened]="drawer.opened">
  <mat-drawer #drawer class="drawer" mode="side" opened="true">
    <mat-toolbar>
       <mat-toolbar-row  class="header">
            <div class="menu-app-image"></div>
              <div class="font-header-menu">Menu</div>
            <mat-icon class="close" (click)="drawer.close()">close</mat-icon>
        </mat-toolbar-row>

        <mat-toolbar-row  class="submenu">
            <div class="font-header-menu">Dashboard</div>
            <mat-icon class="close" (click)="drawer.close()">dashboard</mat-icon>
       </mat-toolbar-row>

        <mat-toolbar-row  class="submenu">
            <div class="font-header-menu">Chamados</div>
            <mat-icon class="close" (click)="drawer.close()">assignment</mat-icon>
       </mat-toolbar-row>

        <mat-toolbar-row  class="submenu">
            <div class="font-header-menu">Companias</div>
            <mat-icon class="close" (click)="drawer.close()">account_balance</mat-icon>
       </mat-toolbar-row>

       <mat-toolbar-row  class="submenu">
            <div class="font-header-menu">Usúarios</div>
            <mat-icon class="close" (click)="drawer.close()">person</mat-icon>
       </mat-toolbar-row>

       <mat-toolbar-row  class="submenu">
            <div class="font-header-menu">Configuração</div>
            <mat-icon class="close" (click)="drawer.close()">settings</mat-icon>
       </mat-toolbar-row>

       <mat-toolbar-row  class="submenu">
            <div class="font-header-menu">Sair</div>
            <mat-icon class="close" (click)="exit()">exit_to_app</mat-icon>
       </mat-toolbar-row>

    </mat-toolbar>
  </mat-drawer>

  <mat-toolbar color="primary">
    <mat-icon *ngIf="!drawer.opened" (click)="drawer.open()" class="font-header-menu">menu</mat-icon>
    <div class="font-header-menu">Iroko App</div>
  </mat-toolbar>

  <div class="main contents">
    <div style="min-height: 100%;">
        <router-outlet></router-outlet>
    </div>
  </div>


  <mat-toolbar class="main footer">
    Footer
  </mat-toolbar>
</mat-drawer-container>
`
})
export class PagesComponent {
  showFiller = false;

  constructor(private router: Router) {
  }

  exit() {
    localStorage.setItem('login', '');
    this.router.navigateByUrl('/login');
  }
}
