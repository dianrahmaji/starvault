import { HttpModule } from '@nestjs/axios';
import { Module } from '@nestjs/common';
import { CreatorModule } from '../creator/creator.module';
import { SchedulerService } from './scheduler.service';

@Module({
  imports: [HttpModule, CreatorModule],
  providers: [SchedulerService],
})
export class SchedulerModule {}
