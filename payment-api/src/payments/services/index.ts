import { paymentRepository } from '@payments/infrastructure/repositories';
import PaymentService from './payment.service';

const paymentService = new PaymentService(paymentRepository);

export { paymentService };
