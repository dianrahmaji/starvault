import { Field, ObjectType } from '@nestjs/graphql';
import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm';

@ObjectType()
@Entity()
export class Creator {
  @Field()
  @PrimaryGeneratedColumn()
  id!: number;

  @Column({ unique: true })
  externalId: string;

  @Field()
  @Column()
  username: string;

  @Field()
  @Column()
  name: string;
}
