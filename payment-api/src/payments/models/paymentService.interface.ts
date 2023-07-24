import BusinessError from '@shared/handler/businessError';
import { Ticket } from './ticket';

interface IPaymentService {
  checkout(
    userId: string,
    username: string,
    paymentMethod: string
  ): Promise<Ticket | BusinessError<'Insufficient stock' | 'Cart not found' | 'Cart without products'>>;
}

export { IPaymentService };
