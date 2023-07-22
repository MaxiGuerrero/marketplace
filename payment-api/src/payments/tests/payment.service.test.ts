import TestService from '@payments/services/payment.service';

const service = new TestService();

describe('Test service', () => {
  describe('#testEndpoint', () => {
    test('Get message', async () => {
      // Arrange
      const response = { status: 200, message: 'Successful operation' };
      // Act
      const result = await service.testEndpoint();
      // Assert
      expect(result).toEqual(response);
    });
  });
});
