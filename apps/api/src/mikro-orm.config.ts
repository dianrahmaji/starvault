import { SqlHighlighter } from '@mikro-orm/sql-highlighter';
import { defineConfig, Options } from '@mikro-orm/postgresql';
import { Dummy } from './entities/Dummy.entity';
import * as dotenv from 'dotenv';
import { Migrator } from '@mikro-orm/migrations';
import { SeedManager } from '@mikro-orm/seeder';

dotenv.config();

const config: Options = {
	dbName: process.env.DB_NAME,
	password: process.env.DB_PASSWORD,
	user: process.env.DB_USER,
	port: parseInt(process.env.DB_PORT),
	highlighter: new SqlHighlighter(),
	debug: process.env.ENV.endsWith('local'),
	extensions: [SeedManager, Migrator],
	migrations: {
		snapshot: false,
	},
	entities: [
		//
		Dummy,
	],
};

export default defineConfig(config);
