import { ArgumentsHost, Catch, ExceptionFilter, HttpException, HttpStatus, Logger } from '@nestjs/common';
import type { Request, Response } from 'express';

type DatabaseError = Error & { code?: string };

@Catch()
export class JsonExceptionFilter implements ExceptionFilter {
  private readonly logger = new Logger(JsonExceptionFilter.name);

  catch(exception: unknown, host: ArgumentsHost): void {
    const response = host.switchToHttp().getResponse<Response>();
    const request = host.switchToHttp().getRequest<Request>();
    const error = exception as DatabaseError;
    let status = HttpStatus.INTERNAL_SERVER_ERROR;
    let code = 'INTERNAL_SERVER_ERROR';
    let message = 'An unexpected error occurred.';

    if (exception instanceof HttpException) {
      status = exception.getStatus();
      const body = exception.getResponse();
      if (typeof body === 'object' && body !== null) {
        const value = body as { code?: string; message?: string | string[] };
        code = value.code ?? (status === 404 ? 'NOT_FOUND' : 'HTTP_ERROR');
        message = Array.isArray(value.message) ? value.message.join(', ') : (value.message ?? exception.message);
      } else {
        message = String(body);
      }
    } else if (error.code === '42P01') {
      status = HttpStatus.SERVICE_UNAVAILABLE;
      code = 'TABLE_NOT_FOUND';
      message = "The database table 'users' does not exist.";
    } else if (error.code && ['ECONNREFUSED', 'ECONNRESET', 'ETIMEDOUT', 'ENOTFOUND', '28P01', '3D000'].includes(error.code)) {
      status = HttpStatus.SERVICE_UNAVAILABLE;
      code = 'DB_CONNECTION_ERROR';
      message = 'Cannot connect to the database.';
    }

    if (!(exception instanceof HttpException) || status >= 500) {
      this.logger.error(`${request.method} ${request.url}: ${error.message ?? String(exception)}`, error.stack);
    }
    response.status(status).json({
      source_code: 'nestjs',
      success: false,
      error: { code, message, details: error.message ?? null },
      timestamp: new Date().toISOString(),
      path: request.url,
    });
  }
}
