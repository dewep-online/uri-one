import { Component, OnInit } from '@angular/core';
import { BaseService } from '../../services/base.service';
import { Config } from '../../services/models';

@Component({
  selector: 'app-license',
  imports: [],
  templateUrl: './license.component.html',
  styleUrl: './license.component.scss',
  standalone: true,
})
export class LicenseComponent implements OnInit {
  config?: Config;

  constructor(
    private readonly base: BaseService,
  ) {
  }

  ngOnInit(): void {
    this.base.getConfig().subscribe(value => this.config = value);
  }
}
