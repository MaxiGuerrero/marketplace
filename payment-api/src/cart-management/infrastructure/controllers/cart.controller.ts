import { NextFunction, Request, Response } from 'express';
import { AddProductRequest, RemoveProductRequest } from '@cart-management/models';
import { User } from '@shared/models/user';
import { ICartService } from '@cart-management/models/cartService.interface';

class CartController {
  constructor(private readonly cartService: ICartService) {}

  addProduct = async (req: Request, res: Response, next: NextFunction) => {
    try {
      const body = req.body as AddProductRequest;
      const user = req.body.user as User;
      const result = await this.cartService.addProduct(body.productId, body.amount, user.userId);
      if (typeof result === 'string') {
        return res.status(400).json({ error: result });
      }
      return res.json(result);
    } catch (err) {
      return next(err);
    }
  };

  removeProduct = async (req: Request, res: Response, next: NextFunction) => {
    try {
      const body = req.body as RemoveProductRequest;
      const user = req.body.user as User;
      const result = await this.cartService.removeProductOnCart(body.productId, user.userId);
      if (typeof result === 'string') {
        return res.status(400).json({ error: result });
      }
      return res.json(result);
    } catch (error) {
      return next(error);
    }
  };
}

export default CartController;
