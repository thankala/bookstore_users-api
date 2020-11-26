package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thankala/bookstore_users-api/logger"
)

func StartApplication() {
	app := fiber.New()
	mapUrls(app)
	logger.Info("About to start the application")
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
