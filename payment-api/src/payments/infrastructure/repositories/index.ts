import { connector } from '@shared/databases';
import { PaymentRepository } from './payment.repository';

const paymentRepository = new PaymentRepository(connector);

export { paymentRepository };
