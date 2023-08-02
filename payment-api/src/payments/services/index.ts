import { paymentRepository } from '@payments/infrastructure/repositories';
import { MqConnector } from '@payments/infrastructure/mq/connector';
import PaymentService from './payment.service';

const connector = new MqConnector();
const paymentService = new PaymentService(paymentRepository, connector);

export { paymentService };
