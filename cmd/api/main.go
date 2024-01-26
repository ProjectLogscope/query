package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/hardeepnarang10/query/cmd/api/internal/config"
	"github.com/hardeepnarang10/query/common/service"
	"github.com/hardeepnarang10/query/service/registry"
)

func main() {
	cfg := config.LoadConfig()
	service.SetName(cfg.ServiceName)
	initLogger(cfg.ServiceLogFilepath, cfg.ServiceLogLevel.GetLevel(), cfg.ServiceLogAddSource)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	svc := &registry.ServiceRegistry{}
	var elasticsearchClose func(context.Context) error
	svc.AuthorizationClient = initAuthorization(authorizationConfig{
		authorizationAccessDefault: cfg.AuthorizationDefault,
		authorizationAccessLimited: cfg.AuthorizationLimited,
	})
	svc.ElasticsearchClient, elasticsearchClose = initElasticSearch(ctx, elasticsearchConfig{
		endpoints: cfg.ElasticsearchEndpoints,
	})
	svc.ServiceConfig = registry.ServiceConfig{
		RequestTimeout:   cfg.ServiceRequestTimeout,
		StoreIndex:       cfg.ElasticsearchIndex,
		EnableValidation: cfg.ServiceRequestValidate,
	}

	server := initServer(ctx, svc, serverConfig{
		serverName:    service.GetName(),
		serverPort:    cfg.ServerHTTPPort,
		enablePrefork: cfg.ServicePreforkEnable,
	})
	go func() {
		if err := server.Start(cfg.ServerHTTPPort); err != nil {
			report(ctx, err)
		}
	}()
	defer func() {
		graceCtx, cancel := context.WithTimeout(context.Background(), cfg.ServiceGracePeriod)
		defer cancel()
		if err := server.Shutdown(graceCtx); err != nil {
			slog.ErrorContext(graceCtx, "server close error", slog.Any("error", err))
			report(graceCtx, err)
		}
		if err := elasticsearchClose(graceCtx); err != nil {
			slog.ErrorContext(graceCtx, "elasticsearch close error", slog.Any("error", err))
			report(graceCtx, err)
		}
	}()
	<-ctx.Done()
}
