import { Migrator } from '@mikro-orm/migrations';
import { defineConfig, Options } from '@mikro-orm/postgresql';
import { SeedManager } from '@mikro-orm/seeder';
import { SqlHighlighter } from '@mikro-orm/sql-highlighter';
import * as dotenv from 'dotenv';

import { Dummy } from './entities/Dummy.entity';

dotenv.config();

const config: Options = {
	dbName: process.env.DB_NAME,
	debug: process.env.ENV!.endsWith('local'),
	entities: [
		//
		Dummy,
	],
	extensions: [SeedManager, Migrator],
	highlighter: new SqlHighlighter(),
	migrations: {
		snapshot: false,
	},
	password: process.env.DB_PASSWORD,
	port: parseInt(process.env.DB_PORT!),
	user: process.env.DB_USER,
};

export default defineConfig(config);
