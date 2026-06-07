import { IsBoolean, IsNotEmpty, IsString } from 'class-validator';

export class UpsertCreatorDto {
  @IsNotEmpty()
  @IsString()
  readonly externalId: string;

  @IsNotEmpty()
  @IsString()
  readonly name: string;

  @IsNotEmpty()
  @IsString()
  readonly username: string;

  @IsNotEmpty()
  @IsBoolean()
  readonly isLivestreaming: boolean;
}
