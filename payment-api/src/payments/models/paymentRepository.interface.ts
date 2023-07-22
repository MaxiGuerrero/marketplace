import { Cart } from './cart';
import { Product } from './product';
import { Ticket } from './ticket';

interface IPaymentInterface {
  findCartByUserId(userId: string): Promise<Cart | null>;
  findProductByID(productId: string): Promise<Product | undefined>;
  checkout(cart: Cart): Promise<void>;
  saveTicket(ticket: Ticket): Promise<void>;
}

export { IPaymentInterface };
