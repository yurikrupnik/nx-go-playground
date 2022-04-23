import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { NestLoggerModule, BackendDocsModule, HealthModule } from '@nx-go-playground/nest/modules'
import { AppController } from './app.controller';
import { AppService } from './app.service';
import {UsersModule} from '@nx-go-playground/api/users'

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      // cache: true,
    }),
    // AuthModule,
    UsersModule,
    HealthModule,
    BackendDocsModule,
    NestLoggerModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
