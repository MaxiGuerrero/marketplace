import logger from '@shared/utils/logger';
import { Connector } from '@shared/databases';
import { ObjectId } from 'mongodb';
import { Cart, ICartRepository, Product } from '@cart-management/models';

class CartRepository implements ICartRepository {
  constructor(private readonly dbConnector: Connector) {}

  async findProductByID(productId: string): Promise<Product | undefined> {
    try {
      const product = await this.dbConnector
        .getCollection<Product>('product')
        .findOne({ _id: new ObjectId(productId) });
      if (!product) {
        return undefined;
      }
      return product;
    } catch (error) {
      logger.error('Error on findProductByID Repository: ', error);
      throw error;
    }
  }

  async updateCart(cart: Cart): Promise<Cart> {
    const result = await this.dbConnector.getCollection<Cart>('cart').updateOne(
      { _id: cart._id },
      {
        $set: {
          products: cart.products,
          updatedAt: new Date(),
          total: cart.total,
        },
      }
    );
    if (result.modifiedCount !== 1) {
      throw new Error('Error on update document cart from database');
    }
    const cartFound = await this.dbConnector.getCollection<Cart>('cart').findOne({ _id: cart._id });
    if (!cartFound) {
      throw new Error('Error on update document cart from database');
    }
    return cartFound;
  }

  async findCartByUserId(userId: string): Promise<Cart | null> {
    return this.dbConnector.getCollection<Cart>('cart').findOne({ userId, paid: false });
  }

  async createCart(userId: string): Promise<Cart> {
    const cartCreated = await this.dbConnector.getCollection<Cart>('cart').insertOne({
      _id: new ObjectId(),
      userId,
      products: [],
      createdAt: new Date(),
      updatedAt: new Date(),
      paid: false,
      total: 0,
    });
    const cart = await this.dbConnector.getCollection<Cart>('cart').findOne({ _id: cartCreated.insertedId });
    if (!cart) {
      throw new Error('Error on create document cart from database');
    }
    return cart;
  }
}

export default CartRepository;
