package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type Middleware interface {
	Filter(*fiber.Ctx) error
	Rank(*fiber.Ctx) error
}
