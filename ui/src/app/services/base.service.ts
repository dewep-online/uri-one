import { Injectable, OnInit } from '@angular/core';
import { RequestService } from '@onega-ui/core';
import { Observable } from 'rxjs';
import { Config, Shorten } from './models';

@Injectable({
  providedIn: 'root',
  deps:[RequestService],
})
export class BaseService {

  constructor(
    private readonly req: RequestService,
  ) { }

  getConfig(): Observable<Config> {
    return this.req.get<Config>('/config.json');
  }

  newShort(m: Shorten): Observable<Shorten> {
    return this.req.post<Shorten>('/shorten/add', m);
  }
}
