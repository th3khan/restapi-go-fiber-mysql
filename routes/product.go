package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/restapi-go-fiber-mysql/database"
	"github.com/th3khan/restapi-go-fiber-mysql/models"
)

type ProductSerializer struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	SerialNumber string  `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) ProductSerializer {
	return ProductSerializer{
		ID:           productModel.ID,
		Name:         productModel.Name,
		Price:        productModel.Price,
		SerialNumber: productModel.SerialNumber,
	}
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)
	responseProducts := []ProductSerializer{}

	for _, product := range products {
		responseProducts = append(responseProducts, CreateResponseProduct(product))
	}

	return c.Status(fiber.StatusOK).JSON(responseProducts)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)

	return c.Status(fiber.StatusCreated).JSON(responseProduct)
}

func FindProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("Product not found")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var product models.Product

	if err := FindProduct(id, &product); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(CreateResponseProduct(product))
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var product models.Product

	if err := FindProduct(id, &product); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	type UpdateProduct struct {
		Name         string  `json:"name"`
		Price        float64 `json:"price"`
		SerialNumber string  `json:"serial_number"`
	}

	var updateProduct UpdateProduct

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product.Name = updateProduct.Name
	product.Price = updateProduct.Price
	product.SerialNumber = updateProduct.SerialNumber

	database.Database.Db.Save(&product)

	return c.Status(fiber.StatusOK).JSON(CreateResponseProduct(product))
}
