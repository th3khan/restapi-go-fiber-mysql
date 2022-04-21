package routes

import (
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
