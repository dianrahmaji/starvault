import { MigrationInterface, QueryRunner } from 'typeorm';

export class Migration1783351025481 implements MigrationInterface {
  name = 'Migration1783351025481';

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE "stream" ADD "createdAt" TIMESTAMP NOT NULL DEFAULT now()`,
    );
    await queryRunner.query(
      `ALTER TABLE "stream" ADD "updatedAt" TIMESTAMP NOT NULL DEFAULT now()`,
    );
    await queryRunner.query(
      `ALTER TABLE "stream" DROP CONSTRAINT "PK_0dc9d7e04ff213c08a096f835f2"`,
    );
    await queryRunner.query(`ALTER TABLE "stream" DROP COLUMN "id"`);
    await queryRunner.query(`ALTER TABLE "stream" ADD "id" uuid NOT NULL`);
    await queryRunner.query(
      `ALTER TABLE "stream" ADD CONSTRAINT "PK_0dc9d7e04ff213c08a096f835f2" PRIMARY KEY ("id")`,
    );
    await queryRunner.query(
      `ALTER TABLE "stream" ALTER COLUMN "startedAt" DROP DEFAULT`,
    );
    await queryRunner.query(
      `ALTER TABLE "stream" DROP CONSTRAINT "FK_2d71d3871be1ad3f6224009e82d"`,
    );
    await queryRunner.query(
      `ALTER TABLE "creator" ALTER COLUMN "id" DROP DEFAULT`,
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
      `ALTER TABLE "creator" ALTER COLUMN "id" SET DEFAULT uuid_generate_v4()`,
    );
    await queryRunner.query(
      `ALTER TABLE "stream" ADD CONSTRAINT "FK_2d71d3871be1ad3f6224009e82d" FOREIGN KEY ("creatorId") REFERENCES "creator"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`,
    );
    await queryRunner.query(
      `ALTER TABLE "stream" ALTER COLUMN "startedAt" SET DEFAULT now()`,
    );
    await queryRunner.query(
      `ALTER TABLE "stream" DROP CONSTRAINT "PK_0dc9d7e04ff213c08a096f835f2"`,
    );
    await queryRunner.query(`ALTER TABLE "stream" DROP COLUMN "id"`);
    await queryRunner.query(`ALTER TABLE "stream" ADD "id" integer NOT NULL`);
    await queryRunner.query(
      `ALTER TABLE "stream" ADD CONSTRAINT "PK_0dc9d7e04ff213c08a096f835f2" PRIMARY KEY ("id")`,
    );
    await queryRunner.query(`ALTER TABLE "stream" DROP COLUMN "updatedAt"`);
    await queryRunner.query(`ALTER TABLE "stream" DROP COLUMN "createdAt"`);
  }
}
