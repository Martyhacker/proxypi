package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":"Server is healthy",
		"message":"You can post your link with 'uri' field to this link. ;)",
	})
}