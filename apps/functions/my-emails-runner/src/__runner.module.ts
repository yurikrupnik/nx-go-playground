import { Module } from '@nestjs/common';
import { RunnerController } from './__runner.controller';

@Module({
  imports: [],
  controllers: [RunnerController],
  providers: [],
})
export class RunnerModule {}
