import {
  Column,
  CreateDateColumn,
  Entity,
  PrimaryColumn,
  UpdateDateColumn,
} from 'typeorm';
import { v7 } from 'uuid';

@Entity()
export class User {
  @PrimaryColumn({ type: 'uuid' })
  id: string = v7();

  @Column({ unique: true })
  username: string;

  @Column({ select: false })
  password: string;

  @Column()
  firstName: string;

  @Column()
  lastName: string;

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;
}
