import { NgIf } from '@angular/common';
import { Component, model, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { YaCaptchaComponent } from '../../components/ya-captcha/ya-captcha.component';
import { BaseService } from '../../services/base.service';
import { Config, Shorten } from '../../services/models';

@Component({
  selector: 'app-shorten',
  imports: [
    FormsModule,
    YaCaptchaComponent,
    NgIf,
  ],
  templateUrl: './shorten.component.html',
  styleUrl: './shorten.component.scss',
  standalone: true,
})
export class ShortenComponent implements OnInit {

  config?: Config;
  model: Shorten = new Shorten();
  err?: string;

  constructor(
    private readonly base: BaseService,
  ) {
  }

  ngOnInit(): void {
    this.base.getConfig().subscribe(value => {
      this.config = value;
    });
  }

  save(): void {
    if (this.model.source.length === 0) {
      this.err = 'Source URL is empty';
      return;
    }
    if (this.config?.captchaUse === true && this.model.token.length === 0) {
      this.err = 'Captcha has not been verified';
      return;
    }
    this.base.newShort(this.model).subscribe({
      next: (v) => {
        this.model = v;
        this.err = undefined;
      },
      error: (e) => this.err = e.error.msg,
    });
  }

  reset(): void {
    window.location.reload();
  }

  captchaCallback(token: string): void {
    this.model.token = token;
  }

}
