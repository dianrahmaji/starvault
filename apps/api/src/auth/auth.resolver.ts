import {
  Args,
  Field,
  InputType,
  Mutation,
  ObjectType,
  Resolver,
} from '@nestjs/graphql';
import { AuthService } from './auth.service';

@InputType()
class SignInInput {
  @Field()
  username: string;

  @Field()
  password: string;
}

@ObjectType()
class AuthToken {
  @Field()
  accessToken: string;
}

@Resolver()
export class AuthResolver {
  constructor(private readonly authService: AuthService) {}

  @Mutation(() => AuthToken)
  signIn(@Args('signInInput') signInInput: SignInInput) {
    return this.authService.signIn(signInInput);
  }
}
