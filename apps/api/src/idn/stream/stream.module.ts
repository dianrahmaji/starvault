import { Module } from '@nestjs/common';
import { StreamResolver } from './stream.resolver';
import { StreamService } from './stream.service';

@Module({
  providers: [StreamService, StreamResolver],
})
export class StreamModule {}
