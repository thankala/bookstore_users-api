package app

import (
	"github.com/gofiber/fiber"
	"github.com/thankala/bookstore_users-api/controllers/ping"
	"github.com/thankala/bookstore_users-api/controllers/users"
)

func mapUrls(app *fiber.App) {
	app.Get("/ping", ping.Ping)
	app.Get("/users/:userID", users.GetUser)
	app.Get("/users/search", users.FindUser)
	app.Post("/users", users.CreateUser)
}
