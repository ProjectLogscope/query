package handler

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/hardeepnarang10/query/service/common/regex"
	"github.com/hardeepnarang10/query/service/handler/v1/internal/jsonmap"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type RequestValidator struct {
	validator *validator.Validate
}

func (v RequestValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	if errs := v.validator.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var e ErrorResponse
			e.FailedField = err.Field()
			e.Tag = err.Tag()
			e.Value = err.Value()
			e.Error = true
			validationErrors = append(validationErrors, e)
		}
	}
	return validationErrors
}

func validateRegexFields(jm jsonmap.JSONMap) error {
	for k, v := range jm {
		switch val := v.(type) {
		case string:
			if !regex.IsRegex(val) {
				continue
			}
			// regexp.CompilePOSIX does not perform semantic pattern validation
			if _, err := regexp.CompilePOSIX(val); err != nil {
				return fmt.Errorf("cannot compile tag %q value %q to regex (POSIX ERE): %w", k, val, err)
			}
		}
	}
	return nil
}
