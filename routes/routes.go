package routes

import (
	"proxypi/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App){
	app.Post("",controllers.PostUri)
}