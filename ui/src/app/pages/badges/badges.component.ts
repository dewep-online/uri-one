import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BaseService } from '../../services/base.service';
import { Config } from '../../services/models';

@Component({
  selector: 'app-badges',
  imports: [ FormsModule ],
  templateUrl: './badges.component.html',
  styleUrl: './badges.component.scss',
  standalone: true,
})
export class BadgesComponent implements OnInit {
  config?: Config;

  builderColor = 'primary';
  builderTitle = 'Title Example';
  builderValue = 'Title Value';

  constructor(
    private readonly base: BaseService,
  ) {
  }

  ngOnInit(): void {
    this.base.getConfig().subscribe(value => this.config = value);
  }
}
