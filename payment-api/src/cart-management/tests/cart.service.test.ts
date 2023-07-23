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

const product: Product = {
  _id: new ObjectId(),
  name: 'product-a',
  description: 'is a good product',
  createdat: new Date(),
  price: 10,
  stock: 1,
  updatedat: new Date(),
};

describe('Cart services', () => {
  describe('#addProductToCart', () => {
    test('Product not found, return error message', async () => {
      // Arrange
      repository.findProductByID.mockReturnValueOnce(undefined);
      // Act
      const result = await service.addProduct(productId, amount, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(result).toEqual('Product not found');
    });
    test('Product does not have stock, return error message', async () => {
      // Arrange
      const amountNotStock = 3;
      repository.findProductByID.mockReturnValueOnce(product);
      // Act
      const result = await service.addProduct(productId, amountNotStock, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(result).toEqual('Insufficient stock');
    });
    test('Product added successfully, case user did not have a cart register before, return error cart', async () => {
      // Arrange
      const cartExpected: Cart = {
        _id: new ObjectId(),
        createdAt: new Date(),
        paid: false,
        products: [
          {
            amount,
            productId: product._id,
          },
        ],
        total: amount * product.price,
        updatedAt: new Date(),
        userId,
      };
      repository.findProductByID.mockReturnValue(product);
      repository.findCartByUserId.mockReturnValueOnce(undefined);
      repository.createCart.mockReturnValueOnce(cartExpected);
      repository.updateCart.mockReturnValueOnce(cartExpected);
      // Act
      const result = await service.addProduct(productId, amount, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(repository.findCartByUserId).toHaveBeenCalledWith(userId);
      expect(repository.createCart).toHaveBeenCalledWith(userId);
      expect(repository.updateCart).toHaveBeenCalledWith(cartExpected);
      expect(result).toEqual(cartExpected);
    });
    test('Product added successfully, case user have a cart register before, return error cart', async () => {
      // Arrange
      const cartExpected: Cart = {
        _id: new ObjectId(),
        createdAt: new Date(),
        paid: false,
        products: [
          {
            amount,
            productId: product._id,
          },
        ],
        total: amount * product.price,
        updatedAt: new Date(),
        userId,
      };
      repository.findProductByID.mockReturnValue(product);
      repository.findCartByUserId.mockReturnValueOnce(cartExpected);
      repository.updateCart.mockReturnValueOnce(cartExpected);
      // Act
      const result = await service.addProduct(productId, amount, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(repository.findCartByUserId).toHaveBeenCalledWith(userId);
      expect(repository.createCart).not.toHaveBeenCalled();
      expect(repository.updateCart).toHaveBeenCalledWith(cartExpected);
      expect(result).toEqual(cartExpected);
    });
  });

  describe('#removeProductToCart', () => {
    test('product not found, return message', async () => {
      // Arrenge
      repository.findProductByID.mockReturnValueOnce(undefined);
      // Act
      const result = await service.removeProductOnCart(productId, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(result).toEqual('Product not found');
    });
    test('cart not found, return message', async () => {
      // Arrenge
      repository.findProductByID.mockReturnValueOnce(product);
      repository.findCartByUserId.mockReturnValueOnce(undefined);
      // Act
      const result = await service.removeProductOnCart(productId, userId);
      // Assert
      expect(repository.findProductByID).toHaveBeenCalledWith(productId);
      expect(repository.findCartByUserId).toHaveBeenCalledWith(userId);
      expect(result).toEqual('Cart not found');
    });
    test('remove product succesfully, return cart updated', async () => {
      // Arrenge
      const cart: Cart = {
        _id: new ObjectId(),
        createdAt: new Date(),
        paid: false,
        products: [
          {
            amount,
            productId: product._id,
          },
        ],
        total: amount * product.price,
        updatedAt: new Date(),
        userId,
      };
      const cartUpdated = { ...cart, total: 0, products: [] };
      repository.findProductByID.mockReturnValueOnce(product);
      repository.findCartByUserId.mockReturnValueOnce(cart);
      repository.updateCart.mockReturnValueOnce(cartUpdated);
      // Act
      const result = await service.removeProductOnCart(productId, userId);
      // Act
      expect(result).toEqual(cartUpdated);
    });
  });
});
