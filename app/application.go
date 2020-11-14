package app

import (
	"github.com/gofiber/fiber"
)

func StartApplication() {
	app := fiber.New()
	mapUrls(app)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
