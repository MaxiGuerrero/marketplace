import { cartService } from '@cart-management/services';
import CartController from './cart.controller';

const cartController = new CartController(cartService);

export { cartController };
