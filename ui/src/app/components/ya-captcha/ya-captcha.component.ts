import { DOCUMENT } from '@angular/common';
import {
  AfterViewInit,
  Component,
  ElementRef,
  EventEmitter, Inject,
  Input,
  OnDestroy,
  Output, Renderer2,
  RendererFactory2,
  ViewChild,
} from '@angular/core';
import { SmartCaptcha } from './ya-captcha.types';

@Component({
  selector: 'app-ya-captcha',
  imports: [],
  templateUrl: './ya-captcha.component.html',
  styleUrl: './ya-captcha.component.scss',
  standalone: true,
})
export class YaCaptchaComponent implements AfterViewInit, OnDestroy {
  private renderer: Renderer2;
  private smartCaptcha?: SmartCaptcha;
  private widgetId?: number;

  @Input({ required: true }) clientKey?: string;
  @Output() onCallback: EventEmitter<string> = new EventEmitter<string>();
  @ViewChild('captcha', { read: ElementRef }) captcha?: ElementRef<HTMLDivElement>;

  constructor(
    private readonly rendererFactory: RendererFactory2,
  ) {
    this.renderer = rendererFactory.createRenderer(null, null);
    this.smartCaptcha = window.smartCaptcha;
  }

  ngAfterViewInit(): void {
    this.widgetId = this.smartCaptcha?.render(this.captcha!.nativeElement, {
      sitekey: this.clientKey,
      callback: (token: string)=>this.onCallback.emit(token),
    });
  }

  ngOnDestroy(): void {
    this.smartCaptcha?.destroy(this.widgetId);
  }
}
