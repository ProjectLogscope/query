package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/hardeepnarang10/query/service/db/elasticsearch/store"
	"github.com/hardeepnarang10/query/service/registry"
)

type handler struct {
	service *registry.ServiceRegistry
	ess     store.Store
}

func New(ctx context.Context, svc *registry.ServiceRegistry) (Handler, error) {
	if svc == nil {
		return nil, errors.New("handler: empty service registery passed")
	}

	ess, err := store.New(ctx, svc.ElasticsearchClient, svc.ServiceConfig.StoreIndex)
	if err != nil {
		return nil, fmt.Errorf("handler: unable to create elasticsearch store: %w", err)
	}

	return &handler{
		service: svc,
		ess:     ess,
	}, nil
}
