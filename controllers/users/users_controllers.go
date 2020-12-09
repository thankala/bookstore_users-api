package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thankala/bookstore_auth-go/auth"
	"github.com/thankala/bookstore_users-api/domain/users"
	"github.com/thankala/bookstore_users-api/services"
	"github.com/thankala/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func Create(c *fiber.Ctx) error {
	var requestBody users.User

	if parseErr := parseUserRequest(c, &requestBody); parseErr != nil {
		return c.Status(parseErr.StatusCode).JSON(parseErr)
	}

	user, saveError := services.UsersService.CreateUser(requestBody)
	if saveError != nil {
		return c.Status(saveError.StatusCode).JSON(saveError)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

func Get(c *fiber.Ctx) error {
	if err := auth.AuthenticateRequest(c.Request()); err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}
	userID, parseErr := parseUserID(c)
	if parseErr != nil {
		return c.Status(parseErr.StatusCode).JSON(parseErr)
	}
	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		return c.Status(getErr.StatusCode).JSON(getErr)
	}
	if auth.GetCallerId(c.Request()) == user.Id {
		return c.Status(http.StatusOK).JSON(user.Marshall(false))

	}
	return c.Status(http.StatusOK).JSON(user.Marshall(auth.IsPublic(c.Request())))
}

func Update(c *fiber.Ctx) error {
	userID, parseErr := parseUserID(c)
	if parseErr != nil {
		return c.Status(parseErr.StatusCode).JSON(parseErr)
	}

	var requestBody users.User

	if parseErr := parseUserRequest(c, &requestBody); parseErr != nil {
		return c.Status(parseErr.StatusCode).JSON(parseErr)
	}

	requestBody.Id = userID

	user, updateErr := services.UsersService.UpdateUser(requestBody)
	if updateErr != nil {
		return c.Status(updateErr.StatusCode).JSON(updateErr)
	}

	return c.Status(http.StatusOK).JSON(user)
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
	usersArray, searchErr := services.UsersService.SearchUser(userStatus)
	if searchErr != nil {
		return c.Status(searchErr.StatusCode).JSON(searchErr)
	}
	return c.Status(http.StatusOK).JSON(usersArray.Marshall(auth.IsPublic(c.Request())))
}

func Login(c *fiber.Ctx) error {
	var requestBody users.LoginRequest

	if parseErr := parseLoginRequest(c, &requestBody); parseErr != nil {
		return c.Status(parseErr.StatusCode).JSON(parseErr)
	}

	user, saveError := services.UsersService.LoginUser(requestBody)
	if saveError != nil {
		return c.Status(saveError.StatusCode).JSON(saveError)
	}

	return c.Status(http.StatusOK).JSON(user.Marshall(c.Get("X-Public") == "true"))
}

func parseUserRequest(c *fiber.Ctx, requestBody *users.User) *errors.RestError {
	parseErr := c.BodyParser(requestBody)
	if parseErr != nil {
		return errors.NewBadRequestError("Invalid JSON Body")
	}
	return nil
}

func parseLoginRequest(c *fiber.Ctx, requestBody *users.LoginRequest) *errors.RestError {
	parseErr := c.BodyParser(requestBody)
	if parseErr != nil {
		return errors.NewBadRequestError("Invalid JSON Body")
	}
	return nil
}

func parseUserID(c *fiber.Ctx) (int64, *errors.RestError) {
	userID, userErr := strconv.ParseInt(c.Params("userId"), 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("Invalid UserID")
	}
	return userID, nil
}
