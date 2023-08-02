package main

import "github.com/gofiber/fiber/v2"

func main() {
	
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	app.Get("/panic", func(c *fiber.Ctx) error {
        panic("problem happen")
    })

    app.Listen(":8080")
}