package message

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/hardeepnarang10/query/service/common/timestring"
)

type Filter struct {
	Level                    *string    `json:"level"`
	Message                  *string    `json:"message"`
	ResourceID               *string    `json:"resourceId"`
	TimestampStart           *time.Time `json:"timestampStart"`
	TimestampEnd             *time.Time `json:"timestampEnd"`
	TraceID                  *string    `json:"traceId"`
	SpanID                   *string    `json:"spanId"`
	Commit                   *string    `json:"commit"`
	MetadataParentResourceID *string    `json:"metadata_parentResourceId"`
}

func (f *Filter) Unmarshal(m map[string]string) error {
	if f == nil {
		return errors.New("parse attempted on uninitialized filter message")
	}

	for k, v := range m {
		switch k {
		case "level":
			f.Level = queryString(v)
		case "message":
			f.Message = queryString(v)
		case "resourceId":
			f.ResourceID = queryString(v)
		case "timestampStart":
			ts, err := time.Parse(timestring.Layout, v)
			if err != nil {
				return fmt.Errorf("unable to parse %q with layout %q: %w", "timestampStart", timestring.Layout, err)
			}
			f.TimestampStart = &ts
		case "timestampEnd":
			te, err := time.Parse(timestring.Layout, v)
			if err != nil {
				return fmt.Errorf("unable to parse %q with layout %q: %w", "timestampEnd", timestring.Layout, err)
			}
			f.TimestampEnd = &te
		case "traceId":
			f.TraceID = queryString(v)
		case "spanId":
			f.SpanID = queryString(v)
		case "commit":
			f.Commit = queryString(v)
		case "metadateParentResourceId":
			f.MetadataParentResourceID = queryString(v)
		default:
			continue
		}
	}
	return nil
}

func (f *Filter) LogValue() slog.Value {
	if f == nil {
		return slog.Value{}
	}
	attrs := []slog.Attr{}

	if f.Level != nil {
		attrs = append(attrs, slog.String("Level", *f.Level))
	}
	if f.Message != nil {
		attrs = append(attrs, slog.String("Message", *f.Message))
	}
	if f.ResourceID != nil {
		attrs = append(attrs, slog.String("ResourceID", *f.ResourceID))
	}
	if f.TimestampStart != nil {
		attrs = append(attrs, slog.Time("Start", *f.TimestampStart))
	}
	if f.TimestampEnd != nil {
		attrs = append(attrs, slog.Time("End", *f.TimestampEnd))
	}
	if f.TraceID != nil {
		attrs = append(attrs, slog.String("TraceID", *f.TraceID))
	}
	if f.SpanID != nil {
		attrs = append(attrs, slog.String("SpanID", *f.SpanID))
	}
	if f.Commit != nil {
		attrs = append(attrs, slog.String("Commit", *f.Commit))
	}
	if f.MetadataParentResourceID != nil {
		attrs = append(attrs, slog.String("ParentResourceID", *f.MetadataParentResourceID))
	}

	return slog.GroupValue(attrs...)
}
