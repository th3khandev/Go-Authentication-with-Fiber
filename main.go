package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Post("/signup", func(c *fiber.Ctx) error {
		return nil
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/private", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"path":    "private",
		})
	})

	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"path":    "public",
		})
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
