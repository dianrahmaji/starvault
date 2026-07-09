import {
  ConflictException,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import * as argon2 from 'argon2';
import { UserService } from '../user/user.service';
import { SignInDto } from './dtos/sign-in.dto';
import { SignUpDto } from './dtos/sign-up.dto';

type Token = {
  accessToken: string;
};

@Injectable()
export class AuthService {
  constructor(
    private readonly userService: UserService,
    private readonly jwtService: JwtService,
  ) {}

  async signIn(credentials: SignInDto): Promise<Token> {
    const user = await this.userService.findOneWithPassword(
      credentials.username,
    );

    if (!user) {
      throw new UnauthorizedException();
    }
    const { password } = user;

    const isPasswordValid = await argon2.verify(password, credentials.password);

    if (!isPasswordValid) {
      throw new UnauthorizedException();
    }

    const payload = { sub: user.id, username: user.username };

    const accessToken = await this.jwtService.signAsync(payload);

    return {
      accessToken,
    };
  }

  async signUp(data: SignUpDto): Promise<Token> {
    const { password, username, ...rest } = data;

    const isUsernameExist =
      await this.userService.getIsExistByUsername(username);

    if (isUsernameExist) {
      throw new ConflictException('Username already taken');
    }

    const hashedPassword = await argon2.hash(password);

    const user = await this.userService.createUser({
      ...rest,
      username,
      password: hashedPassword,
    });

    const payload = { sub: user.id, username: user.username };

    const accessToken = await this.jwtService.signAsync(payload);

    return {
      accessToken,
    };
  }
}
