// setup tutorial https://youtu.be/d4Y2DkKbxM0?t=199
//go run main.go 
package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}
