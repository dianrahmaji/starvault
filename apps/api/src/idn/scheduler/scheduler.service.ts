import { HttpService } from '@nestjs/axios';
import { Injectable } from '@nestjs/common';
import { Cron, CronExpression } from '@nestjs/schedule';
import { firstValueFrom } from 'rxjs';
import { CreatorService } from '../creator/creator.service';

type Creator = {
  name: string;
  uuid: string;
  username: string;
};

enum LivestreamStatus {
  Live = 'live',
  Scheduled = 'scheduled',
}

type LiveStream = {
  title: string;
  image_url: string;
  playback_url: string;
  status: LivestreamStatus;
  creator: Creator;
};

type GraphqlResponse = {
  data: {
    getLivestreams: LiveStream[];
  };
};

const IDN_GRAPHQL_URL = 'https://api.idn.app/graphql';
const GET_LIVESTREAMS_QUERY = `
  query($page: Int) {
    getLivestreams(category: "all", page: $page) {
      title
      image_url
      playback_url
      status
      creator {
        uuid
        name
        username
      }
    }
  }
`;

@Injectable()
export class SchedulerService {
  constructor(
    private readonly httpService: HttpService,
    private readonly creatorService: CreatorService,
  ) {}

  @Cron(CronExpression.EVERY_MINUTE)
  async getLiveStreams() {
    const livestreams = await this.fetchLivestreams();

    const activeExternalIds = livestreams
      .filter((livestream) => {
        return livestream.status === LivestreamStatus.Live;
      })
      .map((livesteram) => {
        return livesteram.creator.uuid;
      });

    const activeExternalIdSet = new Set(activeExternalIds);

    const creatorEntries = livestreams.map(({ creator: { uuid, ...rest } }) => {
      const isLivestreaming = activeExternalIdSet.has(uuid);

      return [uuid, { externalId: uuid, isLivestreaming, ...rest }] as const;
    });

    const creatorsMap = new Map(creatorEntries);

    const creators = [...creatorsMap.values()];

    await this.creatorService.upsertMany(creators);

    await this.creatorService.resetLivestreamingStatus(activeExternalIds);
  }

  private async fetchLivestreams() {
    const allLiveStreams: LiveStream[] = [];

    for (let page = 1; ; page++) {
      const res = await firstValueFrom(
        this.httpService.post<GraphqlResponse>(IDN_GRAPHQL_URL, {
          query: GET_LIVESTREAMS_QUERY,
          variables: {
            page,
          },
        }),
      );

      const liveStreams = res.data.data.getLivestreams;

      if (!liveStreams.length) {
        break;
      }

      allLiveStreams.push(...liveStreams);
    }

    return allLiveStreams;
  }
}
