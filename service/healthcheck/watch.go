package healthcheck

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status string `json:"status"`
}

func (*healthcheck) Watch(c *fiber.Ctx) error {
	return c.JSON(Response{Status: "ok"}, fiber.MIMEApplicationJSON)
}
