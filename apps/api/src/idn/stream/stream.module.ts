import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Stream } from './stream.entity';
import { StreamResolver } from './stream.resolver';
import { StreamService } from './stream.service';

@Module({
  imports: [TypeOrmModule.forFeature([Stream])],
  providers: [StreamService, StreamResolver],
})
export class StreamModule {}
