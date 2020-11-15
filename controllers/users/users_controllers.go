package users

import (
	"github.com/gofiber/fiber"
	"github.com/thankala/bookstore_users-api/domain/users"
	"github.com/thankala/bookstore_users-api/services"
	"github.com/thankala/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func Create(c *fiber.Ctx) error {
	var requestBody users.User

	parseErr := c.BodyParser(&requestBody)
	if parseErr != nil {
		err := errors.NewBadRequestError("Invalid JSON Body")
		return c.Status(err.StatusCode).JSON(err)
	}

	result, saveError := services.CreateUser(requestBody)
	if saveError != nil {
		return c.Status(saveError.StatusCode).JSON(saveError)
	}

	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"result": result,
	})
}

func Get(c *fiber.Ctx) error {
	userID, userErr := strconv.ParseUint(c.Params("userId"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid UserID")
		return c.Status(err.StatusCode).JSON(err)
	}
	result, err := services.GetUser(uint(userID))
	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func Update(c *fiber.Ctx) error {
	userID, userErr := strconv.ParseUint(c.Params("userId"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid UserID")
		return c.Status(err.StatusCode).JSON(err)
	}

	var requestBody users.User

	parseErr := c.BodyParser(&requestBody)
	if parseErr != nil {
		err := errors.NewBadRequestError("Invalid JSON Body")
		return c.Status(err.StatusCode).JSON(err)
	}

	requestBody.ID = uint(userID)

	result, saveError := services.UpdateUser(requestBody)
	if saveError != nil {
		return c.Status(saveError.StatusCode).JSON(saveError)
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func Delete(c *fiber.Ctx) error {
	userID, userErr := strconv.ParseUint(c.Params("userId"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid UserID")
		return c.Status(err.StatusCode).JSON(err)
	}
	err := services.DeleteUser(uint(userID))
	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"result": "deleted",
	})
}

func Search(c *fiber.Ctx) error {
	userStatus := c.Query("status")
	users, err := services.Search(userStatus)
	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(users)
}
