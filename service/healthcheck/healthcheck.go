package healthcheck

import "github.com/gofiber/fiber/v2"

type HealthCheck interface {
	Watch(c *fiber.Ctx) error
}
