package playground

import "log/slog"

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func (e *ErrorResponse) LogValue() slog.Value {
	if e == nil {
		return slog.Value{}
	}
	attrs := []slog.Attr{
		slog.Bool("error", e.Error),
		slog.String("failed_field", e.FailedField),
		slog.String("tag", e.Tag),
		slog.Any("value", e.Value),
	}
	return slog.GroupValue(attrs...)
}
