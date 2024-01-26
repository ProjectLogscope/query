package router

import (
	"context"

	"github.com/hardeepnarang10/query/service/registry"
)

type Router interface {
	Register(svc *registry.ServiceRegistry, opts ...opts) error
	Listen(addr string) error
	Shutdown(ctx context.Context) error
}
