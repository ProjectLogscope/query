package playground

import (
	"github.com/go-playground/validator/v10"
)

type Playground struct {
	Validator *validator.Validate
}

func (p *Playground) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	if errs := p.Validator.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ErrorResponse{
				Error:       true,
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			})
		}
	}
	return validationErrors
}
