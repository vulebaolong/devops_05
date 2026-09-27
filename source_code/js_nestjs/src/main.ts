import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import 'dotenv/config';
import { JsonExceptionFilter } from './http-exception.filter';
import { RequestLoggerMiddleware } from './request-logger.middleware';

async function bootstrap() {
  console.log(`DATABASE_URL=${process.env.DATABASE_URL}`);
  const app = await NestFactory.create(AppModule);
  const requestLogger = new RequestLoggerMiddleware();
  app.use(requestLogger.use.bind(requestLogger));
  app.useGlobalFilters(new JsonExceptionFilter());
  app.enableCors({
    origin: '*',
    methods: '*',
    allowedHeaders: '*',
  });
  await app.listen(process.env.PORT ?? 3000);
}
bootstrap();
