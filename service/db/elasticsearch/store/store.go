package store

import (
	"context"
)

type Store interface {
	SearchExecutor(ctx context.Context, sp SearchParams) (JSONRawMessageSlice, error)
}
