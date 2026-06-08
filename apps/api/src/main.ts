import fastifyBasicAuth from '@fastify/basic-auth';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';
import {
  FastifyAdapter,
  NestFastifyApplication,
} from '@nestjs/platform-fastify';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create<NestFastifyApplication>(
    AppModule,
    new FastifyAdapter(),
  );

  const config = app.get(ConfigService);
  const fastify = app.getHttpAdapter().getInstance();

  await fastify.register(fastifyBasicAuth, {
    validate(username, password, _req, _reply, done) {
      const ok =
        username === config.getOrThrow<string>('BULL_BOARD_USERNAME') &&
        password === config.getOrThrow<string>('BULL_BOARD_PASSWORD');
      done(ok ? undefined : new Error('Unauthorized'));
    },
    authenticate: { realm: 'Bull Board' },
  });

  fastify.addHook('onRequest', (req, reply, done) => {
    if (req.url.startsWith('/queues')) {
      fastify.basicAuth(req, reply, done);
    } else {
      done();
    }
  });

  await app.listen(process.env.PORT ?? 3000);
}
bootstrap();
