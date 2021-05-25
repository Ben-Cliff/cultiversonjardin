package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}

FIX BRANCH: https://stackoverflow.com/questions/32056324/there-is-no-tracking-information-for-the-current-branch


// setup tutorial https://youtu.be/d4Y2DkKbxM0?t=199
//go run main.go