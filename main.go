package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	api := app.Group("/api")
	api.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"users": []string{"john", "doe"}})
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
