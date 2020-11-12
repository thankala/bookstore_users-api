package users

import "C"
import (
	"github.com/gofiber/fiber"
	"github.com/thankala/bookstore_users-api/domain/users"
	"github.com/thankala/bookstore_users-api/services"
	"github.com/thankala/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *fiber.Ctx) error {
	var requestBody users.User

	err := c.BodyParser(&requestBody)
	if err != nil {
		restError := errors.NewBadRequest("Invalid JSON Body")
		return c.Status(restError.StatusCode).JSON(restError)
	}

	result, saveError := services.CreateUser(requestBody)
	if saveError != nil {
		return c.Status(saveError.StatusCode).JSON(saveError)
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
