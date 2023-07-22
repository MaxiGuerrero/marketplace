import { Server } from '@express/index';
import { config } from '@shared/index';
import PaymentRegisterRoutes from '@payments/infrastructure/routes';
import CartRegisterRoutes from '@cart-management/infrastructure/routes';

(async () => {
  const routes = [...(await PaymentRegisterRoutes()), ...(await CartRegisterRoutes())];
  const server = new Server(config.PORT, routes);
  server.start();
})();
