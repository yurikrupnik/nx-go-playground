import { NestFactory } from '@nestjs/core';
import { Logger } from '@nestjs/common';
import { ExpressAdapter } from '@nestjs/platform-express';
import { RunnerModule } from './__runner.module';

async function bootstrap() {
  const app = await NestFactory.create(RunnerModule, new ExpressAdapter());

  await app.listen(8080, '0.0.0.0').then(() => {
    Logger.log(`Functions running on http://localhost:8080}`);
  });
}

bootstrap();
