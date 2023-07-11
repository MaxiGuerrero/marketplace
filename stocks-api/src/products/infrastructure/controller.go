package infrastructure

import (
	"marketplace/stocks-api/src/products/models"
	response "marketplace/stocks-api/src/shared"
	"marketplace/stocks-api/src/shared/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	ProductService models.IProductService
}

func NewProductController(productService models.IProductService) *ProductController{
	return &ProductController{productService}
}

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