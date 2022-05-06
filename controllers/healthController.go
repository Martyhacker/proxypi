package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":"Server is healthy",
		"message":"Tutorial on github.com/Martyhacker/proxypi/README.md",
	})
}