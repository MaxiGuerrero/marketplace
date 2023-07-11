package tests

import (
	"marketplace/stocks-api/src/products/models"
	services "marketplace/stocks-api/src/products/service"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FakeProductRepository struct {
	mock.Mock
}

func (pr *FakeProductRepository) RegisterProduct(name string, description string, price float32, stock int) {
	pr.Called(name,description,price,stock)
}

func (pr *FakeProductRepository) UpdateProduct(productId primitive.ObjectID,name string, description string, price float32) {
	pr.Called(productId,name,description,price)
}

func (pr *FakeProductRepository) GetProductById(productId primitive.ObjectID) *models.Product {
	args := pr.Called(productId)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*models.Product)
}

func (pr *FakeProductRepository) UpdateStock(productId primitive.ObjectID, stock int){
	pr.Called(productId,stock)
}

func TestRegisterProductService(t *testing.T){
	fakeRepository := &FakeProductRepository{}
	service := services.NewProductService(fakeRepository)
	t.Run("Product register successfully",func(t *testing.T){
		// Arrenge
		const name,description = "product-A","Is a product for everybody"
		const price float32 = 23.41
		const stock int = 50
		fakeRepository.On("RegisterProduct",name,description,price,stock).Once()
		// Act
		service.RegisterProduct(name,description,price,stock)
		// Assert
		require.True(t,fakeRepository.AssertCalled(t,"RegisterProduct",name,description,price,stock))
	})
}

func TestUpdateProductService(t *testing.T){
	fakeRepository := &FakeProductRepository{}
	service := services.NewProductService(fakeRepository)
	t.Run("Product update successfully",func(t *testing.T){
		// Arrenge
		const name,description = "product-A","Is a product for everybody"
		const price float32 = 23.41
		id := primitive.NewObjectID()
		product := &models.Product{
			ID: id,
			Name: name,
			Description: description,
			Price: price,
			Stock: 50,
		}
		fakeRepository.On("GetProductById",id).Once().Return(product)
		fakeRepository.On("UpdateProduct",id,name,description,price).Once()
		// Act
		service.UpdateProduct(id,name,description,price)
		// Assert
		require.True(t,fakeRepository.AssertCalled(t,"UpdateProduct",id,name,description,price))
	})
	t.Run("Product does not exists",func(t *testing.T){
		// Arrenge
		const name,description = "product-A","Is a product for everybody"
		const price float32 = 23.41
		id := primitive.NewObjectID()
		fakeRepository.On("GetProductById",id).Once().Return(nil)
		fakeRepository.On("UpdateProduct",id,name,description,price).Once()
		// Act
		err := service.UpdateProduct(id,name,description,price)
		// Assert
		require.Error(t,err)
		require.True(t,fakeRepository.AssertNotCalled(t,"UpdateProduct",id,name,description,price))
	})
}

func TestUpdateStockProductService(t *testing.T){
	fakeRepository := &FakeProductRepository{}
	service := services.NewProductService(fakeRepository)
	t.Run("Stock update successfully",func(t *testing.T){
		// Arrenge
		const stock int = 150
		id := primitive.NewObjectID()
		product := &models.Product{
			ID: id,
			Name: "product",
			Description: "description",
			Price: 5.51,
			Stock: 50,
		}
		fakeRepository.On("GetProductById",id).Once().Return(product)
		fakeRepository.On("UpdateStock",id,stock).Once()
		// Act
		service.UpdateStock(id,stock)
		// Assert
		require.True(t,fakeRepository.AssertCalled(t,"UpdateStock",id,stock))
	})
}