package main

import (
	"os"
	"proxypi/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Setup(app)
	port := os.Getenv("PORT")
	app.Listen(":"+port)
}