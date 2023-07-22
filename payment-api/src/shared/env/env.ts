import dotenv from 'dotenv';

process.env.NODE_ENV = process.env.NODE_ENV || 'development';

if (process.env.NODE_ENV !== 'production') {
  const envFound = dotenv.config({ override: true });
  if (envFound.error) {
    throw new Error("Couldn't find .env file");
  }
}

export default {
  NODE_ENV: process.env.NODE_ENV,
  PORT: Number(process.env.PORT || 8080),
  DOCS_ENDPOINT: '/docs',
  DIR_SWAGGER: './src/shared/docs/swagger.yml',
  DB_CONNECTION: process.env.DB_CONNECTION || 'localhost',
  SECURITY_API_URL: process.env.SECURITY_API_URL || 'http://localhost:8080/token/validate',
};
