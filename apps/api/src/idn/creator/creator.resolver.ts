import { Query, Resolver } from '@nestjs/graphql';
import { Creator } from './creator.entity';
import { CreatorService } from './creator.service';

@Resolver(() => Creator)
export class CreatorResolver {
  constructor(private readonly creatorService: CreatorService) {}

  @Query(() => [Creator])
  async creators() {
    return this.creatorService.findAll();
  }
}
