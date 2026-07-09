import { IsNotEmpty, IsString } from 'class-validator';

export class SignUpDto {
  @IsNotEmpty()
  @IsString()
  readonly username: string;

  @IsNotEmpty()
  @IsString()
  readonly password: string;

  @IsNotEmpty()
  @IsString()
  readonly firstName: string;

  @IsNotEmpty()
  @IsString()
  readonly lastName: string;
}
