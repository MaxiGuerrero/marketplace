import { Cart, IPaymentInterface, Product, Ticket } from '@payments/models';
import { Connector } from '@shared/databases';
import logger from '@shared/utils/logger';
import { ObjectId } from 'mongodb';

class PaymentRepository implements IPaymentInterface {
  constructor(private readonly dbConnector: Connector) {}

  async findCartByUserId(userId: string): Promise<Cart | null> {
    return this.dbConnector.getCollection<Cart>('cart').findOne({ userId, paid: false });
  }

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

  async checkout(cart: Cart): Promise<void> {
    const result = await this.dbConnector.getCollection<Cart>('cart').updateOne(
      { _id: cart._id },
      {
        $set: {
          paid: true,
          updatedAt: new Date(),
        },
      }
    );
    if (result.modifiedCount !== 1) {
      throw new Error('Error on update document cart from database');
    }
  }

  async saveTicket(ticket: Ticket): Promise<void> {
    await this.dbConnector.getCollection<Ticket>('ticket').insertOne(ticket);
  }
}

export { PaymentRepository };
