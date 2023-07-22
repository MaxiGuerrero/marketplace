import { connector } from '@shared/databases';
import CartRepository from './cart.repository';

const cartRepository = new CartRepository(connector);

export default cartRepository;
