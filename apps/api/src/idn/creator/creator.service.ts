import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
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
    await this.creatorRepository.upsert(creators, {
      skipUpdateIfNoValuesChanged: true,
      conflictPaths: {
        externalId: true,
      },
    });
  }
}
