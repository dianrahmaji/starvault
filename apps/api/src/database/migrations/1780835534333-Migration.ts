import { MigrationInterface, QueryRunner } from 'typeorm';

export class Migration1780835534333 implements MigrationInterface {
  name = 'Migration1780835534333';

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE "creator" ADD "isLivestreaming" boolean NOT NULL DEFAULT false`,
    );
    await queryRunner.query(
      `ALTER TABLE "creator" ADD "isRecording" boolean NOT NULL DEFAULT false`,
    );
    await queryRunner.query(
      `ALTER TABLE "creator" ADD "isRecordingEnabled" boolean NOT NULL DEFAULT false`,
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE "creator" DROP COLUMN "isRecordingEnabled"`,
    );
    await queryRunner.query(`ALTER TABLE "creator" DROP COLUMN "isRecording"`);
    await queryRunner.query(
      `ALTER TABLE "creator" DROP COLUMN "isLivestreaming"`,
    );
  }
}
