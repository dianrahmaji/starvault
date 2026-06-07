import { Field, Int, ObjectType } from '@nestjs/graphql';
import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm';

@ObjectType()
@Entity()
export class Creator {
  @Field(() => Int)
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ unique: true })
  externalId: string;

  @Field()
  @Column()
  username: string;

  @Field()
  @Column()
  name: string;
}
