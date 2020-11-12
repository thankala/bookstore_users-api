package app

import (
	"github.com/gofiber/fiber"
)

func StartApplication() {
	app := fiber.New()
	mapUrls(app)
	app.Listen(":3000")

}
