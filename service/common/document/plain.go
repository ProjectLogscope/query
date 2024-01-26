package document

import (
	"log/slog"
	"time"
)

type Metadata struct {
	ParentResourceID *string `json:"parentResourceId,omitempty"`
}

type Plain struct {
	Level      string    `json:"level,omitempty" validate:"required"`
	Message    string    `json:"message,omitempty" validate:"required"`
	ResourceID string    `json:"resourceId,omitempty" validate:"required"`
	Timestamp  time.Time `json:"timestamp,omitempty" validate:"required"`
	TraceID    string    `json:"traceId,omitempty" validate:"required"`
	SpanID     string    `json:"spanId,omitempty" validate:"required"`
	Commit     string    `json:"commit,omitempty" validate:"required"`
	Metadata   *Metadata `json:"metadata,omitempty" validate:"required"`
}

func (p *Plain) Extend() *Extended {
	return &Extended{
		LevelTerm:                      p.Level,
		LevelPhrase:                    p.Level,
		MessageTerm:                    p.Message,
		MessagePhrase:                  p.Message,
		ResourceIDTerm:                 p.ResourceID,
		ResourceIDPhrase:               p.ResourceID,
		Timestamp:                      p.Timestamp,
		TraceIDTerm:                    p.TraceID,
		TraceIDPhrase:                  p.TraceID,
		SpanIDTerm:                     p.SpanID,
		SpanIDPhrase:                   p.SpanID,
		CommitTerm:                     p.Commit,
		CommitPhrase:                   p.Commit,
		MetadataParentResourceIDTerm:   p.Metadata.ParentResourceID,
		MetadataParentResourceIDPhrase: p.Metadata.ParentResourceID,
	}
}

func (p *Plain) LogValue() slog.Value {
	if p == nil {
		return slog.Value{}
	}
	attrs := []slog.Attr{
		slog.String("Level", p.Level),
		slog.String("Message", p.Message),
		slog.String("ResourceID", p.ResourceID),
		slog.Time("Timestamp", p.Timestamp),
		slog.String("TraceID", p.TraceID),
		slog.String("SpanID", p.SpanID),
		slog.String("Commit", p.Commit),
	}
	if p.Metadata.ParentResourceID != nil {
		attrs = append(attrs,
			slog.Any("Metadata",
				slog.GroupValue([]slog.Attr{
					slog.String("ParentResourceID", *p.Metadata.ParentResourceID)}...,
				),
			),
		)
	}
	return slog.GroupValue(attrs...)
}
