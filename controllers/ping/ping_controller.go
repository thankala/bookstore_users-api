package ping

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Ping(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("Ping")
}
