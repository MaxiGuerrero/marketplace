import { paymentService } from '@payments/services';
import PaymentController from './payment.controller';

const paymentController = new PaymentController(paymentService);

export { paymentController };
