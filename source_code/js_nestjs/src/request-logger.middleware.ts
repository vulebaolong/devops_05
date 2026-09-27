import { Injectable, Logger, NestMiddleware } from '@nestjs/common';
import type { NextFunction, Request, Response } from 'express';

@Injectable()
export class RequestLoggerMiddleware implements NestMiddleware {
  private readonly logger = new Logger(RequestLoggerMiddleware.name);

  use(request: Request, response: Response, next: NextFunction): void {
    const startedAt = Date.now();
    this.logger.log(`HTTP request started: ${request.method} ${request.originalUrl}`);
    response.on('finish', () => {
      this.logger.log(`HTTP request completed: ${request.method} ${request.originalUrl} - ${response.statusCode} in ${Date.now() - startedAt}ms`);
    });
    next();
  }
}
