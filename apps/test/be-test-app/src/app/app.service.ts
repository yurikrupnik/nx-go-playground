import { Injectable } from '@nestjs/common';

@Injectable()
export class AppService {
  getData(): { message: string } {
    return { message: 'Welcome to be-test-app!' };
  }

  getItem() {
    return { message: 'item here' };
  }
}
