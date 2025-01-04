export interface RenderParams {
  sitekey?: string;
  callback?: (token: string) => void;
}

// type SubscribeEvent =
//   | 'challenge-visible'
//   | 'challenge-hidden'
//   | 'network-error'
//   | 'javascript-error'
//   | 'success'
//   | 'token-expired';

export interface SmartCaptcha {
  render: (container: HTMLElement | string, params: RenderParams) => number;
  destroy: (widgetId?: number) => void;

  // getResponse: (widgetId?: number) => string;
  // execute: (widgetId?: number) => void;
  // reset: (widgetId?: number) => void;
  // showError: (widgetId?: number) => void;
  // subscribe(
  //   widgetId: number,
  //   event: SubscribeEvent,
  //   callback: Function,
  // ): () => void;
}

declare global {
  interface Window {
    smartCaptcha: SmartCaptcha;
  }
}
