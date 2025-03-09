import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MikroOrmModule } from '@mikro-orm/nestjs';
import { ConfigModule } from '@nestjs/config';
import config from './mikro-orm.config';

@Module({
	imports: [ConfigModule.forRoot(), MikroOrmModule.forRoot(config)],
	controllers: [AppController],
	providers: [AppService],
})
export class AppModule {}
