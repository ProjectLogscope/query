package router

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/hardeepnarang10/query/service/registry"
)

func New(name string, enablePrefork bool) Router {
	app := fiber.New(fiber.Config{
		AppName: name,
		Prefork: enablePrefork,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			slog.ErrorContext(c.Context(), "fiber error sink", slog.Any("error", err))
			return err
		},
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	return &router{
		app: app,
	}
}

type router struct {
	app *fiber.App
}

func (r *router) Register(svc *registry.ServiceRegistry, opts ...opts) error {
	if r == nil {
		return errors.New("server: register attempted on uninitialized server instance")
	}

	for _, opt := range opts {
		if err := opt(r); err != nil {
			return err
		}
	}

	return nil
}

func (r *router) Listen(addr string) error {
	if r.app.Listen(addr) != nil {
		return fmt.Errorf("unable to start server: %w", r.app.Listen(addr))
	}
	return nil
}

func (r *router) Shutdown(ctx context.Context) error {
	if err := r.app.ShutdownWithContext(ctx); err != nil {
		return fmt.Errorf("unable to gracefully shutdown server: %w", err)
	}
	return nil
}
