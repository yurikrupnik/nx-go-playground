import { Controller, Get, Post, Logger, Res, Req } from '@nestjs/common';
import { Request, Response } from 'express';

/* eslint @nrwl/nx/enforce-module-boundaries: 0 */
// HTTP
// import { function } from '../../<path to http function>/src/main'

// PUB-SUB
// import { pubSub  } from ''../../<path to http function>/src/main'

@Controller()
export class RunnerController {
  private readonly logger = new Logger('FunctionsRunner');

  @Get('/*')
  get(@Req() req, @Res() res): Promise<void> {
    return this.executeFunction(req, res);
  }

  @Post('/*')
  post(@Req() req, @Res() res): Promise<void> {
    return this.executeFunction(req, res);
  }

  private async executeFunction(req: Request, res: Response): Promise<void> {
    this.logger.log(`Handle "${req.path}"`);

    switch (req.path) {
      case '/the function endpoint':
        // await function(req, res)
        break;

      case '/pub-sub-event':
        // await this.simulatePubSubEvent(pubSub, req, res)
        break;

      default:
        res.send('Function not found!');
    }
  }

  private async simulatePubSubEvent(
    pubSub,
    req: Request,
    res: Response
  ): Promise<void> {
    if (
      req.method === 'POST' &&
      req.headers['content-type'].indexOf('application/json') > -1 &&
      Object.getPrototypeOf(req.body) === Object.prototype
    ) {
      const message = { attributes: null, data: null };
      if (req.body.attributes) {
        message.attributes = req.body.attributes;
      }

      if (req.body.data) {
        message.data = Buffer.from(
          Object.getPrototypeOf(req.body.data) === Object.prototype
            ? JSON.stringify(req.body.data)
            : req.body.data,
          'binary'
        ).toString('base64');
      }

      try {
        await pubSub(message);

        res.send('ok');
      } catch (err) {
        res.status(500).send(err);
      }
    }

    res.status(405).send();
  }
}
