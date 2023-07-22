import { Router } from 'express';
import { authMiddleware } from '@shared/index';
import { paymentController } from '../controllers';

const router = Router();

router.post('/checkout', authMiddleware, paymentController.checkout);

export default router;
