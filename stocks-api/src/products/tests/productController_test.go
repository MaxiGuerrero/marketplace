package tests

import (
	"errors"
	"fmt"
	controllers "marketplace/stocks-api/src/products/infrastructure"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FakeProductService struct {
	mock.Mock
}

func (ps *FakeProductService) RegisterProduct(name string, description string, price float32, stock int){
	ps.Called(name,description,price,stock)
}

func (ps *FakeProductService) UpdateProduct(productId primitive.ObjectID,name string, description string, price float32) error {
	args := ps.Called(productId,name,description,price)
	return args.Error(0)
}

func (ps *FakeProductService) UpdateStock(productId primitive.ObjectID,stock int) error{
	args := ps.Called(productId,stock)
	return args.Error(0)
}

func TestProductRegisterController(t *testing.T){
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	c.Request().Header.SetContentType("application/json")
	fakeService := &FakeProductService{}
	controller := controllers.NewProductController(fakeService)
	t.Run("Product register successfully, return 200",func(t *testing.T){
		// Arrengue
		const name,description = "product-A","Is a product for everybody"
		const price float32 = 23.41
		const stock int = 50
		request := fmt.Sprintf(`{"name":"%v","description":"%v","price":%v,"stock":%v}`,name,description,price,stock)
		c.Request().SetBody([]byte(request))
		fakeService.On("RegisterProduct",name,description,price,stock).Once().Return(nil)
		// Act
		err := controller.RegisterProduct(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,`{"message":"Successful operation"}`,string(c.Response().Body()))
		utils.AssertEqual(t,200,c.Response().StatusCode())
	})
	t.Run("Send negative numbers, return 400",func(t *testing.T){
		// Arrengue
		const name,description = "product-A","Is a product for everybody"
		const price float32 = -23.41
		const stock int = -50
		request := fmt.Sprintf(`{"name":"%v","description":"%v","price":%v,"stock":%v}`,name,description,price,stock)
		c.Request().SetBody([]byte(request))
		fakeService.On("RegisterProduct",name,description,price,stock).Once().Return(nil)
		// Act
		err := controller.RegisterProduct(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
}

func TestUpdateProductController(t *testing.T){
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	c.Request().Header.SetContentType("application/json")
	fakeService := &FakeProductService{}
	controller := controllers.NewProductController(fakeService)
	t.Run("Product updated successfully, return 200",func(t *testing.T){
		// Arrengue
		const name,description = "product-A","Is a product for everybody"
		const price float32 = 23.41
		id := primitive.NewObjectID()
		request := fmt.Sprintf(`{"productId":"%v","name":"%v","description":"%v","price":%v}`,id.Hex(),name,description,price)
		c.Request().SetBody([]byte(request))
		fakeService.On("UpdateProduct",id,name,description,price).Once().Return(nil)
		// Act
		err := controller.UpdateProduct(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,`{"message":"Successful operation"}`,string(c.Response().Body()))
		utils.AssertEqual(t,200,c.Response().StatusCode())
	})
	t.Run("Bad values setted, return 400",func(t *testing.T){
		// Arrengue
		const name,description = 10,""
		const price float32 = 23.41
		id := 1234
		request := fmt.Sprintf(`{"productId":"%v","name":"%v","description":"%v","price":%v}`,id,name,description,price)
		c.Request().SetBody([]byte(request))
		fakeService.On("UpdateProduct",id,name,description,price).Once().Return(nil)
		// Act
		err := controller.UpdateProduct(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
	t.Run("Send negative values, return 400",func(t *testing.T){
		// Arrengue
		const name,description = "product-a","Is a product for everybody"
		const price float32 = -23.41
		id := primitive.NewObjectID()
		request := fmt.Sprintf(`{"productId":"%v","name":"%v","description":"%v","price":%v}`,id.Hex(),name,description,price)
		c.Request().SetBody([]byte(request))
		fakeService.On("UpdateProduct",id,name,description,price).Once().Return(nil)
		// Act
		err := controller.UpdateProduct(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
	t.Run("Product does not exists, return 400 and message",func(t *testing.T){
		// Arrengue
		const name,description = "product-a","Is a product for everybody"
		const price float32 = 23.41
		id := primitive.NewObjectID()
		request := fmt.Sprintf(`{"productId":"%v","name":"%v","description":"%v","price":%v}`,id.Hex(),name,description,price)
		c.Request().SetBody([]byte(request))
		fakeService.On("UpdateProduct",id,name,description,price).Once().Return(errors.New("product does not exists"))
		// Act
		err := controller.UpdateProduct(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,`{"message":"product does not exists"}`,string(c.Response().Body()))
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
}

func TestUpdateStockProductController(t *testing.T){
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	c.Request().Header.SetContentType("application/json")
	fakeService := &FakeProductService{}
	controller := controllers.NewProductController(fakeService)
	t.Run("Stock updated successfully, return 200",func(t *testing.T){
		// Arrengue
		const stock = 150
		id := primitive.NewObjectID()
		request := fmt.Sprintf(`{"productId":"%v","stock":%v}`,id.Hex(),stock)
		c.Request().SetBody([]byte(request))
		fakeService.On("UpdateStock",id,stock).Once().Return(nil)
		// Act
		err := controller.UpdateStock(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,`{"message":"Successful operation"}`,string(c.Response().Body()))
		utils.AssertEqual(t,200,c.Response().StatusCode())
	})
	t.Run("Send negative values, return 400",func(t *testing.T){
		// Arrengue
		const stock = -150
		id := primitive.NewObjectID()
		request := fmt.Sprintf(`{"productId":"%v","stock":%v}`,id.Hex(),stock)
		c.Request().SetBody([]byte(request))
		fakeService.On("UpdateStock",id,stock).Once().Return(nil)
		// Act
		err := controller.UpdateStock(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
	t.Run("Product does not exists, return 400 and message",func(t *testing.T){
		// Arrengue
		const stock = 150
		id := primitive.NewObjectID()
		request := fmt.Sprintf(`{"productId":"%v","stock":%v}`,id.Hex(),stock)
		c.Request().SetBody([]byte(request))
		fakeService.On("UpdateStock",id,stock).Once().Return(errors.New("product does not exists"))
		// Act
		err := controller.UpdateStock(c)
		// Assert
		require.NoError(t,err,"Controller does not return an error")
		utils.AssertEqual(t,`{"message":"product does not exists"}`,string(c.Response().Body()))
		utils.AssertEqual(t,400,c.Response().StatusCode())
	})
}