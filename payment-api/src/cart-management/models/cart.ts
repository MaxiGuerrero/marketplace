import { ObjectId } from 'mongodb';

type ProductOnCart = {
  productId: ObjectId;
  amount: number;
};

type Cart = {
  _id: ObjectId;
  userId: string;
  products: ProductOnCart[];
  createdAt: Date;
  updatedAt: Date;
  total: number;
  paid: boolean;
};

export { Cart, ProductOnCart };
