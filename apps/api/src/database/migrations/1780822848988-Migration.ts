import { MigrationInterface, QueryRunner } from 'typeorm';

export class Migration1780822848988 implements MigrationInterface {
  name = 'Migration1780822848988';

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE "creator" ADD CONSTRAINT "UQ_0a77348632c3550ddf0805ab422" UNIQUE ("externalId")`,
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE "creator" DROP CONSTRAINT "UQ_0a77348632c3550ddf0805ab422"`,
    );
  }
}
