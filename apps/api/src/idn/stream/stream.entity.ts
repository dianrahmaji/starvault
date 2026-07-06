import {
  Column,
  CreateDateColumn,
  Entity,
  ManyToOne,
  PrimaryColumn,
  UpdateDateColumn,
} from 'typeorm';
import { v7 } from 'uuid';
import { Creator } from '../creator/creator.entity';

@Entity()
export class Stream {
  @PrimaryColumn({ type: 'uuid' })
  id: string = v7();

  @Column()
  title: string;

  @Column({ nullable: true })
  videoUrl?: string;

  @Column({ nullable: true })
  thumbnailUrl?: string;

  @Column({ type: 'timestamptz' })
  startedAt: Date;

  @Column({ type: 'timestamptz', nullable: true })
  endedAt?: Date;

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;

  @ManyToOne(() => Creator, (creator) => creator.streams)
  creator: Creator;
}
