package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/hardeepnarang10/query/service/registry"
	"github.com/hardeepnarang10/query/service/router"
	"github.com/hardeepnarang10/query/service/server"
)

type serverConfig struct {
	serverName    string
	serverPort    uint16
	enablePrefork bool
}

func initServer(ctx context.Context, svc *registry.ServiceRegistry, sc serverConfig) server.Server {
	if svc == nil {
		report(ctx, errors.New("server: empty service registery passed"))
	}

	hooks := router.Hooks{
		Register: func(r fiber.Route) error {
			slog.Info("Route registered", slog.String("Name", r.Name), slog.String("Method", r.Method))
			return nil
		},
		Listen: func(ld fiber.ListenData) error {
			if fiber.IsChild() {
				return nil
			}
			serverHostPort := fmt.Sprintf("%s://%s:%s", map[bool]string{false: "http", true: "https"}[ld.TLS], ld.Host, ld.Port)
			slog.Info("Starting server", slog.String("HostPort", serverHostPort))
			return nil
		},
		Shutdown: func() error {
			slog.Info("Shutting down server", slog.Int("Port", int(sc.serverPort)))
			return nil
		},
	}

	r := router.New(sc.serverName, sc.enablePrefork)
	if err := r.Register(svc,
		router.WithHandler(ctx, svc, nil),
		router.WithHooks(hooks),
		router.WithMiddleware(),
		router.WithDocumentation(),
		router.WithRedirect(),
		router.WithMonitoring("/service/monitor"),
		router.WithHealthCheck(),
	); err != nil {
		report(ctx, err)
	}

	srv, err := server.New(r)
	if err != nil {
		report(ctx, fmt.Errorf("unable to initialize server instance: %w", err))
	}

	return srv
}
