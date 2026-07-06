import { MigrationInterface, QueryRunner } from 'typeorm';

export class Migration1783349589106 implements MigrationInterface {
  name = 'Migration1783349589106';

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE "creator" ADD "createdAt" TIMESTAMP NOT NULL DEFAULT now()`,
    );
    await queryRunner.query(
      `ALTER TABLE "creator" ADD "updatedAt" TIMESTAMP NOT NULL DEFAULT now()`,
    );
    await queryRunner.query(
      `ALTER TABLE "stream" DROP CONSTRAINT "FK_2d71d3871be1ad3f6224009e82d"`,
    );
    await queryRunner.query(`ALTER TABLE "stream" DROP COLUMN "creatorId"`);
    await queryRunner.query(`ALTER TABLE "stream" ADD "creatorId" uuid`);
    await queryRunner.query(
      `ALTER TABLE "creator" DROP CONSTRAINT "PK_43e489c9896f9eb32f7a0b912c2"`,
    );
    await queryRunner.query(`ALTER TABLE "creator" DROP COLUMN "id"`);
    await queryRunner.query(
      `ALTER TABLE "creator" ADD "id" uuid NOT NULL DEFAULT gen_random_uuid()`,
    );
    await queryRunner.query(
      `ALTER TABLE "creator" ADD CONSTRAINT "PK_43e489c9896f9eb32f7a0b912c2" PRIMARY KEY ("id")`,
    );
    await queryRunner.query(
      `ALTER TABLE "stream" ADD CONSTRAINT "FK_2d71d3871be1ad3f6224009e82d" FOREIGN KEY ("creatorId") REFERENCES "creator"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`,
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE "stream" DROP CONSTRAINT "FK_2d71d3871be1ad3f6224009e82d"`,
    );
    await queryRunner.query(
      `ALTER TABLE "creator" DROP CONSTRAINT "PK_43e489c9896f9eb32f7a0b912c2"`,
    );
    await queryRunner.query(`ALTER TABLE "creator" DROP COLUMN "id"`);
    await queryRunner.query(`ALTER TABLE "creator" ADD "id" SERIAL NOT NULL`);
    await queryRunner.query(
      `ALTER TABLE "creator" ADD CONSTRAINT "PK_43e489c9896f9eb32f7a0b912c2" PRIMARY KEY ("id")`,
    );
    await queryRunner.query(`ALTER TABLE "stream" DROP COLUMN "creatorId"`);
    await queryRunner.query(`ALTER TABLE "stream" ADD "creatorId" integer`);
    await queryRunner.query(
      `ALTER TABLE "stream" ADD CONSTRAINT "FK_2d71d3871be1ad3f6224009e82d" FOREIGN KEY ("creatorId") REFERENCES "creator"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`,
    );
    await queryRunner.query(`ALTER TABLE "creator" DROP COLUMN "updatedAt"`);
    await queryRunner.query(`ALTER TABLE "creator" DROP COLUMN "createdAt"`);
  }
}
