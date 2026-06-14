import { MigrationInterface, QueryRunner } from 'typeorm';

export class Migration1781445692952 implements MigrationInterface {
  name = 'Migration1781445692952';

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `CREATE TABLE "stream" ("id" integer NOT NULL, "title" character varying NOT NULL, "videoUrl" character varying, "thumbnailUrl" character varying, "startedAt" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "endedAt" TIMESTAMP WITH TIME ZONE, "creatorId" integer, CONSTRAINT "PK_0dc9d7e04ff213c08a096f835f2" PRIMARY KEY ("id"))`,
    );
    await queryRunner.query(
      `ALTER TABLE "stream" ADD CONSTRAINT "FK_2d71d3871be1ad3f6224009e82d" FOREIGN KEY ("creatorId") REFERENCES "creator"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`,
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE "stream" DROP CONSTRAINT "FK_2d71d3871be1ad3f6224009e82d"`,
    );
    await queryRunner.query(`DROP TABLE "stream"`);
  }
}
