package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", homeRoute)
	log.Fatal(app.Listen(":4000"))
}

func homeRoute(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Welcome to the world of Go Fiber!!!"})
}
