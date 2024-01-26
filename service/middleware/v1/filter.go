package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func (m *middleware) Filter(c *fiber.Ctx) error {
	if m.customValidator.Empty(c.Queries()) {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]string{
				"error":       "empty input",
				"description": "atleast one of non-pagination values must be present in search input",
			})
	}
	if err := m.customValidator.TimeRange(c.Queries()); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]string{
				"error":       "bad time input",
				"description": err.Error(),
			})
	}
	if err := m.customValidator.PageRange(c.Queries()); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]string{
				"error":       "invalid pagination value",
				"description": err.Error(),
			})
	}
	return c.Next()
}
