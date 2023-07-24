import BusinessError from '@shared/handler/businessError';
import { Cart } from './cart';

interface ICartService {
  addProduct(
    productId: string,
    amount: number,
    userId: string
  ): Promise<BusinessError<'Product not found' | 'Insufficient stock'> | Cart>;

  removeProductOnCart(
    productId: string,
    userId: string
  ): Promise<BusinessError<'Product not found' | 'Cart not found'> | Cart>;
}

export { ICartService };
