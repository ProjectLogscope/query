package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/hardeepnarang10/query/service/middleware/v1/internal/validator/custom"
	"github.com/hardeepnarang10/query/service/middleware/v1/internal/validator/playground"
)

type middleware struct {
	playgroundValidator playground.Playground
	customValidator     custom.Validators
}

func New() Middleware {
	return &middleware{
		playgroundValidator: playground.Playground{
			Validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		customValidator: custom.New(),
	}
}
