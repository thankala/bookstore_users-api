package controllers

import (
	"github.com/gofiber/fiber"
	"net/http"
)

func Ping(c *fiber.Ctx) error{
	return c.Status(http.StatusOK).SendString("Ping")
}
