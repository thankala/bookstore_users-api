package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thankala/bookstore_users-api/controllers/ping"
	"github.com/thankala/bookstore_users-api/controllers/users"
)

func mapUrls(app *fiber.App) {
	app.Get("/ping", ping.Ping)

	app.Post("/users", users.Create)
	app.Get("/users/:userId", users.Get)
	app.Patch("/users/:userId", users.Update)
	app.Delete("/users/:userId", users.Delete)
	app.Post("/users/login",users.Login)

	app.Get("/internal/users/search",users.Search)
}
