import { Module, Provider } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Creator } from './creator.entity';
import { CreatorResolver } from './creator.resolver';
import { CreatorService } from './creator.service';

const services: Provider[] = [
  //
  CreatorService,
];

const resolvers: Provider[] = [
  //
  CreatorResolver,
];

@Module({
  imports: [TypeOrmModule.forFeature([Creator])],
  providers: [
    //
    ...services,
    ...resolvers,
  ],
  exports: [CreatorService],
})
export class CreatorModule {}
