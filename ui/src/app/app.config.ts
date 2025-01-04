import { provideHttpClient } from '@angular/common/http';
import { ApplicationConfig, importProvidersFrom, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { CoreModule } from '@onega-ui/core';
import { routes } from './app.routes';
import { BaseService } from './services/base.service';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes),
    provideHttpClient(),
    importProvidersFrom(
      CoreModule.forRoot({
        apiHost: '/api',
      }),
    ),
    { provide: BaseService },
  ],
};
