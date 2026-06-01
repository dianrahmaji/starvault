import { HttpModule } from '@nestjs/axios';
import { Module } from '@nestjs/common';
import { SchedulerService } from './scheduler.service';

@Module({
  imports: [HttpModule],
  providers: [SchedulerService],
})
export class SchedulerModule {}
