import { Module, Provider } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { User } from './user.entity';
import { UserService } from './user.service';

const services: Provider[] = [
  //
  UserService,
];

@Module({
  imports: [TypeOrmModule.forFeature([User])],
  providers: [
    //
    ...services,
  ],
  exports: [
    //
    UserService,
  ],
})
export class UserModule {}
