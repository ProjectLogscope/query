package registry

import (
	"time"
)

type ServiceConfig struct {
	RequestTimeout   time.Duration
	StoreIndex       string
	EnableValidation bool
}
