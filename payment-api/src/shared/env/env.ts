import dotenv from 'dotenv';

process.env.NODE_ENV = process.env.NODE_ENV || 'development';

dotenv.config();

export default {
  NODE_ENV: process.env.NODE_ENV,
  PORT: Number(process.env.PORT || 8080),
  DOCS_ENDPOINT: '/docs',
  DIR_SWAGGER: process.env.DIR_SWAGGER || './src/shared/docs/swagger.yml',
  DB_CONNECTION: process.env.DB_CONNECTION || 'localhost',
  SECURITY_API_URL: process.env.SECURITY_API_URL || 'http://localhost:8080',
  PREFIX_URL: process.env.PREFIX_URL || '/',
  MQ_SERVER_URL: process.env.MQ_SERVER_URL || 'tcp://127.0.0.1:9000',
};
