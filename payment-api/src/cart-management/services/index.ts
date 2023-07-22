import cartRepository from '@cart-management/infrastructure/repositories';
import CartService from './cart.service';

const cartService = new CartService(cartRepository);

export { cartService };
