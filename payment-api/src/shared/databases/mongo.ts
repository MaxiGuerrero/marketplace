import logger from '@shared/utils/logger';
import { Db, Document, MongoClient } from 'mongodb';

class Connector {
  db: Db;

  constructor(url: string) {
    this.initialize(url);
  }

  async initialize(url: string) {
    const client = new MongoClient(url);
    await client.connect();
    this.db = client.db('marketplace');
    logger.info(`Database connected`);
  }

  getCollection<T extends Document>(collection: string) {
    return this.db.collection<T>(collection);
  }
}

export default Connector;
