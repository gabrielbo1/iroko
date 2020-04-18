import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from "./login/login.component";
import {AuthenticateService} from "./AuthenticateService";


const routes: Routes = [
  {
    path:'login', component: LoginComponent
  },
  {
    path: 'pages', loadChildren: 'src/app/pages/pages.module#PagesModule',
    canActivate: [AuthenticateService]
  },
  {
    path: '',
    redirectTo: 'pages',
    pathMatch: 'full',
  },
  {
    path: '**',
    redirectTo: 'pages',
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
