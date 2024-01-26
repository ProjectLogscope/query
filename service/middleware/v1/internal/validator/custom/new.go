package custom

import (
	"github.com/hardeepnarang10/query/service/middleware/v1/internal/validator/custom/empty"
	"github.com/hardeepnarang10/query/service/middleware/v1/internal/validator/custom/pagerange"
	"github.com/hardeepnarang10/query/service/middleware/v1/internal/validator/custom/timerange"
)

type validators struct{}

func New() Validators {
	return &validators{}
}

func (v *validators) Empty(m map[string]string) bool {
	return empty.Validate(m)
}

func (v *validators) TimeRange(m map[string]string) error {
	return timerange.Validate(m)
}

func (v *validators) PageRange(m map[string]string) error {
	return pagerange.Validate(m)
}
