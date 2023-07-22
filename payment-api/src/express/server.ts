import express, { Express, Router } from 'express';
import http from 'http';
import compression from 'compression';
import cors from 'cors';
import helmet from 'helmet';
import swaggerUi from 'swagger-ui-express';
import YAML from 'yamljs';
import { middleware } from 'express-openapi-validator';
import { errorLogger, logger as loggerMiddleware } from 'express-winston';
import util from 'util';
import process from 'process';
import config from '@shared/env/env';
import logger from '@shared/utils/logger';

export default class Server {
  private app: Express;

  private server: http.Server;

  constructor(private readonly port: number, private readonly router: Router[]) {
    this.app = express();
    this.server = http.createServer(this.app);
    this.setServerConfig();
    this.setListeners();
  }

  start(): void {
    this.server.listen(this.port, () => {
      logger.info(`⚡ Listening on ${this.port}`);
    });
  }

  setServerConfig(): void {
    this.app.set('port', this.port);
    this.app.set('trust proxy', 1);
    this.app.use(
      helmet({
        contentSecurityPolicy: false,
      })
    );
    this.app.use(compression());
    this.app.use(cors());
    this.app.use(express.json());
    this.app.use(
      express.urlencoded({
        extended: true,
      })
    );
    // Use logging
    this.app.use(
      loggerMiddleware({
        winstonInstance: logger,
        expressFormat: true,
        colorize: true,
        meta: false,
      })
    );
    // Serve Swagger UI when not in production
    const swaggerDocument = YAML.load(config.DIR_SWAGGER || '');
    if (config.NODE_ENV !== 'production') {
      this.app.use(config.DOCS_ENDPOINT, swaggerUi.serve, swaggerUi.setup(swaggerDocument));
      // Validate requests against defined OpenApi spec
      this.app.use(
        middleware({
          apiSpec: config.DIR_SWAGGER || '',
          validateResponses: false,
          validateRequests: true,
          validateSecurity: false,
        })
      );
    }
    // API routes prefix
    this.app.use(this.router);
    this.app.use(
      errorLogger({
        winstonInstance: logger,
      })
    );
  }

  stop(): void {
    logger.info(`Stopping server. Waiting for connections to end...`);
    this.server.close(() => {
      logger.info(`Server closed successfully`);
    });
  }

  setListeners(): void {
    process.on('uncaughtExceptionMonitor', (error: Error, origin: string) => {
      logger.error(`Caught exception:\n${util.format(error)}`);
      logger.error(`Origin: ${origin}`);
    });
    process.on('unhandledRejection', (reason, promise) => {
      logger.warn(`Unhandled Rejection at:\n${util.format(promise)}`);
    });
  }
}
