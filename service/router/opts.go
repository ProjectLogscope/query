package router

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/hardeepnarang10/query/service/handler/v1"
	"github.com/hardeepnarang10/query/service/healthcheck"
	"github.com/hardeepnarang10/query/service/middleware/v1"
	"github.com/hardeepnarang10/query/service/registry"
)

type opts func(*router) error

type Hooks struct {
	Register func(fiber.Route) error
	Listen   func(fiber.ListenData) error
	Shutdown func() error
}

func WithHooks(hooks Hooks) opts {
	if hooks.Register == nil {
		hooks.Register = func(r fiber.Route) error { return nil }
	}
	if hooks.Listen == nil {
		hooks.Listen = func(r fiber.ListenData) error { return nil }
	}
	if hooks.Shutdown == nil {
		hooks.Shutdown = func() error { return nil }
	}

	return func(r *router) error {
		r.app.Hooks().OnName(hooks.Register)
		r.app.Hooks().OnListen(hooks.Listen)
		r.app.Hooks().OnShutdown(hooks.Shutdown)
		return nil
	}
}

func WithMiddleware(m ...fiber.Handler) opts {
	m = append(m,
		cors.New(cors.Config{
			AllowOrigins:     "*",
			AllowHeaders:     "Origin,X-Requested-With,Content-Type,Accept",
			AllowCredentials: true,
		}),
		favicon.New(),
		recover.New(),
		requestid.New(),
	)

	return func(r *router) error {
		for _, middleware := range m {
			r.app.Use(middleware)
		}
		return nil
	}
}

func WithHandler(ctx context.Context, svc *registry.ServiceRegistry, h handler.Handler) opts {
	if h == nil && svc == nil {
		return func(*router) error {
			return errors.New("server: both handler and service registry cannot be nil")
		}
	}

	if h == nil {
		var err error
		h, err = handler.New(ctx, svc)
		if err != nil {
			return func(*router) error {
				return fmt.Errorf("server: unable to create handler instance: %w", err)
			}
		}
	}

	var mw middleware.Middleware
	if svc.ServiceConfig.EnableValidation {
		mw = middleware.New()
	}

	return func(r *router) error {
		apiV1 := r.app.Group("/api/v1").Name("V1")
		queryV1 := apiV1.Group("/search").Name("Search")

		filterSlice := []func(*fiber.Ctx) error{}
		if svc.ServiceConfig.EnableValidation {
			filterSlice = append(filterSlice, timeout.NewWithContext(mw.Filter, svc.ServiceConfig.RequestTimeout, errors.New("validation timed out")))
		}
		filterSlice = append(filterSlice, timeout.NewWithContext(h.Filter, svc.ServiceConfig.RequestTimeout, errors.New("handler timed out")))
		queryV1.Get("/filter", filterSlice...).Name("Filter")

		rankSlice := []func(*fiber.Ctx) error{}
		if svc.ServiceConfig.EnableValidation {
			rankSlice = append(rankSlice, timeout.NewWithContext(mw.Rank, svc.ServiceConfig.RequestTimeout, errors.New("validation timed out")))
		}
		rankSlice = append(rankSlice, timeout.NewWithContext(h.Rank, svc.ServiceConfig.RequestTimeout, errors.New("handler timed out")))
		queryV1.Get("/rank", rankSlice...).Name("Rank")

		return nil
	}
}

func WithDocumentation() opts {
	return func(r *router) error {
		r.app.Use(swagger.New(swagger.Config{
			Next:     nil,
			BasePath: "/api/",
			FilePath: "/service/swagger.json",
			Path:     "v1",
			Title:    "Query API Docs",
		}))
		return nil
	}
}

func WithRedirect() opts {
	return func(r *router) error {
		r.app.Get("/*", func(c *fiber.Ctx) error { return c.Redirect("/api/v1") })
		return nil
	}
}

func WithMonitoring(prefix string) opts {
	return func(r *router) error {
		r.app.Get(prefix+"/metrics", monitor.New()).Name("Metrics")
		return nil
	}
}

func WithHealthCheck() opts {
	return func(r *router) error {
		hc := healthcheck.New()
		r.app.Get("/healthz", hc.Watch)
		return nil
	}
}
