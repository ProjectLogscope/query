package definition

import "time"

type CommonResponse struct {
	Level      string    `json:"level,omitempty" validate:"required"`
	Message    string    `json:"message,omitempty" validate:"required"`
	ResourceID string    `json:"resourceId,omitempty" validate:"required"`
	Timestamp  time.Time `json:"timestamp,omitempty" validate:"required"`
	TraceID    string    `json:"traceId,omitempty" validate:"required"`
	SpanID     string    `json:"spanId,omitempty" validate:"required"`
	Commit     string    `json:"commit,omitempty" validate:"required"`
	Metadata   *struct {
		ParentResourceID *string `json:"parentResourceId,omitempty"`
	} `json:"metadata,omitempty" validate:"required"`
}
