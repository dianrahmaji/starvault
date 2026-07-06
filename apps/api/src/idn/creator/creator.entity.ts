import { Field, ObjectType } from '@nestjs/graphql';
import {
  Column,
  CreateDateColumn,
  Entity,
  OneToMany,
  PrimaryColumn,
  UpdateDateColumn,
} from 'typeorm';
import { v7 } from 'uuid';
import { Stream } from '../stream/stream.entity';

@ObjectType()
@Entity()
export class Creator {
  @Field()
  @PrimaryColumn({ type: 'uuid' })
  id: string = v7();

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

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;

  @OneToMany(() => Stream, (stream) => stream.creator)
  streams: Stream[];
}
