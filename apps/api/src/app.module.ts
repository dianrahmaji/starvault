import { MikroOrmModule } from '@mikro-orm/nestjs';
import { BullModule } from '@nestjs/bullmq';
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';

import { AppController } from './app.controller';
import { AppService } from './app.service';
import config from './mikro-orm.config';

@Module({
	controllers: [AppController],
	imports: [
		ConfigModule.forRoot(),
		MikroOrmModule.forRoot(config),
		BullModule.forRoot({
			connection: {
				host: process.env.REDIS_HOST,
				port: parseInt(process.env.REDIS_PORT!),
			},
		}),
	],
	providers: [AppService],
})
export class AppModule {}
