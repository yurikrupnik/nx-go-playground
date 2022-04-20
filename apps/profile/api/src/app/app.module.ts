import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { NestLoggerModule, BackendDocsModule } from '@nx-go-playground/nest/modules'
import { AppController } from './app.controller';
import { AppService } from './app.service';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      // cache: true,
    }),
    BackendDocsModule,
    NestLoggerModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
