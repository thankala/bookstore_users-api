package app

import (
	"github.com/gofiber/fiber"
	"github.com/thankala/bookstore_users-api/controllers"
)

func mapUrls(app *fiber.App) {
	app.Get("/ping", controllers.Ping)
	app.Get("/users/:userID",controllers.GetUser)
	app.Get("/users/search", controllers.FindUser)
	app.Post("/users",controllers.CreateUser)
}
