import { Cart, Product } from '@cart-management/models';
import PaymentService from '@payments/services/payment.service';
import { ObjectId } from 'mongodb';

const repository = {
  findCartByUserId: jest.fn(),
  findProductByID: jest.fn(),
  checkout: jest.fn(),
  saveTicket: jest.fn(),
};

const connector = {
  sendMessage: jest.fn(),
};

const service = new PaymentService(repository, connector);

const userId = '1';
const username = 'user';
const paymentMethod = 'cash';

const product: Product = {
  _id: new ObjectId(),
  name: 'product-a',
  description: 'is a good product',
  createdat: new Date(),
  price: 10,
  stock: 1,
  updatedat: new Date(),
};

describe('Payment service', () => {
  describe('#checkout', () => {
    test('Cart not found, return error', async () => {
      // Arrange
      repository.findCartByUserId.mockReturnValueOnce(undefined);
      // Act
      const result = await service.checkout(userId, username, paymentMethod);
      // Assert
      expect(result).toEqual('Cart not found');
    });
    test('Cart not found, return error', async () => {
      // Arrange
      const cart: Cart = {
        _id: new ObjectId(),
        createdAt: new Date(),
        paid: false,
        products: [],
        total: 50,
        updatedAt: new Date(),
        userId,
      };
      repository.findCartByUserId.mockReturnValueOnce(cart);
      // Act
      const result = await service.checkout(userId, username, paymentMethod);
      // Assert
      expect(result).toEqual('Cart without products');
    });
    test('Insufficient stock, return error', async () => {
      // Arrange
      const cart: Cart = {
        _id: new ObjectId(),
        createdAt: new Date(),
        paid: false,
        products: [
          {
            amount: 2,
            productId: product._id,
          },
        ],
        total: 50,
        updatedAt: new Date(),
        userId,
      };
      repository.findCartByUserId.mockReturnValueOnce(cart);
      repository.findProductByID.mockReturnValueOnce(product);
      // Act
      const result = await service.checkout(userId, username, paymentMethod);
      // Assert
      expect(result).toEqual('Insufficient stock');
    });
    test('Payment done, return ticket', async () => {
      // Arrange
      const cart: Cart = {
        _id: new ObjectId(),
        createdAt: new Date(),
        paid: false,
        products: [
          {
            amount: 1,
            productId: product._id,
          },
        ],
        total: 50,
        updatedAt: new Date(),
        userId,
      };
      repository.findCartByUserId.mockReturnValueOnce(cart);
      repository.findProductByID.mockReturnValueOnce(product);
      // Act
      const result = await service.checkout(userId, username, paymentMethod);
      // Assert
      expect(result).not.toBeUndefined();
      expect(repository.findCartByUserId).toHaveBeenCalledWith(userId);
      expect(repository.findProductByID).toHaveBeenCalledWith(product._id.toHexString());
      expect(repository.checkout).toHaveBeenCalledWith(cart);
      expect(repository.saveTicket).toHaveBeenCalled();
      expect(connector.sendMessage).toHaveBeenCalledWith(cart.products);
    });
  });
});
