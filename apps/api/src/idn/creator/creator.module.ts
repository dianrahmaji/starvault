import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Creator } from './creator.entity';
import { CreatorService } from './creator.service';

@Module({
  imports: [TypeOrmModule.forFeature([Creator])],
  providers: [CreatorService],
  exports: [CreatorService],
})
export class CreatorModule {}
