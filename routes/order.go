package routes

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/restapi-go-fiber-mysql/database"
	"github.com/th3khan/restapi-go-fiber-mysql/models"
)

type OrderSerializer struct {
	ID        uint              `json:"id" gorm:"primary_key"`
	ProductId int               `json:"product_id"`
	Product   ProductSerializer `json:"product"`
	UserId    int               `json:"user_id"`
	User      UserSerializer    `json:"user"`
	CreatedAt time.Time         `json:"created_at"`
}

func CreateOrderResponse(order models.Order, user UserSerializer, product ProductSerializer) OrderSerializer {
	return OrderSerializer{
		ID:        order.ID,
		ProductId: order.ProductId,
		Product:   product,
		UserId:    order.UserId,
		User:      user,
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	order.CreatedAt = time.Now()

	if err := database.Database.Db.Create(&order).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user models.User
	if err := FindUser(order.UserId, &user); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var product models.Product
	if err := FindProduct(order.ProductId, &product); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(CreateOrderResponse(order, CreateResponseUser(user), CreateResponseProduct(product)))
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrder := []OrderSerializer{}

	for _, order := range orders {
		var user models.User
		database.Database.Db.Find(&user, "id = ?", order.UserId)
		var product models.Product
		database.Database.Db.Find(&product, "id = ?", order.ProductId)

		responseOrder = append(responseOrder, CreateOrderResponse(order, CreateResponseUser(user), CreateResponseProduct(product)))
	}

	return c.Status(fiber.StatusOK).JSON(responseOrder)
}

func FindOrder(id int, order *models.Order) error {
	database.Database.Db.Find(order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("Order not found")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var order models.Order
	if err := FindOrder(id, &order); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user models.User
	var product models.Product

	database.Database.Db.First(&user, "id = ?", order.UserId)
	database.Database.Db.First(&product, "id = ?", order.ProductId)

	return c.Status(fiber.StatusOK).JSON(CreateOrderResponse(order, CreateResponseUser(user), CreateResponseProduct(product)))
}
