export interface Config {
  address: string
  orgName: string
  serviceName: string
  email: string
  captchaUse: boolean
  captchaClientKey: string
}

export interface Err {
  msg: string
}

export class Shorten {
  url: string;
  source: string;
  token: string;

  constructor() {
    this.url = '';
    this.source = '';
    this.token = '';
  }
}
