import { Controller, Get, NotFoundException } from '@nestjs/common';
import { AppService } from './app.service';
import { Client } from 'pg';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) { }

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }

  @Get('user')
  async getUsers(): Promise<{ source_code: string; users: string[] }> {
    const client = new Client({ connectionString: process.env.DATABASE_URL });
    await client.connect();
    let result;
    try {
      result = await client.query<{ name: string }>('SELECT name FROM users');
    } finally {
      await client.end();
    }
    if (result.rows.length === 0) {
      throw new NotFoundException({ code: 'NO_DATA', message: 'No users found.' });
    }
    return {
      source_code: 'nestjs',
      users: result.rows.map((user) => user.name),
    };
  }
}
