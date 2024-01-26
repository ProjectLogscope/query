package store

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type SearchParams struct {
	Fields  []string
	Queries []types.Query
	Page    int
	Size    int
}

func (s *SearchParams) LogValue() slog.Value {
	if s == nil {
		return slog.Value{}
	}

	attrs := []slog.Attr{}
	for i, field := range s.Fields {
		attrs = append(attrs, slog.String(fmt.Sprintf("field_%02d", i), field))
	}

	for i, query := range s.Queries {
		attrs = append(attrs, slog.String(fmt.Sprintf("query_%02d", i), fmt.Sprintf("%#v", query)))
	}

	attrs = append(attrs, slog.Int("page", s.Page))
	attrs = append(attrs, slog.Int("size", s.Size))

	return slog.GroupValue(attrs...)
}

func (s *store) SearchExecutor(ctx context.Context, sp SearchParams) (JSONRawMessageSlice, error) {
	res, err := s.client.Client().Search().
		Index(s.logIndex).
		Request(
			&search.Request{
				Source_: types.SourceFilter{
					Includes: sp.Fields,
				},
				Query: &types.Query{
					Bool: &types.BoolQuery{
						Must: sp.Queries,
					},
				},
				From: elasticInt((sp.Page - 1) * sp.Size),
				Size: elasticInt(sp.Size),
			},
		).
		Do(ctx)
	if err != nil {
		slog.ErrorContext(ctx,
			"unable to query elasticsearch for document",
			slog.String("index", s.logIndex),
			slog.Any("error", err),
		)
		return nil, fmt.Errorf("unable to query elasticsearch for document: %w", err)
	}

	messageSlice := make([]json.RawMessage, len(res.Hits.Hits))
	for i, hit := range res.Hits.Hits {
		messageSlice[i] = hit.Source_
	}

	return messageSlice, nil
}
