package users

import (
	"github.com/gofiber/fiber"
	"github.com/thankala/bookstore_users-api/domain/users"
	"github.com/thankala/bookstore_users-api/services"
	"net/http"
)

func CreateUser(c *fiber.Ctx) error {
	var requestBody users.User
	err := c.BodyParser(&requestBody)
	if err != nil {
		//TODO: Handle Error
	}
	result, saveError := services.CreateUser(requestBody)
	if saveError != nil {
		//TODO: Handle User Creation
	}
	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"result": result,
		"error":  saveError,
	})
}

func GetUser(c *fiber.Ctx) error {
	return c.Status(http.StatusNotImplemented).SendString("Implement")
}

func FindUser(c *fiber.Ctx) error {
	return c.Status(http.StatusNotImplemented).SendString("Implement")
}
