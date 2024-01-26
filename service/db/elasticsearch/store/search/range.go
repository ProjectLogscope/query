package search

import (
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/hardeepnarang10/query/service/common/timestring"
)

func Range(fields map[string]interface{}) []types.Query {
	var queries []types.Query
	tl, tlok := fields["timestampStart"]
	th, thok := fields["timestampEnd"]
	if (tlok && thok) && (tl != nil) && (th != nil) {
		lte := th.(*time.Time).Format(timestring.Layout)
		gte := tl.(*time.Time).Format(timestring.Layout)
		queries = append(queries, types.Query{
			Range: map[string]types.RangeQuery{
				"timestamp": types.DateRangeQuery{
					Lte: &lte,
					Gte: &gte,
				},
			},
		})
	}
	return queries
}
