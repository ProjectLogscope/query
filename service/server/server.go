package server

import (
	"context"
)

type Server interface {
	Start(port uint16) error
	Shutdown(ctx context.Context) error
}
