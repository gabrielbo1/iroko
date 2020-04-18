import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {PagesComponent} from "./pages.component";
import { DashboardComponent } from './dashboard/dashboard.component';
import {PagesRoutingModule} from "./pages-routing.module";
import {MaterialModule} from "../material.module";



@NgModule({
  declarations: [
    PagesComponent,
    DashboardComponent
  ],
  imports: [
    CommonModule,
    PagesRoutingModule,
    MaterialModule
  ]
})
export class PagesModule { }
