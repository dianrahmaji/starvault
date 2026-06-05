import { Module } from '@nestjs/common';
import { CreatorModule } from './creator/creator.module';
import { SchedulerModule } from './scheduler/scheduler.module';

@Module({
  imports: [SchedulerModule, CreatorModule],
})
export class IdnModule {}
