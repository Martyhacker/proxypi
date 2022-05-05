package main

import (
	"proxypi/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Setup(app)

	app.Listen(":3031")
}