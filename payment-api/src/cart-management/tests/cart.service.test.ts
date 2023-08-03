import { Cart, Product } from '@cart-management/models';
import CartService from '@cart-management/services/cart.service';
import { ObjectId } from 'mongodb';

const repository = {
  findCartByUserId: jest.fn(),
  findProductByID: jest.fn(),
  updateCart: jest.fn(),
  createCart: jest.fn(),
};

const service = new CartService(repository);

const productId = '1';
const amount = 1;
const userId = '1';

const objectId = new ObjectId();
const date = new Date();

const product: Product = {
  _id: objectId,
  name: 'product-a',
  description: 'is a good product',
  createdat: date,
  price: 10,
  stock: 5,
  updatedat: date,
};

const cart: Cart = {
  _id: objectId,
  createdAt: date,
  paid: false,
  products: [
    {
      amount,
      productId: product._id,
    },
  ],
  total: amount * product.price,
  updatedAt: date,
  userId,
};

describe('Cart services', () => {
  describe('#addProductToCart', () => {
    test('Product not found, return error message', async () => {
      // Arrange
      repository.findProductByID.mockReturnValue(undefined);
      // Act
      const result = await service.addProduct(productId, amount, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(result).toEqual('Product not found');
    });
    test('Product does not have stock, return error message', async () => {
      // Arrange
      const amountNotStock = 100;
      repository.findProductByID.mockReturnValue(product);
      // Act
      const result = await service.addProduct(productId, amountNotStock, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(result).toEqual('Insufficient stock');
    });
    test('Product added successfully, case user did not have a cart register before, return cart', async () => {
      // Arrange
      repository.findProductByID.mockReturnValue(product);
      repository.findCartByUserId.mockReturnValue(undefined);
      repository.createCart.mockReturnValue(cart);
      repository.updateCart.mockReturnValue(cart);
      // Act
      const result = await service.addProduct(productId, amount, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(repository.findCartByUserId).toHaveBeenCalledWith(userId);
      expect(repository.createCart).toHaveBeenCalledWith(userId);
      expect(repository.updateCart).toHaveBeenCalledWith(cart);
      expect(result).toEqual(cart);
    });

    test('Product added successfully, case user have a cart register before, return cart', async () => {
      // Arrange
      repository.findProductByID.mockReturnValue(product);
      repository.findCartByUserId.mockReturnValue(cart);
      repository.updateCart.mockReturnValue(cart);
      // Act
      const result = await service.addProduct(productId, amount, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(repository.findCartByUserId).toHaveBeenCalledWith(userId);
      expect(repository.createCart).not.toHaveBeenCalled();
      expect(repository.updateCart).toHaveBeenCalledWith(cart);
      expect(result).toEqual(cart);
    });

    test('Product added again, but the sum of the amount with the cart is major than stock, return error', async () => {
      // Arrange
      repository.findProductByID.mockReturnValue(product);
      repository.findCartByUserId.mockReturnValue(cart);
      // Act
      const result = await service.addProduct(productId, 5, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(repository.findCartByUserId).toHaveBeenCalledWith(userId);
      expect(repository.createCart).not.toHaveBeenCalled();
      expect(repository.updateCart).not.toHaveBeenCalled();
      expect(result).toEqual('Insufficient stock');
    });
  });

  describe('#removeProductToCart', () => {
    test('product not found, return message', async () => {
      // Arrenge
      repository.findProductByID.mockReturnValue(undefined);
      // Act
      const result = await service.removeProductOnCart(productId, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(result).toEqual('Product not found');
    });
    test('cart not found, return message', async () => {
      // Arrenge
      repository.findProductByID.mockReturnValue(product);
      repository.findCartByUserId.mockReturnValue(undefined);
      // Act
      const result = await service.removeProductOnCart(productId, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(repository.findCartByUserId).toHaveBeenCalledWith(userId);
      expect(result).toEqual('Cart not found');
    });
    test('remove product succesfully, return cart updated', async () => {
      // Arrenge
      const cartUpdated = { ...cart, total: 0, products: [] };
      repository.findProductByID.mockReturnValue(product);
      repository.findCartByUserId.mockReturnValue(cart);
      repository.updateCart.mockReturnValue(cartUpdated);
      // Act
      const result = await service.removeProductOnCart(productId, userId);
      // Act
      expect(result).toEqual(cartUpdated);
    });
  });
});
