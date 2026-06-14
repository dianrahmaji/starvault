import { Module } from '@nestjs/common';
import { CreatorModule } from './creator/creator.module';
import { SchedulerModule } from './scheduler/scheduler.module';
import { StreamModule } from './stream/stream.module';

@Module({
  imports: [SchedulerModule, CreatorModule, StreamModule],
})
export class IdnModule {}
