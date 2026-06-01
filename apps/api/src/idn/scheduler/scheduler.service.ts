import { HttpService } from '@nestjs/axios';
import { Injectable } from '@nestjs/common';
import { Cron, CronExpression } from '@nestjs/schedule';
import { firstValueFrom } from 'rxjs';

type Creator = {
  name: string;
  uuid: string;
};

type LiveStream = {
  title: string;
  image_url: string;
  playback_url: string;
  status: string;
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
        name
        uuid
      }
    }
  }
`;

@Injectable()
export class SchedulerService {
  constructor(private readonly httpService: HttpService) {}

  @Cron(CronExpression.EVERY_10_SECONDS)
  async getLiveStreams() {
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
  }
}
