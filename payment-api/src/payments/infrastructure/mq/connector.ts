import zmq from 'zeromq';
import { IMqConnector } from '@payments/models';
import logger from '@shared/utils/logger';
import { config } from '@shared/index';

class MqConnector implements IMqConnector {
  private readonly socket: zmq.Socket;

  constructor() {
    const sock = zmq.socket('push');
    sock.connect(config.MQ_SERVER_URL);
    logger.info(`ZeroMQ connector has been connected to server: ${config.MQ_SERVER_URL}`);
    this.socket = sock;
  }

  async sendMessage<T extends object>(data: T): Promise<void> {
    const jsonParsed = JSON.stringify(data);
    logger.info(`Message sent ${jsonParsed}`);
    this.socket.send(jsonParsed);
  }
}

export { MqConnector };
