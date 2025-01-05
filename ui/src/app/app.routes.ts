import { Routes } from '@angular/router';
import { BadgesComponent } from './pages/badges/badges.component';
import { LicenseComponent } from './pages/license/license.component';
import { Page404Component } from './pages/page404/page404.component';
import { ShortenComponent } from './pages/shorten/shorten.component';

export const routes: Routes = [
  { path: '', component: ShortenComponent, title: 'UriOne | Shorten' },
  { path: 'badges', component: BadgesComponent, title: 'UriOne | Badges' },
  { path: 'license', component: LicenseComponent, title: 'UriOne | License' },
  { path: '**', component: Page404Component, pathMatch: 'full' },
];
