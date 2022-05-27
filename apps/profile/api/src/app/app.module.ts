import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import {
  NestLoggerModule,
  BackendDocsModule,
  HealthModule,
} from '@nx-go-playground/nest/modules';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UsersModule } from '@nx-go-playground/api/users';
import {
  StatusMonitorModule,
  SpansConfiguration,
  // StatusMonitorConfiguration,
  ChartVisibilityConfiguration,
  // HealthCheckConfiguration,
} from 'nestjs-status-monitor';

const spans: SpansConfiguration[] = [
  {
    interval: 1, // Every second
    retention: 60, // Keep 60 datapoints in memory
  },
  {
    interval: 5, // Every 5 seconds
    retention: 60,
  },
  {
    interval: 15, // Every 15 seconds
    retention: 60,
  },
];

const chartVisibility: ChartVisibilityConfiguration = {
  cpu: true,
  mem: true,
  load: true,
  rps: true,
  responseTime: true,
  statusCodes: true,
};

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
    NestLoggerModule,
    StatusMonitorModule.forRoot({
      title: 'NestJS Status', // Default title
      path: '/status',
      socketPath: '/socket.io', // In case you use a custom path
      port: null, // Defaults to NestJS port
      spans,
      chartVisibility,
      ignoreStartsWith: ['/admin'], // paths to ignore for responseTime stats
      healthChecks: [],
    }),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
