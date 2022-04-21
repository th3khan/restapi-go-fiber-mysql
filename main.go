package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/restapi-go-fiber-mysql/database"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	database.ConnectionDb()
	app := fiber.New()

	app.Get("/", helloWorld)

	log.Fatal(app.Listen(":3000"))
}
