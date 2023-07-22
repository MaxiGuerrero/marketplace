import { config } from '@shared/index';
import Connector from './mongo';

const connector = new Connector(config.DB_CONNECTION);

export { connector, Connector };
