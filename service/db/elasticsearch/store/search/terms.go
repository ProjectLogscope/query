package search

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

var (
	TermFieldsNonTimestamp []string = []string{
		"level_term",
		"message_term",
		"resourceId_term",
		"traceId_term",
		"spanId_term",
		"commit_term",
		"metadata_parentResourceId_term",
	}
)

func Terms(fields []string, query string) []types.Query {
	var queries []types.Query
	if len(query) != 0 {
		queries = append(queries,
			types.Query{
				MultiMatch: &types.MultiMatchQuery{
					Fields: fields,
					Query:  query,
				},
			},
		)
	}
	return queries
}
