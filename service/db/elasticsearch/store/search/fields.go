package search

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/hardeepnarang10/query/service/common/regex"
)

func Fields(fields map[string]interface{}) []types.Query {
	var queries []types.Query
	for key, value := range fields {
		var query types.Query
		if rp, ok := value.(*string); ok && rp != nil {
			if key == "timestampStart" ||
				key == "timestampEnd" {
				continue
			}

			val := strings.TrimSpace(*rp)
			if regex.IsRegex(val) {
				val = val[1 : len(val)-1]
				switch key {
				case "level", "message", "resourceId",
					"traceId", "spanId", "commit",
					"metadata_parentResourceId":
					key += "_phrase"
				}
				query = types.Query{
					Regexp: map[string]types.RegexpQuery{
						key: {
							Value: val,
						},
					},
				}
			} else {
				switch key {
				case "level", "message", "resourceId",
					"traceId", "spanId", "commit",
					"metadata_parentResourceId":
					key += "_term"
				}
				query = types.Query{
					MatchPhrase: map[string]types.MatchPhraseQuery{
						key: {
							Query: val,
						},
					},
				}
			}

			queries = append(queries, query)
		}
	}

	return queries
}
