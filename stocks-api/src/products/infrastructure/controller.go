package infrastructure

import (
	"marketplace/stocks-api/src/products/models"
	response "marketplace/stocks-api/src/shared"
	"marketplace/stocks-api/src/shared/utils"

	"github.com/gofiber/fiber/v2"
)

// ProductController is a controller that is responsable to manage request and responses.
type ProductController struct {
	ProductService models.IProductService
}

// Create an instance of the ProductController with injection dependencies.
func NewProductController(productService models.IProductService) *ProductController{
	return &ProductController{productService}
}

// Manage request and responses about register an product in the system.
func (p *ProductController) RegisterProduct(c *fiber.Ctx) error{
	req := models.RegisterProductRequest{}
	if parseError := c.BodyParser(&req); parseError!=nil{
		return c.Status(400).JSON(response.BadRequest(parseError.Error()))
	}
	if badSchema := utils.ValidateSchema(&req); badSchema != nil{
		return c.Status(400).JSON(response.BadRequest(badSchema.Error()))
	}
	p.ProductService.RegisterProduct(req.Name,req.Description,req.Price,req.Stock)
	return c.Status(200).JSON(response.OK())
}

// Manage request and responses about update an product in the system.
func (p *ProductController) UpdateProduct(c *fiber.Ctx) error{
	req := models.UpdateProductRequest{}
	if parseError := c.BodyParser(&req); parseError!=nil{
		return c.Status(400).JSON(response.BadRequest(parseError.Error()))
	}
	if badSchema := utils.ValidateSchema(&req); badSchema != nil{
		return c.Status(400).JSON(response.BadRequest(badSchema.Error()))
	}
	err := p.ProductService.UpdateProduct(req.Id,req.Name,req.Description,req.Price)
	if err != nil {
		return c.Status(400).JSON(response.BadRequest(err.Error()))
	}
	return c.Status(200).JSON(response.OK())
}

// Manage request and responses about update the stock of an product in the system.
func (p *ProductController) UpdateStock(c *fiber.Ctx) error{
	req := models.UpdateStockRequest{}
	if parseError := c.BodyParser(&req); parseError!=nil{
		return c.Status(400).JSON(response.BadRequest(parseError.Error()))
	}
	if badSchema := utils.ValidateSchema(&req); badSchema != nil{
		return c.Status(400).JSON(response.BadRequest(badSchema.Error()))
	}
	err := p.ProductService.UpdateStock(req.Id,req.Stock)
	if err != nil {
		return c.Status(400).JSON(response.BadRequest(err.Error()))
	}
	return c.Status(200).JSON(response.OK())
}

// Manage request and responses about get list of products in the system.
func (p *ProductController) GetProducts(c *fiber.Ctx) error {
	return c.Status(200).JSON(p.ProductService.GetAll())
}