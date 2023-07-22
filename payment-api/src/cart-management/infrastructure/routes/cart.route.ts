import { Router } from 'express';
import { authMiddleware } from '@shared/index';
import { cartController } from '../controllers';

const router = Router();

router.post('/products/add', authMiddleware, cartController.addProduct);
router.post('/products/remove', authMiddleware, cartController.removeProduct);

export default router;
