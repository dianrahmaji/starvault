import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm';

@Entity()
export class Creator {
  @PrimaryGeneratedColumn()
  id!: number;

  @Column({ unique: true })
  externalId: string;

  @Column()
  username: string;

  @Column()
  name: string;
}
