import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateUserDto } from './dtos/create-user.dto';
import { User } from './user.entity';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(User)
    private userRepository: Repository<User>,
  ) {}

  findOneWithPassword(username: string) {
    return this.userRepository.findOne({
      where: { username },
      select: {
        id: true,
        username: true,
        password: true,
      },
    });
  }

  getIsExistByUsername(username: string) {
    return this.userRepository.exists({ where: { username } });
  }

  createUser(data: CreateUserDto) {
    const user = this.userRepository.create(data);

    return this.userRepository.save(user);
  }
}
