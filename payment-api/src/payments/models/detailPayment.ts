import { Product } from './product';

type DetailPayment = Omit<Product, '_id' | 'createdat' | 'updatedat' | 'stock'> & {
  amount: number;
};

export { DetailPayment };
