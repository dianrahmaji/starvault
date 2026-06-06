import { MigrationInterface, QueryRunner } from 'typeorm';

export class Migration1780755299802 implements MigrationInterface {
  name = 'Migration1780755299802';

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `CREATE TABLE "creator" ("id" SERIAL NOT NULL, "externalId" character varying NOT NULL, "username" character varying NOT NULL, "name" character varying NOT NULL, CONSTRAINT "PK_43e489c9896f9eb32f7a0b912c2" PRIMARY KEY ("id"))`,
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`DROP TABLE "creator"`);
  }
}
