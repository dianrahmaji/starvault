import { Field, Int, ObjectType } from '@nestjs/graphql';
import { Column, Entity, OneToMany, PrimaryGeneratedColumn } from 'typeorm';
import { Stream } from '../stream/stream.entity';

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

  @Field()
  @Column({ default: false })
  isLivestreaming: boolean;

  @Field()
  @Column({ default: false })
  isRecording: boolean;

  @Field()
  @Column({ default: false })
  isRecordingEnabled: boolean;

  @OneToMany(() => Stream, (stream) => stream.creator)
  streams: Stream[];
}
