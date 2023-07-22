import { ObjectId } from 'mongodb';

type Product = {
  _id: ObjectId;
  name: string;
  description: string;
  stock: number;
  price: number;
  createdat: Date;
  updatedat: Date;
};

export { Product };
