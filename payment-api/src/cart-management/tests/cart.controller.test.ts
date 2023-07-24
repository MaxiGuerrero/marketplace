import CartController from '@cart-management/infrastructure/controllers/cart.controller';
import { Cart } from '@cart-management/models';
import { Request, Response } from 'express';
import { ObjectId } from 'mongodb';

const cartService = {
  addProduct: jest.fn(),
  removeProductOnCart: jest.fn(),
};

const req = {
  body: {
    user: {
      userId: '1',
      username: 'test',
    },
    productId: '1',
    amount: 1,
  },
} as Request;

const res = {} as Response;
res.status = jest.fn().mockReturnValue(res);
res.json = jest.fn().mockReturnValue(res);

const next = jest.fn();

const cartController = new CartController(cartService);

describe('Cart controller', () => {
  describe('#addProduct', () => {
    test('Product not found, return error 400', async () => {
      // Arrenge
      cartService.addProduct.mockReturnValueOnce('Product not found');
      // Act
      await cartController.addProduct(req, res, next);
      // Assert
      expect(res.status).toHaveBeenCalledWith(400);
      expect(res.json).toHaveBeenCalledWith({ error: 'Product not found' });
    });
    test('Insuficient stock, return error 400', async () => {
      // Arrenge
      cartService.addProduct.mockReturnValueOnce('Insufficient stock');
      // Act
      await cartController.addProduct(req, res, next);
      // Assert
      expect(res.status).toHaveBeenCalledWith(400);
      expect(res.json).toHaveBeenCalledWith({ error: 'Insufficient stock' });
    });
    test('Product added sucessfully, return 200', async () => {
      // Arrenge
      const cartExpected: Cart = {
        _id: new ObjectId(),
        createdAt: new Date(),
        paid: false,
        products: [
          {
            amount: req.body.amount,
            productId: req.body.productId,
          },
        ],
        total: 150,
        updatedAt: new Date(),
        userId: req.body.user.userId,
      };
      cartService.addProduct.mockReturnValueOnce(cartExpected);
      // Act
      await cartController.addProduct(req, res, next);
      // Assert
      expect(res.json).toHaveBeenCalledWith(cartExpected);
    });
  });
  describe('#removeProduct', () => {
    test('Product not found, return error 400', async () => {
      // Arrenge
      cartService.removeProductOnCart.mockReturnValueOnce('Product not found');
      // Act
      await cartController.removeProduct(req, res, next);
      // Assert
      expect(res.status).toHaveBeenCalledWith(400);
      expect(res.json).toHaveBeenCalledWith({ error: 'Product not found' });
    });
    const cartExpected: Cart = {
      _id: new ObjectId(),
      createdAt: new Date(),
      paid: false,
      products: [
        {
          amount: req.body.amount,
          productId: req.body.productId,
        },
      ],
      total: 150,
      updatedAt: new Date(),
      userId: req.body.user.userId,
    };
    test('Product removed from the cart sucessfully, return cart updated', async () => {
      // Arrenge
      cartService.removeProductOnCart.mockReturnValueOnce(cartExpected);
      // Act
      await cartController.removeProduct(req, res, next);
      // Assert
      expect(res.json).toHaveBeenCalledWith(cartExpected);
    });
  });
});
