import {
  Column,
  CreateDateColumn,
  Entity,
  ManyToOne,
  PrimaryColumn,
} from 'typeorm';
import { Creator } from '../creator/creator.entity';

@Entity()
export class Stream {
  @PrimaryColumn()
  id: number;

  @Column()
  title: string;

  @Column({ nullable: true })
  videoUrl?: string;

  @Column({ nullable: true })
  thumbnailUrl?: string;

  @CreateDateColumn({ type: 'timestamptz' })
  startedAt: Date;

  @Column({ type: 'timestamptz', nullable: true })
  endedAt?: Date;

  @ManyToOne(() => Creator, (creator) => creator.streams)
  creator: Creator;
}
