import { Cart, ICartRepository, ProductOnCart } from '@cart-management/models';
import BusinessError from '@shared/handler/businessError';
import logger from '@shared/utils/logger';

class CartService {
  constructor(private readonly repository: ICartRepository) {}

  async addProduct(
    productId: string,
    amount: number,
    userId: string
  ): Promise<BusinessError<'Product not found' | 'Insufficient stock'> | Cart> {
    try {
      const product = await this.repository.findProductByID(productId);
      if (!product) {
        return 'Product not found';
      }
      if (product.stock < amount) {
        return 'Insufficient stock';
      }
      let cart = await this.repository.findCartByUserId(userId);
      if (!cart) {
        cart = await this.repository.createCart(userId);
      }
      const productExistCart = cart.products.find(
        (productCart) => productCart.productId.toHexString() === product._id.toHexString()
      );
      if (productExistCart) {
        productExistCart.amount += amount;
      } else {
        const newArticle: ProductOnCart = {
          productId: product._id,
          amount,
        };
        cart.products.push(newArticle);
      }
      cart.total = await this._calucateTotal(cart.products);
      return this.repository.updateCart(cart);
    } catch (error) {
      logger.error(error);
      throw error;
    }
  }

  async removeProductOnCart(
    productId: string,
    userId: string
  ): Promise<BusinessError<'Product not found' | 'Cart not found'> | Cart> {
    try {
      const product = await this.repository.findProductByID(productId);
      if (!product) {
        return 'Product not found';
      }
      const cart = await this.repository.findCartByUserId(userId);
      if (!cart) {
        return 'Cart not found';
      }
      const index = cart.products.findIndex(
        (productCart) => productCart.productId.toHexString() === product._id.toHexString()
      );
      cart.products.splice(index, 1);
      cart.total = await this._calucateTotal(cart.products);
      return this.repository.updateCart(cart);
    } catch (error) {
      logger.error(error);
      throw error;
    }
  }

  private async _calucateTotal(products: ProductOnCart[]): Promise<number> {
    try {
      const accumulator: number[] = await Promise.all(
        products.map(async (productOnCart) => {
          const product = await this.repository.findProductByID(productOnCart.productId.toHexString());
          if (!product) {
            throw new Error('Inconsistent data on database');
          }
          return productOnCart.amount * product.price;
        })
      );
      return accumulator.reduce((acc, value) => acc + value, 0);
    } catch (error) {
      logger.error('Error on calculate total: ', error);
      throw error;
    }
  }
}

export default CartService;
