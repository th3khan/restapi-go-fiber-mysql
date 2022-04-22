package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/restapi-go-fiber-mysql/database"
	"github.com/th3khan/restapi-go-fiber-mysql/routes"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	// hello world
	app.Get("/", helloWorld)

	// user
	app.Get("/users", routes.GetUsers)
	app.Get("/users/:id", routes.GetUser)
	app.Post("/users", routes.CreateUser)
	app.Put("/users/:id", routes.UpdateUser)
	app.Delete("/users/:id", routes.DeleteUser)

	// product
	app.Get("/products", routes.GetProducts)
	app.Post("/products", routes.CreateProduct)
	app.Get("/products/:id", routes.GetProduct)
	app.Put("/products/:id", routes.UpdateProduct)
	app.Delete("/products/:id", routes.DeleteProduct)

	// order
	app.Post("/orders", routes.CreateOrder)
	app.Get("/orders", routes.GetOrders)
}

func main() {
	database.ConnectionDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
