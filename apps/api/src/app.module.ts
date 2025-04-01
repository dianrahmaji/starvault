import { MikroOrmModule } from '@mikro-orm/nestjs';
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';

import { AppController } from './app.controller';
import { AppService } from './app.service';
import config from './mikro-orm.config';

@Module({
	controllers: [AppController],
	imports: [ConfigModule.forRoot(), MikroOrmModule.forRoot(config)],
	providers: [AppService],
})
export class AppModule {}
