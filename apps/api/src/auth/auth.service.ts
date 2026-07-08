import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { UserService } from '../user/user.service';
import { SignInDto } from './dtos/sign-in.dto';

@Injectable()
export class AuthService {
  constructor(
    private readonly userService: UserService,
    private readonly jwtService: JwtService,
  ) {}

  async signIn(credentials: SignInDto): Promise<{ accessToken: string }> {
    const user = await this.userService.findOneWithPassword(
      credentials.username,
    );

    if (!user) {
      throw new UnauthorizedException();
    }
    const { password } = user;

    // TODO: implement hash compare
    if (password !== credentials.password) {
      throw new UnauthorizedException();
    }

    const payload = { sub: user.id, username: user.username };

    const accessToken = await this.jwtService.signAsync(payload);

    return {
      accessToken,
    };
  }
}
