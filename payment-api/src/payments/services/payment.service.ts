import { IMqConnector, IPaymentInterface, ProductOnCart, Ticket } from '@payments/models';
import BusinessError from '@shared/handler/businessError';
import { ObjectId } from 'mongodb';

class PaymentService {
  constructor(private readonly repository: IPaymentInterface, private readonly mqConnector: IMqConnector) {}

  async checkout(
    userId: string,
    username: string,
    paymentMethod: string
  ): Promise<Ticket | BusinessError<'Insufficient stock' | 'Cart not found' | 'Cart without products'>> {
    const cart = await this.repository.findCartByUserId(userId);
    if (!cart) {
      return 'Cart not found';
    }
    if (cart.products.length === 0) {
      return 'Cart without products';
    }
    const products = await Promise.all(
      cart.products.map(async (product) => {
        const productFound = await this.repository.findProductByID(product.productId.toHexString());
        if (!productFound) {
          throw new Error('Inconsistent Data');
        }
        return productFound;
      })
    );
    let noStock = false;
    products.forEach((product, index) => {
      noStock = product.stock < cart.products[index].amount;
    });
    if (noStock) {
      return 'Insufficient stock';
    }
    await this.repository.checkout(cart);
    const ticket: Ticket = {
      _id: new ObjectId(),
      userId,
      paymentMethod,
      transactionDate: new Date(),
      total: cart.total,
      products: products.map((product, index) => ({
        name: product.name,
        description: product.description,
        price: product.price,
        amount: cart.products[index].amount,
      })),
      username,
    };
    await this.repository.saveTicket(ticket);
    // Call asynchronous to stock-api about products that has been sold
    cart.products.forEach((product) => this.mqConnector.sendMessage<ProductOnCart>(product));
    return ticket;
  }
}

export default PaymentService;
