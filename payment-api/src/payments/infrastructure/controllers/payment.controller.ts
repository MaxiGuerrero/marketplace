import { NextFunction, Request, Response } from 'express';
import PaymentService from '@payments/services/payment.service';
import { CheckoutRequest } from '@payments/models';
import { User } from '@shared/index';

class PaymentController {
  constructor(private readonly paymentService: PaymentService) {}

  checkout = async (req: Request, res: Response, next: NextFunction) => {
    try {
      const body = req.body as CheckoutRequest;
      const user = req.body.user as User;
      const result = await this.paymentService.checkout(user.userId, user.username, body.paymentMethod);
      if (typeof result === 'string') {
        return res.status(400).json({ error: result });
      }
      return res.json(result);
    } catch (err) {
      return next(err);
    }
  };
}

export default PaymentController;
