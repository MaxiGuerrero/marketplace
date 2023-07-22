import { Cart } from './cart';
import { Product } from './product';

interface ICartRepository {
  findProductByID(productId: string): Promise<Product | undefined>;
  updateCart(cart: Cart): Promise<Cart>;
  findCartByUserId(userId: string): Promise<Cart | null>;
  createCart(userId: string): Promise<Cart>;
}

export { ICartRepository };
