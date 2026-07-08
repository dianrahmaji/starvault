import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { v7 } from 'uuid';
import { Creator } from './creator.entity';
import { UpsertCreatorDto } from './dtos/upsert-creator.dto';

@Injectable()
export class CreatorService {
  constructor(
    @InjectRepository(Creator)
    private creatorRepository: Repository<Creator>,
  ) {}

  findAll(): Promise<Creator[]> {
    return this.creatorRepository.find();
  }

  async upsertMany(creators: UpsertCreatorDto[]) {
    await this.creatorRepository.upsert(
      creators.map((creator) => ({ id: v7(), ...creator })),
      {
        skipUpdateIfNoValuesChanged: true,
        conflictPaths: {
          externalId: true,
        },
      },
    );
  }

  async resetLivestreamingStatus(activeExternalIds: string[]) {
    if (!activeExternalIds.length) {
      return;
    }

    await this.creatorRepository
      .createQueryBuilder()
      .update(Creator)
      .set({
        isLivestreaming: false,
      })
      .where('isLivestreaming = true AND "externalId" NOT IN (:...ids)', {
        ids: activeExternalIds,
      })
      .execute();
  }
}
