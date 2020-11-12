package controllers

import "C"
import (
	"github.com/gofiber/fiber"
	"net/http"
)

func CreateUser(c *fiber.Ctx) error{
	return c.Status(http.StatusNotImplemented).SendString("Implement")


}

func GetUser(c *fiber.Ctx) error{
	return c.Status(http.StatusNotImplemented).SendString("Implement")

}

func FindUser(c *fiber.Ctx) error{
	return c.Status(http.StatusNotImplemented).SendString("Implement")
}

