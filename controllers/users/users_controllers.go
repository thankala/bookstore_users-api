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

	if parseErr := parseRequestBody(c, &requestBody); parseErr != nil {
		return c.Status(parseErr.StatusCode).JSON(parseErr)
	}

	result, saveError := services.UsersService.CreateUser(requestBody)
	if saveError != nil {
		return c.Status(saveError.StatusCode).JSON(saveError)
	}

	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"result": result,
	})
}

func Get(c *fiber.Ctx) error {
	userID, parseErr := parseUserID(c)
	if parseErr != nil {
		return c.Status(parseErr.StatusCode).JSON(parseErr)
	}
	user, createErr := services.UsersService.GetUser(userID)
	if createErr != nil {
		return c.Status(createErr.StatusCode).JSON(createErr)
	}
	return c.Status(http.StatusOK).JSON(user.Marshall(string(c.Request().Header.PeekBytes([]byte("X-Public"))) == "true"))
}

func Update(c *fiber.Ctx) error {
	userID, parseErr := parseUserID(c)
	if parseErr != nil {
		return c.Status(parseErr.StatusCode).JSON(parseErr)
	}

	var requestBody users.User

	if parseErr := parseRequestBody(c, &requestBody); parseErr != nil {
		return c.Status(parseErr.StatusCode).JSON(parseErr)
	}

	requestBody.ID = userID

	result, updateErr := services.UsersService.UpdateUser(requestBody)
	if updateErr != nil {
		return c.Status(updateErr.StatusCode).JSON(updateErr)
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func Delete(c *fiber.Ctx) error {
	userID, parseErr := parseUserID(c)
	if parseErr != nil {
		return c.Status(parseErr.StatusCode).JSON(parseErr)
	}
	deleteErr := services.UsersService.DeleteUser(userID)
	if deleteErr != nil {
		return c.Status(deleteErr.StatusCode).JSON(deleteErr)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"result": "deleted",
	})
}

func Search(c *fiber.Ctx) error {
	userStatus := c.Query("status")
	usersArray, searchErr := services.UsersService.Search(userStatus)
	if searchErr != nil {
		return c.Status(searchErr.StatusCode).JSON(searchErr)
	}
	return c.Status(http.StatusOK).JSON(usersArray.Marshall(string(c.Request().Header.PeekBytes([]byte("X-Public"))) == "true"))
}

func parseRequestBody(c *fiber.Ctx, requestBody *users.User) *errors.RestError {
	parseErr := c.BodyParser(requestBody)
	if parseErr != nil {
		return errors.NewBadRequestError("Invalid JSON Body")
	}
	return nil
}

func parseUserID(c *fiber.Ctx) (uint, *errors.RestError) {
	userID, userErr := strconv.ParseUint(c.Params("userId"), 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("Invalid UserID")
	}
	return uint(userID), nil
}
