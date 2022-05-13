import { Controller, Get } from '@nestjs/common';

import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getData() {
    return this.appService.getData();
  }

  @Get('/item')
  getItem() {
    return this.appService.getItem();
  }

  @Get('/items')
  getItems() {
    return this.appService.getItems();
  }
}
