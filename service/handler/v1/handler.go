package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Filter(*fiber.Ctx) error
	Rank(*fiber.Ctx) error
}
