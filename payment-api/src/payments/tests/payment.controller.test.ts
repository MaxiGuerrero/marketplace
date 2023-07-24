import PaymentController from '@payments/infrastructure/controllers/payment.controller';
import { Ticket } from '@payments/models';
import { Request, Response } from 'express';
import { ObjectId } from 'mongodb';

const paymentService = {
  checkout: jest.fn(),
};

const req = {
  body: {
    user: {
      userId: '1',
      username: 'test',
    },
    paymentMethod: 'cash',
  },
} as Request;

const res = {} as Response;
res.status = jest.fn().mockReturnValue(res);
res.json = jest.fn().mockReturnValue(res);

const next = jest.fn();

const paymentController = new PaymentController(paymentService);

describe('Payment controller', () => {
  describe('#checkout', () => {
    test('Cart not found, return error 400', async () => {
      // Arrenge
      paymentService.checkout.mockReturnValueOnce('Cart not found');
      // Act
      await paymentController.checkout(req, res, next);
      // Assert
      expect(res.status).toHaveBeenCalledWith(400);
      expect(res.json).toHaveBeenCalledWith({ error: 'Cart not found' });
    });
    test('Insuficient stock, return error 400', async () => {
      // Arrenge
      paymentService.checkout.mockReturnValueOnce('Insufficient stock');
      // Act
      await paymentController.checkout(req, res, next);
      // Assert
      expect(res.status).toHaveBeenCalledWith(400);
      expect(res.json).toHaveBeenCalledWith({ error: 'Insufficient stock' });
    });
    test('Cart without products, return error 400', async () => {
      // Arrenge
      paymentService.checkout.mockReturnValueOnce('Cart without products');
      // Act
      await paymentController.checkout(req, res, next);
      // Assert
      expect(res.status).toHaveBeenCalledWith(400);
      expect(res.json).toHaveBeenCalledWith({ error: 'Cart without products' });
    });
    test('Payment done sucessfully, return 200', async () => {
      // Arrenge
      const ticket: Ticket = {
        _id: new ObjectId(),
        paymentMethod: req.body.paymentMethod,
        products: [
          {
            amount: 1,
            description: 'product-A',
            name: 'product-A',
            price: 10,
          },
        ],
        total: 10,
        transactionDate: new Date(),
        userId: req.body.user.userId,
        username: req.body.user.username,
      };
      paymentService.checkout.mockReturnValueOnce(ticket);
      // Act
      await paymentController.checkout(req, res, next);
      // Assert
      expect(res.json).toHaveBeenCalledWith(ticket);
    });
  });
});
