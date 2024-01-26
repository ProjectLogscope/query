package message

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/hardeepnarang10/query/service/common/timestring"
)

type Rank struct {
	Query          *string    `json:"query"`
	TimestampStart *time.Time `json:"timestampStart"`
	TimestampEnd   *time.Time `json:"timestampEnd"`
}

func (r *Rank) Unmarshal(m map[string]string) error {
	if r == nil {
		return errors.New("parse attempted on uninitialized rank message")
	}

	for k, v := range m {
		switch k {
		case "query":
			r.Query = queryString(v)
		case "timestampStart":
			ts, err := time.Parse(timestring.Layout, v)
			if err != nil {
				return fmt.Errorf("unable to parse %q with layout %q: %w", "timestampStart", timestring.Layout, err)
			}
			r.TimestampStart = &ts
		case "timestampEnd":
			te, err := time.Parse(timestring.Layout, v)
			if err != nil {
				return fmt.Errorf("unable to parse %q with layout %q: %w", "timestampEnd", timestring.Layout, err)
			}
			r.TimestampEnd = &te
		default:
			continue
		}
	}
	return nil
}

func (r *Rank) LogValue() slog.Value {
	if r == nil {
		return slog.Value{}
	}
	attrs := []slog.Attr{}

	if r.Query != nil {
		attrs = append(attrs, slog.String("Query", *r.Query))
	}
	if r.TimestampStart != nil {
		attrs = append(attrs, slog.Time("Start", *r.TimestampStart))
	}
	if r.TimestampEnd != nil {
		attrs = append(attrs, slog.Time("End", *r.TimestampEnd))
	}

	return slog.GroupValue(attrs...)
}
