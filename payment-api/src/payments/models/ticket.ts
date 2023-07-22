import { User } from '@shared/index';
import { ObjectId } from 'mongodb';
import { DetailPayment } from './detailPayment';

type Ticket = Pick<User, 'username'> & {
  _id: ObjectId;
  userId: string;
  products: DetailPayment[];
  transactionDate: Date;
  total: number;
  paymentMethod: string;
};

export { Ticket };
